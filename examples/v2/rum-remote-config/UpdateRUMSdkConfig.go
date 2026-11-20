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
	body := datadogV2.RUMSdkConfigUpdateRequest{
		Data: datadogV2.RUMSdkConfigUpdateData{
			Attributes: datadogV2.RUMSdkConfigUpdateAttributes{
				Rum: datadogV2.RUMSdkConfigRUMUpdateAttributes{
					AllowedTracingUrls: []datadogV2.RUMSdkConfigTracingUrlConfig{
						{
							Match: datadogV2.RUMSdkConfigMatchOption{
								RcSerializedType: datadogV2.RUMSDKCONFIGMATCHOPTIONSERIALIZEDTYPE_STRING,
								Value:            "https://app.datadoghq.com",
							},
							PropagatorTypes: []datadogV2.RUMSdkConfigTracingUrlPropagatorType{
								datadogV2.RUMSDKCONFIGTRACINGURLPROPAGATORTYPE_DATADOG,
								datadogV2.RUMSDKCONFIGTRACINGURLPROPAGATORTYPE_TRACECONTEXT,
							},
						},
					},
					AllowedTrackingOrigins: []datadogV2.RUMSdkConfigMatchOption{
						{
							RcSerializedType: datadogV2.RUMSDKCONFIGMATCHOPTIONSERIALIZEDTYPE_STRING,
							Value:            "https://app.datadoghq.com",
						},
					},
					Context: []datadogV2.RUMSdkConfigDynamicOptionPair{
						{
							Key: "id",
							Value: datadogV2.RUMSdkConfigDynamicOption{
								Attribute: datadog.PtrString("data-version"),
								Extractor: &datadogV2.RUMSdkConfigSerializedRegex{
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
					User: []datadogV2.RUMSdkConfigDynamicOptionPair{
						{
							Key: "id",
							Value: datadogV2.RUMSdkConfigDynamicOption{
								Attribute: datadog.PtrString("data-version"),
								Extractor: &datadogV2.RUMSdkConfigSerializedRegex{
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
					Version: &datadogV2.RUMSdkConfigDynamicOption{
						Attribute: datadog.PtrString("data-version"),
						Extractor: &datadogV2.RUMSdkConfigSerializedRegex{
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
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateRUMSdkConfig", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRUMRemoteConfigApi(apiClient)
	resp, r, err := api.UpdateRUMSdkConfig(ctx, "config_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMRemoteConfigApi.UpdateRUMSdkConfig`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RUMRemoteConfigApi.UpdateRUMSdkConfig`:\n%s\n", responseContent)
}
