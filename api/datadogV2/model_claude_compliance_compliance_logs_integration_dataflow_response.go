// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ClaudeComplianceComplianceLogsIntegrationDataflowResponse Audit activity logs from the Anthropic Compliance API, such as SSO sign-ins, API key lifecycle events, member and role changes, and Claude chat and file access. [Cloud SIEM](https://docs.datadoghq.com/security/cloud_siem/) uses these logs to detect security-sensitive activity. Requires an Anthropic Enterprise plan with the Compliance API enabled in the organization's Data and privacy settings. Without it, this dataflow collects no data. This dataflow is always on and cannot be turned off, so it appears in responses only.
type ClaudeComplianceComplianceLogsIntegrationDataflowResponse struct {
	// Whether Datadog collects this data. Always `true`, because this collection cannot be turned off.
	Enabled *bool `json:"enabled,omitempty"`
	// Read-only collection status of a dataflow.
	Status *IntegrationAccountDataflowStatus `json:"status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClaudeComplianceComplianceLogsIntegrationDataflowResponse instantiates a new ClaudeComplianceComplianceLogsIntegrationDataflowResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClaudeComplianceComplianceLogsIntegrationDataflowResponse() *ClaudeComplianceComplianceLogsIntegrationDataflowResponse {
	this := ClaudeComplianceComplianceLogsIntegrationDataflowResponse{}
	return &this
}

// NewClaudeComplianceComplianceLogsIntegrationDataflowResponseWithDefaults instantiates a new ClaudeComplianceComplianceLogsIntegrationDataflowResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClaudeComplianceComplianceLogsIntegrationDataflowResponseWithDefaults() *ClaudeComplianceComplianceLogsIntegrationDataflowResponse {
	this := ClaudeComplianceComplianceLogsIntegrationDataflowResponse{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ClaudeComplianceComplianceLogsIntegrationDataflowResponse) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaudeComplianceComplianceLogsIntegrationDataflowResponse) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ClaudeComplianceComplianceLogsIntegrationDataflowResponse) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ClaudeComplianceComplianceLogsIntegrationDataflowResponse) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ClaudeComplianceComplianceLogsIntegrationDataflowResponse) GetStatus() IntegrationAccountDataflowStatus {
	if o == nil || o.Status == nil {
		var ret IntegrationAccountDataflowStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaudeComplianceComplianceLogsIntegrationDataflowResponse) GetStatusOk() (*IntegrationAccountDataflowStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ClaudeComplianceComplianceLogsIntegrationDataflowResponse) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given IntegrationAccountDataflowStatus and assigns it to the Status field.
func (o *ClaudeComplianceComplianceLogsIntegrationDataflowResponse) SetStatus(v IntegrationAccountDataflowStatus) {
	o.Status = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClaudeComplianceComplianceLogsIntegrationDataflowResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClaudeComplianceComplianceLogsIntegrationDataflowResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled *bool                             `json:"enabled,omitempty"`
		Status  *IntegrationAccountDataflowStatus `json:"status,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enabled", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Enabled = all.Enabled
	if all.Status != nil && all.Status.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Status = all.Status

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
