package discord

import (
	"fmt"
	"get-stock-bot/stock"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func SendMessage(stock stock.GetStockResponse) *discordgo.MessageEmbed {

	// var Df DiscordField
	// // var NewArr []DiscordField
	// var AllStock int

	// // 1個でまずは処理を考えてみること
	// for _, v := range stock.Results {
	// 	// 必要なデータのみ抽出する
	// 	Df = DiscordField{
	// 		Name:   v.Itemname,
	// 		Value:  v.Stockqty + "個",
	// 		Inline: true,
	// 	}
	// 	stock, err := strconv.Atoi(v.Stockqty)
	// 	if err != nil {
	// 		fmt.Println("convert error")
	// 		return nil
	// 	}
	// 	AllStock += stock
	// 	// 以下で渡す配列に1つの構造体のデータを加えて新しく作成する
	// 	// NewArr = append(NewArr, Df)
	// }

	// marshaledJson, err := json.Marshal(data)
	// if err != nil {
	// 	fmt.Println("json err:", err)
	// 	return nil
	// }

	// return &data

	// まずMessageEmbedFieldの作成
	var AllStock int
	var MessageEmbedField []*discordgo.MessageEmbedField
	for _, v := range stock.Results {
		message := discordgo.MessageEmbedField{
			Name:   v.Itemname,
			Value:  v.Stockqty + "個",
			Inline: true,
		}
		stock, err := strconv.Atoi(v.Stockqty)
		if err != nil {
			fmt.Println("convert error")
			return nil
		}
		AllStock += stock
		MessageEmbedField = append(MessageEmbedField, &message)
	}
	data := discordgo.MessageEmbed{
		Fields: MessageEmbedField,
	}

	return &data
}

// type MessageEmbed struct {
// 	URL         string                 `json:"url,omitempty"`
// 	Type        EmbedType              `json:"type,omitempty"`
// 	Title       string                 `json:"title,omitempty"`
// 	Description string                 `json:"description,omitempty"`
// 	Timestamp   string                 `json:"timestamp,omitempty"`
// 	Color       int                    `json:"color,omitempty"`
// 	Footer      *MessageEmbedFooter    `json:"footer,omitempty"`
// 	Image       *MessageEmbedImage     `json:"image,omitempty"`
// 	Thumbnail   *MessageEmbedThumbnail `json:"thumbnail,omitempty"`
// 	Video       *MessageEmbedVideo     `json:"video,omitempty"`
// 	Provider    *MessageEmbedProvider  `json:"provider,omitempty"`
// 	Author      *MessageEmbedAuthor    `json:"author,omitempty"`
// 	Fields      []*MessageEmbedField   `json:"fields,omitempty"`
// }

// type MessageEmbedField struct {
// 	Name   string `json:"name,omitempty"`
// 	Value  string `json:"value,omitempty"`
// 	Inline bool   `json:"inline,omitempty"`
// }
