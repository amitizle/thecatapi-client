package thecatapi

import (
	http "github.com/amitizle/thecatapi_client/internal/http_client"
)

type Client struct {
	BaseURL string
	ApiKey  string
	http    *http.HTTPClient

	Favourites *FavouriteService
	Images     *ImageService
}

func NewClient() (*Client, error) {
	httpClient, _ := http.NewClient()
	client := &Client{
		http: httpClient,
	}
	favouriteService, _ := NewFavouriteService(client)
	imageService, _ := NewImageService(client)
	client.Favourites = favouriteService
	client.Images = imageService
	return client, nil
}
