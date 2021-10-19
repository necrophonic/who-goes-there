package github_test

import (
	"context"
	"testing"

	"github.com/ONSdigital/who-goes-there/pkg/github"
	"github.com/machinebox/graphql"
)

func TestRun(t *testing.T) {

	client := github.NewClient("SomeFakeToken")

	ctx := context.Background()

	var resp interface{}
	err := client.Run(ctx, graphql.NewRequest(`sheep`), resp)
	if err != nil {
		t.Error(err)
	}

}
