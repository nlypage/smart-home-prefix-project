package config

import (
	"github.com/glebarez/sqlite"
	"github.com/nlypage/smart-home-prefix-project/internal/domain/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Config struct {
	Database   *gorm.DB
	ListenPort string
	Timezone   time.Location
	Logging    bool
}

func GetConfig() *Config {
	var appLogging bool
	logging := os.Getenv("LOGGING")
	switch logging {
	case "true":
		appLogging = true
	case "false":
		appLogging = false
	default:
		appLogging = false
	}

	var gormConfig *gorm.Config
	if appLogging {
		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: time.Second,
				LogLevel:      logger.Info,
				Colorful:      true,
			},
		)
		gormConfig = &gorm.Config{
			Logger: newLogger,
		}
	}

	database, err := gorm.Open(sqlite.Open("SmartHome.db"), gormConfig)
	if err != nil {
		panic(err)
	} else {
		log.Println("Успешно подключились к базе данных")
	}
	err = database.AutoMigrate(
		&entities.AlarmSystem{},
		&entities.Bulb{},
		&entities.Conditioner{},
		&entities.HeatingControl{},
	)
	if err != nil {
		log.Panic(err)
	}

	port := os.Getenv("LISTEN_PORT")

	return &Config{
		Database:   database,
		ListenPort: port,
		Logging:    appLogging,
	}
}
