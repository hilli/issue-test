package issues

import (
	"context"
	"log"
	"os"
	"time"

	abs "github.com/microsoft/kiota-abstractions-go"
	octokit "github.com/octokit/go-sdk/pkg"
	octokitIssues "github.com/octokit/go-sdk/pkg/github/issues"
	"github.com/octokit/go-sdk/pkg/github/models"
	repos "github.com/octokit/go-sdk/pkg/github/repos"
)

type IssuesClient struct {
	client  *octokit.Client
	headers *abs.RequestHeaders
}

func NewClient() IssuesClient {
	client, err := octokit.NewApiClient(
		octokit.WithTokenAuthentication(os.Getenv("GITHUB_TOKEN")),
		octokit.WithRequestTimeout(5*time.Second),
		octokit.WithBaseUrl("https://api.github.com/"),
	)
	if err != nil {
		log.Fatalf("error creating client: %v", err)
	}

	// create headers that accept json back; GitHub's OpenAPI definition says
	// octet-stream but that's not actually what the API returns in this case
	headers := abs.NewRequestHeaders()
	_ = headers.TryAdd("Accept", "application/vnd.github.v3+json")

	return IssuesClient{
		client:  client,
		headers: headers,
	}
}

// List all open issues for a repository with the `DSR-Request` label.
func (c *IssuesClient) GetOpenIssues(nwo, labels string) []models.Issueable {
	open := octokitIssues.OPEN_GETSTATEQUERYPARAMETERTYPE

	issuesRequestConfig := &abs.RequestConfiguration[octokitIssues.IssuesRequestBuilderGetQueryParameters]{
		QueryParameters: &octokitIssues.IssuesRequestBuilderGetQueryParameters{
			State:  &open,
			Labels: &labels,
		},
		Headers: c.headers,
	}

	issuesList, err := c.client.Issues().WithUrl("https://api.github.com/repos/"+nwo+"/issues").Get(context.Background(), issuesRequestConfig)
	if err != nil {
		log.Fatalf("error getting issues: %v\n", err)
	}
	return issuesList
}

func (c *IssuesClient) GetIssue(owner, repo, labels string) []models.Issueable {
	req := &abs.RequestConfiguration[repos.ItemItemIssuesRequestBuilderGetQueryParameters]{
		QueryParameters: &repos.ItemItemIssuesRequestBuilderGetQueryParameters{
			Labels: &labels,
		},
		Headers: c.headers,
	}

	issues, err := c.client.Repos().ByOwnerId(owner).ByRepoId(repo).Issues().Get(context.Background(), req)
	if err != nil {
		log.Fatalf("error getting issues: %+v\n", err)
	}
	return issues
}

func (c *IssuesClient) GetStargazers(owner, repo string) {
	var zero int32 = 0
	var hundred int32 = 100
	stargazersReq := &abs.RequestConfiguration[repos.ItemItemStargazersRequestBuilderGetQueryParameters]{
		QueryParameters: &repos.ItemItemStargazersRequestBuilderGetQueryParameters{
			Per_page: &hundred,
			Page:     &zero,
		},
		Headers: c.headers,
	}

	stargazers, err := c.client.Repos().ByOwnerId(owner).ByRepoId(repo).Stargazers().Get(context.Background(), stargazersReq)
	if err != nil {
		log.Fatalf("error getting stargazers: %v", err.Error())
	}
	sgz := stargazers.GetItemItemStargazersStargazer()
	log.Printf("Stargazers: %+v", sgz)

}
