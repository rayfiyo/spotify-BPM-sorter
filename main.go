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

	// プレイリストを取得
	playlist, err := client.GetPlaylist(ctx, playlistID)
	if err != nil {
		log.Fatal(err)
	}

	// プレイリストのトラック要素を JSON に変換
	tracksJson, err := json.Marshal(playlist.Tracks)
	if err != nil {
		log.Fatalf("Failed to marshal playlist tracks: %v", err)
	}

	// JSON データを解析し，items キーのみ抽出
	items := gjson.ParseBytes(tracksJson).Get("items")

	// 各アイテムから id を抽出
	items.ForEach(func(key, value gjson.Result) bool {
		// トラック id
		id := value.Get("track.id").String()
		fmt.Println(id)

		// トラック名
		fmt.Println("    " + value.Get("track.name").String())

		// 追加日時
		fmt.Println("    " + value.Get("added_at").String())

		return true
	})

	// プレイリストフォロワーを JSON に変換
	followersJson, err := json.Marshal(playlist.Followers)
	if err != nil {
		log.Fatalf("Failed to marshal playlist followers: %v", err)
	}

	// プレイリストフォロワーを出力（href は現在未実装のためパースせずに出力）
    // 特定のユーザーが特定のプレイリストをフォローしているか調べる方法は開発者オプションとして API あり
    // https://developer.spotify.com/documentation/web-api/reference/check-current-user-follows
	fmt.Println("Followers: "+string(followersJson))
}
