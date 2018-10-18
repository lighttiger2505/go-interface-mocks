package main

import (
	"bytes"
	"testing"
)

func TestUICmd_RunPrint(t *testing.T) {
	ui := NewMockUI()
	cmd := &UICmd{UI: ui}
	cmd.RunPrint()

	got := ui.Out.String()
	want := "hogehoge\n"
	if got != want {
		t.Errorf("invalid output, \ngot=%s\nwant=%s", got, want)
	}
}

func TestUICmd_RunScan(t *testing.T) {
	ui := NewMockUI()
	ui.In = bytes.NewBufferString("hogehoge")
	cmd := &UICmd{UI: ui}
	cmd.RunScan()

	got := ui.Out.String()
	want := "hogehoge\n"
	if got != want {
		t.Errorf("invalid output, \ngot: %#v\nwant:%#v", got, want)
	}
}
