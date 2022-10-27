package csv

import (
	"encoding/csv"
	"log"
	"os"
)

func LoadCSV() []string {

	file, err := os.Open("./setting.csv") // ファイルを開いて構造体`型`の値を返却する
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

	return Setting

}
