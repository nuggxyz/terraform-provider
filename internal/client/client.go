package client

type Client struct {
}

func NewClient() (*Client, error) {
	return &Client{}, nil
}
