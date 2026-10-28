// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMApplicationCreateAttributes RUM application creation attributes.
// When `rum_event_processing_state` is omitted, it defaults to `ALL`. When `product_analytics_retention_state` is omitted, Product Analytics is not enabled. Setting it to `MAX` requires the `product_analytics_apps_write` permission.
type RUMApplicationCreateAttributes struct {
	// Indicates if flat sampling of Session Replay for APM-correlated RUM sessions is enabled. Defaults to `true`.
	ApmRumFlatSamplingReplayEnabled *bool `json:"apm_rum_flat_sampling_replay_enabled,omitempty"`
	// Sample rate (0-100) of Session Replay for APM-correlated RUM sessions when flat sampling is enabled. Defaults to `1`.
	ApmRumFlatSamplingReplaySampleRate *float64 `json:"apm_rum_flat_sampling_replay_sample_rate,omitempty"`
	// Whether to create the default retention filters when the application is created: `Sessions with errors` and `Sessions`, each with a 100% retention rate.
	// Set to `false` to create the application without them, for example when you add your own retention filters right after creating it. Until you add retention filters, the application only retains sessions matched by its [permanent retention filters](https://docs.datadoghq.com/real_user_monitoring/rum_without_limits/retention_filters/#permanent-retention-filters).
	// This setting doesn't affect permanent retention filters (sessions with forced replays, Synthetics sessions and RUM-APM flat sampling).
	CreateDefaultRetentionFilters *bool `json:"create_default_retention_filters,omitempty"`
	// Indicates if Error Tracking exclusion filters are enabled for the application. Defaults to `false`.
	ErrorTrackingExclusionFilterEnabled *bool `json:"error_tracking_exclusion_filter_enabled,omitempty"`
	// Name of the RUM application. Between 1 and 255 characters.
	Name string `json:"name"`
	// Sample rate (0-100) of Session Replay for Product Analytics sessions. Defaults to `100`.
	ProductAnalyticsReplaySampleRate *int32 `json:"product_analytics_replay_sample_rate,omitempty"`
	// Controls the retention policy for Product Analytics data derived from RUM events.
	// Requests also accept `PREVIEW` (the initial state of a new application, reported as `NONE` in responses). Setting `MAX` requires the `product_analytics_apps_write` permission.
	ProductAnalyticsRetentionState *RUMProductAnalyticsRetentionState `json:"product_analytics_retention_state,omitempty"`
	// Configures which RUM events are processed and stored for the application. Defaults to `ALL` when creating an application.
	RumEventProcessingState *RUMEventProcessingState `json:"rum_event_processing_state,omitempty"`
	// List of tags to associate with the RUM application, in `key:value` format.
	Tags []string `json:"tags,omitempty"`
	// Type of the RUM application. Defaults to `browser`. Supported values are `browser`, `ios`, `android`, `react-native`, `flutter`, `roku`, `electron`, `unity`, `kotlin-multiplatform`, `react`, `nextjs`, `angular`, `vuejs`, `nuxt`, `tanstack`, `maui`, `salesforce-lwc`, `shopify`, `cpp`, `rum-cpp`, `unreal`.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRUMApplicationCreateAttributes instantiates a new RUMApplicationCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRUMApplicationCreateAttributes(name string) *RUMApplicationCreateAttributes {
	this := RUMApplicationCreateAttributes{}
	var createDefaultRetentionFilters bool = true
	this.CreateDefaultRetentionFilters = &createDefaultRetentionFilters
	this.Name = name
	return &this
}

// NewRUMApplicationCreateAttributesWithDefaults instantiates a new RUMApplicationCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRUMApplicationCreateAttributesWithDefaults() *RUMApplicationCreateAttributes {
	this := RUMApplicationCreateAttributes{}
	var createDefaultRetentionFilters bool = true
	this.CreateDefaultRetentionFilters = &createDefaultRetentionFilters
	return &this
}

// GetApmRumFlatSamplingReplayEnabled returns the ApmRumFlatSamplingReplayEnabled field value if set, zero value otherwise.
func (o *RUMApplicationCreateAttributes) GetApmRumFlatSamplingReplayEnabled() bool {
	if o == nil || o.ApmRumFlatSamplingReplayEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ApmRumFlatSamplingReplayEnabled
}

