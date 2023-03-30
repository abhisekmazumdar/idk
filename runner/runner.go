package runner

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Run command as lando OR ddev.
func Run(devTool string, args string) {
	fmt.Println("Running: " + devTool + " " + args)
	command := exec.Command(devTool, strings.TrimSpace(args))
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err := command.Run()
	if err != nil {
		fmt.Println(err)
	}
}
