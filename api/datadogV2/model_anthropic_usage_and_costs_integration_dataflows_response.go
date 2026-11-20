// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnthropicUsageAndCostsIntegrationDataflowsResponse Data Datadog collects from Anthropic Usage and Costs, keyed by dataflow ID.
type AnthropicUsageAndCostsIntegrationDataflowsResponse struct {
	// Anthropic costs by model, workspace, and service tier, in [Cloud Cost Management](https://docs.datadoghq.com/cloud_cost_management/). Requires Cloud Cost Management on your Datadog organization. Without it, this dataflow collects no data.
	AnthropicCloudCostMetrics *AnthropicUsageAndCostsCloudCostMetricsIntegrationDataflowResponse `json:"anthropic-cloud-cost-metrics,omitempty"`
	// Anthropic usage metrics, such as input, output, and cache token consumption by model, workspace, and service tier. With an Analytics API key, usage is also broken down by user, product, context window, inference geography, and speed. This dataflow is always on and cannot be turned off, which is why it appears in responses only.
	AnthropicUsageMetrics *AnthropicUsageAndCostsUsageMetricsIntegrationDataflowResponse `json:"anthropic-usage-metrics,omitempty"`
	// Daily Claude usage data used for DORA Metrics. This dataflow is always on and cannot be turned off, which is why it appears in responses only.
	ClaudeDailyUsageDoraMetrics *AnthropicUsageAndCostsDailyUsageDoraMetricsIntegrationDataflowResponse `json:"claude-daily-usage-dora-metrics,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAnthropicUsageAndCostsIntegrationDataflowsResponse instantiates a new AnthropicUsageAndCostsIntegrationDataflowsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAnthropicUsageAndCostsIntegrationDataflowsResponse() *AnthropicUsageAndCostsIntegrationDataflowsResponse {
	this := AnthropicUsageAndCostsIntegrationDataflowsResponse{}
	return &this
}

// NewAnthropicUsageAndCostsIntegrationDataflowsResponseWithDefaults instantiates a new AnthropicUsageAndCostsIntegrationDataflowsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAnthropicUsageAndCostsIntegrationDataflowsResponseWithDefaults() *AnthropicUsageAndCostsIntegrationDataflowsResponse {
	this := AnthropicUsageAndCostsIntegrationDataflowsResponse{}
	return &this
}

// GetAnthropicCloudCostMetrics returns the AnthropicCloudCostMetrics field value if set, zero value otherwise.
func (o *AnthropicUsageAndCostsIntegrationDataflowsResponse) GetAnthropicCloudCostMetrics() AnthropicUsageAndCostsCloudCostMetricsIntegrationDataflowResponse {
	if o == nil || o.AnthropicCloudCostMetrics == nil {
		var ret AnthropicUsageAndCostsCloudCostMetricsIntegrationDataflowResponse
		return ret
	}
	return *o.AnthropicCloudCostMetrics
}

// GetAnthropicCloudCostMetricsOk returns a tuple with the AnthropicCloudCostMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnthropicUsageAndCostsIntegrationDataflowsResponse) GetAnthropicCloudCostMetricsOk() (*AnthropicUsageAndCostsCloudCostMetricsIntegrationDataflowResponse, bool) {
	if o == nil || o.AnthropicCloudCostMetrics == nil {
		return nil, false
	}
	return o.AnthropicCloudCostMetrics, true
}

// HasAnthropicCloudCostMetrics returns a boolean if a field has been set.
func (o *AnthropicUsageAndCostsIntegrationDataflowsResponse) HasAnthropicCloudCostMetrics() bool {
	return o != nil && o.AnthropicCloudCostMetrics != nil
}

// SetAnthropicCloudCostMetrics gets a reference to the given AnthropicUsageAndCostsCloudCostMetricsIntegrationDataflowResponse and assigns it to the AnthropicCloudCostMetrics field.
func (o *AnthropicUsageAndCostsIntegrationDataflowsResponse) SetAnthropicCloudCostMetrics(v AnthropicUsageAndCostsCloudCostMetricsIntegrationDataflowResponse) {
	o.AnthropicCloudCostMetrics = &v
}

// GetAnthropicUsageMetrics returns the AnthropicUsageMetrics field value if set, zero value otherwise.
func (o *AnthropicUsageAndCostsIntegrationDataflowsResponse) GetAnthropicUsageMetrics() AnthropicUsageAndCostsUsageMetricsIntegrationDataflowResponse {
	if o == nil || o.AnthropicUsageMetrics == nil {
		var ret AnthropicUsageAndCostsUsageMetricsIntegrationDataflowResponse
		return ret
	}
	return *o.AnthropicUsageMetrics
}

// GetAnthropicUsageMetricsOk returns a tuple with the AnthropicUsageMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnthropicUsageAndCostsIntegrationDataflowsResponse) GetAnthropicUsageMetricsOk() (*AnthropicUsageAndCostsUsageMetricsIntegrationDataflowResponse, bool) {
	if o == nil || o.AnthropicUsageMetrics == nil {
		return nil, false
	}
	return o.AnthropicUsageMetrics, true
}

// HasAnthropicUsageMetrics returns a boolean if a field has been set.
func (o *AnthropicUsageAndCostsIntegrationDataflowsResponse) HasAnthropicUsageMetrics() bool {
	return o != nil && o.AnthropicUsageMetrics != nil
}

// SetAnthropicUsageMetrics gets a reference to the given AnthropicUsageAndCostsUsageMetricsIntegrationDataflowResponse and assigns it to the AnthropicUsageMetrics field.
func (o *AnthropicUsageAndCostsIntegrationDataflowsResponse) SetAnthropicUsageMetrics(v AnthropicUsageAndCostsUsageMetricsIntegrationDataflowResponse) {
	o.AnthropicUsageMetrics = &v
}

// GetClaudeDailyUsageDoraMetrics returns the ClaudeDailyUsageDoraMetrics field value if set, zero value otherwise.
func (o *AnthropicUsageAndCostsIntegrationDataflowsResponse) GetClaudeDailyUsageDoraMetrics() AnthropicUsageAndCostsDailyUsageDoraMetricsIntegrationDataflowResponse {
	if o == nil || o.ClaudeDailyUsageDoraMetrics == nil {
		var ret AnthropicUsageAndCostsDailyUsageDoraMetricsIntegrationDataflowResponse
		return ret
	}
	return *o.ClaudeDailyUsageDoraMetrics
}

// GetClaudeDailyUsageDoraMetricsOk returns a tuple with the ClaudeDailyUsageDoraMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnthropicUsageAndCostsIntegrationDataflowsResponse) GetClaudeDailyUsageDoraMetricsOk() (*AnthropicUsageAndCostsDailyUsageDoraMetricsIntegrationDataflowResponse, bool) {
	if o == nil || o.ClaudeDailyUsageDoraMetrics == nil {
		return nil, false
	}
	return o.ClaudeDailyUsageDoraMetrics, true
}

// HasClaudeDailyUsageDoraMetrics returns a boolean if a field has been set.
func (o *AnthropicUsageAndCostsIntegrationDataflowsResponse) HasClaudeDailyUsageDoraMetrics() bool {
	return o != nil && o.ClaudeDailyUsageDoraMetrics != nil
}

// SetClaudeDailyUsageDoraMetrics gets a reference to the given AnthropicUsageAndCostsDailyUsageDoraMetricsIntegrationDataflowResponse and assigns it to the ClaudeDailyUsageDoraMetrics field.
func (o *AnthropicUsageAndCostsIntegrationDataflowsResponse) SetClaudeDailyUsageDoraMetrics(v AnthropicUsageAndCostsDailyUsageDoraMetricsIntegrationDataflowResponse) {
	o.ClaudeDailyUsageDoraMetrics = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AnthropicUsageAndCostsIntegrationDataflowsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AnthropicCloudCostMetrics != nil {
		toSerialize["anthropic-cloud-cost-metrics"] = o.AnthropicCloudCostMetrics
	}
	if o.AnthropicUsageMetrics != nil {
		toSerialize["anthropic-usage-metrics"] = o.AnthropicUsageMetrics
	}
	if o.ClaudeDailyUsageDoraMetrics != nil {
		toSerialize["claude-daily-usage-dora-metrics"] = o.ClaudeDailyUsageDoraMetrics
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AnthropicUsageAndCostsIntegrationDataflowsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AnthropicCloudCostMetrics   *AnthropicUsageAndCostsCloudCostMetricsIntegrationDataflowResponse      `json:"anthropic-cloud-cost-metrics,omitempty"`
		AnthropicUsageMetrics       *AnthropicUsageAndCostsUsageMetricsIntegrationDataflowResponse          `json:"anthropic-usage-metrics,omitempty"`
		ClaudeDailyUsageDoraMetrics *AnthropicUsageAndCostsDailyUsageDoraMetricsIntegrationDataflowResponse `json:"claude-daily-usage-dora-metrics,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"anthropic-cloud-cost-metrics", "anthropic-usage-metrics", "claude-daily-usage-dora-metrics"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AnthropicCloudCostMetrics != nil && all.AnthropicCloudCostMetrics.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AnthropicCloudCostMetrics = all.AnthropicCloudCostMetrics
	if all.AnthropicUsageMetrics != nil && all.AnthropicUsageMetrics.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AnthropicUsageMetrics = all.AnthropicUsageMetrics
	if all.ClaudeDailyUsageDoraMetrics != nil && all.ClaudeDailyUsageDoraMetrics.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ClaudeDailyUsageDoraMetrics = all.ClaudeDailyUsageDoraMetrics

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
