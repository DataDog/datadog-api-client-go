// Unassign seats from users for a product code returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.UnassignSeatsUserRequest{
		Data: &datadogV2.UnassignSeatsUserRequestData{
			Attributes: &datadogV2.UnassignSeatsUserRequestDataAttributes{
				UserUuids: []string{},
			},
			Type: datadogV2.SEATASSIGNMENTSDATATYPE_SEAT_ASSIGNMENTS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSeatsApi(apiClient)
	r, err := api.UnassignSeatsUserV2(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeatsApi.UnassignSeatsUserV2`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
