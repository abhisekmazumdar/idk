package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/abhisekmazumdar/idk/ddev"
	"github.com/abhisekmazumdar/idk/lando"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Checking which development tool configuration files are present...")

	if _, err := os.Stat("./.ddev"); !os.IsNotExist(err) {
		if _, err := os.Stat("./.lando.yml"); !os.IsNotExist(err) {
			fmt.Print("Both .lando.yml and .ddev directories exist.\nWhich one would you like to use? Type 'lando' or 'ddev': ")
			option, _ := reader.ReadString('\n')
			option = strings.ToLower(strings.TrimSpace(option))
			switch option {
			case "lando":
				runLandoCommand(os.Args[1:])
			case "ddev":
				runDdevCommand(os.Args[1:])
			default:
				fmt.Println("Invalid option. Exiting.")
				os.Exit(1)
			}
		} else {
			runDdevCommand(os.Args[1:])
		}
	} else if _, err := os.Stat("./.lando.yml"); !os.IsNotExist(err) {
		runLandoCommand(os.Args[1:])
	} else {
		fmt.Println("No ddev or lando configuration files found.")
	}
}

func runLandoCommand(args []string) {
	fmt.Println("Running as lando.")
	lando.Run(strings.Replace(strings.Join(args, " "), "idk", "", -1))
}

func runDdevCommand(args []string) {
	fmt.Println("Running as ddev.")
	ddev.Run(strings.Replace(strings.Join(args, " "), "idk", "", -1))
}
