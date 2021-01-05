# PartialApplicationKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**PartialApplicationKeyAttributes**](PartialApplicationKeyAttributes.md) |  | [optional] 
**Id** | Pointer to **string** | ID of the application key. | [optional] 
**Relationships** | Pointer to [**ApplicationKeyRelationships**](ApplicationKeyRelationships.md) |  | [optional] 
**Type** | Pointer to [**ApplicationKeysType**](ApplicationKeysType.md) |  | [optional] [default to "application_keys"]

## Methods

### NewPartialApplicationKey

`func NewPartialApplicationKey() *PartialApplicationKey`

NewPartialApplicationKey instantiates a new PartialApplicationKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartialApplicationKeyWithDefaults

`func NewPartialApplicationKeyWithDefaults() *PartialApplicationKey`

NewPartialApplicationKeyWithDefaults instantiates a new PartialApplicationKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *PartialApplicationKey) GetAttributes() PartialApplicationKeyAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PartialApplicationKey) GetAttributesOk() (*PartialApplicationKeyAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PartialApplicationKey) SetAttributes(v PartialApplicationKeyAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *PartialApplicationKey) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *PartialApplicationKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PartialApplicationKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PartialApplicationKey) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PartialApplicationKey) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRelationships

`func (o *PartialApplicationKey) GetRelationships() ApplicationKeyRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PartialApplicationKey) GetRelationshipsOk() (*ApplicationKeyRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PartialApplicationKey) SetRelationships(v ApplicationKeyRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PartialApplicationKey) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *PartialApplicationKey) GetType() ApplicationKeysType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PartialApplicationKey) GetTypeOk() (*ApplicationKeysType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PartialApplicationKey) SetType(v ApplicationKeysType)`

SetType sets Type field to given value.

### HasType

`func (o *PartialApplicationKey) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


