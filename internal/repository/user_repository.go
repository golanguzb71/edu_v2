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
			utils.SendMessage(err.Error(), 6805374430)
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
