package models

import (
	"net/http"
	"time"
)

type Positions struct {
	Begin map[string]interface{} `gorm:"serializer:json"`
}

type Location struct {
	Path      string                 `json:"path" bson:"path"`
	Positions map[string]interface{} `gorm:"serializer:json"`
}

type MetaData struct {
	Description string `json:"description" bson:"description"`
	Severity    string `json:"severity" bson:"severity"`
}

type Findings struct {
	Type     string                 `json:"type" bson:"type"`
	RuleId   string                 `json:"ruleId" bson:"ruleId"`
	Location map[string]interface{} `gorm:"serializer:json" json:"location" bson:"location"`
	MetaData MetaData               `gorm:"serializer:json" json:"metadata" bson:"metadata"`
}

type ScanResults struct {
	ID             uint
	RepositoryName string                 `json:"repository_name" bson:"repository_name"`
	RepositoryURL  string                 `json:"repository_url" bson:"repository_url"`
	ScanStatus     string                 `json:"scan_status" bson:"scan_status"`
	Findings       map[string]interface{} `gorm:"serializer:json" json:"findings" bson:"findings"`
	ScanStatedAt   time.Time              `json:"scan_started_at" bson:"scan_started_at"`
	ScanEndedAt    time.Time              `json:"scan_ended_at" bson:"scan_ended_at"`
}

func (*ScanResults) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
