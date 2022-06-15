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

func Update() {
	iFiles, err := ioutil.ReadDir(path.Join(settings.Folders["config"], "installed"))
	if err != nil {
		arrowprint.Err0("cannot get list of installed packages")
		return
	}
	repository.Reload()
	for _, pkg := range iFiles {
		ivb, err := ioutil.ReadFile(path.Join(settings.Folders["config"], "installed", pkg.Name()))
		if err != nil {
			arrowprint.Err0("cannot get installed version of %s", pkg.Name())
			return
		}
		iv := strings.TrimSpace(string(ivb))
		var av string
	outer:
		for _, index := range repository.Indexes {
			for _, p := range index.PackageList {
				if p.Name == pkg.Name() {
					av = p.Version
					break outer
				}
			}
		}
		if utils.SemVersionGreater(av, iv) {
			arrowprint.Suc0("found newer version for %s, updating", pkg.Name())
			repository.InstallPackage(pkg.Name())
		}
	}
}
