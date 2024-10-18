package gissues

import (
	"context"
	"os"

	"github.com/google/go-github/v66/github"
)

type IssuesClient struct {
	client *github.Client
}

func NewClient() IssuesClient {
	issuesclient := IssuesClient{
		client: github.NewClient(nil).WithAuthToken(os.Getenv("GITHUB_TOKEN")),
	}
	return issuesclient
}

func (c *IssuesClient) GetIssues(owner, repo string, labels []string) ([]*github.Issue, error) {
	issues, _, err := c.client.Issues.ListByRepo(context.Background(), owner, repo, &github.IssueListByRepoOptions{
		State:       "open",
		Labels:      labels,
		ListOptions: github.ListOptions{PerPage: 10},
	})
	if err != nil {
		return nil, err
	}
	return issues, nil
}

func (c *IssuesClient) GetIssueById(owner, repo string, number int) (*github.Issue, error) {
	issue, _, err := c.client.Issues.Get(context.Background(), owner, repo, number)
	if err != nil {
		return nil, err
	}
	return issue, nil
}

func (c *IssuesClient) CreateIssue(owner, repo, title, body string, labels []string) (*github.Issue, error) {
	issueRequest := &github.IssueRequest{
		Title:  &title,
		Body:   &body,
		Labels: &labels,
		// Assignee: &owner, // Only one of assignee or assignees can be set
		Assignees: &[]string{
			owner,
		},
	}
	issue, _, err := c.client.Issues.Create(context.Background(), owner, repo, issueRequest)
	if err != nil {
		return nil, err
	}
	return issue, nil
}

func (c *IssuesClient) CloseIssue(owner, repo string, number int) error {
	closed := "closed"
	issueRequest := &github.IssueRequest{
		State: &closed,
	}
	_, _, err := c.client.Issues.Edit(context.Background(), owner, repo, number, issueRequest)
	return err
}

func (c *IssuesClient) CommentIssue(owner, repo string, number int, comment string) error {
	_, _, err := c.client.Issues.CreateComment(context.Background(), owner, repo, number, &github.IssueComment{
		Body: &comment,
	})
	return err
}
