package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/KebaCorp/TechnologyStackAPI/internal/app/model"
)

// Returns all projects
func (s *APIServer) handleProjects() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		types, err := s.store.Project().FindAll()

		if err != nil {
			s.error(w, r, http.StatusNotFound, err)

			return
		}

		s.respond(w, r, http.StatusOK, types)
	}
}

// Create project
func (s *APIServer) handleProjectCreate() http.HandlerFunc {
	type request struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Code        string `json:"code"`
		Image       string `json:"image"`
		IsActive    string `json:"isActive"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)

			return
		}

		t := &model.Project{
			Title:         req.Title,
			Description:   req.Description,
			Code:          req.Code,
			Image:         req.Image,
			IsActive:      true,
			CreatorUserId: 1,
		}

		id, err := s.store.Project().CreateProject(t)

		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
		}

		s.respond(w, r, http.StatusCreated, id)
	}
}

// Delete project
func (s *APIServer) handleProjectDelete() http.HandlerFunc {
	type request struct {
		ID int `json:"id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)

			return
		}

		if err := s.store.Project().DeleteProject(req.ID); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
		}

		s.respond(w, r, http.StatusCreated, true)
	}
}
