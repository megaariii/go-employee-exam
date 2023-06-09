package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/megaariii/go-employee-exam/app"
	"github.com/megaariii/go-employee-exam/controller"
	"github.com/megaariii/go-employee-exam/exception"
	"github.com/megaariii/go-employee-exam/helper"
	"github.com/megaariii/go-employee-exam/repository"
	"github.com/megaariii/go-employee-exam/service"
)

func main() {
	router := httprouter.New()
	validate := validator.New()
	db := app.AppDB()

	examRepository := repository.NewExamRepository()
	examService := service.NewExamService(examRepository, db, validate)
	examController := controller.NewExamController(examService)

	employeeRepository := repository.NewEmployeeRepository()
	employeeService := service.NewEmployeeService(employeeRepository, examRepository, db, validate)
	employeeController := controller.NewEmployeeController(employeeService)

	// Employee
	router.POST("/api/employee", employeeController.AddEmployee)
	router.GET("/api/employee/:employeeId", employeeController.FindEmployeeById)
	router.GET("/api/employees", employeeController.FindAllEmployees)
	router.PUT("/api/employee/:employeeId", employeeController.UpdateEmployee)
	router.DELETE("/api/employee/:employeeId", employeeController.DeleteEmployee)

	// Exam
	router.POST("/api/exam/:employeeId", examController.InsertEmployeeExamResult)
	router.GET("/api/exam/:examId", examController.FindExamEmployeeByExamId)
	router.GET("/api/exam", examController.FindAllExamEmployees)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
