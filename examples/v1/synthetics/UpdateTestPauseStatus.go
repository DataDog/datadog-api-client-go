// Pause or start a test returns "OK - Returns a boolean indicating if the update was successful." response

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
	body := datadog.SyntheticsUpdateTestPauseStatusPayload{
		NewStatus: datadog.SYNTHETICSTESTPAUSESTATUS_LIVE.Ptr(),
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.SyntheticsApi(apiClient)
	resp, r, err := api.UpdateTestPauseStatus(ctx, "public_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.UpdateTestPauseStatus`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.UpdateTestPauseStatus`:\n%s\n", responseContent)
}
