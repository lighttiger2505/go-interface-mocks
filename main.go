package main

func main() {
	uicmd := &UICmd{UI: NewBasicUI()}
	uicmd.RunPrint()

	commandcmd := &CommandCmd{UI: NewBasicUI()}
	commandcmd.RunExec()
}
