package auth

import (
	"encoding/json"
	"get-stock-bot/util"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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
func Auth(setting []string) AuthResponse {

	endpoint := util.BaseUrl + "api/auth"
	method := "POST"

	form := url.Values{}
	form.Add("id", setting[0])
	form.Add("pwd", setting[1])
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
