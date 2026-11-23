package domain

import "testing"

func TestSeedQuests(t *testing.T) {
	quests := SeedQuests()
	if len(quests) != 2 {
		t.Fatalf("expected 2 seed quests, got %d", len(quests))
	}
	for _, q := range quests {
		if q.Status != StatusInProgress {
			t.Errorf("quest %q: expected status %q, got %q", q.ID, StatusInProgress, q.Status)
		}
		if q.Points <= 0 {
			t.Errorf("quest %q: points should be positive, got %d", q.ID, q.Points)
		}
	}
}

func TestDeliveryRemaining(t *testing.T) {
	d := Delivery{Item: "野菜", Collected: 16, Required: 30}
	if got := d.Remaining(); got != 14 {
		t.Errorf("expected 14 remaining, got %d", got)
	}

	done := Delivery{Item: "野菜", Collected: 30, Required: 30}
	if got := done.Remaining(); got != 0 {
		t.Errorf("expected 0 remaining when complete, got %d", got)
	}

	over := Delivery{Item: "野菜", Collected: 35, Required: 30}
	if got := over.Remaining(); got != 0 {
		t.Errorf("expected 0 remaining when over-delivered, got %d", got)
	}
}
