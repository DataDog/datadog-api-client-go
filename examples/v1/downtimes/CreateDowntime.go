// Schedule a downtime returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	body := datadog.Downtime{
		Message:  common.PtrString("Example-Schedule_a_downtime_returns_OK_response"),
		Start:    common.PtrInt64(time.Now().Unix()),
		Timezone: common.PtrString("Etc/UTC"),
		Scope: []string{
			"test:examplescheduleadowntimereturnsokresponse",
		},
		Recurrence: *datadog.NewNullableDowntimeRecurrence(&datadog.DowntimeRecurrence{
			Type:   common.PtrString("weeks"),
			Period: common.PtrInt32(1),
			WeekDays: []string{
				"Mon",
				"Tue",
				"Wed",
				"Thu",
				"Fri",
			},
			UntilDate: *common.NewNullableInt64(common.PtrInt64(time.Now().AddDate(0, 0, 21).Unix())),
		}),
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewDowntimesApi(apiClient)
	resp, r, err := api.CreateDowntime(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DowntimesApi.CreateDowntime`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DowntimesApi.CreateDowntime`:\n%s\n", responseContent)
}
