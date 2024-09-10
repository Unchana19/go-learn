package rest

import (
	"github.com/Unchana19/go-clean-sdd/usecases"
	"github.com/gofiber/fiber/v2"
)

type StudentRestHandler interface {
	GetAllStudents(c *fiber.Ctx) error
}

type studentRestHandler struct {
	studentUseCase usecases.StudentUseCase
}

func NewStudentHandler(studentUseCase usecases.StudentUseCase) StudentRestHandler {
	return &studentRestHandler{
		studentUseCase: studentUseCase,
	}
}

// GetAllStudents implements StudentRestHandler.
func (s *studentRestHandler) GetAllStudents(c *fiber.Ctx) error {
	students := s.studentUseCase.GetAllStudent()

	return c.Status(fiber.StatusOK).JSON(students)
}
