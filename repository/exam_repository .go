package repository

import (
	"context"
	"database/sql"
	"errors"
	"math/rand"
	"strconv"

	"github.com/megaariii/go-employee-exam/helper"
	"github.com/megaariii/go-employee-exam/model/domain"
)

type ExamRepository interface {
	InsertEmployeeToExam(ctx context.Context, tx *sql.Tx, employeeId string)
	InsertEmployeeExamResult(ctx context.Context, tx *sql.Tx, exam domain.Exam) domain.Exam
	FindExamEmployeeByEmployeeId(ctx context.Context, tx *sql.Tx, employeeId string) (domain.Exam, error, bool)
	FindExamEmployeeByExamId(ctx context.Context, tx *sql.Tx, examId string) (domain.Exam, error)
	FindAllExamEmployees(ctx context.Context, tx *sql.Tx) []domain.Exam
	DeleteEmployeeFromExam(ctx context.Context, tx *sql.Tx, employeeId string)
}

type ExamRepositoryImpl struct {
}

func NewExamRepository() ExamRepository {
	return &ExamRepositoryImpl{}
}

func (repository *ExamRepositoryImpl) InsertEmployeeToExam(ctx context.Context, tx *sql.Tx, employeeId string) {
	id := "exam-" + strconv.Itoa(rand.Intn(100000))
	SQL := "insert into exam(id, employee_id) values (?, ?)"
	_, err := tx.ExecContext(ctx, SQL, id, employeeId)
	helper.PanicIfError(err)
}

func (repository *ExamRepositoryImpl) InsertEmployeeExamResult(ctx context.Context, tx *sql.Tx, exam domain.Exam) domain.Exam {
	SQL := "update exam set exam_result = ?, exam_date = ? where employee_id = ?"
	_, err := tx.ExecContext(ctx, SQL, exam.ExamResult, exam.ExamDate, exam.EmployeId)
	helper.PanicIfError(err)

	return exam
}

func (repository *ExamRepositoryImpl) FindExamEmployeeByEmployeeId(ctx context.Context, tx *sql.Tx, employeeId string) (domain.Exam, error, bool) {
	SQL := "select id, employee_id, exam_result, exam_date from exam where employee_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, employeeId)
	helper.PanicIfError(err)
	defer rows.Close()

	exam := domain.Exam{}
	if rows.Next() {
		err := rows.Scan(&exam.Id, &exam.EmployeId, &exam.ExamResult, &exam.ExamDate)
		helper.PanicIfError(err)
		return exam, nil, true
	} else {
		return exam, errors.New("exam is not found"), false
	}
}

func (repository *ExamRepositoryImpl) FindExamEmployeeByExamId(ctx context.Context, tx *sql.Tx, examId string) (domain.Exam, error) {
	SQL := "select id, employee_id, exam_result, exam_date from exam where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, examId)
	helper.PanicIfError(err)
	defer rows.Close()

	exam := domain.Exam{}
	if rows.Next() {
		err := rows.Scan(&exam.Id, &exam.EmployeId, &exam.ExamResult, &exam.ExamDate)
		helper.PanicIfError(err)
		return exam, nil
	} else {
		return exam, errors.New("exam is not found")
	}
}

func (repository *ExamRepositoryImpl) FindAllExamEmployees(ctx context.Context, tx *sql.Tx) []domain.Exam {
	SQL := "select id, employee_id, exam_result, exam_date from exam"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var exams []domain.Exam
	for rows.Next() {
		exam := domain.Exam{}
		err := rows.Scan(&exam.Id, &exam.EmployeId, &exam.ExamResult, &exam.ExamDate)
		helper.PanicIfError(err)
		exams = append(exams, exam)
	}
	return exams
}

func (repository *ExamRepositoryImpl) DeleteEmployeeFromExam(ctx context.Context, tx *sql.Tx, employeeId string) {
	SQL := "delete from exam where employee_id = ?"
	_, err := tx.ExecContext(ctx, SQL, employeeId)
	helper.PanicIfError(err)
}
