# EventsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------ | ------------ | ------------
[**CreateEvent**](EventsApi.md#CreateEvent) | **Post** /api/v1/events | Post an event
[**GetEvent**](EventsApi.md#GetEvent) | **Get** /api/v1/events/{event_id} | Get an event
[**ListEvents**](EventsApi.md#ListEvents) | **Get** /api/v1/events | Query the event stream



## CreateEvent

> EventCreateResponse CreateEvent(ctx, body)

This endpoint allows you to post events to the stream.
Tag them, set priority and event aggregate them with other events.

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

    body := *datadog.NewEventCreateRequest("Oh boy!", "Did you hear the news today?") // EventCreateRequest | Event request object

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.CreateEvent(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.CreateEvent`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEvent`: EventCreateResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from EventsApi.CreateEvent:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**body** | [**EventCreateRequest**](EventCreateRequest.md) | Event request object | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**EventCreateResponse**](EventCreateResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvent

> EventResponse GetEvent(ctx, eventId)

This endpoint allows you to query for event details.

**Note**: If the event you’re querying contains markdown formatting of any kind,
you may see characters such as `%`,`\`,`n` in your output.

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

    eventId := int64(789) // int64 | The ID of the event.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.GetEvent(ctx, eventId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.GetEvent`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEvent`: EventResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from EventsApi.GetEvent:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**eventId** | **int64** | The ID of the event. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**EventResponse**](EventResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEvents

> EventListResponse ListEvents(ctx, start, end, datadog.ListEventsOptionalParameters{})

The event stream can be queried and filtered by time, priority, sources and tags.

**Notes**:
- If the event you’re querying contains markdown formatting of any kind,
you may see characters such as `%`,`\`,`n` in your output.

- This endpoint returns a maximum of `1000` most recent results. To return additional results,
identify the last timestamp of the last result and set that as the `end` query time to
paginate the results. You can also use the page parameter to specify which set of `1000` results to return.

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

    start := int64(789) // int64 | POSIX timestamp.
    end := int64(789) // int64 | POSIX timestamp.
    priority := datadog.EventPriority("normal") // EventPriority | Priority of your events, either `low` or `normal`. (optional)
    sources := "sources_example" // string | A comma separated string of sources. (optional)
    tags := "host:host0" // string | A comma separated list indicating what tags, if any, should be used to filter the list of monitors by scope. (optional)
    unaggregated := true // bool | Set unaggregated to `true` to return all events within the specified [`start`,`end`] timeframe. Otherwise if an event is aggregated to a parent event with a timestamp outside of the timeframe, it won't be available in the output. Aggregated events with `is_aggregate=true` in the response will still be returned unless exclude_aggregate is set to `true.` (optional)
    excludeAggregate := true // bool | Set `exclude_aggregate` to `true` to only return unaggregated events where `is_aggregate=false` in the response. If the `exclude_aggregate` parameter is set to `true`, then the unaggregated parameter is ignored and will be `true` by default. (optional)
    page := int32(56) // int32 | By default 1000 results are returned per request. Set page to the number of the page to return with `0` being the first page. The page parameter can only be used when either unaggregated or exclude_aggregate is set to `true.` (optional)
    optionalParams := datadog.ListEventsOptionalParameters{
        Priority: &priority,
        Sources: &sources,
        Tags: &tags,
        Unaggregated: &unaggregated,
        ExcludeAggregate: &excludeAggregate,
        Page: &page,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.ListEvents(ctx, start, end, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.ListEvents`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEvents`: EventListResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from EventsApi.ListEvents:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**start** | **int64** | POSIX timestamp. |  |
**end** | **int64** | POSIX timestamp. | 


### Optional Parameters


Other parameters are passed through a pointer to a ListEventsOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**priority** | [**EventPriority**](EventPriority.md) | Priority of your events, either &#x60;low&#x60; or &#x60;normal&#x60;. | 
**sources** | **string** | A comma separated string of sources. | 
**tags** | **string** | A comma separated list indicating what tags, if any, should be used to filter the list of monitors by scope. | 
**unaggregated** | **bool** | Set unaggregated to &#x60;true&#x60; to return all events within the specified [&#x60;start&#x60;,&#x60;end&#x60;] timeframe. Otherwise if an event is aggregated to a parent event with a timestamp outside of the timeframe, it won&#39;t be available in the output. Aggregated events with &#x60;is_aggregate&#x3D;true&#x60; in the response will still be returned unless exclude_aggregate is set to &#x60;true.&#x60; | 
**excludeAggregate** | **bool** | Set &#x60;exclude_aggregate&#x60; to &#x60;true&#x60; to only return unaggregated events where &#x60;is_aggregate&#x3D;false&#x60; in the response. If the &#x60;exclude_aggregate&#x60; parameter is set to &#x60;true&#x60;, then the unaggregated parameter is ignored and will be &#x60;true&#x60; by default. | 
**page** | **int32** | By default 1000 results are returned per request. Set page to the number of the page to return with &#x60;0&#x60; being the first page. The page parameter can only be used when either unaggregated or exclude_aggregate is set to &#x60;true.&#x60; | 

### Return type

[**EventListResponse**](EventListResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

