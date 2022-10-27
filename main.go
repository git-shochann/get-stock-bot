package main

import (
	"get-stock-bot/auth"
	"get-stock-bot/csv"
	"get-stock-bot/discord"
	"get-stock-bot/stock"
	"runtime"

	"github.com/robfig/cron/v3"
)

func main() {

	// APIを使うための準備
	authResponse := auth.Auth()

	// 全在庫取得
	getStockResponse := stock.GetStock(authResponse)

	// WebHook送信
	discord.SendWebhook(getStockResponse)

}

// 定期実行させる
func init() {
	c := cron.New()
	setting := csv.LoadCSV()
	monitorDelay := setting[3]
	c.AddFunc("@every "+monitorDelay+"s", main)
	runtime.Goexit()
}
