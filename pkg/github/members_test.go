package github_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ONSdigital/who-goes-there/pkg/github"
	"github.com/stretchr/testify/assert"
)

func TestFetchOrganizationMembersBadResponse(t *testing.T) {

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `xxx`)
	}))
	defer srv.Close()

	// Reset the default URL to our local test server
	github.GithubAPIURL = srv.URL

	client := github.NewClient("SomeFakeToken")

	_, err := client.FetchOrganizationMembers("my-org")
	assert.True(t, strings.Contains(err.Error(), "invalid character 'x' looking for beginning of value"))
}
