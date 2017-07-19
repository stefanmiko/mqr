package medium

import "fmt"

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
