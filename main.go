package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rayfiyo/spotify-BPM-sorter/cmd"
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

	type trackInfo struct {
		Name string
		ID   spotify.ID
	}

	var tracks []trackInfo
	for _, item := range result.Items {
		track := item.Track.Track
		if track == nil {
			continue
		}
		if track.ID == "" {
			continue
		}
		tracks = append(tracks, trackInfo{
			Name: track.Name,
			ID:   track.ID,
		})
	}

	if len(tracks) == 0 {
		log.Println("プレイリストに有効なトラックが見つかりませんでした")
		return
	}

	const batchSize = 100
	featureByID := make(map[spotify.ID]*spotify.AudioFeatures)

	for start := 0; start < len(tracks); start += batchSize {
		end := start + batchSize
		end = min(end, len(tracks))

		ids := make([]spotify.ID, end-start)
		for i := start; i < end; i++ {
			ids[i-start] = tracks[i].ID
		}

		audioFeatures, err := client.GetAudioFeatures(ctx, ids...)
		if err != nil {
			var apiErr spotify.Error
			if errors.As(err, &apiErr) && apiErr.Status == http.StatusForbidden {
				log.Printf("403 Forbidden が発生したため個別に特徴量を取得します: %v", ids)
				for _, id := range ids {
					features, singleErr := client.GetAudioFeatures(ctx, id)
					if singleErr != nil {
						log.Printf("トラック %s の特徴量取得に失敗したためスキップします: %v", id, singleErr)
						continue
					}
					for _, feature := range features {
						if feature == nil {
							continue
						}
						featureByID[feature.ID] = feature
					}
				}
				continue
			}
			log.Fatalf("オーディオ特徴量の取得に失敗しました: %v\n%v", err, ids)
		}

		for _, feature := range audioFeatures {
			if feature == nil {
				continue
			}
			featureByID[feature.ID] = feature
		}
	}

	for _, track := range tracks {
		tempo := "-"
		if feature, ok := featureByID[track.ID]; ok && feature != nil {
			tempo = fmt.Sprintf("%.2f", feature.Tempo)
		}
		fmt.Printf("%s, %s, %s\n", track.Name, tempo, track.ID)
	}
}
