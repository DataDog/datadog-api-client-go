// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleProposalError The reason no rule could be generated. Only present when the status is `failed`.
type SecurityMonitoringRuleProposalError struct {
	// The kind of failure.
	// - `translation_error`: no rule could be generated from the input. Reword the input
	//   or add details, then start a new rule proposal.
	// - `internal_error`: the proposal failed because of an internal error. Start a new
	//   rule proposal to try again.
	Kind string `json:"kind"`
	// A human-readable explanation of why no rule could be generated, when available.
	Message *string `json:"message,omitempty"`
	// The type of failure within its kind, such as `NoRuleProduced`.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringRuleProposalError instantiates a new SecurityMonitoringRuleProposalError object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringRuleProposalError(kind string) *SecurityMonitoringRuleProposalError {
	this := SecurityMonitoringRuleProposalError{}
	this.Kind = kind
	return &this
}

// NewSecurityMonitoringRuleProposalErrorWithDefaults instantiates a new SecurityMonitoringRuleProposalError object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringRuleProposalErrorWithDefaults() *SecurityMonitoringRuleProposalError {
	this := SecurityMonitoringRuleProposalError{}
	return &this
}

// GetKind returns the Kind field value.
func (o *SecurityMonitoringRuleProposalError) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleProposalError) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value.
func (o *SecurityMonitoringRuleProposalError) SetKind(v string) {
	o.Kind = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleProposalError) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleProposalError) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleProposalError) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *SecurityMonitoringRuleProposalError) SetMessage(v string) {
	o.Message = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleProposalError) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleProposalError) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleProposalError) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SecurityMonitoringRuleProposalError) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringRuleProposalError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["kind"] = o.Kind
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringRuleProposalError) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Kind    *string `json:"kind"`
		Message *string `json:"message,omitempty"`
		Type    *string `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Kind == nil {
		return fmt.Errorf("required field kind missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"kind", "message", "type"})
	} else {
		return err
	}
	o.Kind = *all.Kind
	o.Message = all.Message
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
