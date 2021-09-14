# SLOCorrectionCreateData

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Attributes** | Pointer to [**SLOCorrectionCreateRequestAttributes**](SLOCorrectionCreateRequestAttributes.md) |  | [optional] 
**Type** | [**SLOCorrectionType**](SLOCorrectionType.md) |  | [default to SLOCORRECTIONTYPE_CORRECTION]

## Methods

### NewSLOCorrectionCreateData

`func NewSLOCorrectionCreateData(type_ SLOCorrectionType) *SLOCorrectionCreateData`

NewSLOCorrectionCreateData instantiates a new SLOCorrectionCreateData object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSLOCorrectionCreateDataWithDefaults

`func NewSLOCorrectionCreateDataWithDefaults() *SLOCorrectionCreateData`

NewSLOCorrectionCreateDataWithDefaults instantiates a new SLOCorrectionCreateData object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAttributes

`func (o *SLOCorrectionCreateData) GetAttributes() SLOCorrectionCreateRequestAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SLOCorrectionCreateData) GetAttributesOk() (*SLOCorrectionCreateRequestAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SLOCorrectionCreateData) SetAttributes(v SLOCorrectionCreateRequestAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *SLOCorrectionCreateData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetType

`func (o *SLOCorrectionCreateData) GetType() SLOCorrectionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SLOCorrectionCreateData) GetTypeOk() (*SLOCorrectionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SLOCorrectionCreateData) SetType(v SLOCorrectionType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


