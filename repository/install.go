package repository

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"

	builder "github.com/alexcoder04/LeoConsole-apkg-builder/pkg"
	"github.com/alexcoder04/arrowprint"
	"github.com/alexcoder04/gilc"
	"github.com/alexcoder04/lpm/settings"
	"github.com/alexcoder04/lpm/utils"
)

// TODO refactor this file

func InstallPackage(name string) error {
	dlUrl := GetUrlFor(name)
	if dlUrl == "" {
		return errors.New("package " + name + " was not found")
	}
	arrowprint.Suc1("downloading from %s", dlUrl)
	res, err := http.Get(dlUrl)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	out, err := os.Create(path.Join(settings.Folders["temp"], name+"lcp"))
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, res.Body)
	if err != nil {
		return err
	}
	err = InstallArchive(path.Join(settings.Folders["temp"], name+"lcp"))
	if err != nil {
		return err
	}
	arrowprint.Suc0("package %s removed", name)
	return nil
}

func InstallArchive(file string) error {
	arrowprint.Suc0("installing %s", file)
	arrowprint.Suc1("extracting package")
	unzipDir := path.Join(settings.Folders["temp"], strings.Split(file, ".")[0])
	if utils.IsDir(unzipDir) {
		err := os.RemoveAll(unzipDir)
		if err != nil {
			return err
		}
	}
	_, err := utils.Unzip(file, unzipDir)
	if err != nil {
		return err
	}
	arrowprint.Suc1("reading PKGINFO")
	cont, err := ioutil.ReadFile(path.Join(unzipDir, "PKGINFO.json"))
	if err != nil {
		return err
	}
	var pkginfo builder.PKGINFO
	err = json.Unmarshal(cont, &pkginfo)
	if err != nil {
		return err
	}
	if pkginfo.PackageOS != utils.GetOS() && pkginfo.PackageOS != "any" {
		return errors.New("package OS incompatible")
	}
	arrowprint.Suc1("checking integrity")
	conflict, err := AnyFilesAlreadyInstalled(pkginfo.Files)
	if err != nil {
		return err
	}
	if conflict {
		return errors.New("package conflicts with some installed package")
	}
	if utils.StringArrayContains(GetListInstalled(), pkginfo.PackageName) {
		installedVersion, err := ioutil.ReadFile(path.Join(settings.Folders["config"], pkginfo.PackageName, "version"))
		if err != nil {
			return err
		}
		if strings.TrimSpace(string(installedVersion)) == strings.TrimSpace(pkginfo.PackageVersion) {
			ans, err := gilc.YesNoDialog("reinstall same package version", true)
			if err != nil || !ans {
				return err
			}
		} else if utils.SemVersionGreater(string(installedVersion), pkginfo.PackageVersion) {
			ans, err := gilc.YesNoDialog("downgrade package", false)
			if err != nil || !ans {
				return err
			}
		}
		err = RemovePackage(pkginfo.PackageName)
		if err != nil {
			return err
		}
	}
	for i, f := range pkginfo.Files {
		fullPath := path.Join(settings.Folders["root"], f)
		err := os.MkdirAll(path.Dir(fullPath), 0700)
		if err != nil {
			return err
		}
		arrowprint.Info1("installing file %d of %d", i+1, len(pkginfo.Files))
		err = utils.CopyFile(path.Join(unzipDir, f), path.Join(settings.Folders["root"], f))
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
	arrowprint.Suc0("registering %s in database", pkginfo.PackageName)
	err = os.MkdirAll(path.Join(settings.Folders["config"], "installed", pkginfo.PackageName), 0700)
	if err != nil {
		return err
	}
	err = utils.WriteLinesList(path.Join(settings.Folders["config"], "installed", pkginfo.PackageName, "files"), pkginfo.Files)
	if err != nil {
		return err
	}
	err = utils.WriteLinesList(path.Join(settings.Folders["config"], "installed", pkginfo.PackageName, "version"), []string{pkginfo.PackageVersion})
	if err != nil {
		return err
	}
	arrowprint.Suc1("package %s installed", pkginfo.PackageName)
	return nil
}
