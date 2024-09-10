package repositories

import "github.com/Unchana19/go-clean-sdd/entities"

type StudentRepository interface {
	GetAll() []entities.Student
}
