package store

import (
	"database/sql"

	_ "github.com/lib/pq" // ...
)

// Store ...
type Store struct {
	config               *Config
	db                   *sql.DB
	userRepository       *UserRepository
	technologyRepository *TechnologyRepository
}

// New ...
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Open ...
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

// Close ...
func (s *Store) Close() {
	s.db.Close()
}

// User ...
func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}

// Technology ...
func (s *Store) Technology() *TechnologyRepository {
	if s.technologyRepository != nil {
		return s.technologyRepository
	}

	s.technologyRepository = &TechnologyRepository{
		store: s,
	}

	return s.technologyRepository
}
