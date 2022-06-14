package repository

import (
	"errors"
	"io/ioutil"
	"path"
	"strings"

	builder "github.com/alexcoder04/LeoConsole-apkg-builder/pkg"
	"github.com/alexcoder04/arrowprint"
	"github.com/alexcoder04/lpm/settings"
	"github.com/alexcoder04/lpm/utils"
)

func GetUrlFor(pkgname string) string {
	arrowprint.Info1("resolving %s", pkgname)
	for _, r := range Indexes {
		for _, p := range r.PackageList {
			if p.Name == pkgname {
				return p.Url
			}
		}
	}
	return ""
}

func CheckCompatibility(pkginfo builder.PKGINFO) error {
	arrowprint.Suc1("checking compatibility")
	if pkginfo.PackageOS != utils.GetOS() && pkginfo.PackageOS != "any" {
		return errors.New("package OS incompatible")
	}
	conflict, err := AnyFilesAlreadyInstalled(pkginfo.Files)
	if conflict {
		return errors.New("package conflicts with some installed package")
	}
	return err
}

func AnyFilesAlreadyInstalled(files []string) (bool, error) {
	for _, f := range files {
		packageDirs, err := ioutil.ReadDir(path.Join(settings.Folders["config"], "installed"))
		if err != nil {
			return true, err
		}
		for _, pDir := range packageDirs {
			lines, err := utils.ReadLinesList(path.Join(pDir.Name(), "files"))
			if err != nil {
				return true, err
			}
			for _, l := range lines {
				if strings.TrimSpace(l) == strings.TrimSpace(f) {
					return true, nil
				}
			}
		}
	}
	return false, nil
}

func GetListInstalled() []string {
	files, err := ioutil.ReadDir(path.Join(settings.Folders["config"], "installed"))
	if err != nil {
		return []string{}
	}
	var pkgs []string = []string{}
	for _, f := range files {
		pkgs = append(pkgs, f.Name())
	}
	return pkgs
}
