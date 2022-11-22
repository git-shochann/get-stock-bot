package discord

import (
	"fmt"
	"get-stock-bot/auth"
	"get-stock-bot/csv"
	"get-stock-bot/stock"
	"log"

	"github.com/bwmarrin/discordgo"
)

var BotID string
var Setting []string

func StartBot() {

	// CSV読み込み
	Setting := csv.LoadCSV()

	discord, err := discordgo.New("Bot " + Setting[4])
	if err != nil {
		log.Fatalln(err)
	}

	user, err := discord.User("@me")
	if err != nil {
		log.Fatalln(err)
	}

	BotID = user.ID

	// ハンドラーを登録する
	discord.AddHandler(messageHandler)

	// websocketを作成する
	err = discord.Open()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Bot is running!")

}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}
	if m.Content == "!stock" {
		authResponse := auth.Auth(Setting)
		getStockResponse := stock.GetStock(authResponse)
		SendMessage(getStockResponse)
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID)
	}
}
