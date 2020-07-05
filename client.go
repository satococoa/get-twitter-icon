package client

import (
	"context"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2/clientcredentials"
)

type Client struct {
	twClient *twitter.Client
}

func NewClient(ctx context.Context, clientID, clientSecret string) *Client {
	config := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}
	httpClient := config.Client(ctx)

	return &Client{
		twClient: twitter.NewClient(httpClient),
	}
}

func (c *Client) GetTwitterIconURL(screenName string) (string, error) {
	user, _, err := c.twClient.Users.Show(&twitter.UserShowParams{ScreenName: screenName})
	if err != nil {
		return "", err
	}
	normalImageURL := user.ProfileImageURLHttps
	// remove "_normal"
	originalImageURL := strings.Replace(normalImageURL, "_normal", "", 1)
	return originalImageURL, nil
}
