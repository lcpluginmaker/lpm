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
		err := settings.UpdateConfig()
		if err != nil {
			return err
		}
		break
	case 1:
		repository.Reload()
		break
	}
	return nil
}

func pmain(data gilc.IData) {
	arrowprint.Suc0("running PluginMain for lpm")

	err := _init(data, 0)
	if err != nil {
		arrowprint.Err0("cannot load config: %s", err.Error())
		return
	}

	// create folders
	err = initial.CreateVarFolders()
	if err != nil {
		arrowprint.Err0("cannot create var folders: %s", err.Error())
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
}

func pcommand(data gilc.IData, args []string) {
	if len(args) < 1 {
		arrowprint.Err0("you need pass a subcommand")
		return
	}

	err := _init(data, 0)
	if err != nil {
		arrowprint.Err0("cannot load config: %s", err.Error())
		return
	}
	_init(data, 1)

	// TODO debug: build
	switch args[0] {
	case "search", "f":
		commands.Search(args)
		return
	case "help", "h":
		commands.Help()
		return
	case "get-local", "gl":
		if !settings.CurrentConfig.DebugMode {
			arrowprint.Err0("this command is only available in debug mode")
			return
		}
		if len(args) < 2 {
			arrowprint.Err0("you need to pass an argument")
			return
		}
		err := commands.GetLocal(args[1])
		if err != nil {
			arrowprint.Err0("error installing %s: %s", args[1], err.Error())
		}
		return
	case "get", "install", "i":
		if len(args) < 2 {
			arrowprint.Err0("you need to pass the package name")
			return
		}
		err := repository.InstallPackage(args[1])
		if err != nil {
			arrowprint.Err0("error installing %s: %s", args[1], err.Error())
		}
		return
	case "list-available", "la":
		commands.ListAvailable()
		return
	case "list-installed", "li":
		commands.ListInstalled()
		return
	case "remove", "uninstall", "r":
		if len(args) < 2 {
			arrowprint.Err0("you need to pass the package name")
			return
		}
		repository.RemovePackage(args[1])
		return
	case "reload", "sync", "s":
		repository.Reload()
		return
	case "update", "upgrade", "u":
		commands.Update()
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
