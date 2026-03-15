// Register a token returns "OK" response

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
	body := datadogV2.OnPremManagementServiceRegisterTokenRequest{
		Data: datadogV2.OnPremManagementServiceRegisterTokenDataRequest{
			Attributes: datadogV2.OnPremManagementServiceRegisterTokenAttributes{
				AppId:        datadog.PtrUUID(uuid.MustParse("9a93d314-ca37-461d-b18e-0587f03c6e2a")),
				AppVersion:   datadog.PtrInt64(5),
				ConnectionId: uuid.MustParse("2f66ec56-d1f3-4a18-908d-5e8c12dfb3b0"),
				QueryId:      datadog.PtrUUID(uuid.MustParse("8744d459-18aa-405b-821e-df9bb101c01e")),
				RunnerId:     "runner-GBUyh2Tm6oKS6ap4kt8Bbx",
			},
			Type: datadogV2.ONPREMMANAGEMENTSERVICEREGISTERTOKENTYPE_REGISTERTOKENREQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.RegisterOnPremManagementServiceToken", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOnPremManagementServiceApi(apiClient)
	resp, r, err := api.RegisterOnPremManagementServiceToken(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnPremManagementServiceApi.RegisterOnPremManagementServiceToken`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OnPremManagementServiceApi.RegisterOnPremManagementServiceToken`:\n%s\n", responseContent)
}
