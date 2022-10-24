# Get Stock Bot

既存の在庫データを取得し、Discord に通知

## 事前準備

- root に setting.csv を用意(スプレッドシートで管理 -> CSV に変換して DL がおすすめ)

手動で用意するなら...

ヘッダー行に、`id,password,discordwebhook`
2 行目以降にデータを記入

## DL 方法(暫定)

Windows -> exe ファイルとして配布予定

Mac -> dmg ファイルとして配布予定

```shell
    git clone <this repository>
    go run main.go
```
