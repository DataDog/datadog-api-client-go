# ApplicationKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** | Name of your application key. | [optional] 
**Owner** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewApplicationKey

`func NewApplicationKey() *ApplicationKey`

NewApplicationKey instantiates a new ApplicationKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationKeyWithDefaults

`func NewApplicationKeyWithDefaults() *ApplicationKey`

NewApplicationKeyWithDefaults instantiates a new ApplicationKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *ApplicationKey) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ApplicationKey) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ApplicationKey) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *ApplicationKey) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetName

`func (o *ApplicationKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationKey) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplicationKey) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *ApplicationKey) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ApplicationKey) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ApplicationKey) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ApplicationKey) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


