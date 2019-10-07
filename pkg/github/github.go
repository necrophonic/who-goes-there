// package github defines the interface for interacting with the Github graphql api
package github

import "github.com/machinebox/graphql"

// Various fixed values for interacting with the Github graphql api
const (
	APIURL = "https://api.github.com/graphql"
)

// Client wraps a graphql client for communicating with the Github API
type Client struct {
	token  string
	client *graphql.Client
}

// NewClient instansiates a new graphql client
func NewClient(token string) *Client {
	return &Client{
		token:  token,
		client: graphql.NewClient(APIURL),
	}
}
