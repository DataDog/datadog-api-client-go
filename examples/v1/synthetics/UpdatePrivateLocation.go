// Edit a private location returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	body := datadog.SyntheticsPrivateLocation{
		Description: "Description of private location",
		Metadata: &datadog.SyntheticsPrivateLocationMetadata{
			RestrictedRoles: []string{
				"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
			},
		},
		Name: "New private location",
		Tags: []string{
			"team:front",
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewSyntheticsApi(apiClient)
	resp, r, err := api.UpdatePrivateLocation(ctx, "location_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.UpdatePrivateLocation`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.UpdatePrivateLocation`:\n%s\n", responseContent)
}
