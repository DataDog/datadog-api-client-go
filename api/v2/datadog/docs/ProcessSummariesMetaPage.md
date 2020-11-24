# ProcessSummariesMetaPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**After** | Pointer to **string** | The cursor used to get the next results, if any. To make the next request, use the same parameters with the addition of the &#x60;page[cursor]&#x60;. | [optional] 
**Size** | Pointer to **int32** | Number of results returned. | [optional] 

## Methods

### NewProcessSummariesMetaPage

`func NewProcessSummariesMetaPage() *ProcessSummariesMetaPage`

NewProcessSummariesMetaPage instantiates a new ProcessSummariesMetaPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessSummariesMetaPageWithDefaults

`func NewProcessSummariesMetaPageWithDefaults() *ProcessSummariesMetaPage`

NewProcessSummariesMetaPageWithDefaults instantiates a new ProcessSummariesMetaPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfter

`func (o *ProcessSummariesMetaPage) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *ProcessSummariesMetaPage) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *ProcessSummariesMetaPage) SetAfter(v string)`

SetAfter sets After field to given value.

### HasAfter

`func (o *ProcessSummariesMetaPage) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### GetSize

`func (o *ProcessSummariesMetaPage) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ProcessSummariesMetaPage) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ProcessSummariesMetaPage) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ProcessSummariesMetaPage) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


