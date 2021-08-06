# PartialApplicationKeyResponse

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Data** | Pointer to [**PartialApplicationKey**](PartialApplicationKey.md) |  | [optional] 
**Included** | Pointer to [**[]ApplicationKeyResponseIncludedItem**](ApplicationKeyResponseIncludedItem.md) | Array of objects related to the application key. | [optional] 

## Methods

### NewPartialApplicationKeyResponse

`func NewPartialApplicationKeyResponse() *PartialApplicationKeyResponse`

NewPartialApplicationKeyResponse instantiates a new PartialApplicationKeyResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewPartialApplicationKeyResponseWithDefaults

`func NewPartialApplicationKeyResponseWithDefaults() *PartialApplicationKeyResponse`

NewPartialApplicationKeyResponseWithDefaults instantiates a new PartialApplicationKeyResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetData

`func (o *PartialApplicationKeyResponse) GetData() PartialApplicationKey`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PartialApplicationKeyResponse) GetDataOk() (*PartialApplicationKey, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PartialApplicationKeyResponse) SetData(v PartialApplicationKey)`

SetData sets Data field to given value.

### HasData

`func (o *PartialApplicationKeyResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetIncluded

`func (o *PartialApplicationKeyResponse) GetIncluded() []ApplicationKeyResponseIncludedItem`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *PartialApplicationKeyResponse) GetIncludedOk() (*[]ApplicationKeyResponseIncludedItem, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *PartialApplicationKeyResponse) SetIncluded(v []ApplicationKeyResponseIncludedItem)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *PartialApplicationKeyResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


