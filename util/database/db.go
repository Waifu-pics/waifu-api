package database

import (
	"database/sql"
	"log"

	"waifu.pics/util/config"

	_ "github.com/go-sql-driver/mysql" // MySQL
)

// InitSQL : MySQL Initializer
func InitSQL(config config.Config) Database {
	db, err := sql.Open("mysql", config.DB.URL)
	if err != nil {
		log.Panicln("Error: " + err.Error())
	}
	return Database{db: db}
}
