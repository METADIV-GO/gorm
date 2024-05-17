package conn

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

/*
Sqlite connects to a SQLite database and returns a gorm.DB object.
By default, the silent parameter is false, which means that the
gorm.DB object will log all SQL statements to the console.
*/
func Sqlite(path string, silent ...bool) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(path), gormConfig(silent...))
}

/*
SqliteMem connects to a SQLite database in memory and returns a gorm.DB object.
By default, the silent parameter is false, which means that the
gorm.DB object will log all SQL statements to the console.
*/
func SqliteMem(silent ...bool) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("file::memory:?cache=shared&busy_timeout=5000"),
		gormConfig(silent...))
}
