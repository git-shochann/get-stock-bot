package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type GetStockResponse struct {
	Count   int `json:"count"`
	Results Results
}

type Results struct {
	Itemid       string `json:"itemid"`
	Skuid        string `json:"skuid"`
	Itemname     string `json:"itemname"`
	Lot          string `json:"lot"`
	Stockqty     string `json:"stockqty"`
	Itemcategory string `json:"itemcategory"`
}

// 在庫を取得する
func GetStock(auth AuthResponse) GetStockResponse {

	endpoint := BaseUrl + "api/stockresult"
	method := "GET"
	bearerToken := "Bearer " + auth.AccessToken

	fmt.Println(bearerToken)

	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		log.Fatalln("unknown error")
	}
	req.Header.Add("Authorization", bearerToken)

	cliant := http.Client{}
	res, err := cliant.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	// []byteにする
	byteBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var getStockResponse GetStockResponse

	err = json.Unmarshal(byteBody, &getStockResponse)
	if err != nil {
		log.Fatalln("faild to read json")
	}

	return getStockResponse

}
