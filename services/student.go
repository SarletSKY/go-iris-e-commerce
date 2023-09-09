package services

import "github.com/SarletSKY/go-iris-e-commerce/repositories"

type StudentService interface {
	ShowStudentName() string
}

type StudentServiceManager struct {
	repo repositories.StudentRepository
}

func NewStudentServiceManager(repo repositories.StudentRepository) StudentService {
	return &StudentServiceManager{repo: repo}
}

func (m *StudentServiceManager) ShowStudentName() string {
	return m.repo.GetStudentName()
}
