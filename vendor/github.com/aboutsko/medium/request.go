package medium

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//HTTPGetArticle http requests medium article and unmarshal into medium object
func HTTPGetArticle(url string) (*Article, error) {
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
