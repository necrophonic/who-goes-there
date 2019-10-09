package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ONSdigital/who-goes-there/pkg/github"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	// We need...:
	// - Github API token
	// - Organisation name
	//
	token := ""
	if token = os.Getenv("API_TOKEN"); len(token) == 0 {
		log.Fatal("Missing API_TOKEN env var")
	}

	organisation := ""
	if organisation = os.Getenv("GITHUB_ORG_NAME"); len(organisation) == 0 {
		log.Fatal("Missing GITHUB_ORG_NAME")
	}

	repository := ""
	if repository = os.Getenv("REPOSITORY"); len(repository) == 0 {
		log.Fatal("Missing REPOSITORY")
	}

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

}
