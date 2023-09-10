package httpclient

import (
	"github.com/go-resty/resty/v2"
)

type RestClient struct {
	*resty.Client
}

func NewRestClient(baseUrl string) *RestClient {
	httpClient := resty.New()
	httpClient.SetBaseURL(baseUrl)

	return &RestClient{httpClient}
}
