package cmd

import (
	"log"

	"github.com/KKogaa/rio/cmd/create"
	"github.com/KKogaa/rio/cmd/list"
	"github.com/KKogaa/rio/cmd/send"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rio",
	Short: "rio: make a http request based a file structure",
	Long: `
     ____  ___ ___  
    |  _ \|_ _/ _ \ 
    | |_) || | | | |
    |  _ < | | |_| |
    |_| \_\___\___/ 

  jajajajajja

  `,
}

func Execute() {
	rootCmd.AddCommand(send.SendCmd)
	rootCmd.AddCommand(create.CreateCmd)
	rootCmd.AddCommand(list.ListCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("there was an error executing the cli: %s", err)
	}
}
