# SLOCorrectionResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**SLOCorrectionResponseAttributes**](SLOCorrectionResponseAttributes.md) |  | [optional] 
**Id** | Pointer to **string** | The ID of the SLO correction | [optional] 
**Type** | Pointer to **string** | Should always return \&quot;correction\&quot; | [optional] 

## Methods

### NewSLOCorrectionResponseData

`func NewSLOCorrectionResponseData() *SLOCorrectionResponseData`

NewSLOCorrectionResponseData instantiates a new SLOCorrectionResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSLOCorrectionResponseDataWithDefaults

`func NewSLOCorrectionResponseDataWithDefaults() *SLOCorrectionResponseData`

NewSLOCorrectionResponseDataWithDefaults instantiates a new SLOCorrectionResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *SLOCorrectionResponseData) GetAttributes() SLOCorrectionResponseAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SLOCorrectionResponseData) GetAttributesOk() (*SLOCorrectionResponseAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SLOCorrectionResponseData) SetAttributes(v SLOCorrectionResponseAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *SLOCorrectionResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *SLOCorrectionResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SLOCorrectionResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SLOCorrectionResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SLOCorrectionResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *SLOCorrectionResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SLOCorrectionResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SLOCorrectionResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SLOCorrectionResponseData) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


