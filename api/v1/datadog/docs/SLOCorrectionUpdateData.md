# SLOCorrectionUpdateData

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Attributes** | Pointer to [**SLOCorrectionUpdateRequestAttributes**](SLOCorrectionUpdateRequestAttributes.md) |  | [optional] 
**Type** | Pointer to [**SLOCorrectionType**](SLOCorrectionType.md) |  | [optional] [default to SLOCORRECTIONTYPE_CORRECTION]

## Methods

### NewSLOCorrectionUpdateData

`func NewSLOCorrectionUpdateData() *SLOCorrectionUpdateData`

NewSLOCorrectionUpdateData instantiates a new SLOCorrectionUpdateData object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSLOCorrectionUpdateDataWithDefaults

`func NewSLOCorrectionUpdateDataWithDefaults() *SLOCorrectionUpdateData`

NewSLOCorrectionUpdateDataWithDefaults instantiates a new SLOCorrectionUpdateData object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAttributes

`func (o *SLOCorrectionUpdateData) GetAttributes() SLOCorrectionUpdateRequestAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SLOCorrectionUpdateData) GetAttributesOk() (*SLOCorrectionUpdateRequestAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SLOCorrectionUpdateData) SetAttributes(v SLOCorrectionUpdateRequestAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *SLOCorrectionUpdateData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetType

`func (o *SLOCorrectionUpdateData) GetType() SLOCorrectionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SLOCorrectionUpdateData) GetTypeOk() (*SLOCorrectionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SLOCorrectionUpdateData) SetType(v SLOCorrectionType)`

SetType sets Type field to given value.

### HasType

`func (o *SLOCorrectionUpdateData) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


