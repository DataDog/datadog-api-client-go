// Create an annotation returns "OK" response

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
	body := datadogV2.AnnotationCreateRequest{
		Data: datadogV2.AnnotationRequestData{
			Attributes: datadogV2.AnnotationCreateAttributes{
				Color:       datadogV2.ANNOTATIONCOLOR_BLUE,
				Description: "Deployed v2.3.1 to production.",
				PageId:      "dashboard:abc-def-xyz",
				StartTime:   1704067200000,
				Type:        datadogV2.ANNOTATIONKIND_POINT_IN_TIME,
				WidgetIds: []string{
					"1234567890",
				},
			},
			Type: datadogV2.ANNOTATIONTYPE_ANNOTATION,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateAnnotation", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAnnotationsApi(apiClient)
	resp, r, err := api.CreateAnnotation(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnnotationsApi.CreateAnnotation`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AnnotationsApi.CreateAnnotation`:\n%s\n", responseContent)
}
