// Create a Google Security Operations custom destination returns "OK" response

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
	body := datadogV2.CustomDestinationCreateRequest{
		Data: &datadogV2.CustomDestinationCreateRequestDefinition{
			Attributes: datadogV2.CustomDestinationCreateRequestAttributes{
				Enabled:     datadog.PtrBool(false),
				ForwardTags: datadog.PtrBool(false),
				ForwardTagsRestrictionList: []string{
					"datacenter",
					"host",
				},
				ForwardTagsRestrictionListType: datadogV2.CUSTOMDESTINATIONATTRIBUTETAGSRESTRICTIONLISTTYPE_ALLOW_LIST.Ptr(),
				ForwarderDestination: datadogV2.CustomDestinationForwardDestination{
					CustomDestinationForwardDestinationGoogleSecurityOperations: &datadogV2.CustomDestinationForwardDestinationGoogleSecurityOperations{
						Type:             datadogV2.CUSTOMDESTINATIONFORWARDDESTINATIONGOOGLESECURITYOPERATIONSTYPE_GOOGLE_SECURITY_OPERATIONS,
						CustomerId:       "123-456-7890",
						RegionalEndpoint: "https://malachiteingestion-pa.googleapis.com",
						Namespace:        "google-security-operations-namespace",
						Auth: datadogV2.CustomDestinationGoogleSecurityOperationsDestinationAuth{
							Type:         datadogV2.CUSTOMDESTINATIONGOOGLESECURITYOPERATIONSDESTINATIONAUTHTYPE_GCP_PRIVATE_KEY,
							ProjectId:    "gcp-project",
							PrivateKeyId: "abc12345678",
							ClientEmail:  "client@example.com",
							ClientId:     "def123456",
							PrivateKey: `-----BEGIN PRIVATE KEY-----
MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBK...
-----END PRIVATE KEY-----
`,
						},
					}},
				Name:  "Nginx logs",
				Query: datadog.PtrString("source:nginx"),
			},
			Type: datadogV2.CUSTOMDESTINATIONTYPE_CUSTOM_DESTINATION,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLogsCustomDestinationsApi(apiClient)
	resp, r, err := api.CreateLogsCustomDestination(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsCustomDestinationsApi.CreateLogsCustomDestination`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsCustomDestinationsApi.CreateLogsCustomDestination`:\n%s\n", responseContent)
}
