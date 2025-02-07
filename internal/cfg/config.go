package cfg

import (
	"os"

	_ "github.com/joho/godotenv/autoload" // for development to load .env file
)

var (
	PORT       = os.Getenv("PORT")
	AUTHSECRET = os.Getenv("AUTHSECRET")
	CON_STRING = os.Getenv("CON_STRING")
)
