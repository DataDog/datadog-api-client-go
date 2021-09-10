# SnapshotsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------ | ------------ | ------------
[**GetGraphSnapshot**](SnapshotsApi.md#GetGraphSnapshot) | **Get** /api/v1/graph/snapshot | Take graph snapshots



## GetGraphSnapshot

> GraphSnapshot GetGraphSnapshot(ctx, start, end, datadog.GetGraphSnapshotOptionalParameters{})

Take graph snapshots.
**Note**: When a snapshot is created, there is some delay before it is available.

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
    ctx := datadog.NewDefaultContext(context.Background())

    start := int64(789) // int64 | The POSIX timestamp of the start of the query.
    end := int64(789) // int64 | The POSIX timestamp of the end of the query.
    metricQuery := "metricQuery_example" // string | The metric query. (optional)
    eventQuery := "eventQuery_example" // string | A query that adds event bands to the graph. (optional)
    graphDef := "graphDef_example" // string | A JSON document defining the graph. `graph_def` can be used instead of `metric_query`. The JSON document uses the [grammar defined here](https://docs.datadoghq.com/graphing/graphing_json/#grammar) and should be formatted to a single line then URL encoded. (optional)
    title := "title_example" // string | A title for the graph. If no title is specified, the graph does not have a title. (optional)
    optionalParams := datadog.GetGraphSnapshotOptionalParameters{
        MetricQuery: &metricQuery,
        EventQuery: &eventQuery,
        GraphDef: &graphDef,
        Title: &title,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.SnapshotsApi.GetGraphSnapshot(ctx, start, end, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsApi.GetGraphSnapshot`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGraphSnapshot`: GraphSnapshot
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SnapshotsApi.GetGraphSnapshot:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**start** | **int64** | The POSIX timestamp of the start of the query. |  |
**end** | **int64** | The POSIX timestamp of the end of the query. | 


### Optional Parameters


Other parameters are passed through a pointer to a GetGraphSnapshotOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**metricQuery** | **string** | The metric query. | 
**eventQuery** | **string** | A query that adds event bands to the graph. | 
**graphDef** | **string** | A JSON document defining the graph. &#x60;graph_def&#x60; can be used instead of &#x60;metric_query&#x60;. The JSON document uses the [grammar defined here](https://docs.datadoghq.com/graphing/graphing_json/#grammar) and should be formatted to a single line then URL encoded. | 
**title** | **string** | A title for the graph. If no title is specified, the graph does not have a title. | 

### Return type

[**GraphSnapshot**](GraphSnapshot.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

