// Edit a mobile test returns "OK" response

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
		Config: datadogV1.SyntheticsMobileTestConfig{
			Variables: []datadogV1.SyntheticsConfigVariable{
				{
					Name:   "VARIABLE_NAME",
					Secure: datadog.PtrBool(false),
					Type:   datadogV1.SYNTHETICSCONFIGVARIABLETYPE_TEXT,
				},
			},
		},
		DeviceIds: []string{
			"chrome.laptop_large",
		},
		Message: "Notification message",
		Name:    "Example test name",
		Options: datadogV1.SyntheticsMobileTestOptions{
			Bindings: []datadogV1.SyntheticsTestRestrictionPolicyBinding{
				{
					Principals: []string{},
					Relation:   datadogV1.SYNTHETICSTESTRESTRICTIONPOLICYBINDINGRELATION_EDITOR.Ptr(),
				},
			},
			Ci: &datadogV1.SyntheticsTestCiOptions{
				ExecutionRule: datadogV1.SYNTHETICSTESTEXECUTIONRULE_BLOCKING,
			},
			DeviceIds: []string{
				"synthetics:mobile:device:apple_ipad_10th_gen_2022_ios_16",
			},
			MobileApplication: datadogV1.SyntheticsMobileTestsMobileApplication{
				ApplicationId: "00000000-0000-0000-0000-aaaaaaaaaaaa",
				ReferenceId:   "00000000-0000-0000-0000-aaaaaaaaaaab",
				ReferenceType: datadogV1.SYNTHETICSMOBILETESTSMOBILEAPPLICATIONREFERENCETYPE_LATEST,
			},
			MonitorOptions: &datadogV1.SyntheticsTestOptionsMonitorOptions{
				NotificationPresetName: datadogV1.SYNTHETICSTESTOPTIONSMONITOROPTIONSNOTIFICATIONPRESETNAME_SHOW_ALL.Ptr(),
			},
			RestrictedRoles: []string{
				"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
			},
			Retry: &datadogV1.SyntheticsTestOptionsRetry{},
			Scheduling: &datadogV1.SyntheticsTestOptionsScheduling{
				Timeframes: []datadogV1.SyntheticsTestOptionsSchedulingTimeframe{
					{
						Day:  1,
						From: "07:00",
						To:   "16:00",
					},
					{
						Day:  3,
						From: "07:00",
						To:   "16:00",
					},
				},
				Timezone: "America/New_York",
			},
			TickEvery: 300,
		},
		Status: datadogV1.SYNTHETICSTESTPAUSESTATUS_LIVE.Ptr(),
		Steps: []datadogV1.SyntheticsMobileStep{
			{
				Name: "",
				Params: datadogV1.SyntheticsMobileStepParams{
					Check:     datadogV1.SYNTHETICSCHECKTYPE_EQUALS.Ptr(),
					Direction: datadogV1.SYNTHETICSMOBILESTEPPARAMSDIRECTION_UP.Ptr(),
					Element: &datadogV1.SyntheticsMobileStepParamsElement{
						ContextType:      datadogV1.SYNTHETICSMOBILESTEPPARAMSELEMENTCONTEXTTYPE_NATIVE.Ptr(),
						RelativePosition: &datadogV1.SyntheticsMobileStepParamsElementRelativePosition{},
						UserLocator: &datadogV1.SyntheticsMobileStepParamsElementUserLocator{
							Values: []datadogV1.SyntheticsMobileStepParamsElementUserLocatorValuesItems{
								{
									Type: datadogV1.SYNTHETICSMOBILESTEPPARAMSELEMENTUSERLOCATORVALUESITEMSTYPE_ACCESSIBILITY_ID.Ptr(),
								},
							},
						},
					},
					Positions: []datadogV1.SyntheticsMobileStepParamsPositionsItems{
						{},
					},
					Variable: &datadogV1.SyntheticsMobileStepParamsVariable{
						Example: "",
						Name:    "VAR_NAME",
					},
				},
				PublicId: datadog.PtrString("pub-lic-id0"),
				Type:     datadogV1.SYNTHETICSMOBILESTEPTYPE_ASSERTELEMENTCONTENT,
			},
		},
		Tags: []string{
			"env:production",
		},
		Type: datadogV1.SYNTHETICSMOBILETESTTYPE_MOBILE,
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewSyntheticsApi(apiClient)
	resp, r, err := api.UpdateMobileTest(ctx, "public_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.UpdateMobileTest`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.UpdateMobileTest`:\n%s\n", responseContent)
}
