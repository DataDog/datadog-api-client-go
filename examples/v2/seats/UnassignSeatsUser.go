// Unassign seats from users returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	// there is a valid "user" in the system
	UserDataID := os.Getenv("USER_DATA_ID")

	body := datadogV2.UnassignSeatsUserRequest{
		Data: &datadogV2.UnassignSeatsUserRequestData{
			Type: datadogV2.SEATASSIGNMENTSDATATYPE_SEAT_ASSIGNMENTS,
			Attributes: datadogV2.UnassignSeatsUserRequestDataAttributes{
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
	r, err := api.UnassignSeatsUser(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeatsApi.UnassignSeatsUser`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
