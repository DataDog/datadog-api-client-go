# \EventsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEvent**](EventsApi.md#GetEvent) | **Get** /api/v1/events/{event_id} | Get an event
[**ListEvents**](EventsApi.md#ListEvents) | **Get** /api/v1/events | Query the event stream



## GetEvent

> EventResponse GetEvent(ctx, eventId).Execute()

Get an event



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **int64** | The ID of the event. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EventResponse**](EventResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEvents

> EventListResponse ListEvents(ctx).Start(start).End(end).Priority(priority).Sources(sources).Tags(tags).Unaggregated(unaggregated).Execute()

Query the event stream



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int64** | POSIX timestamp. | 
 **end** | **int64** | POSIX timestamp. | 
 **priority** | [**EventPriority**](.md) | Priority of your events, either &#x60;low&#x60; or &#x60;normal&#x60;. | 
 **sources** | **string** | A comma separated string of sources. | 
 **tags** | **string** | A comma separated list indicating what tags, if any, should be used to filter the list of monitors by scope. | 
 **unaggregated** | **bool** | Set unaggregated to &#x60;true&#x60; to return all events within the specified [&#x60;start&#x60;,&#x60;end&#x60;] timeframe. Otherwise if an event is aggregated to a parent event with a timestamp outside of the timeframe, it won&#39;t be available in the output. | 

### Return type

[**EventListResponse**](EventListResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

