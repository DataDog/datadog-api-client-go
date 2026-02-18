// Assign seats to users returns "Created" response

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
	// there is a valid "user" in the system
	UserDataID := os.Getenv("USER_DATA_ID")

	body := datadogV2.AssignSeatsUserRequest{
		Data: &datadogV2.AssignSeatsUserRequestData{
			Type: datadogV2.SEATASSIGNMENTSDATATYPE_SEAT_ASSIGNMENTS,
			Attributes: datadogV2.AssignSeatsUserRequestDataAttributes{
				ProductCode: "incident_response",
				UserUuids: []string{
					UserDataID,
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSeatsApi(apiClient)
	resp, r, err := api.AssignSeatsUser(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeatsApi.AssignSeatsUser`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SeatsApi.AssignSeatsUser`:\n%s\n", responseContent)
}
