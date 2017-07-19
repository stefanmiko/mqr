package main

import (
	"fmt"
	"os"

	"github.com/aboutsko/medium"
)

func main() {
	articles := os.Args[1:]

	if len(articles) == 0 {
		articles = append(articles, "")
	}

	for _, article := range articles {
		article, err := medium.GetArticle(article)
		if err != nil {
			panic(err)
		}
		fmt.Print(article)
	}
}
