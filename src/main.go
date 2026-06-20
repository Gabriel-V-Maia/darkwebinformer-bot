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
	channelid := os.Getenv("CHANNEL_ID")

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

	posts, err := rssreader.ParseRSS(file)

	for _, post := range posts {
		if post.Title == "" || post.Title == "Dark Web Informer" {
			continue
		}

		notifier.SendEmbed(dg, channelid, post)

	}

	select {}
}
