package api

type EmployeeResponse struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Ktp    string `json:"ktp"`
	Status bool   `json:"status"`
	IsExam bool   `json:"is_exam"`
}
