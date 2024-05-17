package gorm

import "os"

var (
	GORM_ENCRYPT_KEY = os.Getenv("GORM_ENCRYPT_KEY")
)
