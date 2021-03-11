package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jperezviloria/student-go/controller"
)

func StudentRouter(app *fiber.App) {
	app.Get("/student/getall", controller.GetAllStudent)
	app.Get("/student/:id", controller.GetStudentById)
	app.Post("/student/save", controller.SaveStudent)
}
