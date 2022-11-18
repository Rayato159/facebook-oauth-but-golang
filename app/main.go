package main

import (
	"log"
	"os"

	"github.com/Rayato159/facebook-oauth-but-go/configs"
	databases "github.com/Rayato159/facebook-oauth-but-go/package/databases/mongodb"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	// Check args[1] is exist? load that .env file: load .env by default
	log.Println("server is starting, please wait...")
	var err error
	if len(os.Args) > 1 {
		err = godotenv.Load(os.Args[1])
	} else {
		err = godotenv.Load(".env")
	}
	if err != nil {
		log.Fatal("error, loading .env file has failed")
	}

	// Set configs
	cfg := &configs.Configs{
		Database: configs.MongoDb{
			Host: os.Getenv(""),
			Port: os.Getenv(""),
		},
		App: configs.App{
			Host:                 os.Getenv(""),
			Port:                 os.Getenv(""),
			ServerRequestTimeout: os.Getenv(""),
			Stage:                os.Getenv(""),
		},
	}
	log.Printf("setting config as \"%s\"", cfg.App.Stage)

	// Connecting to database (MongoDb) and return it as type *mongo.Client
	db, err := databases.NewMongoDbConnect(cfg.Database)
	if err != nil {
		log.Fatal(err)
	}
	// Before the end of the main function -> Close the database connection
	defer databases.MongoDbDisconnect(db)
}
