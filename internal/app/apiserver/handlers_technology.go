package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/KebaCorp/TechnologyStackAPI/internal/app/model"
)

// Returns all technologies
func (s *APIServer) handleTechnologies() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		types, err := s.store.Technology().FindAll()

		if err != nil {
			s.error(w, r, http.StatusNotFound, err)

			return
		}

		s.respond(w, r, http.StatusOK, types)
	}
}

// Create technology
func (s *APIServer) handleTechnologyCreate() http.HandlerFunc {
	type request struct {
		TypeId  int    `json:"typeId"`
		StageId int    `json:"stageId"`
		Title   string `json:"title"`
		Image   string `json:"image"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)

			return
		}

		t := &model.Technology{
			TypeId:        req.TypeId,
			StageId:       req.StageId,
			Title:         req.Title,
			Image:         req.Image,
			IsDeprecated:  false,
			CreatorUserId: 1,
		}

		id, err := s.store.Technology().CreateTechnology(t)

		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
		}

		s.respond(w, r, http.StatusCreated, id)
	}
}

// Delete technology
func (s *APIServer) handleTechnologyDelete() http.HandlerFunc {
	type request struct {
		ID int `json:"id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)

			return
		}

		if err := s.store.Technology().DeleteTechnology(req.ID); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
		}

		s.respond(w, r, http.StatusCreated, true)
	}
}
