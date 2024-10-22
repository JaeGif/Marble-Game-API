package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)


var db *sql.DB

func InitDB() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  dbHost := os.Getenv("DB_HOST")
  dbPort := os.Getenv("DB_PORT")
  dbUser := os.Getenv("DB_USER")
  dbPass := os.Getenv("DB_PASSWORD")
  dbName := os.Getenv("DB_NAME")
  dbExt := os.Getenv("DB_EXT")

  db, err = sql.Open("postgres", fmt.Sprintf("name=%s://user=%s:password=%s@host=%s:port=%s/ext=%s sslmode=disable", dbName, dbUser, dbPass, dbHost, dbPort, dbExt))

  if err != nil {
    panic(err.Error())
  }

  err = db.Ping()
  if err != nil {
    panic(err.Error())
  }

  fmt.Println("Successfully connected to database")
}

func GetDB() *sql.DB{
    return db
}
// postgresql://postgres:WIgzYmcKWHKhPygDhJmDoIzZZLYLFcDP@autorack.proxy.rlwy.net:28107/railway