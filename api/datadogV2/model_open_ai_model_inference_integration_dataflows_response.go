// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpenAIModelInferenceIntegrationDataflowsResponse Data Datadog collects from OpenAI Model Inference, keyed by dataflow ID.
type OpenAIModelInferenceIntegrationDataflowsResponse struct {
	// Makes the key in `authentication` available to run evaluations in [LLM Observability](https://docs.datadoghq.com/llm_observability/).
	OpenaiLlmCompletions *OpenAIModelInferenceLLMCompletionsIntegrationDataflowResponse `json:"openai-llm-completions,omitempty"`
	// Costs of the OpenAI project that owns the key in `authentication`. Analyze these costs in [Cloud Cost Management](https://docs.datadoghq.com/cloud_cost_management/).
	OpenaiProjectCloudCostData *OpenAIModelInferenceProjectCloudCostDataIntegrationDataflowResponse `json:"openai-project-cloud-cost-data,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOpenAIModelInferenceIntegrationDataflowsResponse instantiates a new OpenAIModelInferenceIntegrationDataflowsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpenAIModelInferenceIntegrationDataflowsResponse() *OpenAIModelInferenceIntegrationDataflowsResponse {
	this := OpenAIModelInferenceIntegrationDataflowsResponse{}
	return &this
}

// NewOpenAIModelInferenceIntegrationDataflowsResponseWithDefaults instantiates a new OpenAIModelInferenceIntegrationDataflowsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpenAIModelInferenceIntegrationDataflowsResponseWithDefaults() *OpenAIModelInferenceIntegrationDataflowsResponse {
	this := OpenAIModelInferenceIntegrationDataflowsResponse{}
	return &this
}

// GetOpenaiLlmCompletions returns the OpenaiLlmCompletions field value if set, zero value otherwise.
func (o *OpenAIModelInferenceIntegrationDataflowsResponse) GetOpenaiLlmCompletions() OpenAIModelInferenceLLMCompletionsIntegrationDataflowResponse {
	if o == nil || o.OpenaiLlmCompletions == nil {
		var ret OpenAIModelInferenceLLMCompletionsIntegrationDataflowResponse
		return ret
	}
	return *o.OpenaiLlmCompletions
}

// GetOpenaiLlmCompletionsOk returns a tuple with the OpenaiLlmCompletions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAIModelInferenceIntegrationDataflowsResponse) GetOpenaiLlmCompletionsOk() (*OpenAIModelInferenceLLMCompletionsIntegrationDataflowResponse, bool) {
	if o == nil || o.OpenaiLlmCompletions == nil {
		return nil, false
	}
	return o.OpenaiLlmCompletions, true
}

// HasOpenaiLlmCompletions returns a boolean if a field has been set.
func (o *OpenAIModelInferenceIntegrationDataflowsResponse) HasOpenaiLlmCompletions() bool {
	return o != nil && o.OpenaiLlmCompletions != nil
}

// SetOpenaiLlmCompletions gets a reference to the given OpenAIModelInferenceLLMCompletionsIntegrationDataflowResponse and assigns it to the OpenaiLlmCompletions field.
func (o *OpenAIModelInferenceIntegrationDataflowsResponse) SetOpenaiLlmCompletions(v OpenAIModelInferenceLLMCompletionsIntegrationDataflowResponse) {
	o.OpenaiLlmCompletions = &v
}

// GetOpenaiProjectCloudCostData returns the OpenaiProjectCloudCostData field value if set, zero value otherwise.
func (o *OpenAIModelInferenceIntegrationDataflowsResponse) GetOpenaiProjectCloudCostData() OpenAIModelInferenceProjectCloudCostDataIntegrationDataflowResponse {
	if o == nil || o.OpenaiProjectCloudCostData == nil {
		var ret OpenAIModelInferenceProjectCloudCostDataIntegrationDataflowResponse
		return ret
	}
	return *o.OpenaiProjectCloudCostData
}

// GetOpenaiProjectCloudCostDataOk returns a tuple with the OpenaiProjectCloudCostData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAIModelInferenceIntegrationDataflowsResponse) GetOpenaiProjectCloudCostDataOk() (*OpenAIModelInferenceProjectCloudCostDataIntegrationDataflowResponse, bool) {
	if o == nil || o.OpenaiProjectCloudCostData == nil {
		return nil, false
	}
	return o.OpenaiProjectCloudCostData, true
}

// HasOpenaiProjectCloudCostData returns a boolean if a field has been set.
func (o *OpenAIModelInferenceIntegrationDataflowsResponse) HasOpenaiProjectCloudCostData() bool {
	return o != nil && o.OpenaiProjectCloudCostData != nil
}

// SetOpenaiProjectCloudCostData gets a reference to the given OpenAIModelInferenceProjectCloudCostDataIntegrationDataflowResponse and assigns it to the OpenaiProjectCloudCostData field.
func (o *OpenAIModelInferenceIntegrationDataflowsResponse) SetOpenaiProjectCloudCostData(v OpenAIModelInferenceProjectCloudCostDataIntegrationDataflowResponse) {
	o.OpenaiProjectCloudCostData = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpenAIModelInferenceIntegrationDataflowsResponse) MarshalJSON() ([]byte, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpenAIModelInferenceIntegrationDataflowsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OpenaiLlmCompletions       *OpenAIModelInferenceLLMCompletionsIntegrationDataflowResponse       `json:"openai-llm-completions,omitempty"`
		OpenaiProjectCloudCostData *OpenAIModelInferenceProjectCloudCostDataIntegrationDataflowResponse `json:"openai-project-cloud-cost-data,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"openai-llm-completions", "openai-project-cloud-cost-data"})
	} else {
		return err
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

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
