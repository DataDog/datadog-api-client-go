// Send pipeline event returns "Request accepted for processing" response

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
					CIAppPipelineEventPipeline: &datadogV2.CIAppPipelineEventPipeline{
						End:          time.Now().Add(time.Second * -30),
						Level:        datadogV2.CIAPPPIPELINEEVENTPIPELINELEVEL_PIPELINE,
						Name:         "Deploy to AWS",
						PartialRetry: false,
						Start:        time.Now().Add(time.Second * -120),
						Status:       datadogV2.CIAPPPIPELINEEVENTPIPELINESTATUS_SUCCESS,
						UniqueId:     "3eacb6f3-ff04-4e10-8a9c-46e6d054024a",
						Url:          "https://my-ci-provider.example/pipelines/my-pipeline/run/1",
						Git: *datadogV2.NewNullableCIAppGitInfo(&datadogV2.CIAppGitInfo{
							RepositoryUrl: "https://github.com/DataDog/datadog-agent",
							Sha:           "7f263865994b76066c4612fd1965215e7dcb4cd2",
							AuthorEmail:   "john.doe@email.com",
						}),
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
