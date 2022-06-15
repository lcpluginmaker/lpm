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

var Indexes []RepositoryIndex = getIndexes(false)

func Reload() {
	arrowprint.Suc0("reloading repositories")
	Indexes = getIndexes(true)
}

func loadIndexFromFile(r settings.Repository) (RepositoryIndex, error) {
	cont, err := ioutil.ReadFile(path.Join(settings.Folders["config"], "repos-index", r.Name+".json"))
	if err != nil {
		return RepositoryIndex{}, err
	}
	var repo RepositoryIndex
	err = json.Unmarshal(cont, &repo)
	return repo, err
}

func getIndexes(hard bool) []RepositoryIndex {
	var indexList []RepositoryIndex = make([]RepositoryIndex, 0)
	for _, r := range settings.CurrentConfig.Repositories {
		arrowprint.Suc1("processing %s", r.Name)
		if !hard && utils.IsFile(path.Join(settings.Folders["config"], "repos-index", r.Name+".json")) {
			repo, err := loadIndexFromFile(r)
			if err == nil {
				arrowprint.Info1("using cached index for %s", r.Name)
				indexList = append(indexList, repo)
				arrowprint.Warn1("error processing %s, skipping (%s)", r.Name, err.Error())
				continue
			}
		}
		res, err := http.Get(r.Url)
		if err != nil {
			arrowprint.Warn1("error processing %s, skipping (%s)", r.Name, err.Error())
			continue
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			arrowprint.Warn1("error processing %s, skipping (%s)", r.Name, err.Error())
			continue
		}
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			arrowprint.Warn1("error processing %s, skipping (%s)", r.Name, err.Error())
			continue
		}
		repo := RepositoryIndex{}
		err = json.Unmarshal(body, &repo)
		if err != nil {
			arrowprint.Warn1("error processing %s, skipping (%s)", r.Name, err.Error())
			continue
		}
		indexList = append(indexList, repo)
		err = ioutil.WriteFile(path.Join(settings.Folders["config"], "repo-index", r.Name+".json"), body, 0600)
		if err != nil {
			arrowprint.Warn1("cannot cache repository index: %s", err.Error())
		}
	}
	return indexList
}
