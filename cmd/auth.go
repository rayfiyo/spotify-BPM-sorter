package cmd

import (
	"context"
	"log"

	"github.com/rayfiyo/spotify-BPM-sorter/cmd/config"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2/clientcredentials"
)

// Spotify への認証を行う
func Auth() (context.Context, *spotify.Client, error) {
	spotifyClients, err := config.LoadEnv()
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
