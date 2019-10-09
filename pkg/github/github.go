// Package github defines the interface for interacting with the Github graphql api
package github

import (
	"context"

	"github.com/machinebox/graphql"
)

type (

	// PageInfo represents the pagination information returned from the query
	PageInfo struct {
		EndCursor       string
		HasNextPage     bool
		HasPreviousPage bool
		StartCursor     string
	}
)

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

// Run calls the underlying graphql.Run() and automatically adds in appropriate
// authentication headers and background context
func (c Client) Run(request *graphql.Request, response interface{}) error {
	request.Header.Set("Authorization", "bearer "+c.token)
	return c.client.Run(context.Background(), request, response)
}
