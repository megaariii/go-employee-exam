package api

import "time"

type ExamInsertRequest struct {
	Id         string    `json:"id"`
	EmployeId  string    `validate:"required" json:"employee_id"`
	ExamResult int       `validate:"required" json:"exam_result"`
	ExamDate   time.Time `validate:"required" json:"exam_date"`
}
