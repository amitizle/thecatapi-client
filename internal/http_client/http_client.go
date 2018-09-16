package http_client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	defaultBaseURL = "https://api.thecatapi.com"
)

type HTTPClient struct {
	BaseURL   string
	ClientImp *http.Client
	ApiKey    string
}

type Request struct {
	Path        string
	BaseURL     string
	QueryParams map[string]string
	Headers     map[string]string
	httpClient  *HTTPClient
}

type Response struct {
	Body []byte
}

func NewClient() (*HTTPClient, error) {
	return &HTTPClient{
		BaseURL:   defaultBaseURL,
		ClientImp: http.DefaultClient, // Maybe support different client in the future
	}, nil
}

func (httpClient *HTTPClient) NewRequest() *Request {
	return &Request{
		Path:        "/",
		BaseURL:     httpClient.BaseURL,
		QueryParams: map[string]string{},
		Headers: map[string]string{
			"Content-Type": "application/json",
			"Accept":       "application/json",
			"X-Api-Key":    httpClient.ApiKey,
		},
		httpClient: httpClient,
	}
}

func addHeaders(httpRequest *http.Request, headers map[string]string) {
	for header, value := range headers {
		httpRequest.Header.Add(header, value)
	}
}

func (request *Request) Get() (*Response, error) {
	url, err := request.prepareURL()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", url, nil)
	addHeaders(req, request.Headers)
	resp, err := request.httpClient.ClientImp.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return &Response{Body: body}, nil
}

func (request *Request) Post() (*Response, error) {
	return &Response{}, nil
}

func (request Request) prepareURL() (string, error) {
	urlString := request.httpClient.BaseURL
	url, err := url.Parse(urlString)
	if err != nil {
		return "", err
	}
	url.Path = request.Path
	values := url.Query()
	for key, value := range request.QueryParams {
		values.Add(key, value)
	}
	url.RawQuery = values.Encode()
	fmt.Println(url.String())
	return url.String(), nil
}
