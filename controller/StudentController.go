package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jperezviloria/student-go/model"
	"github.com/jperezviloria/student-go/repository"
	"strconv"
)

func GetAllStudent(ctx *fiber.Ctx) error {
	responseStudents := repository.GetAllStudents()
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": responseStudents,
	})
}

func GetStudentById(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	idStudent, err := strconv.Atoi(idParam)
	if err != nil {
		return err
	}
	responseStudent := repository.GetStudentById(idStudent)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": responseStudent,
	})
}

func SaveStudent(ctx *fiber.Ctx) error {
	student := new(model.Student)
	requestStudent := ctx.BodyParser(student)
	if requestStudent != nil {
		return requestStudent
	}
	responseStudent := repository.SaveStudent(*student)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": responseStudent,
	})
}
