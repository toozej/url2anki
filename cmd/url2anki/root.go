package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/automaxprocs/maxprocs"

	"github.com/toozej/url2anki/internal/url2anki"
	"github.com/toozej/url2anki/pkg/man"
	"github.com/toozej/url2anki/pkg/version"
)

var rootCmd = &cobra.Command{
	Use:              "url2anki",
	Short:            "Generate Anki flashcards from a URL",
	Long:             `Generate Anki-formatted flashcards from a given URL and export them to a file to be imported into Anki`,
	Args:             cobra.ExactArgs(0),
	PersistentPreRun: rootCmdPreRun,
	Run:              url2anki.Run,
}

func rootCmdPreRun(cmd *cobra.Command, args []string) {
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		return
	}
	if viper.GetBool("debug") {
		log.SetLevel(log.DebugLevel)
	}

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func init() {
	_, err := maxprocs.Set()
	if err != nil {
		log.Error("Error setting maxprocs: ", err)
	}

	// create rootCmd-level flags
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "Enable debug-level logging")
	rootCmd.Flags().StringP("url", "u", "", "The URL to scrape for flashcards (EX: https://kubernetes.io/docs/reference/glossary/?all=true)")
	rootCmd.Flags().StringP("question-selector", "q", "", "The HTML selector for the questions (EX: div.term-name)")
	rootCmd.Flags().StringP("answer-selector", "a", "", "The HTML selector for the answers (EX: div.term-definition)")
	rootCmd.Flags().StringP("output-file", "o", "./anki_cards.csv", "The filename (including extension) to export flashcards to")
	rootCmd.Flags().BoolP("preview", "p", false, "Preview the flashcards before exporting")

	// Mark required flags
	_ = cobra.MarkFlagRequired(rootCmd.Flags(), "url")
	_ = cobra.MarkFlagRequired(rootCmd.Flags(), "question-selector")
	_ = cobra.MarkFlagRequired(rootCmd.Flags(), "answer-selector")

	// add sub-commands
	rootCmd.AddCommand(
		man.NewManCmd(),
		version.Command(),
	)
}
