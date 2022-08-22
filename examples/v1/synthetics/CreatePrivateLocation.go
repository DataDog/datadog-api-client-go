// Create a private location returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	// there is a valid "role" in the system
	RoleDataID := os.Getenv("ROLE_DATA_ID")

	body := datadogV1.SyntheticsPrivateLocation{
		Description: "Test Example-Create_a_private_location_returns_OK_response description",
		Metadata: &datadogV1.SyntheticsPrivateLocationMetadata{
			RestrictedRoles: []string{
				RoleDataID,
			},
		},
		Name: "Example-Create_a_private_location_returns_OK_response",
		Tags: []string{
			"test:examplecreateaprivatelocationreturnsokresponse",
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewSyntheticsApi(apiClient)
	resp, r, err := api.CreatePrivateLocation(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.CreatePrivateLocation`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.CreatePrivateLocation`:\n%s\n", responseContent)
}
