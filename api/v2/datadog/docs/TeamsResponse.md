# TeamsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]TeamResponseData**](TeamResponseData.md) | An array of teams. | 
**Included** | Pointer to [**[]TeamIncludedItems**](TeamIncludedItems.md) | Included related resources which the user requested. | [optional] [readonly] 
**Meta** | Pointer to [**ServicesResponseMeta**](ServicesResponse_meta.md) |  | [optional] 

## Methods

### NewTeamsResponse

`func NewTeamsResponse(data []TeamResponseData, ) *TeamsResponse`

NewTeamsResponse instantiates a new TeamsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamsResponseWithDefaults

`func NewTeamsResponseWithDefaults() *TeamsResponse`

NewTeamsResponseWithDefaults instantiates a new TeamsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TeamsResponse) GetData() []TeamResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TeamsResponse) GetDataOk() (*[]TeamResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TeamsResponse) SetData(v []TeamResponseData)`

SetData sets Data field to given value.


### GetIncluded

`func (o *TeamsResponse) GetIncluded() []TeamIncludedItems`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *TeamsResponse) GetIncludedOk() (*[]TeamIncludedItems, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *TeamsResponse) SetIncluded(v []TeamIncludedItems)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *TeamsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetMeta

`func (o *TeamsResponse) GetMeta() ServicesResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *TeamsResponse) GetMetaOk() (*ServicesResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *TeamsResponse) SetMeta(v ServicesResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *TeamsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


