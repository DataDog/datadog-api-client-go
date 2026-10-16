// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringBatchMigrationCreateAttributes Attributes of the batch migration to create.
type SecurityMonitoringBatchMigrationCreateAttributes struct {
	// The description of the batch migration.
	Description *string `json:"description,omitempty"`
	// The name of the batch migration.
	Name *string `json:"name,omitempty"`
	// The initial resources of the batch migration. From 1 to 100 resources.
	Resources []SecurityMonitoringBatchMigrationResourceCreateItem `json:"resources"`
	// Tags of the batch migration. Up to 50 tags, each up to 256 characters.
	Tags []string `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringBatchMigrationCreateAttributes instantiates a new SecurityMonitoringBatchMigrationCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringBatchMigrationCreateAttributes(resources []SecurityMonitoringBatchMigrationResourceCreateItem) *SecurityMonitoringBatchMigrationCreateAttributes {
	this := SecurityMonitoringBatchMigrationCreateAttributes{}
	this.Resources = resources
	return &this
}

// NewSecurityMonitoringBatchMigrationCreateAttributesWithDefaults instantiates a new SecurityMonitoringBatchMigrationCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringBatchMigrationCreateAttributesWithDefaults() *SecurityMonitoringBatchMigrationCreateAttributes {
	this := SecurityMonitoringBatchMigrationCreateAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SecurityMonitoringBatchMigrationCreateAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationCreateAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SecurityMonitoringBatchMigrationCreateAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SecurityMonitoringBatchMigrationCreateAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SecurityMonitoringBatchMigrationCreateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SecurityMonitoringBatchMigrationCreateAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SecurityMonitoringBatchMigrationCreateAttributes) SetName(v string) {
	o.Name = &v
}

// GetResources returns the Resources field value.
func (o *SecurityMonitoringBatchMigrationCreateAttributes) GetResources() []SecurityMonitoringBatchMigrationResourceCreateItem {
	if o == nil {
		var ret []SecurityMonitoringBatchMigrationResourceCreateItem
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationCreateAttributes) GetResourcesOk() (*[]SecurityMonitoringBatchMigrationResourceCreateItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resources, true
}

// SetResources sets field value.
func (o *SecurityMonitoringBatchMigrationCreateAttributes) SetResources(v []SecurityMonitoringBatchMigrationResourceCreateItem) {
	o.Resources = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SecurityMonitoringBatchMigrationCreateAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationCreateAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SecurityMonitoringBatchMigrationCreateAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *SecurityMonitoringBatchMigrationCreateAttributes) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringBatchMigrationCreateAttributes) MarshalJSON() ([]byte, error) {
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
	toSerialize["resources"] = o.Resources
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringBatchMigrationCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string                                               `json:"description,omitempty"`
		Name        *string                                               `json:"name,omitempty"`
		Resources   *[]SecurityMonitoringBatchMigrationResourceCreateItem `json:"resources"`
		Tags        []string                                              `json:"tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Resources == nil {
		return fmt.Errorf("required field resources missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "name", "resources", "tags"})
	} else {
		return err
	}
	o.Description = all.Description
	o.Name = all.Name
	o.Resources = *all.Resources
	o.Tags = all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
