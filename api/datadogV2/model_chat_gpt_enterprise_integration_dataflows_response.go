// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChatGPTEnterpriseIntegrationDataflowsResponse Data Datadog collects from ChatGPT Enterprise, keyed by dataflow id.
type ChatGPTEnterpriseIntegrationDataflowsResponse struct {
	// Activity, token consumption, credits, and estimated usage costs across ChatGPT Chat, ChatGPT Work, and Codex, shown in Agent Console with breakdowns by product, model, and user where OpenAI provides them.
	ChatgptEnterpriseAgentConsole *ChatGPTEnterpriseAgentConsoleIntegrationDataflowResponse `json:"chatgpt-enterprise-agent-console,omitempty"`
	// ChatGPT Enterprise costs from the OpenAI Costs API, so that ChatGPT Enterprise spend can be analyzed in [Cloud Cost Management](https://docs.datadoghq.com/cloud_cost_management/). These costs are excluded from Cloud Cost Management spend-monitored billing. Reported costs are estimates, not final invoice amounts.
	ChatgptEnterpriseCloudCostData *ChatGPTEnterpriseCloudCostDataIntegrationDataflowResponse `json:"chatgpt-enterprise-cloud-cost-data,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewChatGPTEnterpriseIntegrationDataflowsResponse instantiates a new ChatGPTEnterpriseIntegrationDataflowsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewChatGPTEnterpriseIntegrationDataflowsResponse() *ChatGPTEnterpriseIntegrationDataflowsResponse {
	this := ChatGPTEnterpriseIntegrationDataflowsResponse{}
	return &this
}

// NewChatGPTEnterpriseIntegrationDataflowsResponseWithDefaults instantiates a new ChatGPTEnterpriseIntegrationDataflowsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewChatGPTEnterpriseIntegrationDataflowsResponseWithDefaults() *ChatGPTEnterpriseIntegrationDataflowsResponse {
	this := ChatGPTEnterpriseIntegrationDataflowsResponse{}
	return &this
}

// GetChatgptEnterpriseAgentConsole returns the ChatgptEnterpriseAgentConsole field value if set, zero value otherwise.
func (o *ChatGPTEnterpriseIntegrationDataflowsResponse) GetChatgptEnterpriseAgentConsole() ChatGPTEnterpriseAgentConsoleIntegrationDataflowResponse {
	if o == nil || o.ChatgptEnterpriseAgentConsole == nil {
		var ret ChatGPTEnterpriseAgentConsoleIntegrationDataflowResponse
		return ret
	}
	return *o.ChatgptEnterpriseAgentConsole
}

// GetChatgptEnterpriseAgentConsoleOk returns a tuple with the ChatgptEnterpriseAgentConsole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatGPTEnterpriseIntegrationDataflowsResponse) GetChatgptEnterpriseAgentConsoleOk() (*ChatGPTEnterpriseAgentConsoleIntegrationDataflowResponse, bool) {
	if o == nil || o.ChatgptEnterpriseAgentConsole == nil {
		return nil, false
	}
	return o.ChatgptEnterpriseAgentConsole, true
}

// HasChatgptEnterpriseAgentConsole returns a boolean if a field has been set.
func (o *ChatGPTEnterpriseIntegrationDataflowsResponse) HasChatgptEnterpriseAgentConsole() bool {
	return o != nil && o.ChatgptEnterpriseAgentConsole != nil
}

// SetChatgptEnterpriseAgentConsole gets a reference to the given ChatGPTEnterpriseAgentConsoleIntegrationDataflowResponse and assigns it to the ChatgptEnterpriseAgentConsole field.
func (o *ChatGPTEnterpriseIntegrationDataflowsResponse) SetChatgptEnterpriseAgentConsole(v ChatGPTEnterpriseAgentConsoleIntegrationDataflowResponse) {
	o.ChatgptEnterpriseAgentConsole = &v
}

// GetChatgptEnterpriseCloudCostData returns the ChatgptEnterpriseCloudCostData field value if set, zero value otherwise.
func (o *ChatGPTEnterpriseIntegrationDataflowsResponse) GetChatgptEnterpriseCloudCostData() ChatGPTEnterpriseCloudCostDataIntegrationDataflowResponse {
	if o == nil || o.ChatgptEnterpriseCloudCostData == nil {
		var ret ChatGPTEnterpriseCloudCostDataIntegrationDataflowResponse
		return ret
	}
	return *o.ChatgptEnterpriseCloudCostData
}

// GetChatgptEnterpriseCloudCostDataOk returns a tuple with the ChatgptEnterpriseCloudCostData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatGPTEnterpriseIntegrationDataflowsResponse) GetChatgptEnterpriseCloudCostDataOk() (*ChatGPTEnterpriseCloudCostDataIntegrationDataflowResponse, bool) {
	if o == nil || o.ChatgptEnterpriseCloudCostData == nil {
		return nil, false
	}
	return o.ChatgptEnterpriseCloudCostData, true
}

// HasChatgptEnterpriseCloudCostData returns a boolean if a field has been set.
func (o *ChatGPTEnterpriseIntegrationDataflowsResponse) HasChatgptEnterpriseCloudCostData() bool {
	return o != nil && o.ChatgptEnterpriseCloudCostData != nil
}

// SetChatgptEnterpriseCloudCostData gets a reference to the given ChatGPTEnterpriseCloudCostDataIntegrationDataflowResponse and assigns it to the ChatgptEnterpriseCloudCostData field.
func (o *ChatGPTEnterpriseIntegrationDataflowsResponse) SetChatgptEnterpriseCloudCostData(v ChatGPTEnterpriseCloudCostDataIntegrationDataflowResponse) {
	o.ChatgptEnterpriseCloudCostData = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ChatGPTEnterpriseIntegrationDataflowsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ChatgptEnterpriseAgentConsole != nil {
		toSerialize["chatgpt-enterprise-agent-console"] = o.ChatgptEnterpriseAgentConsole
	}
	if o.ChatgptEnterpriseCloudCostData != nil {
		toSerialize["chatgpt-enterprise-cloud-cost-data"] = o.ChatgptEnterpriseCloudCostData
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ChatGPTEnterpriseIntegrationDataflowsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ChatgptEnterpriseAgentConsole  *ChatGPTEnterpriseAgentConsoleIntegrationDataflowResponse  `json:"chatgpt-enterprise-agent-console,omitempty"`
		ChatgptEnterpriseCloudCostData *ChatGPTEnterpriseCloudCostDataIntegrationDataflowResponse `json:"chatgpt-enterprise-cloud-cost-data,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"chatgpt-enterprise-agent-console", "chatgpt-enterprise-cloud-cost-data"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ChatgptEnterpriseAgentConsole != nil && all.ChatgptEnterpriseAgentConsole.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ChatgptEnterpriseAgentConsole = all.ChatgptEnterpriseAgentConsole
	if all.ChatgptEnterpriseCloudCostData != nil && all.ChatgptEnterpriseCloudCostData.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ChatgptEnterpriseCloudCostData = all.ChatgptEnterpriseCloudCostData

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
