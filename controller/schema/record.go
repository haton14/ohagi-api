package schema

type Record struct {
	ID            int    `json:"id,omitempty"`
	Foods         []Food `json:"foods"`
	LastUpdatedAt int64  `json:"last_updated_at,omitempty"`
	CreatedAt     int64  `json:"created_at,omitempty"`
}
