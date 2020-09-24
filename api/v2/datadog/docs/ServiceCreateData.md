# ServiceCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**ServiceCreateAttributes**](ServiceCreateAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**ServiceRelationships**](ServiceRelationships.md) |  | [optional] 
**Type** | [**ServiceType**](ServiceType.md) |  | [default to "services"]

## Methods

### NewServiceCreateData

`func NewServiceCreateData(type_ ServiceType, ) *ServiceCreateData`

NewServiceCreateData instantiates a new ServiceCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceCreateDataWithDefaults

`func NewServiceCreateDataWithDefaults() *ServiceCreateData`

NewServiceCreateDataWithDefaults instantiates a new ServiceCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *ServiceCreateData) GetAttributes() ServiceCreateAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ServiceCreateData) GetAttributesOk() (*ServiceCreateAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ServiceCreateData) SetAttributes(v ServiceCreateAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ServiceCreateData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *ServiceCreateData) GetRelationships() ServiceRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ServiceCreateData) GetRelationshipsOk() (*ServiceRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ServiceCreateData) SetRelationships(v ServiceRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ServiceCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *ServiceCreateData) GetType() ServiceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceCreateData) GetTypeOk() (*ServiceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceCreateData) SetType(v ServiceType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


