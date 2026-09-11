// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatabricksIntegrationDataflowsRequest Data Datadog collects from Databricks, keyed by dataflow id. Each dataflow turns on a distinct kind of collection: set `enabled` to start or stop it, and use `settings` to tune what it gathers. The defaults noted below apply when the account is created; on update, anything left out keeps its current value. Some dataflows have prerequisites, noted on each; unless one is documented as rejecting the request, it is not verified, so a dataflow enabled without it is stored but collects no data.
type DatabricksIntegrationDataflowsRequest struct {
	// Cost data collected from your Databricks system tables. [Cloud Cost Management](https://docs.datadoghq.com/cloud_cost_management/) must be enabled for your organization while this dataflow is on; any request that leaves it enabled without that is rejected with a `422` response.
	DatabricksCloudCostMetrics *DatabricksCloudCostMetricsIntegrationDataflowRequest `json:"databricks-cloud-cost-metrics,omitempty"`
	// Data Jobs Monitoring, which collects performance, reliability, and cost data for your Databricks jobs.
	DatabricksDataObservabilityJobsMonitoring *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowRequest `json:"databricks-data-observability-jobs-monitoring,omitempty"`
	// Data Observability, which collects lineage and data quality information from your Databricks catalogs so you can explore how data flows and detect, resolve, and prevent quality issues.
	DatabricksDataObservabilityQualityMonitoring *DatabricksDataObservabilityQualityMonitoringIntegrationDataflowRequest `json:"databricks-data-observability-quality-monitoring,omitempty"`
	// Health and usage metrics for your Databricks model serving endpoints. Not supported on accounts that authenticate with `private-action-runner`; on those accounts this dataflow collects no data even when enabled.
	DatabricksModelServingMetrics *DatabricksModelServingMetricsIntegrationDataflowRequest `json:"databricks-model-serving-metrics,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewDatabricksIntegrationDataflowsRequest instantiates a new DatabricksIntegrationDataflowsRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatabricksIntegrationDataflowsRequest() *DatabricksIntegrationDataflowsRequest {
	this := DatabricksIntegrationDataflowsRequest{}
	return &this
}

// NewDatabricksIntegrationDataflowsRequestWithDefaults instantiates a new DatabricksIntegrationDataflowsRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatabricksIntegrationDataflowsRequestWithDefaults() *DatabricksIntegrationDataflowsRequest {
	this := DatabricksIntegrationDataflowsRequest{}
	return &this
}

// GetDatabricksCloudCostMetrics returns the DatabricksCloudCostMetrics field value if set, zero value otherwise.
func (o *DatabricksIntegrationDataflowsRequest) GetDatabricksCloudCostMetrics() DatabricksCloudCostMetricsIntegrationDataflowRequest {
	if o == nil || o.DatabricksCloudCostMetrics == nil {
		var ret DatabricksCloudCostMetricsIntegrationDataflowRequest
		return ret
	}
	return *o.DatabricksCloudCostMetrics
}

// GetDatabricksCloudCostMetricsOk returns a tuple with the DatabricksCloudCostMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksIntegrationDataflowsRequest) GetDatabricksCloudCostMetricsOk() (*DatabricksCloudCostMetricsIntegrationDataflowRequest, bool) {
	if o == nil || o.DatabricksCloudCostMetrics == nil {
		return nil, false
	}
	return o.DatabricksCloudCostMetrics, true
}

// HasDatabricksCloudCostMetrics returns a boolean if a field has been set.
func (o *DatabricksIntegrationDataflowsRequest) HasDatabricksCloudCostMetrics() bool {
	return o != nil && o.DatabricksCloudCostMetrics != nil
}

// SetDatabricksCloudCostMetrics gets a reference to the given DatabricksCloudCostMetricsIntegrationDataflowRequest and assigns it to the DatabricksCloudCostMetrics field.
func (o *DatabricksIntegrationDataflowsRequest) SetDatabricksCloudCostMetrics(v DatabricksCloudCostMetricsIntegrationDataflowRequest) {
	o.DatabricksCloudCostMetrics = &v
}

// GetDatabricksDataObservabilityJobsMonitoring returns the DatabricksDataObservabilityJobsMonitoring field value if set, zero value otherwise.
func (o *DatabricksIntegrationDataflowsRequest) GetDatabricksDataObservabilityJobsMonitoring() DatabricksDataObservabilityJobsMonitoringIntegrationDataflowRequest {
	if o == nil || o.DatabricksDataObservabilityJobsMonitoring == nil {
		var ret DatabricksDataObservabilityJobsMonitoringIntegrationDataflowRequest
		return ret
	}
	return *o.DatabricksDataObservabilityJobsMonitoring
}

// GetDatabricksDataObservabilityJobsMonitoringOk returns a tuple with the DatabricksDataObservabilityJobsMonitoring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksIntegrationDataflowsRequest) GetDatabricksDataObservabilityJobsMonitoringOk() (*DatabricksDataObservabilityJobsMonitoringIntegrationDataflowRequest, bool) {
	if o == nil || o.DatabricksDataObservabilityJobsMonitoring == nil {
		return nil, false
	}
	return o.DatabricksDataObservabilityJobsMonitoring, true
}

// HasDatabricksDataObservabilityJobsMonitoring returns a boolean if a field has been set.
func (o *DatabricksIntegrationDataflowsRequest) HasDatabricksDataObservabilityJobsMonitoring() bool {
	return o != nil && o.DatabricksDataObservabilityJobsMonitoring != nil
}

// SetDatabricksDataObservabilityJobsMonitoring gets a reference to the given DatabricksDataObservabilityJobsMonitoringIntegrationDataflowRequest and assigns it to the DatabricksDataObservabilityJobsMonitoring field.
func (o *DatabricksIntegrationDataflowsRequest) SetDatabricksDataObservabilityJobsMonitoring(v DatabricksDataObservabilityJobsMonitoringIntegrationDataflowRequest) {
	o.DatabricksDataObservabilityJobsMonitoring = &v
}

// GetDatabricksDataObservabilityQualityMonitoring returns the DatabricksDataObservabilityQualityMonitoring field value if set, zero value otherwise.
func (o *DatabricksIntegrationDataflowsRequest) GetDatabricksDataObservabilityQualityMonitoring() DatabricksDataObservabilityQualityMonitoringIntegrationDataflowRequest {
	if o == nil || o.DatabricksDataObservabilityQualityMonitoring == nil {
		var ret DatabricksDataObservabilityQualityMonitoringIntegrationDataflowRequest
		return ret
	}
	return *o.DatabricksDataObservabilityQualityMonitoring
}

// GetDatabricksDataObservabilityQualityMonitoringOk returns a tuple with the DatabricksDataObservabilityQualityMonitoring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksIntegrationDataflowsRequest) GetDatabricksDataObservabilityQualityMonitoringOk() (*DatabricksDataObservabilityQualityMonitoringIntegrationDataflowRequest, bool) {
	if o == nil || o.DatabricksDataObservabilityQualityMonitoring == nil {
		return nil, false
	}
	return o.DatabricksDataObservabilityQualityMonitoring, true
}

// HasDatabricksDataObservabilityQualityMonitoring returns a boolean if a field has been set.
func (o *DatabricksIntegrationDataflowsRequest) HasDatabricksDataObservabilityQualityMonitoring() bool {
	return o != nil && o.DatabricksDataObservabilityQualityMonitoring != nil
}

// SetDatabricksDataObservabilityQualityMonitoring gets a reference to the given DatabricksDataObservabilityQualityMonitoringIntegrationDataflowRequest and assigns it to the DatabricksDataObservabilityQualityMonitoring field.
func (o *DatabricksIntegrationDataflowsRequest) SetDatabricksDataObservabilityQualityMonitoring(v DatabricksDataObservabilityQualityMonitoringIntegrationDataflowRequest) {
	o.DatabricksDataObservabilityQualityMonitoring = &v
}

// GetDatabricksModelServingMetrics returns the DatabricksModelServingMetrics field value if set, zero value otherwise.
func (o *DatabricksIntegrationDataflowsRequest) GetDatabricksModelServingMetrics() DatabricksModelServingMetricsIntegrationDataflowRequest {
	if o == nil || o.DatabricksModelServingMetrics == nil {
		var ret DatabricksModelServingMetricsIntegrationDataflowRequest
		return ret
	}
	return *o.DatabricksModelServingMetrics
}

// GetDatabricksModelServingMetricsOk returns a tuple with the DatabricksModelServingMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksIntegrationDataflowsRequest) GetDatabricksModelServingMetricsOk() (*DatabricksModelServingMetricsIntegrationDataflowRequest, bool) {
	if o == nil || o.DatabricksModelServingMetrics == nil {
		return nil, false
	}
	return o.DatabricksModelServingMetrics, true
}

// HasDatabricksModelServingMetrics returns a boolean if a field has been set.
func (o *DatabricksIntegrationDataflowsRequest) HasDatabricksModelServingMetrics() bool {
	return o != nil && o.DatabricksModelServingMetrics != nil
}

// SetDatabricksModelServingMetrics gets a reference to the given DatabricksModelServingMetricsIntegrationDataflowRequest and assigns it to the DatabricksModelServingMetrics field.
func (o *DatabricksIntegrationDataflowsRequest) SetDatabricksModelServingMetrics(v DatabricksModelServingMetricsIntegrationDataflowRequest) {
	o.DatabricksModelServingMetrics = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatabricksIntegrationDataflowsRequest) MarshalJSON() ([]byte, error) {
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
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DatabricksIntegrationDataflowsRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DatabricksCloudCostMetrics                   *DatabricksCloudCostMetricsIntegrationDataflowRequest                   `json:"databricks-cloud-cost-metrics,omitempty"`
		DatabricksDataObservabilityJobsMonitoring    *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowRequest    `json:"databricks-data-observability-jobs-monitoring,omitempty"`
		DatabricksDataObservabilityQualityMonitoring *DatabricksDataObservabilityQualityMonitoringIntegrationDataflowRequest `json:"databricks-data-observability-quality-monitoring,omitempty"`
		DatabricksModelServingMetrics                *DatabricksModelServingMetricsIntegrationDataflowRequest                `json:"databricks-model-serving-metrics,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
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

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
