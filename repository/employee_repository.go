package repository

import (
	"context"
	"database/sql"
	"errors"
	"math/rand"
	"strconv"

	"github.com/megaariii/go-employee-exam/helper"
	"github.com/megaariii/go-employee-exam/model/domain"
)

type EmployeeRepository interface {
	AddEmployee(ctx context.Context, tx *sql.Tx, employee domain.Employee) domain.Employee
	UpdateEmployee(ctx context.Context, tx *sql.Tx, employee domain.Employee) domain.Employee
	DeleteEmployee(ctx context.Context, tx *sql.Tx, employee domain.Employee)
	FindEmployeeById(ctx context.Context, tx *sql.Tx, employeeId string) (domain.Employee, error)
	FindAllEmployees(ctx context.Context, tx *sql.Tx) []domain.Employee
}

type EmployeeRepositoryImpl struct {
}

func NewEmployeeRepository() EmployeeRepository {
	return &EmployeeRepositoryImpl{}
}

func (repository *EmployeeRepositoryImpl) AddEmployee(ctx context.Context, tx *sql.Tx, employee domain.Employee) domain.Employee {
	id := "emp-" + strconv.Itoa(rand.Intn(100000))
	SQL := "insert into employee(id, name, ktp, status) values (?, ?, ?, ?)"
	_, err := tx.ExecContext(ctx, SQL, id, employee.Name, employee.Ktp, employee.Status)
	helper.PanicIfError(err)

	employee.Id = id
	return employee
}

func (repository *EmployeeRepositoryImpl) UpdateEmployee(ctx context.Context, tx *sql.Tx, employee domain.Employee) domain.Employee {
	SQL := "update employee set name = ?, ktp = ?, status = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, employee.Name, employee.Ktp, employee.Status, employee.Id)
	helper.PanicIfError(err)

	return employee
}

func (repository *EmployeeRepositoryImpl) DeleteEmployee(ctx context.Context, tx *sql.Tx, employee domain.Employee) {
	SQL := "delete from employee where id = ?"
	_, err := tx.ExecContext(ctx, SQL, employee.Id)
	helper.PanicIfError(err)
}

func (repository *EmployeeRepositoryImpl) FindEmployeeById(ctx context.Context, tx *sql.Tx, employeeId string) (domain.Employee, error) {
	SQL := "select id, name, ktp, status from employee where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, employeeId)
	helper.PanicIfError(err)
	defer rows.Close()

	employee := domain.Employee{}
	if rows.Next() {
		err := rows.Scan(&employee.Id, &employee.Name, &employee.Ktp, &employee.Status)
		helper.PanicIfError(err)
		return employee, nil
	} else {
		return employee, errors.New("employee is not found")
	}
}

func (repository *EmployeeRepositoryImpl) FindAllEmployees(ctx context.Context, tx *sql.Tx) []domain.Employee {
	SQL := "select id, name, ktp, status from employee"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var employees []domain.Employee
	for rows.Next() {
		employee := domain.Employee{}
		err := rows.Scan(&employee.Id, &employee.Name, &employee.Ktp, &employee.Status)
		helper.PanicIfError(err)
		employees = append(employees, employee)
	}
	return employees
}
