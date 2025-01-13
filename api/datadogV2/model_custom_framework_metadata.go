// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomFrameworkMetadata Response object for an organization's custom frameworks.
type CustomFrameworkMetadata struct {
	// Framework Creation Date
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Framework Creator
	CreatedBy *string `json:"created_by,omitempty"`
	// Framework Description
	Description *string `json:"description,omitempty"`
	// Framework Handle
	Handle string `json:"handle"`
	// Framework Icon URL
	IconUrl *string `json:"icon_url,omitempty"`
	// Custom Framework ID
	Id string `json:"id"`
	// Framework Name
	Name string `json:"name"`
	// Org ID
	OrgId int64 `json:"org_id"`
	// Framework Update Date
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// Framework Version
	Version string `json:"version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomFrameworkMetadata instantiates a new CustomFrameworkMetadata object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomFrameworkMetadata(handle string, id string, name string, orgId int64, version string) *CustomFrameworkMetadata {
	this := CustomFrameworkMetadata{}
	this.Handle = handle
	this.Id = id
	this.Name = name
	this.OrgId = orgId
	this.Version = version
	return &this
}

// NewCustomFrameworkMetadataWithDefaults instantiates a new CustomFrameworkMetadata object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomFrameworkMetadataWithDefaults() *CustomFrameworkMetadata {
	this := CustomFrameworkMetadata{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CustomFrameworkMetadata) GetCreatedAt() int64 {
	if o == nil || o.CreatedAt == nil {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFrameworkMetadata) GetCreatedAtOk() (*int64, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CustomFrameworkMetadata) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *CustomFrameworkMetadata) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *CustomFrameworkMetadata) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFrameworkMetadata) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *CustomFrameworkMetadata) HasCreatedBy() bool {
	return o != nil && o.CreatedBy != nil
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *CustomFrameworkMetadata) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CustomFrameworkMetadata) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFrameworkMetadata) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CustomFrameworkMetadata) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CustomFrameworkMetadata) SetDescription(v string) {
	o.Description = &v
}

// GetHandle returns the Handle field value.
func (o *CustomFrameworkMetadata) GetHandle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Handle
}

// GetHandleOk returns a tuple with the Handle field value
// and a boolean to check if the value has been set.
func (o *CustomFrameworkMetadata) GetHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Handle, true
}

// SetHandle sets field value.
func (o *CustomFrameworkMetadata) SetHandle(v string) {
	o.Handle = v
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise.
func (o *CustomFrameworkMetadata) GetIconUrl() string {
	if o == nil || o.IconUrl == nil {
		var ret string
		return ret
	}
	return *o.IconUrl
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFrameworkMetadata) GetIconUrlOk() (*string, bool) {
	if o == nil || o.IconUrl == nil {
		return nil, false
	}
	return o.IconUrl, true
}

// HasIconUrl returns a boolean if a field has been set.
func (o *CustomFrameworkMetadata) HasIconUrl() bool {
	return o != nil && o.IconUrl != nil
}

// SetIconUrl gets a reference to the given string and assigns it to the IconUrl field.
func (o *CustomFrameworkMetadata) SetIconUrl(v string) {
	o.IconUrl = &v
}

// GetId returns the Id field value.
func (o *CustomFrameworkMetadata) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CustomFrameworkMetadata) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *CustomFrameworkMetadata) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value.
func (o *CustomFrameworkMetadata) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CustomFrameworkMetadata) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CustomFrameworkMetadata) SetName(v string) {
	o.Name = v
}

// GetOrgId returns the OrgId field value.
func (o *CustomFrameworkMetadata) GetOrgId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *CustomFrameworkMetadata) GetOrgIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value.
func (o *CustomFrameworkMetadata) SetOrgId(v int64) {
	o.OrgId = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CustomFrameworkMetadata) GetUpdatedAt() int64 {
	if o == nil || o.UpdatedAt == nil {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFrameworkMetadata) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CustomFrameworkMetadata) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *CustomFrameworkMetadata) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

// GetVersion returns the Version field value.
func (o *CustomFrameworkMetadata) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *CustomFrameworkMetadata) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *CustomFrameworkMetadata) SetVersion(v string) {
	o.Version = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomFrameworkMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	toSerialize["handle"] = o.Handle
	if o.IconUrl != nil {
		toSerialize["icon_url"] = o.IconUrl
	}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["org_id"] = o.OrgId
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomFrameworkMetadata) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt   *int64  `json:"created_at,omitempty"`
		CreatedBy   *string `json:"created_by,omitempty"`
		Description *string `json:"description,omitempty"`
		Handle      *string `json:"handle"`
		IconUrl     *string `json:"icon_url,omitempty"`
		Id          *string `json:"id"`
		Name        *string `json:"name"`
		OrgId       *int64  `json:"org_id"`
		UpdatedAt   *int64  `json:"updated_at,omitempty"`
		Version     *string `json:"version"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Handle == nil {
		return fmt.Errorf("required field handle missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.OrgId == nil {
		return fmt.Errorf("required field org_id missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "created_by", "description", "handle", "icon_url", "id", "name", "org_id", "updated_at", "version"})
	} else {
		return err
	}
	o.CreatedAt = all.CreatedAt
	o.CreatedBy = all.CreatedBy
	o.Description = all.Description
	o.Handle = *all.Handle
	o.IconUrl = all.IconUrl
	o.Id = *all.Id
	o.Name = *all.Name
	o.OrgId = *all.OrgId
	o.UpdatedAt = all.UpdatedAt
	o.Version = *all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
