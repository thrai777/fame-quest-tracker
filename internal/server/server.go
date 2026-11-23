// Package server exposes the tracker as a JSON HTTP API.
package server

import (
	"encoding/json"
	"net/http"

	"github.com/thrai777/fame-quest-tracker/internal/domain"
)

// Server serves the API. It currently reads from an in-memory quest list;
// persistence will be added behind this boundary later.
type Server struct {
	quests []domain.Quest
}

// New creates a Server serving the given quests.
func New(quests []domain.Quest) *Server {
	return &Server{quests: quests}
}

// Handler returns the HTTP routing table for the API.
func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", s.handleHealth)
	mux.HandleFunc("GET /api/quests", s.handleListQuests)
	return mux
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (s *Server) handleListQuests(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, s.quests)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	_ = enc.Encode(v)
}
