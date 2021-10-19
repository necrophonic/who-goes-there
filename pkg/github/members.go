package github

import (
	"context"

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
			Edges      []OrganizationMemberEdge
		}
	}

	// OrganizationMemberEdge represents a user edge
	OrganizationMemberEdge struct {
		HasTwoFactorEnabled bool
		Role                string
		Node                User
	}

	// User represents the information returned for a specific user
	User struct {
		Name  string
		Login string
	}
)

// FetchOrganizationMembers performs a graphql query to fetch the member information
// for a given organization
func (c *Client) FetchOrganizationMembers(ctx context.Context, org string) ([]OrganizationMemberEdge, error) {

	var users []OrganizationMemberEdge
	var next *string // Allow it to be nil

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

	page := 0
	hasNextPage := true
	for hasNextPage {
		page++
		res := &struct {
			Organization Organization
		}{}

		req.Var("after", next)

		if err := c.Run(ctx, req, &res); err != nil {
			return nil, errors.Wrap(err, "failed to fetch members for organisation")
		}

		users = append(users, res.Organization.MembersWithRole.Edges...)
		next = &res.Organization.MembersWithRole.PageInfo.EndCursor
		hasNextPage = res.Organization.MembersWithRole.PageInfo.HasNextPage
	}

	return users, nil
}
