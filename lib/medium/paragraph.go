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

func (p *Paragraph) Format() string {
	formatted := wrap.Wrap(p.Text, 80)
	if p.Type == 3 {
		formatted = fmt.Sprintf("\033[1m%s\033[0m", formatted)
	}

	if p.Type > 1 {
		formatted = fmt.Sprintf("%s", formatted)
	}

	return formatted
}
