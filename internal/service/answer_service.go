package service

import (
	"edu_v2/internal/repository"
)

type AnswerService struct {
	repo *repository.AnswerRepository
}

func NewAnswerService(repo *repository.AnswerRepository) *AnswerService {
	return &AnswerService{repo: repo}
}

func (s *AnswerService) CreateAnswer(answers []*string, isUpdated *bool, collectionId *string) error {
	return s.repo.CreateAnswer(collectionId, answers, isUpdated)
}

func (s *AnswerService) DeleteAnswer(collectionId *string) error {
	return s.repo.DeleteAnswer(collectionId)
}

func (s *AnswerService) CreateStudentAnswer(collectionId string, answers []*string, code string) error {
	return s.repo.CreateStudentAnswer(collectionId, answers, code)
}
