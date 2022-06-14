package initial

import (
	"errors"

	"github.com/alexcoder04/lpm/repository"
)

func CheckRepositories() ([]string, error) {
	repos := repository.GetList()
	if len(repos) >= 1 {
		return repos, nil
	}
	err := repository.EnableDefault()
	if err != nil {
		return []string{}, err
	}
	repos = repository.GetList()
	if len(repos) < 1 {
		return []string{}, errors.New("default repository cannot be enabled")
	}
	return repos, nil
}
