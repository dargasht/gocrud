package cfg

import (
	"os"

	_ "github.com/joho/godotenv/autoload" // for development to load .env file
)

var (
	PORT = os.Getenv("PORT")
)
