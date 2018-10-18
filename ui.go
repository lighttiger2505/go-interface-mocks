package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type UI interface {
	Println(string)
	Scan() (string, error)
}

type BasicUI struct {
	In  io.Reader
	Out io.Writer
}

func NewBasicUI() *BasicUI {
	return &BasicUI{
		In:  os.Stdin,
		Out: os.Stdout,
	}
}

func (u *BasicUI) Println(val string) {
	fmt.Fprint(u.Out, val+"\n")
}

func (u *BasicUI) Scan() (string, error) {
	b, _ := ioutil.ReadAll(u.In)
	return string(b), nil
}

type UICmd struct {
	UI UI
}

func (c *UICmd) RunPrint() {
	c.UI.Println("hogehoge")
}

func (c *UICmd) RunScan() {
	res, _ := c.UI.Scan()
	c.UI.Println(res)
}
