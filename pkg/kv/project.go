package kv

import (
	"context"
	"encoding/json"
	"io/fs"

	"github.com/brittonhayes/staffing"
	"github.com/google/uuid"
	"go.etcd.io/bbolt"
)

type projectRepository struct {
	bucket string
	db     *bbolt.DB
}

func NewProjectRepository(filepath string) staffing.ProjectRepository {
	db, err := bbolt.Open(filepath, fs.ModePerm, nil)
	if err != nil {
		panic(err)
	}
	return &projectRepository{
		bucket: "projects",
		db:     db,
	}
}

func (r *projectRepository) Close() error {
	return r.db.Close()
}

func (r *projectRepository) CreateProject(ctx context.Context, name string) error {
	id := uuid.NewString()
	project := &staffing.Project{
		ID:   staffing.ProjectID(id),
		Name: name,
	}

	b, err := json.Marshal(project)
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

func (r *projectRepository) DeleteProject(ctx context.Context, projectID staffing.ProjectID) error {
	return r.db.Update(func(tx *bbolt.Tx) error {
		return tx.DeleteBucket([]byte(projectID))
	})
}
