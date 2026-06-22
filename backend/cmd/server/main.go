package main

import (
	"backend/internal/config"
	"log"
	"net/http"
)

func main() {
	cfg , err  := config.LoadConfig()
	
	if err != nil {
	   log.Fatal("not able to load .env file")
	}

	


	http.ListenAndServe()
}
