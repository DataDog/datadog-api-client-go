// Edit an API key returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	body := datadog.ApiKey{
		Created:   datadog.PtrString("2019-08-02 15:31:07"),
		CreatedBy: datadog.PtrString("john@example.com"),
		Key:       datadog.PtrString("1234512345123456abcabc912349abcd"),
		Name:      datadog.PtrString("example user"),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.KeyManagementApi.UpdateAPIKey(ctx, "key", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.UpdateAPIKey`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `KeyManagementApi.UpdateAPIKey`:\n%s\n", responseContent)
}
