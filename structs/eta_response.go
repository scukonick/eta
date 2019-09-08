package structs

type ETAResponse struct {
	ID     string `json:"id"`
	ETA    int64  `json:"eta"`
	Status string `json:"status"`
}
