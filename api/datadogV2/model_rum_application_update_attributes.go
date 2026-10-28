// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMApplicationUpdateAttributes RUM application update attributes. Only the attributes you provide are changed, omitted attributes keep their current value.
// Changing `product_analytics_retention_state` or `product_analytics_replay_sample_rate` requires the `product_analytics_apps_write` permission.
type RUMApplicationUpdateAttributes struct {
	// Indicates if flat sampling of Session Replay for APM-correlated RUM sessions is enabled.
	ApmRumFlatSamplingReplayEnabled *bool `json:"apm_rum_flat_sampling_replay_enabled,omitempty"`
	// Sample rate (0-100) of Session Replay for APM-correlated RUM sessions when flat sampling is enabled.
	ApmRumFlatSamplingReplaySampleRate *float64 `json:"apm_rum_flat_sampling_replay_sample_rate,omitempty"`
	// Indicates if Error Tracking exclusion filters are enabled for the application.
	ErrorTrackingExclusionFilterEnabled *bool `json:"error_tracking_exclusion_filter_enabled,omitempty"`
	// Instrumentation type of the RUM application.
	InstrumentationType *string `json:"instrumentation_type,omitempty"`
	// New name of the RUM application.
	Name *string `json:"name,omitempty"`
	// Sample rate (0-100) of Session Replay for Product Analytics sessions.
	ProductAnalyticsReplaySampleRate *int32 `json:"product_analytics_replay_sample_rate,omitempty"`
	// Controls the retention policy for Product Analytics data derived from RUM events.
	// Requests also accept `PREVIEW` (the initial state of a new application, reported as `NONE` in responses). Setting `MAX` requires the `product_analytics_apps_write` permission.
	ProductAnalyticsRetentionState *RUMProductAnalyticsRetentionState `json:"product_analytics_retention_state,omitempty"`
	// ID of the RUM SDK remote configuration to associate with the application.
	RemoteConfigId *string `json:"remote_config_id,omitempty"`
	// Configures which RUM events are processed and stored for the application. Defaults to `ALL` when creating an application.
	RumEventProcessingState *RUMEventProcessingState `json:"rum_event_processing_state,omitempty"`
	// New list of tags of the RUM application, in `key:value` format. Replaces the existing tags.
	Tags []string `json:"tags,omitempty"`
	// New type of the RUM application. Supported values are `browser`, `ios`, `android`, `react-native`, `flutter`, `roku`, `electron`, `unity`, `kotlin-multiplatform`, `react`, `nextjs`, `angular`, `vuejs`, `nuxt`, `tanstack`, `maui`, `salesforce-lwc`, `shopify`, `cpp`, `rum-cpp`, `unreal`.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRUMApplicationUpdateAttributes instantiates a new RUMApplicationUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRUMApplicationUpdateAttributes() *RUMApplicationUpdateAttributes {
	this := RUMApplicationUpdateAttributes{}
	return &this
}

// NewRUMApplicationUpdateAttributesWithDefaults instantiates a new RUMApplicationUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRUMApplicationUpdateAttributesWithDefaults() *RUMApplicationUpdateAttributes {
	this := RUMApplicationUpdateAttributes{}
	return &this
}

// GetApmRumFlatSamplingReplayEnabled returns the ApmRumFlatSamplingReplayEnabled field value if set, zero value otherwise.
func (o *RUMApplicationUpdateAttributes) GetApmRumFlatSamplingReplayEnabled() bool {
	if o == nil || o.ApmRumFlatSamplingReplayEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ApmRumFlatSamplingReplayEnabled
}

// GetApmRumFlatSamplingReplayEnabledOk returns a tuple with the ApmRumFlatSamplingReplayEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationUpdateAttributes) GetApmRumFlatSamplingReplayEnabledOk() (*bool, bool) {
	if o == nil || o.ApmRumFlatSamplingReplayEnabled == nil {
		return nil, false
	}
	return o.ApmRumFlatSamplingReplayEnabled, true
}

// HasApmRumFlatSamplingReplayEnabled returns a boolean if a field has been set.
func (o *RUMApplicationUpdateAttributes) HasApmRumFlatSamplingReplayEnabled() bool {
	return o != nil && o.ApmRumFlatSamplingReplayEnabled != nil
}

