// Delete a notebook returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	// there is a valid "notebook" in the system
	NOTEBOOK_DATA_ID, _ := strconv.ParseInt(os.Getenv("NOTEBOOK_DATA_ID"), 10, 64)

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	r, err := apiClient.NotebooksApi.DeleteNotebook(ctx, NOTEBOOK_DATA_ID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotebooksApi.DeleteNotebook`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
