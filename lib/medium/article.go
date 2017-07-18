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
	Payload Payload `json:"payload"`
}

func (a *Article) Format() string {
	formatted := ""
	for _, paragraph := range a.Payload.Value.Content.BodyModel.Paragraphs {
		formatted = fmt.Sprintf("%s\n%s", formatted, paragraph.Format())
	}

	for uid, post := range a.Payload.References.Posts {
		formatted = fmt.Sprintf("%s\n\033[34m%s\033[0m ", formatted, uid)
		for _, paragraph := range post.PreviewContent.BodyModel.Paragraphs {
			formatted = fmt.Sprintf("%s%s", formatted, paragraph.Format())
		}
	}

	formatted = fmt.Sprintf("%s\n\nurl: \033[32m%s\033[0m\n\n", formatted, a.Payload.Value.CanonicalUrl)

	return formatted
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
