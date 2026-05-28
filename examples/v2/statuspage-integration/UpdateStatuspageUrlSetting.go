// Update a Statuspage URL setting returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.StatuspageUrlSettingUpdateRequest{
		Data: datadogV2.StatuspageUrlSettingUpdateData{
			Attributes: datadogV2.StatuspageUrlSettingUpdateAttributes{
				CustomTags: datadog.PtrString("team:collaboration-integrations,env:prod"),
				Url:        datadog.PtrString("https://example.statuspage.io"),
			},
			Id:   "596da4af-0563-4097-90ff-07230c3f9db3",
			Type: datadogV2.STATUSPAGEURLSETTINGTYPE_STATUSPAGE_URL_SETTING,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewStatuspageIntegrationApi(apiClient)
	resp, r, err := api.UpdateStatuspageUrlSetting(ctx, "statuspage_url_setting_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatuspageIntegrationApi.UpdateStatuspageUrlSetting`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `StatuspageIntegrationApi.UpdateStatuspageUrlSetting`:\n%s\n", responseContent)
}
