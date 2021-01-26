# SLOCorrectionUpdateRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**SLOCorrectionUpdateRequestAttributes**](SLOCorrectionUpdateRequestAttributes.md) |  | [optional] 
**Type** | Pointer to **string** | Should always be set to \&quot;correction\&quot; | [optional] [default to "correction"]

## Methods

### NewSLOCorrectionUpdateRequestData

`func NewSLOCorrectionUpdateRequestData() *SLOCorrectionUpdateRequestData`

NewSLOCorrectionUpdateRequestData instantiates a new SLOCorrectionUpdateRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSLOCorrectionUpdateRequestDataWithDefaults

`func NewSLOCorrectionUpdateRequestDataWithDefaults() *SLOCorrectionUpdateRequestData`

NewSLOCorrectionUpdateRequestDataWithDefaults instantiates a new SLOCorrectionUpdateRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *SLOCorrectionUpdateRequestData) GetAttributes() SLOCorrectionUpdateRequestAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SLOCorrectionUpdateRequestData) GetAttributesOk() (*SLOCorrectionUpdateRequestAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SLOCorrectionUpdateRequestData) SetAttributes(v SLOCorrectionUpdateRequestAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *SLOCorrectionUpdateRequestData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetType

`func (o *SLOCorrectionUpdateRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SLOCorrectionUpdateRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SLOCorrectionUpdateRequestData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SLOCorrectionUpdateRequestData) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


