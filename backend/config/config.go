package config

import "time"

var Configuration = struct {
	JWT_SECRET              string
	JWT_EXPIRATION_DURATION time.Duration
	DB_HOST                 string
	DB_PORT                 string
	DB_USER                 string
	DB_PASSWORD             string
	DB_NAME                 string
	DB_SSL_MODE             string
	DB_TIMEOUT              string
	DB_MAX_CONNECTIONS      string
	DB_MAX_IDLE_CONNECTIONS string
	DB_MAX_LIFE_TIME        string
	DB_MAX_OPEN_REQUESTS    string
	DB_MAX_IDLE_TIME        string
}{
	JWT_SECRET:              "hicoder",
	JWT_EXPIRATION_DURATION: time.Hour * 3,
	DB_HOST:                 "localhost",
	DB_PORT:                 "5432",
	DB_USER:                 "postgres",
	DB_PASSWORD:             "postgres",
	DB_NAME:                 "postgres",
	DB_SSL_MODE:             "disable",
	DB_TIMEOUT:              "30",
	DB_MAX_CONNECTIONS:      "0",
	DB_MAX_IDLE_CONNECTIONS: "0",
	DB_MAX_LIFE_TIME:        "0",
	DB_MAX_OPEN_REQUESTS:    "0",
	DB_MAX_IDLE_TIME:        "0",
}

var KEYPRIVATE = []byte("656#@$$6622)(){]{:^&")

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "HiCoder Company <hicoder224@gmail.com>"
const CONFIG_AUTH_EMAIL = "hicoder224@gmail.com"
const CONFIG_AUTH_PASSWORD = "pkadliooghsdelxk"
