package service

import (
	"edu_v2/graph/model"
	"edu_v2/internal/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetStudentsList(page *int, size *int) ([]*model.Student, error) {
	return s.userRepo.GetStudentsList(page, size)
}
