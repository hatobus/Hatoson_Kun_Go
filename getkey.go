package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Env_load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error! Loading .env file")
	}
}

func main() {
	Env_load()
	message := fmt.Sprintf("APIKEY = %s", os.Getenv("SLACKAPIKEY"))
	fmt.Println(message)
}
