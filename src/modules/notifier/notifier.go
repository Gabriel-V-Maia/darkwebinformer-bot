package notifier

import (
	"strings"

	"darkwebinformer-bot/src/modules/rssreader"

	"github.com/bwmarrin/discordgo"
)

func SendEmbed(session *discordgo.Session, channelID string, item rssreader.Item) {
	targetCompany := "Não especificada"
	if strings.Contains(item.Title, "From ") {
		parts := strings.Split(item.Title, "From ")
		if len(parts) > 1 {
			targetCompany = parts[1]
		}
	} else if strings.Contains(item.Title, "Allegedly Breached") {
		parts := strings.Split(item.Title, " Allegedly")
		targetCompany = parts[0]
	}

	desc := item.Description
	if len(desc) > 1000 {
		desc = desc[:997] + "..."
	}

	embed := &discordgo.MessageEmbed{
		Title:       item.Title,
		URL:         item.Link,
		Color:       0x99ffee,
		Description: desc,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Tópico / Categoria",
				Value:  item.Category,
				Inline: true,
			},
			{
				Name:   "Empresa Alvo",
				Value:  targetCompany,
				Inline: true,
			},
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://cdn.discordapp.com/avatars/119249192806776836/cc32c5c3ee602e1fe252f9f595f9010e.jpg?size=2048",
		},
	}

	if item.PubDate != "" {
		embed.Footer = &discordgo.MessageEmbedFooter{
			Text: "Publicado: " + item.PubDate,
		}
	}

	session.ChannelMessageSendEmbed(channelID, embed)
}
