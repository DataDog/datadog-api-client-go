// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSMetrics AWS Metrics config
type AWSMetrics struct {
	// Enable EC2 automute for AWS metrics
	AutomuteEnabled *bool `json:"automute_enabled,omitempty"`
	// Enable CloudWatch alarms collection
	CollectCloudwatchAlarms *bool `json:"collect_cloudwatch_alarms,omitempty"`
	// Enable custom metrics collection
	CollectCustomMetrics *bool `json:"collect_custom_metrics,omitempty"`
	// Enable AWS metrics collection
	Enabled *bool `json:"enabled,omitempty"`
	// AWS Metrics namespace filters
	NamespaceFilters *AWSNamespacesList `json:"namespace_filters,omitempty"`
	// AWS Metrics tag filters list
	TagFilters []AWSNamespaceTagFilter `json:"tag_filters,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewAWSMetrics instantiates a new AWSMetrics object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSMetrics() *AWSMetrics {
	this := AWSMetrics{}
	return &this
}

// NewAWSMetricsWithDefaults instantiates a new AWSMetrics object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSMetricsWithDefaults() *AWSMetrics {
	this := AWSMetrics{}
	return &this
}

// GetAutomuteEnabled returns the AutomuteEnabled field value if set, zero value otherwise.
func (o *AWSMetrics) GetAutomuteEnabled() bool {
	if o == nil || o.AutomuteEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AutomuteEnabled
}

// GetAutomuteEnabledOk returns a tuple with the AutomuteEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSMetrics) GetAutomuteEnabledOk() (*bool, bool) {
	if o == nil || o.AutomuteEnabled == nil {
		return nil, false
	}
	return o.AutomuteEnabled, true
}

// HasAutomuteEnabled returns a boolean if a field has been set.
func (o *AWSMetrics) HasAutomuteEnabled() bool {
	return o != nil && o.AutomuteEnabled != nil
}

// SetAutomuteEnabled gets a reference to the given bool and assigns it to the AutomuteEnabled field.
func (o *AWSMetrics) SetAutomuteEnabled(v bool) {
	o.AutomuteEnabled = &v
}

// GetCollectCloudwatchAlarms returns the CollectCloudwatchAlarms field value if set, zero value otherwise.
func (o *AWSMetrics) GetCollectCloudwatchAlarms() bool {
	if o == nil || o.CollectCloudwatchAlarms == nil {
		var ret bool
		return ret
	}
	return *o.CollectCloudwatchAlarms
}

// GetCollectCloudwatchAlarmsOk returns a tuple with the CollectCloudwatchAlarms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSMetrics) GetCollectCloudwatchAlarmsOk() (*bool, bool) {
	if o == nil || o.CollectCloudwatchAlarms == nil {
		return nil, false
	}
	return o.CollectCloudwatchAlarms, true
}

// HasCollectCloudwatchAlarms returns a boolean if a field has been set.
func (o *AWSMetrics) HasCollectCloudwatchAlarms() bool {
	return o != nil && o.CollectCloudwatchAlarms != nil
}

// SetCollectCloudwatchAlarms gets a reference to the given bool and assigns it to the CollectCloudwatchAlarms field.
func (o *AWSMetrics) SetCollectCloudwatchAlarms(v bool) {
	o.CollectCloudwatchAlarms = &v
}

// GetCollectCustomMetrics returns the CollectCustomMetrics field value if set, zero value otherwise.
func (o *AWSMetrics) GetCollectCustomMetrics() bool {
	if o == nil || o.CollectCustomMetrics == nil {
		var ret bool
		return ret
	}
	return *o.CollectCustomMetrics
}

// GetCollectCustomMetricsOk returns a tuple with the CollectCustomMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSMetrics) GetCollectCustomMetricsOk() (*bool, bool) {
	if o == nil || o.CollectCustomMetrics == nil {
		return nil, false
	}
	return o.CollectCustomMetrics, true
}

// HasCollectCustomMetrics returns a boolean if a field has been set.
func (o *AWSMetrics) HasCollectCustomMetrics() bool {
	return o != nil && o.CollectCustomMetrics != nil
}

// SetCollectCustomMetrics gets a reference to the given bool and assigns it to the CollectCustomMetrics field.
func (o *AWSMetrics) SetCollectCustomMetrics(v bool) {
	o.CollectCustomMetrics = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AWSMetrics) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSMetrics) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AWSMetrics) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AWSMetrics) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetNamespaceFilters returns the NamespaceFilters field value if set, zero value otherwise.
func (o *AWSMetrics) GetNamespaceFilters() AWSNamespacesList {
	if o == nil || o.NamespaceFilters == nil {
		var ret AWSNamespacesList
		return ret
	}
	return *o.NamespaceFilters
}

// GetNamespaceFiltersOk returns a tuple with the NamespaceFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSMetrics) GetNamespaceFiltersOk() (*AWSNamespacesList, bool) {
	if o == nil || o.NamespaceFilters == nil {
		return nil, false
	}
	return o.NamespaceFilters, true
}

// HasNamespaceFilters returns a boolean if a field has been set.
func (o *AWSMetrics) HasNamespaceFilters() bool {
	return o != nil && o.NamespaceFilters != nil
}

// SetNamespaceFilters gets a reference to the given AWSNamespacesList and assigns it to the NamespaceFilters field.
func (o *AWSMetrics) SetNamespaceFilters(v AWSNamespacesList) {
	o.NamespaceFilters = &v
}

// GetTagFilters returns the TagFilters field value if set, zero value otherwise.
func (o *AWSMetrics) GetTagFilters() []AWSNamespaceTagFilter {
	if o == nil || o.TagFilters == nil {
		var ret []AWSNamespaceTagFilter
		return ret
	}
	return o.TagFilters
}

// GetTagFiltersOk returns a tuple with the TagFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSMetrics) GetTagFiltersOk() (*[]AWSNamespaceTagFilter, bool) {
	if o == nil || o.TagFilters == nil {
		return nil, false
	}
	return &o.TagFilters, true
}

// HasTagFilters returns a boolean if a field has been set.
func (o *AWSMetrics) HasTagFilters() bool {
	return o != nil && o.TagFilters != nil
}

// SetTagFilters gets a reference to the given []AWSNamespaceTagFilter and assigns it to the TagFilters field.
func (o *AWSMetrics) SetTagFilters(v []AWSNamespaceTagFilter) {
	o.TagFilters = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSMetrics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AutomuteEnabled != nil {
		toSerialize["automute_enabled"] = o.AutomuteEnabled
	}
	if o.CollectCloudwatchAlarms != nil {
		toSerialize["collect_cloudwatch_alarms"] = o.CollectCloudwatchAlarms
	}
	if o.CollectCustomMetrics != nil {
		toSerialize["collect_custom_metrics"] = o.CollectCustomMetrics
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.NamespaceFilters != nil {
		toSerialize["namespace_filters"] = o.NamespaceFilters
	}
	if o.TagFilters != nil {
		toSerialize["tag_filters"] = o.TagFilters
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSMetrics) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AutomuteEnabled         *bool                   `json:"automute_enabled,omitempty"`
		CollectCloudwatchAlarms *bool                   `json:"collect_cloudwatch_alarms,omitempty"`
		CollectCustomMetrics    *bool                   `json:"collect_custom_metrics,omitempty"`
		Enabled                 *bool                   `json:"enabled,omitempty"`
		NamespaceFilters        *AWSNamespacesList      `json:"namespace_filters,omitempty"`
		TagFilters              []AWSNamespaceTagFilter `json:"tag_filters,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"automute_enabled", "collect_cloudwatch_alarms", "collect_custom_metrics", "enabled", "namespace_filters", "tag_filters"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AutomuteEnabled = all.AutomuteEnabled
	o.CollectCloudwatchAlarms = all.CollectCloudwatchAlarms
	o.CollectCustomMetrics = all.CollectCustomMetrics
	o.Enabled = all.Enabled
	if all.NamespaceFilters != nil && all.NamespaceFilters.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.NamespaceFilters = all.NamespaceFilters
	o.TagFilters = all.TagFilters

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
