# RelationshipToOrganizationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the organization. | [optional] 
**Type** | Pointer to [**OrganizationsType**](OrganizationsType.md) |  | [optional] [default to "orgs"]

## Methods

### NewRelationshipToOrganizationData

`func NewRelationshipToOrganizationData() *RelationshipToOrganizationData`

NewRelationshipToOrganizationData instantiates a new RelationshipToOrganizationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipToOrganizationDataWithDefaults

`func NewRelationshipToOrganizationDataWithDefaults() *RelationshipToOrganizationData`

NewRelationshipToOrganizationDataWithDefaults instantiates a new RelationshipToOrganizationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RelationshipToOrganizationData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RelationshipToOrganizationData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RelationshipToOrganizationData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RelationshipToOrganizationData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *RelationshipToOrganizationData) GetType() OrganizationsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RelationshipToOrganizationData) GetTypeOk() (*OrganizationsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RelationshipToOrganizationData) SetType(v OrganizationsType)`

SetType sets Type field to given value.

### HasType

`func (o *RelationshipToOrganizationData) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


