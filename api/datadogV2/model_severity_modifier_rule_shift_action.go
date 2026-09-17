// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SeverityModifierRuleShiftAction Shifts matched findings up or down by one severity rank.
type SeverityModifierRuleShiftAction struct {
	// An optional free-form explanation for the severity change.
	Description *string `json:"description,omitempty"`
	// The direction in which to shift the severity of matched findings by one rank.
	SeverityDelta SeverityModifierSeverityDelta `json:"severity_delta"`
	// The type of a severity modifier rule action that shifts the severity by one rank.
	Type SeverityModifierRuleShiftActionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSeverityModifierRuleShiftAction instantiates a new SeverityModifierRuleShiftAction object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSeverityModifierRuleShiftAction(severityDelta SeverityModifierSeverityDelta, typeVar SeverityModifierRuleShiftActionType) *SeverityModifierRuleShiftAction {
	this := SeverityModifierRuleShiftAction{}
	this.SeverityDelta = severityDelta
	this.Type = typeVar
	return &this
}

// NewSeverityModifierRuleShiftActionWithDefaults instantiates a new SeverityModifierRuleShiftAction object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSeverityModifierRuleShiftActionWithDefaults() *SeverityModifierRuleShiftAction {
	this := SeverityModifierRuleShiftAction{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SeverityModifierRuleShiftAction) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeverityModifierRuleShiftAction) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SeverityModifierRuleShiftAction) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SeverityModifierRuleShiftAction) SetDescription(v string) {
	o.Description = &v
}

// GetSeverityDelta returns the SeverityDelta field value.
func (o *SeverityModifierRuleShiftAction) GetSeverityDelta() SeverityModifierSeverityDelta {
	if o == nil {
		var ret SeverityModifierSeverityDelta
		return ret
	}
	return o.SeverityDelta
}

// GetSeverityDeltaOk returns a tuple with the SeverityDelta field value
// and a boolean to check if the value has been set.
func (o *SeverityModifierRuleShiftAction) GetSeverityDeltaOk() (*SeverityModifierSeverityDelta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeverityDelta, true
}

// SetSeverityDelta sets field value.
func (o *SeverityModifierRuleShiftAction) SetSeverityDelta(v SeverityModifierSeverityDelta) {
	o.SeverityDelta = v
}

// GetType returns the Type field value.
func (o *SeverityModifierRuleShiftAction) GetType() SeverityModifierRuleShiftActionType {
	if o == nil {
		var ret SeverityModifierRuleShiftActionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SeverityModifierRuleShiftAction) GetTypeOk() (*SeverityModifierRuleShiftActionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SeverityModifierRuleShiftAction) SetType(v SeverityModifierRuleShiftActionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SeverityModifierRuleShiftAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	toSerialize["severity_delta"] = o.SeverityDelta
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SeverityModifierRuleShiftAction) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description   *string                              `json:"description,omitempty"`
		SeverityDelta *SeverityModifierSeverityDelta       `json:"severity_delta"`
		Type          *SeverityModifierRuleShiftActionType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.SeverityDelta == nil {
		return fmt.Errorf("required field severity_delta missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "severity_delta", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = all.Description
	if !all.SeverityDelta.IsValid() {
		hasInvalidField = true
	} else {
		o.SeverityDelta = *all.SeverityDelta
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
