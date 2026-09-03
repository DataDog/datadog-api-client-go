// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatabricksIntegrationDataflowsRequest Dataflows to configure on the Databricks integration account, keyed by dataflow id. Some dataflows and settings have prerequisites, noted on each. Those prerequisites are not checked when the request is made, so anything left enabled without them is stored but collects no data.
type DatabricksIntegrationDataflowsRequest struct {
	// The Databricks cloud cost metrics dataflow.
	DatabricksCloudCostMetrics *DatabricksCloudCostMetricsIntegrationDataflowRequest `json:"databricks-cloud-cost-metrics,omitempty"`
	// The Databricks Data Jobs Monitoring dataflow.
	DatabricksDataJobMonitoring *DatabricksDataJobMonitoringIntegrationDataflowRequest `json:"databricks-data-job-monitoring,omitempty"`
	// The Databricks data observability dataflow.
	DatabricksDataObservability *DatabricksDataObservabilityIntegrationDataflowRequest `json:"databricks-data-observability,omitempty"`
	// The Databricks model serving metrics dataflow. Not supported on accounts that authenticate with `private-action-runner`; on those accounts this dataflow collects no data even when enabled.
	DatabricksModelServingMetrics *DatabricksModelServingMetricsIntegrationDataflowRequest `json:"databricks-model-serving-metrics,omitempty"`
	// The Databricks serverless jobs dataflow.
	DatabricksServerlessJobs *DatabricksServerlessJobsIntegrationDataflowRequest `json:"databricks-serverless-jobs,omitempty"`
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

// GetDatabricksDataJobMonitoring returns the DatabricksDataJobMonitoring field value if set, zero value otherwise.
func (o *DatabricksIntegrationDataflowsRequest) GetDatabricksDataJobMonitoring() DatabricksDataJobMonitoringIntegrationDataflowRequest {
	if o == nil || o.DatabricksDataJobMonitoring == nil {
		var ret DatabricksDataJobMonitoringIntegrationDataflowRequest
		return ret
	}
	return *o.DatabricksDataJobMonitoring
}

// GetDatabricksDataJobMonitoringOk returns a tuple with the DatabricksDataJobMonitoring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksIntegrationDataflowsRequest) GetDatabricksDataJobMonitoringOk() (*DatabricksDataJobMonitoringIntegrationDataflowRequest, bool) {
	if o == nil || o.DatabricksDataJobMonitoring == nil {
		return nil, false
	}
	return o.DatabricksDataJobMonitoring, true
}

// HasDatabricksDataJobMonitoring returns a boolean if a field has been set.
func (o *DatabricksIntegrationDataflowsRequest) HasDatabricksDataJobMonitoring() bool {
	return o != nil && o.DatabricksDataJobMonitoring != nil
}

// SetDatabricksDataJobMonitoring gets a reference to the given DatabricksDataJobMonitoringIntegrationDataflowRequest and assigns it to the DatabricksDataJobMonitoring field.
func (o *DatabricksIntegrationDataflowsRequest) SetDatabricksDataJobMonitoring(v DatabricksDataJobMonitoringIntegrationDataflowRequest) {
	o.DatabricksDataJobMonitoring = &v
}

// GetDatabricksDataObservability returns the DatabricksDataObservability field value if set, zero value otherwise.
func (o *DatabricksIntegrationDataflowsRequest) GetDatabricksDataObservability() DatabricksDataObservabilityIntegrationDataflowRequest {
	if o == nil || o.DatabricksDataObservability == nil {
		var ret DatabricksDataObservabilityIntegrationDataflowRequest
		return ret
	}
	return *o.DatabricksDataObservability
}

// GetDatabricksDataObservabilityOk returns a tuple with the DatabricksDataObservability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksIntegrationDataflowsRequest) GetDatabricksDataObservabilityOk() (*DatabricksDataObservabilityIntegrationDataflowRequest, bool) {
	if o == nil || o.DatabricksDataObservability == nil {
		return nil, false
	}
	return o.DatabricksDataObservability, true
}

// HasDatabricksDataObservability returns a boolean if a field has been set.
func (o *DatabricksIntegrationDataflowsRequest) HasDatabricksDataObservability() bool {
	return o != nil && o.DatabricksDataObservability != nil
}

