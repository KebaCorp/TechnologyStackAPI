package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/KebaCorp/TechnologyStackAPI/internal/app/model"
)

// Returns all users
func (s *APIServer) handleUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		types, err := s.store.User().FindAll()

		if err != nil {
			s.error(w, r, http.StatusNotFound, err)

			return
		}

		s.respond(w, r, http.StatusOK, types)
	}
}

// Create user
func (s *APIServer) handleUserCreate() http.HandlerFunc {
	type request struct {
		Email      string `json:"email"`
		Username   string `json:"username"`
		FirstName  string `json:"firstName"`
		LastName   string `json:"lastName"`
		MiddleName string `json:"middleName"`
		Image      string `json:"image"`
		IsActive   string `json:"isActive"`
		Password   string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)

			return
		}

		t := &model.User{
			Email:         req.Email,
			Username:      req.Username,
			FirstName:     req.FirstName,
			LastName:      req.LastName,
			MiddleName:    req.MiddleName,
			Image:         req.Image,
			IsActive:      true,
			Password:      req.Password,
			CreatorUserId: 1,
		}

		id, err := s.store.User().CreateUser(t)

		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
		}

		s.respond(w, r, http.StatusCreated, id)
	}
}

// Delete user
func (s *APIServer) handleUserDelete() http.HandlerFunc {
	type request struct {
		ID int `json:"id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)

			return
		}

		if err := s.store.User().DeleteUser(req.ID); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
		}

		s.respond(w, r, http.StatusCreated, true)
	}
}
