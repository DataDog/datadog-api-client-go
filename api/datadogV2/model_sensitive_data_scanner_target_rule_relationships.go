// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SensitiveDataScannerTargetRuleRelationships Relationships of a rule of a scanning group of a target.
type SensitiveDataScannerTargetRuleRelationships struct {
	// A scanning group data.
	Group *SensitiveDataScannerGroupData `json:"group,omitempty"`
	// The standard rule a rule is based on. Not included for custom rules.
	StandardRule *SensitiveDataScannerStandardRule `json:"standard_rule,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSensitiveDataScannerTargetRuleRelationships instantiates a new SensitiveDataScannerTargetRuleRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSensitiveDataScannerTargetRuleRelationships() *SensitiveDataScannerTargetRuleRelationships {
	this := SensitiveDataScannerTargetRuleRelationships{}
	return &this
}

// NewSensitiveDataScannerTargetRuleRelationshipsWithDefaults instantiates a new SensitiveDataScannerTargetRuleRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSensitiveDataScannerTargetRuleRelationshipsWithDefaults() *SensitiveDataScannerTargetRuleRelationships {
	this := SensitiveDataScannerTargetRuleRelationships{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *SensitiveDataScannerTargetRuleRelationships) GetGroup() SensitiveDataScannerGroupData {
	if o == nil || o.Group == nil {
		var ret SensitiveDataScannerGroupData
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerTargetRuleRelationships) GetGroupOk() (*SensitiveDataScannerGroupData, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *SensitiveDataScannerTargetRuleRelationships) HasGroup() bool {
	return o != nil && o.Group != nil
}

// SetGroup gets a reference to the given SensitiveDataScannerGroupData and assigns it to the Group field.
func (o *SensitiveDataScannerTargetRuleRelationships) SetGroup(v SensitiveDataScannerGroupData) {
	o.Group = &v
}

// GetStandardRule returns the StandardRule field value if set, zero value otherwise.
func (o *SensitiveDataScannerTargetRuleRelationships) GetStandardRule() SensitiveDataScannerStandardRule {
	if o == nil || o.StandardRule == nil {
		var ret SensitiveDataScannerStandardRule
		return ret
	}
	return *o.StandardRule
}

// GetStandardRuleOk returns a tuple with the StandardRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerTargetRuleRelationships) GetStandardRuleOk() (*SensitiveDataScannerStandardRule, bool) {
	if o == nil || o.StandardRule == nil {
		return nil, false
	}
	return o.StandardRule, true
}

// HasStandardRule returns a boolean if a field has been set.
func (o *SensitiveDataScannerTargetRuleRelationships) HasStandardRule() bool {
	return o != nil && o.StandardRule != nil
}

// SetStandardRule gets a reference to the given SensitiveDataScannerStandardRule and assigns it to the StandardRule field.
func (o *SensitiveDataScannerTargetRuleRelationships) SetStandardRule(v SensitiveDataScannerStandardRule) {
	o.StandardRule = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SensitiveDataScannerTargetRuleRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Group != nil {
		toSerialize["group"] = o.Group
	}
	if o.StandardRule != nil {
		toSerialize["standard_rule"] = o.StandardRule
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SensitiveDataScannerTargetRuleRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Group        *SensitiveDataScannerGroupData    `json:"group,omitempty"`
		StandardRule *SensitiveDataScannerStandardRule `json:"standard_rule,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"group", "standard_rule"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Group != nil && all.Group.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Group = all.Group
	if all.StandardRule != nil && all.StandardRule.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.StandardRule = all.StandardRule

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
