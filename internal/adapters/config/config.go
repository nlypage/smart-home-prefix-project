package config

import "gorm.io/gorm"

type Config struct {
	Database   *gorm.DB
	ListenPort string
}
