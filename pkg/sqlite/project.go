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

type projectRepository struct {
	db *bun.DB
}

func NewProjectRepository(connection string, inmem bool) staffing.ProjectRepository {
	sqldb, err := sql.Open(sqliteshim.ShimName, connection)
	if err != nil {
		panic(err)
	}
	if inmem {
		sqldb.SetMaxOpenConns(1)
		sqldb.SetMaxIdleConns(1000)
		sqldb.SetConnMaxLifetime(0)
	}

	db := bun.NewDB(sqldb, sqlitedialect.New())

	_, err = db.NewCreateTable().Model((*staffing.Project)(nil)).Exec(context.Background())
	if err != nil {
		panic(err)
	}

	return &projectRepository{
		db: db,
	}
}

func (r *projectRepository) Close() error {
	return r.db.Close()
}

func (r *projectRepository) CreateProject(ctx context.Context, name string) (*staffing.Project, error) {

	project := new(staffing.Project)
	project.ID = staffing.ProjectID(uuid.NewString())
	project.Name = name

	if _, err := r.db.NewInsert().Model(project).Exec(ctx); err != nil {
		return nil, err
	}

	return project, nil
}

func (r *projectRepository) CancelProject(ctx context.Context, projectID staffing.ProjectID) (*staffing.Project, error) {
	return nil, nil
}
