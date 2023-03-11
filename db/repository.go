package db

import (
	"os"
	"path"
	"strconv"
	"time"

	"bitbucket.org/guardrails-go/models"
)

var conn, _ = connect()

func GetAllRepositories() ([]models.Repository, error) {
	result := []models.Repository{}
	conn.Find(&result)
	return result, nil
}

func GetRepositoryById(repoId string) (models.Repository, error) {
	intrepoId, _ := strconv.Atoi(repoId)
	result := models.Repository{}
	err := conn.Where("id = ?", intrepoId).Find(&result)
	return result, err.Error
}

func AddRepository(newRepository *models.Repository) error {
	err := conn.FirstOrCreate(newRepository)
	return err.Error
}

func UpdateRepository(repoId string, repoData models.Repository) (models.Repository, error) {
	intrepoId, _ := strconv.Atoi(repoId)
	//Get document's collection
	repo := &models.Repository{}
	conn.First(&repo, `id=?`, intrepoId)

	if repoData.Name != "" {
		repo.Name = repoData.Name
	}
	if repoData.RepoLink != "" {
		repo.RepoLink = repoData.RepoLink
	}

	err := conn.Model(&repo).Updates(repo)
	if err != nil {
		return *repo, err.Error
	}
	return *repo, nil
}

func DeleteRepository(repoId string) error {
	intrepoId, _ := strconv.Atoi(repoId)
	repo := &models.Repository{}
	err := conn.Where("id = ?", intrepoId).Delete(&repo)
	return err.Error
}

func ScanRepositoryById(repoId string) (models.ScanResults, error) {
	findingsRes := &[]models.Findings{} // initialisse response obj

	// get the url from DB
	intrepoId, _ := strconv.Atoi(repoId)
	result := models.Repository{}
	err := conn.Where("id = ?", intrepoId).Find(&result)
	if err.Error != nil {
		return models.ScanResults{}, err.Error
	}

	fileUrl := result.RepoLink
	base := path.Base(fileUrl)

	clonseRepo(fileUrl)                  //clone to temp dir
	defer os.RemoveAll("./temp/" + base) // remove repo in temp dir

	scanStart := time.Now()
	scanDir("./temp/"+base, findingsRes) //scan the dir
	scanEnd := time.Now()

	scanResObj := createScanResultsObj(findingsRes, scanStart, scanEnd, result)

	return scanResObj, nil
}
