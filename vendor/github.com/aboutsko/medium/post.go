package medium

import "fmt"

//Post is a part of an article
type Post struct {
	Description    string   `json:"description"`
	Title          string   `json:"title"`
	PreviewContent *Content `json:"previewContent"`
}

func (post *Post) String() string {
	formatted := fmt.Sprintf("\033[1m%s\033[0m\n%s", post.Title, post.Description)
	if post.PreviewContent != nil {
		formatted = fmt.Sprintf("%s\n%s", formatted, post.PreviewContent)
	}
	return formatted
}
