// Get service repository information returns "OK" response

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
	body := datadogV2.ServiceRepositoryInfoRequest{
		Data: datadogV2.ServiceRepositoryInfoRequestData{
			Attributes: datadogV2.ServiceRepositoryInfoRequestAttributes{
				DebugId: datadog.PtrUUID(uuid.MustParse("00000000-0000-4000-8000-000000000001")),
			},
			Type: datadogV2.SERVICEREPOSITORYINFODATATYPE_SERVICE_REPOSITORY_INFO,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.GetServiceRepositoryInfo", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRUMApi(apiClient)
	resp, r, err := api.GetServiceRepositoryInfo(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMApi.GetServiceRepositoryInfo`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RUMApi.GetServiceRepositoryInfo`:\n%s\n", responseContent)
}
