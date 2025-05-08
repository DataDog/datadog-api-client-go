// Create tenancy config returns "Created" response

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
	body := datadogV2.CreateTenancyConfig{
		Data: &datadogV2.CreateTenancyConfigData{
			Attributes: &datadogV2.CreateTenancyConfigDataAttributes{
				AuthCredentials: datadogV2.AuthCredentials{
					Fingerprint: "a7:b5:54:f2:da:a2:d7:b0:ed:f4:79:47:93:64:12:b1",
					PrivateKey: `-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCmMo2jwJXWTt0y
k+X6biZycflZSwOAP/iNeAZPTWwhYxj9pxDvd5OfiIe+o/7eupk/3q+fRsSaztPn
JwI/JnbQz5IT5miLi/apIozg870FFxjrgRxSGjo7BNH0dLKITc5nLDLBnOEzxR2Y
k9+0dFaiNlcodFULlg75trqbILRSc6jn9Tp9G8C5e9cj+LYQuUu2JwIqhCJqcNcU
t+lRL5odBJhZ85KlugKyUg6LN3VQIdOpTtPBMXYA1oBgDCbe5Rw5yzgnd0KtSFf3
GOmLfR95gQshLfbGavLOTh9ioaOj/2hT9HrsEe1VWgX3m1WibqKiPc4OA4BGGToN
9tzN/t89AgMBAAECggEAVFKD4JherXwX6Ih3f6cRZLGFBJP1s8VBM225LdUnTo07
6b4w7n6p7KBV1xjXwGPGS0yNqG88YxsbEkWNc0Ltt6YJBIW7d0nNHSVFewDPX1zH
rP01xEZAUx9v8uqehl+LoHchTXBuJlkVWgt0zdbU+bo+YG0dlSJOeM4IQZrHQqlQ
e4PNk73rot9NSqiKQFXUroaoVPTkUHb3idpLX60K3MgIBoAm4DpJ6cMItb4hyHv5
pNZhHQbr9Eciz2tj+OhQTYKCrAd0gJgl0tC+6L3kzkmiYE3ceGphqWfI9bX52Y96
wpgAtYi6o8wTykgRLabLc6vSQ9RegWEh7P8iSAvAlQKBgQDX5wJhYeWDdG4uPqLC
X3EtnR3y5zYgOd7cVtMr1DIvXa4I8PSIOC4Wnb/5A1S03dJ2e8GJ/qSbl5R2fsDr
XhjIm/KeBPI9p2dVZM8fPoWppR3SgDaHY5qxAED111DnEZuTMl5BO87QZXurTSiF
fbGsWaVqdVieRAQ3b5DEkC9TSwKBgQDFEFgui7iyPhQaQafsjnVbWyrWF821xjTG
b6Bo4FO97c9pw/tbkpfM+dcOU4SsZL8HjwGBUhUsDsVOX7m/sWRjZqNM5t/VR+52
9ygIPEjNyh0b3aARgn8AQ8n+RZvl1Z2A32KCO3MFzhpVKnv2sdSc1TNHQkuJH/rq
eUAm3El6lwKBgQCK8w+jIOAXRB2NAZ66PbaXRqD5rTg2cUguwmpRsNVDiqTw+DJI
YO+4enoMhspDROeofWlHqGzD/j/8KwN59ys4ILV6YXCNoWltmd17HD/luHCDAyUU
6VOrSqCEF7jnnXtktmvWy+kEUevPiW7kyspIQ8GjzDXmVZvpGZIwDyOGFQKBgGtS
l3PiDFimjnQuRbIDc86pPA8VL6dLpvpbWNVFNtY9abSEU6RvldTATGs0+RCaXZ9U
NtGjTnyMHtCsOZE4nx+zikQbiNOzNR/9QwQZMN1Csc+3R7HBjEEsqhmc92aYjArf
ndqnXeFPee/gD1svRkeTpTWt2U146UJBfrqrRilJAoGAQp7FtEtps5I9xK92AVpD
Hj2p1JNKzLCRtWQ8j4jthKqR0iTQ9SwQyjiAvcKc7HdMaG11gmr5XbmKAzelVC+f
o9kEwoumo8yHVn5Ztp4F2cxaD6+MzSJ/I6WesPyePUD7sPeorXByg1UNOXyzqDub
/aU4/sNo2f8epM9l7QGiCtY=
-----END PRIVATE KEY-----`,
				},
				ConfigVersion: datadog.PtrInt64(2),
				HomeRegion:    "us-ashburn-1",
				LogsConfig: &datadogV2.OCILogsConfig{
					CompartmentTagFilters: []string{
						"datadog:true",
						"env:prod",
					},
					Enabled: datadog.PtrBool(true),
					EnabledServices: []string{
						"oacnativeproduction",
					},
				},
				MetricsConfig: &datadogV2.OCIMetricsConfig{
					CompartmentTagFilters: []string{
						"datadog:true",
						"env:prod",
					},
					Enabled: datadog.PtrBool(true),
					ExcludedServices: []string{
						"oci_compute",
					},
				},
				ResourceCollectionEnabled: datadog.PtrBool(true),
				UserOcid:                  "ocid1.user.test",
			},
			Id:   "ocid1.tenancy.dummy_value",
			Type: datadogV2.CREATETENANCYCONFIGDATATYPE_OCI_TENANCY,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOCIIntegrationApi(apiClient)
	resp, r, err := api.CreateTenancyConfig(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OCIIntegrationApi.CreateTenancyConfig`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OCIIntegrationApi.CreateTenancyConfig`:\n%s\n", responseContent)
}
