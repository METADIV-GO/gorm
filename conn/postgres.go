package conn

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*
PostgreSQL connects to a MySQL database and returns a gorm.DB object.
By default, the silent parameter is false, which means that the
gorm.DB object will log all SQL statements to the console.
*/
func PostgreSQL(host, port, username, password, database string, silent ...bool) (*gorm.DB, error) {
	return gorm.Open(
		postgres.Open(
			fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
				host, port, username, database, password)),
		gormConfig(silent...),
	)
}

/*
QuickPostgreSQL connects to a MySQL database using the environment variables.
*/
func QuickPostgreSQL() (*gorm.DB, error) {
	return PostgreSQL(GORM_HOST, GORM_PORT, GORM_USERNAME, GORM_PASSWORD, GORM_DATABASE)
}
