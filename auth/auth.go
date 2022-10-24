package auth

import (
	"encoding/csv"
	"encoding/json"
	"get-stock-bot/util"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Issued       string `json:"issued"`
	Expires      string `json:"expires"`
}

// 認証API
func Auth() AuthResponse {

	endpoint := util.BaseUrl + "api/auth"
	method := "POST"

	file, err := os.Open("setting.csv") // ファイルを開いて構造体`型`の値を返却する
	if err != nil {
		log.Fatalln("faild to read setting.csv")
	}
	defer file.Close()

	r := csv.NewReader(file) // コンストラクタ関数で構造体`型`の値を初期化する -> r -> csv.Reader`型` -> フィールドにアクセス可能, またはメソッドにアクセスが可能 -> NewReaderにホバーして何をするかの説明は別に後ほどでいい
	r.FieldsPerRecord = -1

	_, err = r.Read()
	if err != nil {
		log.Fatalln("faild to read setting.csv")
	}

	rows, err := r.ReadAll()
	if err != nil {
		log.Fatalln("faild to read setting.csv")
	}

	var Setting []string

	for _, v := range rows {
		Setting = v
	}

	form := url.Values{}
	form.Add("id", Setting[0])
	form.Add("pwd", Setting[1])
	form.Add("grant_type", "password")

	body := strings.NewReader(form.Encode())

	req, err := http.NewRequest(method, endpoint, body)
	if err != nil {
		log.Fatalln("unknown error")
	}

	req.Header.Set("Content-Type", "application/x-wwww-form-unlencoded")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var authResponse AuthResponse

	err = json.Unmarshal(data, &authResponse)
	if err != nil {
		log.Fatalln("faild to read json")
	}

	return authResponse

}
