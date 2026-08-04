// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SeverityModifierRuleSetAction Sets matched findings to a fixed severity.
type SeverityModifierRuleSetAction struct {
	// An optional free-form explanation for the severity change.
	Description *string `json:"description,omitempty"`
	// The severity to assign to matched findings. `info_none` is not supported for the `iac_misconfiguration`, `runtime_code_vulnerability`, `secret`, or `static_code_vulnerability` finding types.
	Severity SeverityModifierSeverity `json:"severity"`
	// The type of a severity modifier rule action that sets a fixed severity.
	Type SeverityModifierRuleSetActionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSeverityModifierRuleSetAction instantiates a new SeverityModifierRuleSetAction object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSeverityModifierRuleSetAction(severity SeverityModifierSeverity, typeVar SeverityModifierRuleSetActionType) *SeverityModifierRuleSetAction {
	this := SeverityModifierRuleSetAction{}
	this.Severity = severity
	this.Type = typeVar
	return &this
}

// NewSeverityModifierRuleSetActionWithDefaults instantiates a new SeverityModifierRuleSetAction object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSeverityModifierRuleSetActionWithDefaults() *SeverityModifierRuleSetAction {
	this := SeverityModifierRuleSetAction{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SeverityModifierRuleSetAction) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeverityModifierRuleSetAction) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SeverityModifierRuleSetAction) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SeverityModifierRuleSetAction) SetDescription(v string) {
	o.Description = &v
}

// GetSeverity returns the Severity field value.
func (o *SeverityModifierRuleSetAction) GetSeverity() SeverityModifierSeverity {
	if o == nil {
		var ret SeverityModifierSeverity
		return ret
	}
	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *SeverityModifierRuleSetAction) GetSeverityOk() (*SeverityModifierSeverity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value.
func (o *SeverityModifierRuleSetAction) SetSeverity(v SeverityModifierSeverity) {
	o.Severity = v
}

// GetType returns the Type field value.
func (o *SeverityModifierRuleSetAction) GetType() SeverityModifierRuleSetActionType {
	if o == nil {
		var ret SeverityModifierRuleSetActionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SeverityModifierRuleSetAction) GetTypeOk() (*SeverityModifierRuleSetActionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SeverityModifierRuleSetAction) SetType(v SeverityModifierRuleSetActionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SeverityModifierRuleSetAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	toSerialize["severity"] = o.Severity
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SeverityModifierRuleSetAction) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string                            `json:"description,omitempty"`
		Severity    *SeverityModifierSeverity          `json:"severity"`
		Type        *SeverityModifierRuleSetActionType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Severity == nil {
		return fmt.Errorf("required field severity missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "severity", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = all.Description
	if !all.Severity.IsValid() {
		hasInvalidField = true
	} else {
		o.Severity = *all.Severity
	}
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
