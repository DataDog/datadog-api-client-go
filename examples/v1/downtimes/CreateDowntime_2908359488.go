// Schedule a downtime until date

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
		Message: datadog.PtrString("Example-Schedule_a_downtime_until_date"),
		Recurrence: *common.datadog.NewNullableDowntimeRecurrence(&datadog.DowntimeRecurrence{
			Period:    datadog.PtrInt32(1),
			Type:      datadog.PtrString("weeks"),
			UntilDate: *common.NewNullableInt64(datadog.PtrInt64(time.Now().AddDate(0, 0, 21).Unix())),
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
		Start:                         datadog.PtrInt64(time.Now().Unix()),
		End:                           *common.NewNullableInt64(datadog.PtrInt64(time.Now().Add(time.Hour * 1).Unix())),
		Timezone:                      datadog.PtrString("Etc/UTC"),
		MuteFirstRecoveryNotification: datadog.PtrBool(true),
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.DowntimesApi(apiClient)
	resp, r, err := api.CreateDowntime(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DowntimesApi.CreateDowntime`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DowntimesApi.CreateDowntime`:\n%s\n", responseContent)
}
