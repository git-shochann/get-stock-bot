package main

const BaseUrl = "https://www.service-netdepot.jp/NetDepotWebAPI/"

func main() {

	// APIを使うための準備
	authResponse := Auth()

	GetStock(authResponse)

	// WebHook送信
	// discord.SendWebhook()

}
