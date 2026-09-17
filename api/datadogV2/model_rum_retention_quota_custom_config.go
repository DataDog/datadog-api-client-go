// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRetentionQuotaCustomConfig The configuration used when `mode` is `custom`.
type RumRetentionQuotaCustomConfig struct {
	// The time of day when the daily quota resets, in `HH:MM` 24-hour format.
	DailyResetTime string `json:"daily_reset_time"`
	// The timezone offset used for the daily reset time, in `±HH:MM` format.
	DailyResetTimezone string `json:"daily_reset_timezone"`
	// The action to take when the session quota is reached.
	QuotaReachedAction RumRetentionQuotaReachedAction `json:"quota_reached_action"`
	// The maximum number of sessions allowed within the window. Must be at least `1000`.
	SessionLimit int64 `json:"session_limit"`
	// The window type over which the session limit is enforced.
	WindowType RumRetentionQuotaWindowType `json:"window_type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumRetentionQuotaCustomConfig instantiates a new RumRetentionQuotaCustomConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumRetentionQuotaCustomConfig(dailyResetTime string, dailyResetTimezone string, quotaReachedAction RumRetentionQuotaReachedAction, sessionLimit int64, windowType RumRetentionQuotaWindowType) *RumRetentionQuotaCustomConfig {
	this := RumRetentionQuotaCustomConfig{}
	this.DailyResetTime = dailyResetTime
	this.DailyResetTimezone = dailyResetTimezone
	this.QuotaReachedAction = quotaReachedAction
	this.SessionLimit = sessionLimit
	this.WindowType = windowType
	return &this
}

// NewRumRetentionQuotaCustomConfigWithDefaults instantiates a new RumRetentionQuotaCustomConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumRetentionQuotaCustomConfigWithDefaults() *RumRetentionQuotaCustomConfig {
	this := RumRetentionQuotaCustomConfig{}
	return &this
}

// GetDailyResetTime returns the DailyResetTime field value.
func (o *RumRetentionQuotaCustomConfig) GetDailyResetTime() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DailyResetTime
}

// GetDailyResetTimeOk returns a tuple with the DailyResetTime field value
// and a boolean to check if the value has been set.
func (o *RumRetentionQuotaCustomConfig) GetDailyResetTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DailyResetTime, true
}

// SetDailyResetTime sets field value.
func (o *RumRetentionQuotaCustomConfig) SetDailyResetTime(v string) {
	o.DailyResetTime = v
}

// GetDailyResetTimezone returns the DailyResetTimezone field value.
func (o *RumRetentionQuotaCustomConfig) GetDailyResetTimezone() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DailyResetTimezone
}

// GetDailyResetTimezoneOk returns a tuple with the DailyResetTimezone field value
// and a boolean to check if the value has been set.
func (o *RumRetentionQuotaCustomConfig) GetDailyResetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DailyResetTimezone, true
}

// SetDailyResetTimezone sets field value.
func (o *RumRetentionQuotaCustomConfig) SetDailyResetTimezone(v string) {
	o.DailyResetTimezone = v
}

// GetQuotaReachedAction returns the QuotaReachedAction field value.
func (o *RumRetentionQuotaCustomConfig) GetQuotaReachedAction() RumRetentionQuotaReachedAction {
	if o == nil {
		var ret RumRetentionQuotaReachedAction
		return ret
	}
	return o.QuotaReachedAction
}

// GetQuotaReachedActionOk returns a tuple with the QuotaReachedAction field value
// and a boolean to check if the value has been set.
func (o *RumRetentionQuotaCustomConfig) GetQuotaReachedActionOk() (*RumRetentionQuotaReachedAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QuotaReachedAction, true
}

// SetQuotaReachedAction sets field value.
func (o *RumRetentionQuotaCustomConfig) SetQuotaReachedAction(v RumRetentionQuotaReachedAction) {
	o.QuotaReachedAction = v
}

// GetSessionLimit returns the SessionLimit field value.
func (o *RumRetentionQuotaCustomConfig) GetSessionLimit() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.SessionLimit
}

// GetSessionLimitOk returns a tuple with the SessionLimit field value
// and a boolean to check if the value has been set.
func (o *RumRetentionQuotaCustomConfig) GetSessionLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionLimit, true
}

// SetSessionLimit sets field value.
func (o *RumRetentionQuotaCustomConfig) SetSessionLimit(v int64) {
	o.SessionLimit = v
}

// GetWindowType returns the WindowType field value.
func (o *RumRetentionQuotaCustomConfig) GetWindowType() RumRetentionQuotaWindowType {
	if o == nil {
		var ret RumRetentionQuotaWindowType
		return ret
	}
	return o.WindowType
}

// GetWindowTypeOk returns a tuple with the WindowType field value
// and a boolean to check if the value has been set.
func (o *RumRetentionQuotaCustomConfig) GetWindowTypeOk() (*RumRetentionQuotaWindowType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WindowType, true
}

// SetWindowType sets field value.
func (o *RumRetentionQuotaCustomConfig) SetWindowType(v RumRetentionQuotaWindowType) {
	o.WindowType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumRetentionQuotaCustomConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["daily_reset_time"] = o.DailyResetTime
	toSerialize["daily_reset_timezone"] = o.DailyResetTimezone
	toSerialize["quota_reached_action"] = o.QuotaReachedAction
	toSerialize["session_limit"] = o.SessionLimit
	toSerialize["window_type"] = o.WindowType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumRetentionQuotaCustomConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DailyResetTime     *string                         `json:"daily_reset_time"`
		DailyResetTimezone *string                         `json:"daily_reset_timezone"`
		QuotaReachedAction *RumRetentionQuotaReachedAction `json:"quota_reached_action"`
		SessionLimit       *int64                          `json:"session_limit"`
		WindowType         *RumRetentionQuotaWindowType    `json:"window_type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DailyResetTime == nil {
		return fmt.Errorf("required field daily_reset_time missing")
	}
	if all.DailyResetTimezone == nil {
		return fmt.Errorf("required field daily_reset_timezone missing")
	}
	if all.QuotaReachedAction == nil {
		return fmt.Errorf("required field quota_reached_action missing")
	}
	if all.SessionLimit == nil {
		return fmt.Errorf("required field session_limit missing")
	}
	if all.WindowType == nil {
		return fmt.Errorf("required field window_type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"daily_reset_time", "daily_reset_timezone", "quota_reached_action", "session_limit", "window_type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DailyResetTime = *all.DailyResetTime
	o.DailyResetTimezone = *all.DailyResetTimezone
	if !all.QuotaReachedAction.IsValid() {
		hasInvalidField = true
	} else {
		o.QuotaReachedAction = *all.QuotaReachedAction
	}
	o.SessionLimit = *all.SessionLimit
	if !all.WindowType.IsValid() {
		hasInvalidField = true
	} else {
		o.WindowType = *all.WindowType
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
