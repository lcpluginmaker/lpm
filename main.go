package main

import (
	"path"

	"github.com/alexcoder04/arrowprint"
	"github.com/alexcoder04/gilc"
	"github.com/alexcoder04/lpm/commands"
	"github.com/alexcoder04/lpm/initial"
	"github.com/alexcoder04/lpm/repository"
	"github.com/alexcoder04/lpm/settings"
	"github.com/alexcoder04/lpm/utils"
)

func _init(data gilc.IData, stage int) error {
	switch stage {
	case 0:
		settings.Folders["config"] = path.Join(data.SavePath, "var", "lpm")
		settings.Folders["temp"] = path.Join(data.DownloadPath, "apkg")
		settings.Folders["root"] = data.SavePath
		break
	case 1:
		repository.Reload()
		break
	case 2:
		err := settings.UpdateConfig()
		if err != nil {
			return err
		}
		break
	}
	return nil
}

func pmain(data gilc.IData) {
	_init(data, 0)

	// create folders
	err := initial.CreateVarFolders()
	if err != nil {
		arrowprint.Err0("cannot create var folders: %s", err.Error())
		return
	}

	// check repositories
	err = initial.CheckRepositories()
	if err != nil {
		arrowprint.Err0("error checking repositories: %s", err.Error())
		return
	}

	_init(data, 1)

	if !utils.StringArrayContains(repository.GetListInstalled(), "lpm") {
		arrowprint.Warn0("lpm is not installed properly, installing...")
		err := repository.InstallPackage("lpm")
		if err != nil {
			arrowprint.Err0("cannot install lpm: %s", err.Error())
			return
		}
	}

	err = _init(data, 2)
	if err != nil {
		arrowprint.Err0("cannot load config: %s", err.Error())
		return
	}
}

func pcommand(data gilc.IData, args []string) {
	if len(args) < 1 {
		arrowprint.Err0("you need pass a subcommand")
		return
	}
	_init(data, 0)
	_init(data, 1)
	err := _init(data, 2)
	if err != nil {
		arrowprint.Err0("cannot load config: %s", err.Error())
		return
	}
	// TODO: update
	// TODO debug: build, get-local
	switch args[0] {
	case "reload", "sync", "s":
		repository.Reload()
		return
	case "search", "f":
		commands.Search(args)
		return
	case "get", "install", "i":
		if len(args) < 2 {
			arrowprint.Err0("you need to pass the package name")
			return
		}
		repository.InstallPackage(args[1])
		return
	case "remove", "uninstall", "r":
		if len(args) < 2 {
			arrowprint.Err0("you need to pass the package name")
			return
		}
		repository.RemovePackage(args[1])
		return
	case "list-installed", "li":
		commands.ListInstalled()
		return
	case "list-available", "la":
		commands.ListAvailable()
		return
	case "help", "h":
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
