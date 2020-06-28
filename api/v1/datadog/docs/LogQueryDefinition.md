# LogQueryDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compute** | [**LogsQueryCompute**](LogsQueryCompute.md) |  | [optional] 
**GroupBy** | [**[]LogQueryDefinitionGroupBy**](LogQueryDefinition_group_by.md) | List of tag prefixes to group by in the case of a cluster check. | [optional] 
**Index** | **string** | A coma separated-list of index names. Use \&quot;*\&quot; query all indexes at once. [Multiple Indexes](https://docs.datadoghq.com/logs/indexes/#multiple-indexes) | [optional] 
**MultiCompute** | [**[]LogsQueryCompute**](LogsQueryCompute.md) | This field is mutually exclusive with &#x60;compute&#x60;. | [optional] 
**Search** | [**LogQueryDefinitionSearch**](LogQueryDefinition_search.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


