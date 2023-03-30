package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/abhisekmazumdar/idk/runner"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "idk",
		Usage:       "replace idk with dev tool followed by the arguments & flags",
		Description: "idk - Helper CLI that knows for you and runs command as lando OR ddev.",
		Commands: []*cli.Command{
			{
				Name:    "composer",
				Aliases: []string{"c"},
				Usage:   "runs as ddev/lando composer ...",
				Action:  cmdComposer,
			},
			{
				Name:    "drush",
				Aliases: []string{"d"},
				Usage:   "runs as ddev/lando drush ...",
				Action:  cmdDrush,
			},
			{
				Name:  "create-project",
				Usage: "runs as ddev/lando composer create-project ...",
				Action: func(cCtx *cli.Context) error {
					cmdComposerBased("create-project ", cCtx)
					return nil
				},
			},
			{
				Name:    "install",
				Aliases: []string{"i"},
				Usage:   "runs as ddev/lando composer install ...",
				Action: func(cCtx *cli.Context) error {
					cmdComposerBased("install ", cCtx)
					return nil
				},
			},
			{
				Name:    "require",
				Aliases: []string{"r"},
				Usage:   "runs as ddev/lando composer require ...",
				Action: func(cCtx *cli.Context) error {
					cmdComposerBased("require ", cCtx)
					return nil
				},
			},
			{
				Name:    "update",
				Aliases: []string{"u"},
				Usage:   "runs as ddev/lando composer update ...",
				Action: func(cCtx *cli.Context) error {
					cmdComposerBased("update ", cCtx)
					return nil
				},
			},
			{
				Name:    "site-install",
				Aliases: []string{"si"},
				Usage:   "runs as ddev/lando drush site-install profile_name",
				Action: func(cCtx *cli.Context) error {
					cmdDrushBased("site-install ", cCtx)
					return nil
				},
			},
			{
				Name:    "pm:enable",
				Aliases: []string{"en"},
				Usage:   "runs as ddev/lando drush pm-enable module_name",
				Action: func(cCtx *cli.Context) error {
					cmdDrushBased("pm:enable ", cCtx)
					return nil
				},
			},
			{
				Name:    "pm:uninstall",
				Aliases: []string{"pmu"},
				Usage:   "runs as ddev/lando drush pm-uninstall module_name",
				Action: func(cCtx *cli.Context) error {
					cmdDrushBased("pm:uninstall ", cCtx)
					return nil
				},
			},
			{
				Name:    "cache:rebuild",
				Aliases: []string{"cr"},
				Usage:   "runs as ddev/lando drush cache:rebuild",
				Action: func(cCtx *cli.Context) error {
					cmdDrushBased("cache:rebuild", cCtx)
					return nil
				},
			},
		},
		Authors: []*cli.Author{
			{Name: "Abhisek Mazumdar (abhisekmazumdar)"},
		},
		Action: mainAction,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

// Handle when user is not very lazy.
func mainAction(cCtx *cli.Context) error {

	devTool, strArgs := helper(cCtx)
	runner.Run(devTool, strArgs)

	return nil

}

// Handle any command which has composer as first argument.
func cmdComposer(cCtx *cli.Context) error {

	devTool, strArgs := helper(cCtx)
	runner.Run(devTool, "composer "+strArgs)

	return nil

}

// Handle any command which has drush as first argument.
func cmdDrush(cCtx *cli.Context) error {

	devTool, strArgs := helper(cCtx)
	runner.Run(devTool, "drush "+strArgs)

	return nil

}

// Handle any command which has a composer's second argument and idk sub-command.
func cmdComposerBased(subCmd string, cCtx *cli.Context) error {

	devTool, strArgs := helper(cCtx)
	cmd := "composer " + subCmd + strArgs

	runner.Run(devTool, cmd)

	return nil
}

// Handle any command which has a drush's second argument and idk sub-command.
func cmdDrushBased(subCmd string, cCtx *cli.Context) error {

	devTool, strArgs := helper(cCtx)
	cmd := "drush " + subCmd + strArgs

	runner.Run(devTool, cmd)

	return nil
}

// Helper for all commands.
func helper(cCtx *cli.Context) (string, string) {

	strArgs := ""
	strArgs = argsToString(cCtx.Args())

	fmt.Println("Checking which dev tool configuration files are present...")

	devTool := ""
	devTool = checkForDevTool()

	return devTool, strArgs
}

// Check for dev tool's(ddev or lando) config files are present.
func checkForDevTool() string {

	reader := bufio.NewReader(os.Stdin)

	if _, err := os.Stat("./.ddev"); !os.IsNotExist(err) {

		if _, err := os.Stat("./.lando.yml"); !os.IsNotExist(err) {

			fmt.Print("Both .lando.yml and .ddev directories exist.\nWhich one would you like to use? Type 'lando' or 'ddev': ")

			option, _ := reader.ReadString('\n')
			option = strings.ToLower(strings.TrimSpace(option))

			if option != "ddev" && option != "lando" {
				fmt.Println("You 🫵  spoiled all the fun.\nInvalid input, Try Again!!")
				os.Exit(1)
			}
			return option
		} else {
			return "ddev"
		}
	} else if _, err := os.Stat("./.lando.yml"); !os.IsNotExist(err) {
		return "lando"
	} else {
		fmt.Println("No ddev or lando configuration files found 🙁")
		os.Exit(1)
	}
	return ""
}

// Convert the cli.Args into string arguments.
func argsToString(args cli.Args) string {
	var str strings.Builder
	for i := 0; i < args.Len(); i++ {
		str.WriteString(args.Get(i))
		str.WriteString(" ")
	}
	return str.String()
}
