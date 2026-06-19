package main

import (
	"fmt"
	"os"

	"darkwebinformer-bot/src/modules/envreader"
	"darkwebinformer-bot/src/modules/notifier"
	"darkwebinformer-bot/src/modules/rssreader"

	"github.com/bwmarrin/discordgo"
)

func main() {
	envreader.LoadConfigs(".env")
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

	file, err := rssreader.GetRSS("https://darkwebinformer.com/rss", "latest.rss")

	if err != nil {
		fmt.Printf("Error getting rss content: %s", err)
	}

	defer file.Close()

	matches, err := rssreader.ParseRSS(file)

	// fmt.Println(matches)

	notifier.SendEmbed(dg)

	select {}
}
