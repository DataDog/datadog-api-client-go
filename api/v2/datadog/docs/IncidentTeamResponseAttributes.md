# IncidentTeamResponseAttributes

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Created** | Pointer to **time.Time** | Timestamp of when the incident team was created. | [optional] [readonly] 
**Modified** | Pointer to **time.Time** | Timestamp of when the incident team was modified. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the incident team. | [optional] 

## Methods

### NewIncidentTeamResponseAttributes

`func NewIncidentTeamResponseAttributes() *IncidentTeamResponseAttributes`

NewIncidentTeamResponseAttributes instantiates a new IncidentTeamResponseAttributes object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewIncidentTeamResponseAttributesWithDefaults

`func NewIncidentTeamResponseAttributesWithDefaults() *IncidentTeamResponseAttributes`

NewIncidentTeamResponseAttributesWithDefaults instantiates a new IncidentTeamResponseAttributes object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCreated

`func (o *IncidentTeamResponseAttributes) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *IncidentTeamResponseAttributes) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *IncidentTeamResponseAttributes) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *IncidentTeamResponseAttributes) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *IncidentTeamResponseAttributes) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *IncidentTeamResponseAttributes) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *IncidentTeamResponseAttributes) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *IncidentTeamResponseAttributes) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetName

`func (o *IncidentTeamResponseAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IncidentTeamResponseAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IncidentTeamResponseAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IncidentTeamResponseAttributes) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


