// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FullCustomFrameworkDataAttributes Full Framework Data Attributes.
type FullCustomFrameworkDataAttributes struct {
	// Creation Timestamp
	CreatedAt int64 `json:"created_at"`
	// Creator
	CreatedBy string `json:"created_by"`
	// Framework Description
	Description string `json:"description"`
	// Framework Handle
	Handle string `json:"handle"`
	// Framework Icon URL
	IconUrl string `json:"icon_url"`
	// Framework ID
	Id string `json:"id"`
	// Modification Timestamp
	ModifiedAt int64 `json:"modified_at"`
	// Framework Name
	Name string `json:"name"`
	// Organization ID
	OrgId int64 `json:"org_id"`
	// Framework Requirements
	Requirements []CustomFrameworkRequirement `json:"requirements"`
	// Framework Version
	Version string `json:"version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFullCustomFrameworkDataAttributes instantiates a new FullCustomFrameworkDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFullCustomFrameworkDataAttributes(createdAt int64, createdBy string, description string, handle string, iconUrl string, id string, modifiedAt int64, name string, orgId int64, requirements []CustomFrameworkRequirement, version string) *FullCustomFrameworkDataAttributes {
	this := FullCustomFrameworkDataAttributes{}
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.Description = description
	this.Handle = handle
	this.IconUrl = iconUrl
	this.Id = id
	this.ModifiedAt = modifiedAt
	this.Name = name
	this.OrgId = orgId
	this.Requirements = requirements
	this.Version = version
	return &this
}

// NewFullCustomFrameworkDataAttributesWithDefaults instantiates a new FullCustomFrameworkDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFullCustomFrameworkDataAttributesWithDefaults() *FullCustomFrameworkDataAttributes {
	this := FullCustomFrameworkDataAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *FullCustomFrameworkDataAttributes) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *FullCustomFrameworkDataAttributes) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *FullCustomFrameworkDataAttributes) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value.
func (o *FullCustomFrameworkDataAttributes) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *FullCustomFrameworkDataAttributes) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value.
func (o *FullCustomFrameworkDataAttributes) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetDescription returns the Description field value.
func (o *FullCustomFrameworkDataAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *FullCustomFrameworkDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *FullCustomFrameworkDataAttributes) SetDescription(v string) {
	o.Description = v
}

// GetHandle returns the Handle field value.
func (o *FullCustomFrameworkDataAttributes) GetHandle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Handle
}

// GetHandleOk returns a tuple with the Handle field value
// and a boolean to check if the value has been set.
func (o *FullCustomFrameworkDataAttributes) GetHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Handle, true
}

// SetHandle sets field value.
func (o *FullCustomFrameworkDataAttributes) SetHandle(v string) {
	o.Handle = v
}

// GetIconUrl returns the IconUrl field value.
func (o *FullCustomFrameworkDataAttributes) GetIconUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.IconUrl
}

// GetIconUrlOk returns a tuple with the IconUrl field value
// and a boolean to check if the value has been set.
func (o *FullCustomFrameworkDataAttributes) GetIconUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IconUrl, true
}

// SetIconUrl sets field value.
func (o *FullCustomFrameworkDataAttributes) SetIconUrl(v string) {
	o.IconUrl = v
}

// GetId returns the Id field value.
func (o *FullCustomFrameworkDataAttributes) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FullCustomFrameworkDataAttributes) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *FullCustomFrameworkDataAttributes) SetId(v string) {
	o.Id = v
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *FullCustomFrameworkDataAttributes) GetModifiedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *FullCustomFrameworkDataAttributes) GetModifiedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *FullCustomFrameworkDataAttributes) SetModifiedAt(v int64) {
	o.ModifiedAt = v
}

// GetName returns the Name field value.
func (o *FullCustomFrameworkDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FullCustomFrameworkDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *FullCustomFrameworkDataAttributes) SetName(v string) {
	o.Name = v
}

// GetOrgId returns the OrgId field value.
func (o *FullCustomFrameworkDataAttributes) GetOrgId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *FullCustomFrameworkDataAttributes) GetOrgIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value.
func (o *FullCustomFrameworkDataAttributes) SetOrgId(v int64) {
	o.OrgId = v
}

// GetRequirements returns the Requirements field value.
func (o *FullCustomFrameworkDataAttributes) GetRequirements() []CustomFrameworkRequirement {
	if o == nil {
		var ret []CustomFrameworkRequirement
		return ret
	}
	return o.Requirements
}

// GetRequirementsOk returns a tuple with the Requirements field value
// and a boolean to check if the value has been set.
func (o *FullCustomFrameworkDataAttributes) GetRequirementsOk() (*[]CustomFrameworkRequirement, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requirements, true
}

// SetRequirements sets field value.
func (o *FullCustomFrameworkDataAttributes) SetRequirements(v []CustomFrameworkRequirement) {
	o.Requirements = v
}

// GetVersion returns the Version field value.
func (o *FullCustomFrameworkDataAttributes) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *FullCustomFrameworkDataAttributes) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *FullCustomFrameworkDataAttributes) SetVersion(v string) {
	o.Version = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FullCustomFrameworkDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["created_by"] = o.CreatedBy
	toSerialize["description"] = o.Description
	toSerialize["handle"] = o.Handle
	toSerialize["icon_url"] = o.IconUrl
	toSerialize["id"] = o.Id
	toSerialize["modified_at"] = o.ModifiedAt
	toSerialize["name"] = o.Name
	toSerialize["org_id"] = o.OrgId
	toSerialize["requirements"] = o.Requirements
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FullCustomFrameworkDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt    *int64                        `json:"created_at"`
		CreatedBy    *string                       `json:"created_by"`
		Description  *string                       `json:"description"`
		Handle       *string                       `json:"handle"`
		IconUrl      *string                       `json:"icon_url"`
		Id           *string                       `json:"id"`
		ModifiedAt   *int64                        `json:"modified_at"`
		Name         *string                       `json:"name"`
		OrgId        *int64                        `json:"org_id"`
		Requirements *[]CustomFrameworkRequirement `json:"requirements"`
		Version      *string                       `json:"version"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.CreatedBy == nil {
		return fmt.Errorf("required field created_by missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Handle == nil {
		return fmt.Errorf("required field handle missing")
	}
	if all.IconUrl == nil {
		return fmt.Errorf("required field icon_url missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.ModifiedAt == nil {
		return fmt.Errorf("required field modified_at missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.OrgId == nil {
		return fmt.Errorf("required field org_id missing")
	}
	if all.Requirements == nil {
		return fmt.Errorf("required field requirements missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "created_by", "description", "handle", "icon_url", "id", "modified_at", "name", "org_id", "requirements", "version"})
	} else {
		return err
	}
	o.CreatedAt = *all.CreatedAt
	o.CreatedBy = *all.CreatedBy
	o.Description = *all.Description
	o.Handle = *all.Handle
	o.IconUrl = *all.IconUrl
	o.Id = *all.Id
	o.ModifiedAt = *all.ModifiedAt
	o.Name = *all.Name
	o.OrgId = *all.OrgId
	o.Requirements = *all.Requirements
	o.Version = *all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
