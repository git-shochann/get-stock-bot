package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// テスト段階
func SendWebhook(stock GetStockResponse) {

	webHookUrl := "https://discord.com/api/webhooks/1033295235188523008/SEneGDA02Lp0M1-NzJuMJxUyC2Dxj2g2Pa8wQafzZWmQf5H_9yuky3uIeZj3AhQqvOkk"

	dw := &DiscordWebhook{UserName: "Egitee"}
	dw.Embeds = []DiscordEmbed{
		DiscordEmbed{
			Title: "残りの在庫",
			URL:   "https://www.service-netdepot.jp/Contents/StockList.aspx",
			Color: 3066993,
			Fields: []DiscordField{
				DiscordField{Name: "品番1", Value: "XXX", Inline: true},
				DiscordField{Name: "品番2", Value: "XXX", Inline: true},
			},
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
