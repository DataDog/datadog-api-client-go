// Create a new RUM application returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/v2/api/v2/datadog"
)

func main() {
	body := datadog.RUMApplicationCreateRequest{
		Data: datadog.RUMApplicationCreate{
			Attributes: datadog.RUMApplicationCreateAttributes{
				Name: "my_new_rum_application",
				Type: common.PtrString("ios"),
			},
			Type: datadog.RUMAPPLICATIONCREATETYPE_RUM_APPLICATION_CREATE,
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewRUMApi(apiClient)
	resp, r, err := api.CreateRUMApplication(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMApi.CreateRUMApplication`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RUMApi.CreateRUMApplication`:\n%s\n", responseContent)
}
