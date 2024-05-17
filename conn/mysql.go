package conn

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
MySQL connects to a MySQL database and returns a gorm.DB object.
By default, the silent parameter is false, which means that the
gorm.DB object will log all SQL statements to the console.
*/
func MySQL(host, port, username, password, database string, silent ...bool) (*gorm.DB, error) {
	return gorm.Open(
		mysql.Open(
			fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				username, password, host, port, database)),
		gormConfig(silent...),
	)
}

/*
QuickMySQL connects to a MySQL database using the environment variables.
*/
func QuickMySQL(silent ...bool) (*gorm.DB, error) {
	return MySQL(GORM_HOST, GORM_PORT, GORM_USERNAME, GORM_PASSWORD, GORM_DATABASE, silent...)
}
