package models

type Bookmark struct {
	CommentRaw      string   `json:"comment_raw"`
	Private         bool     `json:"private"`
	Eid             string   `json:"eid"`
	CreatedEpoch    int      `json:"created_epoch"`
	Tags            []string `json:"tags"`
	Permalink       string   `json:"permalink"`
	Comment         string   `json:"comment"`
	CreatedDatetime string   `json:"created_datetime"`
	User            string   `json:"user"`
}
