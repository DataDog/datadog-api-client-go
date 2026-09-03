// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatabricksIntegrationDataflowsResponse Dataflows configured on the Databricks integration account, keyed by dataflow id.
type DatabricksIntegrationDataflowsResponse struct {
	// The Databricks cloud cost metrics dataflow.
	DatabricksCloudCostMetrics *DatabricksCloudCostMetricsIntegrationDataflowResponse `json:"databricks-cloud-cost-metrics,omitempty"`
	// The Databricks Data Jobs Monitoring dataflow.
	DatabricksDataJobMonitoring *DatabricksDataJobMonitoringIntegrationDataflowResponse `json:"databricks-data-job-monitoring,omitempty"`
	// The Databricks data observability dataflow.
	DatabricksDataObservability *DatabricksDataObservabilityIntegrationDataflowResponse `json:"databricks-data-observability,omitempty"`
	// The Databricks model serving metrics dataflow. Not supported on accounts that authenticate with `private-action-runner`; on those accounts this dataflow collects no data even when enabled.
	DatabricksModelServingMetrics *DatabricksModelServingMetricsIntegrationDataflowResponse `json:"databricks-model-serving-metrics,omitempty"`
	// The Databricks serverless jobs dataflow.
	DatabricksServerlessJobs *DatabricksServerlessJobsIntegrationDataflowResponse `json:"databricks-serverless-jobs,omitempty"`
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

// GetDatabricksDataJobMonitoring returns the DatabricksDataJobMonitoring field value if set, zero value otherwise.
func (o *DatabricksIntegrationDataflowsResponse) GetDatabricksDataJobMonitoring() DatabricksDataJobMonitoringIntegrationDataflowResponse {
	if o == nil || o.DatabricksDataJobMonitoring == nil {
		var ret DatabricksDataJobMonitoringIntegrationDataflowResponse
		return ret
	}
	return *o.DatabricksDataJobMonitoring
}

// GetDatabricksDataJobMonitoringOk returns a tuple with the DatabricksDataJobMonitoring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksIntegrationDataflowsResponse) GetDatabricksDataJobMonitoringOk() (*DatabricksDataJobMonitoringIntegrationDataflowResponse, bool) {
	if o == nil || o.DatabricksDataJobMonitoring == nil {
		return nil, false
	}
	return o.DatabricksDataJobMonitoring, true
}

// HasDatabricksDataJobMonitoring returns a boolean if a field has been set.
func (o *DatabricksIntegrationDataflowsResponse) HasDatabricksDataJobMonitoring() bool {
	return o != nil && o.DatabricksDataJobMonitoring != nil
}

// SetDatabricksDataJobMonitoring gets a reference to the given DatabricksDataJobMonitoringIntegrationDataflowResponse and assigns it to the DatabricksDataJobMonitoring field.
func (o *DatabricksIntegrationDataflowsResponse) SetDatabricksDataJobMonitoring(v DatabricksDataJobMonitoringIntegrationDataflowResponse) {
	o.DatabricksDataJobMonitoring = &v
}

// GetDatabricksDataObservability returns the DatabricksDataObservability field value if set, zero value otherwise.
func (o *DatabricksIntegrationDataflowsResponse) GetDatabricksDataObservability() DatabricksDataObservabilityIntegrationDataflowResponse {
	if o == nil || o.DatabricksDataObservability == nil {
		var ret DatabricksDataObservabilityIntegrationDataflowResponse
		return ret
	}
	return *o.DatabricksDataObservability
}

// GetDatabricksDataObservabilityOk returns a tuple with the DatabricksDataObservability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksIntegrationDataflowsResponse) GetDatabricksDataObservabilityOk() (*DatabricksDataObservabilityIntegrationDataflowResponse, bool) {
	if o == nil || o.DatabricksDataObservability == nil {
		return nil, false
	}
	return o.DatabricksDataObservability, true
}

// HasDatabricksDataObservability returns a boolean if a field has been set.
func (o *DatabricksIntegrationDataflowsResponse) HasDatabricksDataObservability() bool {
	return o != nil && o.DatabricksDataObservability != nil
}

// SetDatabricksDataObservability gets a reference to the given DatabricksDataObservabilityIntegrationDataflowResponse and assigns it to the DatabricksDataObservability field.
func (o *DatabricksIntegrationDataflowsResponse) SetDatabricksDataObservability(v DatabricksDataObservabilityIntegrationDataflowResponse) {
	o.DatabricksDataObservability = &v
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

// GetDatabricksServerlessJobs returns the DatabricksServerlessJobs field value if set, zero value otherwise.
func (o *DatabricksIntegrationDataflowsResponse) GetDatabricksServerlessJobs() DatabricksServerlessJobsIntegrationDataflowResponse {
	if o == nil || o.DatabricksServerlessJobs == nil {
		var ret DatabricksServerlessJobsIntegrationDataflowResponse
		return ret
	}
	return *o.DatabricksServerlessJobs
}

// GetDatabricksServerlessJobsOk returns a tuple with the DatabricksServerlessJobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksIntegrationDataflowsResponse) GetDatabricksServerlessJobsOk() (*DatabricksServerlessJobsIntegrationDataflowResponse, bool) {
	if o == nil || o.DatabricksServerlessJobs == nil {
		return nil, false
	}
	return o.DatabricksServerlessJobs, true
}

// HasDatabricksServerlessJobs returns a boolean if a field has been set.
func (o *DatabricksIntegrationDataflowsResponse) HasDatabricksServerlessJobs() bool {
	return o != nil && o.DatabricksServerlessJobs != nil
}

// SetDatabricksServerlessJobs gets a reference to the given DatabricksServerlessJobsIntegrationDataflowResponse and assigns it to the DatabricksServerlessJobs field.
func (o *DatabricksIntegrationDataflowsResponse) SetDatabricksServerlessJobs(v DatabricksServerlessJobsIntegrationDataflowResponse) {
	o.DatabricksServerlessJobs = &v
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DatabricksIntegrationDataflowsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DatabricksCloudCostMetrics    *DatabricksCloudCostMetricsIntegrationDataflowResponse    `json:"databricks-cloud-cost-metrics,omitempty"`
		DatabricksDataJobMonitoring   *DatabricksDataJobMonitoringIntegrationDataflowResponse   `json:"databricks-data-job-monitoring,omitempty"`
		DatabricksDataObservability   *DatabricksDataObservabilityIntegrationDataflowResponse   `json:"databricks-data-observability,omitempty"`
		DatabricksModelServingMetrics *DatabricksModelServingMetricsIntegrationDataflowResponse `json:"databricks-model-serving-metrics,omitempty"`
		DatabricksServerlessJobs      *DatabricksServerlessJobsIntegrationDataflowResponse      `json:"databricks-serverless-jobs,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"databricks-cloud-cost-metrics", "databricks-data-job-monitoring", "databricks-data-observability", "databricks-model-serving-metrics", "databricks-serverless-jobs"})
	} else {
		return err
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

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
