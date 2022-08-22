// Take graph snapshots returns "OK" response

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
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewSnapshotsApi(apiClient)
	resp, r, err := api.GetGraphSnapshot(ctx, time.Now().AddDate(0, 0, -1).Unix(), time.Now().Unix(), *datadogV1.NewGetGraphSnapshotOptionalParameters().WithMetricQuery("avg:system.load.1{*}").WithTitle("System load").WithHeight(400).WithWidth(600))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsApi.GetGraphSnapshot`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsApi.GetGraphSnapshot`:\n%s\n", responseContent)
}
