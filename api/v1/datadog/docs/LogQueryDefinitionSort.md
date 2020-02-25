# LogQueryDefinitionSort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregation** | Pointer to **string** |  | 
**Facet** | Pointer to **string** |  | [optional] 
**Order** | Pointer to [**WidgetSort**](WidgetSort.md) |  | 

## Methods

### NewLogQueryDefinitionSort

`func NewLogQueryDefinitionSort(aggregation string, order WidgetSort, ) *LogQueryDefinitionSort`

NewLogQueryDefinitionSort instantiates a new LogQueryDefinitionSort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogQueryDefinitionSortWithDefaults

`func NewLogQueryDefinitionSortWithDefaults() *LogQueryDefinitionSort`

NewLogQueryDefinitionSortWithDefaults instantiates a new LogQueryDefinitionSort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregation

`func (o *LogQueryDefinitionSort) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *LogQueryDefinitionSort) GetAggregationOk() (string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAggregation

`func (o *LogQueryDefinitionSort) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### SetAggregation

`func (o *LogQueryDefinitionSort) SetAggregation(v string)`

SetAggregation gets a reference to the given string and assigns it to the Aggregation field.

### GetFacet

`func (o *LogQueryDefinitionSort) GetFacet() string`

GetFacet returns the Facet field if non-nil, zero value otherwise.

### GetFacetOk

`func (o *LogQueryDefinitionSort) GetFacetOk() (string, bool)`

GetFacetOk returns a tuple with the Facet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFacet

`func (o *LogQueryDefinitionSort) HasFacet() bool`

HasFacet returns a boolean if a field has been set.

### SetFacet

`func (o *LogQueryDefinitionSort) SetFacet(v string)`

SetFacet gets a reference to the given string and assigns it to the Facet field.

### GetOrder

`func (o *LogQueryDefinitionSort) GetOrder() WidgetSort`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *LogQueryDefinitionSort) GetOrderOk() (WidgetSort, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOrder

`func (o *LogQueryDefinitionSort) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### SetOrder

`func (o *LogQueryDefinitionSort) SetOrder(v WidgetSort)`

SetOrder gets a reference to the given WidgetSort and assigns it to the Order field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


