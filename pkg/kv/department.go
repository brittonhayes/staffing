package kv

import (
	"context"
	"encoding/json"
	"io/fs"

	"github.com/brittonhayes/staffing"
	"github.com/google/uuid"
	"go.etcd.io/bbolt"
)

type departmentRepository struct {
	bucket string
	db     *bbolt.DB
}

func NewDepartmentRepository(filepath string) staffing.DepartmentRepository {
	db, err := bbolt.Open(filepath, fs.ModePerm, nil)
	if err != nil {
		panic(err)
	}
	return &departmentRepository{
		bucket: "departments",
		db:     db,
	}
}

func (r *departmentRepository) Close() error {
	return r.db.Close()
}

func (r *departmentRepository) CreateDepartment(ctx context.Context, name string) error {
	id := uuid.NewString()
	department := &staffing.Department{
		ID:   staffing.DepartmentID(id),
		Name: name,
	}

	b, err := json.Marshal(department)
	if err != nil {
		return err
	}

	return r.db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(r.bucket))
		if err != nil {
			return err
		}

		return bucket.Put([]byte(id), b)
	})
}

func (r *departmentRepository) DeleteDepartment(ctx context.Context, departmentID staffing.DepartmentID) error {
	return r.db.Update(func(tx *bbolt.Tx) error {
		return tx.DeleteBucket([]byte(departmentID))
	})
}
