// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SnowflakeOrganizationUsageMetricsIntegrationDataflowSettingsRequest Settings of the organization usage metrics dataflow. Only the fields provided are changed.
type SnowflakeOrganizationUsageMetricsIntegrationDataflowSettingsRequest struct {
	// Period each metric aggregates over. Set to `true` to aggregate the past 24 hours on a rolling basis, or `false` to aggregate the current day so far. Defaults to `false`.
	OrganizationUsageMetricsAggregateLast24h *bool `json:"organization_usage_metrics_aggregate_last_24h,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewSnowflakeOrganizationUsageMetricsIntegrationDataflowSettingsRequest instantiates a new SnowflakeOrganizationUsageMetricsIntegrationDataflowSettingsRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSnowflakeOrganizationUsageMetricsIntegrationDataflowSettingsRequest() *SnowflakeOrganizationUsageMetricsIntegrationDataflowSettingsRequest {
	this := SnowflakeOrganizationUsageMetricsIntegrationDataflowSettingsRequest{}
	return &this
}

// NewSnowflakeOrganizationUsageMetricsIntegrationDataflowSettingsRequestWithDefaults instantiates a new SnowflakeOrganizationUsageMetricsIntegrationDataflowSettingsRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSnowflakeOrganizationUsageMetricsIntegrationDataflowSettingsRequestWithDefaults() *SnowflakeOrganizationUsageMetricsIntegrationDataflowSettingsRequest {
	this := SnowflakeOrganizationUsageMetricsIntegrationDataflowSettingsRequest{}
	return &this
}

// GetOrganizationUsageMetricsAggregateLast24h returns the OrganizationUsageMetricsAggregateLast24h field value if set, zero value otherwise.
func (o *SnowflakeOrganizationUsageMetricsIntegrationDataflowSettingsRequest) GetOrganizationUsageMetricsAggregateLast24h() bool {
	if o == nil || o.OrganizationUsageMetricsAggregateLast24h == nil {
		var ret bool
		return ret
	}
	return *o.OrganizationUsageMetricsAggregateLast24h
}

// GetOrganizationUsageMetricsAggregateLast24hOk returns a tuple with the OrganizationUsageMetricsAggregateLast24h field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeOrganizationUsageMetricsIntegrationDataflowSettingsRequest) GetOrganizationUsageMetricsAggregateLast24hOk() (*bool, bool) {
	if o == nil || o.OrganizationUsageMetricsAggregateLast24h == nil {
		return nil, false
	}
	return o.OrganizationUsageMetricsAggregateLast24h, true
}

// HasOrganizationUsageMetricsAggregateLast24h returns a boolean if a field has been set.
func (o *SnowflakeOrganizationUsageMetricsIntegrationDataflowSettingsRequest) HasOrganizationUsageMetricsAggregateLast24h() bool {
	return o != nil && o.OrganizationUsageMetricsAggregateLast24h != nil
}

// SetOrganizationUsageMetricsAggregateLast24h gets a reference to the given bool and assigns it to the OrganizationUsageMetricsAggregateLast24h field.
func (o *SnowflakeOrganizationUsageMetricsIntegrationDataflowSettingsRequest) SetOrganizationUsageMetricsAggregateLast24h(v bool) {
	o.OrganizationUsageMetricsAggregateLast24h = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SnowflakeOrganizationUsageMetricsIntegrationDataflowSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.OrganizationUsageMetricsAggregateLast24h != nil {
		toSerialize["organization_usage_metrics_aggregate_last_24h"] = o.OrganizationUsageMetricsAggregateLast24h
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SnowflakeOrganizationUsageMetricsIntegrationDataflowSettingsRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OrganizationUsageMetricsAggregateLast24h *bool `json:"organization_usage_metrics_aggregate_last_24h,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.OrganizationUsageMetricsAggregateLast24h = all.OrganizationUsageMetricsAggregateLast24h

	return nil
}
