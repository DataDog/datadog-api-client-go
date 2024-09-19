// Create a new SLO report returns "OK" response

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
	body := datadogV2.SloReportCreateRequest{
Data: datadogV2.SloReportCreateRequestData{
Attributes: datadogV2.SloReportCreateRequestAttributes{
FromTs: time.Now().AddDate(0, 0, -40).Unix(),
ToTs: time.Now().Unix(),
Query: `slo_type:metric "SLO Reporting Test"`,
Interval: datadogV2.SLOREPORTINTERVAL_MONTHLY.Ptr(),
Timezone: datadog.PtrString("America/New_York"),
},
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateSLOReportJob", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewServiceLevelObjectivesApi(apiClient)
	resp, r, err := api.CreateSLOReportJob(ctx, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectivesApi.CreateSLOReportJob`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ServiceLevelObjectivesApi.CreateSLOReportJob`:\n%s\n", responseContent)
}