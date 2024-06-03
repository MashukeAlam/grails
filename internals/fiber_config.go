package internals

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func FiberAppStart(app *fiber.App) *fiber.App {
	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	// Setup static files
	app.Static("/js", "./static/public/js")
	app.Static("/img", "./static/public/img")
	app.Static("/css", "./static/public/css")
	return app
}
