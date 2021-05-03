# LogQueryDefinitionGroupBySort

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Aggregation** | **string** | The aggregation method. | 
**Facet** | Pointer to **string** | Facet name. | [optional] 
**Order** | [**WidgetSort**](WidgetSort.md) |  | 

## Methods

### NewLogQueryDefinitionGroupBySort

`func NewLogQueryDefinitionGroupBySort(aggregation string, order WidgetSort, ) *LogQueryDefinitionGroupBySort`

NewLogQueryDefinitionGroupBySort instantiates a new LogQueryDefinitionGroupBySort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogQueryDefinitionGroupBySortWithDefaults

`func NewLogQueryDefinitionGroupBySortWithDefaults() *LogQueryDefinitionGroupBySort`

NewLogQueryDefinitionGroupBySortWithDefaults instantiates a new LogQueryDefinitionGroupBySort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregation

`func (o *LogQueryDefinitionGroupBySort) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *LogQueryDefinitionGroupBySort) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *LogQueryDefinitionGroupBySort) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.


### GetFacet

`func (o *LogQueryDefinitionGroupBySort) GetFacet() string`

GetFacet returns the Facet field if non-nil, zero value otherwise.

### GetFacetOk

`func (o *LogQueryDefinitionGroupBySort) GetFacetOk() (*string, bool)`

GetFacetOk returns a tuple with the Facet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacet

`func (o *LogQueryDefinitionGroupBySort) SetFacet(v string)`

SetFacet sets Facet field to given value.

### HasFacet

`func (o *LogQueryDefinitionGroupBySort) HasFacet() bool`

HasFacet returns a boolean if a field has been set.

### GetOrder

`func (o *LogQueryDefinitionGroupBySort) GetOrder() WidgetSort`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *LogQueryDefinitionGroupBySort) GetOrderOk() (*WidgetSort, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *LogQueryDefinitionGroupBySort) SetOrder(v WidgetSort)`

SetOrder sets Order field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


