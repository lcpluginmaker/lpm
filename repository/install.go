package repository

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/alexcoder04/arrowprint"
	"github.com/alexcoder04/lpm/settings"
	"github.com/alexcoder04/lpm/utils"
)

func InstallPackage(name string) error {
	dlUrl := GetUrlFor(name)
	if dlUrl == "" {
		return errors.New("package " + name + " was not found")
	}
	arrowprint.Suc1("downloading from %s", dlUrl)
	tempFile := path.Join(settings.Folders["temp"], name+"lcp")
	err := utils.DownloadFile(dlUrl, tempFile)
	if err != nil {
		return err
	}
	err = InstallArchive(tempFile)
	if err != nil {
		return err
	}
	arrowprint.Suc0("package %s installed", name)
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

	pkginfo, err := ReadPkginfo(unzipDir)
	if err != nil {
		return err
	}

	// TODO throws an error even if same package
	err = CheckCompatibility(pkginfo)
	if err != nil {
		return err
	}

	if utils.StringArrayContains(GetListInstalled(), pkginfo.PackageName) {
		installed, err := ioutil.ReadFile(path.Join(settings.Folders["config"], pkginfo.PackageName, "version"))
		if err != nil {
			return err
		}
		err = HandleVersioning(string(installed), pkginfo.PackageVersion, pkginfo.PackageName)
		if err != nil {
			return err
		}
	}

	err = InstallFiles(pkginfo.Files, unzipDir)
	if err != nil {
		return err
	}

	err = RegisterPackage(pkginfo)
	if err != nil {
		return err
	}

	arrowprint.Suc1("package %s installed", pkginfo.PackageName)
	return nil
}
