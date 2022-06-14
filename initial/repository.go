package initial

import (
	"errors"

	"github.com/alexcoder04/lpm/repository"
)

func CheckRepositories() error {
	repos := repository.GetListRepositories()
	if len(repos) >= 1 {
		return nil
	}
	err := repository.EnableDefault()
	if err != nil {
		return err
	}
	repos = repository.GetListRepositories()
	if len(repos) < 1 {
		return errors.New("default repository cannot be enabled")
	}
	return nil
}
