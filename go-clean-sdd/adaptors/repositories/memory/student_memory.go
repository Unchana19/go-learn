package memory

import (
	"github.com/Unchana19/go-clean-sdd/entities"
	"github.com/Unchana19/go-clean-sdd/repositories"
)

type studentMemory struct {
	students []entities.Student
}

func NewStudentMemory() repositories.StudentRepository {
	return &studentMemory{
		students: []entities.Student{
			{
				ID: "6510405865",
				Name: "Oanchana Changcharoen",
			},
		},
	}
}

// GetAll implements repositories.StudentRepository.
func (s *studentMemory) GetAll() []entities.Student {
	return s.students
}
