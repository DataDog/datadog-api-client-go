// Create Scanning Rule with should_save_match returns "OK" response

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
	// there is a valid "scanning_group" in the system
	GroupDataID := os.Getenv("GROUP_DATA_ID")


	body := datadogV2.SensitiveDataScannerRuleCreateRequest{
Meta: datadogV2.SensitiveDataScannerMetaVersionOnly{
},
Data: datadogV2.SensitiveDataScannerRuleCreate{
Type: datadogV2.SENSITIVEDATASCANNERRULETYPE_SENSITIVE_DATA_SCANNER_RULE,
Attributes: datadogV2.SensitiveDataScannerRuleAttributes{
Name: datadog.PtrString("Example-Sensitive-Data-Scanner"),
Pattern: datadog.PtrString("pattern"),
TextReplacement: &datadogV2.SensitiveDataScannerTextReplacement{
Type: datadogV2.SENSITIVEDATASCANNERTEXTREPLACEMENTTYPE_REPLACEMENT_STRING.Ptr(),
ReplacementString: datadog.PtrString("REDACTED"),
ShouldSaveMatch: datadog.PtrBool(true),
},
Tags: []string{
"sensitive_data:true",
},
IsEnabled: datadog.PtrBool(true),
Priority: datadog.PtrInt64(1),
},
Relationships: datadogV2.SensitiveDataScannerRuleRelationships{
Group: &datadogV2.SensitiveDataScannerGroupData{
Data: &datadogV2.SensitiveDataScannerGroup{
Type: datadogV2.SENSITIVEDATASCANNERGROUPTYPE_SENSITIVE_DATA_SCANNER_GROUP.Ptr(),
Id: datadog.PtrString(GroupDataID),
},
},
},
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSensitiveDataScannerApi(apiClient)
	resp, r, err := api.CreateScanningRule(ctx, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SensitiveDataScannerApi.CreateScanningRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SensitiveDataScannerApi.CreateScanningRule`:\n%s\n", responseContent)
}