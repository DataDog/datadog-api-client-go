// Create Vercel access token returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.VercelTokenCreateRequest{
		AuthGrantCode:         "code",
		VercelConfigurationId: "icfg_abc123",
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewVercelApi(apiClient)
	r, err := api.CreateVercelToken(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VercelApi.CreateVercelToken`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
