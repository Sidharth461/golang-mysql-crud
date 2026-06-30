package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		user,
		password,
		host,
		port,
		database,
	)

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal("Database not Reachable:", err)
	}

	log.Println("Connected to MySQL Successfully")
}
