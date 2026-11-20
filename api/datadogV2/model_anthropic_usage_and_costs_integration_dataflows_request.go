// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnthropicUsageAndCostsIntegrationDataflowsRequest Data Datadog collects from Anthropic Usage and Costs, keyed by dataflow ID. Datadog always collects usage metrics. Each dataflow in this object adds another type of collection. Set `enabled` to start or stop a dataflow. The default for each dataflow applies when the account is created. On update, omitted fields keep their current values. If a dataflow depends on a Datadog product, its description names that product. If you enable a dataflow without that product, Datadog stores the setting but collects no data.
type AnthropicUsageAndCostsIntegrationDataflowsRequest struct {
	// Anthropic costs by model, workspace, and service tier, in [Cloud Cost Management](https://docs.datadoghq.com/cloud_cost_management/). Requires Cloud Cost Management on your Datadog organization. Without it, this dataflow collects no data. Cost data appears in Cloud Cost Management within 24 to 48 hours of enabling this dataflow.
	AnthropicCloudCostMetrics *AnthropicUsageAndCostsCloudCostMetricsIntegrationDataflowRequest `json:"anthropic-cloud-cost-metrics,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewAnthropicUsageAndCostsIntegrationDataflowsRequest instantiates a new AnthropicUsageAndCostsIntegrationDataflowsRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAnthropicUsageAndCostsIntegrationDataflowsRequest() *AnthropicUsageAndCostsIntegrationDataflowsRequest {
	this := AnthropicUsageAndCostsIntegrationDataflowsRequest{}
	return &this
}

// NewAnthropicUsageAndCostsIntegrationDataflowsRequestWithDefaults instantiates a new AnthropicUsageAndCostsIntegrationDataflowsRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAnthropicUsageAndCostsIntegrationDataflowsRequestWithDefaults() *AnthropicUsageAndCostsIntegrationDataflowsRequest {
	this := AnthropicUsageAndCostsIntegrationDataflowsRequest{}
	return &this
}

// GetAnthropicCloudCostMetrics returns the AnthropicCloudCostMetrics field value if set, zero value otherwise.
func (o *AnthropicUsageAndCostsIntegrationDataflowsRequest) GetAnthropicCloudCostMetrics() AnthropicUsageAndCostsCloudCostMetricsIntegrationDataflowRequest {
	if o == nil || o.AnthropicCloudCostMetrics == nil {
		var ret AnthropicUsageAndCostsCloudCostMetricsIntegrationDataflowRequest
		return ret
	}
	return *o.AnthropicCloudCostMetrics
}

// GetAnthropicCloudCostMetricsOk returns a tuple with the AnthropicCloudCostMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnthropicUsageAndCostsIntegrationDataflowsRequest) GetAnthropicCloudCostMetricsOk() (*AnthropicUsageAndCostsCloudCostMetricsIntegrationDataflowRequest, bool) {
	if o == nil || o.AnthropicCloudCostMetrics == nil {
		return nil, false
	}
	return o.AnthropicCloudCostMetrics, true
}

// HasAnthropicCloudCostMetrics returns a boolean if a field has been set.
func (o *AnthropicUsageAndCostsIntegrationDataflowsRequest) HasAnthropicCloudCostMetrics() bool {
	return o != nil && o.AnthropicCloudCostMetrics != nil
}

// SetAnthropicCloudCostMetrics gets a reference to the given AnthropicUsageAndCostsCloudCostMetricsIntegrationDataflowRequest and assigns it to the AnthropicCloudCostMetrics field.
func (o *AnthropicUsageAndCostsIntegrationDataflowsRequest) SetAnthropicCloudCostMetrics(v AnthropicUsageAndCostsCloudCostMetricsIntegrationDataflowRequest) {
	o.AnthropicCloudCostMetrics = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AnthropicUsageAndCostsIntegrationDataflowsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AnthropicCloudCostMetrics != nil {
		toSerialize["anthropic-cloud-cost-metrics"] = o.AnthropicCloudCostMetrics
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AnthropicUsageAndCostsIntegrationDataflowsRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AnthropicCloudCostMetrics *AnthropicUsageAndCostsCloudCostMetricsIntegrationDataflowRequest `json:"anthropic-cloud-cost-metrics,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	hasInvalidField := false
	if all.AnthropicCloudCostMetrics != nil && all.AnthropicCloudCostMetrics.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AnthropicCloudCostMetrics = all.AnthropicCloudCostMetrics

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
