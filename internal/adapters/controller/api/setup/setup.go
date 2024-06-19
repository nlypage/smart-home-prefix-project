package setup

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/nlypage/smart-home-prefix-project/cmd/app"
	v1 "github.com/nlypage/smart-home-prefix-project/internal/adapters/controller/api/v1"
)

func Setup(app *app.SmartHomeApp) {
	app.Fiber.Use(cors.New())

	if app.Logging {
		app.Fiber.Use(logger.New())
	}

	app.Fiber.Get("/ping", func(c *fiber.Ctx) error {
		c.Status(fiber.StatusOK)
		return c.JSON("pong")
	})

	// Setup api v1 routes
	apiV1 := app.Fiber.Group("/api/v1")

	// Setup alarm system routes
	alarmSystemHandler := v1.NewAlarmSystemHandler(app)
	alarmSystemHandler.Setup(apiV1)
}
