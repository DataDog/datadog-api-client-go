// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ElasticCloudIntegrationDataflowsRequest Dataflows to configure on the Elastic Cloud integration account, keyed by dataflow id.
type ElasticCloudIntegrationDataflowsRequest struct {
	// The Elastic Cloud detailed index stats dataflow.
	ElasticCloudDetailedIndexStats *ElasticCloudDetailedIndexStatsIntegrationDataflowRequest `json:"elastic-cloud-detailed-index-stats,omitempty"`
	// The Elastic Cloud index stats dataflow.
	ElasticCloudIndexStats *ElasticCloudIndexStatsIntegrationDataflowRequest `json:"elastic-cloud-index-stats,omitempty"`
	// The Elastic Cloud pending task stats dataflow.
	ElasticCloudPendingTaskStats *ElasticCloudPendingTaskStatsIntegrationDataflowRequest `json:"elastic-cloud-pending-task-stats,omitempty"`
	// The Elastic Cloud primary shard graceful timeout dataflow.
	ElasticCloudPrimaryShardGracefulTimeout *ElasticCloudPrimaryShardGracefulTimeoutIntegrationDataflowRequest `json:"elastic-cloud-primary-shard-graceful-timeout,omitempty"`
	// The Elastic Cloud primary shard stats dataflow.
	ElasticCloudPrimaryShardStats *ElasticCloudPrimaryShardStatsIntegrationDataflowRequest `json:"elastic-cloud-primary-shard-stats,omitempty"`
	// The Elastic Cloud shard allocation stats dataflow.
	ElasticCloudShardAllocationStats *ElasticCloudShardAllocationStatsIntegrationDataflowRequest `json:"elastic-cloud-shard-allocation-stats,omitempty"`
	// The Elastic Cloud snapshot lifecycle management stats dataflow.
	ElasticCloudSlmStats *ElasticCloudSlmStatsIntegrationDataflowRequest `json:"elastic-cloud-slm-stats,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewElasticCloudIntegrationDataflowsRequest instantiates a new ElasticCloudIntegrationDataflowsRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewElasticCloudIntegrationDataflowsRequest() *ElasticCloudIntegrationDataflowsRequest {
	this := ElasticCloudIntegrationDataflowsRequest{}
	return &this
}

// NewElasticCloudIntegrationDataflowsRequestWithDefaults instantiates a new ElasticCloudIntegrationDataflowsRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewElasticCloudIntegrationDataflowsRequestWithDefaults() *ElasticCloudIntegrationDataflowsRequest {
	this := ElasticCloudIntegrationDataflowsRequest{}
	return &this
}

// GetElasticCloudDetailedIndexStats returns the ElasticCloudDetailedIndexStats field value if set, zero value otherwise.
func (o *ElasticCloudIntegrationDataflowsRequest) GetElasticCloudDetailedIndexStats() ElasticCloudDetailedIndexStatsIntegrationDataflowRequest {
	if o == nil || o.ElasticCloudDetailedIndexStats == nil {
		var ret ElasticCloudDetailedIndexStatsIntegrationDataflowRequest
		return ret
	}
	return *o.ElasticCloudDetailedIndexStats
}

// GetElasticCloudDetailedIndexStatsOk returns a tuple with the ElasticCloudDetailedIndexStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticCloudIntegrationDataflowsRequest) GetElasticCloudDetailedIndexStatsOk() (*ElasticCloudDetailedIndexStatsIntegrationDataflowRequest, bool) {
	if o == nil || o.ElasticCloudDetailedIndexStats == nil {
		return nil, false
	}
	return o.ElasticCloudDetailedIndexStats, true
}

// HasElasticCloudDetailedIndexStats returns a boolean if a field has been set.
func (o *ElasticCloudIntegrationDataflowsRequest) HasElasticCloudDetailedIndexStats() bool {
	return o != nil && o.ElasticCloudDetailedIndexStats != nil
}

// SetElasticCloudDetailedIndexStats gets a reference to the given ElasticCloudDetailedIndexStatsIntegrationDataflowRequest and assigns it to the ElasticCloudDetailedIndexStats field.
func (o *ElasticCloudIntegrationDataflowsRequest) SetElasticCloudDetailedIndexStats(v ElasticCloudDetailedIndexStatsIntegrationDataflowRequest) {
	o.ElasticCloudDetailedIndexStats = &v
}

// GetElasticCloudIndexStats returns the ElasticCloudIndexStats field value if set, zero value otherwise.
func (o *ElasticCloudIntegrationDataflowsRequest) GetElasticCloudIndexStats() ElasticCloudIndexStatsIntegrationDataflowRequest {
	if o == nil || o.ElasticCloudIndexStats == nil {
		var ret ElasticCloudIndexStatsIntegrationDataflowRequest
		return ret
	}
	return *o.ElasticCloudIndexStats
}

// GetElasticCloudIndexStatsOk returns a tuple with the ElasticCloudIndexStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticCloudIntegrationDataflowsRequest) GetElasticCloudIndexStatsOk() (*ElasticCloudIndexStatsIntegrationDataflowRequest, bool) {
	if o == nil || o.ElasticCloudIndexStats == nil {
		return nil, false
	}
	return o.ElasticCloudIndexStats, true
}

// HasElasticCloudIndexStats returns a boolean if a field has been set.
func (o *ElasticCloudIntegrationDataflowsRequest) HasElasticCloudIndexStats() bool {
	return o != nil && o.ElasticCloudIndexStats != nil
}

// SetElasticCloudIndexStats gets a reference to the given ElasticCloudIndexStatsIntegrationDataflowRequest and assigns it to the ElasticCloudIndexStats field.
func (o *ElasticCloudIntegrationDataflowsRequest) SetElasticCloudIndexStats(v ElasticCloudIndexStatsIntegrationDataflowRequest) {
	o.ElasticCloudIndexStats = &v
}

// GetElasticCloudPendingTaskStats returns the ElasticCloudPendingTaskStats field value if set, zero value otherwise.
func (o *ElasticCloudIntegrationDataflowsRequest) GetElasticCloudPendingTaskStats() ElasticCloudPendingTaskStatsIntegrationDataflowRequest {
	if o == nil || o.ElasticCloudPendingTaskStats == nil {
		var ret ElasticCloudPendingTaskStatsIntegrationDataflowRequest
		return ret
	}
	return *o.ElasticCloudPendingTaskStats
}

// GetElasticCloudPendingTaskStatsOk returns a tuple with the ElasticCloudPendingTaskStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticCloudIntegrationDataflowsRequest) GetElasticCloudPendingTaskStatsOk() (*ElasticCloudPendingTaskStatsIntegrationDataflowRequest, bool) {
	if o == nil || o.ElasticCloudPendingTaskStats == nil {
		return nil, false
	}
	return o.ElasticCloudPendingTaskStats, true
}

// HasElasticCloudPendingTaskStats returns a boolean if a field has been set.
func (o *ElasticCloudIntegrationDataflowsRequest) HasElasticCloudPendingTaskStats() bool {
	return o != nil && o.ElasticCloudPendingTaskStats != nil
}

// SetElasticCloudPendingTaskStats gets a reference to the given ElasticCloudPendingTaskStatsIntegrationDataflowRequest and assigns it to the ElasticCloudPendingTaskStats field.
func (o *ElasticCloudIntegrationDataflowsRequest) SetElasticCloudPendingTaskStats(v ElasticCloudPendingTaskStatsIntegrationDataflowRequest) {
	o.ElasticCloudPendingTaskStats = &v
}

// GetElasticCloudPrimaryShardGracefulTimeout returns the ElasticCloudPrimaryShardGracefulTimeout field value if set, zero value otherwise.
func (o *ElasticCloudIntegrationDataflowsRequest) GetElasticCloudPrimaryShardGracefulTimeout() ElasticCloudPrimaryShardGracefulTimeoutIntegrationDataflowRequest {
	if o == nil || o.ElasticCloudPrimaryShardGracefulTimeout == nil {
		var ret ElasticCloudPrimaryShardGracefulTimeoutIntegrationDataflowRequest
		return ret
	}
	return *o.ElasticCloudPrimaryShardGracefulTimeout
}

// GetElasticCloudPrimaryShardGracefulTimeoutOk returns a tuple with the ElasticCloudPrimaryShardGracefulTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticCloudIntegrationDataflowsRequest) GetElasticCloudPrimaryShardGracefulTimeoutOk() (*ElasticCloudPrimaryShardGracefulTimeoutIntegrationDataflowRequest, bool) {
	if o == nil || o.ElasticCloudPrimaryShardGracefulTimeout == nil {
		return nil, false
	}
	return o.ElasticCloudPrimaryShardGracefulTimeout, true
}

// HasElasticCloudPrimaryShardGracefulTimeout returns a boolean if a field has been set.
func (o *ElasticCloudIntegrationDataflowsRequest) HasElasticCloudPrimaryShardGracefulTimeout() bool {
	return o != nil && o.ElasticCloudPrimaryShardGracefulTimeout != nil
}

// SetElasticCloudPrimaryShardGracefulTimeout gets a reference to the given ElasticCloudPrimaryShardGracefulTimeoutIntegrationDataflowRequest and assigns it to the ElasticCloudPrimaryShardGracefulTimeout field.
func (o *ElasticCloudIntegrationDataflowsRequest) SetElasticCloudPrimaryShardGracefulTimeout(v ElasticCloudPrimaryShardGracefulTimeoutIntegrationDataflowRequest) {
	o.ElasticCloudPrimaryShardGracefulTimeout = &v
}

// GetElasticCloudPrimaryShardStats returns the ElasticCloudPrimaryShardStats field value if set, zero value otherwise.
func (o *ElasticCloudIntegrationDataflowsRequest) GetElasticCloudPrimaryShardStats() ElasticCloudPrimaryShardStatsIntegrationDataflowRequest {
	if o == nil || o.ElasticCloudPrimaryShardStats == nil {
		var ret ElasticCloudPrimaryShardStatsIntegrationDataflowRequest
		return ret
	}
	return *o.ElasticCloudPrimaryShardStats
}

// GetElasticCloudPrimaryShardStatsOk returns a tuple with the ElasticCloudPrimaryShardStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticCloudIntegrationDataflowsRequest) GetElasticCloudPrimaryShardStatsOk() (*ElasticCloudPrimaryShardStatsIntegrationDataflowRequest, bool) {
	if o == nil || o.ElasticCloudPrimaryShardStats == nil {
		return nil, false
	}
	return o.ElasticCloudPrimaryShardStats, true
}

// HasElasticCloudPrimaryShardStats returns a boolean if a field has been set.
func (o *ElasticCloudIntegrationDataflowsRequest) HasElasticCloudPrimaryShardStats() bool {
	return o != nil && o.ElasticCloudPrimaryShardStats != nil
}

// SetElasticCloudPrimaryShardStats gets a reference to the given ElasticCloudPrimaryShardStatsIntegrationDataflowRequest and assigns it to the ElasticCloudPrimaryShardStats field.
func (o *ElasticCloudIntegrationDataflowsRequest) SetElasticCloudPrimaryShardStats(v ElasticCloudPrimaryShardStatsIntegrationDataflowRequest) {
	o.ElasticCloudPrimaryShardStats = &v
}

// GetElasticCloudShardAllocationStats returns the ElasticCloudShardAllocationStats field value if set, zero value otherwise.
func (o *ElasticCloudIntegrationDataflowsRequest) GetElasticCloudShardAllocationStats() ElasticCloudShardAllocationStatsIntegrationDataflowRequest {
	if o == nil || o.ElasticCloudShardAllocationStats == nil {
		var ret ElasticCloudShardAllocationStatsIntegrationDataflowRequest
		return ret
	}
	return *o.ElasticCloudShardAllocationStats
}

// GetElasticCloudShardAllocationStatsOk returns a tuple with the ElasticCloudShardAllocationStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticCloudIntegrationDataflowsRequest) GetElasticCloudShardAllocationStatsOk() (*ElasticCloudShardAllocationStatsIntegrationDataflowRequest, bool) {
	if o == nil || o.ElasticCloudShardAllocationStats == nil {
		return nil, false
	}
	return o.ElasticCloudShardAllocationStats, true
}

// HasElasticCloudShardAllocationStats returns a boolean if a field has been set.
func (o *ElasticCloudIntegrationDataflowsRequest) HasElasticCloudShardAllocationStats() bool {
	return o != nil && o.ElasticCloudShardAllocationStats != nil
}

// SetElasticCloudShardAllocationStats gets a reference to the given ElasticCloudShardAllocationStatsIntegrationDataflowRequest and assigns it to the ElasticCloudShardAllocationStats field.
func (o *ElasticCloudIntegrationDataflowsRequest) SetElasticCloudShardAllocationStats(v ElasticCloudShardAllocationStatsIntegrationDataflowRequest) {
	o.ElasticCloudShardAllocationStats = &v
}

// GetElasticCloudSlmStats returns the ElasticCloudSlmStats field value if set, zero value otherwise.
func (o *ElasticCloudIntegrationDataflowsRequest) GetElasticCloudSlmStats() ElasticCloudSlmStatsIntegrationDataflowRequest {
	if o == nil || o.ElasticCloudSlmStats == nil {
		var ret ElasticCloudSlmStatsIntegrationDataflowRequest
		return ret
	}
	return *o.ElasticCloudSlmStats
}

// GetElasticCloudSlmStatsOk returns a tuple with the ElasticCloudSlmStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticCloudIntegrationDataflowsRequest) GetElasticCloudSlmStatsOk() (*ElasticCloudSlmStatsIntegrationDataflowRequest, bool) {
	if o == nil || o.ElasticCloudSlmStats == nil {
		return nil, false
	}
	return o.ElasticCloudSlmStats, true
}

// HasElasticCloudSlmStats returns a boolean if a field has been set.
func (o *ElasticCloudIntegrationDataflowsRequest) HasElasticCloudSlmStats() bool {
	return o != nil && o.ElasticCloudSlmStats != nil
}

// SetElasticCloudSlmStats gets a reference to the given ElasticCloudSlmStatsIntegrationDataflowRequest and assigns it to the ElasticCloudSlmStats field.
func (o *ElasticCloudIntegrationDataflowsRequest) SetElasticCloudSlmStats(v ElasticCloudSlmStatsIntegrationDataflowRequest) {
	o.ElasticCloudSlmStats = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ElasticCloudIntegrationDataflowsRequest) MarshalJSON() ([]byte, error) {
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
func (o *ElasticCloudIntegrationDataflowsRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ElasticCloudDetailedIndexStats          *ElasticCloudDetailedIndexStatsIntegrationDataflowRequest          `json:"elastic-cloud-detailed-index-stats,omitempty"`
		ElasticCloudIndexStats                  *ElasticCloudIndexStatsIntegrationDataflowRequest                  `json:"elastic-cloud-index-stats,omitempty"`
		ElasticCloudPendingTaskStats            *ElasticCloudPendingTaskStatsIntegrationDataflowRequest            `json:"elastic-cloud-pending-task-stats,omitempty"`
		ElasticCloudPrimaryShardGracefulTimeout *ElasticCloudPrimaryShardGracefulTimeoutIntegrationDataflowRequest `json:"elastic-cloud-primary-shard-graceful-timeout,omitempty"`
		ElasticCloudPrimaryShardStats           *ElasticCloudPrimaryShardStatsIntegrationDataflowRequest           `json:"elastic-cloud-primary-shard-stats,omitempty"`
		ElasticCloudShardAllocationStats        *ElasticCloudShardAllocationStatsIntegrationDataflowRequest        `json:"elastic-cloud-shard-allocation-stats,omitempty"`
		ElasticCloudSlmStats                    *ElasticCloudSlmStatsIntegrationDataflowRequest                    `json:"elastic-cloud-slm-stats,omitempty"`
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
