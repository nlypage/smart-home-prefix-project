package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nlypage/smart-home-prefix-project/internal/adapters/config"
	smartHomeValidator "github.com/nlypage/smart-home-prefix-project/internal/adapters/controller/api/validator"
	"gorm.io/gorm"
	"time"
)

// SmartHomeApp app is a struct that contains the fiber app, database connection, listen port and timezone.
type SmartHomeApp struct {
	Fiber      *fiber.App
	Db         *gorm.DB
	listenPort string
	Timezone   time.Location
	Validator  *smartHomeValidator.Validator
	Logging    bool
}

// NewSmartHomeApp New is a function that creates a new app struct
func NewSmartHomeApp(config *config.Config) *SmartHomeApp {
	fiberApp := fiber.New(fiber.Config{
		// Global custom error handler
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusBadRequest).JSON(smartHomeValidator.GlobalErrorHandlerResp{
				Success: false,
				Message: err.Error(),
			})
		},
	},
	)

	return &SmartHomeApp{
		Fiber:      fiberApp,
		Db:         config.Database,
		listenPort: config.ListenPort,
		Timezone:   config.Timezone,
		Validator:  smartHomeValidator.New(),
		Logging:    config.Logging,
	}
}

// Start is a function that starts the app
func (a *SmartHomeApp) Start() {
	if err := a.Fiber.Listen(":" + a.listenPort); err != nil {
		panic(err)
	}
}
