// Package domain holds the core data types and rules of the tracker,
// independent of HTTP or storage concerns.
package domain

// Status represents the state of a quest (人助け).
type Status string

const (
	StatusNotAccepted Status = "not_accepted" // 未受注
	StatusInProgress  Status = "in_progress"  // 進行中
	StatusCompleted   Status = "completed"    // 完了
	StatusFailed      Status = "failed"       // 失敗
)

// Delivery tracks progress for delivery quests (納入系クエスト),
// e.g. 16 out of 30 vegetables delivered.
type Delivery struct {
	Item      string `json:"item"`
	Collected int    `json:"collected"`
	Required  int    `json:"required"`
}

// Remaining returns how many items are still needed.
func (d Delivery) Remaining() int {
	r := d.Required - d.Collected
	if r < 0 {
		return 0
	}
	return r
}

// Quest is a single fame quest (人助け).
type Quest struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Points   int       `json:"points"`             // 名声ポイント
	Route    string    `json:"route"`              // ルート名 (e.g. カイ編)
	Deadline string    `json:"deadline,omitempty"` // ゲーム内期日 (e.g. "9/1")
	Status   Status    `json:"status"`
	Note     string    `json:"note,omitempty"`
	Delivery *Delivery `json:"delivery,omitempty"`
}

// SeedQuests returns the initial sample data set.
// It mirrors the quests currently in progress in the author's playthrough,
// and will be replaced by persistent storage later.
func SeedQuests() []Quest {
	return []Quest{
		{
			ID:       "eternal-love-bouquet",
			Name:     "永遠の愛の花束",
			Points:   130,
			Route:    "カイ編",
			Deadline: "9/1",
			Status:   StatusInProgress,
			Note:     "リナリアが必要",
		},
		{
			ID:     "salamis-vegetables",
			Name:   "サラミス迎賓館への野菜納入",
			Points: 200,
			Route:  "カイ編",
			Status: StatusInProgress,
			Delivery: &Delivery{
				Item:      "野菜",
				Collected: 16,
				Required:  30,
			},
		},
	}
}