// SetApmRumFlatSamplingReplayEnabled gets a reference to the given bool and assigns it to the ApmRumFlatSamplingReplayEnabled field.
func (o *RUMApplicationUpdateAttributes) SetApmRumFlatSamplingReplayEnabled(v bool) {
	o.ApmRumFlatSamplingReplayEnabled = &v
}

// GetApmRumFlatSamplingReplaySampleRate returns the ApmRumFlatSamplingReplaySampleRate field value if set, zero value otherwise.
func (o *RUMApplicationUpdateAttributes) GetApmRumFlatSamplingReplaySampleRate() float64 {
	if o == nil || o.ApmRumFlatSamplingReplaySampleRate == nil {
		var ret float64
		return ret
	}
	return *o.ApmRumFlatSamplingReplaySampleRate
}

// GetApmRumFlatSamplingReplaySampleRateOk returns a tuple with the ApmRumFlatSamplingReplaySampleRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationUpdateAttributes) GetApmRumFlatSamplingReplaySampleRateOk() (*float64, bool) {
	if o == nil || o.ApmRumFlatSamplingReplaySampleRate == nil {
		return nil, false
	}
	return o.ApmRumFlatSamplingReplaySampleRate, true
}

// HasApmRumFlatSamplingReplaySampleRate returns a boolean if a field has been set.
func (o *RUMApplicationUpdateAttributes) HasApmRumFlatSamplingReplaySampleRate() bool {
	return o != nil && o.ApmRumFlatSamplingReplaySampleRate != nil
}

// SetApmRumFlatSamplingReplaySampleRate gets a reference to the given float64 and assigns it to the ApmRumFlatSamplingReplaySampleRate field.
func (o *RUMApplicationUpdateAttributes) SetApmRumFlatSamplingReplaySampleRate(v float64) {
	o.ApmRumFlatSamplingReplaySampleRate = &v
}

// GetErrorTrackingExclusionFilterEnabled returns the ErrorTrackingExclusionFilterEnabled field value if set, zero value otherwise.
func (o *RUMApplicationUpdateAttributes) GetErrorTrackingExclusionFilterEnabled() bool {
	if o == nil || o.ErrorTrackingExclusionFilterEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ErrorTrackingExclusionFilterEnabled
}

// GetErrorTrackingExclusionFilterEnabledOk returns a tuple with the ErrorTrackingExclusionFilterEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationUpdateAttributes) GetErrorTrackingExclusionFilterEnabledOk() (*bool, bool) {
	if o == nil || o.ErrorTrackingExclusionFilterEnabled == nil {
		return nil, false
	}
	return o.ErrorTrackingExclusionFilterEnabled, true
}

// HasErrorTrackingExclusionFilterEnabled returns a boolean if a field has been set.
func (o *RUMApplicationUpdateAttributes) HasErrorTrackingExclusionFilterEnabled() bool {
	return o != nil && o.ErrorTrackingExclusionFilterEnabled != nil
}

// SetErrorTrackingExclusionFilterEnabled gets a reference to the given bool and assigns it to the ErrorTrackingExclusionFilterEnabled field.
func (o *RUMApplicationUpdateAttributes) SetErrorTrackingExclusionFilterEnabled(v bool) {
	o.ErrorTrackingExclusionFilterEnabled = &v
}

// GetInstrumentationType returns the InstrumentationType field value if set, zero value otherwise.
func (o *RUMApplicationUpdateAttributes) GetInstrumentationType() string {
	if o == nil || o.InstrumentationType == nil {
		var ret string
		return ret
	}
	return *o.InstrumentationType
}

// GetInstrumentationTypeOk returns a tuple with the InstrumentationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationUpdateAttributes) GetInstrumentationTypeOk() (*string, bool) {
	if o == nil || o.InstrumentationType == nil {
		return nil, false
	}
	return o.InstrumentationType, true
}

// HasInstrumentationType returns a boolean if a field has been set.
func (o *RUMApplicationUpdateAttributes) HasInstrumentationType() bool {
	return o != nil && o.InstrumentationType != nil
}

// SetInstrumentationType gets a reference to the given string and assigns it to the InstrumentationType field.
func (o *RUMApplicationUpdateAttributes) SetInstrumentationType(v string) {
	o.InstrumentationType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RUMApplicationUpdateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RUMApplicationUpdateAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RUMApplicationUpdateAttributes) SetName(v string) {
	o.Name = &v
}

