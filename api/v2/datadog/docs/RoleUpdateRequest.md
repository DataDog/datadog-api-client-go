# RoleUpdateRequest

## Properties

| Name     | Type                                    | Description | Notes |
| -------- | --------------------------------------- | ----------- | ----- |
| **Data** | [**RoleUpdateData**](RoleUpdateData.md) |             |

## Methods

### NewRoleUpdateRequest

`func NewRoleUpdateRequest(data RoleUpdateData) *RoleUpdateRequest`

NewRoleUpdateRequest instantiates a new RoleUpdateRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewRoleUpdateRequestWithDefaults

`func NewRoleUpdateRequestWithDefaults() *RoleUpdateRequest`

NewRoleUpdateRequestWithDefaults instantiates a new RoleUpdateRequest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetData

`func (o *RoleUpdateRequest) GetData() RoleUpdateData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RoleUpdateRequest) GetDataOk() (*RoleUpdateData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RoleUpdateRequest) SetData(v RoleUpdateData)`

SetData sets Data field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
