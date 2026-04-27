// Create a Synthetics downtime returns "Created" response

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
	body := datadogV2.SyntheticsDowntimeRequest{
		Data: datadogV2.SyntheticsDowntimeDataRequest{
			Attributes: datadogV2.SyntheticsDowntimeDataAttributesRequest{
				IsEnabled: true,
				Name:      "Weekly maintenance",
				TestIds: []string{
					"abc-def-123",
				},
				TimeSlots: []datadogV2.SyntheticsDowntimeTimeSlotRequest{
					{
						Duration: 3600,
						Start: datadogV2.SyntheticsDowntimeTimeSlotDate{
							Day:    15,
							Hour:   10,
							Minute: 30,
							Month:  1,
							Year:   2024,
						},
						Timezone: "Europe/Paris",
					},
				},
			},
			Type: datadogV2.SYNTHETICSDOWNTIMERESOURCETYPE_DOWNTIME,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSyntheticsApi(apiClient)
	resp, r, err := api.CreateSyntheticsDowntime(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.CreateSyntheticsDowntime`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.CreateSyntheticsDowntime`:\n%s\n", responseContent)
}
