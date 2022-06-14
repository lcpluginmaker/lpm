package repository

type RepositoryIndex struct {
	IndexVersion float64          `json:"indexVersion"`
	Name         string           `json:"name"`
	Url          string           `json:"url"`
	Project      IndexProjectData `json:"project"`
	PackageList  []RepoPackage    `json:"packageList"`
}

type IndexProjectData struct {
	Description string `json:"description"`
	Maintainer  string `json:"maintainer"`
	Email       string `json:"email"`
	Homepage    string `json:"homepage"`
	BugTracker  string `json:"bugTracker"`
}

type RepoPackage struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     string `json:"version"`
	OS          string `json:"os"`
	Url         string `json:"url"`
}
