package commands

import (
	"fmt"

	"github.com/alexcoder04/arrowprint"
	"github.com/alexcoder04/lpm/settings"
)

func Config(args []string) {
	if len(args) < 2 {
		arrowprint.Err0("you need to pass a subcommand, try 'lpm config help'")
		return
	}

	switch args[1] {
	case "get":
		if len(args) < 3 {
			arrowprint.Err0("you need to pass a key")
			return
		}
		switch args[2] {
		case "FirstRun":
			fmt.Println(settings.CurrentConfig.FirstRun)
			return
		case "DebugMode":
			fmt.Println(settings.CurrentConfig.DebugMode)
			return
		case "Repositories":
			for _, r := range settings.CurrentConfig.Repositories {
				fmt.Printf("%s: %s\n", r.Name, r.Url)
			}
			return
		default:
			arrowprint.Err0("unknown config key")
			return
		}
	case "help":
		fmt.Println(`
Available subcommands for 'lpm config'

"get" <key>:
	show config value
"help":
	show this help`)
	default:
		arrowprint.Err0("unknown subcommand: '%s', try 'lpm config help'", args[1])
		return
	}
}
