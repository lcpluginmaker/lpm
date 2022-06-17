package commands

import (
	"io/ioutil"
	"path"
	"strings"

	"github.com/alexcoder04/arrowprint"
	"github.com/alexcoder04/lpm/repository"
	"github.com/alexcoder04/lpm/settings"
	"github.com/alexcoder04/lpm/utils"
)

func ListAvailable() {
	arrowprint.Suc0("available packages:")
	for _, i := range repository.Indexes {
		arrowprint.Info0("repo %s:", i.Name)
		for _, p := range i.PackageList {
			if p.OS == utils.GetOS() || p.OS == "any" {
				arrowprint.Suc1("%s v%s: %s", p.Name, p.Version, p.Description)
			}
		}
	}
}

func ListInstalled() {
	arrowprint.Suc0("installed packages:")
	f, err := ioutil.ReadDir(path.Join(settings.Folders["config"], "installed"))
	if err != nil {
		arrowprint.Err1("cannot open package database")
		return
	}
	for _, f := range f {
		v, err := ioutil.ReadFile(path.Join(settings.Folders["config"], "installed", f.Name(), "version"))
		if err != nil {
			arrowprint.Err1("%s: cannot check version", f.Name())
			continue
		}
		arrowprint.Suc1("%s v%s", f.Name(), strings.TrimSpace(string(v)))
	}
}
