package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/megaariii/go-employee-exam/helper"
	"github.com/megaariii/go-employee-exam/model/api"
	"github.com/megaariii/go-employee-exam/service"
)

type EmployeeController interface {
	AddEmployee(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateEmployee(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteEmployee(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindEmployeeById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAllEmployees(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type EmployeeControllerImpl struct {
	EmployeeService service.EmployeeService
}

func NewEmployeeController(employeeService service.EmployeeService) EmployeeController {
	return &EmployeeControllerImpl{
		EmployeeService: employeeService,
	}
}

func (controller *EmployeeControllerImpl) AddEmployee(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	employeeCreateRequest := api.EmployeeCreateRequest{}
	helper.DecodeRequestBody(request, &employeeCreateRequest)

	employeeResponse := controller.EmployeeService.AddEmployee(request.Context(), employeeCreateRequest)
	response := api.APIResponse{
		Code:   200,
		Status: "OK",
		Data:   employeeResponse,
	}

	helper.EncodeResponse(writer, response)
}

func (controller *EmployeeControllerImpl) UpdateEmployee(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	employeeUpdateRequest := api.EmployeeUpdateRequest{}
	helper.DecodeRequestBody(request, &employeeUpdateRequest)

	employeeUpdateRequest.Id = params.ByName("employeeId")

	employeeResponse := controller.EmployeeService.UpdateEmployee(request.Context(), employeeUpdateRequest)
	response := api.APIResponse{
		Code:   200,
		Status: "OK",
		Data:   employeeResponse,
	}

	helper.EncodeResponse(writer, response)
}

func (controller *EmployeeControllerImpl) DeleteEmployee(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	employeeId := params.ByName("employeeId")

	controller.EmployeeService.DeleteEmployee(request.Context(), employeeId)
	response := api.APIResponse{
		Code:   200,
		Status: "OK",
	}

	helper.EncodeResponse(writer, response)
}

func (controller *EmployeeControllerImpl) FindEmployeeById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	employeeId := params.ByName("employeeId")

	employeeResponse := controller.EmployeeService.FindEmployeeById(request.Context(), employeeId)
	response := api.APIResponse{
		Code:   200,
		Status: "OK",
		Data:   employeeResponse,
	}

	helper.EncodeResponse(writer, response)
}

func (controller *EmployeeControllerImpl) FindAllEmployees(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	employeeResponses := controller.EmployeeService.FindAllEmployees(request.Context())
	response := api.APIResponse{
		Code:   200,
		Status: "OK",
		Data:   employeeResponses,
	}

	helper.EncodeResponse(writer, response)
}
