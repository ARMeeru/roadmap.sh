package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const baseURL = "https://api.github.com/users/"

type GitHubEvent struct {
	Type string `json:"type"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
}

// FetchGitHubActivity retrieves the public GitHub events for a specified user.
//
// Parameters:
//   - username: The GitHub username for which to fetch the activity.
//
// Returns:
//   - A slice of GitHubEvent structs containing the user's public events.
//   - An error if there was an issue fetching or decoding the data.
//
// Possible errors:
//   - If the GitHub API is unreachable.
//   - If the user is not found (HTTP 404).
//   - If there is any other HTTP error.
//   - If there is an error decoding the JSON response.
func FetchGitHubActivity(username string) ([]GitHubEvent, error) {
	url := fmt.Sprintf("%s%s/events", baseURL, username)
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("unable to reach GitHub: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("user not found")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data, status code: %d", resp.StatusCode)
	}

	var events []GitHubEvent
	err = json.NewDecoder(resp.Body).Decode(&events)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return events, nil
}