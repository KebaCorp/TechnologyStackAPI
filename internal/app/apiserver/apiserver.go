package apiserver

import (
	"encoding/json"
	"net/http"

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
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{s.config.CorsOrigin})))
	s.router.HandleFunc("/api/v1/technologies", s.handleTechnologies())
}

func (s *APIServer) handleTechnologies() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		technologies, err := s.store.Technology().FindAll()
		stages, err := s.store.Stage().FindAll()
		types, err := s.store.Type().FindAll()

		for tsKey, tsValue := range types {
			for _, stValue := range stages {
				for _, tchValue := range technologies {
					if tsValue.ID == tchValue.TypeId && stValue.ID == tchValue.StageId {
						stValue.Technologies = append(stValue.Technologies, tchValue)
						break
					}
				}
				types[tsKey].Stages = append(types[tsKey].Stages, stValue)
			}
		}

		js, err := json.Marshal(types)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}
