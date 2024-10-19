package list

import (
	"log"

	"github.com/KKogaa/rio/internal/core/services"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all the rio specs in the working directory for use with the send cmd",
	Run: func(cmd *cobra.Command, args []string) {
		fileService := services.NewFileService()

		list, err := fileService.SearchCwd()
		if err != nil {
			log.Fatal(err)
		}
		green := color.New(color.FgGreen)
		green.Println("List of specs in cwd:")
		for _, spec := range list {
			green.Printf("%-15s %-15s %x\n", spec.SpecName, spec.Filename, spec.Hash)
		}
	},
}
