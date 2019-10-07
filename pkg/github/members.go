package github

import (
	"context"
	"log"

	"github.com/machinebox/graphql"
	"github.com/pkg/errors"
)

type (
	// Organization represents the top level information returned from the
	// graphql query
	Organization struct {
		MembersWithRole struct {
			TotalCount int
			PageInfo   PageInfo
			Edges      []User
		}
	}

	// PageInfo represents the pagination information returned from the query
	PageInfo struct {
		EndCursor       string
		HasNextPage     bool
		HasPreviousPage bool
		StartCursor     string
	}

	// User represents the information returned for a specific user
	User struct {
		HasTwoFactorEnabled bool
		Role                string
		Node                struct {
			Name  string
			Login string
		}
	}

	organizationResponse struct {
		Organization Organization
	}
)

// FetchOrganizationMembers performs a graphql query to fetch the member information
// for a given organization
func (c *Client) FetchOrganizationMembers(org string) ([]User, error) {

	var users []User
	var next *string // Allow it to be nil\

	req := graphql.NewRequest(`
	query($organization: String!, $after: String) {
		organization(login: $organization){
			membersWithRole(after: $after, first: 100){
				totalCount
				pageInfo{
					hasNextPage
					endCursor
				}
				edges{
					hasTwoFactorEnabled
					role
					node{
						name
						login
					}
				}
			}
		}
	}`)
	req.Var("organization", org)
	req.Header.Set("Authorization", "bearer "+c.token)

	page := 0
	for {
		page++
		log.Printf("Fetch page %d", page)
		res := &organizationResponse{}

		req.Var("after", next)

		if err := c.client.Run(context.Background(), req, &res); err != nil {
			return nil, errors.Wrap(err, "failed to fetch members for organisation")
		}
		users = append(users, res.Organization.MembersWithRole.Edges...)
		next = &res.Organization.MembersWithRole.PageInfo.EndCursor

		if next == nil || *next == "" {
			break
		}
	}

	return users, nil
}
