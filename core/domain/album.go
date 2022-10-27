package domain

type Album struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Label      string   `json:"label"`
	Popularity int      `json:"popularity"`
	Feedback   Feedback `json:"feedBack"`
}
