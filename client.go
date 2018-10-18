package main

import "fmt"

type Issue struct {
	ID          int
	Title       string
	Description string
}

type IssueClient interface {
	GetIssue(repo string, id int) (*Issue, error)
	CreateIssue(repo, title, description string) (*Issue, error)
}

type IssueCmd struct {
	UI     UI
	Client IssueClient
}

func (c *IssueCmd) RunGetIssue() {
	issue, _ := c.Client.GetIssue("foo/bar", 12)
	res := fmt.Sprintf("ID:%d, Title:%s, Description:%s", issue.ID, issue.Title, issue.Description)
	c.UI.Println(res)
}

func (c *IssueCmd) RunCreateIssue() {
	issue, _ := c.Client.CreateIssue("foo/bar", "newtitle", "newdescription")
	res := fmt.Sprintf("created issue ID:%d", issue.ID)
	c.UI.Println(res)
}
