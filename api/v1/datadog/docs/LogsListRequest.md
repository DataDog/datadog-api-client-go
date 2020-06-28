# LogsListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **string** | For multi-index organizations, the log index in which the request is performed. Default to &#39;*&#39; (all indexes). | [optional] 
**Limit** | **int32** | Number of logs return in the response. | [optional] 
**Query** | **string** | The search query - following the log search syntax. | 
**Sort** | [**LogsSort**](LogsSort.md) |  | [optional] 
**StartAt** | **string** | Hash identifier of the first log to return in the list, available in a log &#x60;id&#x60; attribute. This parameter is used for the pagination feature.  **Note**: This parameter is ignored if the corresponding log is out of the scope of the specified time window. | [optional] 
**Time** | [**LogsListRequestTime**](LogsListRequest_time.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


