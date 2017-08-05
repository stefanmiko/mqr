package medium

//Value is the post detail
type Value struct {
	Content      *Content         `json:"content"`
	CanonicalURL string           `json:"canonicalUrl"`
	Posts        map[string]*Post `json:"posts"`
}
