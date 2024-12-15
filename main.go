package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/rayfiyo/spotify-BPM-sorter/cmd"
	"github.com/tidwall/gjson"
	"github.com/zmb3/spotify/v2"
)

func main() {
	// 認証とコンテキスト作成
	ctx, client, err := cmd.Auth()
	if err != nil {
		log.Fatal(err)
	}

	// 実行時引数からプレイリストのIDを取得
	if len(os.Args) < 2 {
		log.Fatal("Expected index out of range. Insufficient argument missing.")
	}
	playlistID := spotify.ID(os.Args[1])

	// プレイリストアイテムを取得
	result, err := client.GetPlaylistItems(ctx, playlistID)
	if err != nil {
		log.Fatal(err)
	}

	// プレイリストアイテムを JSON に変換
	jsonData, err := json.Marshal(result.Items)
	if err != nil {
		log.Fatalf("Failed to marshal playlist items: %v", err)
	}

	// JSON データを解析
	items := gjson.ParseBytes(jsonData)

	// 各アイテムから id を抽出
	items.ForEach(func(key, value gjson.Result) bool {
		id := value.Get("track").String()
		fmt.Println(id)
		return true
	})
}
