package service

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/megaariii/go-employee-exam/exception"
	"github.com/megaariii/go-employee-exam/helper"
	"github.com/megaariii/go-employee-exam/model/api"
	"github.com/megaariii/go-employee-exam/repository"
)

type ExamService interface {
	InsertEmployeeExamResult(ctx context.Context, request api.ExamInsertRequest) api.ExamInsertRequest
	FindExamEmployeeByExamId(ctx context.Context, examId string) api.ExamResponse
	FindAllExamEmployees(ctx context.Context) []api.ExamResponse
}

type ExamServiceImpl struct {
	ExamRepository repository.ExamRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewExamService(ExamRepository repository.ExamRepository, db *sql.DB, validate *validator.Validate) ExamService {
	return &ExamServiceImpl{
		ExamRepository: ExamRepository,
		DB:             db,
		Validate:       validate,
	}
}

func (service *ExamServiceImpl) InsertEmployeeExamResult(ctx context.Context, request api.ExamInsertRequest) api.ExamInsertRequest {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	exam, err, _ := service.ExamRepository.FindExamEmployeeByEmployeeId(ctx, tx, request.EmployeId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	exam.ExamResult = request.ExamResult
	exam.ExamDate = request.ExamDate

	exam = service.ExamRepository.InsertEmployeeExamResult(ctx, tx, exam)

	return helper.ConvertDomainToResponseExam(exam)
}

func (service *ExamServiceImpl) FindExamEmployeeByExamId(ctx context.Context, examId string) api.ExamResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	exam, err := service.ExamRepository.FindExamEmployeeByExamId(ctx, tx, examId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ConvertDomainToResponseExamFind(exam)
}

func (service *ExamServiceImpl) FindAllExamEmployees(ctx context.Context) []api.ExamResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	exams := service.ExamRepository.FindAllExamEmployees(ctx, tx)

	var examsResponse []api.ExamResponse
	for _, exam := range exams {
		examsResponse = append(examsResponse, helper.ConvertDomainToResponseExamFind(exam))
	}

	return examsResponse
}
