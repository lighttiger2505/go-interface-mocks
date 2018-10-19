package main

import (
	"fmt"
	"io"
	"os"
)

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

// UI Stdin/Stdoutによるコマンド入力/出力するインタフェース
type UI interface {
	Println(string)
}

// BasicUI 単純なUI
type BasicUI struct {
	Out io.Writer
}

// NewBasicUI コンストラクタ
func NewBasicUI() UI {
	// 単純にStdoutで初期化
	return &BasicUI{
		Out: os.Stdout,
	}
}

// Println fmt.Printlnのラッパー
func (u *BasicUI) Println(val string) {
	fmt.Fprint(u.Out, val+"\n")
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

// GitlabIssueClient GitlabのIssueを取得する構造体
type GitlabIssueClient struct{}

// NewGitlabIssueClient コンストラクタ
func NewGitlabIssueClient() IssueClient {
	return &GitlabIssueClient{}
}

// GetIssue 割愛(実際にリクエストするコードがあると思ってください)
func (m *GitlabIssueClient) GetIssue(repo string, id int) (*Issue, error) {
	return &Issue{}, nil
}

// CreateIssue 割愛(実際にリクエストするコードがあると思ってください)
func (m *GitlabIssueClient) CreateIssue(repo, title, description string) (*Issue, error) {
	return &Issue{}, nil
}

func main() {
	// 実行するコマンドを作成
	issueCmd := &IssueCmd{
		UI:     NewBasicUI(),
		Client: NewGitlabIssueClient(),
	}
	// Issue取得を実行
	issueCmd.RunGetIssue()
}
