package thirdparty

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// GitHubReposInfo represents the minimal repository information we need
type GitHubReposInfo struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Private    bool   `json:"private"`
	Watchers   int    `json:"watchers"`
	OpenIssues int    `json:"open_issues"`
}

type GitHubAPI struct {
	BaseURL string
	Client  *http.Client
}

func NewGitHubAPI() *GitHubAPI {
	return &GitHubAPI{
		BaseURL: "https://api.github.com",
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (g *GitHubAPI) FetchRepositories(username string) ([]GitHubReposInfo, error) {
	url := fmt.Sprintf("%s/users/%s/repos", g.BaseURL, username)

	resp, err := g.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch repositories from GitHub")
	}

	// Decode full GitHub API response into a structured slice
	var repos []struct {
		ID         int64  `json:"id"`
		Name       string `json:"name"`
		Private    bool   `json:"private"`
		Watchers   int    `json:"watchers_count"`
		OpenIssues int    `json:"open_issues_count"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		return nil, err
	}

	// Convert to GitHubReposInfo slice
	var reposInfo []GitHubReposInfo
	for _, repo := range repos {
		reposInfo = append(reposInfo, GitHubReposInfo{
			ID:         repo.ID,
			Name:       repo.Name,
			Private:    repo.Private,
			Watchers:   repo.Watchers,
			OpenIssues: repo.OpenIssues,
		})
	}

	return reposInfo, nil
}
