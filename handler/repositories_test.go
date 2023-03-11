package handler

import (
	"log"
	"net/http"
	"testing"

	"bitbucket.org/guardrails-go/mock_http"
	"bitbucket.org/guardrails-go/models"
)

func TestCreateRepository(t *testing.T) {
	newList := models.NewRepositoryList()

	headers := http.Header{}
	headers.Add("content-type", "application/json")

	w := &mock_http.ResponseWriter{}
	r := &http.Request{
		Header: headers,
	}

	r.Body = mock_http.RequestBody(map[string]string{
		"name":     "hello",
		"repo_url": "world",
	})

	createRepository(w, r)

	result := w.GetBodyString()
	log.Println(result)

	if len(newList.GetAll()) != 1 {
		t.Errorf("Item did not add")
	}

	if newList.GetAll()[0].Name != "hello" {
		t.Errorf("Item bad")
	}
}
