package cmd

import (
	"log"

	"github.com/KKogaa/rio/cmd/create"
	"github.com/KKogaa/rio/cmd/send"
	"github.com/spf13/cobra"
)

// which commands should this handle? and how the structure should be
// rio send <path to the file?> first do this implmentation
// rio list to show all the request based on the current location
// optional add --path or do some .config
// rio <httpverb> <request url> <params like httpie>

var rootCmd = &cobra.Command{
	Use:   "rio",
	Short: "rio: make a http request based a file structure",
}

func Execute() {
	rootCmd.AddCommand(send.SendCmd)
	rootCmd.AddCommand(create.CreateCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("there was an error executing the cli: %s", err)
	}
}
