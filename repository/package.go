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
	"github.com/alexcoder04/gilc"
	"github.com/alexcoder04/lpm/settings"
	"github.com/alexcoder04/lpm/utils"
)

func GetUrlFor(pkgname string) string {
	for _, r := range Indexes {
		for _, p := range r.PackageList {
			if p.Name == pkgname {
				return p.Url
			}
		}
	}
	return ""
}

func RemovePackage(name string) error {
	if !utils.StringArrayContains(GetListInstalled(), name) {
		return nil
	}
	files, err := utils.ReadLinesList(path.Join(settings.Folders["config"], "installed", name, "files"))
	if err != nil {
		return err
	}
	for _, f := range files {
		err := os.Remove(path.Join(settings.Folders["root"], f))
		if err != nil {
			return err
		}
	}
	return os.RemoveAll(path.Join(settings.Folders["config"], "installed", name))
}

func InstallPackage(name string) error {
	dlUrl := GetUrlFor(name)
	if dlUrl == "" {
		return errors.New("package " + name + " was not found")
	}
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
	return InstallArchive(path.Join(settings.Folders["temp"], name+"lcp"))
}

func InstallArchive(file string) error {
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
	cont, err := ioutil.ReadFile(path.Join(unzipDir, "PKGINFO.json"))
	if err != nil {
		return err
	}
	var pkginfo builder.PKGINFO
	err = json.Unmarshal(cont, &pkginfo)
	if err != nil {
		return err
	}
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
	for _, f := range pkginfo.Files {
		fullPath := path.Join(settings.Folders["root"], f)
		err := os.MkdirAll(path.Dir(fullPath), 0700)
		if err != nil {
			return err
		}
		err = utils.CopyFile(path.Join(unzipDir, f), path.Join(settings.Folders["root"], f))
		if err != nil {
			return err
		}
		if strings.HasPrefix(f, "share/scripts") || strings.HasPrefix(f, "share/go-plugin") {
			err := os.Chmod(path.Join(settings.Folders["root"], f), 0700)
			if err != nil {
				return err
			}
		}
	}
	err = os.MkdirAll(path.Join(settings.Folders["config"], "installed", pkginfo.PackageName), 0700)
	if err != nil {
		return err
	}
	err = utils.WriteLinesList(path.Join(settings.Folders["config"], "installed", pkginfo.PackageName, "files"), pkginfo.Files)
	if err != nil {
		return err
	}
	return utils.WriteLinesList(path.Join(settings.Folders["config"], "installed", pkginfo.PackageName, "version"), []string{pkginfo.PackageVersion})
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
