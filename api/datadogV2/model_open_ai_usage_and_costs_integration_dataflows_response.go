// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpenAIUsageAndCostsIntegrationDataflowsResponse Data Datadog collects from OpenAI Usage and Costs, keyed by dataflow ID.
type OpenAIUsageAndCostsIntegrationDataflowsResponse struct {
	// OpenAI organization costs, which you can analyze in [Cloud Cost Management](https://docs.datadoghq.com/cloud_cost_management/). See [collected data](https://docs.datadoghq.com/cloud_cost_management/saas_costs/?tab=openai#data-collected) for what is included.
	OpenaiCloudCostData *OpenAIUsageAndCostsCloudCostDataIntegrationDataflowResponse `json:"openai-cloud-cost-data,omitempty"`
	// Organization usage metrics under the `openai.*` namespace, such as `openai.completions.*`, `openai.embeddings.*`, `openai.images.*`, `openai.moderations.*`, `openai.audio_speeches.*`, `openai.audio_transcriptions.*`, `openai.code_interpreter_sessions.*`, and `openai.vector_stores.*`. This dataflow is always on and cannot be turned off, so it appears only in responses.
	OpenaiUsageMetrics *OpenAIUsageAndCostsUsageMetricsIntegrationDataflowResponse `json:"openai-usage-metrics,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOpenAIUsageAndCostsIntegrationDataflowsResponse instantiates a new OpenAIUsageAndCostsIntegrationDataflowsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpenAIUsageAndCostsIntegrationDataflowsResponse() *OpenAIUsageAndCostsIntegrationDataflowsResponse {
	this := OpenAIUsageAndCostsIntegrationDataflowsResponse{}
	return &this
}

// NewOpenAIUsageAndCostsIntegrationDataflowsResponseWithDefaults instantiates a new OpenAIUsageAndCostsIntegrationDataflowsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpenAIUsageAndCostsIntegrationDataflowsResponseWithDefaults() *OpenAIUsageAndCostsIntegrationDataflowsResponse {
	this := OpenAIUsageAndCostsIntegrationDataflowsResponse{}
	return &this
}

// GetOpenaiCloudCostData returns the OpenaiCloudCostData field value if set, zero value otherwise.
func (o *OpenAIUsageAndCostsIntegrationDataflowsResponse) GetOpenaiCloudCostData() OpenAIUsageAndCostsCloudCostDataIntegrationDataflowResponse {
	if o == nil || o.OpenaiCloudCostData == nil {
		var ret OpenAIUsageAndCostsCloudCostDataIntegrationDataflowResponse
		return ret
	}
	return *o.OpenaiCloudCostData
}

// GetOpenaiCloudCostDataOk returns a tuple with the OpenaiCloudCostData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAIUsageAndCostsIntegrationDataflowsResponse) GetOpenaiCloudCostDataOk() (*OpenAIUsageAndCostsCloudCostDataIntegrationDataflowResponse, bool) {
	if o == nil || o.OpenaiCloudCostData == nil {
		return nil, false
	}
	return o.OpenaiCloudCostData, true
}

// HasOpenaiCloudCostData returns a boolean if a field has been set.
func (o *OpenAIUsageAndCostsIntegrationDataflowsResponse) HasOpenaiCloudCostData() bool {
	return o != nil && o.OpenaiCloudCostData != nil
}

// SetOpenaiCloudCostData gets a reference to the given OpenAIUsageAndCostsCloudCostDataIntegrationDataflowResponse and assigns it to the OpenaiCloudCostData field.
func (o *OpenAIUsageAndCostsIntegrationDataflowsResponse) SetOpenaiCloudCostData(v OpenAIUsageAndCostsCloudCostDataIntegrationDataflowResponse) {
	o.OpenaiCloudCostData = &v
}

// GetOpenaiUsageMetrics returns the OpenaiUsageMetrics field value if set, zero value otherwise.
func (o *OpenAIUsageAndCostsIntegrationDataflowsResponse) GetOpenaiUsageMetrics() OpenAIUsageAndCostsUsageMetricsIntegrationDataflowResponse {
	if o == nil || o.OpenaiUsageMetrics == nil {
		var ret OpenAIUsageAndCostsUsageMetricsIntegrationDataflowResponse
		return ret
	}
	return *o.OpenaiUsageMetrics
}

// GetOpenaiUsageMetricsOk returns a tuple with the OpenaiUsageMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAIUsageAndCostsIntegrationDataflowsResponse) GetOpenaiUsageMetricsOk() (*OpenAIUsageAndCostsUsageMetricsIntegrationDataflowResponse, bool) {
	if o == nil || o.OpenaiUsageMetrics == nil {
		return nil, false
	}
	return o.OpenaiUsageMetrics, true
}

// HasOpenaiUsageMetrics returns a boolean if a field has been set.
func (o *OpenAIUsageAndCostsIntegrationDataflowsResponse) HasOpenaiUsageMetrics() bool {
	return o != nil && o.OpenaiUsageMetrics != nil
}

// SetOpenaiUsageMetrics gets a reference to the given OpenAIUsageAndCostsUsageMetricsIntegrationDataflowResponse and assigns it to the OpenaiUsageMetrics field.
func (o *OpenAIUsageAndCostsIntegrationDataflowsResponse) SetOpenaiUsageMetrics(v OpenAIUsageAndCostsUsageMetricsIntegrationDataflowResponse) {
	o.OpenaiUsageMetrics = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpenAIUsageAndCostsIntegrationDataflowsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.OpenaiCloudCostData != nil {
		toSerialize["openai-cloud-cost-data"] = o.OpenaiCloudCostData
	}
	if o.OpenaiUsageMetrics != nil {
		toSerialize["openai-usage-metrics"] = o.OpenaiUsageMetrics
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpenAIUsageAndCostsIntegrationDataflowsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OpenaiCloudCostData *OpenAIUsageAndCostsCloudCostDataIntegrationDataflowResponse `json:"openai-cloud-cost-data,omitempty"`
		OpenaiUsageMetrics  *OpenAIUsageAndCostsUsageMetricsIntegrationDataflowResponse  `json:"openai-usage-metrics,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"openai-cloud-cost-data", "openai-usage-metrics"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.OpenaiCloudCostData != nil && all.OpenaiCloudCostData.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.OpenaiCloudCostData = all.OpenaiCloudCostData
	if all.OpenaiUsageMetrics != nil && all.OpenaiUsageMetrics.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.OpenaiUsageMetrics = all.OpenaiUsageMetrics

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
