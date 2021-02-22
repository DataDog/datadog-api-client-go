# LogsCategoryProcessorCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**LogsFilter**](LogsFilter.md) |  | [optional] 
**Name** | Pointer to **string** | Value to assign to the target attribute. | [optional] 

## Methods

### NewLogsCategoryProcessorCategory

`func NewLogsCategoryProcessorCategory() *LogsCategoryProcessorCategory`

NewLogsCategoryProcessorCategory instantiates a new LogsCategoryProcessorCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsCategoryProcessorCategoryWithDefaults

`func NewLogsCategoryProcessorCategoryWithDefaults() *LogsCategoryProcessorCategory`

NewLogsCategoryProcessorCategoryWithDefaults instantiates a new LogsCategoryProcessorCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *LogsCategoryProcessorCategory) GetFilter() LogsFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *LogsCategoryProcessorCategory) GetFilterOk() (*LogsFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *LogsCategoryProcessorCategory) SetFilter(v LogsFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *LogsCategoryProcessorCategory) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetName

`func (o *LogsCategoryProcessorCategory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogsCategoryProcessorCategory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LogsCategoryProcessorCategory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LogsCategoryProcessorCategory) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


