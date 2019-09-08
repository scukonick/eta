package structs

type Task struct {
	ID  string  `json:"id"`
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
	Tag uint64  `json:"-"`
	ETA int     `json:"-"`
}
