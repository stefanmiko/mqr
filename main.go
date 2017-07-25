package main

import (
	"fmt"
	"os"

	"github.com/aboutsko/medium"
	"github.com/aboutsko/mqr/formatter"
)

func getMostPopular() (*medium.Article, error) {
	article, err := medium.GetMostPopularPosts()
	if err != nil {
		return nil, err
	}
	return article, nil
}

func getArticles(articles ...string) ([]*medium.Article, error) {
	articlesFetched := []*medium.Article{}
	for _, articleID := range articles {
		article, err := medium.GetArticle(articleID)
		if err != nil {
			panic(err)
		}
		articlesFetched = append(articlesFetched, article)
	}
	return articlesFetched, nil
}

func main() {
	articles := os.Args[1:]
	articlesToPrint := []*medium.Article{}

	if len(articles) == 0 {
		article, err := getMostPopular()
		if err != nil {
			panic(err)
		}
		articlesToPrint = append(articlesToPrint, article)
	} else {
		var err error
		articlesToPrint, err = getArticles(articles...)
		if err != nil {
			panic(err)
		}
	}

	for _, article := range articlesToPrint {
		fmt.Print(formatter.FormatArticle(article, &formatter.FormatOptions{
			FormatReference:        true,
			FormatReferenceContent: false,
			FormatReferenceTitle:   true,
			FormatReferenceUID:     true,
			FormatValue:            true,
			FormatValueURL:         true,
		}))
	}
}
