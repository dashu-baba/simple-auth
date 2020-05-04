package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	//Configuration defines config variables
	Configuration Config
)

// Config godoc
// Defines the configuration structure
type Config struct {
	//Db denotes database configuration
	Db struct {
		EndPoint string
		Name     string
	}
	//Server denotes server configuration
	Server struct {
		Port string
	}
}

//New creates a new instance of Config
func New() Config {
	config := Config{}
	config.Db.EndPoint = os.Getenv("DB_CONN")
	config.Db.Name = os.Getenv("DB_NAME")
	config.Server.Port = os.Getenv("PORT")
	return config
}

//init sets up the initial states
func init() {
	// Load environment variable from .env file
	err := godotenv.Load("dev.env")
	if err != nil {
		log.Printf("Error loading env file %v", err)
	}

	Configuration = New()
}
