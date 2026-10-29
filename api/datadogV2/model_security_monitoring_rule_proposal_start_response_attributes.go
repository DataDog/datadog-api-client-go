// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleProposalStartResponseAttributes Attributes of the started rule proposal.
type SecurityMonitoringRuleProposalStartResponseAttributes struct {
	// The status of the rule proposal, one of `pending`, `in_progress`, `completed`, or
	// `failed`. For a newly started rule proposal, the status is `pending`.
	Status string `json:"status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringRuleProposalStartResponseAttributes instantiates a new SecurityMonitoringRuleProposalStartResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringRuleProposalStartResponseAttributes(status string) *SecurityMonitoringRuleProposalStartResponseAttributes {
	this := SecurityMonitoringRuleProposalStartResponseAttributes{}
	this.Status = status
	return &this
}

// NewSecurityMonitoringRuleProposalStartResponseAttributesWithDefaults instantiates a new SecurityMonitoringRuleProposalStartResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringRuleProposalStartResponseAttributesWithDefaults() *SecurityMonitoringRuleProposalStartResponseAttributes {
	this := SecurityMonitoringRuleProposalStartResponseAttributes{}
	return &this
}

// GetStatus returns the Status field value.
func (o *SecurityMonitoringRuleProposalStartResponseAttributes) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleProposalStartResponseAttributes) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *SecurityMonitoringRuleProposalStartResponseAttributes) SetStatus(v string) {
	o.Status = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringRuleProposalStartResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringRuleProposalStartResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Status *string `json:"status"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"status"})
	} else {
		return err
	}
	o.Status = *all.Status

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
