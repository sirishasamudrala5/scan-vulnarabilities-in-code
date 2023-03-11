package models

import (
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type Repository struct {
	gorm.Model
	Name     string `json:"name" bson:"name"`
	RepoLink string `json:"repo_link" bson:"repo_link"`
}

func (i *Repository) Bind(r *http.Request) error {
	if i.Name == "" {
		return fmt.Errorf("name is a required field")
	}
	return nil
}

func (*Repository) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type Getter interface {
	GetAll() []Repository
}

type Adder interface {
	Add(repositoryData Repository)
}

type RepositoryList struct {
	Repositories []Repository
}

func NewRepositoryList() *RepositoryList {
	return &RepositoryList{
		Repositories: []Repository{},
	}
}

func (r *RepositoryList) Add(item Repository) {
	r.Repositories = append(r.Repositories, item)
}

func (r *RepositoryList) GetAll() []Repository {
	return r.Repositories
}
