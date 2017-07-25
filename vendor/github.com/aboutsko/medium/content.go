package medium

import "fmt"

//Content defines an article content
type Content struct {
	BodyModel *BodyModel `json:"bodyModel"`
}

func (content *Content) String() string {
	if content.BodyModel != nil {
		return fmt.Sprintf("%s", content.BodyModel)

	}
	return ""
}
