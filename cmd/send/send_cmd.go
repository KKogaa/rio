package send

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/KKogaa/rio/internal/core/services"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/tidwall/pretty"
)

var SendCmd = &cobra.Command{
	Use:   "send",
	Short: "Sends an http request based on file structure",
	Run: func(cmd *cobra.Command, args []string) {
		//TODO: add the uber dependency injection
		// log.Println("executing send command")
		// log.Println("args", args)
		fileService := services.NewFileService()
		requestService := services.NewRequestService()
		requestFacade := services.NewRequestFacade(fileService, requestService)

		request, response, err := requestFacade.Send(args[0])
		if err != nil {
			log.Println("error", err)
		}

    green := color.New(color.FgGreen)
    fmt.Println(green.Sprint("Request:"))
		PrettyPrintStruct(request)
    fmt.Println(green.Sprint("Response:"))
		PrettyPrintStruct(response)
	},
}

func PrettyPrintStruct[T any](data T) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("error marshalling data: ", err)
	}
	prettyJsonData := pretty.Pretty(jsonData)
	fmt.Println(string(pretty.Color(prettyJsonData, nil)))
}
