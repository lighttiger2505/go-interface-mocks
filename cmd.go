package main

import (
	"os/exec"
	"strings"
)

// For os/exec test
var execCommand = exec.Command

// CommandCmd sample command mockking
type CommandCmd struct {
	UI UI
}

// RunExec run command exec
func (c *CommandCmd) RunExec() {
	output, _ := execCommand("ls").CombinedOutput()
	files := strings.Fields(string(output))
	for _, val := range files {
		c.UI.Println(val)
	}
}
