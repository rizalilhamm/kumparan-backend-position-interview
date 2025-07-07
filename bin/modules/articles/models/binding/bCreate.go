package binding

import "time"

type Create struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	AuthorID  string    `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
}
