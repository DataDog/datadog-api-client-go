// Assign seats to users for a product code returns "Created" response

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
	body := datadogV2.AssignSeatsUserRequest{
		Data: &datadogV2.AssignSeatsUserRequestData{
			Attributes: &datadogV2.AssignSeatsUserRequestDataAttributes{
				UserUuids: []string{},
			},
			Type: datadogV2.SEATASSIGNMENTSDATATYPE_SEAT_ASSIGNMENTS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSeatsApi(apiClient)
	resp, r, err := api.AssignSeatsUserV2(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeatsApi.AssignSeatsUserV2`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SeatsApi.AssignSeatsUserV2`:\n%s\n", responseContent)
}
