package github

import (
	"context"

	"github.com/machinebox/graphql"
	"github.com/pkg/errors"
)

type (

	// Repository represents a github repository
	Repository struct {
		Issues struct {
			TotalCount int
			PageInfo   PageInfo
			Nodes      []Issue
		}
	}

	// Issue represents a github issue
	Issue struct {
		ID        string
		Title     string
		Body      string
		BodyText  string
		CreatedAt string
		Comments  struct {
			TotalCount int
			Nodes      []Comment
		}
	}

	// Comment represents a comment on an issue
	Comment struct {
		ID        string
		Body      string
		BodyText  string
		CreatedAt string
	}
)

// FetchAllOpenIssues will return all the issues currently open on the
// specified repository
func (c *Client) FetchAllOpenIssues(ctx context.Context, owner, repo string) ([]Issue, error) {

	var issues []Issue
	var next *string // Allow it to be nil

	req := graphql.NewRequest(`
	query($owner: String!, $repo: String!, $after: String) {
		repository(owner: $owner, name: $repo) {
			issues(after: $after, first: 100, states: [OPEN]) {
				totalCount
				pageInfo{
					hasNextPage
					endCursor
				}
				nodes{
					id
					title
					body
					createdAt
					comments(first: 10) {
						totalCount
						nodes {
							id
							body
							bodyText
							createdAt
						}
					}
				}
			}
		}
	}`)
	req.Var("owner", owner)
	req.Var("repo", repo)

	page := 0
	hasNextPage := true
	for hasNextPage {
		page++
		res := &struct{ Repository Repository }{}

		req.Var("after", next)

		if err := c.Run(ctx, req, &res); err != nil {
			return nil, errors.Wrap(err, "failed to fetch issues for repo")
		}
		issues = append(issues, res.Repository.Issues.Nodes...)
		next = &res.Repository.Issues.PageInfo.EndCursor
		hasNextPage = res.Repository.Issues.PageInfo.HasNextPage
	}

	return issues, nil
}
