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
	arrowprint.Suc0("available packages")
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
	arrowprint.Suc0("installed packages")
	for _, i := range repository.Indexes {
		for _, p := range i.PackageList {
			if p.OS != utils.GetOS() && p.OS != "any" {
				continue
			}
			if utils.IsDir(path.Join(settings.Folders["config"], "installed", p.Name)) {
				v, err := ioutil.ReadFile(path.Join(settings.Folders["config"], "installed", p.Name, "version"))
				if err != nil {
					arrowprint.Err1("%s: cannot check version", p.Name)
					continue
				}
				arrowprint.Suc1("%s v%s: %s", p.Name, strings.TrimSpace(string(v)), p.Description)
			}
		}
	}
}
