// Package github handles interacting with the Github graphql api
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

var (
	// GithubAPIURL is the default api endpoint
	GithubAPIURL = "https://api.github.com/graphql"
)

// Client wraps a graphql client for communicating with the Github API
type Client struct {
	token string
	q     *graphql.Client
}

// NewClient instansiates a new graphql client
func NewClient(token string) *Client {
	return &Client{
		token: token,
		q:     graphql.NewClient(GithubAPIURL),
	}
}

// Run calls the underlying graphql.Run() and automatically adds in appropriate
// authentication headers and background context
func (c Client) Run(ctx context.Context, request *graphql.Request, response interface{}) error {
	request.Header.Set("Authorization", "bearer "+c.token)
	err := c.q.Run(ctx, request, response)
	return err
}
