// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringBatchMigrationPublishAttributes Attributes of the batch migration publish request.
type SecurityMonitoringBatchMigrationPublishAttributes struct {
	// Unique UUIDs of the resources to publish. Up to 1000 IDs. When omitted or empty,
	// every `completed` resource of the batch migration is published.
	ResourceIds []uuid.UUID `json:"resourceIds,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringBatchMigrationPublishAttributes instantiates a new SecurityMonitoringBatchMigrationPublishAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringBatchMigrationPublishAttributes() *SecurityMonitoringBatchMigrationPublishAttributes {
	this := SecurityMonitoringBatchMigrationPublishAttributes{}
	return &this
}

// NewSecurityMonitoringBatchMigrationPublishAttributesWithDefaults instantiates a new SecurityMonitoringBatchMigrationPublishAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringBatchMigrationPublishAttributesWithDefaults() *SecurityMonitoringBatchMigrationPublishAttributes {
	this := SecurityMonitoringBatchMigrationPublishAttributes{}
	return &this
}

// GetResourceIds returns the ResourceIds field value if set, zero value otherwise.
func (o *SecurityMonitoringBatchMigrationPublishAttributes) GetResourceIds() []uuid.UUID {
	if o == nil || o.ResourceIds == nil {
		var ret []uuid.UUID
		return ret
	}
	return o.ResourceIds
}

// GetResourceIdsOk returns a tuple with the ResourceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationPublishAttributes) GetResourceIdsOk() (*[]uuid.UUID, bool) {
	if o == nil || o.ResourceIds == nil {
		return nil, false
	}
	return &o.ResourceIds, true
}

// HasResourceIds returns a boolean if a field has been set.
func (o *SecurityMonitoringBatchMigrationPublishAttributes) HasResourceIds() bool {
	return o != nil && o.ResourceIds != nil
}

// SetResourceIds gets a reference to the given []uuid.UUID and assigns it to the ResourceIds field.
func (o *SecurityMonitoringBatchMigrationPublishAttributes) SetResourceIds(v []uuid.UUID) {
	o.ResourceIds = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringBatchMigrationPublishAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ResourceIds != nil {
		toSerialize["resourceIds"] = o.ResourceIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringBatchMigrationPublishAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ResourceIds []uuid.UUID `json:"resourceIds,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"resourceIds"})
	} else {
		return err
	}
	o.ResourceIds = all.ResourceIds

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
