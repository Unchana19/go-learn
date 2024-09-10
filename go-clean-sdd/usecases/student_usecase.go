package usecases

import (
	"github.com/Unchana19/go-clean-sdd/entities"
	"github.com/Unchana19/go-clean-sdd/repositories"
)

type StudentUseCase interface {
	GetAllStudent() []entities.Student
}

type studentService struct {
	studentRepo repositories.StudentRepository
}

func NewStudentService(studentRepo repositories.StudentRepository) StudentUseCase {
	return &studentService{
		studentRepo: studentRepo,
	}
}

// GetAllStudent implements StudentUseCase.
func (s *studentService) GetAllStudent() []entities.Student {
	return s.studentRepo.GetAll()
}
