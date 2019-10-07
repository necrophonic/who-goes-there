package main

import (
	"log"
	"os"

	"github.com/ONSdigital/who-goes-there/pkg/github"
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

	organisationName := ""
	if organisationName = os.Getenv("GITHUB_ORG_NAME"); len(organisationName) == 0 {
		log.Fatal("Missing GITHUB_ORG_NAME")
	}

	client := github.NewClient(token)

	users, err := client.FetchOrganizationMembers(organisationName)
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

}
