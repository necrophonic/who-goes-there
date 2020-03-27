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

// GraphQLSpec contains the values required for connecting to the GraphQL api
type GraphQLSpec struct {
	Token        string `required:"true"`
	Organisation string `required:"true"`
	// Repository   string `envconfig:"GITHUB_REPOSITORY"`
}

// Handler recieves a CloudWatch Event and triggers the rules engine
func Handler(ctx context.Context, cwEvent events.CloudWatchEvent) (*report.Report, error) {

	var g GraphQLSpec
	err := envconfig.Process("GITHUB", &g)
	if err != nil {
		return nil, err
	}

	client := github.NewClient(g.Token)
	users, err := client.FetchOrganizationMembers(g.Organisation)
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve members")
	}

	report := report.Report{}

	for _, user := range users {
		report.Summary.TotalUsers++
		if !user.HasTwoFactorEnabled {
			report.Summary.UsersMissingMFA++
		}
		if user.Role == "ADMIN" {
			report.Summary.AdminUsers++
			if !user.HasTwoFactorEnabled {
				report.Summary.AdminUsersFailingRules++
			}
		}
	}

	// This service is defined with a destination that puts our response onto
	// an SQS queue. The content of *report gets automatically unrolled into
	// the SQS payload.
	// This means we don't need to manually connect to and send the message to
	// a queue!
	log.Println("Publishing report")
	return &report, nil
}

func main() {
	lambda.Start(Handler)
}
