package main

import (
	"fmt"
	"os"
	"time"

	"github.com/khulnasoft/evolution/utils/pkg/utils"
	"github.com/spf13/cobra"
)

var (
	latestUpdateTextStartTag = "<!-- LATEST-UPDATE -->"
	latestUpdateTextEndTag   = "<!-- /LATEST-UPDATE -->"
)

var rootCmd = &cobra.Command{
	Use:   "utils",
	Short: "utils - CLI tool for managing khulnasoft/evolution",
	Run: func(c *cobra.Command, args []string) {
		c.Help()
	},
}

func latestUpdateTextEditor(s string) (string, error) {
	if len(s) == 0 {
		s = latestUpdateTextStartTag + latestUpdateTextEndTag
	}

	str := time.Now().UTC().Format(time.RFC3339)
	return utils.ReplaceTextTags(s, latestUpdateTextStartTag, latestUpdateTextEndTag, str)
}

func main() {
	rootCmd.AddCommand(readmeCmd)
	rootCmd.AddCommand(maintainersCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "CLI error: %s\n", err)
		os.Exit(1)
	}
}
