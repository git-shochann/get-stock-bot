package main

import (
	"get-stock-bot/auth"
	"get-stock-bot/csv"
	"get-stock-bot/discord"
	"get-stock-bot/stock"
)

func main() {

	// CSV読み込み
	setting := csv.LoadCSV()

	// discordBot起動
	// discord.StartBot()
	// fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	// sc := make(chan os.Signal, 1)
	// signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	// <-sc

	// APIを使うための準備
	authResponse := auth.Auth(setting)

	// 全在庫取得
	getStockResponse := stock.GetStock(authResponse)

	// WebHook送信
	discord.SendWebhook(getStockResponse)

}
