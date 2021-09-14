# SLOCorrectionCreateRequestAttributes

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Category** | [**SLOCorrectionCategory**](SLOCorrectionCategory.md) |  | 
**Description** | Pointer to **string** | Description of the correction being made. | [optional] 
**End** | **int64** | Ending time of the correction in epoch seconds. | 
**SloId** | **string** | ID of the SLO that this correction will be applied to. | 
**Start** | **int64** | Starting time of the correction in epoch seconds. | 
**Timezone** | Pointer to **string** | The timezone to display in the UI for the correction times (defaults to \&quot;UTC\&quot;). | [optional] 

## Methods

### NewSLOCorrectionCreateRequestAttributes

`func NewSLOCorrectionCreateRequestAttributes(category SLOCorrectionCategory, end int64, sloId string, start int64) *SLOCorrectionCreateRequestAttributes`

NewSLOCorrectionCreateRequestAttributes instantiates a new SLOCorrectionCreateRequestAttributes object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSLOCorrectionCreateRequestAttributesWithDefaults

`func NewSLOCorrectionCreateRequestAttributesWithDefaults() *SLOCorrectionCreateRequestAttributes`

NewSLOCorrectionCreateRequestAttributesWithDefaults instantiates a new SLOCorrectionCreateRequestAttributes object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCategory

`func (o *SLOCorrectionCreateRequestAttributes) GetCategory() SLOCorrectionCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *SLOCorrectionCreateRequestAttributes) GetCategoryOk() (*SLOCorrectionCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *SLOCorrectionCreateRequestAttributes) SetCategory(v SLOCorrectionCategory)`

SetCategory sets Category field to given value.


### GetDescription

`func (o *SLOCorrectionCreateRequestAttributes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SLOCorrectionCreateRequestAttributes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SLOCorrectionCreateRequestAttributes) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SLOCorrectionCreateRequestAttributes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnd

`func (o *SLOCorrectionCreateRequestAttributes) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *SLOCorrectionCreateRequestAttributes) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *SLOCorrectionCreateRequestAttributes) SetEnd(v int64)`

SetEnd sets End field to given value.


### GetSloId

`func (o *SLOCorrectionCreateRequestAttributes) GetSloId() string`

GetSloId returns the SloId field if non-nil, zero value otherwise.

### GetSloIdOk

`func (o *SLOCorrectionCreateRequestAttributes) GetSloIdOk() (*string, bool)`

GetSloIdOk returns a tuple with the SloId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloId

`func (o *SLOCorrectionCreateRequestAttributes) SetSloId(v string)`

SetSloId sets SloId field to given value.


### GetStart

`func (o *SLOCorrectionCreateRequestAttributes) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *SLOCorrectionCreateRequestAttributes) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *SLOCorrectionCreateRequestAttributes) SetStart(v int64)`

SetStart sets Start field to given value.


### GetTimezone

`func (o *SLOCorrectionCreateRequestAttributes) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *SLOCorrectionCreateRequestAttributes) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *SLOCorrectionCreateRequestAttributes) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *SLOCorrectionCreateRequestAttributes) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


