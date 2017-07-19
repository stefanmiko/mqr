package medium

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var URL_TOP = "https://medium.com/topic/popular?format=json"
var URL_POST = "https://medium.com/p/%s?format=json"

type Article struct {
	Payload *Payload `json:"payload"`
}

func (article *Article) String() string {
	return fmt.Sprintf("%s", article.Payload)
}

func GetArticle(articleId string) (*Article, error) {
	url := URL_TOP
	if articleId != "" {
		url = fmt.Sprintf(URL_POST, articleId)
	}
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
