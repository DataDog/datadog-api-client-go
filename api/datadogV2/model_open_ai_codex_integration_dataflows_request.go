// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpenAICodexIntegrationDataflowsRequest Data Datadog collects from OpenAI Codex, keyed by dataflow id. Daily Codex analytics metrics are always collected; each dataflow here adds a further kind of collection on top of that baseline, so set `enabled` to start or stop it. Defaults listed on each dataflow apply when the account is created; on update, omitted fields keep their current values. Every dataflow requires an OpenAI enterprise subscription with Codex enabled; a dataflow enabled without it is stored but collects no data.
type OpenAICodexIntegrationDataflowsRequest struct {
	// A daily per-engineer engagement signal that [AI Impact](https://docs.datadoghq.com/delivery_performance/ai_impact/) uses to attribute AI-assisted pull requests.
	OpenaiCodexDora *OpenAICodexDoraIntegrationDataflowRequest `json:"openai-codex-dora,omitempty"`
	// Hourly turn-level metrics from the OpenAI Compliance Logs Platform, including turn and session identity, Work Insights classification, token usage, credits, estimated cost, duration, and skill, plugin, and artifact activity. Turn cost is a provider estimate and can differ from invoiced charges. Requires the `chatgpt.enterprise.compliance_logs_platform.workspace_analytics.read` scope on the key in `authentication`; without it this dataflow collects no data.
	OpenaiCodexTurnMetrics *OpenAICodexTurnMetricsIntegrationDataflowRequest `json:"openai-codex-turn-metrics,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewOpenAICodexIntegrationDataflowsRequest instantiates a new OpenAICodexIntegrationDataflowsRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpenAICodexIntegrationDataflowsRequest() *OpenAICodexIntegrationDataflowsRequest {
	this := OpenAICodexIntegrationDataflowsRequest{}
	return &this
}

// NewOpenAICodexIntegrationDataflowsRequestWithDefaults instantiates a new OpenAICodexIntegrationDataflowsRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpenAICodexIntegrationDataflowsRequestWithDefaults() *OpenAICodexIntegrationDataflowsRequest {
	this := OpenAICodexIntegrationDataflowsRequest{}
	return &this
}

// GetOpenaiCodexDora returns the OpenaiCodexDora field value if set, zero value otherwise.
func (o *OpenAICodexIntegrationDataflowsRequest) GetOpenaiCodexDora() OpenAICodexDoraIntegrationDataflowRequest {
	if o == nil || o.OpenaiCodexDora == nil {
		var ret OpenAICodexDoraIntegrationDataflowRequest
		return ret
	}
	return *o.OpenaiCodexDora
}

// GetOpenaiCodexDoraOk returns a tuple with the OpenaiCodexDora field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAICodexIntegrationDataflowsRequest) GetOpenaiCodexDoraOk() (*OpenAICodexDoraIntegrationDataflowRequest, bool) {
	if o == nil || o.OpenaiCodexDora == nil {
		return nil, false
	}
	return o.OpenaiCodexDora, true
}

// HasOpenaiCodexDora returns a boolean if a field has been set.
func (o *OpenAICodexIntegrationDataflowsRequest) HasOpenaiCodexDora() bool {
	return o != nil && o.OpenaiCodexDora != nil
}

// SetOpenaiCodexDora gets a reference to the given OpenAICodexDoraIntegrationDataflowRequest and assigns it to the OpenaiCodexDora field.
func (o *OpenAICodexIntegrationDataflowsRequest) SetOpenaiCodexDora(v OpenAICodexDoraIntegrationDataflowRequest) {
	o.OpenaiCodexDora = &v
}

// GetOpenaiCodexTurnMetrics returns the OpenaiCodexTurnMetrics field value if set, zero value otherwise.
func (o *OpenAICodexIntegrationDataflowsRequest) GetOpenaiCodexTurnMetrics() OpenAICodexTurnMetricsIntegrationDataflowRequest {
	if o == nil || o.OpenaiCodexTurnMetrics == nil {
		var ret OpenAICodexTurnMetricsIntegrationDataflowRequest
		return ret
	}
	return *o.OpenaiCodexTurnMetrics
}

// GetOpenaiCodexTurnMetricsOk returns a tuple with the OpenaiCodexTurnMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAICodexIntegrationDataflowsRequest) GetOpenaiCodexTurnMetricsOk() (*OpenAICodexTurnMetricsIntegrationDataflowRequest, bool) {
	if o == nil || o.OpenaiCodexTurnMetrics == nil {
		return nil, false
	}
	return o.OpenaiCodexTurnMetrics, true
}

// HasOpenaiCodexTurnMetrics returns a boolean if a field has been set.
func (o *OpenAICodexIntegrationDataflowsRequest) HasOpenaiCodexTurnMetrics() bool {
	return o != nil && o.OpenaiCodexTurnMetrics != nil
}

// SetOpenaiCodexTurnMetrics gets a reference to the given OpenAICodexTurnMetricsIntegrationDataflowRequest and assigns it to the OpenaiCodexTurnMetrics field.
func (o *OpenAICodexIntegrationDataflowsRequest) SetOpenaiCodexTurnMetrics(v OpenAICodexTurnMetricsIntegrationDataflowRequest) {
	o.OpenaiCodexTurnMetrics = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpenAICodexIntegrationDataflowsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.OpenaiCodexDora != nil {
		toSerialize["openai-codex-dora"] = o.OpenaiCodexDora
	}
	if o.OpenaiCodexTurnMetrics != nil {
		toSerialize["openai-codex-turn-metrics"] = o.OpenaiCodexTurnMetrics
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpenAICodexIntegrationDataflowsRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OpenaiCodexDora        *OpenAICodexDoraIntegrationDataflowRequest        `json:"openai-codex-dora,omitempty"`
		OpenaiCodexTurnMetrics *OpenAICodexTurnMetricsIntegrationDataflowRequest `json:"openai-codex-turn-metrics,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	hasInvalidField := false
	if all.OpenaiCodexDora != nil && all.OpenaiCodexDora.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.OpenaiCodexDora = all.OpenaiCodexDora
	if all.OpenaiCodexTurnMetrics != nil && all.OpenaiCodexTurnMetrics.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.OpenaiCodexTurnMetrics = all.OpenaiCodexTurnMetrics

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
