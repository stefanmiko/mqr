package medium

import "fmt"

type Url string

type Value struct {
	Content      *Content `json:"content"`
	CanonicalURL Url      `json:"canonicalUrl"`
}

func (url *Url) String() string {
	return fmt.Sprintf("url: \033[32m%s\033[0m", url)
}

func (value *Value) String() string {
	if value != nil {
		return fmt.Sprintf("%s\n%s", value.Content, value.CanonicalURL)
	}
	return ""

}
