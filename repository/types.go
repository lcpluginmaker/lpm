package repository

type RepositoryIndex struct {
	IndexVersion float64               `json:"indexVersion"`
	Name         string                `json:"name"`
	Url          string                `json:"url"`
	Project      RepositoryProjectData `json:"project"`
	PackageList  []RepositoryPackage   `json:"packageList"`
}

type RepositoryProjectData struct {
	Description string `json:"description"`
	Maintainer  string `json:"maintainer"`
	Email       string `json:"email"`
	Homepage    string `json:"homepage"`
	BugTracker  string `json:"bugTracker"`
}

type RepositoryPackage struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     string `json:"version"`
	OS          string `json:"os"`
	Url         string `json:"url"`
}
