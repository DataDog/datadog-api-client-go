# SLOListResponse

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Data** | Pointer to [**[]ServiceLevelObjective**](ServiceLevelObjective.md) | An array of service level objective objects. | [optional] 
**Errors** | Pointer to **[]string** | An array of error messages. Each endpoint documents how/whether this field is used. | [optional] 
**Metadata** | Pointer to [**SLOListResponseMetadata**](SLOListResponseMetadata.md) |  | [optional] 

## Methods

### NewSLOListResponse

`func NewSLOListResponse() *SLOListResponse`

NewSLOListResponse instantiates a new SLOListResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSLOListResponseWithDefaults

`func NewSLOListResponseWithDefaults() *SLOListResponse`

NewSLOListResponseWithDefaults instantiates a new SLOListResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetData

`func (o *SLOListResponse) GetData() []ServiceLevelObjective`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SLOListResponse) GetDataOk() (*[]ServiceLevelObjective, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SLOListResponse) SetData(v []ServiceLevelObjective)`

SetData sets Data field to given value.

### HasData

`func (o *SLOListResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetErrors

`func (o *SLOListResponse) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *SLOListResponse) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *SLOListResponse) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *SLOListResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetMetadata

`func (o *SLOListResponse) GetMetadata() SLOListResponseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SLOListResponse) GetMetadataOk() (*SLOListResponseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SLOListResponse) SetMetadata(v SLOListResponseMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SLOListResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


