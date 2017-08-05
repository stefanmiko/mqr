package formatter

import (
	"fmt"

	"github.com/aboutsko/medium"
	"github.com/bbrks/wrap"
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

func bold(str string) string {
	return fmt.Sprintf("\033[1m%s\033[0m", str)
}

func FormatContent(content *medium.Content, options *FormatOptions) string {
	formatted := ""
	if content == nil || content.BodyModel == nil {
		return ""
	}
	for _, paragraph := range content.BodyModel.Paragraphs {
		wrappedText := wrap.Wrap(paragraph.Text, 80)
		if paragraph.Type > 2 {
			formatted = fmt.Sprintf("%s%s\n", formatted, bold(wrappedText))
		} else {
			formatted = fmt.Sprintf("%s%s\n", formatted, wrappedText)
		}
	}
	return formatted
}

func FormatPost(postUID string, post *medium.Post, options *FormatOptions) string {
	formatted := ""
	if options.FormatReferenceUID {
		formatted = fmt.Sprintf("%s%s\t", formatted, bold(postUID))
	}
	if options.FormatReferenceTitle {
		formatted = fmt.Sprintf("%s%s\n", formatted, post.Title)
	}

	if options.FormatReferenceContent {
		formatted = fmt.Sprintf("%s\n%s\n", formatted, post.PreviewContent)
	}
	return formatted
}

func FormatValue(value *medium.Value, options *FormatOptions) string {
	formatted := ""
	if options.FormatValueURL {
		formatted = fmt.Sprintf("%s\nurl: %s \n", formatted, value.CanonicalURL)
	}

	formatted = fmt.Sprintf("%s%s\n", formatted, FormatContent(value.Content, options))

	for uid, post := range value.Posts {
		formatted = fmt.Sprintf("%s%s", formatted, FormatPost(uid, post, options))
	}

	return formatted
}

func FormatReference(reference *medium.References, options *FormatOptions) string {
	formatted := ""
	for uid, post := range reference.Posts {
		formatted = fmt.Sprintf("%s%s", formatted, FormatPost(uid, post, options))
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
