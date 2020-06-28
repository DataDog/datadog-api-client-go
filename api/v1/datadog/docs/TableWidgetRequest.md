# TableWidgetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregator** | [**WidgetAggregator**](WidgetAggregator.md) |  | [optional] 
**Alias** | **string** | The column name (defaults to the metric name). | [optional] 
**ApmQuery** | [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 
**ConditionalFormats** | [**[]WidgetConditionalFormat**](WidgetConditionalFormat.md) | List of conditional formats. | [optional] 
**EventQuery** | [**EventQueryDefinition**](EventQueryDefinition.md) |  | [optional] 
**Limit** | **int64** | For metric queries, the number of lines to show in the table. Only one request should have this property. | [optional] 
**LogQuery** | [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 
**NetworkQuery** | [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 
**Order** | [**WidgetSort**](WidgetSort.md) |  | [optional] 
**ProcessQuery** | [**ProcessQueryDefinition**](ProcessQueryDefinition.md) |  | [optional] 
**Q** | **string** | Query definition. | [optional] 
**RumQuery** | [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