// GetApmRumFlatSamplingReplayEnabledOk returns a tuple with the ApmRumFlatSamplingReplayEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationCreateAttributes) GetApmRumFlatSamplingReplayEnabledOk() (*bool, bool) {
	if o == nil || o.ApmRumFlatSamplingReplayEnabled == nil {
		return nil, false
	}
	return o.ApmRumFlatSamplingReplayEnabled, true
}

// HasApmRumFlatSamplingReplayEnabled returns a boolean if a field has been set.
func (o *RUMApplicationCreateAttributes) HasApmRumFlatSamplingReplayEnabled() bool {
	return o != nil && o.ApmRumFlatSamplingReplayEnabled != nil
}

// SetApmRumFlatSamplingReplayEnabled gets a reference to the given bool and assigns it to the ApmRumFlatSamplingReplayEnabled field.
func (o *RUMApplicationCreateAttributes) SetApmRumFlatSamplingReplayEnabled(v bool) {
	o.ApmRumFlatSamplingReplayEnabled = &v
}

// GetApmRumFlatSamplingReplaySampleRate returns the ApmRumFlatSamplingReplaySampleRate field value if set, zero value otherwise.
func (o *RUMApplicationCreateAttributes) GetApmRumFlatSamplingReplaySampleRate() float64 {
	if o == nil || o.ApmRumFlatSamplingReplaySampleRate == nil {
		var ret float64
		return ret
	}
	return *o.ApmRumFlatSamplingReplaySampleRate
}

// GetApmRumFlatSamplingReplaySampleRateOk returns a tuple with the ApmRumFlatSamplingReplaySampleRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationCreateAttributes) GetApmRumFlatSamplingReplaySampleRateOk() (*float64, bool) {
	if o == nil || o.ApmRumFlatSamplingReplaySampleRate == nil {
		return nil, false
	}
	return o.ApmRumFlatSamplingReplaySampleRate, true
}

// HasApmRumFlatSamplingReplaySampleRate returns a boolean if a field has been set.
func (o *RUMApplicationCreateAttributes) HasApmRumFlatSamplingReplaySampleRate() bool {
	return o != nil && o.ApmRumFlatSamplingReplaySampleRate != nil
}

// SetApmRumFlatSamplingReplaySampleRate gets a reference to the given float64 and assigns it to the ApmRumFlatSamplingReplaySampleRate field.
func (o *RUMApplicationCreateAttributes) SetApmRumFlatSamplingReplaySampleRate(v float64) {
	o.ApmRumFlatSamplingReplaySampleRate = &v
}

// GetCreateDefaultRetentionFilters returns the CreateDefaultRetentionFilters field value if set, zero value otherwise.
func (o *RUMApplicationCreateAttributes) GetCreateDefaultRetentionFilters() bool {
	if o == nil || o.CreateDefaultRetentionFilters == nil {
		var ret bool
		return ret
	}
	return *o.CreateDefaultRetentionFilters
}

// GetCreateDefaultRetentionFiltersOk returns a tuple with the CreateDefaultRetentionFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationCreateAttributes) GetCreateDefaultRetentionFiltersOk() (*bool, bool) {
	if o == nil || o.CreateDefaultRetentionFilters == nil {
		return nil, false
	}
	return o.CreateDefaultRetentionFilters, true
}

// HasCreateDefaultRetentionFilters returns a boolean if a field has been set.
func (o *RUMApplicationCreateAttributes) HasCreateDefaultRetentionFilters() bool {
	return o != nil && o.CreateDefaultRetentionFilters != nil
}

// SetCreateDefaultRetentionFilters gets a reference to the given bool and assigns it to the CreateDefaultRetentionFilters field.
func (o *RUMApplicationCreateAttributes) SetCreateDefaultRetentionFilters(v bool) {
	o.CreateDefaultRetentionFilters = &v
}

// GetErrorTrackingExclusionFilterEnabled returns the ErrorTrackingExclusionFilterEnabled field value if set, zero value otherwise.
func (o *RUMApplicationCreateAttributes) GetErrorTrackingExclusionFilterEnabled() bool {
	if o == nil || o.ErrorTrackingExclusionFilterEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ErrorTrackingExclusionFilterEnabled
}

// GetErrorTrackingExclusionFilterEnabledOk returns a tuple with the ErrorTrackingExclusionFilterEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationCreateAttributes) GetErrorTrackingExclusionFilterEnabledOk() (*bool, bool) {
	if o == nil || o.ErrorTrackingExclusionFilterEnabled == nil {
		return nil, false
	}
	return o.ErrorTrackingExclusionFilterEnabled, true
}

