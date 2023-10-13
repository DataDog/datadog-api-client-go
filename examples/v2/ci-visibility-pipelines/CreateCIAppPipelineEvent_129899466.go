// Send pipeline job event returns "Request accepted for processing" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.CIAppCreatePipelineEventRequest{
		Data: &datadogV2.CIAppCreatePipelineEventRequestData{
			Attributes: &datadogV2.CIAppCreatePipelineEventRequestAttributes{
				Resource: datadogV2.CIAppCreatePipelineEventRequestAttributesResource{
					CIAppPipelineEventJob: &datadogV2.CIAppPipelineEventJob{
						End:              time.Now().Add(time.Second * -30),
						Level:            datadogV2.CIAPPPIPELINEEVENTJOBLEVEL_JOB,
						Name:             "Build image",
						Start:            time.Now().Add(time.Second * -120),
						Status:           datadogV2.CIAPPPIPELINEEVENTJOBSTATUS_ERROR,
						Id:               "cf9456de-8b9e-4c27-aa79-27b1e78c1a33",
						PipelineUniqueId: "3eacb6f3-ff04-4e10-8a9c-46e6d054024a",
						PipelineName:     "Deploy to AWS",
						Url:              "https://my-ci-provider.example/jobs/my-jobs/run/1",
					}},
			},
			Type: datadogV2.CIAPPCREATEPIPELINEEVENTREQUESTDATATYPE_CIPIPELINE_RESOURCE_REQUEST.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCIVisibilityPipelinesApi(apiClient)
	resp, r, err := api.CreateCIAppPipelineEvent(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CIVisibilityPipelinesApi.CreateCIAppPipelineEvent`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CIVisibilityPipelinesApi.CreateCIAppPipelineEvent`:\n%s\n", responseContent)
}
