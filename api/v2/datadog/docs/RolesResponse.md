# RolesResponse

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Data** | Pointer to [**[]Role**](Role.md) | Array of returned roles. | [optional] 
**Meta** | Pointer to [**ResponseMetaAttributes**](ResponseMetaAttributes.md) |  | [optional] 

## Methods

### NewRolesResponse

`func NewRolesResponse() *RolesResponse`

NewRolesResponse instantiates a new RolesResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewRolesResponseWithDefaults

`func NewRolesResponseWithDefaults() *RolesResponse`

NewRolesResponseWithDefaults instantiates a new RolesResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetData

`func (o *RolesResponse) GetData() []Role`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RolesResponse) GetDataOk() (*[]Role, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RolesResponse) SetData(v []Role)`

SetData sets Data field to given value.

### HasData

`func (o *RolesResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *RolesResponse) GetMeta() ResponseMetaAttributes`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *RolesResponse) GetMetaOk() (*ResponseMetaAttributes, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *RolesResponse) SetMeta(v ResponseMetaAttributes)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *RolesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


