// Update App Tags returns "No Content" response

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
	body := datadogV2.UpdateAppTagsRequest{
		Data: &datadogV2.UpdateAppTagsRequestData{
			Attributes: &datadogV2.UpdateAppTagsRequestDataAttributes{
				Tags: []string{
					"team:platform",
					"service:ops",
				},
			},
			Type: datadogV2.APPTAGSTYPE_TAGS.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAppBuilderApi(apiClient)
	r, err := api.UpdateAppTags(ctx, uuid.MustParse("9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d"), body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppBuilderApi.UpdateAppTags`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
