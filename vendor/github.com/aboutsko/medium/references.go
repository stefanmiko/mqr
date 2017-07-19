package medium

import "fmt"

type References struct {
	Posts map[string]*Post `json:"Post"`
}

func (references *References) String() string {
	formatted := ""
	for uid, post := range references.Posts {
		if post != nil {
			formatted = fmt.Sprintf(
				"%s\n\033[34m%s\033[0m %s",
				formatted,
				uid,
				post,
			)
		}
	}

	return formatted

}
