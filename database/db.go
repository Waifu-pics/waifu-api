package database

import (
	"database/sql"
	"log"

	"github.com/Waifu-pics/waifu-api/config"

	_ "github.com/go-sql-driver/mysql" // MySQL
)

// InitSQL : MySQL Initializer
func InitSQL(config config.Config) Database {
	db, err := sql.Open("mysql", config.DatabaseUrl)
	if err != nil {
		log.Panicln("Error: " + err.Error())
	}
	return Database{db: db}
}
