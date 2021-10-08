package main

import (
	"context"
	"log"

	"github.com/ONSdigital/who-goes-there/pkg/github"
	"github.com/ONSdigital/who-goes-there/pkg/report"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

// GraphQLSpec contains the values required for connecting to the GraphQL api.
// See README for required token scopes
type GraphQLSpec struct {
	Token        string `envconfig:"TOKEN" required:"true"`
	Organisation string `envconfig:"ORGANISATION" required:"true"`
	Repository   string `envconfig:"REPOSITORY"`
}

// Handler recieves a CloudWatch Event and triggers the rules engine
func Handler(ctx context.Context, cwEvent events.CloudWatchEvent) (*report.Report, error) {

	// Import the environment variables using envconfig
	// Ideally we'd make sure we only need to do this once if the lambda is
	// already pre-warmed, but for our purposes given the infrequent run
	// of this lambda this is adequate and less complex.
	var g GraphQLSpec
	err := envconfig.Process("GITHUB", &g)
	if err != nil {
		return nil, err
	}

	// Connect to github and fetch all the members for the configured organisation
	client := github.NewClient(g.Token)
	users, err := client.FetchOrganizationMembers(g.Organisation)
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve members")
	}

	// Create a basic summary report - right now this is based on
	// whether users are normal or admin and whether they have
	// MFA enabled.
	// We could run a multitude of rules here (ideally via a pluggable
	// rules engine).
	rep := report.New()
	for _, user := range users {
		rep.Summary.TotalUsers++
		if !user.HasTwoFactorEnabled {
			rep.Summary.UsersMissingMFA++
		}
		if user.Role == "ADMIN" {
			rep.Summary.AdminUsers++
			if !user.HasTwoFactorEnabled {
				rep.Summary.AdminUsersFailingRules++
			}
		}
	}

	// This service is defined with an on_success destination that puts our
	// response onto an SQS queue. The content of *report gets automatically
	// unrolled into the SQS payload.
	// This means we don't need to manually connect to and send the message to
	// a queue!
	log.Println("Publishing report")
	return rep, nil
}

func main() {
	lambda.Start(Handler)
}
