package issues

import "testing"

func TestGetIssues(t *testing.T) {
	nwo := "hilli/issue-test"
	labels := "DSR-Request"

	client := NewClient()
	issues := client.GetOpenIssues(nwo, labels)

	for _, issue := range issues {
		t.Logf("Issue: %s", issue)
	}
}

func TestGetIssue(t *testing.T) {
	client := NewClient()
	issues := client.GetIssue("hilli", "issue-test", "DSR-Request")

	for _, issue := range issues {
		t.Logf("Issue: %s", issue)
	}
}

func TestGetStargazers(t *testing.T) {
	client := NewClient()
	client.GetStargazers("hilli", "issue-test")
}
