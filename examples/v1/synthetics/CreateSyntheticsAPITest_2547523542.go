// Create an API test with MCP steps returns "OK - Returns the created test details." response

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
			Steps: []datadogV1.SyntheticsAPIStep{
				datadogV1.SyntheticsAPIStep{
					SyntheticsAPITestStep: &datadogV1.SyntheticsAPITestStep{
						Name:         "Initialize MCP session",
						Subtype:      datadogV1.SYNTHETICSAPITESTSTEPSUBTYPE_MCP,
						AllowFailure: datadog.PtrBool(false),
						IsCritical:   datadog.PtrBool(true),
						Retry: &datadogV1.SyntheticsTestOptionsRetry{
							Count:    datadog.PtrInt64(0),
							Interval: datadog.PtrFloat64(300),
						},
						Assertions: []datadogV1.SyntheticsAssertion{
							datadogV1.SyntheticsAssertion{
								SyntheticsAssertionTarget: &datadogV1.SyntheticsAssertionTarget{
									Operator: datadogV1.SYNTHETICSASSERTIONOPERATOR_IS,
									Type:     datadogV1.SYNTHETICSASSERTIONTYPE_STATUS_CODE,
									Target: datadogV1.SyntheticsAssertionTargetValue{
										SyntheticsAssertionTargetValueNumber: datadog.PtrFloat64(200)},
								}},
							datadogV1.SyntheticsAssertion{
								SyntheticsAssertionMCPRespectsSpecification: &datadogV1.SyntheticsAssertionMCPRespectsSpecification{
									Type: datadogV1.SYNTHETICSASSERTIONMCPRESPECTSSPECIFICATIONTYPE_MCP_RESPECTS_SPECIFICATION,
								}},
							datadogV1.SyntheticsAssertion{
								SyntheticsAssertionMCPServerCapabilitiesTarget: &datadogV1.SyntheticsAssertionMCPServerCapabilitiesTarget{
									Operator: datadogV1.SYNTHETICSASSERTIONOPERATOR_CONTAINS,
									Type:     datadogV1.SYNTHETICSASSERTIONMCPSERVERCAPABILITIESTYPE_MCP_SERVER_CAPABILITIES,
									Target: []datadogV1.SyntheticsMCPServerCapability{
										datadogV1.SYNTHETICSMCPSERVERCAPABILITY_TOOLS,
									},
								}},
						},
						Request: datadogV1.SyntheticsTestRequest{
							Url:                datadog.PtrString("https://example.org/mcp"),
							CallType:           datadogV1.SYNTHETICSTESTCALLTYPE_INIT.Ptr(),
							McpProtocolVersion: datadogV1.SYNTHETICSMCPPROTOCOLVERSION_VERSION_2025_06_18.Ptr(),
							Headers: map[string]string{
								"DD-API-KEY":         "<YOUR-API-KEY>",
								"DD-APPLICATION-KEY": "<YOUR-APP-KEY>",
							},
						},
					}},
				datadogV1.SyntheticsAPIStep{
					SyntheticsAPITestStep: &datadogV1.SyntheticsAPITestStep{
						Name:         "List MCP tools",
						Subtype:      datadogV1.SYNTHETICSAPITESTSTEPSUBTYPE_MCP,
						AllowFailure: datadog.PtrBool(false),
						IsCritical:   datadog.PtrBool(true),
						Retry: &datadogV1.SyntheticsTestOptionsRetry{
							Count:    datadog.PtrInt64(0),
							Interval: datadog.PtrFloat64(300),
						},
						Assertions: []datadogV1.SyntheticsAssertion{
							datadogV1.SyntheticsAssertion{
								SyntheticsAssertionTarget: &datadogV1.SyntheticsAssertionTarget{
									Operator: datadogV1.SYNTHETICSASSERTIONOPERATOR_IS,
									Type:     datadogV1.SYNTHETICSASSERTIONTYPE_STATUS_CODE,
									Target: datadogV1.SyntheticsAssertionTargetValue{
										SyntheticsAssertionTargetValueNumber: datadog.PtrFloat64(200)},
								}},
							datadogV1.SyntheticsAssertion{
								SyntheticsAssertionTarget: &datadogV1.SyntheticsAssertionTarget{
									Operator: datadogV1.SYNTHETICSASSERTIONOPERATOR_MORE_THAN,
									Type:     datadogV1.SYNTHETICSASSERTIONTYPE_MCP_TOOL_COUNT,
									Target: datadogV1.SyntheticsAssertionTargetValue{
										SyntheticsAssertionTargetValueNumber: datadog.PtrFloat64(0)},
								}},
							datadogV1.SyntheticsAssertion{
								SyntheticsAssertionTarget: &datadogV1.SyntheticsAssertionTarget{
									Operator: datadogV1.SYNTHETICSASSERTIONOPERATOR_LESS_THAN,
									Type:     datadogV1.SYNTHETICSASSERTIONTYPE_MCP_TOOL_NAME_LENGTH,
									Target: datadogV1.SyntheticsAssertionTargetValue{
										SyntheticsAssertionTargetValueNumber: datadog.PtrFloat64(64)},
								}},
							datadogV1.SyntheticsAssertion{
								SyntheticsAssertionMCPRespectsSpecification: &datadogV1.SyntheticsAssertionMCPRespectsSpecification{
									Type: datadogV1.SYNTHETICSASSERTIONMCPRESPECTSSPECIFICATIONTYPE_MCP_RESPECTS_SPECIFICATION,
								}},
						},
						Request: datadogV1.SyntheticsTestRequest{
							Url:                datadog.PtrString("https://example.org/mcp"),
							CallType:           datadogV1.SYNTHETICSTESTCALLTYPE_TOOL_LIST.Ptr(),
							McpProtocolVersion: datadogV1.SYNTHETICSMCPPROTOCOLVERSION_VERSION_2025_06_18.Ptr(),
							Headers: map[string]string{
								"DD-API-KEY":         "<YOUR-API-KEY>",
								"DD-APPLICATION-KEY": "<YOUR-APP-KEY>",
							},
						},
					}},
				datadogV1.SyntheticsAPIStep{
					SyntheticsAPITestStep: &datadogV1.SyntheticsAPITestStep{
						Name:         "Call MCP search tool",
						Subtype:      datadogV1.SYNTHETICSAPITESTSTEPSUBTYPE_MCP,
						AllowFailure: datadog.PtrBool(false),
						IsCritical:   datadog.PtrBool(true),
						Retry: &datadogV1.SyntheticsTestOptionsRetry{
							Count:    datadog.PtrInt64(0),
							Interval: datadog.PtrFloat64(300),
						},
						Assertions: []datadogV1.SyntheticsAssertion{
							datadogV1.SyntheticsAssertion{
								SyntheticsAssertionTarget: &datadogV1.SyntheticsAssertionTarget{
									Operator: datadogV1.SYNTHETICSASSERTIONOPERATOR_IS,
									Type:     datadogV1.SYNTHETICSASSERTIONTYPE_STATUS_CODE,
									Target: datadogV1.SyntheticsAssertionTargetValue{
										SyntheticsAssertionTargetValueNumber: datadog.PtrFloat64(200)},
								}},
							datadogV1.SyntheticsAssertion{
								SyntheticsAssertionTarget: &datadogV1.SyntheticsAssertionTarget{
									Operator: datadogV1.SYNTHETICSASSERTIONOPERATOR_LESS_THAN,
									Type:     datadogV1.SYNTHETICSASSERTIONTYPE_RESPONSE_TIME,
									Target: datadogV1.SyntheticsAssertionTargetValue{
										SyntheticsAssertionTargetValueNumber: datadog.PtrFloat64(5000)},
								}},
							datadogV1.SyntheticsAssertion{
								SyntheticsAssertionMCPRespectsSpecification: &datadogV1.SyntheticsAssertionMCPRespectsSpecification{
									Type: datadogV1.SYNTHETICSASSERTIONMCPRESPECTSSPECIFICATIONTYPE_MCP_RESPECTS_SPECIFICATION,
								}},
						},
						Request: datadogV1.SyntheticsTestRequest{
							Url:                datadog.PtrString("https://example.org/mcp"),
							CallType:           datadogV1.SYNTHETICSTESTCALLTYPE_TOOL_CALL.Ptr(),
							McpProtocolVersion: datadogV1.SYNTHETICSMCPPROTOCOLVERSION_VERSION_2025_06_18.Ptr(),
							ToolName:           datadog.PtrString("search"),
							ToolArgs: map[string]interface{}{
								"limit": 5,
								"query": "datadog synthetics",
							},
							Headers: map[string]string{
								"DD-API-KEY":         "<YOUR-API-KEY>",
								"DD-APPLICATION-KEY": "<YOUR-APP-KEY>",
							},
						},
					}},
			},
		},
		Locations: []string{
			"aws:us-east-2",
		},
		Message: "BDD test payload: synthetics_api_test_mcp_payload.json",
		Name:    "Example-Synthetic",
		Options: datadogV1.SyntheticsTestOptions{
			MinFailureDuration: datadog.PtrInt64(10),
			MinLocationFailed:  datadog.PtrInt64(1),
			MonitorName:        datadog.PtrString("Example-Synthetic"),
			MonitorPriority:    datadog.PtrInt32(5),
			Retry: &datadogV1.SyntheticsTestOptionsRetry{
				Count:    datadog.PtrInt64(3),
				Interval: datadog.PtrFloat64(1000),
			},
			TickEvery: datadog.PtrInt64(900),
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
