// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringBatchMigrationUpdateAttributes Attributes of the batch migration to update. At least one attribute is required.
type SecurityMonitoringBatchMigrationUpdateAttributes struct {
	// The description of the batch migration.
	Description *string `json:"description,omitempty"`
	// The name of the batch migration.
	Name *string `json:"name,omitempty"`
	// Tags of the batch migration. Up to 50 tags, each up to 256 characters.
	Tags []string `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringBatchMigrationUpdateAttributes instantiates a new SecurityMonitoringBatchMigrationUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringBatchMigrationUpdateAttributes() *SecurityMonitoringBatchMigrationUpdateAttributes {
	this := SecurityMonitoringBatchMigrationUpdateAttributes{}
	return &this
}

// NewSecurityMonitoringBatchMigrationUpdateAttributesWithDefaults instantiates a new SecurityMonitoringBatchMigrationUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringBatchMigrationUpdateAttributesWithDefaults() *SecurityMonitoringBatchMigrationUpdateAttributes {
	this := SecurityMonitoringBatchMigrationUpdateAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SecurityMonitoringBatchMigrationUpdateAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationUpdateAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SecurityMonitoringBatchMigrationUpdateAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SecurityMonitoringBatchMigrationUpdateAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SecurityMonitoringBatchMigrationUpdateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SecurityMonitoringBatchMigrationUpdateAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SecurityMonitoringBatchMigrationUpdateAttributes) SetName(v string) {
	o.Name = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SecurityMonitoringBatchMigrationUpdateAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationUpdateAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SecurityMonitoringBatchMigrationUpdateAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *SecurityMonitoringBatchMigrationUpdateAttributes) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringBatchMigrationUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringBatchMigrationUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string  `json:"description,omitempty"`
		Name        *string  `json:"name,omitempty"`
		Tags        []string `json:"tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.Description = all.Description
	o.Name = all.Name
	o.Tags = all.Tags

	return nil
}
