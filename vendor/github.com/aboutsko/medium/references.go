package medium

//References are links to other articles
type References struct {
	Posts map[string]*Post `json:"Post"`
}
