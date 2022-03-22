// Update host tags returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	body := datadog.HostTags{
		Host: datadog.PtrString("test.host"),
		Tags: &[]string{
			"environment:production",
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.TagsApi.UpdateHostTags(ctx, "host_name", body, *datadog.NewUpdateHostTagsOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.UpdateHostTags`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `TagsApi.UpdateHostTags`:\n%s\n", responseContent)
}
