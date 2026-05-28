// Delete an annotation returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	// there is a valid "annotation" in the system
	AnnotationDataID := uuid.MustParse(os.Getenv("ANNOTATION_DATA_ID"))

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.DeleteAnnotation", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAnnotationsApi(apiClient)
	r, err := api.DeleteAnnotation(ctx, AnnotationDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnnotationsApi.DeleteAnnotation`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