// HasErrorTrackingExclusionFilterEnabled returns a boolean if a field has been set.
func (o *RUMApplicationCreateAttributes) HasErrorTrackingExclusionFilterEnabled() bool {
	return o != nil && o.ErrorTrackingExclusionFilterEnabled != nil
}

// SetErrorTrackingExclusionFilterEnabled gets a reference to the given bool and assigns it to the ErrorTrackingExclusionFilterEnabled field.
func (o *RUMApplicationCreateAttributes) SetErrorTrackingExclusionFilterEnabled(v bool) {
	o.ErrorTrackingExclusionFilterEnabled = &v
}

// GetName returns the Name field value.
func (o *RUMApplicationCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RUMApplicationCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *RUMApplicationCreateAttributes) SetName(v string) {
	o.Name = v
}

// GetProductAnalyticsReplaySampleRate returns the ProductAnalyticsReplaySampleRate field value if set, zero value otherwise.
func (o *RUMApplicationCreateAttributes) GetProductAnalyticsReplaySampleRate() int32 {
	if o == nil || o.ProductAnalyticsReplaySampleRate == nil {
		var ret int32
		return ret
	}
	return *o.ProductAnalyticsReplaySampleRate
}

// GetProductAnalyticsReplaySampleRateOk returns a tuple with the ProductAnalyticsReplaySampleRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationCreateAttributes) GetProductAnalyticsReplaySampleRateOk() (*int32, bool) {
	if o == nil || o.ProductAnalyticsReplaySampleRate == nil {
		return nil, false
	}
	return o.ProductAnalyticsReplaySampleRate, true
}

// HasProductAnalyticsReplaySampleRate returns a boolean if a field has been set.
func (o *RUMApplicationCreateAttributes) HasProductAnalyticsReplaySampleRate() bool {
	return o != nil && o.ProductAnalyticsReplaySampleRate != nil
}

// SetProductAnalyticsReplaySampleRate gets a reference to the given int32 and assigns it to the ProductAnalyticsReplaySampleRate field.
func (o *RUMApplicationCreateAttributes) SetProductAnalyticsReplaySampleRate(v int32) {
	o.ProductAnalyticsReplaySampleRate = &v
}

// GetProductAnalyticsRetentionState returns the ProductAnalyticsRetentionState field value if set, zero value otherwise.
func (o *RUMApplicationCreateAttributes) GetProductAnalyticsRetentionState() RUMProductAnalyticsRetentionState {
	if o == nil || o.ProductAnalyticsRetentionState == nil {
		var ret RUMProductAnalyticsRetentionState
		return ret
	}
	return *o.ProductAnalyticsRetentionState
}

// GetProductAnalyticsRetentionStateOk returns a tuple with the ProductAnalyticsRetentionState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationCreateAttributes) GetProductAnalyticsRetentionStateOk() (*RUMProductAnalyticsRetentionState, bool) {
	if o == nil || o.ProductAnalyticsRetentionState == nil {
		return nil, false
	}
	return o.ProductAnalyticsRetentionState, true
}

// HasProductAnalyticsRetentionState returns a boolean if a field has been set.
func (o *RUMApplicationCreateAttributes) HasProductAnalyticsRetentionState() bool {
	return o != nil && o.ProductAnalyticsRetentionState != nil
}

// SetProductAnalyticsRetentionState gets a reference to the given RUMProductAnalyticsRetentionState and assigns it to the ProductAnalyticsRetentionState field.
func (o *RUMApplicationCreateAttributes) SetProductAnalyticsRetentionState(v RUMProductAnalyticsRetentionState) {
	o.ProductAnalyticsRetentionState = &v
}

// GetRumEventProcessingState returns the RumEventProcessingState field value if set, zero value otherwise.
func (o *RUMApplicationCreateAttributes) GetRumEventProcessingState() RUMEventProcessingState {
	if o == nil || o.RumEventProcessingState == nil {
		var ret RUMEventProcessingState
		return ret
	}
	return *o.RumEventProcessingState
}

// GetRumEventProcessingStateOk returns a tuple with the RumEventProcessingState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationCreateAttributes) GetRumEventProcessingStateOk() (*RUMEventProcessingState, bool) {
	if o == nil || o.RumEventProcessingState == nil {
		return nil, false
	}
	return o.RumEventProcessingState, true
}

