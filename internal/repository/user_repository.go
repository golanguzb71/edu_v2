package repository

import (
	"database/sql"
	"edu_v2/graph/model"
	"edu_v2/internal/utils"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetStudentsList(page *int, size *int) ([]*model.Student, error) {
	offset := utils.OffSetGenerator(page, size)
	rows, err := r.db.Query(`SELECT id, phone_number, full_name FROM users order by created_at desc LIMIT $1 OFFSET $2`, size, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var students []*model.Student
	for rows.Next() {
		var student model.Student
		err = rows.Scan(&student.ID, &student.PhoneNumber, &student.FullName)
		if err != nil {
			return nil, err
		}
		students = append(students, &student)
	}
	return students, nil
}
