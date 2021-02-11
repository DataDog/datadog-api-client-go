# \LogsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListLogs**](LogsApi.md#ListLogs) | **Post** /api/v1/logs-queries/list | Search logs



## ListLogs

> LogsListResponse ListLogs(ctx).Body(body).Execute()

Search logs



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    "time"
    datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
    ctx := context.WithValue(
        context.Background(),
        datadog.ContextAPIKeys,
        map[string]datadog.APIKey{
            "apiKeyAuth": {
                Key: os.Getenv("DD_CLIENT_API_KEY"),
            },
            "appKeyAuth": {
                Key: os.Getenv("DD_CLIENT_APP_KEY"),
            },
        },
    )

    if site, ok := os.LookupEnv("DD_SITE"); ok {
        ctx = context.WithValue(
            ctx,
            datadog.ContextServerVariables,
            map[string]string{"site": site},
        )
    }

    body := *datadog.NewLogsListRequest(*datadog.NewLogsListRequestTime(time.Now(), time.Now())) // LogsListRequest | Logs filter

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.LogsApi.ListLogs(ctx).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.ListLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogs`: LogsListResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from LogsApi.ListLogs:\n%s\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**LogsListRequest**](LogsListRequest.md) | Logs filter | 

### Return type

[**LogsListResponse**](LogsListResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

