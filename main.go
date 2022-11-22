package main

import (
	"get-stock-bot/csv"
	"get-stock-bot/discord"
)

func main() {

	// CSV読み込み
	setting := csv.LoadCSV()

	// discordBot起動
	discord.StartBot(setting)
	<-make(chan struct{})
	return

	// APIを使うための準備
	// authResponse := auth.Auth(setting)

	// 全在庫取得
	// getStockResponse := stock.GetStock(authResponse)

	// WebHook送信
	// discord.SendWebhook(getStockResponse)

}
