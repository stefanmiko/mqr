package cmd

import (
	"fmt"
	"os"

	"github.com/aboutsko/medium"
	"github.com/aboutsko/mqr/formatter"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "mqr",
	Short: "Medium Quick Reader",
	Long:  `mqr makes reading medium in your terminal easy and friendly`,
	Run: func(cmd *cobra.Command, args []string) {
		article, err := medium.GetMostPopularPosts()
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
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mqr.yaml)")
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".mqr")
	}

	viper.AutomaticEnv() // read in environment variables that match

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
