package commands

import (
	"strings"

	"github.com/alexcoder04/arrowprint"
	"github.com/alexcoder04/lpm/repository"
	"github.com/alexcoder04/lpm/utils"
)

func GetLocal(arg string) error {
	if strings.HasSuffix(arg, ".lcp") {
		arrowprint.Suc0("installing local .lcp file")
		return repository.InstallArchive(arg)
	}

	if utils.IsDir(arg) {
		outFile, err := repository.BuildFolder(arg)
		if err != nil {
			arrowprint.Err0("compiling plugin failed: %s", err.Error())
			return err
		}
		return repository.InstallArchive(outFile)
	}

	arrowprint.Err0("the argument is not valid")
	return nil
}
