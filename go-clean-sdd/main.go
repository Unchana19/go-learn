package main

import (
	"github.com/Unchana19/go-clean-sdd/adaptors/handler/rest"
	"github.com/Unchana19/go-clean-sdd/adaptors/repositories/memory"
	"github.com/Unchana19/go-clean-sdd/usecases"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	studentRepo := memory.NewStudentMemory()
	studentService := usecases.NewStudentService(studentRepo)
	studentHandler := rest.NewStudentHandler(studentService)

	app.Get("/students", studentHandler.GetAllStudents)

	app.Listen(":8000")
}
