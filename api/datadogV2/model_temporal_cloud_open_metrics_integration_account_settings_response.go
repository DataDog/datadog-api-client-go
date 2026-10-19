// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TemporalCloudOpenMetricsIntegrationAccountSettingsResponse Settings configured on the Temporal Cloud OpenMetrics integration account.
type TemporalCloudOpenMetricsIntegrationAccountSettingsResponse struct {
	// When enabled, Datadog collects Temporal Cloud's [opt-in labels](https://docs.temporal.io/cloud/metrics/openmetrics/metrics-reference#opt-in-labels) (activity type, worker deployment name, and worker build ID) for more granularity on your existing metrics. Enabling them may increase metric collection latency.
	IncludeOptInLabels *bool `json:"include_opt_in_labels,omitempty"`
	// Comma-separated list of Temporal Cloud namespaces to collect metrics for. An entry can use `*` as a wildcard; for example, `prod-*` matches every namespace starting with `prod-`. Empty or absent when metrics are collected for all namespaces.
	NamespacesToInclude *string `json:"namespaces_to_include,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTemporalCloudOpenMetricsIntegrationAccountSettingsResponse instantiates a new TemporalCloudOpenMetricsIntegrationAccountSettingsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTemporalCloudOpenMetricsIntegrationAccountSettingsResponse() *TemporalCloudOpenMetricsIntegrationAccountSettingsResponse {
	this := TemporalCloudOpenMetricsIntegrationAccountSettingsResponse{}
	return &this
}

// NewTemporalCloudOpenMetricsIntegrationAccountSettingsResponseWithDefaults instantiates a new TemporalCloudOpenMetricsIntegrationAccountSettingsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTemporalCloudOpenMetricsIntegrationAccountSettingsResponseWithDefaults() *TemporalCloudOpenMetricsIntegrationAccountSettingsResponse {
	this := TemporalCloudOpenMetricsIntegrationAccountSettingsResponse{}
	return &this
}

// GetIncludeOptInLabels returns the IncludeOptInLabels field value if set, zero value otherwise.
func (o *TemporalCloudOpenMetricsIntegrationAccountSettingsResponse) GetIncludeOptInLabels() bool {
	if o == nil || o.IncludeOptInLabels == nil {
		var ret bool
		return ret
	}
	return *o.IncludeOptInLabels
}

// GetIncludeOptInLabelsOk returns a tuple with the IncludeOptInLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemporalCloudOpenMetricsIntegrationAccountSettingsResponse) GetIncludeOptInLabelsOk() (*bool, bool) {
	if o == nil || o.IncludeOptInLabels == nil {
		return nil, false
	}
	return o.IncludeOptInLabels, true
}

// HasIncludeOptInLabels returns a boolean if a field has been set.
func (o *TemporalCloudOpenMetricsIntegrationAccountSettingsResponse) HasIncludeOptInLabels() bool {
	return o != nil && o.IncludeOptInLabels != nil
}

// SetIncludeOptInLabels gets a reference to the given bool and assigns it to the IncludeOptInLabels field.
func (o *TemporalCloudOpenMetricsIntegrationAccountSettingsResponse) SetIncludeOptInLabels(v bool) {
	o.IncludeOptInLabels = &v
}

// GetNamespacesToInclude returns the NamespacesToInclude field value if set, zero value otherwise.
func (o *TemporalCloudOpenMetricsIntegrationAccountSettingsResponse) GetNamespacesToInclude() string {
	if o == nil || o.NamespacesToInclude == nil {
		var ret string
		return ret
	}
	return *o.NamespacesToInclude
}

// GetNamespacesToIncludeOk returns a tuple with the NamespacesToInclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemporalCloudOpenMetricsIntegrationAccountSettingsResponse) GetNamespacesToIncludeOk() (*string, bool) {
	if o == nil || o.NamespacesToInclude == nil {
		return nil, false
	}
	return o.NamespacesToInclude, true
}

// HasNamespacesToInclude returns a boolean if a field has been set.
func (o *TemporalCloudOpenMetricsIntegrationAccountSettingsResponse) HasNamespacesToInclude() bool {
	return o != nil && o.NamespacesToInclude != nil
}

// SetNamespacesToInclude gets a reference to the given string and assigns it to the NamespacesToInclude field.
func (o *TemporalCloudOpenMetricsIntegrationAccountSettingsResponse) SetNamespacesToInclude(v string) {
	o.NamespacesToInclude = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TemporalCloudOpenMetricsIntegrationAccountSettingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.IncludeOptInLabels != nil {
		toSerialize["include_opt_in_labels"] = o.IncludeOptInLabels
	}
	if o.NamespacesToInclude != nil {
		toSerialize["namespaces_to_include"] = o.NamespacesToInclude
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TemporalCloudOpenMetricsIntegrationAccountSettingsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IncludeOptInLabels  *bool   `json:"include_opt_in_labels,omitempty"`
		NamespacesToInclude *string `json:"namespaces_to_include,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"include_opt_in_labels", "namespaces_to_include"})
	} else {
		return err
	}
	o.IncludeOptInLabels = all.IncludeOptInLabels
	o.NamespacesToInclude = all.NamespacesToInclude

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
