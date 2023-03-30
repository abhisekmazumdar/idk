package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/abhisekmazumdar/idk/ddev"
	"github.com/abhisekmazumdar/idk/lando"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "idk",
		Usage:       "replace idk with lando OR ddev followed by the arguments & flags",
		Description: "idk - Helper CLI that knows for you and runs command as lando OR ddev.",
		Commands: []*cli.Command{
			{
				Name:    "composer",
				Aliases: []string{"c"},
				Usage:   "runs as ddev/lando composer ...",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("Run composer: ", cCtx.Args().First())
					return nil
				},
			},
			{
				Name:    "drush",
				Aliases: []string{"d"},
				Usage:   "runs as ddev/lando drush ...",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("Run drush: ", cCtx.Args().First())
					return nil
				},
			},
			{
				Name:    "require",
				Aliases: []string{"r"},
				Usage:   "runs as ddev/lando composer require ...",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("Run composer require: ", cCtx.Args().First())
					return nil
				},
			},
			{
				Name:    "site-install",
				Aliases: []string{"si"},
				Usage:   "runs as ddev/lando drush site-install ...",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("Run drush site-install: ", cCtx.Args().First())
					return nil
				},
			},
		},
		Authors: []*cli.Author{
			{Name: "Abhisek Mazumdar (abhisekmazumdar)"},
		},
		Action: func(cCtx *cli.Context) error {

			strArgs := ""

			strArgs = argsToString(cCtx.Args())
			fmt.Println("Checking which development tool configuration files are present...")

			localConfig := ""
			localConfig = checkLocalConfig()

			switch localConfig {
			case "lando":
				lando.Run(strArgs)
			case "ddev":
				ddev.Run(strArgs)
			default:
				fmt.Println("You 🫵 spoiled all the fun, Try Again!!.")
				os.Exit(1)
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func checkLocalConfig() string {
	reader := bufio.NewReader(os.Stdin)
	if _, err := os.Stat("./.ddev"); !os.IsNotExist(err) {
		if _, err := os.Stat("./.lando.yml"); !os.IsNotExist(err) {
			fmt.Print("Both .lando.yml and .ddev directories exist.\nWhich one would you like to use? Type 'lando' or 'ddev': ")
			option, _ := reader.ReadString('\n')
			option = strings.ToLower(strings.TrimSpace(option))
			return option
		} else {
			return "ddev"
		}
	} else if _, err := os.Stat("./.lando.yml"); !os.IsNotExist(err) {
		return "lando"
	} else {
		fmt.Println("No ddev or lando configuration files found.")
		os.Exit(1)
	}
	return ""
}

func argsToString(args cli.Args) string {
	var str strings.Builder
	for i := 0; i < args.Len(); i++ {
		str.WriteString(args.Get(i))
		str.WriteString(" ")
	}
	return str.String()
}
