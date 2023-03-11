package db

import (
	"bitbucket.org/guardrails-go/models"
)

func AddScanResult(newScanResult *models.ScanResults) (models.ScanResults, error) {
	err := conn.Create(&newScanResult)
	return *newScanResult, err.Error
}

func GetAllScannedResults() ([]models.ScanResults, error) {
	result := []models.ScanResults{}
	conn.Find(&result)
	return result, nil
}
