package requests

const ()

type Client struct {
	Url  string
	Host string
	Key  string
}

func NewClient(url, host, key string) *Client {
	return &Client{
		Url:  url,
		Host: host,
		Key:  key,
	}
}
