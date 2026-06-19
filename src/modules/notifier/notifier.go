package notifier

import (
	"github.com/bwmarrin/discordgo"
)

func SendEmbed(session *discordgo.Session) {
	embed := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       0x99ffee,
		Description: "{SEÇÃO}",
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "Tópico",
				Value:  "I am a value",
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "Empresa alvo",
				Value:  "I am a value",
				Inline: true,
			},
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://cdn.discordapp.com/avatars/119249192806776836/cc32c5c3ee602e1fe252f9f595f9010e.jpg?size=2048",
		},
		Title: "I am an Embed",
	}

	session.ChannelMessageSendEmbed("999", embed)
}
