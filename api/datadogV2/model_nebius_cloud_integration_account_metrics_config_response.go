// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NebiusCloudIntegrationAccountMetricsConfigResponse Metric collection settings configured on the Nebius Cloud integration account. Absent or empty when no filters are configured.
type NebiusCloudIntegrationAccountMetricsConfigResponse struct {
	// Tag-filter rules limiting metric collection to matching Kubernetes clusters.
	ClusterFilters []string `json:"cluster_filters,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNebiusCloudIntegrationAccountMetricsConfigResponse instantiates a new NebiusCloudIntegrationAccountMetricsConfigResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNebiusCloudIntegrationAccountMetricsConfigResponse() *NebiusCloudIntegrationAccountMetricsConfigResponse {
	this := NebiusCloudIntegrationAccountMetricsConfigResponse{}
	return &this
}

// NewNebiusCloudIntegrationAccountMetricsConfigResponseWithDefaults instantiates a new NebiusCloudIntegrationAccountMetricsConfigResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNebiusCloudIntegrationAccountMetricsConfigResponseWithDefaults() *NebiusCloudIntegrationAccountMetricsConfigResponse {
	this := NebiusCloudIntegrationAccountMetricsConfigResponse{}
	return &this
}

// GetClusterFilters returns the ClusterFilters field value if set, zero value otherwise.
func (o *NebiusCloudIntegrationAccountMetricsConfigResponse) GetClusterFilters() []string {
	if o == nil || o.ClusterFilters == nil {
		var ret []string
		return ret
	}
	return o.ClusterFilters
}

// GetClusterFiltersOk returns a tuple with the ClusterFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NebiusCloudIntegrationAccountMetricsConfigResponse) GetClusterFiltersOk() (*[]string, bool) {
	if o == nil || o.ClusterFilters == nil {
		return nil, false
	}
	return &o.ClusterFilters, true
}

// HasClusterFilters returns a boolean if a field has been set.
func (o *NebiusCloudIntegrationAccountMetricsConfigResponse) HasClusterFilters() bool {
	return o != nil && o.ClusterFilters != nil
}

// SetClusterFilters gets a reference to the given []string and assigns it to the ClusterFilters field.
func (o *NebiusCloudIntegrationAccountMetricsConfigResponse) SetClusterFilters(v []string) {
	o.ClusterFilters = v
}

// MarshalJSON serializes the struct using spec logic.
func (o NebiusCloudIntegrationAccountMetricsConfigResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ClusterFilters != nil {
		toSerialize["cluster_filters"] = o.ClusterFilters
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NebiusCloudIntegrationAccountMetricsConfigResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ClusterFilters []string `json:"cluster_filters,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cluster_filters"})
	} else {
		return err
	}
	o.ClusterFilters = all.ClusterFilters

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
