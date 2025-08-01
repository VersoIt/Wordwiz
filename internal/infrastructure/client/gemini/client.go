package gemini

import (
	"net/http"
	"wordwiz/config"
)

type Client struct {
	cfg    config.Config
	client http.Client
}

func New(cfg config.Config, client http.Client) *Client {
	return &Client{
		cfg:    cfg,
		client: client,
	}
}
