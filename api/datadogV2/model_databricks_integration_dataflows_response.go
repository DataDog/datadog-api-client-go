// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatabricksIntegrationDataflowsResponse Data Datadog collects from Databricks, keyed by dataflow id.
type DatabricksIntegrationDataflowsResponse struct {
	// Cost data collected from your Databricks system tables. Requires [Cloud Cost Management](https://docs.datadoghq.com/cloud_cost_management/) to be set up for your organization.
	DatabricksCloudCostMetrics *DatabricksCloudCostMetricsIntegrationDataflowResponse `json:"databricks-cloud-cost-metrics,omitempty"`
	// Data Jobs Monitoring, which collects performance, reliability, and cost data for your Databricks jobs.
	DatabricksDataObservabilityJobsMonitoring *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowResponse `json:"databricks-data-observability-jobs-monitoring,omitempty"`
	// Data Observability, which collects lineage and data quality information from your Databricks catalogs so you can explore how data flows and detect, resolve, and prevent quality issues.
	DatabricksDataObservabilityQualityMonitoring *DatabricksDataObservabilityQualityMonitoringIntegrationDataflowResponse `json:"databricks-data-observability-quality-monitoring,omitempty"`
	// Health and usage metrics for your Databricks model serving endpoints. Not supported on accounts that authenticate with `private-action-runner`; on those accounts this dataflow collects no data even when enabled.
	DatabricksModelServingMetrics *DatabricksModelServingMetricsIntegrationDataflowResponse `json:"databricks-model-serving-metrics,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDatabricksIntegrationDataflowsResponse instantiates a new DatabricksIntegrationDataflowsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatabricksIntegrationDataflowsResponse() *DatabricksIntegrationDataflowsResponse {
	this := DatabricksIntegrationDataflowsResponse{}
	return &this
}

// NewDatabricksIntegrationDataflowsResponseWithDefaults instantiates a new DatabricksIntegrationDataflowsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatabricksIntegrationDataflowsResponseWithDefaults() *DatabricksIntegrationDataflowsResponse {
	this := DatabricksIntegrationDataflowsResponse{}
	return &this
}

// GetDatabricksCloudCostMetrics returns the DatabricksCloudCostMetrics field value if set, zero value otherwise.
func (o *DatabricksIntegrationDataflowsResponse) GetDatabricksCloudCostMetrics() DatabricksCloudCostMetricsIntegrationDataflowResponse {
	if o == nil || o.DatabricksCloudCostMetrics == nil {
		var ret DatabricksCloudCostMetricsIntegrationDataflowResponse
		return ret
	}
	return *o.DatabricksCloudCostMetrics
}

// GetDatabricksCloudCostMetricsOk returns a tuple with the DatabricksCloudCostMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksIntegrationDataflowsResponse) GetDatabricksCloudCostMetricsOk() (*DatabricksCloudCostMetricsIntegrationDataflowResponse, bool) {
	if o == nil || o.DatabricksCloudCostMetrics == nil {
		return nil, false
	}
	return o.DatabricksCloudCostMetrics, true
}

// HasDatabricksCloudCostMetrics returns a boolean if a field has been set.
func (o *DatabricksIntegrationDataflowsResponse) HasDatabricksCloudCostMetrics() bool {
	return o != nil && o.DatabricksCloudCostMetrics != nil
}

// SetDatabricksCloudCostMetrics gets a reference to the given DatabricksCloudCostMetricsIntegrationDataflowResponse and assigns it to the DatabricksCloudCostMetrics field.
func (o *DatabricksIntegrationDataflowsResponse) SetDatabricksCloudCostMetrics(v DatabricksCloudCostMetricsIntegrationDataflowResponse) {
	o.DatabricksCloudCostMetrics = &v
}

// GetDatabricksDataObservabilityJobsMonitoring returns the DatabricksDataObservabilityJobsMonitoring field value if set, zero value otherwise.
func (o *DatabricksIntegrationDataflowsResponse) GetDatabricksDataObservabilityJobsMonitoring() DatabricksDataObservabilityJobsMonitoringIntegrationDataflowResponse {
	if o == nil || o.DatabricksDataObservabilityJobsMonitoring == nil {
		var ret DatabricksDataObservabilityJobsMonitoringIntegrationDataflowResponse
		return ret
	}
	return *o.DatabricksDataObservabilityJobsMonitoring
}

// GetDatabricksDataObservabilityJobsMonitoringOk returns a tuple with the DatabricksDataObservabilityJobsMonitoring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksIntegrationDataflowsResponse) GetDatabricksDataObservabilityJobsMonitoringOk() (*DatabricksDataObservabilityJobsMonitoringIntegrationDataflowResponse, bool) {
	if o == nil || o.DatabricksDataObservabilityJobsMonitoring == nil {
		return nil, false
	}
	return o.DatabricksDataObservabilityJobsMonitoring, true
}

// HasDatabricksDataObservabilityJobsMonitoring returns a boolean if a field has been set.
func (o *DatabricksIntegrationDataflowsResponse) HasDatabricksDataObservabilityJobsMonitoring() bool {
	return o != nil && o.DatabricksDataObservabilityJobsMonitoring != nil
}

// SetDatabricksDataObservabilityJobsMonitoring gets a reference to the given DatabricksDataObservabilityJobsMonitoringIntegrationDataflowResponse and assigns it to the DatabricksDataObservabilityJobsMonitoring field.
func (o *DatabricksIntegrationDataflowsResponse) SetDatabricksDataObservabilityJobsMonitoring(v DatabricksDataObservabilityJobsMonitoringIntegrationDataflowResponse) {
	o.DatabricksDataObservabilityJobsMonitoring = &v
}

// GetDatabricksDataObservabilityQualityMonitoring returns the DatabricksDataObservabilityQualityMonitoring field value if set, zero value otherwise.
func (o *DatabricksIntegrationDataflowsResponse) GetDatabricksDataObservabilityQualityMonitoring() DatabricksDataObservabilityQualityMonitoringIntegrationDataflowResponse {
	if o == nil || o.DatabricksDataObservabilityQualityMonitoring == nil {
		var ret DatabricksDataObservabilityQualityMonitoringIntegrationDataflowResponse
		return ret
	}
	return *o.DatabricksDataObservabilityQualityMonitoring
}

// GetDatabricksDataObservabilityQualityMonitoringOk returns a tuple with the DatabricksDataObservabilityQualityMonitoring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksIntegrationDataflowsResponse) GetDatabricksDataObservabilityQualityMonitoringOk() (*DatabricksDataObservabilityQualityMonitoringIntegrationDataflowResponse, bool) {
	if o == nil || o.DatabricksDataObservabilityQualityMonitoring == nil {
		return nil, false
	}
	return o.DatabricksDataObservabilityQualityMonitoring, true
}

// HasDatabricksDataObservabilityQualityMonitoring returns a boolean if a field has been set.
func (o *DatabricksIntegrationDataflowsResponse) HasDatabricksDataObservabilityQualityMonitoring() bool {
	return o != nil && o.DatabricksDataObservabilityQualityMonitoring != nil
}

// SetDatabricksDataObservabilityQualityMonitoring gets a reference to the given DatabricksDataObservabilityQualityMonitoringIntegrationDataflowResponse and assigns it to the DatabricksDataObservabilityQualityMonitoring field.
func (o *DatabricksIntegrationDataflowsResponse) SetDatabricksDataObservabilityQualityMonitoring(v DatabricksDataObservabilityQualityMonitoringIntegrationDataflowResponse) {
	o.DatabricksDataObservabilityQualityMonitoring = &v
}

// GetDatabricksModelServingMetrics returns the DatabricksModelServingMetrics field value if set, zero value otherwise.
func (o *DatabricksIntegrationDataflowsResponse) GetDatabricksModelServingMetrics() DatabricksModelServingMetricsIntegrationDataflowResponse {
	if o == nil || o.DatabricksModelServingMetrics == nil {
		var ret DatabricksModelServingMetricsIntegrationDataflowResponse
		return ret
	}
	return *o.DatabricksModelServingMetrics
}

// GetDatabricksModelServingMetricsOk returns a tuple with the DatabricksModelServingMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksIntegrationDataflowsResponse) GetDatabricksModelServingMetricsOk() (*DatabricksModelServingMetricsIntegrationDataflowResponse, bool) {
	if o == nil || o.DatabricksModelServingMetrics == nil {
		return nil, false
	}
	return o.DatabricksModelServingMetrics, true
}

// HasDatabricksModelServingMetrics returns a boolean if a field has been set.
func (o *DatabricksIntegrationDataflowsResponse) HasDatabricksModelServingMetrics() bool {
	return o != nil && o.DatabricksModelServingMetrics != nil
}

// SetDatabricksModelServingMetrics gets a reference to the given DatabricksModelServingMetricsIntegrationDataflowResponse and assigns it to the DatabricksModelServingMetrics field.
func (o *DatabricksIntegrationDataflowsResponse) SetDatabricksModelServingMetrics(v DatabricksModelServingMetricsIntegrationDataflowResponse) {
	o.DatabricksModelServingMetrics = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatabricksIntegrationDataflowsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DatabricksCloudCostMetrics != nil {
		toSerialize["databricks-cloud-cost-metrics"] = o.DatabricksCloudCostMetrics
	}
	if o.DatabricksDataObservabilityJobsMonitoring != nil {
		toSerialize["databricks-data-observability-jobs-monitoring"] = o.DatabricksDataObservabilityJobsMonitoring
	}
	if o.DatabricksDataObservabilityQualityMonitoring != nil {
		toSerialize["databricks-data-observability-quality-monitoring"] = o.DatabricksDataObservabilityQualityMonitoring
	}
	if o.DatabricksModelServingMetrics != nil {
		toSerialize["databricks-model-serving-metrics"] = o.DatabricksModelServingMetrics
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DatabricksIntegrationDataflowsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DatabricksCloudCostMetrics                   *DatabricksCloudCostMetricsIntegrationDataflowResponse                   `json:"databricks-cloud-cost-metrics,omitempty"`
		DatabricksDataObservabilityJobsMonitoring    *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowResponse    `json:"databricks-data-observability-jobs-monitoring,omitempty"`
		DatabricksDataObservabilityQualityMonitoring *DatabricksDataObservabilityQualityMonitoringIntegrationDataflowResponse `json:"databricks-data-observability-quality-monitoring,omitempty"`
		DatabricksModelServingMetrics                *DatabricksModelServingMetricsIntegrationDataflowResponse                `json:"databricks-model-serving-metrics,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"databricks-cloud-cost-metrics", "databricks-data-observability-jobs-monitoring", "databricks-data-observability-quality-monitoring", "databricks-model-serving-metrics"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.DatabricksCloudCostMetrics != nil && all.DatabricksCloudCostMetrics.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DatabricksCloudCostMetrics = all.DatabricksCloudCostMetrics
	if all.DatabricksDataObservabilityJobsMonitoring != nil && all.DatabricksDataObservabilityJobsMonitoring.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DatabricksDataObservabilityJobsMonitoring = all.DatabricksDataObservabilityJobsMonitoring
	if all.DatabricksDataObservabilityQualityMonitoring != nil && all.DatabricksDataObservabilityQualityMonitoring.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DatabricksDataObservabilityQualityMonitoring = all.DatabricksDataObservabilityQualityMonitoring
	if all.DatabricksModelServingMetrics != nil && all.DatabricksModelServingMetrics.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DatabricksModelServingMetrics = all.DatabricksModelServingMetrics

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
