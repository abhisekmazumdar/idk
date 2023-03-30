package lando

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Run a Lando command after replacing "idk" with "lando"
func Run(args string) {
	fmt.Println("Running: lando " + args)
	command := exec.Command("lando", strings.TrimSpace(args))
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err := command.Run()
	if err != nil {
		fmt.Println(err)
	}
}
