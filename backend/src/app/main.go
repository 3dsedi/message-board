package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/sedi/messageBoard/api"
	"github.com/sedi/messageBoard/db"
)

func main() {
	initSystemVar()
	userStore, messageStore, err := initDatabase()
	if err != nil {
		log.Fatal(err)
		return
	}

	server := api.NewAPIServer(initSystemVar(), userStore, messageStore)
	server.Run()

}

func initDatabase() (*db.PostgresUserStorage, *db.PostgresMessageStorage, error) {
	err := db.InitDBConnection()
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}

	userStore := db.NewPostgresUserStorage()
	messageStore := db.NewPostgresMessageStorage()
	
	if err := userStore.Init(); err != nil {
		log.Fatal(err)
		return nil, nil, err
	}

	if err := messageStore.Init(); err != nil {
		log.Fatal(err)
		return nil, nil, err
	}

	log.Printf("app started successfully")
	return userStore, messageStore, nil
}

func initSystemVar() string {
	envFile, _ := godotenv.Read(".env")
	return fmt.Sprintf(":%s", envFile["PORT"])
}
