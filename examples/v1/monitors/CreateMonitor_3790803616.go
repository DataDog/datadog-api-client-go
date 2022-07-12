// Create a ci-pipelines monitor returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	body := datadog.Monitor{
		Name:    common.PtrString("Example-Create_a_ci_pipelines_monitor_returns_OK_response"),
		Type:    datadog.MONITORTYPE_CI_PIPELINES_ALERT,
		Query:   `ci-pipelines("ci_level:pipeline @git.branch:staging* @ci.status:error").rollup("count").by("@git.branch,@ci.pipeline.name").last("5m") >= 1`,
		Message: common.PtrString("some message Notify: @hipchat-channel"),
		Tags: []string{
			"test:examplecreateacipipelinesmonitorreturnsokresponse",
			"env:ci",
		},
		Priority: *common.NewNullableInt64(common.PtrInt64(3)),
		Options: &datadog.MonitorOptions{
			Thresholds: &datadog.MonitorThresholds{
				Critical: common.PtrFloat64(1),
			},
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewMonitorsApi(apiClient)
	resp, r, err := api.CreateMonitor(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.CreateMonitor`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MonitorsApi.CreateMonitor`:\n%s\n", responseContent)
}
