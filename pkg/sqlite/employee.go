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

type employeeRepository struct {
	db *bun.DB
}

func NewEmployeeRepository(connection string) staffing.EmployeeRepository {

	sqldb, err := sql.Open(sqliteshim.ShimName, connection)
	if err != nil {
		panic(err)
	}
	sqldb.SetMaxOpenConns(1)

	db := bun.NewDB(sqldb, sqlitedialect.New())

	_, err = db.NewCreateTable().Model((*staffing.Employee)(nil)).Exec(context.Background())
	if err != nil {
		panic(err)
	}

	return &employeeRepository{
		db: db,
	}
}

func (r *employeeRepository) Close() error {
	return r.db.Close()
}

func (r *employeeRepository) CreateEmployee(ctx context.Context, name string) (*staffing.Employee, error) {

	id := staffing.EmployeeID(uuid.NewString())

	employee := new(staffing.Employee)
	employee.ID = id
	employee.Name = name

	if _, err := r.db.NewInsert().Model(employee).Exec(ctx); err != nil {
		return nil, err
	}

	return employee, nil
}

func (r *employeeRepository) DeleteEmployee(ctx context.Context, employeeID staffing.EmployeeID) error {
	return nil
}

func (r *employeeRepository) AssignDepartment(ctx context.Context, employeeID staffing.EmployeeID, departmentID staffing.DepartmentID) error {
	return nil
}

func (r *employeeRepository) UnassignDepartment(ctx context.Context, employeeID staffing.EmployeeID) error {
	return nil
}

func (r *employeeRepository) AssignProject(ctx context.Context, employeeID staffing.EmployeeID, projectID staffing.ProjectID) error {
	return nil
}

func (r *employeeRepository) UnassignProject(ctx context.Context, employeeID staffing.EmployeeID) error {
	return nil
}

func (r *employeeRepository) InsertFeedback(ctx context.Context, feedback *staffing.Feedback) error {
	return nil
}
