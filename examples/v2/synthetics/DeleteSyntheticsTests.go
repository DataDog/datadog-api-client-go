// Synthetics: Bulk delete tests returns "OK" response

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
	body := datadogV2.DeletedTestsRequestDeleteRequest{
		Data: datadogV2.DeletedTestsRequestDelete{
			Attributes: datadogV2.DeletedTestsRequestDeleteAttributes{
				PublicIds: []string{
					"",
				},
			},
			Type: datadogV2.DELETEDTESTSREQUESTTYPE_DELETE_TESTS_REQUEST.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSyntheticsApi(apiClient)
	resp, r, err := api.DeleteSyntheticsTests(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.DeleteSyntheticsTests`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.DeleteSyntheticsTests`:\n%s\n", responseContent)
}
