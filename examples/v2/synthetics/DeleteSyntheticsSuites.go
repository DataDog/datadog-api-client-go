// Bulk delete suites returns "OK" response

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
	body := datadogV2.DeletedSuitesRequestDeleteRequest{
		Data: datadogV2.DeletedSuitesRequestDelete{
			Attributes: datadogV2.DeletedSuitesRequestDeleteAttributes{
				PublicIds: []string{
					"",
				},
			},
			Type: datadogV2.DELETEDSUITESREQUESTTYPE_DELETE_SUITES_REQUEST.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSyntheticsApi(apiClient)
	resp, r, err := api.DeleteSyntheticsSuites(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.DeleteSyntheticsSuites`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.DeleteSyntheticsSuites`:\n%s\n", responseContent)
}
