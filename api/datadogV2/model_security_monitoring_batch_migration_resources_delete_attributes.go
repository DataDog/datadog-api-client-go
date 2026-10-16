// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringBatchMigrationResourcesDeleteAttributes Attributes of the resources to delete from a batch migration.
type SecurityMonitoringBatchMigrationResourcesDeleteAttributes struct {
	// Unique UUIDs of batch migration resources. From 1 to 1000 IDs.
	ResourceIds []uuid.UUID `json:"resourceIds"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringBatchMigrationResourcesDeleteAttributes instantiates a new SecurityMonitoringBatchMigrationResourcesDeleteAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringBatchMigrationResourcesDeleteAttributes(resourceIds []uuid.UUID) *SecurityMonitoringBatchMigrationResourcesDeleteAttributes {
	this := SecurityMonitoringBatchMigrationResourcesDeleteAttributes{}
	this.ResourceIds = resourceIds
	return &this
}

// NewSecurityMonitoringBatchMigrationResourcesDeleteAttributesWithDefaults instantiates a new SecurityMonitoringBatchMigrationResourcesDeleteAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringBatchMigrationResourcesDeleteAttributesWithDefaults() *SecurityMonitoringBatchMigrationResourcesDeleteAttributes {
	this := SecurityMonitoringBatchMigrationResourcesDeleteAttributes{}
	return &this
}

// GetResourceIds returns the ResourceIds field value.
func (o *SecurityMonitoringBatchMigrationResourcesDeleteAttributes) GetResourceIds() []uuid.UUID {
	if o == nil {
		var ret []uuid.UUID
		return ret
	}
	return o.ResourceIds
}

// GetResourceIdsOk returns a tuple with the ResourceIds field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationResourcesDeleteAttributes) GetResourceIdsOk() (*[]uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceIds, true
}

// SetResourceIds sets field value.
func (o *SecurityMonitoringBatchMigrationResourcesDeleteAttributes) SetResourceIds(v []uuid.UUID) {
	o.ResourceIds = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringBatchMigrationResourcesDeleteAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["resourceIds"] = o.ResourceIds

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringBatchMigrationResourcesDeleteAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ResourceIds *[]uuid.UUID `json:"resourceIds"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ResourceIds == nil {
		return fmt.Errorf("required field resourceIds missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"resourceIds"})
	} else {
		return err
	}
	o.ResourceIds = *all.ResourceIds

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
