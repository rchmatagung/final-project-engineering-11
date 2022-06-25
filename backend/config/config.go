package config

import (
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load()

	duration, _ := strconv.Atoi(os.Getenv("JWT_EXPIRATION_DURATION"))

	JWT_EXPIRATION_DURATION = time.Hour * time.Duration(duration)
	JWT_SECRET = os.Getenv("JWT_SECRET")
	PORT = os.Getenv("PORT")
	smptpport, _ := strconv.Atoi(os.Getenv("CONFIG_SMTP_PORT"))
	CONFIG_SMTP_HOST = os.Getenv("CONFIG_SMTP_HOST")
	CONFIG_SMTP_PORT = smptpport
	CONFIG_SENDER_NAME = os.Getenv("CONFIG_SENDER_NAME")
	CONFIG_AUTH_EMAIL = os.Getenv("CONFIG_AUTH_EMAIL")
	CONFIG_AUTH_PASSWORD = os.Getenv("CONFIG_AUTH_PASSWORD")

}

var JWT_SECRET = ""
var JWT_EXPIRATION_DURATION time.Duration
var PORT = ""

var CONFIG_SMTP_HOST = "hicoder"
var CONFIG_SMTP_PORT = 5
var CONFIG_SENDER_NAME = "HiCoder Company <hicoder224@gmail.com>"
var CONFIG_AUTH_EMAIL = "hicoder224@gmail.com"
var CONFIG_AUTH_PASSWORD = "pkadliooghsdelxk"

var Mutex sync.Mutex
