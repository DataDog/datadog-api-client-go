// Get bandit returns "OK" response

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
	// there is a valid "bandit" in the system
	BanditDataID := uuid.MustParse(os.Getenv("BANDIT_DATA_ID"))

	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateBandit", true)
	configuration.SetUnstableOperationEnabled("v2.GetBandit", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewExperimentsApi(apiClient)
	resp, r, err := api.GetBandit(ctx, BanditDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsApi.GetBandit`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsApi.GetBandit`:\n%s\n", responseContent)
}
