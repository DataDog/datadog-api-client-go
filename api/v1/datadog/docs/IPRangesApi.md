# IPRangesApi

All URIs are relative to *https://api.datadoghq.com*

| Method                                        | HTTP request | Description    |
| --------------------------------------------- | ------------ | -------------- |
| [**GetIPRanges**](IPRangesApi.md#GetIPRanges) | **Get** /    | List IP Ranges |

## GetIPRanges

> IPRanges GetIPRanges(ctx)

Get information about Datadog IP ranges.

### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
    ctx := context.Background()


    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.IPRangesApi.GetIPRanges(ctx)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPRangesApi.GetIPRanges`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIPRanges`: IPRanges
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from IPRangesApi.GetIPRanges:\n%s\n", responseContent)
}
```

### Required Parameters

This endpoint does not need any parameter.

### Optional Parameters

This endpoint does not have optional parameters.

### Return type

[**IPRanges**](IPRanges.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
