package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

type MockUI struct {
	In  io.Reader
	Out *bytes.Buffer
}

func NewMockUI() *MockUI {
	return &MockUI{
		Out: new(bytes.Buffer),
	}
}

func (u *MockUI) Println(val string) {
	fmt.Fprint(u.Out, val+"\n")
}

func (u *MockUI) Scan() (string, error) {
	b, err := ioutil.ReadAll(u.In)
	if err != nil {
		return "", fmt.Errorf("failed read stdin. %s", err)
	}
	return string(b), nil
}
