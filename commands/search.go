package commands

import (
	"strings"

	"github.com/alexcoder04/arrowprint"
	"github.com/alexcoder04/lpm/repository"
)

func Search(args []string) {
	if len(args) < 2 {
		arrowprint.Err0("you need to pass the package name")
		return
	}
	for _, r := range repository.Indexes {
		for _, p := range r.PackageList {
			if strings.Contains(p.Name, args[1]) {
				arrowprint.Suc1(p.Name)
			}
		}
	}
}
