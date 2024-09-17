// Schedule a downtime with until occurrences

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	body := datadogV1.Downtime{
		Message: *datadog.NewNullableString(datadog.PtrString("Example-Downtime")),
		Recurrence: *datadogV1.NewNullableDowntimeRecurrence(&datadogV1.DowntimeRecurrence{
			Period:           datadog.PtrInt32(1),
			Type:             datadog.PtrString("weeks"),
			UntilOccurrences: *datadog.NewNullableInt32(datadog.PtrInt32(3)),
			WeekDays: *datadog.NewNullableList(&[]string{
				"Mon",
				"Tue",
				"Wed",
				"Thu",
				"Fri",
			}),
		}),
		Scope: []string{
			"*",
		},
		Start:    datadog.PtrInt64(time.Now().Unix()),
		End:      *datadog.NewNullableInt64(datadog.PtrInt64(time.Now().Add(time.Hour * 1).Unix())),
		Timezone: datadog.PtrString("Etc/UTC"),
		MonitorTags: []string{
			"tag0",
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewDowntimesApi(apiClient)
	resp, r, err := api.CreateDowntime(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DowntimesApi.CreateDowntime`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DowntimesApi.CreateDowntime`:\n%s\n", responseContent)
}
