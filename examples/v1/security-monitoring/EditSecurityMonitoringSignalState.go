// Change the triage state of a security signal returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	body := datadogV1.SignalStateUpdateRequest{
		ArchiveReason: datadogV1.SIGNALARCHIVEREASON_NONE.Ptr(),
		State:         datadogV1.SIGNALTRIAGESTATE_OPEN,
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.EditSecurityMonitoringSignalState(ctx, "AQAAAYDiB_Ol8PbzFAAAAABBWURpQl9PbEFBQU0yeXhGTG9ZV2JnQUE", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.EditSecurityMonitoringSignalState`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.EditSecurityMonitoringSignalState`:\n%s\n", responseContent)
}
