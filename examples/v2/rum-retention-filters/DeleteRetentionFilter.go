// Delete a RUM retention filter returns "No Content" response

package main


import (
	"context"
	"encoding/json"
	"fmt"
	"os"

    "github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRumRetentionFiltersApi(apiClient)
	r, err := api.DeleteRetentionFilter(ctx, "a33671aa-24fd-4dcd-ba4b-5bbdbafe7690", "fe34ee09-14cf-4976-9362-08044c0dea80", )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumRetentionFiltersApi.DeleteRetentionFilter`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}