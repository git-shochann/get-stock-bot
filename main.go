package main

import (
	"get-stock-bot/auth"
	"get-stock-bot/discord"
	"get-stock-bot/stock"
)

func main() {

	// APIを使うための準備
	authResponse := auth.Auth()

	// 全在庫取得
	getStockResponse := stock.GetStock(authResponse)

	// WebHook送信
	discord.SendWebhook(getStockResponse)

}
