# \LogsPipelinesApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLogsPipeline**](LogsPipelinesApi.md#CreateLogsPipeline) | **Post** /api/v1/logs/config/pipelines | Create a pipeline
[**DeleteLogsPipeline**](LogsPipelinesApi.md#DeleteLogsPipeline) | **Delete** /api/v1/logs/config/pipelines/{pipeline_id} | Delete a pipeline
[**GetLogsPipeline**](LogsPipelinesApi.md#GetLogsPipeline) | **Get** /api/v1/logs/config/pipelines/{pipeline_id} | Get a pipeline
[**GetLogsPipelineOrder**](LogsPipelinesApi.md#GetLogsPipelineOrder) | **Get** /api/v1/logs/config/pipeline-order | Get pipeline order
[**ListLogsPipelines**](LogsPipelinesApi.md#ListLogsPipelines) | **Get** /api/v1/logs/config/pipelines | Get all pipelines
[**UpdateLogsPipeline**](LogsPipelinesApi.md#UpdateLogsPipeline) | **Put** /api/v1/logs/config/pipelines/{pipeline_id} | Update a pipeline
[**UpdateLogsPipelineOrder**](LogsPipelinesApi.md#UpdateLogsPipelineOrder) | **Put** /api/v1/logs/config/pipeline-order | Update pipeline order



## CreateLogsPipeline

> LogsPipeline CreateLogsPipeline(ctx).Body(body).Execute()

Create a pipeline



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := openapiclient.LogsPipeline{Filter: openapiclient.LogsFilter{Query: "Query_example"}, Id: "Id_example", IsEnabled: false, IsReadOnly: false, Name: "Name_example", Processors: []LogsProcessor{openapiclient.LogsProcessor{Grok: openapiclient.LogsGrokParserRules{MatchRules: "MatchRules_example", SupportRules: "SupportRules_example"}, Samples: []string{"Samples_example"), Source: "Source_example", Type: "Type_example", Sources: []string{"Sources_example"), OverrideOnConflict: false, PreserveSource: false, SourceType: "SourceType_example", Target: "Target_example", TargetType: "TargetType_example", NormalizeEndingSlashes: false, IsEncoded: false, Categories: []LogsCategoryProcessorCategories{openapiclient.LogsCategoryProcessor_categories{Filter: openapiclient.LogsFilter{Query: "Query_example"}, Name: "Name_example"}), Expression: "Expression_example", IsReplaceMissing: false, Template: "Template_example", Filter: , Processors: []LogsProcessor{openapiclient.LogsProcessor{Grok: openapiclient.LogsGrokParserRules{MatchRules: "MatchRules_example", SupportRules: "SupportRules_example"}, Samples: []string{"Samples_example"), Source: "Source_example", Type: "Type_example", Sources: []string{"Sources_example"), OverrideOnConflict: false, PreserveSource: false, SourceType: "SourceType_example", Target: "Target_example", TargetType: "TargetType_example", NormalizeEndingSlashes: false, IsEncoded: false, Categories: []LogsCategoryProcessorCategories{openapiclient.LogsCategoryProcessor_categories{Filter: , Name: "Name_example"}), Expression: "Expression_example", IsReplaceMissing: false, Template: "Template_example", Filter: , Processors: []LogsProcessor{), DefaultLookup: "DefaultLookup_example", LookupTable: []string{"LookupTable_example")}), DefaultLookup: "DefaultLookup_example", LookupTable: []string{"LookupTable_example")}), Type: "Type_example"} // LogsPipeline | Definition of the new pipeline.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogsPipelinesApi.CreateLogsPipeline(context.Background(), body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsPipelinesApi.CreateLogsPipeline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLogsPipeline`: LogsPipeline
    fmt.Fprintf(os.Stdout, "Response from `LogsPipelinesApi.CreateLogsPipeline`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLogsPipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**LogsPipeline**](LogsPipeline.md) | Definition of the new pipeline. | 

### Return type

[**LogsPipeline**](LogsPipeline.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLogsPipeline

> DeleteLogsPipeline(ctx, pipelineId).Execute()

Delete a pipeline



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pipelineId := "pipelineId_example" // string | ID of the pipeline to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogsPipelinesApi.DeleteLogsPipeline(context.Background(), pipelineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsPipelinesApi.DeleteLogsPipeline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipelineId** | **string** | ID of the pipeline to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogsPipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogsPipeline

> LogsPipeline GetLogsPipeline(ctx, pipelineId).Execute()

Get a pipeline



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pipelineId := "pipelineId_example" // string | ID of the pipeline to get.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogsPipelinesApi.GetLogsPipeline(context.Background(), pipelineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsPipelinesApi.GetLogsPipeline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogsPipeline`: LogsPipeline
    fmt.Fprintf(os.Stdout, "Response from `LogsPipelinesApi.GetLogsPipeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipelineId** | **string** | ID of the pipeline to get. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogsPipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LogsPipeline**](LogsPipeline.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogsPipelineOrder

> LogsPipelinesOrder GetLogsPipelineOrder(ctx).Execute()

Get pipeline order



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogsPipelinesApi.GetLogsPipelineOrder(context.Background(), ).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsPipelinesApi.GetLogsPipelineOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogsPipelineOrder`: LogsPipelinesOrder
    fmt.Fprintf(os.Stdout, "Response from `LogsPipelinesApi.GetLogsPipelineOrder`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogsPipelineOrderRequest struct via the builder pattern


### Return type

[**LogsPipelinesOrder**](LogsPipelinesOrder.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLogsPipelines

> []LogsPipeline ListLogsPipelines(ctx).Execute()

Get all pipelines



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogsPipelinesApi.ListLogsPipelines(context.Background(), ).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsPipelinesApi.ListLogsPipelines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogsPipelines`: []LogsPipeline
    fmt.Fprintf(os.Stdout, "Response from `LogsPipelinesApi.ListLogsPipelines`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLogsPipelinesRequest struct via the builder pattern


### Return type

[**[]LogsPipeline**](LogsPipeline.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLogsPipeline

> LogsPipeline UpdateLogsPipeline(ctx, pipelineId).Body(body).Execute()

Update a pipeline



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pipelineId := "pipelineId_example" // string | ID of the pipeline to delete.
    body := openapiclient.LogsPipeline{Filter: , Id: "Id_example", IsEnabled: false, IsReadOnly: false, Name: "Name_example", Processors: []LogsProcessor{), Type: "Type_example"} // LogsPipeline | New definition of the pipeline.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogsPipelinesApi.UpdateLogsPipeline(context.Background(), pipelineId, body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsPipelinesApi.UpdateLogsPipeline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLogsPipeline`: LogsPipeline
    fmt.Fprintf(os.Stdout, "Response from `LogsPipelinesApi.UpdateLogsPipeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipelineId** | **string** | ID of the pipeline to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogsPipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**LogsPipeline**](LogsPipeline.md) | New definition of the pipeline. | 

### Return type

[**LogsPipeline**](LogsPipeline.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLogsPipelineOrder

> LogsPipelinesOrder UpdateLogsPipelineOrder(ctx).Body(body).Execute()

Update pipeline order



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := openapiclient.LogsPipelinesOrder{PipelineIds: []string{"PipelineIds_example")} // LogsPipelinesOrder | Object containing the new ordered list of pipeline IDs.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogsPipelinesApi.UpdateLogsPipelineOrder(context.Background(), body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsPipelinesApi.UpdateLogsPipelineOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLogsPipelineOrder`: LogsPipelinesOrder
    fmt.Fprintf(os.Stdout, "Response from `LogsPipelinesApi.UpdateLogsPipelineOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogsPipelineOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**LogsPipelinesOrder**](LogsPipelinesOrder.md) | Object containing the new ordered list of pipeline IDs. | 

### Return type

[**LogsPipelinesOrder**](LogsPipelinesOrder.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

