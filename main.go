package main

import (
	"fmt"
	"gorm.io/gorm"
  	"gorm.io/driver/mysql"
	"os"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := os.Getenv("DBConnectionStr")
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(db, err)
}