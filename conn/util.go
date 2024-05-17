package conn

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

/*
Initialize the GORM configuration with the default logger.
*/
func gormConfig(silent ...bool) *gorm.Config {
	if len(silent) > 0 && silent[0] {
		return &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		}
	}
	return &gorm.Config{}
}
