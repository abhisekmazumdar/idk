package runner

import (
	"fmt"
	"os"
	"os/exec"
)

// Run command as lando OR ddev.
func Run(devTool string, args ...string) {

	msg := "Running: " + devTool + " "
	for i := 0; i < len(args); i++ {
		msg += args[i] + " "
	}
	fmt.Println(msg)
	fmt.Println()

	command := exec.Command(devTool, args...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err := command.Run()

	if err != nil {
		fmt.Println(err)
	}
}
