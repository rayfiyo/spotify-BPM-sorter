package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2/clientcredentials"
)

type SpotifyClients struct {
	ID      string
	Sercret string
}

const (
	envSpotifyID     = "SPOTIFY_ID"
	envSpotifySecret = "SPOTIFY_SECRET"
)

func main() {
	ctx, client, err := Auth()
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) < 2 {
		log.Fatal("Expected index out of range. Insufficient argument missing.")
	}

	playlistID := spotify.ID(os.Args[1])
	result, err := client.GetPlaylistItems(ctx, playlistID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}

func Auth() (context.Context, *spotify.Client, error) {
	spotifyClients, err := LoadEnv()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	config := &clientcredentials.Config{
		ClientID:     spotifyClients.ID,
		ClientSecret: spotifyClients.Sercret,
		TokenURL:     spotifyauth.TokenURL,
	}
	token, err := config.Token(ctx)
	if err != nil {
		log.Fatalf("Failed to retrieve token: %v", err)
	}

	httpClient := spotifyauth.New().Client(ctx, token)

	return ctx, spotify.New(httpClient), nil
}

func LoadEnv() (*SpotifyClients, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}

	clientID := os.Getenv(envSpotifyID)
	clientSecret := os.Getenv(envSpotifySecret)

	return &SpotifyClients{
		ID:      clientID,
		Sercret: clientSecret,
	}, nil
}
