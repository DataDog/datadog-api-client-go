// Unregister an App Key returns "No Content" response

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
	api := datadogV2.NewActionConnectionApi(apiClient)
	r, err := api.UnregisterAppKey(ctx, "57cc69ae-9214-4ecc-8df8-43ecc1d92d99", )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionConnectionApi.UnregisterAppKey`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}