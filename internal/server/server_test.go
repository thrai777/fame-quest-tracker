package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/thrai777/fame-quest-tracker/internal/domain"
)

func TestHealth(t *testing.T) {
	s := New(domain.SeedQuests())
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	s.Handler().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
	var body map[string]string
	if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
		t.Fatalf("decode health response: %v", err)
	}
	if body["status"] != "ok" {
		t.Errorf("expected status ok, got %q", body["status"])
	}
}

func TestListQuests(t *testing.T) {
	seed := domain.SeedQuests()
	s := New(seed)
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/quests", nil)
	s.Handler().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
	var quests []domain.Quest
	if err := json.NewDecoder(rec.Body).Decode(&quests); err != nil {
		t.Fatalf("decode quests response: %v", err)
	}
	if len(quests) != len(seed) {
		t.Fatalf("expected %d quests, got %d", len(seed), len(quests))
	}
	if quests[0].Name != seed[0].Name {
		t.Errorf("expected first quest %q, got %q", seed[0].Name, quests[0].Name)
	}
}
