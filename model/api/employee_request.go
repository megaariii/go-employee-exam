package api

type EmployeeCreateRequest struct {
	Name   string `validate:"required" json:"name"`
	Ktp    string `validate:"required" json:"ktp"`
	Status bool   `json:"status"`
}

type EmployeeUpdateRequest struct {
	Id     string `validate:"required" json:"id"`
	Name   string `validate:"required" json:"name"`
	Ktp    string `validate:"required" json:"ktp"`
	Status bool   `json:"status"`
}
