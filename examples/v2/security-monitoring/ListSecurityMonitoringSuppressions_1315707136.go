// Get all suppression rules returns "OK" response with sort ascending

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	// there is a valid "suppression" in the system

	// there is a valid "suppression2" in the system

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.ListSecurityMonitoringSuppressions(ctx, *datadogV2.NewListSecurityMonitoringSuppressionsOptionalParameters().WithSort(datadogV2.SECURITYMONITORINGSUPPRESSIONSORT_NAME).WithQuery("id:3dd-0uc-h1s OR id:886e6c3e-e543-049c-ee1b-56a1110295c0"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.ListSecurityMonitoringSuppressions`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.ListSecurityMonitoringSuppressions`:\n%s\n", responseContent)
}
