// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// CloudConfigurationComplianceRuleOptions Options for cloud_configuration rules.
type CloudConfigurationComplianceRuleOptions struct {
	// Whether the rule is a complex one.
	// Must be set to true if `regoRule.resourceTypes` contains more than one item. Defaults to false.
	//
	ComplexRule *bool `json:"complexRule,omitempty"`
	// Rule details.
	RegoRule CloudConfigurationRegoRule `json:"regoRule"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewCloudConfigurationComplianceRuleOptions instantiates a new CloudConfigurationComplianceRuleOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCloudConfigurationComplianceRuleOptions(regoRule CloudConfigurationRegoRule) *CloudConfigurationComplianceRuleOptions {
	this := CloudConfigurationComplianceRuleOptions{}
	this.RegoRule = regoRule
	return &this
}

// NewCloudConfigurationComplianceRuleOptionsWithDefaults instantiates a new CloudConfigurationComplianceRuleOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCloudConfigurationComplianceRuleOptionsWithDefaults() *CloudConfigurationComplianceRuleOptions {
	this := CloudConfigurationComplianceRuleOptions{}
	return &this
}

// GetComplexRule returns the ComplexRule field value if set, zero value otherwise.
func (o *CloudConfigurationComplianceRuleOptions) GetComplexRule() bool {
	if o == nil || o.ComplexRule == nil {
		var ret bool
		return ret
	}
	return *o.ComplexRule
}

// GetComplexRuleOk returns a tuple with the ComplexRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudConfigurationComplianceRuleOptions) GetComplexRuleOk() (*bool, bool) {
	if o == nil || o.ComplexRule == nil {
		return nil, false
	}
	return o.ComplexRule, true
}

// HasComplexRule returns a boolean if a field has been set.
func (o *CloudConfigurationComplianceRuleOptions) HasComplexRule() bool {
	return o != nil && o.ComplexRule != nil
}

// SetComplexRule gets a reference to the given bool and assigns it to the ComplexRule field.
func (o *CloudConfigurationComplianceRuleOptions) SetComplexRule(v bool) {
	o.ComplexRule = &v
}

// GetRegoRule returns the RegoRule field value.
func (o *CloudConfigurationComplianceRuleOptions) GetRegoRule() CloudConfigurationRegoRule {
	if o == nil {
		var ret CloudConfigurationRegoRule
		return ret
	}
	return o.RegoRule
}

// GetRegoRuleOk returns a tuple with the RegoRule field value
// and a boolean to check if the value has been set.
func (o *CloudConfigurationComplianceRuleOptions) GetRegoRuleOk() (*CloudConfigurationRegoRule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegoRule, true
}

// SetRegoRule sets field value.
func (o *CloudConfigurationComplianceRuleOptions) SetRegoRule(v CloudConfigurationRegoRule) {
	o.RegoRule = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CloudConfigurationComplianceRuleOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.ComplexRule != nil {
		toSerialize["complexRule"] = o.ComplexRule
	}
	toSerialize["regoRule"] = o.RegoRule

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CloudConfigurationComplianceRuleOptions) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		RegoRule *CloudConfigurationRegoRule `json:"regoRule"`
	}{}
	all := struct {
		ComplexRule *bool                      `json:"complexRule,omitempty"`
		RegoRule    CloudConfigurationRegoRule `json:"regoRule"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.RegoRule == nil {
		return fmt.Errorf("required field regoRule missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.ComplexRule = all.ComplexRule
	if all.RegoRule.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.RegoRule = all.RegoRule
	return nil
}
