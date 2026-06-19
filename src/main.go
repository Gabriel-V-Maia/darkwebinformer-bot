package main

import (
	"fmt"
	"os"

	"darkwebinformer-bot/src/modules/envreader"
	"github.com/bwmarrin/discordgo"
)

func main() {
	envreader.LoadEnv(".env")

	token := os.Getenv("TOKEN")

	dg, err := discordgo.New("Bot " + token)

	if err != nil {
		fmt.Println("Error creating discord session", err)
	}

	dg.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Printf("Logged on as %s\n", s.State.User.Username)
	})

	err = dg.Open()

	if err != nil {
		fmt.Println("Error opening websocket to discord: ", err)
	}

	select {}
}
