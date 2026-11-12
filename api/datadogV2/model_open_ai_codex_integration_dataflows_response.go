// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpenAICodexIntegrationDataflowsResponse Data Datadog collects from OpenAI Codex, keyed by dataflow id.
type OpenAICodexIntegrationDataflowsResponse struct {
	// A daily per-engineer engagement signal that [AI Impact](https://docs.datadoghq.com/delivery_performance/ai_impact/) uses to attribute AI-assisted pull requests.
	OpenaiCodexDora *OpenAICodexDoraIntegrationDataflowResponse `json:"openai-codex-dora,omitempty"`
	// Daily Codex analytics metrics under the `openai.codex.*` namespace: usage totals, token counts by model, thread and turn counts by client, lines of code attributable to Codex, and code reviews. This dataflow is always on and cannot be turned off, which is why it appears in responses only.
	OpenaiCodexMetrics *OpenAICodexMetricsIntegrationDataflowResponse `json:"openai-codex-metrics,omitempty"`
	// Hourly turn-level metrics from the OpenAI Compliance Logs Platform, including turn and session identity, Work Insights classification, token usage, credits, estimated cost, duration, and skill, plugin, and artifact activity. Turn cost is a provider estimate and can differ from invoiced charges. Requires the `chatgpt.enterprise.compliance_logs_platform.workspace_analytics.read` scope on the key in `authentication`; without it this dataflow collects no data.
	OpenaiCodexTurnMetrics *OpenAICodexTurnMetricsIntegrationDataflowResponse `json:"openai-codex-turn-metrics,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOpenAICodexIntegrationDataflowsResponse instantiates a new OpenAICodexIntegrationDataflowsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpenAICodexIntegrationDataflowsResponse() *OpenAICodexIntegrationDataflowsResponse {
	this := OpenAICodexIntegrationDataflowsResponse{}
	return &this
}

// NewOpenAICodexIntegrationDataflowsResponseWithDefaults instantiates a new OpenAICodexIntegrationDataflowsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpenAICodexIntegrationDataflowsResponseWithDefaults() *OpenAICodexIntegrationDataflowsResponse {
	this := OpenAICodexIntegrationDataflowsResponse{}
	return &this
}

// GetOpenaiCodexDora returns the OpenaiCodexDora field value if set, zero value otherwise.
func (o *OpenAICodexIntegrationDataflowsResponse) GetOpenaiCodexDora() OpenAICodexDoraIntegrationDataflowResponse {
	if o == nil || o.OpenaiCodexDora == nil {
		var ret OpenAICodexDoraIntegrationDataflowResponse
		return ret
	}
	return *o.OpenaiCodexDora
}

// GetOpenaiCodexDoraOk returns a tuple with the OpenaiCodexDora field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAICodexIntegrationDataflowsResponse) GetOpenaiCodexDoraOk() (*OpenAICodexDoraIntegrationDataflowResponse, bool) {
	if o == nil || o.OpenaiCodexDora == nil {
		return nil, false
	}
	return o.OpenaiCodexDora, true
}

// HasOpenaiCodexDora returns a boolean if a field has been set.
func (o *OpenAICodexIntegrationDataflowsResponse) HasOpenaiCodexDora() bool {
	return o != nil && o.OpenaiCodexDora != nil
}

// SetOpenaiCodexDora gets a reference to the given OpenAICodexDoraIntegrationDataflowResponse and assigns it to the OpenaiCodexDora field.
func (o *OpenAICodexIntegrationDataflowsResponse) SetOpenaiCodexDora(v OpenAICodexDoraIntegrationDataflowResponse) {
	o.OpenaiCodexDora = &v
}

// GetOpenaiCodexMetrics returns the OpenaiCodexMetrics field value if set, zero value otherwise.
func (o *OpenAICodexIntegrationDataflowsResponse) GetOpenaiCodexMetrics() OpenAICodexMetricsIntegrationDataflowResponse {
	if o == nil || o.OpenaiCodexMetrics == nil {
		var ret OpenAICodexMetricsIntegrationDataflowResponse
		return ret
	}
	return *o.OpenaiCodexMetrics
}

// GetOpenaiCodexMetricsOk returns a tuple with the OpenaiCodexMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAICodexIntegrationDataflowsResponse) GetOpenaiCodexMetricsOk() (*OpenAICodexMetricsIntegrationDataflowResponse, bool) {
	if o == nil || o.OpenaiCodexMetrics == nil {
		return nil, false
	}
	return o.OpenaiCodexMetrics, true
}

// HasOpenaiCodexMetrics returns a boolean if a field has been set.
func (o *OpenAICodexIntegrationDataflowsResponse) HasOpenaiCodexMetrics() bool {
	return o != nil && o.OpenaiCodexMetrics != nil
}

// SetOpenaiCodexMetrics gets a reference to the given OpenAICodexMetricsIntegrationDataflowResponse and assigns it to the OpenaiCodexMetrics field.
func (o *OpenAICodexIntegrationDataflowsResponse) SetOpenaiCodexMetrics(v OpenAICodexMetricsIntegrationDataflowResponse) {
	o.OpenaiCodexMetrics = &v
}

// GetOpenaiCodexTurnMetrics returns the OpenaiCodexTurnMetrics field value if set, zero value otherwise.
func (o *OpenAICodexIntegrationDataflowsResponse) GetOpenaiCodexTurnMetrics() OpenAICodexTurnMetricsIntegrationDataflowResponse {
	if o == nil || o.OpenaiCodexTurnMetrics == nil {
		var ret OpenAICodexTurnMetricsIntegrationDataflowResponse
		return ret
	}
	return *o.OpenaiCodexTurnMetrics
}

// GetOpenaiCodexTurnMetricsOk returns a tuple with the OpenaiCodexTurnMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAICodexIntegrationDataflowsResponse) GetOpenaiCodexTurnMetricsOk() (*OpenAICodexTurnMetricsIntegrationDataflowResponse, bool) {
	if o == nil || o.OpenaiCodexTurnMetrics == nil {
		return nil, false
	}
	return o.OpenaiCodexTurnMetrics, true
}

// HasOpenaiCodexTurnMetrics returns a boolean if a field has been set.
func (o *OpenAICodexIntegrationDataflowsResponse) HasOpenaiCodexTurnMetrics() bool {
	return o != nil && o.OpenaiCodexTurnMetrics != nil
}

// SetOpenaiCodexTurnMetrics gets a reference to the given OpenAICodexTurnMetricsIntegrationDataflowResponse and assigns it to the OpenaiCodexTurnMetrics field.
func (o *OpenAICodexIntegrationDataflowsResponse) SetOpenaiCodexTurnMetrics(v OpenAICodexTurnMetricsIntegrationDataflowResponse) {
	o.OpenaiCodexTurnMetrics = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpenAICodexIntegrationDataflowsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.OpenaiCodexDora != nil {
		toSerialize["openai-codex-dora"] = o.OpenaiCodexDora
	}
	if o.OpenaiCodexMetrics != nil {
		toSerialize["openai-codex-metrics"] = o.OpenaiCodexMetrics
	}
	if o.OpenaiCodexTurnMetrics != nil {
		toSerialize["openai-codex-turn-metrics"] = o.OpenaiCodexTurnMetrics
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpenAICodexIntegrationDataflowsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OpenaiCodexDora        *OpenAICodexDoraIntegrationDataflowResponse        `json:"openai-codex-dora,omitempty"`
		OpenaiCodexMetrics     *OpenAICodexMetricsIntegrationDataflowResponse     `json:"openai-codex-metrics,omitempty"`
		OpenaiCodexTurnMetrics *OpenAICodexTurnMetricsIntegrationDataflowResponse `json:"openai-codex-turn-metrics,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"openai-codex-dora", "openai-codex-metrics", "openai-codex-turn-metrics"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.OpenaiCodexDora != nil && all.OpenaiCodexDora.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.OpenaiCodexDora = all.OpenaiCodexDora
	if all.OpenaiCodexMetrics != nil && all.OpenaiCodexMetrics.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.OpenaiCodexMetrics = all.OpenaiCodexMetrics
	if all.OpenaiCodexTurnMetrics != nil && all.OpenaiCodexTurnMetrics.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.OpenaiCodexTurnMetrics = all.OpenaiCodexTurnMetrics

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
