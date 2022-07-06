// Take graph snapshots returns "OK" response

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
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.SnapshotsApi(apiClient)
	resp, r, err := api.GetGraphSnapshot(ctx, time.Now().AddDate(0, 0, -1).Unix(), time.Now().Unix(), *datadog.NewGetGraphSnapshotOptionalParameters().WithMetricQuery("avg:system.load.1{*}").WithTitle("System load").WithHeight(400).WithWidth(600))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsApi.GetGraphSnapshot`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsApi.GetGraphSnapshot`:\n%s\n", responseContent)
}
