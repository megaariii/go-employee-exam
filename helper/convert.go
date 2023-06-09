package helper

import (
	"github.com/megaariii/go-employee-exam/model/api"
	"github.com/megaariii/go-employee-exam/model/domain"
)

func ConvertDomainToResponse(employee domain.Employee) api.EmployeeResponse {
	return api.EmployeeResponse{
		Id:     employee.Id,
		Name:   employee.Name,
		Ktp:    employee.Ktp,
		Status: employee.Status,
		IsExam: employee.Status,
	}
}

func ConvertDomainToResponseExam(exam domain.Exam) api.ExamInsertRequest {
	return api.ExamInsertRequest{
		Id:         exam.Id,
		EmployeId:  exam.EmployeId,
		ExamResult: exam.ExamResult,
		ExamDate:   exam.ExamDate,
	}
}

func ConvertDomainToResponseExamFind(exam domain.Exam) api.ExamResponse {
	return api.ExamResponse{
		Id:         exam.Id,
		EmployeeId: exam.EmployeId,
		ExamResult: exam.ExamResult,
		ExamDate:   exam.ExamDate,
	}
}
