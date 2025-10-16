// Send pipeline event with custom provider returns "Request accepted for processing" response

package main


import (
	"context"
	"fmt"
	"os"

    "github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	body := datadogV2.CIAppCreatePipelineEventRequest{
Data: &datadogV2.CIAppCreatePipelineEventRequestDataSingleOrArray{
CIAppCreatePipelineEventRequestData: &datadogV2.CIAppCreatePipelineEventRequestData{
Attributes: &datadogV2.CIAppCreatePipelineEventRequestAttributes{
ProviderName: datadog.PtrString("example-provider"),
Resource: datadogV2.CIAppCreatePipelineEventRequestAttributesResource{
CIAppPipelineEventPipeline: &datadogV2.CIAppPipelineEventPipeline{
CIAppPipelineEventFinishedPipeline: &datadogV2.CIAppPipelineEventFinishedPipeline{
Level: datadogV2.CIAPPPIPELINEEVENTPIPELINELEVEL_PIPELINE,
UniqueId: "3eacb6f3-ff04-4e10-8a9c-46e6d054024a",
Name: "Deploy to AWS",
Url: "https://my-ci-provider.example/pipelines/my-pipeline/run/1",
Start: time.Now().Add(time.Second*-120),
End: time.Now().Add(time.Second*-30),
Status: datadogV2.CIAPPPIPELINEEVENTPIPELINESTATUS_SUCCESS,
PartialRetry: false,
Git: *datadogV2.NewNullableCIAppGitInfo(&datadogV2.CIAppGitInfo{
RepositoryUrl: "https://github.com/DataDog/datadog-agent",
Sha: "7f263865994b76066c4612fd1965215e7dcb4cd2",
AuthorEmail: "john.doe@email.com",
}),
}}},
},
Type: datadogV2.CIAPPCREATEPIPELINEEVENTREQUESTDATATYPE_CIPIPELINE_RESOURCE_REQUEST.Ptr(),
}},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCIVisibilityPipelinesApi(apiClient)
	resp, r, err := api.CreateCIAppPipelineEvent(ctx, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CIVisibilityPipelinesApi.CreateCIAppPipelineEvent`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CIVisibilityPipelinesApi.CreateCIAppPipelineEvent`:\n%s\n", responseContent)
}