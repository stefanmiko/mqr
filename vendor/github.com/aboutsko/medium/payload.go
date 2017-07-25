package medium

import "fmt"

//Payload is the higher structure of an article
type Payload struct {
	Value      *Value      `json:"value"`
	References *References `json:"references"`
}

func (payload *Payload) String() string {
	formatted := ""
	if payload.References != nil {
		formatted = fmt.Sprintf("%s%s", formatted, payload.References)
	}

	if payload.Value != nil {
		formatted = fmt.Sprintf("%s%s", formatted, payload.Value)
	}

	return formatted

}
