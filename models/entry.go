package models

type Entry struct {
	Title    string `json:"title"`
	URL      string `json:"url"`
	Excerpt  string `json:"excerpt"`
	ImageURL string `json:"imageUrl"`
}
