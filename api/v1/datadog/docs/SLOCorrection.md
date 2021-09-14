# SLOCorrection

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Attributes** | Pointer to [**SLOCorrectionResponseAttributes**](SLOCorrectionResponseAttributes.md) |  | [optional] 
**Id** | Pointer to **string** | The ID of the SLO correction. | [optional] 
**Type** | Pointer to [**SLOCorrectionType**](SLOCorrectionType.md) |  | [optional] [default to SLOCORRECTIONTYPE_CORRECTION]

## Methods

### NewSLOCorrection

`func NewSLOCorrection() *SLOCorrection`

NewSLOCorrection instantiates a new SLOCorrection object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSLOCorrectionWithDefaults

`func NewSLOCorrectionWithDefaults() *SLOCorrection`

NewSLOCorrectionWithDefaults instantiates a new SLOCorrection object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAttributes

`func (o *SLOCorrection) GetAttributes() SLOCorrectionResponseAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SLOCorrection) GetAttributesOk() (*SLOCorrectionResponseAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SLOCorrection) SetAttributes(v SLOCorrectionResponseAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *SLOCorrection) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *SLOCorrection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SLOCorrection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SLOCorrection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SLOCorrection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *SLOCorrection) GetType() SLOCorrectionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SLOCorrection) GetTypeOk() (*SLOCorrectionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SLOCorrection) SetType(v SLOCorrectionType)`

SetType sets Type field to given value.

### HasType

`func (o *SLOCorrection) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


