package medium

//Paragraph is a section of a bodyModel
type Paragraph struct {
	Name string `json:"name"`
	Text string `json:"text"`
	Type int    `json:"type"`
}
