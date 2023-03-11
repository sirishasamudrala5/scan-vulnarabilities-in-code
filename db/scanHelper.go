package db

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"bitbucket.org/guardrails-go/models"
)

func clonseRepo(fileUrl string) {
	// clone to temp dir
	cmd := exec.Command("git", "clone", fileUrl)
	cmd.Dir = "./temp"

	err1 := cmd.Run()

	if err1 != nil {
		// something went wrong
		fmt.Println(err1)
	}
}

func scanFile(filePath string, findingsRes *[]models.Findings) {

	f, _ := os.Open(filePath)
	defer f.Close()

	// Splits on newlines by default.
	scanner := bufio.NewScanner(f)

	line := 1
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "public_key") {
			appendToResObj(filePath, "public_key", line, findingsRes)
		}
		if strings.Contains(scanner.Text(), "private_key") {
			appendToResObj(filePath, "private_key", line, findingsRes)
		}
		line++
	}

	if err := scanner.Err(); err != nil {
		// Handle the error
		fmt.Println(err)
	}
}

func scanDir(dirPath string, findingsRes *[]models.Findings) {
	// walk through the dir
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		// for each file - scan for the vulnarability key words
		scanFile(path, findingsRes)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}

func appendToResObj(filePath string, key string, line int, findingsRes *[]models.Findings) {
	if key != "" {
		newFinding := &models.Findings{}
		newFinding.Type = "sast"
		if key == "public_key" {
			newFinding.RuleId = "G402"
		} else if key == "private_key" {
			newFinding.RuleId = "G404"
		} else {
			newFinding.RuleId = "unknown"
		}

		positions := &models.Positions{}
		positions.Begin = map[string]interface{}{"line": line}
		newFinding.Location = map[string]interface{}{"path": strings.ReplaceAll(filePath, "temp/", ""), "positions": *positions}
		newFinding.MetaData.Description = "Use of Public or Private keys vulnerability"
		newFinding.MetaData.Severity = "High"
		*findingsRes = append(*findingsRes, *newFinding)
	}
}

func createScanResultsObj(findings *[]models.Findings, scanStart time.Time, scanEnd time.Time, repo models.Repository) models.ScanResults {
	scanRes := &models.ScanResults{}
	scanRes.Findings = map[string]interface{}{"findings": *findings}
	scanRes.RepositoryName = repo.Name
	scanRes.RepositoryURL = repo.RepoLink
	scanRes.ScanStatus = "success"
	scanRes.ScanStatedAt = scanStart
	scanRes.ScanEndedAt = scanEnd

	return *scanRes

}
