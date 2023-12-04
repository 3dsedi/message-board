package db

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"net/http"
	"github.com/joho/godotenv"
	"github.com/sedi/message-board/model"
)

var (
	conn      *sql.DB
	initDB    sync.Once
	dbInitErr error
)

func InitDBConnection() error {
	initDB.Do(func() {
		var err error
		conn, err = CreateDBconnection()
		if err != nil {
			dbInitErr = err
		}
	})
	return dbInitErr
}

func CreateDBconnection() (*sql.DB, error) {
	envFile, _ := godotenv.Read(".env")
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		envFile["DB_HOST"],
		envFile["DB_PORT"],
		envFile["DB_USER"],
		envFile["DB_PASSWORD"],
		envFile["DB_NAME"],
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("error opening database connection: %w, please check docker daemon/postgre engine is running", err)
		return nil, model.MessageBoardError{
			Message: "can not connect to database",
			Code:    http.StatusInternalServerError,
		}
	}
	if err := db.Ping(); err != nil {
		log.Fatal("error pinging database: %w", err)
		return nil, model.MessageBoardError{
			Message: "can not connect to database",
			Code:    http.StatusInternalServerError,
		}
	}
	log.Println("Connected to PostgreSQL database!")
	return db, nil
}
