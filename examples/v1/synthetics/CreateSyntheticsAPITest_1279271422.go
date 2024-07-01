// Create an API test with multi subtype returns "OK - Returns the created test details." response

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
	body := datadogV1.SyntheticsAPITest{
		Config: datadogV1.SyntheticsAPITestConfig{
			ConfigVariables: []datadogV1.SyntheticsConfigVariable{
				{
					Example: datadog.PtrString("content-type"),
					Name:    "PROPERTY",
					Pattern: datadog.PtrString("content-type"),
					Type:    datadogV1.SYNTHETICSCONFIGVARIABLETYPE_TEXT,
				},
			},
			Steps: []datadogV1.SyntheticsAPIStep{
				datadogV1.SyntheticsAPIStep{
					SyntheticsAPITestStep: &datadogV1.SyntheticsAPITestStep{
						AllowFailure: datadog.PtrBool(true),
						Assertions: []datadogV1.SyntheticsAssertion{
							datadogV1.SyntheticsAssertion{
								SyntheticsAssertionTarget: &datadogV1.SyntheticsAssertionTarget{
									Operator: datadogV1.SYNTHETICSASSERTIONOPERATOR_IS,
									Type:     datadogV1.SYNTHETICSASSERTIONTYPE_STATUS_CODE,
									Target:   200,
								}},
						},
						ExtractedValues: []datadogV1.SyntheticsParsingOptions{
							{
								Field: datadog.PtrString("server"),
								Name:  datadog.PtrString("EXTRACTED_VALUE"),
								Parser: &datadogV1.SyntheticsVariableParser{
									Type: datadogV1.SYNTHETICSGLOBALVARIABLEPARSERTYPE_RAW,
								},
								Type:   datadogV1.SYNTHETICSGLOBALVARIABLEPARSETESTOPTIONSTYPE_HTTP_HEADER.Ptr(),
								Secure: datadog.PtrBool(true),
							},
						},
						IsCritical: datadog.PtrBool(true),
						Name:       "request is sent",
						Request: datadogV1.SyntheticsTestRequest{
							Method:      datadog.PtrString("GET"),
							Timeout:     datadog.PtrFloat64(10),
							Url:         datadog.PtrString("https://datadoghq.com"),
							HttpVersion: datadogV1.SYNTHETICSTESTOPTIONSHTTPVERSION_HTTP2.Ptr(),
						},
						Retry: &datadogV1.SyntheticsTestOptionsRetry{
							Count:    datadog.PtrInt64(5),
							Interval: datadog.PtrFloat64(1000),
						},
						Subtype: datadogV1.SYNTHETICSAPITESTSTEPSUBTYPE_HTTP,
					}},
				datadogV1.SyntheticsAPIStep{
					SyntheticsAPIWaitStep: &datadogV1.SyntheticsAPIWaitStep{
						Name:    "Wait",
						Subtype: datadogV1.SYNTHETICSAPIWAITSTEPSUBTYPE_WAIT,
						Value:   1,
					}},
				datadogV1.SyntheticsAPIStep{
					SyntheticsAPITestStep: &datadogV1.SyntheticsAPITestStep{
						Name:            "GRPC CALL",
						Subtype:         datadogV1.SYNTHETICSAPITESTSTEPSUBTYPE_GRPC,
						ExtractedValues: []datadogV1.SyntheticsParsingOptions{},
						AllowFailure:    datadog.PtrBool(false),
						IsCritical:      datadog.PtrBool(true),
						Retry: &datadogV1.SyntheticsTestOptionsRetry{
							Count:    datadog.PtrInt64(0),
							Interval: datadog.PtrFloat64(300),
						},
						Assertions: []datadogV1.SyntheticsAssertion{
							datadogV1.SyntheticsAssertion{
								SyntheticsAssertionTarget: &datadogV1.SyntheticsAssertionTarget{
									Operator: datadogV1.SYNTHETICSASSERTIONOPERATOR_LESS_THAN,
									Type:     datadogV1.SYNTHETICSASSERTIONTYPE_RESPONSE_TIME,
									Target:   1000,
								}},
						},
						Request: datadogV1.SyntheticsTestRequest{
							Host:                     datadog.PtrString("grpcbin.test.k6.io"),
							Port:                     datadog.PtrInt64(9000),
							Service:                  datadog.PtrString("grpcbin.GRPCBin"),
							Method:                   datadog.PtrString("Index"),
							Message:                  datadog.PtrString("{}"),
							CompressedJsonDescriptor: datadog.PtrString("eJy1lU1z2yAQhv+Lzj74I3ETH506bQ7OZOSm1w4Wa4epBARQppqM/3v5koCJJdvtxCdW77vPssCO3zMKUgHOFu/ZXvBiS6hZho/f8qe7pftYgXphWJrlA8XwxywEvNba+6PhkC2yVcVVswYp0R6ykRYlZ1SCV21SDrxsssPIeS9FJKqGfK2rqnmmSBwhWa2XlKgtaQPiDcRGCUDVfwGD2sKUqKEtc1cSoOrsMlaMOec1sySYCCgUYRSVLv2zSva2u+FQkB0pVkIw8bFuIudOOn3pOaKYVT3Iy97Pd0AYhOx5QcMsnxvRHlnuLf8ETDd3CNtrv2nejkDpRnANCmGkkFn/hsYzpBKE7jVbufgnKnV9HRM9zRPDDKPttYT61n0TdWkAAjggk9AhuxIeaXd69CYTcsGw7cBTakLVbNpRzGEgyWjkSOpMbZXkhGL6oX30R49qt3GoHrap7i0XdD41WQ+2icCNm5p1hmFqnHNlcla0riKmDZ183crDxChjbnurtxHPRE784sVhWvDfGP+SsTKibU3o5NtWHuZFGZOxP6P5VXqIOvaOSec4eYohyd7NslHuJbd1bewds85xYrNxkr2d+5IhFWF3NvaO684xjE2S5ulY+tu64Pna0fCPJgzw6vF5/WucLcYjt5xoq19O3UDptOg/OamJQRaCcPPnMTQ2QDFn+uhPvUfnCrMc99upyQY4Ui9Dlc/YoG3R/v4Cs9YE+g=="),
							Metadata:                 map[string]string{},
							CallType:                 datadogV1.SYNTHETICSTESTCALLTYPE_UNARY.Ptr(),
						},
					}},
			},
		},
		Locations: []string{
			"aws:us-east-2",
		},
		Message: "BDD test payload: synthetics_api_test_multi_step_payload.json",
		Name:    "Example-Synthetic",
		Options: datadogV1.SyntheticsTestOptions{
			AcceptSelfSigned:   datadog.PtrBool(false),
			AllowInsecure:      datadog.PtrBool(true),
			FollowRedirects:    datadog.PtrBool(true),
			MinFailureDuration: datadog.PtrInt64(10),
			MinLocationFailed:  datadog.PtrInt64(1),
			MonitorName:        datadog.PtrString("Example-Synthetic"),
			MonitorPriority:    datadog.PtrInt32(5),
			Retry: &datadogV1.SyntheticsTestOptionsRetry{
				Count:    datadog.PtrInt64(3),
				Interval: datadog.PtrFloat64(1000),
			},
			TickEvery: datadog.PtrInt64(60),
		},
		Subtype: datadogV1.SYNTHETICSTESTDETAILSSUBTYPE_MULTI.Ptr(),
		Tags: []string{
			"testing:api",
		},
		Type: datadogV1.SYNTHETICSAPITESTTYPE_API,
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewSyntheticsApi(apiClient)
	resp, r, err := api.CreateSyntheticsAPITest(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.CreateSyntheticsAPITest`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.CreateSyntheticsAPITest`:\n%s\n", responseContent)
}
