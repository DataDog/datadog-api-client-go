// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpenAILLMEvaluationsIntegrationDataflowsResponse Data Datadog collects from OpenAI LLM Evaluations, keyed by dataflow ID.
type OpenAILLMEvaluationsIntegrationDataflowsResponse struct {
	// Makes the key in `authentication` available to run evaluations in [LLM Observability](https://docs.datadoghq.com/llm_observability/).
	OpenaiLlmCompletions *OpenAILLMEvaluationsLLMCompletionsIntegrationDataflowResponse `json:"openai-llm-completions,omitempty"`
	// Costs of the OpenAI project that the key in `authentication` belongs to, which you can analyze in [Cloud Cost Management](https://docs.datadoghq.com/cloud_cost_management/).
	OpenaiProjectCloudCostData *OpenAILLMEvaluationsProjectCloudCostDataIntegrationDataflowResponse `json:"openai-project-cloud-cost-data,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOpenAILLMEvaluationsIntegrationDataflowsResponse instantiates a new OpenAILLMEvaluationsIntegrationDataflowsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpenAILLMEvaluationsIntegrationDataflowsResponse() *OpenAILLMEvaluationsIntegrationDataflowsResponse {
	this := OpenAILLMEvaluationsIntegrationDataflowsResponse{}
	return &this
}

// NewOpenAILLMEvaluationsIntegrationDataflowsResponseWithDefaults instantiates a new OpenAILLMEvaluationsIntegrationDataflowsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpenAILLMEvaluationsIntegrationDataflowsResponseWithDefaults() *OpenAILLMEvaluationsIntegrationDataflowsResponse {
	this := OpenAILLMEvaluationsIntegrationDataflowsResponse{}
	return &this
}

// GetOpenaiLlmCompletions returns the OpenaiLlmCompletions field value if set, zero value otherwise.
func (o *OpenAILLMEvaluationsIntegrationDataflowsResponse) GetOpenaiLlmCompletions() OpenAILLMEvaluationsLLMCompletionsIntegrationDataflowResponse {
	if o == nil || o.OpenaiLlmCompletions == nil {
		var ret OpenAILLMEvaluationsLLMCompletionsIntegrationDataflowResponse
		return ret
	}
	return *o.OpenaiLlmCompletions
}

// GetOpenaiLlmCompletionsOk returns a tuple with the OpenaiLlmCompletions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAILLMEvaluationsIntegrationDataflowsResponse) GetOpenaiLlmCompletionsOk() (*OpenAILLMEvaluationsLLMCompletionsIntegrationDataflowResponse, bool) {
	if o == nil || o.OpenaiLlmCompletions == nil {
		return nil, false
	}
	return o.OpenaiLlmCompletions, true
}

// HasOpenaiLlmCompletions returns a boolean if a field has been set.
func (o *OpenAILLMEvaluationsIntegrationDataflowsResponse) HasOpenaiLlmCompletions() bool {
	return o != nil && o.OpenaiLlmCompletions != nil
}

// SetOpenaiLlmCompletions gets a reference to the given OpenAILLMEvaluationsLLMCompletionsIntegrationDataflowResponse and assigns it to the OpenaiLlmCompletions field.
func (o *OpenAILLMEvaluationsIntegrationDataflowsResponse) SetOpenaiLlmCompletions(v OpenAILLMEvaluationsLLMCompletionsIntegrationDataflowResponse) {
	o.OpenaiLlmCompletions = &v
}

// GetOpenaiProjectCloudCostData returns the OpenaiProjectCloudCostData field value if set, zero value otherwise.
func (o *OpenAILLMEvaluationsIntegrationDataflowsResponse) GetOpenaiProjectCloudCostData() OpenAILLMEvaluationsProjectCloudCostDataIntegrationDataflowResponse {
	if o == nil || o.OpenaiProjectCloudCostData == nil {
		var ret OpenAILLMEvaluationsProjectCloudCostDataIntegrationDataflowResponse
		return ret
	}
	return *o.OpenaiProjectCloudCostData
}

// GetOpenaiProjectCloudCostDataOk returns a tuple with the OpenaiProjectCloudCostData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAILLMEvaluationsIntegrationDataflowsResponse) GetOpenaiProjectCloudCostDataOk() (*OpenAILLMEvaluationsProjectCloudCostDataIntegrationDataflowResponse, bool) {
	if o == nil || o.OpenaiProjectCloudCostData == nil {
		return nil, false
	}
	return o.OpenaiProjectCloudCostData, true
}

// HasOpenaiProjectCloudCostData returns a boolean if a field has been set.
func (o *OpenAILLMEvaluationsIntegrationDataflowsResponse) HasOpenaiProjectCloudCostData() bool {
	return o != nil && o.OpenaiProjectCloudCostData != nil
}

// SetOpenaiProjectCloudCostData gets a reference to the given OpenAILLMEvaluationsProjectCloudCostDataIntegrationDataflowResponse and assigns it to the OpenaiProjectCloudCostData field.
func (o *OpenAILLMEvaluationsIntegrationDataflowsResponse) SetOpenaiProjectCloudCostData(v OpenAILLMEvaluationsProjectCloudCostDataIntegrationDataflowResponse) {
	o.OpenaiProjectCloudCostData = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpenAILLMEvaluationsIntegrationDataflowsResponse) MarshalJSON() ([]byte, error) {
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
func (o *OpenAILLMEvaluationsIntegrationDataflowsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OpenaiLlmCompletions       *OpenAILLMEvaluationsLLMCompletionsIntegrationDataflowResponse       `json:"openai-llm-completions,omitempty"`
		OpenaiProjectCloudCostData *OpenAILLMEvaluationsProjectCloudCostDataIntegrationDataflowResponse `json:"openai-project-cloud-cost-data,omitempty"`
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
