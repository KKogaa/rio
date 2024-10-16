package send

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/KKogaa/rio/internal/core/entities"
	"github.com/KKogaa/rio/internal/core/services"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/tidwall/pretty"
)

var SendCmd = &cobra.Command{
	Use:   "send",
	Short: "sends an http request based on a rio spec",
	Run: func(cmd *cobra.Command, args []string) {

		verbose, _ := cmd.Flags().GetBool("verbose")
		requestName, _ := cmd.Flags().GetString("name")

		requestFacade := services.NewRequestFacade(services.NewFileService(),
			services.NewRequestService())

		var request entities.Request
		var response entities.Response
		var err error

		if requestName != "" {
			request, response, err = requestFacade.Send(requestName)
		} else {
			if len(args) == 0 {
				log.Fatal("No request name or file path provided")
			}
			request, response, err = requestFacade.SendByPath(args[0])
		}

		if err != nil {
			log.Fatal(err)
		}

		green := color.New(color.FgGreen)

		if verbose {
			fmt.Println(green.Sprint("Request:"))
			PrettyPrintStruct(request)
			fmt.Println(green.Sprint("Response:"))
		}
		PrettyPrintStruct(response)
	},
}

func init() {
	SendCmd.Flags().BoolP("verbose", "v", false,
		"Show verbose output")
	SendCmd.Flags().StringP("name", "n", "",
		"Execute a request by name")
}

func PrettyPrintStruct[T any](data T) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("error marshalling data: ", err)
	}
	prettyJsonData := pretty.Pretty(jsonData)
	fmt.Println(string(pretty.Color(prettyJsonData, nil)))
}
