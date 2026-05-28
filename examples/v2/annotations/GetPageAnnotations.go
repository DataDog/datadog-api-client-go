// Get annotations for a page returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	// there is a valid "annotation" in the system
	AnnotationDataAttributesPageID := os.Getenv("ANNOTATION_DATA_ATTRIBUTES_PAGE_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.GetPageAnnotations", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAnnotationsApi(apiClient)
	resp, r, err := api.GetPageAnnotations(ctx, AnnotationDataAttributesPageID, 1704067200000, 1704153600000)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnnotationsApi.GetPageAnnotations`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AnnotationsApi.GetPageAnnotations`:\n%s\n", responseContent)
}
