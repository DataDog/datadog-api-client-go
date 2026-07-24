// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleBulkDeleteResponseAttributes Attributes for the bulk delete response.
type SecurityMonitoringRuleBulkDeleteResponseAttributes struct {
	// List of successfully deleted rule IDs.
	DeletedRules []string `json:"deletedRules,omitempty"`
	// List of rule IDs that could not be deleted.
	FailedRules []string `json:"failedRules,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringRuleBulkDeleteResponseAttributes instantiates a new SecurityMonitoringRuleBulkDeleteResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringRuleBulkDeleteResponseAttributes() *SecurityMonitoringRuleBulkDeleteResponseAttributes {
	this := SecurityMonitoringRuleBulkDeleteResponseAttributes{}
	return &this
}

// NewSecurityMonitoringRuleBulkDeleteResponseAttributesWithDefaults instantiates a new SecurityMonitoringRuleBulkDeleteResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringRuleBulkDeleteResponseAttributesWithDefaults() *SecurityMonitoringRuleBulkDeleteResponseAttributes {
	this := SecurityMonitoringRuleBulkDeleteResponseAttributes{}
	return &this
}

// GetDeletedRules returns the DeletedRules field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleBulkDeleteResponseAttributes) GetDeletedRules() []string {
	if o == nil || o.DeletedRules == nil {
		var ret []string
		return ret
	}
	return o.DeletedRules
}

// GetDeletedRulesOk returns a tuple with the DeletedRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleBulkDeleteResponseAttributes) GetDeletedRulesOk() (*[]string, bool) {
	if o == nil || o.DeletedRules == nil {
		return nil, false
	}
	return &o.DeletedRules, true
}

// HasDeletedRules returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleBulkDeleteResponseAttributes) HasDeletedRules() bool {
	return o != nil && o.DeletedRules != nil
}

// SetDeletedRules gets a reference to the given []string and assigns it to the DeletedRules field.
func (o *SecurityMonitoringRuleBulkDeleteResponseAttributes) SetDeletedRules(v []string) {
	o.DeletedRules = v
}

// GetFailedRules returns the FailedRules field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleBulkDeleteResponseAttributes) GetFailedRules() []string {
	if o == nil || o.FailedRules == nil {
		var ret []string
		return ret
	}
	return o.FailedRules
}

// GetFailedRulesOk returns a tuple with the FailedRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleBulkDeleteResponseAttributes) GetFailedRulesOk() (*[]string, bool) {
	if o == nil || o.FailedRules == nil {
		return nil, false
	}
	return &o.FailedRules, true
}

// HasFailedRules returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleBulkDeleteResponseAttributes) HasFailedRules() bool {
	return o != nil && o.FailedRules != nil
}

// SetFailedRules gets a reference to the given []string and assigns it to the FailedRules field.
func (o *SecurityMonitoringRuleBulkDeleteResponseAttributes) SetFailedRules(v []string) {
	o.FailedRules = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringRuleBulkDeleteResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DeletedRules != nil {
		toSerialize["deletedRules"] = o.DeletedRules
	}
	if o.FailedRules != nil {
		toSerialize["failedRules"] = o.FailedRules
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringRuleBulkDeleteResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DeletedRules []string `json:"deletedRules,omitempty"`
		FailedRules  []string `json:"failedRules,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"deletedRules", "failedRules"})
	} else {
		return err
	}
	o.DeletedRules = all.DeletedRules
	o.FailedRules = all.FailedRules

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
