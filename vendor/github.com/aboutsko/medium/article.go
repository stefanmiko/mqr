package medium

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var urlTop = "https://medium.com/topic/popular?format=json"
var urlPost = "https://medium.com/p/%s?format=json"

//Article is the root response structure from medium get apis
type Article struct {
	Payload *Payload `json:"payload"`
}

func (article *Article) String() string {
	return fmt.Sprintf("%s", article.Payload)
}

//GetMostPopularPosts returns unlogged most popular posts (medium.com/topic/popular homepage)
func GetMostPopularPosts() (*Article, error) {
	return getURL(urlTop)
}

//GetArticle returns one article identified by its uid
func GetArticle(articleID string) (*Article, error) {
	return getURL(fmt.Sprintf(urlPost, articleID))
}

func getURL(url string) (*Article, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	body = body[16:]

	article := &Article{}
	err = json.Unmarshal(body, article)
	if err != nil {
		return nil, err
	}

	return article, err
}
