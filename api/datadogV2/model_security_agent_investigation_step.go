// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityAgentInvestigationStep A step of the investigation.
type SecurityAgentInvestigationStep struct {
	// Name of the step, usually the question it answers.
	Name string `json:"name"`
	// Summary of what the step found. It can be empty.
	Summary string `json:"summary"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityAgentInvestigationStep instantiates a new SecurityAgentInvestigationStep object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityAgentInvestigationStep(name string, summary string) *SecurityAgentInvestigationStep {
	this := SecurityAgentInvestigationStep{}
	this.Name = name
	this.Summary = summary
	return &this
}

// NewSecurityAgentInvestigationStepWithDefaults instantiates a new SecurityAgentInvestigationStep object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityAgentInvestigationStepWithDefaults() *SecurityAgentInvestigationStep {
	this := SecurityAgentInvestigationStep{}
	return &this
}

// GetName returns the Name field value.
func (o *SecurityAgentInvestigationStep) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationStep) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SecurityAgentInvestigationStep) SetName(v string) {
	o.Name = v
}

// GetSummary returns the Summary field value.
func (o *SecurityAgentInvestigationStep) GetSummary() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationStep) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value.
func (o *SecurityAgentInvestigationStep) SetSummary(v string) {
	o.Summary = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityAgentInvestigationStep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["summary"] = o.Summary

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityAgentInvestigationStep) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name    *string `json:"name"`
		Summary *string `json:"summary"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Summary == nil {
		return fmt.Errorf("required field summary missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "summary"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Summary = *all.Summary

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
