// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpenAIModelInferenceIntegrationDataflowsRequest Data Datadog collects from OpenAI Model Inference, keyed by dataflow ID. Each dataflow turns on a distinct type of collection. Set `enabled` to start or stop a dataflow. The default for each dataflow applies when the account is created. On update, omitted fields keep their current values.
type OpenAIModelInferenceIntegrationDataflowsRequest struct {
	// Makes the key in `authentication` available to run evaluations in [LLM Observability](https://docs.datadoghq.com/llm_observability/). While this dataflow is enabled, the API returns a `422` response for a key that cannot call the OpenAI chat completions endpoint.
	OpenaiLlmCompletions *OpenAIModelInferenceLLMCompletionsIntegrationDataflowRequest `json:"openai-llm-completions,omitempty"`
	// Costs of the OpenAI project that owns the key in `authentication`. Analyze these costs in [Cloud Cost Management](https://docs.datadoghq.com/cloud_cost_management/). Cost data appears in Cloud Cost Management within 24 hours after you enable this dataflow.
	OpenaiProjectCloudCostData *OpenAIModelInferenceProjectCloudCostDataIntegrationDataflowRequest `json:"openai-project-cloud-cost-data,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewOpenAIModelInferenceIntegrationDataflowsRequest instantiates a new OpenAIModelInferenceIntegrationDataflowsRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpenAIModelInferenceIntegrationDataflowsRequest() *OpenAIModelInferenceIntegrationDataflowsRequest {
	this := OpenAIModelInferenceIntegrationDataflowsRequest{}
	return &this
}

// NewOpenAIModelInferenceIntegrationDataflowsRequestWithDefaults instantiates a new OpenAIModelInferenceIntegrationDataflowsRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpenAIModelInferenceIntegrationDataflowsRequestWithDefaults() *OpenAIModelInferenceIntegrationDataflowsRequest {
	this := OpenAIModelInferenceIntegrationDataflowsRequest{}
	return &this
}

// GetOpenaiLlmCompletions returns the OpenaiLlmCompletions field value if set, zero value otherwise.
func (o *OpenAIModelInferenceIntegrationDataflowsRequest) GetOpenaiLlmCompletions() OpenAIModelInferenceLLMCompletionsIntegrationDataflowRequest {
	if o == nil || o.OpenaiLlmCompletions == nil {
		var ret OpenAIModelInferenceLLMCompletionsIntegrationDataflowRequest
		return ret
	}
	return *o.OpenaiLlmCompletions
}

// GetOpenaiLlmCompletionsOk returns a tuple with the OpenaiLlmCompletions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAIModelInferenceIntegrationDataflowsRequest) GetOpenaiLlmCompletionsOk() (*OpenAIModelInferenceLLMCompletionsIntegrationDataflowRequest, bool) {
	if o == nil || o.OpenaiLlmCompletions == nil {
		return nil, false
	}
	return o.OpenaiLlmCompletions, true
}

// HasOpenaiLlmCompletions returns a boolean if a field has been set.
func (o *OpenAIModelInferenceIntegrationDataflowsRequest) HasOpenaiLlmCompletions() bool {
	return o != nil && o.OpenaiLlmCompletions != nil
}

// SetOpenaiLlmCompletions gets a reference to the given OpenAIModelInferenceLLMCompletionsIntegrationDataflowRequest and assigns it to the OpenaiLlmCompletions field.
func (o *OpenAIModelInferenceIntegrationDataflowsRequest) SetOpenaiLlmCompletions(v OpenAIModelInferenceLLMCompletionsIntegrationDataflowRequest) {
	o.OpenaiLlmCompletions = &v
}

// GetOpenaiProjectCloudCostData returns the OpenaiProjectCloudCostData field value if set, zero value otherwise.
func (o *OpenAIModelInferenceIntegrationDataflowsRequest) GetOpenaiProjectCloudCostData() OpenAIModelInferenceProjectCloudCostDataIntegrationDataflowRequest {
	if o == nil || o.OpenaiProjectCloudCostData == nil {
		var ret OpenAIModelInferenceProjectCloudCostDataIntegrationDataflowRequest
		return ret
	}
	return *o.OpenaiProjectCloudCostData
}

// GetOpenaiProjectCloudCostDataOk returns a tuple with the OpenaiProjectCloudCostData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAIModelInferenceIntegrationDataflowsRequest) GetOpenaiProjectCloudCostDataOk() (*OpenAIModelInferenceProjectCloudCostDataIntegrationDataflowRequest, bool) {
	if o == nil || o.OpenaiProjectCloudCostData == nil {
		return nil, false
	}
	return o.OpenaiProjectCloudCostData, true
}

// HasOpenaiProjectCloudCostData returns a boolean if a field has been set.
func (o *OpenAIModelInferenceIntegrationDataflowsRequest) HasOpenaiProjectCloudCostData() bool {
	return o != nil && o.OpenaiProjectCloudCostData != nil
}

// SetOpenaiProjectCloudCostData gets a reference to the given OpenAIModelInferenceProjectCloudCostDataIntegrationDataflowRequest and assigns it to the OpenaiProjectCloudCostData field.
func (o *OpenAIModelInferenceIntegrationDataflowsRequest) SetOpenaiProjectCloudCostData(v OpenAIModelInferenceProjectCloudCostDataIntegrationDataflowRequest) {
	o.OpenaiProjectCloudCostData = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpenAIModelInferenceIntegrationDataflowsRequest) MarshalJSON() ([]byte, error) {
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
func (o *OpenAIModelInferenceIntegrationDataflowsRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OpenaiLlmCompletions       *OpenAIModelInferenceLLMCompletionsIntegrationDataflowRequest       `json:"openai-llm-completions,omitempty"`
		OpenaiProjectCloudCostData *OpenAIModelInferenceProjectCloudCostDataIntegrationDataflowRequest `json:"openai-project-cloud-cost-data,omitempty"`
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
