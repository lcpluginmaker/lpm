package commands

import (
	"fmt"

	"github.com/alexcoder04/lpm/settings"
)

func Help() {
	fmt.Println(`
Available subcommands

"list-available", "a":
	list all packages from repositories
"config", "conf", "cfg", "c":
	viewing and manipulating the config
"search", "f":
	search for a package in repositories
"help", "h":
	print this help
"get", "install", "i":
	install a package from repositories
"list-installed", "l":
	list all installed packages
"remove", "uninstall", "r":
	remove an installed package
"reload", "sync", "s":
	reload the repository cache
"update", "upgrade", "u":
	update your installed packages`)
	if settings.CurrentConfig.DebugMode {
		fmt.Println(`

Available subcommands in debug mode

"build", "b":
	build a plugin folder
"get-local", "d":
	build a plugin folder and install it`)
	}

}
