package repository

import (
	"database/sql"
	"edu_v2/graph/model"
	"edu_v2/internal/utils"
	"errors"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetStudentsList(page *int, size *int) (*model.PaginatedResult, error) {
	offset := utils.OffSetGenerator(page, size)
	var totalRecords int
	err := r.db.QueryRow(`SELECT COUNT(*) FROM users`).Scan(&totalRecords)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			totalRecords = 0
		} else {
			return nil, err
		}
	}

	totalPages := utils.CalculateTotalPages(totalRecords, size)

	rows, err := r.db.Query(`SELECT id, phone_number, full_name FROM users order by created_at desc LIMIT $1 OFFSET $2`, size, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var students []model.PaginatedItems
	for rows.Next() {
		var student model.Student
		err = rows.Scan(&student.ID, &student.PhoneNumber, &student.FullName)
		if err != nil {
			return nil, err
		}
		students = append(students, &student)
	}

	return &model.PaginatedResult{
		Items:     students,
		TotalPage: &totalPages,
	}, nil
}

func (r *UserRepository) SearchStudent(value string, page, size *int) (*model.PaginatedResult, error) {
	if value == "" {
		return nil, errors.New("value not inserted")
	}

	offset := utils.OffSetGenerator(page, size)

	var query string
	var args []interface{}

	if value[0] == '+' {

		query = `SELECT id, phone_number, full_name FROM users WHERE phone_number LIKE $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`
		args = []interface{}{"%" + value[1:] + "%", *size, offset}
	} else {

		query = `SELECT id, phone_number, full_name FROM users WHERE full_name LIKE $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`
		args = []interface{}{"%" + value + "%", *size, offset}
	}
	countQuery := `SELECT COUNT(*) FROM users WHERE phone_number LIKE $1 OR full_name LIKE $2`
	var totalRecords int
	err := r.db.QueryRow(countQuery, "%"+value+"%", "%"+value+"%").Scan(&totalRecords)
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []model.PaginatedItems
	for rows.Next() {
		var student model.Student
		err = rows.Scan(&student.ID, &student.PhoneNumber, &student.FullName)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	totalPages := utils.CalculateTotalPages(totalRecords, size)

	return &model.PaginatedResult{
		Items:     students,
		TotalPage: &totalPages,
	}, nil
}
