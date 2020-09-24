# ServiceResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**ServiceResponseAttributes**](ServiceResponseAttributes.md) |  | [optional] 
**Id** | **string** | The service&#39;s ID. | 
**Relationships** | Pointer to [**ServiceRelationships**](ServiceRelationships.md) |  | [optional] 
**Type** | [**ServiceType**](ServiceType.md) |  | [default to "services"]

## Methods

### NewServiceResponseData

`func NewServiceResponseData(id string, type_ ServiceType, ) *ServiceResponseData`

NewServiceResponseData instantiates a new ServiceResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceResponseDataWithDefaults

`func NewServiceResponseDataWithDefaults() *ServiceResponseData`

NewServiceResponseDataWithDefaults instantiates a new ServiceResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *ServiceResponseData) GetAttributes() ServiceResponseAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ServiceResponseData) GetAttributesOk() (*ServiceResponseAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ServiceResponseData) SetAttributes(v ServiceResponseAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ServiceResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *ServiceResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetRelationships

`func (o *ServiceResponseData) GetRelationships() ServiceRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ServiceResponseData) GetRelationshipsOk() (*ServiceRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ServiceResponseData) SetRelationships(v ServiceRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ServiceResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *ServiceResponseData) GetType() ServiceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceResponseData) GetTypeOk() (*ServiceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceResponseData) SetType(v ServiceType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


