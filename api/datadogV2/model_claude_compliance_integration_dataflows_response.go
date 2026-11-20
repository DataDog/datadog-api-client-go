// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ClaudeComplianceIntegrationDataflowsResponse Data Datadog collects from Claude Compliance, keyed by dataflow ID.
type ClaudeComplianceIntegrationDataflowsResponse struct {
	// Audit activity logs from the Anthropic Compliance API, such as SSO sign-ins, API key lifecycle events, member and role changes, and Claude chat and file access. [Cloud SIEM](https://docs.datadoghq.com/security/cloud_siem/) uses these logs to detect security-sensitive activity. Requires an Anthropic Enterprise plan with the Compliance API enabled in the organization's Data and privacy settings. Without it, this dataflow collects no data. This dataflow is always on and cannot be turned off, so it appears in responses only.
	AnthropicComplianceLogs *ClaudeComplianceComplianceLogsIntegrationDataflowResponse `json:"anthropic-compliance-logs,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClaudeComplianceIntegrationDataflowsResponse instantiates a new ClaudeComplianceIntegrationDataflowsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClaudeComplianceIntegrationDataflowsResponse() *ClaudeComplianceIntegrationDataflowsResponse {
	this := ClaudeComplianceIntegrationDataflowsResponse{}
	return &this
}

// NewClaudeComplianceIntegrationDataflowsResponseWithDefaults instantiates a new ClaudeComplianceIntegrationDataflowsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClaudeComplianceIntegrationDataflowsResponseWithDefaults() *ClaudeComplianceIntegrationDataflowsResponse {
	this := ClaudeComplianceIntegrationDataflowsResponse{}
	return &this
}

// GetAnthropicComplianceLogs returns the AnthropicComplianceLogs field value if set, zero value otherwise.
func (o *ClaudeComplianceIntegrationDataflowsResponse) GetAnthropicComplianceLogs() ClaudeComplianceComplianceLogsIntegrationDataflowResponse {
	if o == nil || o.AnthropicComplianceLogs == nil {
		var ret ClaudeComplianceComplianceLogsIntegrationDataflowResponse
		return ret
	}
	return *o.AnthropicComplianceLogs
}

// GetAnthropicComplianceLogsOk returns a tuple with the AnthropicComplianceLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaudeComplianceIntegrationDataflowsResponse) GetAnthropicComplianceLogsOk() (*ClaudeComplianceComplianceLogsIntegrationDataflowResponse, bool) {
	if o == nil || o.AnthropicComplianceLogs == nil {
		return nil, false
	}
	return o.AnthropicComplianceLogs, true
}

// HasAnthropicComplianceLogs returns a boolean if a field has been set.
func (o *ClaudeComplianceIntegrationDataflowsResponse) HasAnthropicComplianceLogs() bool {
	return o != nil && o.AnthropicComplianceLogs != nil
}

// SetAnthropicComplianceLogs gets a reference to the given ClaudeComplianceComplianceLogsIntegrationDataflowResponse and assigns it to the AnthropicComplianceLogs field.
func (o *ClaudeComplianceIntegrationDataflowsResponse) SetAnthropicComplianceLogs(v ClaudeComplianceComplianceLogsIntegrationDataflowResponse) {
	o.AnthropicComplianceLogs = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClaudeComplianceIntegrationDataflowsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AnthropicComplianceLogs != nil {
		toSerialize["anthropic-compliance-logs"] = o.AnthropicComplianceLogs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClaudeComplianceIntegrationDataflowsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AnthropicComplianceLogs *ClaudeComplianceComplianceLogsIntegrationDataflowResponse `json:"anthropic-compliance-logs,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"anthropic-compliance-logs"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AnthropicComplianceLogs != nil && all.AnthropicComplianceLogs.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AnthropicComplianceLogs = all.AnthropicComplianceLogs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
