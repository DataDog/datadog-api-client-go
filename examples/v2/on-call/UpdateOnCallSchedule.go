// Update on-call schedule returns "OK" response

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
	ScheduleDataRelationshipsLayersData0ID := os.Getenv("SCHEDULE_DATA_RELATIONSHIPS_LAYERS_DATA_0_ID")

	// there is a valid "user" in the system
	UserDataID := os.Getenv("USER_DATA_ID")

	// there is a valid "dd_team" in the system
	DdTeamDataID := os.Getenv("DD_TEAM_DATA_ID")

	body := datadogV2.ScheduleUpdateRequest{
		Data: datadogV2.ScheduleUpdateRequestData{
			Id: ScheduleDataID,
			Attributes: datadogV2.ScheduleUpdateRequestDataAttributes{
				Layers: []datadogV2.ScheduleUpdateRequestDataAttributesLayersItems{
					{
						Id:            datadog.PtrString(ScheduleDataRelationshipsLayersData0ID),
						EffectiveDate: datadog.PtrTime(time.Now().AddDate(0, 0, -10)),
						EndDate:       datadog.PtrTime(time.Now().AddDate(0, 0, 10)),
						Interval: &datadogV2.ScheduleUpdateRequestDataAttributesLayersItemsInterval{
							Seconds: datadog.PtrInt64(300),
						},
						Members: []datadogV2.ScheduleUpdateRequestDataAttributesLayersItemsMembersItems{
							{
								User: &datadogV2.ScheduleUpdateRequestDataAttributesLayersItemsMembersItemsUser{
									Id: datadog.PtrString(UserDataID),
								},
							},
						},
						Name: datadog.PtrString("Layer 1"),
						Restrictions: []datadogV2.ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItems{
							{
								EndDay:    datadogV2.SCHEDULEUPDATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSENDDAY_FRIDAY.Ptr(),
								EndTime:   datadog.PtrString("17:00:00"),
								StartDay:  datadogV2.SCHEDULEUPDATEREQUESTDATAATTRIBUTESLAYERSITEMSRESTRICTIONSITEMSSTARTDAY_MONDAY.Ptr(),
								StartTime: datadog.PtrString("09:00:00"),
							},
						},
						RotationStart: datadog.PtrTime(time.Now().AddDate(0, 0, -5)),
					},
				},
				Name: "Example-On-Call",
				Tags: []string{
					"tag1",
					"tag2",
					"tag3",
				},
				TimeZone: "America/New_York",
			},
			Relationships: &datadogV2.ScheduleUpdateRequestDataRelationships{
				Teams: &datadogV2.ScheduleUpdateRequestDataRelationshipsTeams{
					Data: []datadogV2.ScheduleUpdateRequestDataRelationshipsTeamsDataItems{
						{
							Id:   datadog.PtrString(DdTeamDataID),
							Type: datadogV2.SCHEDULEUPDATEREQUESTDATARELATIONSHIPSTEAMSDATAITEMSTYPE_TEAMS.Ptr(),
						},
					},
				},
			},
			Type: datadogV2.SCHEDULEUPDATEREQUESTDATATYPE_SCHEDULES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOnCallApi(apiClient)
	resp, r, err := api.UpdateOnCallSchedule(ctx, ScheduleDataID, body, *datadogV2.NewUpdateOnCallScheduleOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnCallApi.UpdateOnCallSchedule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OnCallApi.UpdateOnCallSchedule`:\n%s\n", responseContent)
}
