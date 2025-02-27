package dao

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

var db *sql.DB

type configFile struct {
	DB database `toml:"database"`
}

type database struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SSLmode  string
}

//InitDB reads a TOML file and creates a database connection
func InitDB() {
	var conf configFile
	var err error

	if _, err = toml.DecodeFile("config.toml", &conf); err != nil {
		log.Println(err)
		return
	}

	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		conf.DB.Host, conf.DB.Port, conf.DB.User, conf.DB.Password, conf.DB.Name, conf.DB.SSLmode)
	db, err = sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Printf("host: %s, port: %d, user: %s, name: %s",
		conf.DB.Host, conf.DB.Port, conf.DB.User, conf.DB.Name)
}

//CloseDB close the database connection
func CloseDB() {
	db.Close()
}
