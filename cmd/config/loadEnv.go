package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/rayfiyo/spotify-BPM-sorter/models"
)

const (
	envSpotifyID     = "SPOTIFY_ID"
	envSpotifySecret = "SPOTIFY_SECRET"
)

// .env ファイルから認証に必要な ID と シークレットトークンを読み込む
func LoadEnv() (*models.SpotifyClients, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}

	clientID := os.Getenv(envSpotifyID)
	clientSecret := os.Getenv(envSpotifySecret)

	return &models.SpotifyClients{
		ID:      clientID,
		Sercret: clientSecret,
	}, nil
}
