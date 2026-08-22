package parser

import (
	"net/http"
	"net/url"
	"time"
)

type ParserClient struct {
	http.Client
}

// TODO Transport config
func NewParserClient() *ParserClient {
	proxyURL, err := url.Parse("https://46.47.197.210:3128")
	if err != nil {
		panic("invalid proxy config")
	}

	return &ParserClient{
		Client: http.Client{
			Timeout: 15 * time.Second,
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyURL),
			},
		},
	}
}
