# OrgListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Orgs** | Pointer to [**[]Org**](Org.md) |  | [optional] 

## Methods

### NewOrgListResponse

`func NewOrgListResponse() *OrgListResponse`

NewOrgListResponse instantiates a new OrgListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgListResponseWithDefaults

`func NewOrgListResponseWithDefaults() *OrgListResponse`

NewOrgListResponseWithDefaults instantiates a new OrgListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgs

`func (o *OrgListResponse) GetOrgs() []Org`

GetOrgs returns the Orgs field if non-nil, zero value otherwise.

### GetOrgsOk

`func (o *OrgListResponse) GetOrgsOk() ([]Org, bool)`

GetOrgsOk returns a tuple with the Orgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOrgs

`func (o *OrgListResponse) HasOrgs() bool`

HasOrgs returns a boolean if a field has been set.

### SetOrgs

`func (o *OrgListResponse) SetOrgs(v []Org)`

SetOrgs gets a reference to the given []Org and assigns it to the Orgs field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


