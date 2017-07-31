package medium

//Post is a part of an article
type Post struct {
	Description    string   `json:"description"`
	Title          string   `json:"title"`
	PreviewContent *Content `json:"previewContent"`
}
