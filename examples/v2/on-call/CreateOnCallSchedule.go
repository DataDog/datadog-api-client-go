// Create On-Call schedule returns "Created" response

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

	// there is a valid "dd_team" in the system
	DdTeamDataID := os.Getenv("DD_TEAM_DATA_ID")

	body := datadogV2.ScheduleCreateRequest{
		Data: datadogV2.ScheduleCreateRequestData{
			Attributes: datadogV2.ScheduleCreateRequestDataAttributes{
				Layers: []datadogV2.ScheduleCreateRequestDataAttributesLayersItems{
					{
						EffectiveDate: time.Now().AddDate(0, 0, -10),
						EndDate:       datadog.PtrTime(time.Now().AddDate(0, 0, 10)),
						Interval: datadogV2.LayerAttributesInterval{
							Days: datadog.PtrInt32(1),
						},
						Members: []datadogV2.ScheduleRequestDataAttributesLayersItemsMembersItems{
							{
								User: &datadogV2.ScheduleRequestDataAttributesLayersItemsMembersItemsUser{
									Id: datadog.PtrString(UserDataID),
								},
							},
						},
						Name: "Layer 1",
						Restrictions: []datadogV2.TimeRestriction{
							{
								EndDay:    datadogV2.WEEKDAY_FRIDAY.Ptr(),
								EndTime:   datadog.PtrString("17:00:00"),
								StartDay:  datadogV2.WEEKDAY_MONDAY.Ptr(),
								StartTime: datadog.PtrString("09:00:00"),
							},
						},
						RotationStart: time.Now().AddDate(0, 0, -5),
					},
				},
				Name:     "Example-On-Call",
				TimeZone: "America/New_York",
			},
			Relationships: &datadogV2.ScheduleCreateRequestDataRelationships{
				Teams: &datadogV2.DataRelationshipsTeams{
					Data: []datadogV2.DataRelationshipsTeamsDataItems{
						{
							Id:   DdTeamDataID,
							Type: datadogV2.DATARELATIONSHIPSTEAMSDATAITEMSTYPE_TEAMS,
						},
					},
				},
			},
			Type: datadogV2.SCHEDULECREATEREQUESTDATATYPE_SCHEDULES,
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
