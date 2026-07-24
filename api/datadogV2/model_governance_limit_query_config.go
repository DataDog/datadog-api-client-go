// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GovernanceLimitQueryConfig The query execution context used to visualize a limit and its usage.
type GovernanceLimitQueryConfig struct {
	// The chart type used to visualize the limit and its usage.
	ChartType *string `json:"chart_type,omitempty"`
	// The time shift applied to compare current usage against a prior period.
	ComparisonShift *string `json:"comparison_shift,omitempty"`
	// The default value used for the limit when no explicit value is configured.
	DefaultValue *int64 `json:"default_value,omitempty"`
	// The direction in which usage approaches the limit, indicating whether higher or lower values are closer to the limit.
	Directionality *string `json:"directionality,omitempty"`
	// The number of days of data the limit query evaluates over.
	EffectiveTimeWindowDays *int64 `json:"effective_time_window_days,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGovernanceLimitQueryConfig instantiates a new GovernanceLimitQueryConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGovernanceLimitQueryConfig() *GovernanceLimitQueryConfig {
	this := GovernanceLimitQueryConfig{}
	return &this
}

// NewGovernanceLimitQueryConfigWithDefaults instantiates a new GovernanceLimitQueryConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGovernanceLimitQueryConfigWithDefaults() *GovernanceLimitQueryConfig {
	this := GovernanceLimitQueryConfig{}
	return &this
}

// GetChartType returns the ChartType field value if set, zero value otherwise.
func (o *GovernanceLimitQueryConfig) GetChartType() string {
	if o == nil || o.ChartType == nil {
		var ret string
		return ret
	}
	return *o.ChartType
}

// GetChartTypeOk returns a tuple with the ChartType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceLimitQueryConfig) GetChartTypeOk() (*string, bool) {
	if o == nil || o.ChartType == nil {
		return nil, false
	}
	return o.ChartType, true
}

// HasChartType returns a boolean if a field has been set.
func (o *GovernanceLimitQueryConfig) HasChartType() bool {
	return o != nil && o.ChartType != nil
}

// SetChartType gets a reference to the given string and assigns it to the ChartType field.
func (o *GovernanceLimitQueryConfig) SetChartType(v string) {
	o.ChartType = &v
}

// GetComparisonShift returns the ComparisonShift field value if set, zero value otherwise.
func (o *GovernanceLimitQueryConfig) GetComparisonShift() string {
	if o == nil || o.ComparisonShift == nil {
		var ret string
		return ret
	}
	return *o.ComparisonShift
}

// GetComparisonShiftOk returns a tuple with the ComparisonShift field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceLimitQueryConfig) GetComparisonShiftOk() (*string, bool) {
	if o == nil || o.ComparisonShift == nil {
		return nil, false
	}
	return o.ComparisonShift, true
}

// HasComparisonShift returns a boolean if a field has been set.
func (o *GovernanceLimitQueryConfig) HasComparisonShift() bool {
	return o != nil && o.ComparisonShift != nil
}

// SetComparisonShift gets a reference to the given string and assigns it to the ComparisonShift field.
func (o *GovernanceLimitQueryConfig) SetComparisonShift(v string) {
	o.ComparisonShift = &v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *GovernanceLimitQueryConfig) GetDefaultValue() int64 {
	if o == nil || o.DefaultValue == nil {
		var ret int64
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceLimitQueryConfig) GetDefaultValueOk() (*int64, bool) {
	if o == nil || o.DefaultValue == nil {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *GovernanceLimitQueryConfig) HasDefaultValue() bool {
	return o != nil && o.DefaultValue != nil
}

// SetDefaultValue gets a reference to the given int64 and assigns it to the DefaultValue field.
func (o *GovernanceLimitQueryConfig) SetDefaultValue(v int64) {
	o.DefaultValue = &v
}

// GetDirectionality returns the Directionality field value if set, zero value otherwise.
func (o *GovernanceLimitQueryConfig) GetDirectionality() string {
	if o == nil || o.Directionality == nil {
		var ret string
		return ret
	}
	return *o.Directionality
}

// GetDirectionalityOk returns a tuple with the Directionality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceLimitQueryConfig) GetDirectionalityOk() (*string, bool) {
	if o == nil || o.Directionality == nil {
		return nil, false
	}
	return o.Directionality, true
}

// HasDirectionality returns a boolean if a field has been set.
func (o *GovernanceLimitQueryConfig) HasDirectionality() bool {
	return o != nil && o.Directionality != nil
}

// SetDirectionality gets a reference to the given string and assigns it to the Directionality field.
func (o *GovernanceLimitQueryConfig) SetDirectionality(v string) {
	o.Directionality = &v
}

// GetEffectiveTimeWindowDays returns the EffectiveTimeWindowDays field value if set, zero value otherwise.
func (o *GovernanceLimitQueryConfig) GetEffectiveTimeWindowDays() int64 {
	if o == nil || o.EffectiveTimeWindowDays == nil {
		var ret int64
		return ret
	}
	return *o.EffectiveTimeWindowDays
}

// GetEffectiveTimeWindowDaysOk returns a tuple with the EffectiveTimeWindowDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceLimitQueryConfig) GetEffectiveTimeWindowDaysOk() (*int64, bool) {
	if o == nil || o.EffectiveTimeWindowDays == nil {
		return nil, false
	}
	return o.EffectiveTimeWindowDays, true
}

// HasEffectiveTimeWindowDays returns a boolean if a field has been set.
func (o *GovernanceLimitQueryConfig) HasEffectiveTimeWindowDays() bool {
	return o != nil && o.EffectiveTimeWindowDays != nil
}

// SetEffectiveTimeWindowDays gets a reference to the given int64 and assigns it to the EffectiveTimeWindowDays field.
func (o *GovernanceLimitQueryConfig) SetEffectiveTimeWindowDays(v int64) {
	o.EffectiveTimeWindowDays = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GovernanceLimitQueryConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ChartType != nil {
		toSerialize["chart_type"] = o.ChartType
	}
	if o.ComparisonShift != nil {
		toSerialize["comparison_shift"] = o.ComparisonShift
	}
	if o.DefaultValue != nil {
		toSerialize["default_value"] = o.DefaultValue
	}
	if o.Directionality != nil {
		toSerialize["directionality"] = o.Directionality
	}
	if o.EffectiveTimeWindowDays != nil {
		toSerialize["effective_time_window_days"] = o.EffectiveTimeWindowDays
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GovernanceLimitQueryConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ChartType               *string `json:"chart_type,omitempty"`
		ComparisonShift         *string `json:"comparison_shift,omitempty"`
		DefaultValue            *int64  `json:"default_value,omitempty"`
		Directionality          *string `json:"directionality,omitempty"`
		EffectiveTimeWindowDays *int64  `json:"effective_time_window_days,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"chart_type", "comparison_shift", "default_value", "directionality", "effective_time_window_days"})
	} else {
		return err
	}
	o.ChartType = all.ChartType
	o.ComparisonShift = all.ComparisonShift
	o.DefaultValue = all.DefaultValue
	o.Directionality = all.Directionality
	o.EffectiveTimeWindowDays = all.EffectiveTimeWindowDays

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
