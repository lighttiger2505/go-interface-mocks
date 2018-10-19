package main

import (
	"bytes"
	"fmt"
	"testing"
)

//=======================================================================
// Step 1/3
// シンプルにMockするよ！！！
//=======================================================================

// MockUI UIの入出力をMockする構造体
type MockUI struct {
	Out *bytes.Buffer
}

// NewMockUI コンストラクタ
func NewMockUI() *MockUI {
	return &MockUI{
		Out: new(bytes.Buffer),
	}
}

// Println fmt.Printlnのラッパー
func (u *MockUI) Println(val string) {
	fmt.Fprint(u.Out, val+"\n")
}

// MockIssueClientSimple シンプルにMockするための構造体
type MockIssueClientSimple struct{}

// GetIssue 単純に固定でIssueを返す
func (m *MockIssueClientSimple) GetIssue(repo string, id int) (*Issue, error) {
	return &Issue{
		ID:          12,
		Title:       "Title12",
		Description: "Description12",
	}, nil
}

// CreateIssue 割愛
func (m *MockIssueClientSimple) CreateIssue(repo, title, description string) (*Issue, error) {
	return nil, nil
}

// TestIssueCmd_RunGetIssue_Simple 単純なMockのテスト
func TestIssueCmd_RunGetIssue_Simple(t *testing.T) {

	ui := NewMockUI()
	cmd := &IssueCmd{
		// Stdoutではなくbytes.Bufferを渡すMockUIを使うことで後で出力文字列を取れる
		UI: ui,
		// 実際のIssueを取得するクライアントではなくNewMockUIを
		Client: &MockIssueClientSimple{},
	}

	// 実行する
	cmd.RunGetIssue()

	// bytes.Bufferにしておけば気軽に出力文字列を取れる
	got := ui.Out.String()
	want := "ID:12, Title:Title12, Description:Description12\n"
	if got != want {
		t.Errorf("invalid output, \ngot: %#v\nwant:%#v", got, want)
	}
}

//=======================================================================
// いやいや固定値を返すだけじゃIssueの状態によって動作変更するパターン無理じゃん
//
// Step 2/3
// というわけで柔軟にMockするよ！！！
//=======================================================================

// MockIssueClientSpecific 柔軟にMockするための構造体
type MockIssueClientSpecific struct {
	// Mock用にInterfaceで定義した関数と同様の定義を持つ関数を持つのがミソ
	MockGetIssue    func(repo string, id int) (*Issue, error)
	MockCreateIssue func(repo, title, description string) (*Issue, error)
}

// GetIssue ここでは関数を定義せずにMock用関数をそのまま実行する
func (m *MockIssueClientSpecific) GetIssue(repo string, id int) (*Issue, error) {
	return m.MockGetIssue(repo, id)
}

// CreateIssue 同上
func (m *MockIssueClientSpecific) CreateIssue(repo, title, description string) (*Issue, error) {
	return m.MockCreateIssue(repo, title, description)
}

// TestIssueCmd_RunGetIssue_Specific テスト毎にGetIssueの動作を変更する柔軟なMockのテスト
func TestIssueCmd_RunGetIssue_Specific(t *testing.T) {
	ui := NewMockUI()
	cmd := &IssueCmd{
		UI: ui,
		// Mockする関数に対して関数を定義して渡す
		Client: &MockIssueClientSpecific{
			MockGetIssue: func(repo string, id int) (*Issue, error) {
				return &Issue{
					ID:          12,
					Title:       "Title12",
					Description: "Description12",
				}, nil
			},
		},
	}

	cmd.RunGetIssue()

	got := ui.Out.String()
	want := "ID:12, Title:Title12, Description:Description12\n"
	if got != want {
		t.Errorf("invalid output, \ngot: %#v\nwant:%#v", got, want)
	}
}

//=======================================================================
// いやいやちゃんと渡した値が想定どおりの値になっていることを保証出来なきゃ駄目でしょ
//
// Step 3/3
// というわけでかっちりMockするよ！！！
//=======================================================================

// TestIssueCmd_RunGetIssue_Specific テスト毎にCreateIssueの動作を変更し、値の検証まで行う
func TestIssueCmd_RunCreateIssue(t *testing.T) {
	ui := NewMockUI()
	cmd := &IssueCmd{
		UI: ui,
		Client: &MockIssueClientSpecific{
			MockCreateIssue: func(repo, title, description string) (*Issue, error) {
				// testing.Tがこのスコープにおいても有効であるので
				// 引数に対してそれぞれの値が想定どおりであることを確認する
				wantTitle := "newtitle"
				if title != wantTitle {
					t.Errorf("invalid title \ngot: %s\nwant:%s", title, wantTitle)
				}
				wantDescription := "newdescription"
				if description != wantDescription {
					t.Errorf("invalid description \ngot: %s\nwant:%s", description, wantDescription)
				}
				return &Issue{
					ID: 12,
				}, nil
			},
		},
	}
	cmd.RunCreateIssue()

	got := ui.Out.String()
	want := "created issue ID:12\n"
	if got != want {
		t.Errorf("invalid output, \ngot: %#v\nwant:%#v", got, want)
	}
}

//=======================================================================
// ご清聴ありがとうございました
//=======================================================================