// GetProductAnalyticsReplaySampleRate returns the ProductAnalyticsReplaySampleRate field value if set, zero value otherwise.
func (o *RUMApplicationUpdateAttributes) GetProductAnalyticsReplaySampleRate() int32 {
	if o == nil || o.ProductAnalyticsReplaySampleRate == nil {
		var ret int32
		return ret
	}
	return *o.ProductAnalyticsReplaySampleRate
}

// GetProductAnalyticsReplaySampleRateOk returns a tuple with the ProductAnalyticsReplaySampleRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationUpdateAttributes) GetProductAnalyticsReplaySampleRateOk() (*int32, bool) {
	if o == nil || o.ProductAnalyticsReplaySampleRate == nil {
		return nil, false
	}
	return o.ProductAnalyticsReplaySampleRate, true
}

// HasProductAnalyticsReplaySampleRate returns a boolean if a field has been set.
func (o *RUMApplicationUpdateAttributes) HasProductAnalyticsReplaySampleRate() bool {
	return o != nil && o.ProductAnalyticsReplaySampleRate != nil
}

// SetProductAnalyticsReplaySampleRate gets a reference to the given int32 and assigns it to the ProductAnalyticsReplaySampleRate field.
func (o *RUMApplicationUpdateAttributes) SetProductAnalyticsReplaySampleRate(v int32) {
	o.ProductAnalyticsReplaySampleRate = &v
}

// GetProductAnalyticsRetentionState returns the ProductAnalyticsRetentionState field value if set, zero value otherwise.
func (o *RUMApplicationUpdateAttributes) GetProductAnalyticsRetentionState() RUMProductAnalyticsRetentionState {
	if o == nil || o.ProductAnalyticsRetentionState == nil {
		var ret RUMProductAnalyticsRetentionState
		return ret
	}
	return *o.ProductAnalyticsRetentionState
}

// GetProductAnalyticsRetentionStateOk returns a tuple with the ProductAnalyticsRetentionState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationUpdateAttributes) GetProductAnalyticsRetentionStateOk() (*RUMProductAnalyticsRetentionState, bool) {
	if o == nil || o.ProductAnalyticsRetentionState == nil {
		return nil, false
	}
	return o.ProductAnalyticsRetentionState, true
}

// HasProductAnalyticsRetentionState returns a boolean if a field has been set.
func (o *RUMApplicationUpdateAttributes) HasProductAnalyticsRetentionState() bool {
	return o != nil && o.ProductAnalyticsRetentionState != nil
}

// SetProductAnalyticsRetentionState gets a reference to the given RUMProductAnalyticsRetentionState and assigns it to the ProductAnalyticsRetentionState field.
func (o *RUMApplicationUpdateAttributes) SetProductAnalyticsRetentionState(v RUMProductAnalyticsRetentionState) {
	o.ProductAnalyticsRetentionState = &v
}

// GetRemoteConfigId returns the RemoteConfigId field value if set, zero value otherwise.
func (o *RUMApplicationUpdateAttributes) GetRemoteConfigId() string {
	if o == nil || o.RemoteConfigId == nil {
		var ret string
		return ret
	}
	return *o.RemoteConfigId
}

// GetRemoteConfigIdOk returns a tuple with the RemoteConfigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationUpdateAttributes) GetRemoteConfigIdOk() (*string, bool) {
	if o == nil || o.RemoteConfigId == nil {
		return nil, false
	}
	return o.RemoteConfigId, true
}

// HasRemoteConfigId returns a boolean if a field has been set.
func (o *RUMApplicationUpdateAttributes) HasRemoteConfigId() bool {
	return o != nil && o.RemoteConfigId != nil
}

// SetRemoteConfigId gets a reference to the given string and assigns it to the RemoteConfigId field.
func (o *RUMApplicationUpdateAttributes) SetRemoteConfigId(v string) {
	o.RemoteConfigId = &v
}

// GetRumEventProcessingState returns the RumEventProcessingState field value if set, zero value otherwise.
func (o *RUMApplicationUpdateAttributes) GetRumEventProcessingState() RUMEventProcessingState {
	if o == nil || o.RumEventProcessingState == nil {
		var ret RUMEventProcessingState
		return ret
	}
	return *o.RumEventProcessingState
}

