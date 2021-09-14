# SLOCorrectionListResponse

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Data** | Pointer to [**[]SLOCorrection**](SLOCorrection.md) | The list of of SLO corrections objects. | [optional] 
**Meta** | Pointer to [**ResponseMetaAttributes**](ResponseMetaAttributes.md) |  | [optional] 

## Methods

### NewSLOCorrectionListResponse

`func NewSLOCorrectionListResponse() *SLOCorrectionListResponse`

NewSLOCorrectionListResponse instantiates a new SLOCorrectionListResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSLOCorrectionListResponseWithDefaults

`func NewSLOCorrectionListResponseWithDefaults() *SLOCorrectionListResponse`

NewSLOCorrectionListResponseWithDefaults instantiates a new SLOCorrectionListResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetData

`func (o *SLOCorrectionListResponse) GetData() []SLOCorrection`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SLOCorrectionListResponse) GetDataOk() (*[]SLOCorrection, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SLOCorrectionListResponse) SetData(v []SLOCorrection)`

SetData sets Data field to given value.

### HasData

`func (o *SLOCorrectionListResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *SLOCorrectionListResponse) GetMeta() ResponseMetaAttributes`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SLOCorrectionListResponse) GetMetaOk() (*ResponseMetaAttributes, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SLOCorrectionListResponse) SetMeta(v ResponseMetaAttributes)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SLOCorrectionListResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


