// Create a mobile test returns "OK - Returns the created test details." response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	body := datadogV1.SyntheticsMobileTest{
		Name:   "Example-Synthetic",
		Status: datadogV1.SYNTHETICSTESTPAUSESTATUS_PAUSED.Ptr(),
		Type:   datadogV1.SYNTHETICSMOBILETESTTYPE_MOBILE,
		Config: datadogV1.SyntheticsMobileTestConfig{
			Variables: []datadogV1.SyntheticsConfigVariable{},
		},
		Message: datadog.PtrString(""),
		Options: datadogV1.SyntheticsMobileTestOptions{
			DeviceIds: []datadogV1.SyntheticsMobileDeviceID{
				datadogV1.SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_IPHONE_15_IOS_17,
			},
			MobileApplication: &datadogV1.SyntheticsMobileTestsMobileApplication{
				ApplicationId: datadog.PtrString("ab0e0aed-536d-411a-9a99-5428c27d8f8e"),
				ReferenceId:   datadog.PtrString("6115922a-5f5d-455e-bc7e-7955a57f3815"),
				ReferenceType: datadogV1.SYNTHETICSMOBILETESTSMOBILEAPPLICATIONREFERENCETYPE_VERSION.Ptr(),
			},
			TickEvery: datadog.PtrInt64(3600),
		},
		Steps: []datadogV1.SyntheticsMobileStep{},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewSyntheticsApi(apiClient)
	resp, r, err := api.CreateSyntheticsMobileTest(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.CreateSyntheticsMobileTest`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.CreateSyntheticsMobileTest`:\n%s\n", responseContent)
}
