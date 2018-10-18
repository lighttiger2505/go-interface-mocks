package main

import "testing"

type MockIssueClientSimple struct {
}

func (m *MockIssueClientSimple) GetIssue(repo string, id int) (*Issue, error) {
	return &Issue{
		ID:          12,
		Title:       "Title12",
		Description: "Description12",
	}, nil
}

func (m *MockIssueClientSimple) CreateIssue(repo, title, description string) (*Issue, error) {
	return nil, nil
}

func TestIssueCmd_RunGetIssue_Simple(t *testing.T) {
	ui := NewMockUI()
	cmd := &IssueCmd{
		UI:     ui,
		Client: &MockIssueClientSimple{},
	}
	cmd.RunGetIssue()

	got := ui.Out.String()
	want := "ID:12, Title:Title12, Description:Description12\n"
	if got != want {
		t.Errorf("invalid output, \ngot: %#v\nwant:%#v", got, want)
	}
}

type MockIssueClientSpecific struct {
	MockGetIssue    func(repo string, id int) (*Issue, error)
	MockCreateIssue func(repo, title, description string) (*Issue, error)
}

func (m *MockIssueClientSpecific) GetIssue(repo string, id int) (*Issue, error) {
	return m.MockGetIssue(repo, id)
}

func (m *MockIssueClientSpecific) CreateIssue(repo, title, description string) (*Issue, error) {
	return m.MockCreateIssue(repo, title, description)
}

func TestIssueCmd_RunGetIssue_Specific(t *testing.T) {
	ui := NewMockUI()
	cmd := &IssueCmd{
		UI: ui,
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

func TestIssueCmd_RunCreateIssue(t *testing.T) {
	ui := NewMockUI()
	cmd := &IssueCmd{
		UI: ui,
		Client: &MockIssueClientSpecific{
			MockCreateIssue: func(repo, title, description string) (*Issue, error) {
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