// GetRumEventProcessingStateOk returns a tuple with the RumEventProcessingState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationUpdateAttributes) GetRumEventProcessingStateOk() (*RUMEventProcessingState, bool) {
	if o == nil || o.RumEventProcessingState == nil {
		return nil, false
	}
	return o.RumEventProcessingState, true
}

// HasRumEventProcessingState returns a boolean if a field has been set.
func (o *RUMApplicationUpdateAttributes) HasRumEventProcessingState() bool {
	return o != nil && o.RumEventProcessingState != nil
}

// SetRumEventProcessingState gets a reference to the given RUMEventProcessingState and assigns it to the RumEventProcessingState field.
func (o *RUMApplicationUpdateAttributes) SetRumEventProcessingState(v RUMEventProcessingState) {
	o.RumEventProcessingState = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *RUMApplicationUpdateAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationUpdateAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *RUMApplicationUpdateAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *RUMApplicationUpdateAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RUMApplicationUpdateAttributes) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationUpdateAttributes) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RUMApplicationUpdateAttributes) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RUMApplicationUpdateAttributes) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RUMApplicationUpdateAttributes) MarshalJSON() ([]byte, error) {
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
	if o.ErrorTrackingExclusionFilterEnabled != nil {
		toSerialize["error_tracking_exclusion_filter_enabled"] = o.ErrorTrackingExclusionFilterEnabled
	}
	if o.InstrumentationType != nil {
		toSerialize["instrumentation_type"] = o.InstrumentationType
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ProductAnalyticsReplaySampleRate != nil {
		toSerialize["product_analytics_replay_sample_rate"] = o.ProductAnalyticsReplaySampleRate
	}
	if o.ProductAnalyticsRetentionState != nil {
		toSerialize["product_analytics_retention_state"] = o.ProductAnalyticsRetentionState
	}
	if o.RemoteConfigId != nil {
		toSerialize["remote_config_id"] = o.RemoteConfigId
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
func (o *RUMApplicationUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApmRumFlatSamplingReplayEnabled     *bool                              `json:"apm_rum_flat_sampling_replay_enabled,omitempty"`
		ApmRumFlatSamplingReplaySampleRate  *float64                           `json:"apm_rum_flat_sampling_replay_sample_rate,omitempty"`
		ErrorTrackingExclusionFilterEnabled *bool                              `json:"error_tracking_exclusion_filter_enabled,omitempty"`
		InstrumentationType                 *string                            `json:"instrumentation_type,omitempty"`
		Name                                *string                            `json:"name,omitempty"`
		ProductAnalyticsReplaySampleRate    *int32                             `json:"product_analytics_replay_sample_rate,omitempty"`
		ProductAnalyticsRetentionState      *RUMProductAnalyticsRetentionState `json:"product_analytics_retention_state,omitempty"`
		RemoteConfigId                      *string                            `json:"remote_config_id,omitempty"`
		RumEventProcessingState             *RUMEventProcessingState           `json:"rum_event_processing_state,omitempty"`
		Tags                                []string                           `json:"tags,omitempty"`
		Type                                *string                            `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"apm_rum_flat_sampling_replay_enabled", "apm_rum_flat_sampling_replay_sample_rate", "error_tracking_exclusion_filter_enabled", "instrumentation_type", "name", "product_analytics_replay_sample_rate", "product_analytics_retention_state", "remote_config_id", "rum_event_processing_state", "tags", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ApmRumFlatSamplingReplayEnabled = all.ApmRumFlatSamplingReplayEnabled
	o.ApmRumFlatSamplingReplaySampleRate = all.ApmRumFlatSamplingReplaySampleRate
	o.ErrorTrackingExclusionFilterEnabled = all.ErrorTrackingExclusionFilterEnabled
	o.InstrumentationType = all.InstrumentationType
	o.Name = all.Name
	o.ProductAnalyticsReplaySampleRate = all.ProductAnalyticsReplaySampleRate
	if all.ProductAnalyticsRetentionState != nil && !all.ProductAnalyticsRetentionState.IsValid() {
		hasInvalidField = true
	} else {
		o.ProductAnalyticsRetentionState = all.ProductAnalyticsRetentionState
	}
	o.RemoteConfigId = all.RemoteConfigId
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
