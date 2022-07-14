// Schedule a downtime once a year

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	body := datadog.Downtime{
		Message: datadog.PtrString("Example-Schedule_a_downtime_once_a_year"),
		Recurrence: *datadog.NewNullableDowntimeRecurrence(&datadog.DowntimeRecurrence{
			Period: datadog.PtrInt32(1),
			Type:   datadog.PtrString("years"),
		}),
		Scope: []string{
			"*",
		},
		Start:                         datadog.PtrInt64(time.Now().Unix()),
		End:                           *datadog.NewNullableInt64(datadog.PtrInt64(time.Now().Add(time.Hour * 1).Unix())),
		Timezone:                      datadog.PtrString("Etc/UTC"),
		MuteFirstRecoveryNotification: datadog.PtrBool(true),
		MonitorTags: []string{
			"tag0",
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.DowntimesApi.CreateDowntime(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DowntimesApi.CreateDowntime`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DowntimesApi.CreateDowntime`:\n%s\n", responseContent)
}
