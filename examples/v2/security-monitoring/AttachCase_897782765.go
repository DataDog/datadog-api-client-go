// Attach security finding to a case returns "OK" response

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
	body := datadogV2.AttachCaseRequest{
		Data: &datadogV2.AttachCaseRequestData{
			Id: "7d16945b-baf8-411e-ab2a-20fe43af1ea3",
			Relationships: &datadogV2.AttachCaseRequestDataRelationships{
				Findings: datadogV2.Findings{
					Data: []datadogV2.FindingData{
						{
							Id:   "ZGZhMDI3ZjdjMDM3YjJmNzcxNTlhZGMwMjdmZWNiNTZ-MTVlYTNmYWU3NjNlOTNlYTE2YjM4N2JmZmI4Yjk5N2Y=",
							Type: datadogV2.FINDINGDATATYPE_FINDINGS,
						},
					},
				},
			},
			Type: datadogV2.CASEDATATYPE_CASES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.AttachCase(ctx, "7d16945b-baf8-411e-ab2a-20fe43af1ea3", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.AttachCase`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.AttachCase`:\n%s\n", responseContent)
}
