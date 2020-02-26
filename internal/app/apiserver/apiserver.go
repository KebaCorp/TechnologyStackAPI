package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/KebaCorp/TechnologyStackAPI/internal/app/model"
	"github.com/KebaCorp/TechnologyStackAPI/internal/app/store"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// APIServer ...
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

// New ...
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start ...
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("Starting API Server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st

	return nil
}

func (s *APIServer) configureRouter() {
	s.router.Use(handlers.CORS(
		handlers.AllowedOrigins([]string{s.config.CorsOrigin}),
		handlers.AllowedHeaders([]string{"content-type"}),
	))

	s.router.HandleFunc("/api/v1/technologies", s.handleTechnologies())

	s.router.HandleFunc("/api/v1/types", s.handleTypes())
	s.router.HandleFunc("/api/v1/type/create", s.handleTypeCreate()).Methods(http.MethodPost, http.MethodOptions)
	s.router.HandleFunc("/api/v1/type/delete", s.handleTypeDelete()).Methods(http.MethodPost, http.MethodOptions)

	s.router.HandleFunc("/api/v1/stages", s.handleStages())
	s.router.HandleFunc("/api/v1/stage/create", s.handleStageCreate()).Methods(http.MethodPost, http.MethodOptions)
	s.router.HandleFunc("/api/v1/stage/delete", s.handleStageDelete()).Methods(http.MethodPost, http.MethodOptions)
}

// Returns all technologies
func (s *APIServer) handleTechnologies() http.HandlerFunc {
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

// Returns all stages
func (s *APIServer) handleStages() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		types, err := s.store.Stage().FindAll()

		if err != nil {
			s.error(w, r, http.StatusNotFound, err)

			return
		}

		s.respond(w, r, http.StatusOK, types)
	}
}

// Create stage
func (s *APIServer) handleStageCreate() http.HandlerFunc {
	type request struct {
		Title string `json:"title"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)

			return
		}

		t := &model.Stage{
			Title:         req.Title,
			IsDeleted:     false,
			CreatorUserId: 1,
		}

		id, err := s.store.Stage().CreateStage(t)

		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
		}

		s.respond(w, r, http.StatusCreated, id)
	}
}

// Delete stage
func (s *APIServer) handleStageDelete() http.HandlerFunc {
	type request struct {
		ID int `json:"id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)

			return
		}

		if err := s.store.Stage().DeleteStage(req.ID); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
		}

		s.respond(w, r, http.StatusCreated, true)
	}
}

func (s *APIServer) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *APIServer) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)

	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
