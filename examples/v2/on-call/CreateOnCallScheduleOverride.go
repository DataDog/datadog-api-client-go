// Create an override returns "Created" response

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

	body := datadogV2.OverrideRequest{
		Data: []datadogV2.OverrideCreateData{
			{
				Attributes: datadogV2.OverrideCreateDataAttributes{
					Start: time.Now(),
					End:   time.Now().Add(time.Hour * 1),
				},
				Relationships: &datadogV2.OverrideCreateDataRelationships{
					User: &datadogV2.OnCallUserRelationship{
						Data: &datadogV2.OnCallUserRelationshipData{
							Id:   datadog.PtrString(UserDataID),
							Type: datadogV2.ONCALLUSERRELATIONSHIPTYPE_USERS.Ptr(),
						},
					},
				},
				Type: datadogV2.OVERRIDECREATEDATATYPE_OVERRIDES,
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOnCallApi(apiClient)
	resp, r, err := api.CreateOnCallScheduleOverride(ctx, ScheduleDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnCallApi.CreateOnCallScheduleOverride`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OnCallApi.CreateOnCallScheduleOverride`:\n%s\n", responseContent)
}
