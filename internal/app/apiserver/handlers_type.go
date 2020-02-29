package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/KebaCorp/TechnologyStackAPI/internal/app/model"
)

// Returns all types
func (s *APIServer) handleTypes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		types, err := s.store.Type().FindAll()

		if err != nil {
			s.error(w, r, http.StatusNotFound, err)

			return
		}

		s.respond(w, r, http.StatusOK, types)
	}
}

// Create type
func (s *APIServer) handleTypeCreate() http.HandlerFunc {
	type request struct {
		Title string `json:"title"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)

			return
		}

		t := &model.Type{
			Title:         req.Title,
			IsDeleted:     false,
			CreatorUserId: 1,
		}

		id, err := s.store.Type().CreateType(t)

		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
		}

		s.respond(w, r, http.StatusCreated, id)
	}
}

// Delete type
func (s *APIServer) handleTypeDelete() http.HandlerFunc {
	type request struct {
		ID int `json:"id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)

			return
		}

		if err := s.store.Type().DeleteType(req.ID); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
		}

		s.respond(w, r, http.StatusCreated, true)
	}
}
