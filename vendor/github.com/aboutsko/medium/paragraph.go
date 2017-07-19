package medium

import (
	"fmt"

	"github.com/bbrks/wrap"
)

type Paragraph struct {
	Name string `json:"name"`
	Text string `json:"text"`
	Type int    `json:"type"`
}

func (paragraph *Paragraph) String() string {
	formatted := wrap.Wrap(paragraph.Text, 80)
	if paragraph.Type == 3 {
		formatted = fmt.Sprintf("\033[1m%s\033[0m", formatted)
	}

	if paragraph.Type > 1 {
		formatted = fmt.Sprintf("%s", formatted)
	}

	return formatted

}
