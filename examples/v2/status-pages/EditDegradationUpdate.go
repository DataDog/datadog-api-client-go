// Edit degradation update returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	body := datadogV2.PatchDegradationUpdateRequest{
		Data: &datadogV2.PatchDegradationUpdateRequestData{
			Attributes: &datadogV2.PatchDegradationUpdateRequestDataAttributes{
				Description: datadog.PtrString("We've identified the source of the latency increase and are deploying a fix."),
				Status:      datadogV2.PATCHDEGRADATIONUPDATEREQUESTDATAATTRIBUTESSTATUS_IDENTIFIED.Ptr(),
			},
			Id:   datadog.PtrString("00000000-0000-0000-0000-000000000000"),
			Type: datadogV2.PATCHDEGRADATIONUPDATEREQUESTDATATYPE_DEGRADATION_UPDATES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewStatusPagesApi(apiClient)
	resp, r, err := api.EditDegradationUpdate(ctx, uuid.MustParse("9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d"), uuid.MustParse("9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d"), uuid.MustParse("9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d"), body, *datadogV2.NewEditDegradationUpdateOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatusPagesApi.EditDegradationUpdate`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `StatusPagesApi.EditDegradationUpdate`:\n%s\n", responseContent)
}
