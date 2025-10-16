// Add tags to a host returns "Created" response

package main


import (
	"context"
	"fmt"
	"os"

    "github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/google/uuid"
)

func main() {
	body := datadogV1.HostTags{
Host: datadog.PtrString("test.host"),
Tags: []string{
"environment:production",
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewTagsApi(apiClient)
	resp, r, err := api.CreateHostTags(ctx, "host_name", body, *datadogV1.NewCreateHostTagsOptionalParameters(), )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.CreateHostTags`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `TagsApi.CreateHostTags`:\n%s\n", responseContent)
}