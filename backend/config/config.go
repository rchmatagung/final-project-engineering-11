package config

import (
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	duration, err := strconv.Atoi(os.Getenv("JWT_EXPIRATION_DURATION"))
	if err != nil {
		panic(err)
	}
	JWT_EXPIRATION_DURATION = time.Hour * time.Duration(duration)
	JWT_SECRET = os.Getenv("JWT_SECRET")

}

var JWT_SECRET = ""
var JWT_EXPIRATION_DURATION time.Duration

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "HiCoder Company <hicoder224@gmail.com>"
const CONFIG_AUTH_EMAIL = "hicoder224@gmail.com"
const CONFIG_AUTH_PASSWORD = "pkadliooghsdelxk"

var Mutex sync.Mutex
