package cmd

import (
	"fmt"
	"os"

	"github.com/ARMeeru/github-user-activity/internal/github"
	"github.com/ARMeeru/github-user-activity/internal/output"
)

// Execute is the main entry point for the CLI
func Execute() {
	// Ensure a username is provided as an argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: Provide a GitHub username as an argument")
		os.Exit(1)
	}

	username := os.Args[1]
	events, err := github.FetchGitHubActivity(username)

	if err != nil {
		fmt.Printf("Error fetching activity: %v\n", err)
		os.Exit(1)
	}

	output.DisplayActivity(events)
}