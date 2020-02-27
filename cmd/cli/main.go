package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ONSdigital/who-goes-there/pkg/github"
	"github.com/davecgh/go-spew/spew"
	"github.com/urfave/cli/v2"
)

var (
	// Version is set by build flags
	Version = "0.0.0"
)

func main() {
	// We need...:
	// - Github API token
	// - Organisation name
	//
	app := &cli.App{
		Name:    "who",
		Version: Version,
		Usage:   "cli interface for running Who Goes There",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "token",
				Aliases:  []string{"t"},
				Usage:    "github token for connecting to api",
				EnvVars:  []string{"API_TOKEN"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "org",
				Aliases:  []string{"o"},
				Usage:    "github organisation",
				EnvVars:  []string{"GITHUB_ORG_NAME"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "repo",
				Aliases:  []string{"r"},
				Usage:    "repository for compliance issues",
				EnvVars:  []string{"REPOSITORY"},
				Required: true,
			},
		},
		Commands: []*cli.Command{
			{
				Name:   "run",
				Usage:  "run basic report",
				Action: runCommand,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func runCommand(c *cli.Context) error {

	token := c.String("token")
	organisation := c.String("org")
	repository := c.String("repo")

	// token := ""
	// if token = os.Getenv("API_TOKEN"); len(token) == 0 {
	// 	fmt.Println("Missing API_TOKEN env var")
	// 	os.Exit(1)
	// }

	// organisation := ""
	// if organisation = os.Getenv("GITHUB_ORG_NAME"); len(organisation) == 0 {
	// 	fmt.Println("Missing GITHUB_ORG_NAME")
	// 	os.Exit(1)
	// }

	// repository := ""
	// if repository = os.Getenv("REPOSITORY"); len(repository) == 0 {
	// 	fmt.Println("Missing REPOSITORY")
	// 	os.Exit(1)
	// }

	fmt.Printf("Org: %s\nRepo: %s\n", organisation, repository)

	client := github.NewClient(token)

	users, err := client.FetchOrganizationMembers(organisation)
	if err != nil {
		log.Fatalf("Failed to retrieve members: %v", err)
	}

	// TODO: Temporary reporting of results - these obviously need to properly
	// 		 run through rules
	hasMFA, isAdmin, badAdmin := 0, 0, 0
	for _, user := range users {
		log.Printf("%15s %s\n", user.Node.Name, user.Node.Login)
		if user.HasTwoFactorEnabled {
			hasMFA++
		} else {
			log.Printf("%15s %s\n", user.Node.Name, user.Node.Login)
		}
		if user.Role == "ADMIN" {
			isAdmin++
		}
		if !user.HasTwoFactorEnabled && user.Role == "ADMIN" {
			badAdmin++
		}
	}
	log.Println("Users with MFA:", hasMFA)
	log.Println("Admin users:", isAdmin)
	log.Println("Bad admin:", badAdmin)

	issues, err := client.FetchAllOpenIssues(organisation, repository)
	if err != nil {
		log.Fatalf("Failed to retrieve issues: %v", err)
	}
	spew.Dump(issues)

	return nil
}
