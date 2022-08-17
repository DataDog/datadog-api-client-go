// Delete a notebook returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	// there is a valid "notebook" in the system
	NotebookDataID, _ := strconv.ParseInt(os.Getenv("NOTEBOOK_DATA_ID"), 10, 64)

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewNotebooksApi(apiClient)
	r, err := api.DeleteNotebook(ctx, NotebookDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotebooksApi.DeleteNotebook`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
