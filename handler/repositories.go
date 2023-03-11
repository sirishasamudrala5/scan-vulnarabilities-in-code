package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"bitbucket.org/guardrails-go/db"
	"bitbucket.org/guardrails-go/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

var repositoryIDKey = "repositoryId"

func repositories(router chi.Router) {
	router.Get("/", getAllRepositories)
	router.Post("/", createRepository)
	router.Get("/scan-results", getScannedResults)
	router.Route("/{repositoryId}", func(router chi.Router) {
		router.Use(RepositoryContext)
		router.Get("/", getRepositoryById)
		router.Get("/scan", scanRepository)
		router.Put("/", updateRepository)
		router.Delete("/", deleteRepository)
	})
}

func RepositoryContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		repositoryId := chi.URLParam(r, "repositoryId")
		if repositoryId == "" {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("repositoryId is required")))
			return
		}
		ctx := context.WithValue(r.Context(), repositoryIDKey, repositoryId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllRepositories(w http.ResponseWriter, r *http.Request) {

	items, err := db.GetAllRepositories()
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}

	json_items, _ := json.Marshal(items)
	w.Write(json_items)
}

func getRepositoryById(w http.ResponseWriter, r *http.Request) {
	repositoryID := r.Context().Value(repositoryIDKey).(string)
	repository, err := db.GetRepositoryById(repositoryID)
	if err != nil {
		if err == ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrorRenderer(err))
		}
		return
	}
	if err := render.Render(w, r, &repository); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func createRepository(w http.ResponseWriter, r *http.Request) {
	repository := &models.Repository{}
	if err := render.Bind(r, repository); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	if err := db.AddRepository(repository); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, repository); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func updateRepository(w http.ResponseWriter, r *http.Request) {
	repositoryId := r.Context().Value(repositoryIDKey).(string)

	fmt.Println("in update handler", repositoryId)

	repositoryData := models.Repository{}
	if err := render.Bind(r, &repositoryData); err != nil {
		fmt.Println("bind error", err)
		render.Render(w, r, ErrBadRequest)
		return
	}
	repository, err := db.UpdateRepository(repositoryId, repositoryData)
	if err != nil {
		if err == ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
	if err := render.Render(w, r, &repository); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func deleteRepository(w http.ResponseWriter, r *http.Request) {
	repositoryId := r.Context().Value(repositoryIDKey).(string)
	err := db.DeleteRepository(repositoryId)
	if err != nil {
		if err == ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
}

func scanRepository(w http.ResponseWriter, r *http.Request) {
	repositoryID := r.Context().Value(repositoryIDKey).(string)
	res, err := db.ScanRepositoryById(repositoryID)

	if err != nil {
		if err == ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrorRenderer(err))
		}
	}

	// store scan results
	response, err := db.AddScanResult(&res)
	if err != nil {
		if err == ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
	if err := render.Render(w, r, &response); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func getScannedResults(w http.ResponseWriter, r *http.Request) {

	items, err := db.GetAllScannedResults()
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}

	json_items, _ := json.Marshal(items)
	w.Write(json_items)
}
