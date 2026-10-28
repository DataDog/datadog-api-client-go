// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleProposalResponseAttributes Attributes of the rule proposal.
type SecurityMonitoringRuleProposalResponseAttributes struct {
	// The reason no rule could be generated. Only present when the status is `failed`.
	Error *SecurityMonitoringRuleProposalError `json:"error,omitempty"`
	// The proposed Cloud SIEM detection rule in the same format as the create a detection
	// rule endpoint. It is not saved in Cloud SIEM. Only present when the status is
	// `completed`. Review the rule, then create it with the [create a
	// detection rule](https://docs.datadoghq.com/api/latest/security-monitoring/#create-a-detection-rule)
	// endpoint.
	Rule map[string]interface{} `json:"rule,omitempty"`
	// The status of the rule proposal. `pending` and `in_progress` mean the proposed rule is
	// still being generated. Poll the endpoint again later. `completed` and `failed` are final:
	// `completed` proposals contain the proposed rule in `rule`, and `failed` proposals
	// contain the reason in `error`.
	Status string `json:"status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringRuleProposalResponseAttributes instantiates a new SecurityMonitoringRuleProposalResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringRuleProposalResponseAttributes(status string) *SecurityMonitoringRuleProposalResponseAttributes {
	this := SecurityMonitoringRuleProposalResponseAttributes{}
	this.Status = status
	return &this
}

// NewSecurityMonitoringRuleProposalResponseAttributesWithDefaults instantiates a new SecurityMonitoringRuleProposalResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringRuleProposalResponseAttributesWithDefaults() *SecurityMonitoringRuleProposalResponseAttributes {
	this := SecurityMonitoringRuleProposalResponseAttributes{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleProposalResponseAttributes) GetError() SecurityMonitoringRuleProposalError {
	if o == nil || o.Error == nil {
		var ret SecurityMonitoringRuleProposalError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleProposalResponseAttributes) GetErrorOk() (*SecurityMonitoringRuleProposalError, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleProposalResponseAttributes) HasError() bool {
	return o != nil && o.Error != nil
}

// SetError gets a reference to the given SecurityMonitoringRuleProposalError and assigns it to the Error field.
func (o *SecurityMonitoringRuleProposalResponseAttributes) SetError(v SecurityMonitoringRuleProposalError) {
	o.Error = &v
}

// GetRule returns the Rule field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleProposalResponseAttributes) GetRule() map[string]interface{} {
	if o == nil || o.Rule == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Rule
}

// GetRuleOk returns a tuple with the Rule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleProposalResponseAttributes) GetRuleOk() (*map[string]interface{}, bool) {
	if o == nil || o.Rule == nil {
		return nil, false
	}
	return &o.Rule, true
}

// HasRule returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleProposalResponseAttributes) HasRule() bool {
	return o != nil && o.Rule != nil
}

// SetRule gets a reference to the given map[string]interface{} and assigns it to the Rule field.
func (o *SecurityMonitoringRuleProposalResponseAttributes) SetRule(v map[string]interface{}) {
	o.Rule = v
}

// GetStatus returns the Status field value.
func (o *SecurityMonitoringRuleProposalResponseAttributes) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleProposalResponseAttributes) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *SecurityMonitoringRuleProposalResponseAttributes) SetStatus(v string) {
	o.Status = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringRuleProposalResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.Rule != nil {
		toSerialize["rule"] = o.Rule
	}
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringRuleProposalResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Error  *SecurityMonitoringRuleProposalError `json:"error,omitempty"`
		Rule   map[string]interface{}               `json:"rule,omitempty"`
		Status *string                              `json:"status"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"error", "rule", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Error != nil && all.Error.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Error = all.Error
	o.Rule = all.Rule
	o.Status = *all.Status

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