// SetDatabricksDataObservability gets a reference to the given DatabricksDataObservabilityIntegrationDataflowRequest and assigns it to the DatabricksDataObservability field.
func (o *DatabricksIntegrationDataflowsRequest) SetDatabricksDataObservability(v DatabricksDataObservabilityIntegrationDataflowRequest) {
	o.DatabricksDataObservability = &v
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

// GetDatabricksServerlessJobs returns the DatabricksServerlessJobs field value if set, zero value otherwise.
func (o *DatabricksIntegrationDataflowsRequest) GetDatabricksServerlessJobs() DatabricksServerlessJobsIntegrationDataflowRequest {
	if o == nil || o.DatabricksServerlessJobs == nil {
		var ret DatabricksServerlessJobsIntegrationDataflowRequest
		return ret
	}
	return *o.DatabricksServerlessJobs
}

// GetDatabricksServerlessJobsOk returns a tuple with the DatabricksServerlessJobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksIntegrationDataflowsRequest) GetDatabricksServerlessJobsOk() (*DatabricksServerlessJobsIntegrationDataflowRequest, bool) {
	if o == nil || o.DatabricksServerlessJobs == nil {
		return nil, false
	}
	return o.DatabricksServerlessJobs, true
}

// HasDatabricksServerlessJobs returns a boolean if a field has been set.
func (o *DatabricksIntegrationDataflowsRequest) HasDatabricksServerlessJobs() bool {
	return o != nil && o.DatabricksServerlessJobs != nil
}

// SetDatabricksServerlessJobs gets a reference to the given DatabricksServerlessJobsIntegrationDataflowRequest and assigns it to the DatabricksServerlessJobs field.
func (o *DatabricksIntegrationDataflowsRequest) SetDatabricksServerlessJobs(v DatabricksServerlessJobsIntegrationDataflowRequest) {
	o.DatabricksServerlessJobs = &v
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
	if o.DatabricksDataJobMonitoring != nil {
		toSerialize["databricks-data-job-monitoring"] = o.DatabricksDataJobMonitoring
	}
	if o.DatabricksDataObservability != nil {
		toSerialize["databricks-data-observability"] = o.DatabricksDataObservability
	}
	if o.DatabricksModelServingMetrics != nil {
		toSerialize["databricks-model-serving-metrics"] = o.DatabricksModelServingMetrics
	}
	if o.DatabricksServerlessJobs != nil {
		toSerialize["databricks-serverless-jobs"] = o.DatabricksServerlessJobs
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DatabricksIntegrationDataflowsRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DatabricksCloudCostMetrics    *DatabricksCloudCostMetricsIntegrationDataflowRequest    `json:"databricks-cloud-cost-metrics,omitempty"`
		DatabricksDataJobMonitoring   *DatabricksDataJobMonitoringIntegrationDataflowRequest   `json:"databricks-data-job-monitoring,omitempty"`
		DatabricksDataObservability   *DatabricksDataObservabilityIntegrationDataflowRequest   `json:"databricks-data-observability,omitempty"`
		DatabricksModelServingMetrics *DatabricksModelServingMetricsIntegrationDataflowRequest `json:"databricks-model-serving-metrics,omitempty"`
		DatabricksServerlessJobs      *DatabricksServerlessJobsIntegrationDataflowRequest      `json:"databricks-serverless-jobs,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	hasInvalidField := false
	if all.DatabricksCloudCostMetrics != nil && all.DatabricksCloudCostMetrics.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DatabricksCloudCostMetrics = all.DatabricksCloudCostMetrics
	if all.DatabricksDataJobMonitoring != nil && all.DatabricksDataJobMonitoring.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DatabricksDataJobMonitoring = all.DatabricksDataJobMonitoring
	if all.DatabricksDataObservability != nil && all.DatabricksDataObservability.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DatabricksDataObservability = all.DatabricksDataObservability
	if all.DatabricksModelServingMetrics != nil && all.DatabricksModelServingMetrics.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DatabricksModelServingMetrics = all.DatabricksModelServingMetrics
	if all.DatabricksServerlessJobs != nil && all.DatabricksServerlessJobs.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DatabricksServerlessJobs = all.DatabricksServerlessJobs

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
