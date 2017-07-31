package medium

//BodyModel defines the structure of a post
type BodyModel struct {
	Paragraphs []*Paragraph `json:"paragraphs"`
}
