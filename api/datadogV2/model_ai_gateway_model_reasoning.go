// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelReasoning Supported reasoning controls. Absence means controls are unknown or unavailable.
type AiGatewayModelReasoning struct {
	//
	Controls []AiGatewayModelReasoningControl `json:"controls"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayModelReasoning instantiates a new AiGatewayModelReasoning object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayModelReasoning(controls []AiGatewayModelReasoningControl) *AiGatewayModelReasoning {
	this := AiGatewayModelReasoning{}
	this.Controls = controls
	return &this
}

// NewAiGatewayModelReasoningWithDefaults instantiates a new AiGatewayModelReasoning object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayModelReasoningWithDefaults() *AiGatewayModelReasoning {
	this := AiGatewayModelReasoning{}
	return &this
}

// GetControls returns the Controls field value.
func (o *AiGatewayModelReasoning) GetControls() []AiGatewayModelReasoningControl {
	if o == nil {
		var ret []AiGatewayModelReasoningControl
		return ret
	}
	return o.Controls
}

// GetControlsOk returns a tuple with the Controls field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelReasoning) GetControlsOk() (*[]AiGatewayModelReasoningControl, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Controls, true
}

// SetControls sets field value.
func (o *AiGatewayModelReasoning) SetControls(v []AiGatewayModelReasoningControl) {
	o.Controls = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayModelReasoning) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["controls"] = o.Controls

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayModelReasoning) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Controls *[]AiGatewayModelReasoningControl `json:"controls"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Controls == nil {
		return fmt.Errorf("required field controls missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"controls"})
	} else {
		return err
	}
	o.Controls = *all.Controls

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
