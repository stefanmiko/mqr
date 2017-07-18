package medium

type BodyModel struct {
	Paragraphs []Paragraph `json:"paragraphs"`
}

type Content struct {
	BodyModel BodyModel `json:"bodyModel"`
}

type Value struct {
	Content      Content `json:"content"`
	CanonicalUrl string  `json:"canonicalUrl"`
}
type Post struct {
	Description    string  `json:"description"`
	Title          string  `json:"title"`
	PreviewContent Content `json:"previewContent"`
}

type References struct {
	Posts map[string]*Post `json:"Post"`
}

type Payload struct {
	Value      Value      `json:"value"`
	References References `json:"references"`
}
