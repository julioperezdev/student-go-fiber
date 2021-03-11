package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jperezviloria/student-go/router"
)

func main() {

	app := fiber.New()

	//middleware
	app.Use(logger.New())

	//routes
	router.StudentRouter(app)

	app.Listen(":3001")
}
