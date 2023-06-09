package service

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/megaariii/go-employee-exam/exception"
	"github.com/megaariii/go-employee-exam/helper"
	"github.com/megaariii/go-employee-exam/model/api"
	"github.com/megaariii/go-employee-exam/model/domain"
	"github.com/megaariii/go-employee-exam/repository"
)

type EmployeeService interface {
	AddEmployee(ctx context.Context, request api.EmployeeCreateRequest) api.EmployeeResponse
	UpdateEmployee(ctx context.Context, request api.EmployeeUpdateRequest) api.EmployeeResponse
	DeleteEmployee(ctx context.Context, employeeId string)
	FindEmployeeById(ctx context.Context, employeeId string) api.EmployeeResponse
	FindAllEmployees(ctx context.Context) []api.EmployeeResponse
}

type EmployeeServiceImpl struct {
	EmployeeRepository repository.EmployeeRepository
	ExamRepository     repository.ExamRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewEmployeeService(employeeRepository repository.EmployeeRepository, ExamRepository repository.ExamRepository, db *sql.DB, validate *validator.Validate) EmployeeService {
	return &EmployeeServiceImpl{
		EmployeeRepository: employeeRepository,
		ExamRepository:     ExamRepository,
		DB:                 db,
		Validate:           validate,
	}
}

func (service *EmployeeServiceImpl) AddEmployee(ctx context.Context, request api.EmployeeCreateRequest) api.EmployeeResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	employee := domain.Employee{
		Name:   request.Name,
		Ktp:    request.Ktp,
		Status: request.Status,
	}

	employee = service.EmployeeRepository.AddEmployee(ctx, tx, employee)
	if employee.Status {
		service.ExamRepository.InsertEmployeeToExam(ctx, tx, employee.Id)
	}

	return helper.ConvertDomainToResponse(employee)
}

func (service *EmployeeServiceImpl) UpdateEmployee(ctx context.Context, request api.EmployeeUpdateRequest) api.EmployeeResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	employee, err := service.EmployeeRepository.FindEmployeeById(ctx, tx, request.Id)
	_, _, examFound := service.ExamRepository.FindExamEmployeeByEmployeeId(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	employee.Name = request.Name
	employee.Ktp = request.Ktp
	employee.Status = request.Status

	employee = service.EmployeeRepository.UpdateEmployee(ctx, tx, employee)
	if employee.Status && !examFound {
		service.ExamRepository.InsertEmployeeToExam(ctx, tx, employee.Id)
	} else if !employee.Status {
		service.ExamRepository.DeleteEmployeeFromExam(ctx, tx, employee.Id)
	}

	return helper.ConvertDomainToResponse(employee)
}

func (service *EmployeeServiceImpl) DeleteEmployee(ctx context.Context, employeeId string) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	employee, err := service.EmployeeRepository.FindEmployeeById(ctx, tx, employeeId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.EmployeeRepository.DeleteEmployee(ctx, tx, employee)
}

func (service *EmployeeServiceImpl) FindEmployeeById(ctx context.Context, employeeId string) api.EmployeeResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	employee, err := service.EmployeeRepository.FindEmployeeById(ctx, tx, employeeId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ConvertDomainToResponse(employee)
}

func (service *EmployeeServiceImpl) FindAllEmployees(ctx context.Context) []api.EmployeeResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	employees := service.EmployeeRepository.FindAllEmployees(ctx, tx)

	var employeesResponse []api.EmployeeResponse
	for _, employee := range employees {
		employeesResponse = append(employeesResponse, helper.ConvertDomainToResponse(employee))
	}

	return employeesResponse
}
