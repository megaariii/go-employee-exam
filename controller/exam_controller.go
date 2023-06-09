package controller

import (
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/megaariii/go-employee-exam/helper"
	"github.com/megaariii/go-employee-exam/model/api"
	"github.com/megaariii/go-employee-exam/service"
)

type ExamController interface {
	InsertEmployeeExamResult(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindExamEmployeeByExamId(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAllExamEmployees(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type ExamControllerImpl struct {
	examService service.ExamService
}

func NewExamController(examService service.ExamService) ExamController {
	return &ExamControllerImpl{
		examService: examService,
	}
}

func (controller *ExamControllerImpl) InsertEmployeeExamResult(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	examInsertRequest := api.ExamInsertRequest{}
	helper.DecodeRequestBody(request, &examInsertRequest)

	examInsertRequest.EmployeId = params.ByName("employeeId")
	examInsertRequest.ExamDate = time.Now()

	examResponse := controller.examService.InsertEmployeeExamResult(request.Context(), examInsertRequest)
	response := api.APIResponse{
		Code:   200,
		Status: "OK",
		Data:   examResponse,
	}

	helper.EncodeResponse(writer, response)
}

func (controller *ExamControllerImpl) FindExamEmployeeByExamId(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	examId := params.ByName("examId")

	examResponse := controller.examService.FindExamEmployeeByExamId(request.Context(), examId)
	response := api.APIResponse{
		Code:   200,
		Status: "OK",
		Data:   examResponse,
	}

	helper.EncodeResponse(writer, response)
}

func (controller *ExamControllerImpl) FindAllExamEmployees(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	examResponses := controller.examService.FindAllExamEmployees(request.Context())
	response := api.APIResponse{
		Code:   200,
		Status: "OK",
		Data:   examResponses,
	}

	helper.EncodeResponse(writer, response)
}
