package domain

import "time"

type Exam struct {
	Id         string
	EmployeId  string
	ExamResult int
	ExamDate   time.Time
}
