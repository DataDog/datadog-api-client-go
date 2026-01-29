// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringDatasetAttributesResponse Attributes of a security monitoring dataset.
type SecurityMonitoringDatasetAttributesResponse struct {
	// The creation timestamp of the dataset.
	CreatedAt time.Time `json:"createdAt"`
	// The handle of the user who created the dataset.
	CreatedByHandle *string `json:"createdByHandle,omitempty"`
	// The name of the user who created the dataset.
	CreatedByName *string `json:"createdByName,omitempty"`
	// The definition of a dataset, including its data source, name, indexes, and columns.
	Definition SecurityMonitoringDatasetDefinition `json:"definition"`
	// A description of the dataset.
	Description string `json:"description"`
	// The last modification timestamp of the dataset.
	ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
	// The name of the dataset.
	Name string `json:"name"`
	// The organization ID.
	OrgId int64 `json:"orgId"`
	// The handle of the user who last updated the dataset.
	UpdatedByHandle *string `json:"updatedByHandle,omitempty"`
	// The name of the user who last updated the dataset.
	UpdatedByName *string `json:"updatedByName,omitempty"`
	// The version of the dataset.
	Version int64 `json:"version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringDatasetAttributesResponse instantiates a new SecurityMonitoringDatasetAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringDatasetAttributesResponse(createdAt time.Time, definition SecurityMonitoringDatasetDefinition, description string, name string, orgId int64, version int64) *SecurityMonitoringDatasetAttributesResponse {
	this := SecurityMonitoringDatasetAttributesResponse{}
	this.CreatedAt = createdAt
	this.Definition = definition
	this.Description = description
	this.Name = name
	this.OrgId = orgId
	this.Version = version
	return &this
}

// NewSecurityMonitoringDatasetAttributesResponseWithDefaults instantiates a new SecurityMonitoringDatasetAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringDatasetAttributesResponseWithDefaults() *SecurityMonitoringDatasetAttributesResponse {
	this := SecurityMonitoringDatasetAttributesResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *SecurityMonitoringDatasetAttributesResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetAttributesResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *SecurityMonitoringDatasetAttributesResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedByHandle returns the CreatedByHandle field value if set, zero value otherwise.
func (o *SecurityMonitoringDatasetAttributesResponse) GetCreatedByHandle() string {
	if o == nil || o.CreatedByHandle == nil {
		var ret string
		return ret
	}
	return *o.CreatedByHandle
}

// GetCreatedByHandleOk returns a tuple with the CreatedByHandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetAttributesResponse) GetCreatedByHandleOk() (*string, bool) {
	if o == nil || o.CreatedByHandle == nil {
		return nil, false
	}
	return o.CreatedByHandle, true
}

// HasCreatedByHandle returns a boolean if a field has been set.
func (o *SecurityMonitoringDatasetAttributesResponse) HasCreatedByHandle() bool {
	return o != nil && o.CreatedByHandle != nil
}

// SetCreatedByHandle gets a reference to the given string and assigns it to the CreatedByHandle field.
func (o *SecurityMonitoringDatasetAttributesResponse) SetCreatedByHandle(v string) {
	o.CreatedByHandle = &v
}

// GetCreatedByName returns the CreatedByName field value if set, zero value otherwise.
func (o *SecurityMonitoringDatasetAttributesResponse) GetCreatedByName() string {
	if o == nil || o.CreatedByName == nil {
		var ret string
		return ret
	}
	return *o.CreatedByName
}

// GetCreatedByNameOk returns a tuple with the CreatedByName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetAttributesResponse) GetCreatedByNameOk() (*string, bool) {
	if o == nil || o.CreatedByName == nil {
		return nil, false
	}
	return o.CreatedByName, true
}

// HasCreatedByName returns a boolean if a field has been set.
func (o *SecurityMonitoringDatasetAttributesResponse) HasCreatedByName() bool {
	return o != nil && o.CreatedByName != nil
}

// SetCreatedByName gets a reference to the given string and assigns it to the CreatedByName field.
func (o *SecurityMonitoringDatasetAttributesResponse) SetCreatedByName(v string) {
	o.CreatedByName = &v
}

// GetDefinition returns the Definition field value.
func (o *SecurityMonitoringDatasetAttributesResponse) GetDefinition() SecurityMonitoringDatasetDefinition {
	if o == nil {
		var ret SecurityMonitoringDatasetDefinition
		return ret
	}
	return o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetAttributesResponse) GetDefinitionOk() (*SecurityMonitoringDatasetDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Definition, true
}

// SetDefinition sets field value.
func (o *SecurityMonitoringDatasetAttributesResponse) SetDefinition(v SecurityMonitoringDatasetDefinition) {
	o.Definition = v
}

// GetDescription returns the Description field value.
func (o *SecurityMonitoringDatasetAttributesResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetAttributesResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *SecurityMonitoringDatasetAttributesResponse) SetDescription(v string) {
	o.Description = v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *SecurityMonitoringDatasetAttributesResponse) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetAttributesResponse) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *SecurityMonitoringDatasetAttributesResponse) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt != nil
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *SecurityMonitoringDatasetAttributesResponse) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetName returns the Name field value.
func (o *SecurityMonitoringDatasetAttributesResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetAttributesResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SecurityMonitoringDatasetAttributesResponse) SetName(v string) {
	o.Name = v
}

// GetOrgId returns the OrgId field value.
func (o *SecurityMonitoringDatasetAttributesResponse) GetOrgId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetAttributesResponse) GetOrgIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value.
func (o *SecurityMonitoringDatasetAttributesResponse) SetOrgId(v int64) {
	o.OrgId = v
}

// GetUpdatedByHandle returns the UpdatedByHandle field value if set, zero value otherwise.
func (o *SecurityMonitoringDatasetAttributesResponse) GetUpdatedByHandle() string {
	if o == nil || o.UpdatedByHandle == nil {
		var ret string
		return ret
	}
	return *o.UpdatedByHandle
}

// GetUpdatedByHandleOk returns a tuple with the UpdatedByHandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetAttributesResponse) GetUpdatedByHandleOk() (*string, bool) {
	if o == nil || o.UpdatedByHandle == nil {
		return nil, false
	}
	return o.UpdatedByHandle, true
}

// HasUpdatedByHandle returns a boolean if a field has been set.
func (o *SecurityMonitoringDatasetAttributesResponse) HasUpdatedByHandle() bool {
	return o != nil && o.UpdatedByHandle != nil
}

// SetUpdatedByHandle gets a reference to the given string and assigns it to the UpdatedByHandle field.
func (o *SecurityMonitoringDatasetAttributesResponse) SetUpdatedByHandle(v string) {
	o.UpdatedByHandle = &v
}

// GetUpdatedByName returns the UpdatedByName field value if set, zero value otherwise.
func (o *SecurityMonitoringDatasetAttributesResponse) GetUpdatedByName() string {
	if o == nil || o.UpdatedByName == nil {
		var ret string
		return ret
	}
	return *o.UpdatedByName
}

// GetUpdatedByNameOk returns a tuple with the UpdatedByName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetAttributesResponse) GetUpdatedByNameOk() (*string, bool) {
	if o == nil || o.UpdatedByName == nil {
		return nil, false
	}
	return o.UpdatedByName, true
}

// HasUpdatedByName returns a boolean if a field has been set.
func (o *SecurityMonitoringDatasetAttributesResponse) HasUpdatedByName() bool {
	return o != nil && o.UpdatedByName != nil
}

// SetUpdatedByName gets a reference to the given string and assigns it to the UpdatedByName field.
func (o *SecurityMonitoringDatasetAttributesResponse) SetUpdatedByName(v string) {
	o.UpdatedByName = &v
}

// GetVersion returns the Version field value.
func (o *SecurityMonitoringDatasetAttributesResponse) GetVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetAttributesResponse) GetVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *SecurityMonitoringDatasetAttributesResponse) SetVersion(v int64) {
	o.Version = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringDatasetAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.CreatedByHandle != nil {
		toSerialize["createdByHandle"] = o.CreatedByHandle
	}
	if o.CreatedByName != nil {
		toSerialize["createdByName"] = o.CreatedByName
	}
	toSerialize["definition"] = o.Definition
	toSerialize["description"] = o.Description
	if o.ModifiedAt != nil {
		if o.ModifiedAt.Nanosecond() == 0 {
			toSerialize["modifiedAt"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["modifiedAt"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	toSerialize["name"] = o.Name
	toSerialize["orgId"] = o.OrgId
	if o.UpdatedByHandle != nil {
		toSerialize["updatedByHandle"] = o.UpdatedByHandle
	}
	if o.UpdatedByName != nil {
		toSerialize["updatedByName"] = o.UpdatedByName
	}
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringDatasetAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt       *time.Time                           `json:"createdAt"`
		CreatedByHandle *string                              `json:"createdByHandle,omitempty"`
		CreatedByName   *string                              `json:"createdByName,omitempty"`
		Definition      *SecurityMonitoringDatasetDefinition `json:"definition"`
		Description     *string                              `json:"description"`
		ModifiedAt      *time.Time                           `json:"modifiedAt,omitempty"`
		Name            *string                              `json:"name"`
		OrgId           *int64                               `json:"orgId"`
		UpdatedByHandle *string                              `json:"updatedByHandle,omitempty"`
		UpdatedByName   *string                              `json:"updatedByName,omitempty"`
		Version         *int64                               `json:"version"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	if all.Definition == nil {
		return fmt.Errorf("required field definition missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.OrgId == nil {
		return fmt.Errorf("required field orgId missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"createdAt", "createdByHandle", "createdByName", "definition", "description", "modifiedAt", "name", "orgId", "updatedByHandle", "updatedByName", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = *all.CreatedAt
	o.CreatedByHandle = all.CreatedByHandle
	o.CreatedByName = all.CreatedByName
	if all.Definition.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Definition = *all.Definition
	o.Description = *all.Description
	o.ModifiedAt = all.ModifiedAt
	o.Name = *all.Name
	o.OrgId = *all.OrgId
	o.UpdatedByHandle = all.UpdatedByHandle
	o.UpdatedByName = all.UpdatedByName
	o.Version = *all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
