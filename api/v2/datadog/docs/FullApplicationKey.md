# FullApplicationKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**FullApplicationKeyAttributes**](FullApplicationKeyAttributes.md) |  | [optional] 
**Id** | Pointer to **string** | ID of the application key. | [optional] 
**Relationships** | Pointer to [**ApplicationKeyRelationships**](ApplicationKeyRelationships.md) |  | [optional] 
**Type** | Pointer to [**ApplicationKeysType**](ApplicationKeysType.md) |  | [optional] [default to "application_keys"]

## Methods

### NewFullApplicationKey

`func NewFullApplicationKey() *FullApplicationKey`

NewFullApplicationKey instantiates a new FullApplicationKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullApplicationKeyWithDefaults

`func NewFullApplicationKeyWithDefaults() *FullApplicationKey`

NewFullApplicationKeyWithDefaults instantiates a new FullApplicationKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *FullApplicationKey) GetAttributes() FullApplicationKeyAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FullApplicationKey) GetAttributesOk() (*FullApplicationKeyAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FullApplicationKey) SetAttributes(v FullApplicationKeyAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *FullApplicationKey) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *FullApplicationKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FullApplicationKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FullApplicationKey) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FullApplicationKey) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRelationships

`func (o *FullApplicationKey) GetRelationships() ApplicationKeyRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *FullApplicationKey) GetRelationshipsOk() (*ApplicationKeyRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *FullApplicationKey) SetRelationships(v ApplicationKeyRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *FullApplicationKey) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *FullApplicationKey) GetType() ApplicationKeysType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FullApplicationKey) GetTypeOk() (*ApplicationKeysType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FullApplicationKey) SetType(v ApplicationKeysType)`

SetType sets Type field to given value.

### HasType

`func (o *FullApplicationKey) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


