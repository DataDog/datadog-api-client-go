# SLOCorrectionUpdateRequestAttributes

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Category** | Pointer to [**SLOCorrectionCategory**](SLOCorrectionCategory.md) |  | [optional] 
**Description** | Pointer to **string** | Description of the correction being made. | [optional] 
**End** | Pointer to **int64** | Ending time of the correction in epoch seconds. | [optional] 
**Start** | Pointer to **int64** | Starting time of the correction in epoch seconds. | [optional] 
**Timezone** | Pointer to **string** | The timezone to display in the UI for the correction times (defaults to \&quot;UTC\&quot;). | [optional] 

## Methods

### NewSLOCorrectionUpdateRequestAttributes

`func NewSLOCorrectionUpdateRequestAttributes() *SLOCorrectionUpdateRequestAttributes`

NewSLOCorrectionUpdateRequestAttributes instantiates a new SLOCorrectionUpdateRequestAttributes object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSLOCorrectionUpdateRequestAttributesWithDefaults

`func NewSLOCorrectionUpdateRequestAttributesWithDefaults() *SLOCorrectionUpdateRequestAttributes`

NewSLOCorrectionUpdateRequestAttributesWithDefaults instantiates a new SLOCorrectionUpdateRequestAttributes object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCategory

`func (o *SLOCorrectionUpdateRequestAttributes) GetCategory() SLOCorrectionCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *SLOCorrectionUpdateRequestAttributes) GetCategoryOk() (*SLOCorrectionCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *SLOCorrectionUpdateRequestAttributes) SetCategory(v SLOCorrectionCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *SLOCorrectionUpdateRequestAttributes) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *SLOCorrectionUpdateRequestAttributes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SLOCorrectionUpdateRequestAttributes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SLOCorrectionUpdateRequestAttributes) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SLOCorrectionUpdateRequestAttributes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnd

`func (o *SLOCorrectionUpdateRequestAttributes) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *SLOCorrectionUpdateRequestAttributes) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *SLOCorrectionUpdateRequestAttributes) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *SLOCorrectionUpdateRequestAttributes) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetStart

`func (o *SLOCorrectionUpdateRequestAttributes) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *SLOCorrectionUpdateRequestAttributes) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *SLOCorrectionUpdateRequestAttributes) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *SLOCorrectionUpdateRequestAttributes) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetTimezone

`func (o *SLOCorrectionUpdateRequestAttributes) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *SLOCorrectionUpdateRequestAttributes) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *SLOCorrectionUpdateRequestAttributes) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *SLOCorrectionUpdateRequestAttributes) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


