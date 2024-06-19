package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nlypage/smart-home-prefix-project/cmd/app"
	handlersDto "github.com/nlypage/smart-home-prefix-project/internal/adapters/controller/api/dto"
	smartHomeValidator "github.com/nlypage/smart-home-prefix-project/internal/adapters/controller/api/validator"
	"github.com/nlypage/smart-home-prefix-project/internal/adapters/database/sqlite"
	"github.com/nlypage/smart-home-prefix-project/internal/domain/dto"
	"github.com/nlypage/smart-home-prefix-project/internal/domain/entities"
	"github.com/nlypage/smart-home-prefix-project/internal/domain/services"
)

// AlarmSystemService is an interface that contains methods to interact with the database.
type AlarmSystemService interface {
	Create(createAlarmSystem *dto.CreateAlarmSystem) (*entities.AlarmSystem, error)
	GetByUUID(uuid string) (*entities.AlarmSystem, error)
	GetAll(limit, offset int) ([]*entities.AlarmSystem, error)
	Update(updateAlarmSystem *dto.UpdateAlarmSystem) (*entities.AlarmSystem, error)
	Delete(uuid string) error
}

// AlarmSystemHandler is a struct that contains a pointer to a AlarmSystemService instance.
type AlarmSystemHandler struct {
	alarmSystemService AlarmSystemService
	validator          *smartHomeValidator.Validator
}

// NewAlarmSystemHandler is a function that returns a new instance of alarmSystemHandler.
func NewAlarmSystemHandler(smartHomeApp *app.SmartHomeApp) *AlarmSystemHandler {
	alarmSystemStorage := sqlite.NewAlarmSystemStorage(smartHomeApp.Db)
	alarmSystemService := services.NewAlarmSystemService(alarmSystemStorage)

	return &AlarmSystemHandler{
		alarmSystemService: alarmSystemService,
		validator:          smartHomeApp.Validator,
	}
}

// Create is handler for creating alarm system.
func (h AlarmSystemHandler) Create(c *fiber.Ctx) error {
	var createAlarmSystem dto.CreateAlarmSystem
	if err := c.BodyParser(&createAlarmSystem); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	errValidate := h.validator.ValidateData(createAlarmSystem)
	if errValidate != nil {
		return errValidate
	}

	alarmSystem, err := h.alarmSystemService.Create(&createAlarmSystem)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": true,
		"body":   alarmSystem,
	})
}

// GetByUUID is handler for getting alarm system by UUID.
func (h AlarmSystemHandler) GetByUUID(c *fiber.Ctx) error {
	var getByUUID handlersDto.UUID
	getByUUID.UUID = c.Params("uuid")

	errValidate := h.validator.ValidateData(getByUUID)
	if errValidate != nil {
		return errValidate
	}

	alarmSystem, err := h.alarmSystemService.GetByUUID(getByUUID.UUID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": true,
		"body":   alarmSystem,
	})
}

// GetAll is handler for getting all alarm systems.
func (h AlarmSystemHandler) GetAll(c *fiber.Ctx) error {
	limit, offset := h.validator.GetLimitAndOffset(c)

	alarmSystems, err := h.alarmSystemService.GetAll(limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": true,
		"body":   alarmSystems,
	})
}

// Update is handler for updating alarm system.
func (h AlarmSystemHandler) Update(c *fiber.Ctx) error {
	var updateAlarmSystem dto.UpdateAlarmSystem
	if err := c.BodyParser(&updateAlarmSystem); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	errValidate := h.validator.ValidateData(updateAlarmSystem)
	if errValidate != nil {
		return errValidate
	}

	alarmSystem, err := h.alarmSystemService.Update(&updateAlarmSystem)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": true,
		"body":   alarmSystem,
	})
}

// Delete is handler for deleting alarm system.
func (h AlarmSystemHandler) Delete(c *fiber.Ctx) error {
	var deleteAlarmSystem handlersDto.UUID
	deleteAlarmSystem.UUID = c.Params("uuid")

	errValidate := h.validator.ValidateData(deleteAlarmSystem)
	if errValidate != nil {
		return errValidate
	}

	err := h.alarmSystemService.Delete(deleteAlarmSystem.UUID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": true,
	})
}

// Setup is a function that registers all routes for the alarm system.
func (h AlarmSystemHandler) Setup(router fiber.Router) {
	alarmSystemGroup := router.Group("/alarm-system")
	alarmSystemGroup.Post("/", h.Create)
	alarmSystemGroup.Get("/:uuid", h.GetByUUID)
	alarmSystemGroup.Get("/", h.GetAll)
	alarmSystemGroup.Put("/", h.Update)
	alarmSystemGroup.Delete("/:uuid", h.Delete)
}
