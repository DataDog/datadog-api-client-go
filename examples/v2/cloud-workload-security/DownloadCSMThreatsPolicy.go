// Get the latest CSM Threats policy returns "OK" response

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.DownloadCSMThreatsPolicy", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCloudWorkloadSecurityApi(apiClient)
	resp, r, err := api.DownloadCSMThreatsPolicy(ctx)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudWorkloadSecurityApi.DownloadCSMThreatsPolicy`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := ioutil.ReadAll(resp)
	fmt.Fprintf(os.Stdout, "Response from `CloudWorkloadSecurityApi.DownloadCSMThreatsPolicy`:\n%s\n", responseContent)
}
