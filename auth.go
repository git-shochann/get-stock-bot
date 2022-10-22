package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type AuthRequest struct {
	Id        string `json:"id"`
	Pwd       string `json:"pwd"`
	GrantType string `json:"grant_type"`
}

type AuthResponse struct {
}

func Auth(baseUrl string) {

	endpoint := BaseUrl + "api/auth"
	method := "POST"

	file, err := os.Open("profile.csv") // ファイルを開いて構造体`型`の値を返却する
	if err != nil {
		log.Fatalln("faild to read profile.csv")
	}
	defer file.Close()

	r := csv.NewReader(file) // コンストラクタ関数で構造体`型`の値を初期化する -> r -> csv.Reader`型` -> フィールドにアクセス可能, またはメソッドにアクセスが可能 -> NewReaderにホバーして何をするかの説明は別に後ほどでいい
	r.FieldsPerRecord = -1

	_, err = r.Read()
	if err != nil {
		log.Fatalln("faild to read profile.csv")
	}

	rows, err := r.ReadAll()
	if err != nil {
		log.Fatalln("faild to read profile.csv")
	}

	var Profile []string

	for _, v := range rows {
		Profile = v
	}

	requestBody := AuthRequest{
		Id:        Profile[0],
		Pwd:       Profile[1],
		GrantType: "password",
	}

	encodedJson, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatalln("faild to read profile.csv")
	}

	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(encodedJson))
	if err != nil {
		log.Fatalln("unknown error")
	}

	req.Header.Set("Content-Type", "application/json")

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

	fmt.Println(string(data))

}