// HasRumEventProcessingState returns a boolean if a field has been set.
func (o *RUMApplicationCreateAttributes) HasRumEventProcessingState() bool {
	return o != nil && o.RumEventProcessingState != nil
}

// SetRumEventProcessingState gets a reference to the given RUMEventProcessingState and assigns it to the RumEventProcessingState field.
func (o *RUMApplicationCreateAttributes) SetRumEventProcessingState(v RUMEventProcessingState) {
	o.RumEventProcessingState = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *RUMApplicationCreateAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationCreateAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *RUMApplicationCreateAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *RUMApplicationCreateAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RUMApplicationCreateAttributes) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationCreateAttributes) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RUMApplicationCreateAttributes) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RUMApplicationCreateAttributes) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RUMApplicationCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ApmRumFlatSamplingReplayEnabled != nil {
		toSerialize["apm_rum_flat_sampling_replay_enabled"] = o.ApmRumFlatSamplingReplayEnabled
	}
	if o.ApmRumFlatSamplingReplaySampleRate != nil {
		toSerialize["apm_rum_flat_sampling_replay_sample_rate"] = o.ApmRumFlatSamplingReplaySampleRate
	}
	if o.CreateDefaultRetentionFilters != nil {
		toSerialize["create_default_retention_filters"] = o.CreateDefaultRetentionFilters
	}
	if o.ErrorTrackingExclusionFilterEnabled != nil {
		toSerialize["error_tracking_exclusion_filter_enabled"] = o.ErrorTrackingExclusionFilterEnabled
	}
	toSerialize["name"] = o.Name
	if o.ProductAnalyticsReplaySampleRate != nil {
		toSerialize["product_analytics_replay_sample_rate"] = o.ProductAnalyticsReplaySampleRate
	}
	if o.ProductAnalyticsRetentionState != nil {
		toSerialize["product_analytics_retention_state"] = o.ProductAnalyticsRetentionState
	}
	if o.RumEventProcessingState != nil {
		toSerialize["rum_event_processing_state"] = o.RumEventProcessingState
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RUMApplicationCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApmRumFlatSamplingReplayEnabled     *bool                              `json:"apm_rum_flat_sampling_replay_enabled,omitempty"`
		ApmRumFlatSamplingReplaySampleRate  *float64                           `json:"apm_rum_flat_sampling_replay_sample_rate,omitempty"`
		CreateDefaultRetentionFilters       *bool                              `json:"create_default_retention_filters,omitempty"`
		ErrorTrackingExclusionFilterEnabled *bool                              `json:"error_tracking_exclusion_filter_enabled,omitempty"`
		Name                                *string                            `json:"name"`
		ProductAnalyticsReplaySampleRate    *int32                             `json:"product_analytics_replay_sample_rate,omitempty"`
		ProductAnalyticsRetentionState      *RUMProductAnalyticsRetentionState `json:"product_analytics_retention_state,omitempty"`
		RumEventProcessingState             *RUMEventProcessingState           `json:"rum_event_processing_state,omitempty"`
		Tags                                []string                           `json:"tags,omitempty"`
		Type                                *string                            `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"apm_rum_flat_sampling_replay_enabled", "apm_rum_flat_sampling_replay_sample_rate", "create_default_retention_filters", "error_tracking_exclusion_filter_enabled", "name", "product_analytics_replay_sample_rate", "product_analytics_retention_state", "rum_event_processing_state", "tags", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ApmRumFlatSamplingReplayEnabled = all.ApmRumFlatSamplingReplayEnabled
	o.ApmRumFlatSamplingReplaySampleRate = all.ApmRumFlatSamplingReplaySampleRate
	o.CreateDefaultRetentionFilters = all.CreateDefaultRetentionFilters
	o.ErrorTrackingExclusionFilterEnabled = all.ErrorTrackingExclusionFilterEnabled
	o.Name = *all.Name
	o.ProductAnalyticsReplaySampleRate = all.ProductAnalyticsReplaySampleRate
	if all.ProductAnalyticsRetentionState != nil && !all.ProductAnalyticsRetentionState.IsValid() {
		hasInvalidField = true
	} else {
		o.ProductAnalyticsRetentionState = all.ProductAnalyticsRetentionState
	}
	if all.RumEventProcessingState != nil && !all.RumEventProcessingState.IsValid() {
		hasInvalidField = true
	} else {
		o.RumEventProcessingState = all.RumEventProcessingState
	}
	o.Tags = all.Tags
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
