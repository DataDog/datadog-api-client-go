// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GeneralInvestigationTrigger A trigger created from a general investigation request.
type GeneralInvestigationTrigger struct {
	// Attributes for a general investigation, not tied to a specific monitor alert.
	GeneralInvestigation GeneralInvestigationAttributes `json:"general_investigation"`
	// The type of general investigation trigger.
	Type GeneralInvestigationTriggerType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewGeneralInvestigationTrigger instantiates a new GeneralInvestigationTrigger object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGeneralInvestigationTrigger(generalInvestigation GeneralInvestigationAttributes, typeVar GeneralInvestigationTriggerType) *GeneralInvestigationTrigger {
	this := GeneralInvestigationTrigger{}
	this.GeneralInvestigation = generalInvestigation
	this.Type = typeVar
	return &this
}

// NewGeneralInvestigationTriggerWithDefaults instantiates a new GeneralInvestigationTrigger object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGeneralInvestigationTriggerWithDefaults() *GeneralInvestigationTrigger {
	this := GeneralInvestigationTrigger{}
	return &this
}

// GetGeneralInvestigation returns the GeneralInvestigation field value.
func (o *GeneralInvestigationTrigger) GetGeneralInvestigation() GeneralInvestigationAttributes {
	if o == nil {
		var ret GeneralInvestigationAttributes
		return ret
	}
	return o.GeneralInvestigation
}

// GetGeneralInvestigationOk returns a tuple with the GeneralInvestigation field value
// and a boolean to check if the value has been set.
func (o *GeneralInvestigationTrigger) GetGeneralInvestigationOk() (*GeneralInvestigationAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GeneralInvestigation, true
}

// SetGeneralInvestigation sets field value.
func (o *GeneralInvestigationTrigger) SetGeneralInvestigation(v GeneralInvestigationAttributes) {
	o.GeneralInvestigation = v
}

// GetType returns the Type field value.
func (o *GeneralInvestigationTrigger) GetType() GeneralInvestigationTriggerType {
	if o == nil {
		var ret GeneralInvestigationTriggerType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GeneralInvestigationTrigger) GetTypeOk() (*GeneralInvestigationTriggerType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *GeneralInvestigationTrigger) SetType(v GeneralInvestigationTriggerType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GeneralInvestigationTrigger) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["general_investigation"] = o.GeneralInvestigation
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GeneralInvestigationTrigger) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		GeneralInvestigation *GeneralInvestigationAttributes  `json:"general_investigation"`
		Type                 *GeneralInvestigationTriggerType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.GeneralInvestigation == nil {
		return fmt.Errorf("required field general_investigation missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.GeneralInvestigation = *all.GeneralInvestigation
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
