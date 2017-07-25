package formatter

import (
	"fmt"

	"github.com/aboutsko/medium"
)

type FormatOptions struct {
	Color                  bool
	FormatReference        bool
	FormatReferenceTitle   bool
	FormatReferenceUID     bool
	FormatReferenceContent bool
	FormatValue            bool
	FormatValueURL         bool
}

func FormatValue(value *medium.Value, options *FormatOptions) string {
	formatted := ""
	if options.FormatValueURL {
		formatted = fmt.Sprintf("%s\nurl: %s \n", formatted, value.CanonicalURL)
	}

	formatted = fmt.Sprintf("%s%s\n", formatted, value.Content)
	return formatted
}

func FormatReference(reference *medium.References, options *FormatOptions) string {
	formatted := ""
	for uid, post := range reference.Posts {
		if options.FormatReferenceUID {
			formatted = fmt.Sprintf("%s%s\t", formatted, uid)
		}
		if options.FormatReferenceTitle {
			formatted = fmt.Sprintf("%s%s\n", formatted, post.Title)
		}

		if options.FormatReferenceContent {
			formatted = fmt.Sprintf("%s\n%s\n", formatted, post.PreviewContent)
		}
	}

	return formatted
}

func FormatArticle(article *medium.Article, options *FormatOptions) string {
	if article == nil || article.Payload == nil {
		return ""
	}
	formatted := ""
	if options.FormatValue && article.Payload.Value != nil {
		formatted = fmt.Sprintf("%s%s", formatted, FormatValue(article.Payload.Value, options))
	}

	if options.FormatValue && article.Payload.References != nil {
		formatted = fmt.Sprintf("%s%s", formatted, FormatReference(article.Payload.References, options))
	}

	return formatted
}
