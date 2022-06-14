package repository

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"strings"

	builder "github.com/alexcoder04/LeoConsole-apkg-builder/pkg"
	"github.com/alexcoder04/arrowprint"
	"github.com/alexcoder04/gilc"
	"github.com/alexcoder04/lpm/settings"
	"github.com/alexcoder04/lpm/utils"
)

func RegisterPackage(pkginfo builder.PKGINFO) error {
	arrowprint.Suc0("registering %s in database", pkginfo.PackageName)
	err := os.MkdirAll(path.Join(settings.Folders["config"], "installed", pkginfo.PackageName), 0700)
	if err != nil {
		return err
	}
	err = utils.WriteLinesList(path.Join(settings.Folders["config"], "installed", pkginfo.PackageName, "files"), pkginfo.Files)
	if err != nil {
		return err
	}
	err = utils.WriteLinesList(path.Join(settings.Folders["config"], "installed", pkginfo.PackageName, "version"), []string{pkginfo.PackageVersion})
	return err
}

func ReadPkginfo(folder string) (builder.PKGINFO, error) {
	arrowprint.Suc1("reading PKGINFO")
	cont, err := ioutil.ReadFile(path.Join(folder, "PKGINFO.json"))
	if err != nil {
		return builder.PKGINFO{}, err
	}
	var pkginfo builder.PKGINFO
	err = json.Unmarshal(cont, &pkginfo)
	return pkginfo, err
}

func HandleVersioning(iv string, av string, name string) error {
	iv = strings.TrimSpace(iv)
	av = strings.TrimSpace(av)
	if iv == av {
		ans, err := gilc.YesNoDialog("reinstall same package version", true)
		if err != nil || !ans {
			return err
		}
	} else if utils.SemVersionGreater(iv, av) {
		ans, err := gilc.YesNoDialog("downgrade package", false)
		if err != nil || !ans {
			return err
		}
	}
	err := RemovePackage(name)
	if err != nil {
		return err
	}
	return nil
}

func InstallFiles(files []string, tempDir string) error {
	for i, f := range files {
		fullPath := path.Join(settings.Folders["root"], f)
		err := os.MkdirAll(path.Dir(fullPath), 0700)
		if err != nil {
			return err
		}
		arrowprint.Info1("installing file %d of %d", i+1, len(files))
		err = utils.CopyFile(path.Join(tempDir, f), path.Join(settings.Folders["root"], f))
		if err != nil {
			return err
		}
		if strings.HasPrefix(f, "share/scripts") || strings.HasPrefix(f, "share/go-plugin") {
			arrowprint.Info1("marking %s as executable", f)
			err := os.Chmod(path.Join(settings.Folders["root"], f), 0700)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
