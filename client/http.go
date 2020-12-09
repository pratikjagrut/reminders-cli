package client

import "net/http"

type HTTPClient struct {
	client *http.Client
	BackendURI string
}

func NewHttpClient(uri string) HTTPClient {
	return HTTPClient{
		BackendURI: uri,
		client: &http.Client{},
	}
}
