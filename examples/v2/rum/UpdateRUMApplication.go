// Update a RUM application returns "OK" response

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
	body := datadog.RUMApplicationUpdateRequest{
		Data: datadog.RUMApplicationUpdate{
			Attributes: &datadog.RUMApplicationUpdateAttributes{
				Name: common.PtrString("updated_name_for_my_existing_rum_application"),
				Type: common.PtrString("browser|ios|android|react-native|flutter"),
			},
			Id:   "abcd1234-0000-0000-abcd-1234abcd5678",
			Type: datadog.RUMAPPLICATIONUPDATETYPE_RUM_APPLICATION_UPDATE,
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewRUMApi(apiClient)
	resp, r, err := api.UpdateRUMApplication(ctx, "id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMApi.UpdateRUMApplication`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RUMApi.UpdateRUMApplication`:\n%s\n", responseContent)
}
