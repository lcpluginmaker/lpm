package repository

import (
	"path"

	"github.com/alexcoder04/lpm/settings"
	"github.com/alexcoder04/lpm/utils"
)

func EnableDefault() error {
	return utils.WriteLinesList(
		path.Join(settings.Folders["config"], "repos.txt"),
		[]string{settings.DEFAULT_REPO}[:])
}

func GetList() []string {
	lines, err := utils.ReadLinesList(path.Join(settings.Folders["config"], "repos.txt"))
	if err != nil {
		return []string{}
	}
	return lines
}
