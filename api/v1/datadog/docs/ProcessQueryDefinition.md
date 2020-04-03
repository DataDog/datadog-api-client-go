# ProcessQueryDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterBy** | Pointer to **[]string** |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Metric** | Pointer to **string** |  | 
**SearchBy** | Pointer to **string** |  | [optional] 

## Methods

### NewProcessQueryDefinition

`func NewProcessQueryDefinition(metric string, ) *ProcessQueryDefinition`

NewProcessQueryDefinition instantiates a new ProcessQueryDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessQueryDefinitionWithDefaults

`func NewProcessQueryDefinitionWithDefaults() *ProcessQueryDefinition`

NewProcessQueryDefinitionWithDefaults instantiates a new ProcessQueryDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterBy

`func (o *ProcessQueryDefinition) GetFilterBy() []string`

GetFilterBy returns the FilterBy field if non-nil, zero value otherwise.

### GetFilterByOk

`func (o *ProcessQueryDefinition) GetFilterByOk() (*[]string, bool)`

GetFilterByOk returns a tuple with the FilterBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterBy

`func (o *ProcessQueryDefinition) SetFilterBy(v []string)`

SetFilterBy sets FilterBy field to given value.

### HasFilterBy

`func (o *ProcessQueryDefinition) HasFilterBy() bool`

HasFilterBy returns a boolean if a field has been set.

### GetLimit

`func (o *ProcessQueryDefinition) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ProcessQueryDefinition) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ProcessQueryDefinition) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ProcessQueryDefinition) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetMetric

`func (o *ProcessQueryDefinition) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *ProcessQueryDefinition) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *ProcessQueryDefinition) SetMetric(v string)`

SetMetric sets Metric field to given value.


### GetSearchBy

`func (o *ProcessQueryDefinition) GetSearchBy() string`

GetSearchBy returns the SearchBy field if non-nil, zero value otherwise.

### GetSearchByOk

`func (o *ProcessQueryDefinition) GetSearchByOk() (*string, bool)`

GetSearchByOk returns a tuple with the SearchBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBy

`func (o *ProcessQueryDefinition) SetSearchBy(v string)`

SetSearchBy sets SearchBy field to given value.

### HasSearchBy

`func (o *ProcessQueryDefinition) HasSearchBy() bool`

HasSearchBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


