package processor

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/julioc98/robin/pkg/client"
)

type Client struct {
	token  string
	client client.Requester
}

func NewClient(token string, client client.Requester) *Client {
	return &Client{
		token:  token,
		client: client,
	}
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	id := uuid.New()
	req.Header.Add("accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("IdempotencyKey", id.String())
	req.Header.Add("API-Key", c.token)
	return c.client.Do(req)
}
