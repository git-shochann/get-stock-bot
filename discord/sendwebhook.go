package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"get-stock-bot/stock"
	"io/ioutil"
	"log"
	"net/http"
)

const webHookUrl = "https://discord.com/api/webhooks/1033295235188523008/SEneGDA02Lp0M1-NzJuMJxUyC2Dxj2g2Pa8wQafzZWmQf5H_9yuky3uIeZj3AhQqvOkk"

// テスト段階
func SendWebhook(stock stock.GetStockResponse) {

	var Df DiscordField
	var NewArr []DiscordField

	// 1個でまずは処理を考えてみること
	for _, v := range stock.Results {
		// 必要なデータのみ抽出する
		Df = DiscordField{
			Name:   v.Itemname,
			Value:  v.Stockqty + "個",
			Inline: true,
		}
		// 以下で渡す配列に1つの構造体のデータを加えて新しく作成する
		NewArr = append(NewArr, Df)
	}

	dw := &DiscordWebhook{UserName: "Egitee"}
	dw.Embeds = []DiscordEmbed{
		DiscordEmbed{
			// Title: "残りの在庫数: " + strconv.Itoa(stock.Count),
			Title:  "残りの在庫数: 100",
			URL:    "https://www.service-netdepot.jp/Contents/StockList.aspx",
			Color:  3066993,
			Fields: NewArr,
		},
	}

	j, err := json.Marshal(dw)
	if err != nil {
		fmt.Println("json err:", err)
		return
	}

	req, err := http.NewRequest("POST", webHookUrl, bytes.NewBuffer(j))
	if err != nil {
		fmt.Println("new request err:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("client err:", err)
		return
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(data))

	fmt.Println("webhook sent!")
}
