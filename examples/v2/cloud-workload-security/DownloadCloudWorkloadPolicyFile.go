// Get the latest Cloud Workload Security policy returns "OK" response

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudWorkloadSecurityApi.DownloadCloudWorkloadPolicyFile(ctx)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudWorkloadSecurityApi.DownloadCloudWorkloadPolicyFile`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := ioutil.ReadAll(resp)
	fmt.Fprintf(os.Stdout, "Response from `CloudWorkloadSecurityApi.DownloadCloudWorkloadPolicyFile`:\n%s\n", responseContent)
}
