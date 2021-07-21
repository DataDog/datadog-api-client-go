# OrganizationsResponse

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Data** | Pointer to [**[]Organization**](Organization.md) | Array of returned organizations. | [optional] 
**Meta** | Pointer to [**ResponseMetaAttributes**](ResponseMetaAttributes.md) |  | [optional] 

## Methods

### NewOrganizationsResponse

`func NewOrganizationsResponse() *OrganizationsResponse`

NewOrganizationsResponse instantiates a new OrganizationsResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewOrganizationsResponseWithDefaults

`func NewOrganizationsResponseWithDefaults() *OrganizationsResponse`

NewOrganizationsResponseWithDefaults instantiates a new OrganizationsResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetData

`func (o *OrganizationsResponse) GetData() []Organization`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OrganizationsResponse) GetDataOk() (*[]Organization, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OrganizationsResponse) SetData(v []Organization)`

SetData sets Data field to given value.

### HasData

`func (o *OrganizationsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *OrganizationsResponse) GetMeta() ResponseMetaAttributes`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *OrganizationsResponse) GetMetaOk() (*ResponseMetaAttributes, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *OrganizationsResponse) SetMeta(v ResponseMetaAttributes)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *OrganizationsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


