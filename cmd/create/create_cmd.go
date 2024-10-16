package create

import (
	"fmt"
	"log"

	"github.com/KKogaa/rio/internal/core/services"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var file, requestName, method, url string

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a json file structure to use for the send command",
	Run: func(cmd *cobra.Command, args []string) {
		if file == "sample.json" && len(args) == 1 {
			file = args[0]
		}

		fileService := services.NewFileService()

		name, err := fileService.CreateRequestFile(file, requestName, method, url)
		if err != nil {
			log.Fatal(err)
		}
		green := color.New(color.FgGreen)
		fmt.Println(green.Sprint("File created at: "), name)
	},
}

func init() {
	CreateCmd.Flags().StringVarP(&file, "output", "o", "sample.json",
		"File path to output the request")
	CreateCmd.Flags().StringVarP(&requestName, "requestName", "r", "TestRequest",
		"Brief description of the request")
	CreateCmd.Flags().StringVarP(&method, "method", "m", "GET",
		"HTTP Method of the request")
	CreateCmd.Flags().StringVarP(&url, "url", "u", "http://localhost:8080",
		"URL to send the HTTP request")
}
