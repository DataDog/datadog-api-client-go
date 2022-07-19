// Schedule a downtime with until occurrences

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
		Message: common.PtrString("Example-Schedule_a_downtime_with_until_occurrences"),
		Recurrence: *datadog.NewNullableDowntimeRecurrence(&datadog.DowntimeRecurrence{
			Period:           common.PtrInt32(1),
			Type:             common.PtrString("weeks"),
			UntilOccurrences: *common.NewNullableInt32(common.PtrInt32(3)),
			WeekDays: []string{
				"Mon",
				"Tue",
				"Wed",
				"Thu",
				"Fri",
			},
		}),
		Scope: []string{
			"*",
		},
		Start:    common.PtrInt64(time.Now().Unix()),
		End:      *common.NewNullableInt64(common.PtrInt64(time.Now().Add(time.Hour * 1).Unix())),
		Timezone: common.PtrString("Etc/UTC"),
		MonitorTags: []string{
			"tag0",
		},
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
