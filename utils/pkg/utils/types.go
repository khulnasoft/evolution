package utils

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Maintainer struct {
	Name     string   `yaml:"name"`
	Github   string   `yaml:"github"`
	Company  string   `yaml:"company"`
	Projects []string `yaml:"projects"`
}

type Maintainers []Maintainer

type RepositoryScope string

type RepositoryStatus string

type Repository struct {
	Name        string           `yaml:"name"`
	Description string           `yaml:"description,omitempty"`
	Scope       RepositoryScope  `yaml:"scope"`
	Status      RepositoryStatus `yaml:"status,omitempty"`
}

type Repositories []Repository

const (
	RepositoryStatusStable     RepositoryStatus = "Stable"
	RepositoryStatusIncubating RepositoryStatus = "Incubating"
	RepositoryStatusSandbox    RepositoryStatus = "Sandbox"
	RepositoryStatusDeprecated RepositoryStatus = "Deprecated"
)

const (
	RepositoryScopeCore      RepositoryScope = "Core"
	RepositoryScopeEcosystem RepositoryScope = "Ecosystem"
	RepositoryScopeInfra     RepositoryScope = "Infra"
	RepositoryScopeSpecial   RepositoryScope = "Special"
)

func (r RepositoryStatus) String() string {
	return string(r)
}

func (r RepositoryScope) String() string {
	return string(r)
}

func (r *Repository) URL() string {
	return fmt.Sprintf("https://github.com/khulnasoft/%s", r.Name)
}

func readFromFile(path string, out interface{}) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, out)
}

func ReadRepositoriesFromFile(path string) (Repositories, error) {
	var res Repositories
	if err := readFromFile(path, &res); err != nil {
		return nil, err
	}
	return res, nil
}

func ReadMaintainersFromFile(path string) (Maintainers, error) {
	var res Maintainers
	if err := readFromFile(path, &res); err != nil {
		return nil, err
	}
	return res, nil
}
