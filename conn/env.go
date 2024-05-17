package conn

import "os"

var (
	GORM_HOST     = os.Getenv("GORM_HOST")
	GORM_PORT     = os.Getenv("GORM_PORT")
	GORM_USERNAME = os.Getenv("GORM_USERNAME")
	GORM_PASSWORD = os.Getenv("GORM_PASSWORD")
	GORM_DATABASE = os.Getenv("GORM_DATABASE")
)
