// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FrameworkControl Framework Control.
type FrameworkControl struct {
	// Control Name.
	Name string `json:"name"`
	// Rule IDs.
	RuleIds []string `json:"rule_ids"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFrameworkControl instantiates a new FrameworkControl object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFrameworkControl(name string, ruleIds []string) *FrameworkControl {
	this := FrameworkControl{}
	this.Name = name
	this.RuleIds = ruleIds
	return &this
}

// NewFrameworkControlWithDefaults instantiates a new FrameworkControl object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFrameworkControlWithDefaults() *FrameworkControl {
	this := FrameworkControl{}
	return &this
}

// GetName returns the Name field value.
func (o *FrameworkControl) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FrameworkControl) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *FrameworkControl) SetName(v string) {
	o.Name = v
}

// GetRuleIds returns the RuleIds field value.
func (o *FrameworkControl) GetRuleIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RuleIds
}

// GetRuleIdsOk returns a tuple with the RuleIds field value
// and a boolean to check if the value has been set.
func (o *FrameworkControl) GetRuleIdsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleIds, true
}

// SetRuleIds sets field value.
func (o *FrameworkControl) SetRuleIds(v []string) {
	o.RuleIds = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FrameworkControl) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["rule_ids"] = o.RuleIds

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FrameworkControl) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name    *string   `json:"name"`
		RuleIds *[]string `json:"rule_ids"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.RuleIds == nil {
		return fmt.Errorf("required field rule_ids missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "rule_ids"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.RuleIds = *all.RuleIds

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
