package main

import (
	"os"
)

type Config struct {
	Port         string
	Host         string
	ReadTimeout  int
	WriteTimeout int
}

func LoadConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	host := os.Getenv("HOST")
	if host == "" {
		host = "0.0.0.0"
	}

	return &Config{
		Port:         port,
		Host:         host,
		ReadTimeout:  15,
		WriteTimeout: 15,
	}
}


