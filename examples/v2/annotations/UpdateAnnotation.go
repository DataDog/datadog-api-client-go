// Update an annotation returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	// there is a valid "annotation" in the system
	AnnotationDataID := uuid.MustParse(os.Getenv("ANNOTATION_DATA_ID"))

	body := datadogV2.AnnotationUpdateRequest{
		Data: datadogV2.AnnotationRequestData{
			Attributes: datadogV2.AnnotationCreateAttributes{
				Color:       datadogV2.ANNOTATIONCOLOR_GREEN,
				Description: "Updated annotation.",
				PageId:      "dashboard:abc-def-xyz",
				StartTime:   1704067200000,
				Type:        datadogV2.ANNOTATIONKIND_POINT_IN_TIME,
			},
			Type: datadogV2.ANNOTATIONTYPE_ANNOTATION,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateAnnotation", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAnnotationsApi(apiClient)
	resp, r, err := api.UpdateAnnotation(ctx, AnnotationDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnnotationsApi.UpdateAnnotation`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AnnotationsApi.UpdateAnnotation`:\n%s\n", responseContent)
}
