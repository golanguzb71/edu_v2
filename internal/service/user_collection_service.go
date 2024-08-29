package service

import (
	"edu_v2/graph/model"
	"edu_v2/internal/repository"
)

type UserCollectionService struct {
	userCollectionRepo *repository.UserCollectionRepository
}

func NewUserCollectionUserService(UCR *repository.UserCollectionRepository) *UserCollectionService {
	return &UserCollectionService{userCollectionRepo: UCR}
}

func (s *UserCollectionService) GetStudentTestExams(code *string, studentId *string, page *int, size *int) (*model.PaginatedResult, error) {
	return s.userCollectionRepo.GetStudentTestExams(code, studentId, page, size)
}
