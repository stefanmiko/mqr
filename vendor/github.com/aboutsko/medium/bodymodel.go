package medium

import "fmt"

//BodyModel defines the structure of a post
type BodyModel struct {
	Paragraphs []*Paragraph `json:"paragraphs"`
}

func (bodyModel *BodyModel) String() string {
	formatted := ""
	for _, paragraph := range bodyModel.Paragraphs {
		if paragraph != nil {
			formatted = fmt.Sprintf("%s\n%s", formatted, paragraph)
		}
	}
	return formatted
}
