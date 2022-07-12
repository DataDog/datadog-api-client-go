// Schedule a downtime once a year

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
		Message: datadog.PtrString("Example-Schedule_a_downtime_once_a_year"),
		Recurrence: *common.datadog.NewNullableDowntimeRecurrence(&datadog.DowntimeRecurrence{
			Period: datadog.PtrInt32(1),
			Type:   datadog.PtrString("years"),
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
	api := datadog.NewDowntimesApi(apiClient)
	resp, r, err := api.CreateDowntime(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DowntimesApi.CreateDowntime`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DowntimesApi.CreateDowntime`:\n%s\n", responseContent)
}
