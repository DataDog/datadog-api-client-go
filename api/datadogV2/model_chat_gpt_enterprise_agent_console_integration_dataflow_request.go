// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChatGPTEnterpriseAgentConsoleIntegrationDataflowRequest Activity, token consumption, credits, and estimated usage costs across ChatGPT Chat, ChatGPT Work, and Codex, shown in Agent Console with breakdowns by product, model, and user where OpenAI provides them. Data already available from OpenAI appears within 1 hour of enabling this dataflow.
type ChatGPTEnterpriseAgentConsoleIntegrationDataflowRequest struct {
	// Whether Datadog collects this data. Defaults to `true`; set to `false` to stop collection.
	Enabled bool `json:"enabled"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewChatGPTEnterpriseAgentConsoleIntegrationDataflowRequest instantiates a new ChatGPTEnterpriseAgentConsoleIntegrationDataflowRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewChatGPTEnterpriseAgentConsoleIntegrationDataflowRequest(enabled bool) *ChatGPTEnterpriseAgentConsoleIntegrationDataflowRequest {
	this := ChatGPTEnterpriseAgentConsoleIntegrationDataflowRequest{}
	this.Enabled = enabled
	return &this
}

// NewChatGPTEnterpriseAgentConsoleIntegrationDataflowRequestWithDefaults instantiates a new ChatGPTEnterpriseAgentConsoleIntegrationDataflowRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewChatGPTEnterpriseAgentConsoleIntegrationDataflowRequestWithDefaults() *ChatGPTEnterpriseAgentConsoleIntegrationDataflowRequest {
	this := ChatGPTEnterpriseAgentConsoleIntegrationDataflowRequest{}
	return &this
}

// GetEnabled returns the Enabled field value.
func (o *ChatGPTEnterpriseAgentConsoleIntegrationDataflowRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ChatGPTEnterpriseAgentConsoleIntegrationDataflowRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *ChatGPTEnterpriseAgentConsoleIntegrationDataflowRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ChatGPTEnterpriseAgentConsoleIntegrationDataflowRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["enabled"] = o.Enabled
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ChatGPTEnterpriseAgentConsoleIntegrationDataflowRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled *bool `json:"enabled"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	o.Enabled = *all.Enabled

	return nil
}
