package cmd

import (
	"fmt"
	"strings"

	"github.com/aboutsko/medium"
	"github.com/aboutsko/mqr/formatter"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "search fetches posts listing from your query",
	Long:  `search fetches posts listing from your query`,
	Run: func(cmd *cobra.Command, args []string) {
		article, err := medium.SearchArticles(strings.Join(args, " "))
		if err != nil {
			fmt.Print(err)
		} else {
			fmt.Print(formatter.FormatArticle(article, &formatter.FormatOptions{
				FormatReference:        true,
				FormatReferenceContent: false,
				FormatReferenceTitle:   true,
				FormatReferenceUID:     true,
				FormatValue:            true,
				FormatValueURL:         false,
			}))
		}
	},
}

func init() {
	RootCmd.AddCommand(searchCmd)
}
