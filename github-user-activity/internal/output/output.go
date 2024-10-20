package output

import (
	"fmt"

	"github.com/ARMeeru/github-user-activity/internal/github"
)

// DisplayActivity displays the GitHub activity in a human-readable format
func DisplayActivity(events []github.GitHubEvent) {
	if len(events) == 0 {
		fmt.Println("No recent activity found.")
		return
	}

	fmt.Println("Recent GitHub activity:")
	for _, event := range events {
		switch event.Type {
		case "PushEvent":
			fmt.Printf("- Pushed to repository %s\n", event.Repo.Name)
		case "IssuesEvent":
			fmt.Printf("- Opened an issue in repository %s\n", event.Repo.Name)
		case "WatchEvent":
			fmt.Printf("- Starred repository %s\n", event.Repo.Name)
		default:
			fmt.Printf("- %s in repository %s\n", event.Type, event.Repo.Name)
		}
	}
}