# OrganizationAttributes

## Properties

| Name            | Type                     | Description                                       | Notes      |
| --------------- | ------------------------ | ------------------------------------------------- | ---------- |
| **CreatedAt**   | Pointer to **time.Time** | Creation time of the organization.                | [optional] |
| **Description** | Pointer to **string**    | Description of the organization.                  | [optional] |
| **Disabled**    | Pointer to **bool**      | Whether or not the organization is disabled.      | [optional] |
| **ModifiedAt**  | Pointer to **time.Time** | Time of last organization modification.           | [optional] |
| **Name**        | Pointer to **string**    | Name of the organization.                         | [optional] |
| **PublicId**    | Pointer to **string**    | Public ID of the organization.                    | [optional] |
| **Sharing**     | Pointer to **string**    | Sharing type of the organization.                 | [optional] |
| **Url**         | Pointer to **string**    | URL of the site that this organization exists at. | [optional] |

## Methods

### NewOrganizationAttributes

`func NewOrganizationAttributes() *OrganizationAttributes`

NewOrganizationAttributes instantiates a new OrganizationAttributes object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewOrganizationAttributesWithDefaults

`func NewOrganizationAttributesWithDefaults() *OrganizationAttributes`

NewOrganizationAttributesWithDefaults instantiates a new OrganizationAttributes object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCreatedAt

`func (o *OrganizationAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OrganizationAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *OrganizationAttributes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrganizationAttributes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrganizationAttributes) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrganizationAttributes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisabled

`func (o *OrganizationAttributes) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *OrganizationAttributes) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *OrganizationAttributes) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *OrganizationAttributes) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetModifiedAt

`func (o *OrganizationAttributes) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *OrganizationAttributes) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *OrganizationAttributes) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *OrganizationAttributes) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetName

`func (o *OrganizationAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublicId

`func (o *OrganizationAttributes) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *OrganizationAttributes) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *OrganizationAttributes) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *OrganizationAttributes) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

### GetSharing

`func (o *OrganizationAttributes) GetSharing() string`

GetSharing returns the Sharing field if non-nil, zero value otherwise.

### GetSharingOk

`func (o *OrganizationAttributes) GetSharingOk() (*string, bool)`

GetSharingOk returns a tuple with the Sharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharing

`func (o *OrganizationAttributes) SetSharing(v string)`

SetSharing sets Sharing field to given value.

### HasSharing

`func (o *OrganizationAttributes) HasSharing() bool`

HasSharing returns a boolean if a field has been set.

### GetUrl

`func (o *OrganizationAttributes) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OrganizationAttributes) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OrganizationAttributes) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *OrganizationAttributes) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
