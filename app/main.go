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
	cfg := &configs.Config{
		Database: configs.MongoDb{
			Host:     os.Getenv("MONGODB_HOST"),
			Port:     os.Getenv("MONGODB_PORT"),
			Username: os.Getenv("MONGODB_USERNAME"),
			Password: os.Getenv("MONGODB_PASSWORD"),
		},
		App: configs.App{
			Host:                 os.Getenv("FIBER_HOST"),
			Port:                 os.Getenv("FIBER_PORT"),
			ServerRequestTimeout: os.Getenv("FIBER_REQUEST_TIMEOUT"),
			Stage:                os.Getenv("STAGE"),
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
