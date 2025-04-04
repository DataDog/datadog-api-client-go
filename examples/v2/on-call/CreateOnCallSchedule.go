// Create on-call schedule returns "Created" response

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
	// there is a valid "user" in the system
	UserDataID := os.Getenv("USER_DATA_ID")

	// there is a valid "team" in the system
	TeamDataID := os.Getenv("TEAM_DATA_ID")

	body := datadogV2.ScheduleCreateRequest{
		Data: &datadogV2.ScheduleCreateRequestData{
			Attributes: &datadogV2.ScheduleCreateRequestDataAttributes{
				Layers: []datadogV2.ScheduleCreateRequestDataAttributesLayersItems{
					{
						EffectiveDate: time.Date(2025, 2, 3, 5, 0, 0, 0, time.UTC),
						EndDate:       datadog.PtrTime(time.Date(2025, 12, 31, 0, 0, 0, 0, time.UTC)),
						Interval: datadogV2.ScheduleCreateRequestDataAttributesLayersItemsInterval{
							Days: datadog.PtrInt32(1),
						},
						Members: []datadogV2.ScheduleCreateRequestDataAttributesLayersItemsMembersItems{
							{
								User: &datadogV2.ScheduleCreateRequestDataAttributesLayersItemsMembersItemsUser{
									Id: datadog.PtrString(UserDataID),
								},
							},
						},
						Name: "Layer 1",
						Restrictions: []datadogV2.ScheduleCreateRequestDataAttributesLayersItemsRestrictionsItems{
							{
								EndDay:    datadogV2.SCHEDULECREATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSENDDAY_FRIDAY.Ptr(),
								EndTime:   datadog.PtrString("17:00:00"),
								StartDay:  datadogV2.SCHEDULECREATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSSTARTDAY_MONDAY.Ptr(),
								StartTime: datadog.PtrString("09:00:00"),
							},
						},
						RotationStart: time.Date(2025, 2, 1, 0, 0, 0, 0, time.UTC),
					},
				},
				Name: "Example-On-Call",
				Tags: []string{
					"tag1",
					"tag2",
				},
				TimeZone: "America/New_York",
			},
			Relationships: &datadogV2.ScheduleCreateRequestDataRelationships{
				Teams: &datadogV2.ScheduleCreateRequestDataRelationshipsTeams{
					Data: []datadogV2.ScheduleCreateRequestDataRelationshipsTeamsDataItems{
						{
							Id:   datadog.PtrString(TeamDataID),
							Type: datadogV2.SCHEDULECREATEREQUESTDATARELATIONSHIPSTEAMSDATAITEMSTYPE_TEAMS.Ptr(),
						},
					},
				},
			},
			Type: datadogV2.SCHEDULECREATEREQUESTDATATYPE_SCHEDULES.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOnCallApi(apiClient)
	resp, r, err := api.CreateOnCallSchedule(ctx, body, *datadogV2.NewCreateOnCallScheduleOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnCallApi.CreateOnCallSchedule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OnCallApi.CreateOnCallSchedule`:\n%s\n", responseContent)
}
