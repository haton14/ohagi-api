package schema

type Food struct {
	ID            *int   `json:"id,omitempty"`
	Name          string `json:"name"`
	Amount        int    `json:"amount,omitempty"`
	Unit          string `json:"unit"`
	LastUpdatedAt *int64 `json:"last_updated_at,omitempty"`
}
