// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMSdkConfigRUMUpdateAttributes The RUM SDK settings to apply when updating a configuration.
type RUMSdkConfigRUMUpdateAttributes struct {
	// A list of URL configurations for distributed tracing.
	AllowedTracingUrls []RUMSdkConfigTracingUrlConfig `json:"allowed_tracing_urls,omitempty"`
	// A list of origin patterns allowed for cross-origin session tracking.
	AllowedTrackingOrigins []RUMSdkConfigMatchOption `json:"allowed_tracking_origins,omitempty"`
	// A list of dynamic option key-value pairs.
	Context []RUMSdkConfigDynamicOptionPair `json:"context,omitempty"`
	// The default privacy masking level applied to all RUM data.
	DefaultPrivacyLevel string `json:"default_privacy_level"`
	// Whether to mask user-interaction action names for privacy.
	EnablePrivacyForActionName bool `json:"enable_privacy_for_action_name"`
	// The environment tag for the RUM application.
	Env *string `json:"env,omitempty"`
	// The service name tag for the RUM application.
	Service *string `json:"service,omitempty"`
	// The percentage of collected sessions for which a replay is captured (0–100).
	SessionReplaySampleRate int64 `json:"session_replay_sample_rate"`
	// The percentage of user sessions to collect (0–100).
	SessionSampleRate int64 `json:"session_sample_rate"`
	// The percentage of requests to forward as APM traces (0–100).
	TraceSampleRate *int64 `json:"trace_sample_rate,omitempty"`
	// Whether to share a session across subdomains of the same site.
	TrackSessionAcrossSubdomains *bool `json:"track_session_across_subdomains,omitempty"`
	// A list of dynamic option key-value pairs.
	User []RUMSdkConfigDynamicOptionPair `json:"user,omitempty"`
	// A dynamic configuration option that extracts a value at runtime using a specified strategy.
	Version *RUMSdkConfigDynamicOption `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRUMSdkConfigRUMUpdateAttributes instantiates a new RUMSdkConfigRUMUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRUMSdkConfigRUMUpdateAttributes(defaultPrivacyLevel string, enablePrivacyForActionName bool, sessionReplaySampleRate int64, sessionSampleRate int64) *RUMSdkConfigRUMUpdateAttributes {
	this := RUMSdkConfigRUMUpdateAttributes{}
	this.DefaultPrivacyLevel = defaultPrivacyLevel
	this.EnablePrivacyForActionName = enablePrivacyForActionName
	this.SessionReplaySampleRate = sessionReplaySampleRate
	this.SessionSampleRate = sessionSampleRate
	return &this
}

// NewRUMSdkConfigRUMUpdateAttributesWithDefaults instantiates a new RUMSdkConfigRUMUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRUMSdkConfigRUMUpdateAttributesWithDefaults() *RUMSdkConfigRUMUpdateAttributes {
	this := RUMSdkConfigRUMUpdateAttributes{}
	return &this
}

// GetAllowedTracingUrls returns the AllowedTracingUrls field value if set, zero value otherwise.
func (o *RUMSdkConfigRUMUpdateAttributes) GetAllowedTracingUrls() []RUMSdkConfigTracingUrlConfig {
	if o == nil || o.AllowedTracingUrls == nil {
		var ret []RUMSdkConfigTracingUrlConfig
		return ret
	}
	return o.AllowedTracingUrls
}

// GetAllowedTracingUrlsOk returns a tuple with the AllowedTracingUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMSdkConfigRUMUpdateAttributes) GetAllowedTracingUrlsOk() (*[]RUMSdkConfigTracingUrlConfig, bool) {
	if o == nil || o.AllowedTracingUrls == nil {
		return nil, false
	}
	return &o.AllowedTracingUrls, true
}

// HasAllowedTracingUrls returns a boolean if a field has been set.
func (o *RUMSdkConfigRUMUpdateAttributes) HasAllowedTracingUrls() bool {
	return o != nil && o.AllowedTracingUrls != nil
}

// SetAllowedTracingUrls gets a reference to the given []RUMSdkConfigTracingUrlConfig and assigns it to the AllowedTracingUrls field.
func (o *RUMSdkConfigRUMUpdateAttributes) SetAllowedTracingUrls(v []RUMSdkConfigTracingUrlConfig) {
	o.AllowedTracingUrls = v
}

// GetAllowedTrackingOrigins returns the AllowedTrackingOrigins field value if set, zero value otherwise.
func (o *RUMSdkConfigRUMUpdateAttributes) GetAllowedTrackingOrigins() []RUMSdkConfigMatchOption {
	if o == nil || o.AllowedTrackingOrigins == nil {
		var ret []RUMSdkConfigMatchOption
		return ret
	}
	return o.AllowedTrackingOrigins
}

// GetAllowedTrackingOriginsOk returns a tuple with the AllowedTrackingOrigins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMSdkConfigRUMUpdateAttributes) GetAllowedTrackingOriginsOk() (*[]RUMSdkConfigMatchOption, bool) {
	if o == nil || o.AllowedTrackingOrigins == nil {
		return nil, false
	}
	return &o.AllowedTrackingOrigins, true
}

// HasAllowedTrackingOrigins returns a boolean if a field has been set.
func (o *RUMSdkConfigRUMUpdateAttributes) HasAllowedTrackingOrigins() bool {
	return o != nil && o.AllowedTrackingOrigins != nil
}

// SetAllowedTrackingOrigins gets a reference to the given []RUMSdkConfigMatchOption and assigns it to the AllowedTrackingOrigins field.
func (o *RUMSdkConfigRUMUpdateAttributes) SetAllowedTrackingOrigins(v []RUMSdkConfigMatchOption) {
	o.AllowedTrackingOrigins = v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *RUMSdkConfigRUMUpdateAttributes) GetContext() []RUMSdkConfigDynamicOptionPair {
	if o == nil || o.Context == nil {
		var ret []RUMSdkConfigDynamicOptionPair
		return ret
	}
	return o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMSdkConfigRUMUpdateAttributes) GetContextOk() (*[]RUMSdkConfigDynamicOptionPair, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return &o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *RUMSdkConfigRUMUpdateAttributes) HasContext() bool {
	return o != nil && o.Context != nil
}

// SetContext gets a reference to the given []RUMSdkConfigDynamicOptionPair and assigns it to the Context field.
func (o *RUMSdkConfigRUMUpdateAttributes) SetContext(v []RUMSdkConfigDynamicOptionPair) {
	o.Context = v
}

// GetDefaultPrivacyLevel returns the DefaultPrivacyLevel field value.
func (o *RUMSdkConfigRUMUpdateAttributes) GetDefaultPrivacyLevel() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DefaultPrivacyLevel
}

// GetDefaultPrivacyLevelOk returns a tuple with the DefaultPrivacyLevel field value
// and a boolean to check if the value has been set.
func (o *RUMSdkConfigRUMUpdateAttributes) GetDefaultPrivacyLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultPrivacyLevel, true
}

// SetDefaultPrivacyLevel sets field value.
func (o *RUMSdkConfigRUMUpdateAttributes) SetDefaultPrivacyLevel(v string) {
	o.DefaultPrivacyLevel = v
}

// GetEnablePrivacyForActionName returns the EnablePrivacyForActionName field value.
func (o *RUMSdkConfigRUMUpdateAttributes) GetEnablePrivacyForActionName() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.EnablePrivacyForActionName
}

// GetEnablePrivacyForActionNameOk returns a tuple with the EnablePrivacyForActionName field value
// and a boolean to check if the value has been set.
func (o *RUMSdkConfigRUMUpdateAttributes) GetEnablePrivacyForActionNameOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnablePrivacyForActionName, true
}

// SetEnablePrivacyForActionName sets field value.
func (o *RUMSdkConfigRUMUpdateAttributes) SetEnablePrivacyForActionName(v bool) {
	o.EnablePrivacyForActionName = v
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *RUMSdkConfigRUMUpdateAttributes) GetEnv() string {
	if o == nil || o.Env == nil {
		var ret string
		return ret
	}
	return *o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMSdkConfigRUMUpdateAttributes) GetEnvOk() (*string, bool) {
	if o == nil || o.Env == nil {
		return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *RUMSdkConfigRUMUpdateAttributes) HasEnv() bool {
	return o != nil && o.Env != nil
}

// SetEnv gets a reference to the given string and assigns it to the Env field.
func (o *RUMSdkConfigRUMUpdateAttributes) SetEnv(v string) {
	o.Env = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *RUMSdkConfigRUMUpdateAttributes) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMSdkConfigRUMUpdateAttributes) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *RUMSdkConfigRUMUpdateAttributes) HasService() bool {
	return o != nil && o.Service != nil
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *RUMSdkConfigRUMUpdateAttributes) SetService(v string) {
	o.Service = &v
}

// GetSessionReplaySampleRate returns the SessionReplaySampleRate field value.
func (o *RUMSdkConfigRUMUpdateAttributes) GetSessionReplaySampleRate() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.SessionReplaySampleRate
}

// GetSessionReplaySampleRateOk returns a tuple with the SessionReplaySampleRate field value
// and a boolean to check if the value has been set.
func (o *RUMSdkConfigRUMUpdateAttributes) GetSessionReplaySampleRateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionReplaySampleRate, true
}

// SetSessionReplaySampleRate sets field value.
func (o *RUMSdkConfigRUMUpdateAttributes) SetSessionReplaySampleRate(v int64) {
	o.SessionReplaySampleRate = v
}

// GetSessionSampleRate returns the SessionSampleRate field value.
func (o *RUMSdkConfigRUMUpdateAttributes) GetSessionSampleRate() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.SessionSampleRate
}

// GetSessionSampleRateOk returns a tuple with the SessionSampleRate field value
// and a boolean to check if the value has been set.
func (o *RUMSdkConfigRUMUpdateAttributes) GetSessionSampleRateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionSampleRate, true
}

// SetSessionSampleRate sets field value.
func (o *RUMSdkConfigRUMUpdateAttributes) SetSessionSampleRate(v int64) {
	o.SessionSampleRate = v
}

// GetTraceSampleRate returns the TraceSampleRate field value if set, zero value otherwise.
func (o *RUMSdkConfigRUMUpdateAttributes) GetTraceSampleRate() int64 {
	if o == nil || o.TraceSampleRate == nil {
		var ret int64
		return ret
	}
	return *o.TraceSampleRate
}

// GetTraceSampleRateOk returns a tuple with the TraceSampleRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMSdkConfigRUMUpdateAttributes) GetTraceSampleRateOk() (*int64, bool) {
	if o == nil || o.TraceSampleRate == nil {
		return nil, false
	}
	return o.TraceSampleRate, true
}

// HasTraceSampleRate returns a boolean if a field has been set.
func (o *RUMSdkConfigRUMUpdateAttributes) HasTraceSampleRate() bool {
	return o != nil && o.TraceSampleRate != nil
}

// SetTraceSampleRate gets a reference to the given int64 and assigns it to the TraceSampleRate field.
func (o *RUMSdkConfigRUMUpdateAttributes) SetTraceSampleRate(v int64) {
	o.TraceSampleRate = &v
}

// GetTrackSessionAcrossSubdomains returns the TrackSessionAcrossSubdomains field value if set, zero value otherwise.
func (o *RUMSdkConfigRUMUpdateAttributes) GetTrackSessionAcrossSubdomains() bool {
	if o == nil || o.TrackSessionAcrossSubdomains == nil {
		var ret bool
		return ret
	}
	return *o.TrackSessionAcrossSubdomains
}

// GetTrackSessionAcrossSubdomainsOk returns a tuple with the TrackSessionAcrossSubdomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMSdkConfigRUMUpdateAttributes) GetTrackSessionAcrossSubdomainsOk() (*bool, bool) {
	if o == nil || o.TrackSessionAcrossSubdomains == nil {
		return nil, false
	}
	return o.TrackSessionAcrossSubdomains, true
}

// HasTrackSessionAcrossSubdomains returns a boolean if a field has been set.
func (o *RUMSdkConfigRUMUpdateAttributes) HasTrackSessionAcrossSubdomains() bool {
	return o != nil && o.TrackSessionAcrossSubdomains != nil
}

// SetTrackSessionAcrossSubdomains gets a reference to the given bool and assigns it to the TrackSessionAcrossSubdomains field.
func (o *RUMSdkConfigRUMUpdateAttributes) SetTrackSessionAcrossSubdomains(v bool) {
	o.TrackSessionAcrossSubdomains = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *RUMSdkConfigRUMUpdateAttributes) GetUser() []RUMSdkConfigDynamicOptionPair {
	if o == nil || o.User == nil {
		var ret []RUMSdkConfigDynamicOptionPair
		return ret
	}
	return o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMSdkConfigRUMUpdateAttributes) GetUserOk() (*[]RUMSdkConfigDynamicOptionPair, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return &o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *RUMSdkConfigRUMUpdateAttributes) HasUser() bool {
	return o != nil && o.User != nil
}

// SetUser gets a reference to the given []RUMSdkConfigDynamicOptionPair and assigns it to the User field.
func (o *RUMSdkConfigRUMUpdateAttributes) SetUser(v []RUMSdkConfigDynamicOptionPair) {
	o.User = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *RUMSdkConfigRUMUpdateAttributes) GetVersion() RUMSdkConfigDynamicOption {
	if o == nil || o.Version == nil {
		var ret RUMSdkConfigDynamicOption
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMSdkConfigRUMUpdateAttributes) GetVersionOk() (*RUMSdkConfigDynamicOption, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *RUMSdkConfigRUMUpdateAttributes) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given RUMSdkConfigDynamicOption and assigns it to the Version field.
func (o *RUMSdkConfigRUMUpdateAttributes) SetVersion(v RUMSdkConfigDynamicOption) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RUMSdkConfigRUMUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AllowedTracingUrls != nil {
		toSerialize["allowed_tracing_urls"] = o.AllowedTracingUrls
	}
	if o.AllowedTrackingOrigins != nil {
		toSerialize["allowed_tracking_origins"] = o.AllowedTrackingOrigins
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	toSerialize["default_privacy_level"] = o.DefaultPrivacyLevel
	toSerialize["enable_privacy_for_action_name"] = o.EnablePrivacyForActionName
	if o.Env != nil {
		toSerialize["env"] = o.Env
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	toSerialize["session_replay_sample_rate"] = o.SessionReplaySampleRate
	toSerialize["session_sample_rate"] = o.SessionSampleRate
	if o.TraceSampleRate != nil {
		toSerialize["trace_sample_rate"] = o.TraceSampleRate
	}
	if o.TrackSessionAcrossSubdomains != nil {
		toSerialize["track_session_across_subdomains"] = o.TrackSessionAcrossSubdomains
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RUMSdkConfigRUMUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AllowedTracingUrls           []RUMSdkConfigTracingUrlConfig  `json:"allowed_tracing_urls,omitempty"`
		AllowedTrackingOrigins       []RUMSdkConfigMatchOption       `json:"allowed_tracking_origins,omitempty"`
		Context                      []RUMSdkConfigDynamicOptionPair `json:"context,omitempty"`
		DefaultPrivacyLevel          *string                         `json:"default_privacy_level"`
		EnablePrivacyForActionName   *bool                           `json:"enable_privacy_for_action_name"`
		Env                          *string                         `json:"env,omitempty"`
		Service                      *string                         `json:"service,omitempty"`
		SessionReplaySampleRate      *int64                          `json:"session_replay_sample_rate"`
		SessionSampleRate            *int64                          `json:"session_sample_rate"`
		TraceSampleRate              *int64                          `json:"trace_sample_rate,omitempty"`
		TrackSessionAcrossSubdomains *bool                           `json:"track_session_across_subdomains,omitempty"`
		User                         []RUMSdkConfigDynamicOptionPair `json:"user,omitempty"`
		Version                      *RUMSdkConfigDynamicOption      `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DefaultPrivacyLevel == nil {
		return fmt.Errorf("required field default_privacy_level missing")
	}
	if all.EnablePrivacyForActionName == nil {
		return fmt.Errorf("required field enable_privacy_for_action_name missing")
	}
	if all.SessionReplaySampleRate == nil {
		return fmt.Errorf("required field session_replay_sample_rate missing")
	}
	if all.SessionSampleRate == nil {
		return fmt.Errorf("required field session_sample_rate missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"allowed_tracing_urls", "allowed_tracking_origins", "context", "default_privacy_level", "enable_privacy_for_action_name", "env", "service", "session_replay_sample_rate", "session_sample_rate", "trace_sample_rate", "track_session_across_subdomains", "user", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AllowedTracingUrls = all.AllowedTracingUrls
	o.AllowedTrackingOrigins = all.AllowedTrackingOrigins
	o.Context = all.Context
	o.DefaultPrivacyLevel = *all.DefaultPrivacyLevel
	o.EnablePrivacyForActionName = *all.EnablePrivacyForActionName
	o.Env = all.Env
	o.Service = all.Service
	o.SessionReplaySampleRate = *all.SessionReplaySampleRate
	o.SessionSampleRate = *all.SessionSampleRate
	o.TraceSampleRate = all.TraceSampleRate
	o.TrackSessionAcrossSubdomains = all.TrackSessionAcrossSubdomains
	o.User = all.User
	if all.Version != nil && all.Version.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
