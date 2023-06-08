package kv

import (
	"context"
	"io/fs"

	"github.com/brittonhayes/staffing"
	"go.etcd.io/bbolt"
)

type employeeRepository struct {
	bucket string
	db     *bbolt.DB
}

func NewEmployeeRepository(filepath string) staffing.EmployeeRepository {
	db, err := bbolt.Open(filepath, fs.ModePerm, nil)
	if err != nil {
		panic(err)
	}
	return &employeeRepository{
		bucket: "employees",
		db:     db,
	}
}

func (r *employeeRepository) Close() error {
	return r.db.Close()
}

func (r *employeeRepository) CreateEmployee(ctx context.Context, name string) error {
	return nil
}

func (r *employeeRepository) DeleteEmployee(ctx context.Context, employeeID staffing.EmployeeID) error {
	return nil
}

func (r *employeeRepository) AssignDepartment(ctx context.Context, departmentID staffing.DepartmentID, employeeID staffing.EmployeeID) error {
	return nil
}

func (r *employeeRepository) UnassignDepartment(ctx context.Context, employeeID staffing.EmployeeID) error {
	return nil
}

func (r *employeeRepository) AssignProject(ctx context.Context, projectID staffing.ProjectID, employeeID staffing.EmployeeID) error {
	return nil
}

func (r *employeeRepository) UnassignProject(ctx context.Context, employeeID staffing.EmployeeID) error {
	return nil
}
