package utils

import (
	"os"
	"os/exec"
)

//ClearScreen clears the screen using a windows command
func ClearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
