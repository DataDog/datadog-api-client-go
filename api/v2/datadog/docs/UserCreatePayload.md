# UserCreatePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**UserCreateData**](UserCreateData.md) |  | [optional] 

## Methods

### NewUserCreatePayload

`func NewUserCreatePayload() *UserCreatePayload`

NewUserCreatePayload instantiates a new UserCreatePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCreatePayloadWithDefaults

`func NewUserCreatePayloadWithDefaults() *UserCreatePayload`

NewUserCreatePayloadWithDefaults instantiates a new UserCreatePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *UserCreatePayload) GetData() UserCreateData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UserCreatePayload) GetDataOk() (*UserCreateData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UserCreatePayload) SetData(v UserCreateData)`

SetData sets Data field to given value.

### HasData

`func (o *UserCreatePayload) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


