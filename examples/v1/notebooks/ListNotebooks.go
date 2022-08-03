// Get all notebooks returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/v2/api/v1/datadog"
)

func main() {
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewNotebooksApi(apiClient)
	resp, r, err := api.ListNotebooks(ctx, *datadog.NewListNotebooksOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotebooksApi.ListNotebooks`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `NotebooksApi.ListNotebooks`:\n%s\n", responseContent)
}
