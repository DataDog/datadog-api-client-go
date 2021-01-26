# ApplicationKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**FullApplicationKey**](FullApplicationKey.md) |  | [optional] 
**Included** | Pointer to [**[]ApplicationKeyResponseIncludedItem**](ApplicationKeyResponseIncludedItem.md) | Array of objects related to the application key. | [optional] 

## Methods

### NewApplicationKeyResponse

`func NewApplicationKeyResponse() *ApplicationKeyResponse`

NewApplicationKeyResponse instantiates a new ApplicationKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationKeyResponseWithDefaults

`func NewApplicationKeyResponseWithDefaults() *ApplicationKeyResponse`

NewApplicationKeyResponseWithDefaults instantiates a new ApplicationKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ApplicationKeyResponse) GetData() FullApplicationKey`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApplicationKeyResponse) GetDataOk() (*FullApplicationKey, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApplicationKeyResponse) SetData(v FullApplicationKey)`

SetData sets Data field to given value.

### HasData

`func (o *ApplicationKeyResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetIncluded

`func (o *ApplicationKeyResponse) GetIncluded() []ApplicationKeyResponseIncludedItem`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *ApplicationKeyResponse) GetIncludedOk() (*[]ApplicationKeyResponseIncludedItem, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *ApplicationKeyResponse) SetIncluded(v []ApplicationKeyResponseIncludedItem)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *ApplicationKeyResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


