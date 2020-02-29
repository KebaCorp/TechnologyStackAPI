package apiserver

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/KebaCorp/TechnologyStackAPI/internal/app/model"
)

// Authorization login
func (s *APIServer) handleAuthorizationLogin() http.HandlerFunc {
	type request struct {
		Username   string `json:"username"`
		Password   string `json:"password"`
		RememberMe bool   `json:"rememberMe"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)

			return
		}

		user, err := s.store.User().FindByUsernameOrEmail(req.Username, req.Username)
		if err != nil {
			s.error(w, r, http.StatusNotFound, err)

			return
		}

		t := &model.Token{
			UserId:    user.ID,
			UserAgent: r.Header.Get("User-Agent"),
			Ip:        GetIP(r),
			ExpiresAt: time.Now(),
		}

		token, err := s.store.Token().UpdateOrCreateToken(t)
		if err != nil {
			s.error(w, r, http.StatusNotFound, err)
		}

		s.respond(w, r, http.StatusOK, token)
	}
}
