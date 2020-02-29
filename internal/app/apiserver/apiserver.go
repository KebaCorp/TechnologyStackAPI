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
	// CORS middleware
	s.router.Use(handlers.CORS(
		handlers.AllowedOrigins([]string{s.config.CorsOrigin}),
		handlers.AllowedHeaders([]string{"content-type"}),
	))

	// Authorization handlers
	s.router.HandleFunc("/api/v1/authorization/login", s.handleAuthorizationLogin())

	// Dashboard handlers
	s.router.HandleFunc("/api/v1/dashboard", s.handleDashboard())

	// Technology handlers
	s.router.HandleFunc("/api/v1/technologies", s.handleTechnologies())
	s.router.HandleFunc("/api/v1/technology/create", s.handleTechnologyCreate()).Methods(http.MethodPost, http.MethodOptions)
	s.router.HandleFunc("/api/v1/technology/delete", s.handleTechnologyDelete()).Methods(http.MethodPost, http.MethodOptions)

	// Technology item handlers
	s.router.HandleFunc("/api/v1/technology-items", s.handleTechnologyItems())
	s.router.HandleFunc("/api/v1/technology-item/create", s.handleTechnologyItemCreate()).Methods(http.MethodPost, http.MethodOptions)
	s.router.HandleFunc("/api/v1/technology-item/delete", s.handleTechnologyItemDelete()).Methods(http.MethodPost, http.MethodOptions)

	// Type handlers
	s.router.HandleFunc("/api/v1/types", s.handleTypes())
	s.router.HandleFunc("/api/v1/type/create", s.handleTypeCreate()).Methods(http.MethodPost, http.MethodOptions)
	s.router.HandleFunc("/api/v1/type/delete", s.handleTypeDelete()).Methods(http.MethodPost, http.MethodOptions)

	// Stage handlers
	s.router.HandleFunc("/api/v1/stages", s.handleStages())
	s.router.HandleFunc("/api/v1/stage/create", s.handleStageCreate()).Methods(http.MethodPost, http.MethodOptions)
	s.router.HandleFunc("/api/v1/stage/delete", s.handleStageDelete()).Methods(http.MethodPost, http.MethodOptions)

	// Project handlers
	s.router.HandleFunc("/api/v1/projects", s.handleProjects())
	s.router.HandleFunc("/api/v1/project/create", s.handleProjectCreate()).Methods(http.MethodPost, http.MethodOptions)
	s.router.HandleFunc("/api/v1/project/delete", s.handleProjectDelete()).Methods(http.MethodPost, http.MethodOptions)

	// User handlers
	s.router.HandleFunc("/api/v1/users", s.handleUsers())
	s.router.HandleFunc("/api/v1/user/create", s.handleUserCreate()).Methods(http.MethodPost, http.MethodOptions)
	s.router.HandleFunc("/api/v1/user/delete", s.handleUserDelete()).Methods(http.MethodPost, http.MethodOptions)
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

// Get IP
func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}
