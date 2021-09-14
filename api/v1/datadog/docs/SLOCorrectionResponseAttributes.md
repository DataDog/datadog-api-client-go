# SLOCorrectionResponseAttributes

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Category** | Pointer to [**SLOCorrectionCategory**](SLOCorrectionCategory.md) |  | [optional] 
**Creator** | Pointer to [**Creator**](Creator.md) |  | [optional] 
**Description** | Pointer to **string** | Description of the correction being made. | [optional] 
**End** | Pointer to **int64** | Ending time of the correction in epoch seconds. | [optional] 
**SloId** | Pointer to **string** | ID of the SLO that this correction will be applied to. | [optional] 
**Start** | Pointer to **int64** | Starting time of the correction in epoch seconds. | [optional] 
**Timezone** | Pointer to **string** | The timezone to display in the UI for the correction times (defaults to \&quot;UTC\&quot;). | [optional] 

## Methods

### NewSLOCorrectionResponseAttributes

`func NewSLOCorrectionResponseAttributes() *SLOCorrectionResponseAttributes`

NewSLOCorrectionResponseAttributes instantiates a new SLOCorrectionResponseAttributes object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSLOCorrectionResponseAttributesWithDefaults

`func NewSLOCorrectionResponseAttributesWithDefaults() *SLOCorrectionResponseAttributes`

NewSLOCorrectionResponseAttributesWithDefaults instantiates a new SLOCorrectionResponseAttributes object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCategory

`func (o *SLOCorrectionResponseAttributes) GetCategory() SLOCorrectionCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *SLOCorrectionResponseAttributes) GetCategoryOk() (*SLOCorrectionCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *SLOCorrectionResponseAttributes) SetCategory(v SLOCorrectionCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *SLOCorrectionResponseAttributes) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCreator

`func (o *SLOCorrectionResponseAttributes) GetCreator() Creator`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *SLOCorrectionResponseAttributes) GetCreatorOk() (*Creator, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *SLOCorrectionResponseAttributes) SetCreator(v Creator)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *SLOCorrectionResponseAttributes) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDescription

`func (o *SLOCorrectionResponseAttributes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SLOCorrectionResponseAttributes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SLOCorrectionResponseAttributes) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SLOCorrectionResponseAttributes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnd

`func (o *SLOCorrectionResponseAttributes) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *SLOCorrectionResponseAttributes) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *SLOCorrectionResponseAttributes) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *SLOCorrectionResponseAttributes) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetSloId

`func (o *SLOCorrectionResponseAttributes) GetSloId() string`

GetSloId returns the SloId field if non-nil, zero value otherwise.

### GetSloIdOk

`func (o *SLOCorrectionResponseAttributes) GetSloIdOk() (*string, bool)`

GetSloIdOk returns a tuple with the SloId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloId

`func (o *SLOCorrectionResponseAttributes) SetSloId(v string)`

SetSloId sets SloId field to given value.

### HasSloId

`func (o *SLOCorrectionResponseAttributes) HasSloId() bool`

HasSloId returns a boolean if a field has been set.

### GetStart

`func (o *SLOCorrectionResponseAttributes) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *SLOCorrectionResponseAttributes) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *SLOCorrectionResponseAttributes) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *SLOCorrectionResponseAttributes) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetTimezone

`func (o *SLOCorrectionResponseAttributes) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *SLOCorrectionResponseAttributes) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *SLOCorrectionResponseAttributes) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *SLOCorrectionResponseAttributes) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


