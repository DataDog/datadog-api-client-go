// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SnowflakeAccountUsageMetricsIntegrationDataflowSettingsResponse Settings of the account usage metrics dataflow.
type SnowflakeAccountUsageMetricsIntegrationDataflowSettingsResponse struct {
	// Period each metric aggregates over. When `true`, metrics aggregate the past 24 hours on a rolling basis; when `false`, they aggregate the current day so far.
	AccountUsageMetricsAggregateLast24h *bool `json:"account_usage_metrics_aggregate_last_24h,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSnowflakeAccountUsageMetricsIntegrationDataflowSettingsResponse instantiates a new SnowflakeAccountUsageMetricsIntegrationDataflowSettingsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSnowflakeAccountUsageMetricsIntegrationDataflowSettingsResponse() *SnowflakeAccountUsageMetricsIntegrationDataflowSettingsResponse {
	this := SnowflakeAccountUsageMetricsIntegrationDataflowSettingsResponse{}
	return &this
}

// NewSnowflakeAccountUsageMetricsIntegrationDataflowSettingsResponseWithDefaults instantiates a new SnowflakeAccountUsageMetricsIntegrationDataflowSettingsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSnowflakeAccountUsageMetricsIntegrationDataflowSettingsResponseWithDefaults() *SnowflakeAccountUsageMetricsIntegrationDataflowSettingsResponse {
	this := SnowflakeAccountUsageMetricsIntegrationDataflowSettingsResponse{}
	return &this
}

// GetAccountUsageMetricsAggregateLast24h returns the AccountUsageMetricsAggregateLast24h field value if set, zero value otherwise.
func (o *SnowflakeAccountUsageMetricsIntegrationDataflowSettingsResponse) GetAccountUsageMetricsAggregateLast24h() bool {
	if o == nil || o.AccountUsageMetricsAggregateLast24h == nil {
		var ret bool
		return ret
	}
	return *o.AccountUsageMetricsAggregateLast24h
}

// GetAccountUsageMetricsAggregateLast24hOk returns a tuple with the AccountUsageMetricsAggregateLast24h field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeAccountUsageMetricsIntegrationDataflowSettingsResponse) GetAccountUsageMetricsAggregateLast24hOk() (*bool, bool) {
	if o == nil || o.AccountUsageMetricsAggregateLast24h == nil {
		return nil, false
	}
	return o.AccountUsageMetricsAggregateLast24h, true
}

// HasAccountUsageMetricsAggregateLast24h returns a boolean if a field has been set.
func (o *SnowflakeAccountUsageMetricsIntegrationDataflowSettingsResponse) HasAccountUsageMetricsAggregateLast24h() bool {
	return o != nil && o.AccountUsageMetricsAggregateLast24h != nil
}

// SetAccountUsageMetricsAggregateLast24h gets a reference to the given bool and assigns it to the AccountUsageMetricsAggregateLast24h field.
func (o *SnowflakeAccountUsageMetricsIntegrationDataflowSettingsResponse) SetAccountUsageMetricsAggregateLast24h(v bool) {
	o.AccountUsageMetricsAggregateLast24h = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SnowflakeAccountUsageMetricsIntegrationDataflowSettingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AccountUsageMetricsAggregateLast24h != nil {
		toSerialize["account_usage_metrics_aggregate_last_24h"] = o.AccountUsageMetricsAggregateLast24h
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SnowflakeAccountUsageMetricsIntegrationDataflowSettingsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountUsageMetricsAggregateLast24h *bool `json:"account_usage_metrics_aggregate_last_24h,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_usage_metrics_aggregate_last_24h"})
	} else {
		return err
	}
	o.AccountUsageMetricsAggregateLast24h = all.AccountUsageMetricsAggregateLast24h

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
