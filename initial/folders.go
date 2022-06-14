package initial

import (
	"os"
	"path"

	"github.com/alexcoder04/lpm/settings"
)

func CreateVarFolders() error {
	var foldersRequired []string = []string{
		settings.Folders["config"],
		path.Join(settings.Folders["config"], "installed"),
		path.Join(settings.Folders["config"], "repos-index"),
	}
	for _, f := range foldersRequired {
		err := os.MkdirAll(f, 0700)
		if err != nil {
			return err
		}
	}
	return nil
}
