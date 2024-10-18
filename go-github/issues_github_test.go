package gissues

import "testing"

var (
	owner = "hilli"
	repo  = "issue-test"
)

func TestCreateIssue(t *testing.T) {
	client := NewClient()
	issue, err := client.CreateIssue(owner, repo, "New issue", "New issue body", []string{"DSR-Request", "testing"})
	if err != nil {
		t.Fatalf("error creating issue: %v", err)
	}

	t.Logf("Issue: %s", issue)

	// Delete the issue
	err = client.CloseIssue(owner, repo, issue.GetNumber())
	if err != nil {
		t.Fatalf("error deleting issue: %v", err)
	}
}

func TestGetIssues(t *testing.T) {
	client := NewClient()
	issues, err := client.GetIssues(owner, repo, []string{"DSR-Request"})
	if err != nil {
		t.Fatalf("error getting issues: %v", err)
	}

	for _, issue := range issues {
		t.Logf("Issue: %s", issue)
	}
}

func TestCloseIssue(t *testing.T) {
	client := NewClient()
	issue, err := client.CreateIssue(owner, repo, "New issue", "New issue body for closing", []string{"DSR-Request", "testing"})
	if err != nil {
		t.Fatalf("error creating issue: %v", err)
	}

	t.Logf("Issue: %s", issue)

	// Delete the issue
	err = client.CloseIssue(owner, repo, issue.GetNumber())
	if err != nil {
		t.Fatalf("error deleting issue: %v", err)
	}
}

func TestCleanupAll(t *testing.T) {
	client := NewClient()
	issues, err := client.GetIssues(owner, repo, []string{"testing"})
	if err != nil {
		t.Fatalf("error getting issues: %v", err)
	}

	for _, issue := range issues {
		// Delete the issue
		err = client.CloseIssue(owner, repo, issue.GetNumber())
		if err != nil {
			t.Fatalf("error deleting issue: %v", err)
		}
	}
}

func TestCreateComment(t *testing.T) {
	client := NewClient()
	issue, err := client.CreateIssue(owner, repo, "New issue", "New issue body", []string{"comment-issue", "testing"})
	if err != nil {
		t.Fatalf("error creating issue: %v", err)
	}

	t.Logf("Issue: %s", issue.GetURL())

	// Create a comment
	err = client.CommentIssue(owner, repo, issue.GetNumber(), "This is a comment")
	if err != nil {
		t.Fatalf("error creating comment: %v", err)
	}
}
