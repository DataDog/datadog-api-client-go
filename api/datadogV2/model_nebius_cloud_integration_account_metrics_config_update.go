// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NebiusCloudIntegrationAccountMetricsConfigUpdate Configuration for metric collection. Only the fields provided are changed.
type NebiusCloudIntegrationAccountMetricsConfigUpdate struct {
	// Tag-filter rules limiting metric collection to matching Kubernetes clusters. Sending metrics_config replaces the stored object's cluster_filters list; omit it to keep the stored value. Include rules (`key:value`) are combined with OR, exclusion rules (`!key:value`) are then applied with AND, and `?` and `*` are wildcards.
	ClusterFilters []string `json:"cluster_filters,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewNebiusCloudIntegrationAccountMetricsConfigUpdate instantiates a new NebiusCloudIntegrationAccountMetricsConfigUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNebiusCloudIntegrationAccountMetricsConfigUpdate() *NebiusCloudIntegrationAccountMetricsConfigUpdate {
	this := NebiusCloudIntegrationAccountMetricsConfigUpdate{}
	return &this
}

// NewNebiusCloudIntegrationAccountMetricsConfigUpdateWithDefaults instantiates a new NebiusCloudIntegrationAccountMetricsConfigUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNebiusCloudIntegrationAccountMetricsConfigUpdateWithDefaults() *NebiusCloudIntegrationAccountMetricsConfigUpdate {
	this := NebiusCloudIntegrationAccountMetricsConfigUpdate{}
	return &this
}

// GetClusterFilters returns the ClusterFilters field value if set, zero value otherwise.
func (o *NebiusCloudIntegrationAccountMetricsConfigUpdate) GetClusterFilters() []string {
	if o == nil || o.ClusterFilters == nil {
		var ret []string
		return ret
	}
	return o.ClusterFilters
}

// GetClusterFiltersOk returns a tuple with the ClusterFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NebiusCloudIntegrationAccountMetricsConfigUpdate) GetClusterFiltersOk() (*[]string, bool) {
	if o == nil || o.ClusterFilters == nil {
		return nil, false
	}
	return &o.ClusterFilters, true
}

// HasClusterFilters returns a boolean if a field has been set.
func (o *NebiusCloudIntegrationAccountMetricsConfigUpdate) HasClusterFilters() bool {
	return o != nil && o.ClusterFilters != nil
}

// SetClusterFilters gets a reference to the given []string and assigns it to the ClusterFilters field.
func (o *NebiusCloudIntegrationAccountMetricsConfigUpdate) SetClusterFilters(v []string) {
	o.ClusterFilters = v
}

// MarshalJSON serializes the struct using spec logic.
func (o NebiusCloudIntegrationAccountMetricsConfigUpdate) MarshalJSON() ([]byte, error) {
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
func (o *NebiusCloudIntegrationAccountMetricsConfigUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ClusterFilters []string `json:"cluster_filters,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.ClusterFilters = all.ClusterFilters

	return nil
}
