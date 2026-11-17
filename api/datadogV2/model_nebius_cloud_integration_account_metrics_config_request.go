// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NebiusCloudIntegrationAccountMetricsConfigRequest Configuration for metric collection from this Nebius Cloud account.
type NebiusCloudIntegrationAccountMetricsConfigRequest struct {
	// Tag-filter rules limiting metric collection to matching Kubernetes clusters, for cost savings. Include rules (`key:value`) are combined with OR, exclusion rules (`!key:value`) are then applied with AND, and `?` and `*` are wildcards; for example `cluster:prod-*` collects only matching clusters and `!cluster:dev` excludes a cluster. Defaults to no filters: every cluster is collected. Only cluster-scoped metrics are filtered; billing and object storage metrics are always collected.
	ClusterFilters []string `json:"cluster_filters,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewNebiusCloudIntegrationAccountMetricsConfigRequest instantiates a new NebiusCloudIntegrationAccountMetricsConfigRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNebiusCloudIntegrationAccountMetricsConfigRequest() *NebiusCloudIntegrationAccountMetricsConfigRequest {
	this := NebiusCloudIntegrationAccountMetricsConfigRequest{}
	return &this
}

// NewNebiusCloudIntegrationAccountMetricsConfigRequestWithDefaults instantiates a new NebiusCloudIntegrationAccountMetricsConfigRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNebiusCloudIntegrationAccountMetricsConfigRequestWithDefaults() *NebiusCloudIntegrationAccountMetricsConfigRequest {
	this := NebiusCloudIntegrationAccountMetricsConfigRequest{}
	return &this
}

// GetClusterFilters returns the ClusterFilters field value if set, zero value otherwise.
func (o *NebiusCloudIntegrationAccountMetricsConfigRequest) GetClusterFilters() []string {
	if o == nil || o.ClusterFilters == nil {
		var ret []string
		return ret
	}
	return o.ClusterFilters
}

// GetClusterFiltersOk returns a tuple with the ClusterFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NebiusCloudIntegrationAccountMetricsConfigRequest) GetClusterFiltersOk() (*[]string, bool) {
	if o == nil || o.ClusterFilters == nil {
		return nil, false
	}
	return &o.ClusterFilters, true
}

// HasClusterFilters returns a boolean if a field has been set.
func (o *NebiusCloudIntegrationAccountMetricsConfigRequest) HasClusterFilters() bool {
	return o != nil && o.ClusterFilters != nil
}

// SetClusterFilters gets a reference to the given []string and assigns it to the ClusterFilters field.
func (o *NebiusCloudIntegrationAccountMetricsConfigRequest) SetClusterFilters(v []string) {
	o.ClusterFilters = v
}

// MarshalJSON serializes the struct using spec logic.
func (o NebiusCloudIntegrationAccountMetricsConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ClusterFilters != nil {
		toSerialize["cluster_filters"] = o.ClusterFilters
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NebiusCloudIntegrationAccountMetricsConfigRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ClusterFilters []string `json:"cluster_filters,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.ClusterFilters = all.ClusterFilters

	return nil
}
