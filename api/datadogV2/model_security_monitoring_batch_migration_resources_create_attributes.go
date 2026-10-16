// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringBatchMigrationResourcesCreateAttributes Attributes of the resources to add to a batch migration.
type SecurityMonitoringBatchMigrationResourcesCreateAttributes struct {
	// The resources to add. From 1 to 1000 resources.
	Resources []SecurityMonitoringBatchMigrationResourceCreateItem `json:"resources"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringBatchMigrationResourcesCreateAttributes instantiates a new SecurityMonitoringBatchMigrationResourcesCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringBatchMigrationResourcesCreateAttributes(resources []SecurityMonitoringBatchMigrationResourceCreateItem) *SecurityMonitoringBatchMigrationResourcesCreateAttributes {
	this := SecurityMonitoringBatchMigrationResourcesCreateAttributes{}
	this.Resources = resources
	return &this
}

// NewSecurityMonitoringBatchMigrationResourcesCreateAttributesWithDefaults instantiates a new SecurityMonitoringBatchMigrationResourcesCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringBatchMigrationResourcesCreateAttributesWithDefaults() *SecurityMonitoringBatchMigrationResourcesCreateAttributes {
	this := SecurityMonitoringBatchMigrationResourcesCreateAttributes{}
	return &this
}

// GetResources returns the Resources field value.
func (o *SecurityMonitoringBatchMigrationResourcesCreateAttributes) GetResources() []SecurityMonitoringBatchMigrationResourceCreateItem {
	if o == nil {
		var ret []SecurityMonitoringBatchMigrationResourceCreateItem
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationResourcesCreateAttributes) GetResourcesOk() (*[]SecurityMonitoringBatchMigrationResourceCreateItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resources, true
}

// SetResources sets field value.
func (o *SecurityMonitoringBatchMigrationResourcesCreateAttributes) SetResources(v []SecurityMonitoringBatchMigrationResourceCreateItem) {
	o.Resources = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringBatchMigrationResourcesCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["resources"] = o.Resources

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringBatchMigrationResourcesCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Resources *[]SecurityMonitoringBatchMigrationResourceCreateItem `json:"resources"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Resources == nil {
		return fmt.Errorf("required field resources missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"resources"})
	} else {
		return err
	}
	o.Resources = *all.Resources

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
