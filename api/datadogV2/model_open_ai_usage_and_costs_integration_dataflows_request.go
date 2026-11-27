// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpenAIUsageAndCostsIntegrationDataflowsRequest Data Datadog collects from OpenAI Usage and Costs, keyed by dataflow ID. Datadog always collects organization usage metrics. Each dataflow in this object adds another type of collection. Set `enabled` to start or stop a dataflow. The default for each dataflow applies when the account is created. On update, omitted fields keep their current values.
type OpenAIUsageAndCostsIntegrationDataflowsRequest struct {
	// OpenAI organization costs, which you can analyze in [Cloud Cost Management](https://docs.datadoghq.com/cloud_cost_management/). See [collected data](https://docs.datadoghq.com/cloud_cost_management/saas_costs/?tab=openai#data-collected) for what is included. Cost data appears in Cloud Cost Management within 24 hours after you enable this dataflow.
	OpenaiCloudCostData *OpenAIUsageAndCostsCloudCostDataIntegrationDataflowRequest `json:"openai-cloud-cost-data,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewOpenAIUsageAndCostsIntegrationDataflowsRequest instantiates a new OpenAIUsageAndCostsIntegrationDataflowsRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpenAIUsageAndCostsIntegrationDataflowsRequest() *OpenAIUsageAndCostsIntegrationDataflowsRequest {
	this := OpenAIUsageAndCostsIntegrationDataflowsRequest{}
	return &this
}

// NewOpenAIUsageAndCostsIntegrationDataflowsRequestWithDefaults instantiates a new OpenAIUsageAndCostsIntegrationDataflowsRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpenAIUsageAndCostsIntegrationDataflowsRequestWithDefaults() *OpenAIUsageAndCostsIntegrationDataflowsRequest {
	this := OpenAIUsageAndCostsIntegrationDataflowsRequest{}
	return &this
}

// GetOpenaiCloudCostData returns the OpenaiCloudCostData field value if set, zero value otherwise.
func (o *OpenAIUsageAndCostsIntegrationDataflowsRequest) GetOpenaiCloudCostData() OpenAIUsageAndCostsCloudCostDataIntegrationDataflowRequest {
	if o == nil || o.OpenaiCloudCostData == nil {
		var ret OpenAIUsageAndCostsCloudCostDataIntegrationDataflowRequest
		return ret
	}
	return *o.OpenaiCloudCostData
}

// GetOpenaiCloudCostDataOk returns a tuple with the OpenaiCloudCostData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAIUsageAndCostsIntegrationDataflowsRequest) GetOpenaiCloudCostDataOk() (*OpenAIUsageAndCostsCloudCostDataIntegrationDataflowRequest, bool) {
	if o == nil || o.OpenaiCloudCostData == nil {
		return nil, false
	}
	return o.OpenaiCloudCostData, true
}

// HasOpenaiCloudCostData returns a boolean if a field has been set.
func (o *OpenAIUsageAndCostsIntegrationDataflowsRequest) HasOpenaiCloudCostData() bool {
	return o != nil && o.OpenaiCloudCostData != nil
}

// SetOpenaiCloudCostData gets a reference to the given OpenAIUsageAndCostsCloudCostDataIntegrationDataflowRequest and assigns it to the OpenaiCloudCostData field.
func (o *OpenAIUsageAndCostsIntegrationDataflowsRequest) SetOpenaiCloudCostData(v OpenAIUsageAndCostsCloudCostDataIntegrationDataflowRequest) {
	o.OpenaiCloudCostData = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpenAIUsageAndCostsIntegrationDataflowsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.OpenaiCloudCostData != nil {
		toSerialize["openai-cloud-cost-data"] = o.OpenaiCloudCostData
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpenAIUsageAndCostsIntegrationDataflowsRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OpenaiCloudCostData *OpenAIUsageAndCostsCloudCostDataIntegrationDataflowRequest `json:"openai-cloud-cost-data,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	hasInvalidField := false
	if all.OpenaiCloudCostData != nil && all.OpenaiCloudCostData.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.OpenaiCloudCostData = all.OpenaiCloudCostData

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
