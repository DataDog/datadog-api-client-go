# SLOCorrectionListResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**SLOCorrectionResponseAttributes**](SLOCorrectionResponseAttributes.md) |  | [optional] 
**Id** | Pointer to **string** | The ID of the SLO correction | [optional] 
**Type** | Pointer to **string** | Should always be set to \&quot;correction\&quot; | [optional] [default to "correction"]

## Methods

### NewSLOCorrectionListResponseData

`func NewSLOCorrectionListResponseData() *SLOCorrectionListResponseData`

NewSLOCorrectionListResponseData instantiates a new SLOCorrectionListResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSLOCorrectionListResponseDataWithDefaults

`func NewSLOCorrectionListResponseDataWithDefaults() *SLOCorrectionListResponseData`

NewSLOCorrectionListResponseDataWithDefaults instantiates a new SLOCorrectionListResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *SLOCorrectionListResponseData) GetAttributes() SLOCorrectionResponseAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SLOCorrectionListResponseData) GetAttributesOk() (*SLOCorrectionResponseAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SLOCorrectionListResponseData) SetAttributes(v SLOCorrectionResponseAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *SLOCorrectionListResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *SLOCorrectionListResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SLOCorrectionListResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SLOCorrectionListResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SLOCorrectionListResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *SLOCorrectionListResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SLOCorrectionListResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SLOCorrectionListResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SLOCorrectionListResponseData) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


