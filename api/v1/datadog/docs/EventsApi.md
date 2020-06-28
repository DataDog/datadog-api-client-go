# \EventsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEvent**](EventsApi.md#GetEvent) | **Get** /api/v1/events/{event_id} | Get an event
[**ListEvents**](EventsApi.md#ListEvents) | **Get** /api/v1/events | Query the event stream



## GetEvent

> EventResponse GetEvent(ctx, eventId)

Get an event

This endpoint allows you to query for event details.  **Note**: If the event you’re querying contains markdown formatting of any kind, you may see characters such as `%`,`\\`,`n` in your output.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **int64**| The ID of the event. | 

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

> EventListResponse ListEvents(ctx, start, end, optional)

Query the event stream

The event stream can be queried and filtered by time, priority, sources and tags.  **Note**: If the event you’re querying contains markdown formatting of any kind, you may see characters such as %,\\,n in your output.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**start** | **int64**| POSIX timestamp. | 
**end** | **int64**| POSIX timestamp. | 
 **optional** | ***ListEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListEventsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **priority** | [**optional.Interface of EventPriority**](.md)| Priority of your events, either &#x60;low&#x60; or &#x60;normal&#x60;. | 
 **sources** | **optional.String**| A comma separated string of sources. | 
 **tags** | **optional.String**| A comma separated list indicating what tags, if any, should be used to filter the list of monitors by scope. | 
 **unaggregated** | **optional.Bool**| Set unaggregated to &#x60;true&#x60; to return all events within the specified [&#x60;start&#x60;,&#x60;end&#x60;] timeframe. Otherwise if an event is aggregated to a parent event with a timestamp outside of the timeframe, it won&#39;t be available in the output. | 

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

