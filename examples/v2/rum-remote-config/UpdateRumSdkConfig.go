// Update a RUM SDK configuration returns "OK" response

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
	body := datadogV2.RumSdkConfigUpdateRequest{
		Data: datadogV2.RumSdkConfigUpdateData{
			Attributes: datadogV2.RumSdkConfigUpdateAttributes{
				Rum: datadogV2.RumSdkConfigRumUpdateAttributes{
					AllowedTracingUrls: []datadogV2.RumSdkConfigTracingUrlConfig{
						{
							Match: datadogV2.RumSdkConfigMatchOption{
								RcSerializedType: datadogV2.RUMSDKCONFIGMATCHOPTIONSERIALIZEDTYPE_STRING,
								Value:            "https://app.datadoghq.com",
							},
							PropagatorTypes: []datadogV2.RumSdkConfigTracingUrlPropagatorType{
								datadogV2.RUMSDKCONFIGTRACINGURLPROPAGATORTYPE_DATADOG,
								datadogV2.RUMSDKCONFIGTRACINGURLPROPAGATORTYPE_TRACECONTEXT,
							},
						},
					},
					AllowedTrackingOrigins: []datadogV2.RumSdkConfigMatchOption{
						{
							RcSerializedType: datadogV2.RUMSDKCONFIGMATCHOPTIONSERIALIZEDTYPE_STRING,
							Value:            "https://app.datadoghq.com",
						},
					},
					Context: []datadogV2.RumSdkConfigDynamicOptionPair{
						{
							Key: "id",
							Value: datadogV2.RumSdkConfigDynamicOption{
								Attribute: datadog.PtrString("data-version"),
								Extractor: &datadogV2.RumSdkConfigSerializedRegex{
									RcSerializedType: datadogV2.RUMSDKCONFIGSERIALIZEDREGEXTYPE_REGEX,
									Value:            "^https://app-.*.datadoghq.com",
								},
								Key:              datadog.PtrString("app.version"),
								Name:             datadog.PtrString("app_version"),
								Path:             datadog.PtrString("application.version"),
								RcSerializedType: datadogV2.RUMSDKCONFIGDYNAMICOPTIONSERIALIZEDTYPE_DYNAMIC,
								Selector:         datadog.PtrString("#app-version"),
								Strategy:         datadogV2.RUMSDKCONFIGDYNAMICOPTIONSTRATEGY_JS,
							},
						},
					},
					DefaultPrivacyLevel:          "mask",
					EnablePrivacyForActionName:   true,
					Env:                          datadog.PtrString("production"),
					Service:                      datadog.PtrString("my-service"),
					SessionReplaySampleRate:      20,
					SessionSampleRate:            75,
					TraceSampleRate:              datadog.PtrInt64(100),
					TrackSessionAcrossSubdomains: datadog.PtrBool(false),
					User: []datadogV2.RumSdkConfigDynamicOptionPair{
						{
							Key: "id",
							Value: datadogV2.RumSdkConfigDynamicOption{
								Attribute: datadog.PtrString("data-version"),
								Extractor: &datadogV2.RumSdkConfigSerializedRegex{
									RcSerializedType: datadogV2.RUMSDKCONFIGSERIALIZEDREGEXTYPE_REGEX,
									Value:            "^https://app-.*.datadoghq.com",
								},
								Key:              datadog.PtrString("app.version"),
								Name:             datadog.PtrString("app_version"),
								Path:             datadog.PtrString("application.version"),
								RcSerializedType: datadogV2.RUMSDKCONFIGDYNAMICOPTIONSERIALIZEDTYPE_DYNAMIC,
								Selector:         datadog.PtrString("#app-version"),
								Strategy:         datadogV2.RUMSDKCONFIGDYNAMICOPTIONSTRATEGY_JS,
							},
						},
					},
					Version: &datadogV2.RumSdkConfigDynamicOption{
						Attribute: datadog.PtrString("data-version"),
						Extractor: &datadogV2.RumSdkConfigSerializedRegex{
							RcSerializedType: datadogV2.RUMSDKCONFIGSERIALIZEDREGEXTYPE_REGEX,
							Value:            "^https://app-.*.datadoghq.com",
						},
						Key:              datadog.PtrString("app.version"),
						Name:             datadog.PtrString("app_version"),
						Path:             datadog.PtrString("application.version"),
						RcSerializedType: datadogV2.RUMSDKCONFIGDYNAMICOPTIONSERIALIZEDTYPE_DYNAMIC,
						Selector:         datadog.PtrString("#app-version"),
						Strategy:         datadogV2.RUMSDKCONFIGDYNAMICOPTIONSTRATEGY_JS,
					},
				},
			},
			Id:   "abc12345-1234-5678-abcd-ef1234567890",
			Type: datadogV2.RUMSDKCONFIGTYPE_RUM_SDK_CONFIG,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateRumSdkConfig", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRUMRemoteConfigApi(apiClient)
	resp, r, err := api.UpdateRumSdkConfig(ctx, "config_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMRemoteConfigApi.UpdateRumSdkConfig`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RUMRemoteConfigApi.UpdateRumSdkConfig`:\n%s\n", responseContent)
}
