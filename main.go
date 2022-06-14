package main

import (
	"path"

	"github.com/alexcoder04/arrowprint"
	"github.com/alexcoder04/gilc"
	"github.com/alexcoder04/lpm/commands"
	"github.com/alexcoder04/lpm/initial"
	"github.com/alexcoder04/lpm/settings"
)

func pmain(data gilc.IData) {
	// create folders
	settings.Folders["config"] = path.Join(data.SavePath, "var", "lpm")
	err := initial.CreateVarFolders()
	if err != nil {
		arrowprint.Err0("cannot create var folders: %s", err.Error())
		return
	}

	// check repositories
	_, err = initial.CheckRepositories()
	if err != nil {
		arrowprint.Err0("error checking repositories: %s", err.Error())
		return
	}

	// TODO: reload
	// TODO: install itself if not installed
	// TODO: read config file
}

func pcommand(data gilc.IData, args []string) {
	if len(args) < 1 {
		arrowprint.Err0("you need pass a subcommand")
		return
	}
	// TODO: get, reload, list(available/installed), search, info, remove, update
	// TODO debug: build, get-local
	switch args[0] {
	case "help":
		commands.Help()
		return
	default:
		arrowprint.Err0("unknown subcomand: '%s'", args[0])
		return
	}
}

func pshutdown(data gilc.IData) {
}

func main() {
	gilc.Setup("leoconsole package manager", pmain, pcommand, pshutdown)
	gilc.Run()
}
