package repository

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path"

	"github.com/alexcoder04/arrowprint"
	"github.com/alexcoder04/lpm/settings"
	"github.com/alexcoder04/lpm/utils"
)

var Indexes []RepositoryIndex = getIndexes()

// TODO save to file and restore from file
func Reload() {
	arrowprint.Suc0("reloading repositories")
	Indexes = getIndexes()
}

func getIndexes() []RepositoryIndex {
	var indexList []RepositoryIndex = make([]RepositoryIndex, 0)
	for _, r := range GetListRepositories() {
		arrowprint.Suc1("processing %s", r)
		res, err := http.Get(r)
		if err != nil {
			continue
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			continue
		}
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			continue
		}
		repo := RepositoryIndex{}
		err = json.Unmarshal(body, &repo)
		if err != nil {
			continue
		}
		indexList = append(indexList, repo)
	}
	return indexList
}

func EnableDefault() error {
	arrowprint.Suc1("enabling default repository (%s)", settings.DEFAULT_REPO)
	return utils.WriteLinesList(
		path.Join(settings.Folders["config"], "repos.txt"),
		[]string{settings.DEFAULT_REPO}[:])
}

func GetListRepositories() []string {
	lines, err := utils.ReadLinesList(path.Join(settings.Folders["config"], "repos.txt"))
	if err != nil {
		return []string{}
	}
	return lines
}
