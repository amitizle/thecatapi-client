package thecatapi

import (
	"encoding/json"
	http "github.com/amitizle/thecatapi_client/internal/http_client"
	"strconv"
	"strings"
)

type ImageService struct {
	client             *Client // TODO needed?
	downloadHTTPClient *http.HTTPClient
}

type SearchResponse struct {
	Id         string   `json:"id"`
	Url        string   `json:"url"`
	Breeds     []string `json:"breeds"`
	Categories []string `json:"categories"`
}

// type SearchRequest struct {
// 	limit     int
// 	mimeTypes []string
// 	format    string
// 	size      string
// 	order     string
// }

func NewImageService(client *Client) (*ImageService, error) {
	downloadHTTPClient, err := http.NewClient()
	if err != nil {
		return &ImageService{}, err
	}
	downloadHTTPClient.BaseURL = ""
	return &ImageService{
		client:             client,
		downloadHTTPClient: downloadHTTPClient,
	}, nil
}

// func (imageService *ImageService) Search() (*SearchResponse, error) {}

// TODO support pagination
// TODO struct?
func (imageService *ImageService) Search(mimeTypes []string, format string, limit int) ([]SearchResponse, error) {
	searchRequest := imageService.client.http.NewRequest()
	searchRequest.Path = "/v1/images/search"
	searchRequest.QueryParams["limit"] = strconv.Itoa(limit) // casting due to http client limitation
	searchRequest.QueryParams["format"] = format             // json / src TODO currently supporting only json
	searchRequest.QueryParams["mime_types"] = strings.Join(mimeTypes, ",")
	searchResponse, err := searchRequest.Get()
	if err != nil {
		return []SearchResponse{}, err
	}
	searchResult := make([]SearchResponse, 0)
	json.Unmarshal(searchResponse.Body, &searchResult)
	return searchResult, nil
}
