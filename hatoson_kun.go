package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/nlopes/slack"
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
	SLACKAPIKEY := os.Getenv("SLACKAPIKEY")
	api := slack.New(SLACKAPIKEY)

	groups, err := api.GetGroups(false)

	if err != nil {
		fmt.Println("%s \n", err)
		return
	}

	for _, group := range groups {
		fmt.Println("ID : %s, Name : %s\n", group.ID, group.Name)
	}
}
