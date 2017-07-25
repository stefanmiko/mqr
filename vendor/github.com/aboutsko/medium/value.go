package medium

import "fmt"

//Value is the post detail
type Value struct {
	Content      *Content `json:"content"`
	CanonicalURL string   `json:"canonicalUrl"`
}

func (value *Value) String() string {
	if value != nil {
		return fmt.Sprintf("%s\n%s", value.Content, value.CanonicalURL)
	}
	return ""

}
