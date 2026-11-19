// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityAgentInvestigationCreateAttributes Attributes of the investigation to start.
type SecurityAgentInvestigationCreateAttributes struct {
	// Start a new investigation even if the signal already has one. If the signal's investigation is still
	// running, the response has that investigation instead. A rerun counts toward the monthly investigation
	// limit like any investigation.
	Rerun *bool `json:"rerun,omitempty"`
	// ID of the signal to investigate, in its V1 form, which starts with `AQ`, or in its longer V3 form, which
	// starts with `Aw`. Investigations carry the V1 form, so their `signal_id` can differ from the ID you pass.
	SignalId string `json:"signal_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityAgentInvestigationCreateAttributes instantiates a new SecurityAgentInvestigationCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityAgentInvestigationCreateAttributes(signalId string) *SecurityAgentInvestigationCreateAttributes {
	this := SecurityAgentInvestigationCreateAttributes{}
	var rerun bool = false
	this.Rerun = &rerun
	this.SignalId = signalId
	return &this
}

// NewSecurityAgentInvestigationCreateAttributesWithDefaults instantiates a new SecurityAgentInvestigationCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityAgentInvestigationCreateAttributesWithDefaults() *SecurityAgentInvestigationCreateAttributes {
	this := SecurityAgentInvestigationCreateAttributes{}
	var rerun bool = false
	this.Rerun = &rerun
	return &this
}

// GetRerun returns the Rerun field value if set, zero value otherwise.
func (o *SecurityAgentInvestigationCreateAttributes) GetRerun() bool {
	if o == nil || o.Rerun == nil {
		var ret bool
		return ret
	}
	return *o.Rerun
}

// GetRerunOk returns a tuple with the Rerun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationCreateAttributes) GetRerunOk() (*bool, bool) {
	if o == nil || o.Rerun == nil {
		return nil, false
	}
	return o.Rerun, true
}

// HasRerun returns a boolean if a field has been set.
func (o *SecurityAgentInvestigationCreateAttributes) HasRerun() bool {
	return o != nil && o.Rerun != nil
}

// SetRerun gets a reference to the given bool and assigns it to the Rerun field.
func (o *SecurityAgentInvestigationCreateAttributes) SetRerun(v bool) {
	o.Rerun = &v
}

// GetSignalId returns the SignalId field value.
func (o *SecurityAgentInvestigationCreateAttributes) GetSignalId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SignalId
}

// GetSignalIdOk returns a tuple with the SignalId field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationCreateAttributes) GetSignalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignalId, true
}

// SetSignalId sets field value.
func (o *SecurityAgentInvestigationCreateAttributes) SetSignalId(v string) {
	o.SignalId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityAgentInvestigationCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Rerun != nil {
		toSerialize["rerun"] = o.Rerun
	}
	toSerialize["signal_id"] = o.SignalId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityAgentInvestigationCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Rerun    *bool   `json:"rerun,omitempty"`
		SignalId *string `json:"signal_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.SignalId == nil {
		return fmt.Errorf("required field signal_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"rerun", "signal_id"})
	} else {
		return err
	}
	o.Rerun = all.Rerun
	o.SignalId = *all.SignalId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
