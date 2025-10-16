// Create a RUM retention filter returns "Created" response

package main


import (
	"context"
	"fmt"
	"os"

    "github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	body := datadogV2.RumRetentionFilterCreateRequest{
Data: datadogV2.RumRetentionFilterCreateData{
Type: datadogV2.RUMRETENTIONFILTERTYPE_RETENTION_FILTERS,
Attributes: datadogV2.RumRetentionFilterCreateAttributes{
Name: "Test creating retention filter",
EventType: datadogV2.RUMRETENTIONFILTEREVENTTYPE_SESSION,
Query: datadog.PtrString("custom_query"),
SampleRate: 50,
Enabled: datadog.PtrBool(true),
},
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRumRetentionFiltersApi(apiClient)
	resp, r, err := api.CreateRetentionFilter(ctx, "a33671aa-24fd-4dcd-ba4b-5bbdbafe7690", body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumRetentionFiltersApi.CreateRetentionFilter`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RumRetentionFiltersApi.CreateRetentionFilter`:\n%s\n", responseContent)
}