package initial

import (
	"os"
	"path"

	"github.com/alexcoder04/arrowprint"
	"github.com/alexcoder04/lpm/settings"
	"github.com/alexcoder04/lpm/utils"
)

func CreateVarFolders() error {
	// link to apkg folder if it exists
	apkgDir := path.Join(settings.Folders["root"], "var", "apkg")
	if utils.IsDir(apkgDir) && !utils.IsDir(settings.Folders["config"]) {
		arrowprint.Suc1("found apkg folder, creating symlink")
		err := os.Symlink(apkgDir, settings.Folders["config"])
		if err != nil {
			return err
		}
	}
	// create required folders
	var foldersRequired []string = []string{
		path.Join(settings.Folders["config"], "installed"),
		path.Join(settings.Folders["config"], "repos-index"),
		settings.Folders["temp"],
	}
	arrowprint.Suc1("checking required folders")
	for _, f := range foldersRequired {
		err := os.MkdirAll(f, 0700)
		if err != nil && !os.IsExist(err) {
			return err
		}
	}
	return nil
}
