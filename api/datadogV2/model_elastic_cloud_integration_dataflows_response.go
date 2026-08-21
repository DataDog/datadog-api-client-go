// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ElasticCloudIntegrationDataflowsResponse Dataflows configured on the Elastic Cloud integration account, keyed by dataflow id.
type ElasticCloudIntegrationDataflowsResponse struct {
	// The Elastic Cloud detailed index stats dataflow.
	ElasticCloudDetailedIndexStats *ElasticCloudDetailedIndexStatsIntegrationDataflowResponse `json:"elastic-cloud-detailed-index-stats,omitempty"`
	// The Elastic Cloud index stats dataflow.
	ElasticCloudIndexStats *ElasticCloudIndexStatsIntegrationDataflowResponse `json:"elastic-cloud-index-stats,omitempty"`
	// The Elastic Cloud metrics dataflow.
	ElasticCloudMetrics *ElasticCloudMetricsIntegrationDataflowResponse `json:"elastic-cloud-metrics,omitempty"`
	// The Elastic Cloud pending task stats dataflow.
	ElasticCloudPendingTaskStats *ElasticCloudPendingTaskStatsIntegrationDataflowResponse `json:"elastic-cloud-pending-task-stats,omitempty"`
	// The Elastic Cloud primary shard graceful timeout dataflow.
	ElasticCloudPrimaryShardGracefulTimeout *ElasticCloudPrimaryShardGracefulTimeoutIntegrationDataflowResponse `json:"elastic-cloud-primary-shard-graceful-timeout,omitempty"`
	// The Elastic Cloud primary shard stats dataflow.
	ElasticCloudPrimaryShardStats *ElasticCloudPrimaryShardStatsIntegrationDataflowResponse `json:"elastic-cloud-primary-shard-stats,omitempty"`
	// The Elastic Cloud shard allocation stats dataflow.
	ElasticCloudShardAllocationStats *ElasticCloudShardAllocationStatsIntegrationDataflowResponse `json:"elastic-cloud-shard-allocation-stats,omitempty"`
	// The Elastic Cloud snapshot lifecycle management stats dataflow.
	ElasticCloudSlmStats *ElasticCloudSlmStatsIntegrationDataflowResponse `json:"elastic-cloud-slm-stats,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewElasticCloudIntegrationDataflowsResponse instantiates a new ElasticCloudIntegrationDataflowsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewElasticCloudIntegrationDataflowsResponse() *ElasticCloudIntegrationDataflowsResponse {
	this := ElasticCloudIntegrationDataflowsResponse{}
	return &this
}

// NewElasticCloudIntegrationDataflowsResponseWithDefaults instantiates a new ElasticCloudIntegrationDataflowsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewElasticCloudIntegrationDataflowsResponseWithDefaults() *ElasticCloudIntegrationDataflowsResponse {
	this := ElasticCloudIntegrationDataflowsResponse{}
	return &this
}

// GetElasticCloudDetailedIndexStats returns the ElasticCloudDetailedIndexStats field value if set, zero value otherwise.
func (o *ElasticCloudIntegrationDataflowsResponse) GetElasticCloudDetailedIndexStats() ElasticCloudDetailedIndexStatsIntegrationDataflowResponse {
	if o == nil || o.ElasticCloudDetailedIndexStats == nil {
		var ret ElasticCloudDetailedIndexStatsIntegrationDataflowResponse
		return ret
	}
	return *o.ElasticCloudDetailedIndexStats
}

// GetElasticCloudDetailedIndexStatsOk returns a tuple with the ElasticCloudDetailedIndexStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticCloudIntegrationDataflowsResponse) GetElasticCloudDetailedIndexStatsOk() (*ElasticCloudDetailedIndexStatsIntegrationDataflowResponse, bool) {
	if o == nil || o.ElasticCloudDetailedIndexStats == nil {
		return nil, false
	}
	return o.ElasticCloudDetailedIndexStats, true
}

// HasElasticCloudDetailedIndexStats returns a boolean if a field has been set.
func (o *ElasticCloudIntegrationDataflowsResponse) HasElasticCloudDetailedIndexStats() bool {
	return o != nil && o.ElasticCloudDetailedIndexStats != nil
}

// SetElasticCloudDetailedIndexStats gets a reference to the given ElasticCloudDetailedIndexStatsIntegrationDataflowResponse and assigns it to the ElasticCloudDetailedIndexStats field.
func (o *ElasticCloudIntegrationDataflowsResponse) SetElasticCloudDetailedIndexStats(v ElasticCloudDetailedIndexStatsIntegrationDataflowResponse) {
	o.ElasticCloudDetailedIndexStats = &v
}

// GetElasticCloudIndexStats returns the ElasticCloudIndexStats field value if set, zero value otherwise.
func (o *ElasticCloudIntegrationDataflowsResponse) GetElasticCloudIndexStats() ElasticCloudIndexStatsIntegrationDataflowResponse {
	if o == nil || o.ElasticCloudIndexStats == nil {
		var ret ElasticCloudIndexStatsIntegrationDataflowResponse
		return ret
	}
	return *o.ElasticCloudIndexStats
}

// GetElasticCloudIndexStatsOk returns a tuple with the ElasticCloudIndexStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticCloudIntegrationDataflowsResponse) GetElasticCloudIndexStatsOk() (*ElasticCloudIndexStatsIntegrationDataflowResponse, bool) {
	if o == nil || o.ElasticCloudIndexStats == nil {
		return nil, false
	}
	return o.ElasticCloudIndexStats, true
}

// HasElasticCloudIndexStats returns a boolean if a field has been set.
func (o *ElasticCloudIntegrationDataflowsResponse) HasElasticCloudIndexStats() bool {
	return o != nil && o.ElasticCloudIndexStats != nil
}

// SetElasticCloudIndexStats gets a reference to the given ElasticCloudIndexStatsIntegrationDataflowResponse and assigns it to the ElasticCloudIndexStats field.
func (o *ElasticCloudIntegrationDataflowsResponse) SetElasticCloudIndexStats(v ElasticCloudIndexStatsIntegrationDataflowResponse) {
	o.ElasticCloudIndexStats = &v
}

// GetElasticCloudMetrics returns the ElasticCloudMetrics field value if set, zero value otherwise.
func (o *ElasticCloudIntegrationDataflowsResponse) GetElasticCloudMetrics() ElasticCloudMetricsIntegrationDataflowResponse {
	if o == nil || o.ElasticCloudMetrics == nil {
		var ret ElasticCloudMetricsIntegrationDataflowResponse
		return ret
	}
	return *o.ElasticCloudMetrics
}

// GetElasticCloudMetricsOk returns a tuple with the ElasticCloudMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticCloudIntegrationDataflowsResponse) GetElasticCloudMetricsOk() (*ElasticCloudMetricsIntegrationDataflowResponse, bool) {
	if o == nil || o.ElasticCloudMetrics == nil {
		return nil, false
	}
	return o.ElasticCloudMetrics, true
}

// HasElasticCloudMetrics returns a boolean if a field has been set.
func (o *ElasticCloudIntegrationDataflowsResponse) HasElasticCloudMetrics() bool {
	return o != nil && o.ElasticCloudMetrics != nil
}

// SetElasticCloudMetrics gets a reference to the given ElasticCloudMetricsIntegrationDataflowResponse and assigns it to the ElasticCloudMetrics field.
func (o *ElasticCloudIntegrationDataflowsResponse) SetElasticCloudMetrics(v ElasticCloudMetricsIntegrationDataflowResponse) {
	o.ElasticCloudMetrics = &v
}

// GetElasticCloudPendingTaskStats returns the ElasticCloudPendingTaskStats field value if set, zero value otherwise.
func (o *ElasticCloudIntegrationDataflowsResponse) GetElasticCloudPendingTaskStats() ElasticCloudPendingTaskStatsIntegrationDataflowResponse {
	if o == nil || o.ElasticCloudPendingTaskStats == nil {
		var ret ElasticCloudPendingTaskStatsIntegrationDataflowResponse
		return ret
	}
	return *o.ElasticCloudPendingTaskStats
}

// GetElasticCloudPendingTaskStatsOk returns a tuple with the ElasticCloudPendingTaskStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticCloudIntegrationDataflowsResponse) GetElasticCloudPendingTaskStatsOk() (*ElasticCloudPendingTaskStatsIntegrationDataflowResponse, bool) {
	if o == nil || o.ElasticCloudPendingTaskStats == nil {
		return nil, false
	}
	return o.ElasticCloudPendingTaskStats, true
}

// HasElasticCloudPendingTaskStats returns a boolean if a field has been set.
func (o *ElasticCloudIntegrationDataflowsResponse) HasElasticCloudPendingTaskStats() bool {
	return o != nil && o.ElasticCloudPendingTaskStats != nil
}

// SetElasticCloudPendingTaskStats gets a reference to the given ElasticCloudPendingTaskStatsIntegrationDataflowResponse and assigns it to the ElasticCloudPendingTaskStats field.
func (o *ElasticCloudIntegrationDataflowsResponse) SetElasticCloudPendingTaskStats(v ElasticCloudPendingTaskStatsIntegrationDataflowResponse) {
	o.ElasticCloudPendingTaskStats = &v
}

// GetElasticCloudPrimaryShardGracefulTimeout returns the ElasticCloudPrimaryShardGracefulTimeout field value if set, zero value otherwise.
func (o *ElasticCloudIntegrationDataflowsResponse) GetElasticCloudPrimaryShardGracefulTimeout() ElasticCloudPrimaryShardGracefulTimeoutIntegrationDataflowResponse {
	if o == nil || o.ElasticCloudPrimaryShardGracefulTimeout == nil {
		var ret ElasticCloudPrimaryShardGracefulTimeoutIntegrationDataflowResponse
		return ret
	}
	return *o.ElasticCloudPrimaryShardGracefulTimeout
}

// GetElasticCloudPrimaryShardGracefulTimeoutOk returns a tuple with the ElasticCloudPrimaryShardGracefulTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticCloudIntegrationDataflowsResponse) GetElasticCloudPrimaryShardGracefulTimeoutOk() (*ElasticCloudPrimaryShardGracefulTimeoutIntegrationDataflowResponse, bool) {
	if o == nil || o.ElasticCloudPrimaryShardGracefulTimeout == nil {
		return nil, false
	}
	return o.ElasticCloudPrimaryShardGracefulTimeout, true
}

// HasElasticCloudPrimaryShardGracefulTimeout returns a boolean if a field has been set.
func (o *ElasticCloudIntegrationDataflowsResponse) HasElasticCloudPrimaryShardGracefulTimeout() bool {
	return o != nil && o.ElasticCloudPrimaryShardGracefulTimeout != nil
}

// SetElasticCloudPrimaryShardGracefulTimeout gets a reference to the given ElasticCloudPrimaryShardGracefulTimeoutIntegrationDataflowResponse and assigns it to the ElasticCloudPrimaryShardGracefulTimeout field.
func (o *ElasticCloudIntegrationDataflowsResponse) SetElasticCloudPrimaryShardGracefulTimeout(v ElasticCloudPrimaryShardGracefulTimeoutIntegrationDataflowResponse) {
	o.ElasticCloudPrimaryShardGracefulTimeout = &v
}

// GetElasticCloudPrimaryShardStats returns the ElasticCloudPrimaryShardStats field value if set, zero value otherwise.
func (o *ElasticCloudIntegrationDataflowsResponse) GetElasticCloudPrimaryShardStats() ElasticCloudPrimaryShardStatsIntegrationDataflowResponse {
	if o == nil || o.ElasticCloudPrimaryShardStats == nil {
		var ret ElasticCloudPrimaryShardStatsIntegrationDataflowResponse
		return ret
	}
	return *o.ElasticCloudPrimaryShardStats
}

// GetElasticCloudPrimaryShardStatsOk returns a tuple with the ElasticCloudPrimaryShardStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticCloudIntegrationDataflowsResponse) GetElasticCloudPrimaryShardStatsOk() (*ElasticCloudPrimaryShardStatsIntegrationDataflowResponse, bool) {
	if o == nil || o.ElasticCloudPrimaryShardStats == nil {
		return nil, false
	}
	return o.ElasticCloudPrimaryShardStats, true
}

// HasElasticCloudPrimaryShardStats returns a boolean if a field has been set.
func (o *ElasticCloudIntegrationDataflowsResponse) HasElasticCloudPrimaryShardStats() bool {
	return o != nil && o.ElasticCloudPrimaryShardStats != nil
}

// SetElasticCloudPrimaryShardStats gets a reference to the given ElasticCloudPrimaryShardStatsIntegrationDataflowResponse and assigns it to the ElasticCloudPrimaryShardStats field.
func (o *ElasticCloudIntegrationDataflowsResponse) SetElasticCloudPrimaryShardStats(v ElasticCloudPrimaryShardStatsIntegrationDataflowResponse) {
	o.ElasticCloudPrimaryShardStats = &v
}

// GetElasticCloudShardAllocationStats returns the ElasticCloudShardAllocationStats field value if set, zero value otherwise.
func (o *ElasticCloudIntegrationDataflowsResponse) GetElasticCloudShardAllocationStats() ElasticCloudShardAllocationStatsIntegrationDataflowResponse {
	if o == nil || o.ElasticCloudShardAllocationStats == nil {
		var ret ElasticCloudShardAllocationStatsIntegrationDataflowResponse
		return ret
	}
	return *o.ElasticCloudShardAllocationStats
}

// GetElasticCloudShardAllocationStatsOk returns a tuple with the ElasticCloudShardAllocationStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticCloudIntegrationDataflowsResponse) GetElasticCloudShardAllocationStatsOk() (*ElasticCloudShardAllocationStatsIntegrationDataflowResponse, bool) {
	if o == nil || o.ElasticCloudShardAllocationStats == nil {
		return nil, false
	}
	return o.ElasticCloudShardAllocationStats, true
}

// HasElasticCloudShardAllocationStats returns a boolean if a field has been set.
func (o *ElasticCloudIntegrationDataflowsResponse) HasElasticCloudShardAllocationStats() bool {
	return o != nil && o.ElasticCloudShardAllocationStats != nil
}

// SetElasticCloudShardAllocationStats gets a reference to the given ElasticCloudShardAllocationStatsIntegrationDataflowResponse and assigns it to the ElasticCloudShardAllocationStats field.
func (o *ElasticCloudIntegrationDataflowsResponse) SetElasticCloudShardAllocationStats(v ElasticCloudShardAllocationStatsIntegrationDataflowResponse) {
	o.ElasticCloudShardAllocationStats = &v
}

// GetElasticCloudSlmStats returns the ElasticCloudSlmStats field value if set, zero value otherwise.
func (o *ElasticCloudIntegrationDataflowsResponse) GetElasticCloudSlmStats() ElasticCloudSlmStatsIntegrationDataflowResponse {
	if o == nil || o.ElasticCloudSlmStats == nil {
		var ret ElasticCloudSlmStatsIntegrationDataflowResponse
		return ret
	}
	return *o.ElasticCloudSlmStats
}

// GetElasticCloudSlmStatsOk returns a tuple with the ElasticCloudSlmStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticCloudIntegrationDataflowsResponse) GetElasticCloudSlmStatsOk() (*ElasticCloudSlmStatsIntegrationDataflowResponse, bool) {
	if o == nil || o.ElasticCloudSlmStats == nil {
		return nil, false
	}
	return o.ElasticCloudSlmStats, true
}

// HasElasticCloudSlmStats returns a boolean if a field has been set.
func (o *ElasticCloudIntegrationDataflowsResponse) HasElasticCloudSlmStats() bool {
	return o != nil && o.ElasticCloudSlmStats != nil
}

// SetElasticCloudSlmStats gets a reference to the given ElasticCloudSlmStatsIntegrationDataflowResponse and assigns it to the ElasticCloudSlmStats field.
func (o *ElasticCloudIntegrationDataflowsResponse) SetElasticCloudSlmStats(v ElasticCloudSlmStatsIntegrationDataflowResponse) {
	o.ElasticCloudSlmStats = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ElasticCloudIntegrationDataflowsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ElasticCloudDetailedIndexStats != nil {
		toSerialize["elastic-cloud-detailed-index-stats"] = o.ElasticCloudDetailedIndexStats
	}
	if o.ElasticCloudIndexStats != nil {
		toSerialize["elastic-cloud-index-stats"] = o.ElasticCloudIndexStats
	}
	if o.ElasticCloudMetrics != nil {
		toSerialize["elastic-cloud-metrics"] = o.ElasticCloudMetrics
	}
	if o.ElasticCloudPendingTaskStats != nil {
		toSerialize["elastic-cloud-pending-task-stats"] = o.ElasticCloudPendingTaskStats
	}
	if o.ElasticCloudPrimaryShardGracefulTimeout != nil {
		toSerialize["elastic-cloud-primary-shard-graceful-timeout"] = o.ElasticCloudPrimaryShardGracefulTimeout
	}
	if o.ElasticCloudPrimaryShardStats != nil {
		toSerialize["elastic-cloud-primary-shard-stats"] = o.ElasticCloudPrimaryShardStats
	}
	if o.ElasticCloudShardAllocationStats != nil {
		toSerialize["elastic-cloud-shard-allocation-stats"] = o.ElasticCloudShardAllocationStats
	}
	if o.ElasticCloudSlmStats != nil {
		toSerialize["elastic-cloud-slm-stats"] = o.ElasticCloudSlmStats
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ElasticCloudIntegrationDataflowsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ElasticCloudDetailedIndexStats          *ElasticCloudDetailedIndexStatsIntegrationDataflowResponse          `json:"elastic-cloud-detailed-index-stats,omitempty"`
		ElasticCloudIndexStats                  *ElasticCloudIndexStatsIntegrationDataflowResponse                  `json:"elastic-cloud-index-stats,omitempty"`
		ElasticCloudMetrics                     *ElasticCloudMetricsIntegrationDataflowResponse                     `json:"elastic-cloud-metrics,omitempty"`
		ElasticCloudPendingTaskStats            *ElasticCloudPendingTaskStatsIntegrationDataflowResponse            `json:"elastic-cloud-pending-task-stats,omitempty"`
		ElasticCloudPrimaryShardGracefulTimeout *ElasticCloudPrimaryShardGracefulTimeoutIntegrationDataflowResponse `json:"elastic-cloud-primary-shard-graceful-timeout,omitempty"`
		ElasticCloudPrimaryShardStats           *ElasticCloudPrimaryShardStatsIntegrationDataflowResponse           `json:"elastic-cloud-primary-shard-stats,omitempty"`
		ElasticCloudShardAllocationStats        *ElasticCloudShardAllocationStatsIntegrationDataflowResponse        `json:"elastic-cloud-shard-allocation-stats,omitempty"`
		ElasticCloudSlmStats                    *ElasticCloudSlmStatsIntegrationDataflowResponse                    `json:"elastic-cloud-slm-stats,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	hasInvalidField := false
	if all.ElasticCloudDetailedIndexStats != nil && all.ElasticCloudDetailedIndexStats.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ElasticCloudDetailedIndexStats = all.ElasticCloudDetailedIndexStats
	if all.ElasticCloudIndexStats != nil && all.ElasticCloudIndexStats.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ElasticCloudIndexStats = all.ElasticCloudIndexStats
	if all.ElasticCloudMetrics != nil && all.ElasticCloudMetrics.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ElasticCloudMetrics = all.ElasticCloudMetrics
	if all.ElasticCloudPendingTaskStats != nil && all.ElasticCloudPendingTaskStats.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ElasticCloudPendingTaskStats = all.ElasticCloudPendingTaskStats
	if all.ElasticCloudPrimaryShardGracefulTimeout != nil && all.ElasticCloudPrimaryShardGracefulTimeout.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ElasticCloudPrimaryShardGracefulTimeout = all.ElasticCloudPrimaryShardGracefulTimeout
	if all.ElasticCloudPrimaryShardStats != nil && all.ElasticCloudPrimaryShardStats.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ElasticCloudPrimaryShardStats = all.ElasticCloudPrimaryShardStats
	if all.ElasticCloudShardAllocationStats != nil && all.ElasticCloudShardAllocationStats.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ElasticCloudShardAllocationStats = all.ElasticCloudShardAllocationStats
	if all.ElasticCloudSlmStats != nil && all.ElasticCloudSlmStats.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ElasticCloudSlmStats = all.ElasticCloudSlmStats

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
