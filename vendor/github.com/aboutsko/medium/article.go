package medium

import (
	"fmt"

	"net/url"
)

var (
	urlTop    = "https://medium.com/topic/popular?format=json"
	urlPost   = "https://medium.com/p/%s?format=json"
	urlSearch = "https://medium.com/search?q=%s&format=json"
)

//Article is the root response structure from medium get apis
type Article struct {
	Payload *Payload `json:"payload"`
}

//GetMostPopularPosts returns unlogged most popular posts (medium.com/topic/popular homepage)
func GetMostPopularPosts() (*Article, error) {
	return HTTPGetArticle(urlTop)
}

//GetArticle returns one article identified by its uid
func GetArticle(articleID string) (*Article, error) {
	return HTTPGetArticle(fmt.Sprintf(urlPost, articleID))
}

//SearchArticles returns posts matching search query
func SearchArticles(query string) (*Article, error) {
	return HTTPGetArticle(fmt.Sprintf(urlSearch, url.QueryEscape(query)))
}
