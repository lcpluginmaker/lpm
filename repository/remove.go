package repository

import (
	"os"
	"path"
	"strings"

	"github.com/alexcoder04/arrowprint"
	"github.com/alexcoder04/lpm/settings"
	"github.com/alexcoder04/lpm/utils"
)

func RemovePackage(name string) error {
	arrowprint.Suc0("removing %s", name)
	if !utils.StringArrayContains(GetListInstalled(), name) {
		arrowprint.Info0("%s is not installed", name)
		return nil
	}
	files, err := utils.ReadLinesList(path.Join(settings.Folders["config"], "installed", name, "files"))
	if err != nil {
		return err
	}
	for i, f := range files {
		if strings.TrimSpace(f) == "" {
			continue
		}
		arrowprint.Info1("removing file %d of %d", i+1, len(files))
		err := os.Remove(path.Join(settings.Folders["root"], strings.TrimSpace(f)))
		if err != nil {
			return err
		}
	}
	arrowprint.Info1("unregistering %s from database", name)
	err = os.RemoveAll(path.Join(settings.Folders["config"], "installed", name))
	if err != nil {
		return err
	}
	arrowprint.Suc0("package %s removed", name)
	return nil
}
