package github_test

import (
	"testing"

	"github.com/ONSdigital/who-goes-there/pkg/github"
	"github.com/machinebox/graphql"
)

func TestRun(t *testing.T) {

	client := github.NewClient("SomeFakeToken")

	var resp interface{}
	err := client.Run(graphql.NewRequest(`sheep`), resp)
	if err != nil {
		t.Error(err)
	}

}
