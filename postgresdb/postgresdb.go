package postgresdb

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const host string = "host=localhost"
const user string = "user=gorm"
const passwd string = "password=gorm"
const database string = "dbname=test"
const port string = "port=9920"
const ssl string = "sslmode=disable"
const timezone string = "TimeZone=Asia/Karachi"

func ORM() (*gorm.DB, error) {
	dns := host + " " + user + " " + passwd + " " + database + " " + port + " " + ssl + " " + timezone
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		log.Println("Error opening database:", err)
		return &gorm.DB{}, err
	}

	log.Println("ORM:", db)
	return db, nil
}
