// Create On-Call schedule overrides returns "Created" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	// there is a valid "schedule" in the system
	ScheduleDataID := os.Getenv("SCHEDULE_DATA_ID")

	// there is a valid "user" in the system
	UserDataID := os.Getenv("USER_DATA_ID")

	body := datadogV2.CreateOverridesRequest{
		Data: []datadogV2.CreateOverrideRequestData{
			{
				Attributes: datadogV2.CreateOverrideRequestAttributes{
					End:   time.Now().AddDate(0, 0, 1),
					Start: time.Now(),
				},
				Relationships: &datadogV2.CreateOverrideRequestRelationships{
					User: &datadogV2.OverrideRelationshipsUser{
						Data: datadogV2.OverrideRelationshipsUserData{
							Id:   UserDataID,
							Type: datadogV2.OVERRIDERELATIONSHIPSUSERDATATYPE_USERS,
						},
					},
				},
				Type: datadogV2.OVERRIDEDATATYPE_OVERRIDES,
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOnCallApi(apiClient)
	resp, r, err := api.CreateScheduleOverrides(ctx, ScheduleDataID, body, *datadogV2.NewCreateScheduleOverridesOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnCallApi.CreateScheduleOverrides`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OnCallApi.CreateScheduleOverrides`:\n%s\n", responseContent)
}
