package main

import "fmt"

// IssueCmd Issueの操作を実行するコマンド
// インタフェースにて任意の振る舞いを注入する
type IssueCmd struct {
	// UI Stdin/Stdoutによるコマンド入力/出力するインタフェース
	UI UI
	// Client GitLabAPIにリクエストを送りIssueを操作するインタフェース
	Client IssueClient
}

// RunGetIssue Issue取得を行う処理
func (c *IssueCmd) RunGetIssue() {
	// Issueの取得
	issue, _ := c.Client.GetIssue("foo/bar", 12)

	// Issueの情報をStdoutに出力
	res := fmt.Sprintf("ID:%d, Title:%s, Description:%s", issue.ID, issue.Title, issue.Description)
	c.UI.Println(res)
}

// RunCreateIssue Issue更新を行う処理
func (c *IssueCmd) RunCreateIssue() {
	// Issueの作成
	issue, _ := c.Client.CreateIssue("foo/bar", "newtitle", "newdescription")

	// 作成したIssueの情報を出力(IDのみ)
	res := fmt.Sprintf("created issue ID:%d", issue.ID)
	c.UI.Println(res)
}

// Issue Issueのデータを格納する構造体
type Issue struct {
	// ID ID
	ID int
	// Title タイトル
	Title string
	// Description タイトル
	Description string
}

// IssueClient Issueの操作を行うインタフェース
type IssueClient interface {
	// GetIssue 指定されたIDのIssueを取得
	GetIssue(repo string, id int) (*Issue, error)
	// CreateIssue 指定されたIDのIssueを取得
	CreateIssue(repo, title, description string) (*Issue, error)
}
