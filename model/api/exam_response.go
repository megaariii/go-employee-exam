package api

import "time"

type ExamResponse struct {
	Id         string    `json:"id"`
	EmployeeId string    `json:"employee_id"`
	ExamResult int       `json:"exam_result"`
	ExamDate   time.Time `json:"exam_date"`
}
