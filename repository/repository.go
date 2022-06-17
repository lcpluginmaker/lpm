package repository

import (
	"encoding/json"
	"errors"
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

func getCachedIndex(repo settings.Repository) (RepositoryIndex, error) {
	if !utils.IsFile(path.Join(settings.Folders["config"], "repos-index", repo.Name+".json")) {
		return RepositoryIndex{}, errors.New("cache not available")
	}
	data, err := loadIndexFromFile(repo)
	if err == nil {
		arrowprint.Info1("using cached index for %s", repo.Name)
	}
	return data, err
}

func getIndexes(hard bool) []RepositoryIndex {
	var indexList []RepositoryIndex = make([]RepositoryIndex, 0)
	for _, r := range settings.CurrentConfig.Repositories {
		arrowprint.Suc1("processing %s", r.Name)
		if !hard && utils.IsFile(path.Join(settings.Folders["config"], "repos-index", r.Name+".json")) {
			repo, err := getCachedIndex(r)
			if err == nil {
				indexList = append(indexList, repo)
				continue
			}
		}
		res, err := http.Get(r.Url)
		if err != nil {
			repo, err := getCachedIndex(r)
			if err == nil {
				arrowprint.Warn1("cannot update %s, using cached index", r.Name)
				indexList = append(indexList, repo)
				continue
			}
			arrowprint.Warn1("error processing %s, skipping (%s)", r.Name, err.Error())
			continue
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			repo, err := getCachedIndex(r)
			if err == nil {
				arrowprint.Warn1("cannot update %s, using cached index", r.Name)
				indexList = append(indexList, repo)
				continue
			}
			arrowprint.Warn1("error processing %s, skipping (%s)", r.Name, err.Error())
			continue
		}
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			repo, err := getCachedIndex(r)
			if err == nil {
				arrowprint.Warn1("cannot update %s, using cached index", r.Name)
				indexList = append(indexList, repo)
				continue
			}
			arrowprint.Warn1("error processing %s, skipping (%s)", r.Name, err.Error())
			continue
		}
		repo := RepositoryIndex{}
		err = json.Unmarshal(body, &repo)
		if err != nil {
			repo, err := getCachedIndex(r)
			if err == nil {
				arrowprint.Warn1("cannot update %s, using cached index", r.Name)
				indexList = append(indexList, repo)
				continue
			}
			arrowprint.Warn1("error processing %s, skipping (%s)", r.Name, err.Error())
			continue
		}
		indexList = append(indexList, repo)
		err = ioutil.WriteFile(path.Join(settings.Folders["config"], "repos-index", r.Name+".json"), body, 0600)
		if err != nil {
			arrowprint.Warn1("cannot cache repository index: %s", err.Error())
		}
	}
	return indexList
}
