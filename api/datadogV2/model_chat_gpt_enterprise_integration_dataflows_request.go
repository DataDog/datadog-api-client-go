// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChatGPTEnterpriseIntegrationDataflowsRequest Data Datadog collects from ChatGPT Enterprise, keyed by dataflow id. Each dataflow turns on a distinct kind of collection: set `enabled` to start or stop it. Defaults listed on each dataflow apply when the account is created; on update, omitted fields keep their current values. Every dataflow requires a ChatGPT Enterprise workspace with access to the OpenAI Costs API; a dataflow enabled without it is stored but collects no data.
type ChatGPTEnterpriseIntegrationDataflowsRequest struct {
	// Activity, token consumption, credits, and estimated usage costs across ChatGPT Chat, ChatGPT Work, and Codex, shown in Agent Console with breakdowns by product, model, and user where OpenAI provides them. Data already available from OpenAI appears within 1 hour of enabling this dataflow.
	ChatgptEnterpriseAgentConsole *ChatGPTEnterpriseAgentConsoleIntegrationDataflowRequest `json:"chatgpt-enterprise-agent-console,omitempty"`
	// ChatGPT Enterprise costs from the OpenAI Costs API, so that ChatGPT Enterprise spend can be analyzed in [Cloud Cost Management](https://docs.datadoghq.com/cloud_cost_management/). These costs are excluded from Cloud Cost Management spend-monitored billing. Reported costs are estimates, not final invoice amounts.
	ChatgptEnterpriseCloudCostData *ChatGPTEnterpriseCloudCostDataIntegrationDataflowRequest `json:"chatgpt-enterprise-cloud-cost-data,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewChatGPTEnterpriseIntegrationDataflowsRequest instantiates a new ChatGPTEnterpriseIntegrationDataflowsRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewChatGPTEnterpriseIntegrationDataflowsRequest() *ChatGPTEnterpriseIntegrationDataflowsRequest {
	this := ChatGPTEnterpriseIntegrationDataflowsRequest{}
	return &this
}

// NewChatGPTEnterpriseIntegrationDataflowsRequestWithDefaults instantiates a new ChatGPTEnterpriseIntegrationDataflowsRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewChatGPTEnterpriseIntegrationDataflowsRequestWithDefaults() *ChatGPTEnterpriseIntegrationDataflowsRequest {
	this := ChatGPTEnterpriseIntegrationDataflowsRequest{}
	return &this
}

// GetChatgptEnterpriseAgentConsole returns the ChatgptEnterpriseAgentConsole field value if set, zero value otherwise.
func (o *ChatGPTEnterpriseIntegrationDataflowsRequest) GetChatgptEnterpriseAgentConsole() ChatGPTEnterpriseAgentConsoleIntegrationDataflowRequest {
	if o == nil || o.ChatgptEnterpriseAgentConsole == nil {
		var ret ChatGPTEnterpriseAgentConsoleIntegrationDataflowRequest
		return ret
	}
	return *o.ChatgptEnterpriseAgentConsole
}

// GetChatgptEnterpriseAgentConsoleOk returns a tuple with the ChatgptEnterpriseAgentConsole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatGPTEnterpriseIntegrationDataflowsRequest) GetChatgptEnterpriseAgentConsoleOk() (*ChatGPTEnterpriseAgentConsoleIntegrationDataflowRequest, bool) {
	if o == nil || o.ChatgptEnterpriseAgentConsole == nil {
		return nil, false
	}
	return o.ChatgptEnterpriseAgentConsole, true
}

// HasChatgptEnterpriseAgentConsole returns a boolean if a field has been set.
func (o *ChatGPTEnterpriseIntegrationDataflowsRequest) HasChatgptEnterpriseAgentConsole() bool {
	return o != nil && o.ChatgptEnterpriseAgentConsole != nil
}

// SetChatgptEnterpriseAgentConsole gets a reference to the given ChatGPTEnterpriseAgentConsoleIntegrationDataflowRequest and assigns it to the ChatgptEnterpriseAgentConsole field.
func (o *ChatGPTEnterpriseIntegrationDataflowsRequest) SetChatgptEnterpriseAgentConsole(v ChatGPTEnterpriseAgentConsoleIntegrationDataflowRequest) {
	o.ChatgptEnterpriseAgentConsole = &v
}

// GetChatgptEnterpriseCloudCostData returns the ChatgptEnterpriseCloudCostData field value if set, zero value otherwise.
func (o *ChatGPTEnterpriseIntegrationDataflowsRequest) GetChatgptEnterpriseCloudCostData() ChatGPTEnterpriseCloudCostDataIntegrationDataflowRequest {
	if o == nil || o.ChatgptEnterpriseCloudCostData == nil {
		var ret ChatGPTEnterpriseCloudCostDataIntegrationDataflowRequest
		return ret
	}
	return *o.ChatgptEnterpriseCloudCostData
}

// GetChatgptEnterpriseCloudCostDataOk returns a tuple with the ChatgptEnterpriseCloudCostData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatGPTEnterpriseIntegrationDataflowsRequest) GetChatgptEnterpriseCloudCostDataOk() (*ChatGPTEnterpriseCloudCostDataIntegrationDataflowRequest, bool) {
	if o == nil || o.ChatgptEnterpriseCloudCostData == nil {
		return nil, false
	}
	return o.ChatgptEnterpriseCloudCostData, true
}

// HasChatgptEnterpriseCloudCostData returns a boolean if a field has been set.
func (o *ChatGPTEnterpriseIntegrationDataflowsRequest) HasChatgptEnterpriseCloudCostData() bool {
	return o != nil && o.ChatgptEnterpriseCloudCostData != nil
}

// SetChatgptEnterpriseCloudCostData gets a reference to the given ChatGPTEnterpriseCloudCostDataIntegrationDataflowRequest and assigns it to the ChatgptEnterpriseCloudCostData field.
func (o *ChatGPTEnterpriseIntegrationDataflowsRequest) SetChatgptEnterpriseCloudCostData(v ChatGPTEnterpriseCloudCostDataIntegrationDataflowRequest) {
	o.ChatgptEnterpriseCloudCostData = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ChatGPTEnterpriseIntegrationDataflowsRequest) MarshalJSON() ([]byte, error) {
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
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ChatGPTEnterpriseIntegrationDataflowsRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ChatgptEnterpriseAgentConsole  *ChatGPTEnterpriseAgentConsoleIntegrationDataflowRequest  `json:"chatgpt-enterprise-agent-console,omitempty"`
		ChatgptEnterpriseCloudCostData *ChatGPTEnterpriseCloudCostDataIntegrationDataflowRequest `json:"chatgpt-enterprise-cloud-cost-data,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
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

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
