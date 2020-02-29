package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/KebaCorp/TechnologyStackAPI/internal/app/model"
)

// Returns all technology items
func (s *APIServer) handleTechnologyItems() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		types, err := s.store.TechnologyItem().FindAll()

		if err != nil {
			s.error(w, r, http.StatusNotFound, err)

			return
		}

		s.respond(w, r, http.StatusOK, types)
	}
}

// Create technology item
func (s *APIServer) handleTechnologyItemCreate() http.HandlerFunc {
	type request struct {
		TechnologyId int    `json:"technologyId"`
		ParentId     int    `json:"parentId"`
		Title        string `json:"title"`
		Description  string `json:"description"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)

			return
		}

		t := &model.TechnologyItem{
			TechnologyId:  req.TechnologyId,
			ParentId:      req.ParentId,
			Title:         req.Title,
			Description:   req.Description,
			CreatorUserId: 1,
		}

		id, err := s.store.TechnologyItem().CreateTechnologyItem(t)

		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
		}

		s.respond(w, r, http.StatusCreated, id)
	}
}

// Delete technology item
func (s *APIServer) handleTechnologyItemDelete() http.HandlerFunc {
	type request struct {
		ID int `json:"id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)

			return
		}

		if err := s.store.TechnologyItem().DeleteTechnologyItem(req.ID); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
		}

		s.respond(w, r, http.StatusCreated, true)
	}
}
