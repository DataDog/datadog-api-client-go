// Get users with seats returns "OK" response

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
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSeatsApi(apiClient)
	resp, r, err := api.GetSeatsUsers(ctx, "incident_response", *datadogV2.NewGetSeatsUsersOptionalParameters().WithPageLimit(100))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeatsApi.GetSeatsUsers`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SeatsApi.GetSeatsUsers`:\n%s\n", responseContent)
}
