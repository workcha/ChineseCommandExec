package main

import "github.com/workcha/ChineseCommandExec/command"

func main() {
	result := command.CommonExecute("ipconfig /all")
	println(result)
}
