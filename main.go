package main

import "get-stock-bot/discord"

const BaseUrl = "https://www.service-netdepot.jp/NetDepotWebAPI/"

func main() {

	// APIを使うための準備
	authResponse := Auth()

	// 全在庫取得
	getStockResponse := GetStock(authResponse)

	// WebHook送信
	discord.SendWebhook(getStockResponse)

}
