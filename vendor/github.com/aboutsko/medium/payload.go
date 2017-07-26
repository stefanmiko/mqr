package medium

//Payload is the higher structure of an article
type Payload struct {
	Value      *Value      `json:"value"`
	References *References `json:"references"`
}
