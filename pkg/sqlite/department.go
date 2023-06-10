package sqlite

import (
	"context"
	"database/sql"

	"github.com/brittonhayes/staffing"
	"github.com/google/uuid"
	"github.com/uptrace/bun"

	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

type departmentRepository struct {
	db *bun.DB
}

func NewDepartmentRepository(connection string) staffing.DepartmentRepository {

	sqldb, err := sql.Open(sqliteshim.ShimName, connection)
	if err != nil {
		panic(err)
	}
	sqldb.SetMaxOpenConns(1)

	db := bun.NewDB(sqldb, sqlitedialect.New())

	_, err = db.NewCreateTable().Model((*staffing.Department)(nil)).Exec(context.Background())
	if err != nil {
		panic(err)
	}

	return &departmentRepository{
		db: db,
	}
}

func (r *departmentRepository) Close() error {
	return r.db.Close()
}

func (r *departmentRepository) CreateDepartment(ctx context.Context, name string) error {

	id := staffing.DepartmentID(uuid.NewString())

	department := new(staffing.Department)
	department.ID = id
	department.Name = name

	if _, err := r.db.NewInsert().Model(department).Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (r *departmentRepository) DeleteDepartment(ctx context.Context, departmentID staffing.DepartmentID) error {
	return nil
}
