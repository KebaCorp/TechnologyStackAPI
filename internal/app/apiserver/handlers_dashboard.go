package apiserver

import (
	"net/http"

	"github.com/KebaCorp/TechnologyStackAPI/internal/app/model"
)

// Returns technologies dashboard
func (s *APIServer) handleDashboard() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		technologies, err := s.store.Technology().FindAll()
		stages, err := s.store.Stage().FindAll()
		types, err := s.store.Type().FindAll()

		if err != nil {
			s.error(w, r, http.StatusNotFound, err)

			return
		}

		for tsKey, tsValue := range types {
			for _, stValue := range stages {
				newStage := &model.Stage{ID: stValue.ID, Title: stValue.Title}

				for _, tchValue := range technologies {
					if tsValue.ID == tchValue.TypeId && stValue.ID == tchValue.StageId {
						newStage.Technologies = append(newStage.Technologies, tchValue)
					}
				}

				types[tsKey].Stages = append(types[tsKey].Stages, newStage)
			}
		}

		s.respond(w, r, http.StatusOK, types)
	}
}
