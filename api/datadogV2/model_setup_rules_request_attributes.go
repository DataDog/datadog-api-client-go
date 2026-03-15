// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SetupRulesRequestAttributes Attributes for setting up rules.
type SetupRulesRequestAttributes struct {
	// Array of default rule IDs to disable.
	DisabledDefaultRules []string `json:"disabled_default_rules,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSetupRulesRequestAttributes instantiates a new SetupRulesRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSetupRulesRequestAttributes() *SetupRulesRequestAttributes {
	this := SetupRulesRequestAttributes{}
	return &this
}

// NewSetupRulesRequestAttributesWithDefaults instantiates a new SetupRulesRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSetupRulesRequestAttributesWithDefaults() *SetupRulesRequestAttributes {
	this := SetupRulesRequestAttributes{}
	return &this
}

// GetDisabledDefaultRules returns the DisabledDefaultRules field value if set, zero value otherwise.
func (o *SetupRulesRequestAttributes) GetDisabledDefaultRules() []string {
	if o == nil || o.DisabledDefaultRules == nil {
		var ret []string
		return ret
	}
	return o.DisabledDefaultRules
}

// GetDisabledDefaultRulesOk returns a tuple with the DisabledDefaultRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetupRulesRequestAttributes) GetDisabledDefaultRulesOk() (*[]string, bool) {
	if o == nil || o.DisabledDefaultRules == nil {
		return nil, false
	}
	return &o.DisabledDefaultRules, true
}

// HasDisabledDefaultRules returns a boolean if a field has been set.
func (o *SetupRulesRequestAttributes) HasDisabledDefaultRules() bool {
	return o != nil && o.DisabledDefaultRules != nil
}

// SetDisabledDefaultRules gets a reference to the given []string and assigns it to the DisabledDefaultRules field.
func (o *SetupRulesRequestAttributes) SetDisabledDefaultRules(v []string) {
	o.DisabledDefaultRules = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SetupRulesRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DisabledDefaultRules != nil {
		toSerialize["disabled_default_rules"] = o.DisabledDefaultRules
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SetupRulesRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DisabledDefaultRules []string `json:"disabled_default_rules,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"disabled_default_rules"})
	} else {
		return err
	}
	o.DisabledDefaultRules = all.DisabledDefaultRules

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
