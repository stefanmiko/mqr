package cmd

import (
	"fmt"

	"github.com/aboutsko/medium"
	"github.com/aboutsko/mqr/formatter"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get fetches an article",
	Long:  `get takes an article uid as second argument fetches and display it`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, articleUID := range args {
			article, err := medium.GetArticle(articleUID)
			if err != nil {
				fmt.Print(err)
			} else {
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
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}
