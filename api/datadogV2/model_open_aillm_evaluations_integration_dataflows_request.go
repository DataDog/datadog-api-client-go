// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpenAILLMEvaluationsIntegrationDataflowsRequest Data Datadog collects from OpenAI LLM Evaluations, keyed by dataflow ID. Each dataflow turns on a distinct type of collection. Set `enabled` to start or stop a dataflow. The default for each dataflow applies when the account is created. On update, omitted fields keep their current values.
type OpenAILLMEvaluationsIntegrationDataflowsRequest struct {
	// Makes the key in `authentication` available to run evaluations in [LLM Observability](https://docs.datadoghq.com/llm_observability/). While this dataflow is enabled, a key that cannot call the OpenAI chat completions endpoint is rejected with a `422` response.
	OpenaiLlmCompletions *OpenAILLMEvaluationsLLMCompletionsIntegrationDataflowRequest `json:"openai-llm-completions,omitempty"`
	// Costs of the OpenAI project that the key in `authentication` belongs to, which you can analyze in [Cloud Cost Management](https://docs.datadoghq.com/cloud_cost_management/). Cost data appears in Cloud Cost Management within 24 hours after you enable this dataflow.
	OpenaiProjectCloudCostData *OpenAILLMEvaluationsProjectCloudCostDataIntegrationDataflowRequest `json:"openai-project-cloud-cost-data,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewOpenAILLMEvaluationsIntegrationDataflowsRequest instantiates a new OpenAILLMEvaluationsIntegrationDataflowsRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpenAILLMEvaluationsIntegrationDataflowsRequest() *OpenAILLMEvaluationsIntegrationDataflowsRequest {
	this := OpenAILLMEvaluationsIntegrationDataflowsRequest{}
	return &this
}

// NewOpenAILLMEvaluationsIntegrationDataflowsRequestWithDefaults instantiates a new OpenAILLMEvaluationsIntegrationDataflowsRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpenAILLMEvaluationsIntegrationDataflowsRequestWithDefaults() *OpenAILLMEvaluationsIntegrationDataflowsRequest {
	this := OpenAILLMEvaluationsIntegrationDataflowsRequest{}
	return &this
}

// GetOpenaiLlmCompletions returns the OpenaiLlmCompletions field value if set, zero value otherwise.
func (o *OpenAILLMEvaluationsIntegrationDataflowsRequest) GetOpenaiLlmCompletions() OpenAILLMEvaluationsLLMCompletionsIntegrationDataflowRequest {
	if o == nil || o.OpenaiLlmCompletions == nil {
		var ret OpenAILLMEvaluationsLLMCompletionsIntegrationDataflowRequest
		return ret
	}
	return *o.OpenaiLlmCompletions
}

// GetOpenaiLlmCompletionsOk returns a tuple with the OpenaiLlmCompletions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAILLMEvaluationsIntegrationDataflowsRequest) GetOpenaiLlmCompletionsOk() (*OpenAILLMEvaluationsLLMCompletionsIntegrationDataflowRequest, bool) {
	if o == nil || o.OpenaiLlmCompletions == nil {
		return nil, false
	}
	return o.OpenaiLlmCompletions, true
}

// HasOpenaiLlmCompletions returns a boolean if a field has been set.
func (o *OpenAILLMEvaluationsIntegrationDataflowsRequest) HasOpenaiLlmCompletions() bool {
	return o != nil && o.OpenaiLlmCompletions != nil
}

// SetOpenaiLlmCompletions gets a reference to the given OpenAILLMEvaluationsLLMCompletionsIntegrationDataflowRequest and assigns it to the OpenaiLlmCompletions field.
func (o *OpenAILLMEvaluationsIntegrationDataflowsRequest) SetOpenaiLlmCompletions(v OpenAILLMEvaluationsLLMCompletionsIntegrationDataflowRequest) {
	o.OpenaiLlmCompletions = &v
}

// GetOpenaiProjectCloudCostData returns the OpenaiProjectCloudCostData field value if set, zero value otherwise.
func (o *OpenAILLMEvaluationsIntegrationDataflowsRequest) GetOpenaiProjectCloudCostData() OpenAILLMEvaluationsProjectCloudCostDataIntegrationDataflowRequest {
	if o == nil || o.OpenaiProjectCloudCostData == nil {
		var ret OpenAILLMEvaluationsProjectCloudCostDataIntegrationDataflowRequest
		return ret
	}
	return *o.OpenaiProjectCloudCostData
}

// GetOpenaiProjectCloudCostDataOk returns a tuple with the OpenaiProjectCloudCostData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAILLMEvaluationsIntegrationDataflowsRequest) GetOpenaiProjectCloudCostDataOk() (*OpenAILLMEvaluationsProjectCloudCostDataIntegrationDataflowRequest, bool) {
	if o == nil || o.OpenaiProjectCloudCostData == nil {
		return nil, false
	}
	return o.OpenaiProjectCloudCostData, true
}

// HasOpenaiProjectCloudCostData returns a boolean if a field has been set.
func (o *OpenAILLMEvaluationsIntegrationDataflowsRequest) HasOpenaiProjectCloudCostData() bool {
	return o != nil && o.OpenaiProjectCloudCostData != nil
}

// SetOpenaiProjectCloudCostData gets a reference to the given OpenAILLMEvaluationsProjectCloudCostDataIntegrationDataflowRequest and assigns it to the OpenaiProjectCloudCostData field.
func (o *OpenAILLMEvaluationsIntegrationDataflowsRequest) SetOpenaiProjectCloudCostData(v OpenAILLMEvaluationsProjectCloudCostDataIntegrationDataflowRequest) {
	o.OpenaiProjectCloudCostData = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpenAILLMEvaluationsIntegrationDataflowsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.OpenaiLlmCompletions != nil {
		toSerialize["openai-llm-completions"] = o.OpenaiLlmCompletions
	}
	if o.OpenaiProjectCloudCostData != nil {
		toSerialize["openai-project-cloud-cost-data"] = o.OpenaiProjectCloudCostData
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpenAILLMEvaluationsIntegrationDataflowsRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OpenaiLlmCompletions       *OpenAILLMEvaluationsLLMCompletionsIntegrationDataflowRequest       `json:"openai-llm-completions,omitempty"`
		OpenaiProjectCloudCostData *OpenAILLMEvaluationsProjectCloudCostDataIntegrationDataflowRequest `json:"openai-project-cloud-cost-data,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	hasInvalidField := false
	if all.OpenaiLlmCompletions != nil && all.OpenaiLlmCompletions.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.OpenaiLlmCompletions = all.OpenaiLlmCompletions
	if all.OpenaiProjectCloudCostData != nil && all.OpenaiProjectCloudCostData.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.OpenaiProjectCloudCostData = all.OpenaiProjectCloudCostData

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
