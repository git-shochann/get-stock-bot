package discord

import (
	"get-stock-bot/stock"

	"github.com/bwmarrin/discordgo"
)

func CreateMessage(stock stock.GetStockResponse) *discordgo.MessageEmbed {

	// MessageEmbedFieldの作成
	var MessageEmbedField []*discordgo.MessageEmbedField
	for _, v := range stock.Results {
		message := discordgo.MessageEmbedField{
			Name:   v.Itemname,
			Value:  v.Stockqty + "個",
			Inline: true,
		}
		MessageEmbedField = append(MessageEmbedField, &message)
	}

	// MessageEmbedの作成
	messageEmbed := discordgo.MessageEmbed{
		Fields: MessageEmbedField,
	}

	return &messageEmbed
}
