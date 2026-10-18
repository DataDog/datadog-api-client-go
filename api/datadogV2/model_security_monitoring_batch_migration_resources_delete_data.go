// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringBatchMigrationResourcesDeleteData Data of the resources to delete from a batch migration.
type SecurityMonitoringBatchMigrationResourcesDeleteData struct {
	// Attributes of the resources to delete from a batch migration.
	Attributes SecurityMonitoringBatchMigrationResourcesDeleteAttributes `json:"attributes"`
	// The type of resource for a request on a set of batch migration resources.
	Type SecurityMonitoringBatchMigrationResourcesType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringBatchMigrationResourcesDeleteData instantiates a new SecurityMonitoringBatchMigrationResourcesDeleteData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringBatchMigrationResourcesDeleteData(attributes SecurityMonitoringBatchMigrationResourcesDeleteAttributes, typeVar SecurityMonitoringBatchMigrationResourcesType) *SecurityMonitoringBatchMigrationResourcesDeleteData {
	this := SecurityMonitoringBatchMigrationResourcesDeleteData{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewSecurityMonitoringBatchMigrationResourcesDeleteDataWithDefaults instantiates a new SecurityMonitoringBatchMigrationResourcesDeleteData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringBatchMigrationResourcesDeleteDataWithDefaults() *SecurityMonitoringBatchMigrationResourcesDeleteData {
	this := SecurityMonitoringBatchMigrationResourcesDeleteData{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *SecurityMonitoringBatchMigrationResourcesDeleteData) GetAttributes() SecurityMonitoringBatchMigrationResourcesDeleteAttributes {
	if o == nil {
		var ret SecurityMonitoringBatchMigrationResourcesDeleteAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationResourcesDeleteData) GetAttributesOk() (*SecurityMonitoringBatchMigrationResourcesDeleteAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *SecurityMonitoringBatchMigrationResourcesDeleteData) SetAttributes(v SecurityMonitoringBatchMigrationResourcesDeleteAttributes) {
	o.Attributes = v
}

// GetType returns the Type field value.
func (o *SecurityMonitoringBatchMigrationResourcesDeleteData) GetType() SecurityMonitoringBatchMigrationResourcesType {
	if o == nil {
		var ret SecurityMonitoringBatchMigrationResourcesType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationResourcesDeleteData) GetTypeOk() (*SecurityMonitoringBatchMigrationResourcesType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SecurityMonitoringBatchMigrationResourcesDeleteData) SetType(v SecurityMonitoringBatchMigrationResourcesType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringBatchMigrationResourcesDeleteData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringBatchMigrationResourcesDeleteData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *SecurityMonitoringBatchMigrationResourcesDeleteAttributes `json:"attributes"`
		Type       *SecurityMonitoringBatchMigrationResourcesType             `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
