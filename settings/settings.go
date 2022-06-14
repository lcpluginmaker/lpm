package settings

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
)

type Config struct {
	FirstRun bool `json:"firstRun"`
	DevMode  bool `json:"devMode"`
}

const DEFAULT_REPO = "https://raw.githubusercontent.com/alexcoder04/LeoConsole-repo-main/main/index.json"

var Folders map[string]string = make(map[string]string)
var CurrentConfig Config = Config{}

func UpdateConfig() error {
	configFile := path.Join(Folders["config"], "config.json")
	cont, err := ioutil.ReadFile(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			f, err := os.Create(configFile)
			if err != nil {
				return err
			}
			defer f.Close()
		}
	}
	err = json.Unmarshal(cont, &CurrentConfig)
	if err != nil {
		return err
	}
	return nil
}
