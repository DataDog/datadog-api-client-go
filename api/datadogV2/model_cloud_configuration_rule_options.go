// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// CloudConfigurationRuleOptions Options on cloud configuration rules.
type CloudConfigurationRuleOptions struct {
	// Rego rule details.
	RegoRule CloudConfigurationRegoRule `json:"regoRule"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewCloudConfigurationRuleOptions instantiates a new CloudConfigurationRuleOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCloudConfigurationRuleOptions(regoRule CloudConfigurationRegoRule) *CloudConfigurationRuleOptions {
	this := CloudConfigurationRuleOptions{}
	this.RegoRule = regoRule
	return &this
}

// NewCloudConfigurationRuleOptionsWithDefaults instantiates a new CloudConfigurationRuleOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCloudConfigurationRuleOptionsWithDefaults() *CloudConfigurationRuleOptions {
	this := CloudConfigurationRuleOptions{}
	return &this
}

// GetRegoRule returns the RegoRule field value.
func (o *CloudConfigurationRuleOptions) GetRegoRule() CloudConfigurationRegoRule {
	if o == nil {
		var ret CloudConfigurationRegoRule
		return ret
	}
	return o.RegoRule
}

// GetRegoRuleOk returns a tuple with the RegoRule field value
// and a boolean to check if the value has been set.
func (o *CloudConfigurationRuleOptions) GetRegoRuleOk() (*CloudConfigurationRegoRule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegoRule, true
}

// SetRegoRule sets field value.
func (o *CloudConfigurationRuleOptions) SetRegoRule(v CloudConfigurationRegoRule) {
	o.RegoRule = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CloudConfigurationRuleOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["regoRule"] = o.RegoRule

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CloudConfigurationRuleOptions) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		RegoRule *CloudConfigurationRegoRule `json:"regoRule"`
	}{}
	all := struct {
		RegoRule CloudConfigurationRegoRule `json:"regoRule"`
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
