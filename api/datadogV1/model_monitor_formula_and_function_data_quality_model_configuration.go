// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorFormulaAndFunctionDataQualityModelConfiguration Tuning options for the anomaly detection model used by the monitor.
type MonitorFormulaAndFunctionDataQualityModelConfiguration struct {
	// Number of days after which an open alert is automatically resolved.
	// When unset, alerts stay open until the measure returns within bounds.
	AutoResolveDays *int32 `json:"auto_resolve_days,omitempty"`
	// Whether to alert when the measure stops changing entirely.
	// Defaults to `true`.
	EnableFlatlineDetection *bool `json:"enable_flatline_detection,omitempty"`
	// Function applied to the measure before it is compared against the predicted bounds.
	Function *MonitorFormulaAndFunctionDataQualityDiffFunction `json:"function,omitempty"`
	// Minimum distance between the predicted value and the lower bound. Widening the
	// lower bound to at least this size suppresses alerts on small downward deviations.
	// When unset, no minimum is enforced.
	MinLowerBoundSize *float64 `json:"min_lower_bound_size,omitempty"`
	// Minimum distance between the predicted value and the upper bound. Widening the
	// upper bound to at least this size suppresses alerts on small upward deviations.
	// When unset, no minimum is enforced.
	MinUpperBoundSize *float64 `json:"min_upper_bound_size,omitempty"`
	// Restricts which predicted bound the monitor alerts on. `UPPER_ONLY` alerts only when
	// the measure rises above the upper bound, `LOWER_ONLY` only when it falls below the
	// lower bound. When unset, the monitor alerts on both.
	ModelBoundsOverride *MonitorFormulaAndFunctionDataQualityModelBoundsOverride `json:"model_bounds_override,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMonitorFormulaAndFunctionDataQualityModelConfiguration instantiates a new MonitorFormulaAndFunctionDataQualityModelConfiguration object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorFormulaAndFunctionDataQualityModelConfiguration() *MonitorFormulaAndFunctionDataQualityModelConfiguration {
	this := MonitorFormulaAndFunctionDataQualityModelConfiguration{}
	return &this
}

// NewMonitorFormulaAndFunctionDataQualityModelConfigurationWithDefaults instantiates a new MonitorFormulaAndFunctionDataQualityModelConfiguration object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorFormulaAndFunctionDataQualityModelConfigurationWithDefaults() *MonitorFormulaAndFunctionDataQualityModelConfiguration {
	this := MonitorFormulaAndFunctionDataQualityModelConfiguration{}
	return &this
}

// GetAutoResolveDays returns the AutoResolveDays field value if set, zero value otherwise.
func (o *MonitorFormulaAndFunctionDataQualityModelConfiguration) GetAutoResolveDays() int32 {
	if o == nil || o.AutoResolveDays == nil {
		var ret int32
		return ret
	}
	return *o.AutoResolveDays
}

// GetAutoResolveDaysOk returns a tuple with the AutoResolveDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionDataQualityModelConfiguration) GetAutoResolveDaysOk() (*int32, bool) {
	if o == nil || o.AutoResolveDays == nil {
		return nil, false
	}
	return o.AutoResolveDays, true
}

// HasAutoResolveDays returns a boolean if a field has been set.
func (o *MonitorFormulaAndFunctionDataQualityModelConfiguration) HasAutoResolveDays() bool {
	return o != nil && o.AutoResolveDays != nil
}

// SetAutoResolveDays gets a reference to the given int32 and assigns it to the AutoResolveDays field.
func (o *MonitorFormulaAndFunctionDataQualityModelConfiguration) SetAutoResolveDays(v int32) {
	o.AutoResolveDays = &v
}

// GetEnableFlatlineDetection returns the EnableFlatlineDetection field value if set, zero value otherwise.
func (o *MonitorFormulaAndFunctionDataQualityModelConfiguration) GetEnableFlatlineDetection() bool {
	if o == nil || o.EnableFlatlineDetection == nil {
		var ret bool
		return ret
	}
	return *o.EnableFlatlineDetection
}

// GetEnableFlatlineDetectionOk returns a tuple with the EnableFlatlineDetection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionDataQualityModelConfiguration) GetEnableFlatlineDetectionOk() (*bool, bool) {
	if o == nil || o.EnableFlatlineDetection == nil {
		return nil, false
	}
	return o.EnableFlatlineDetection, true
}

// HasEnableFlatlineDetection returns a boolean if a field has been set.
func (o *MonitorFormulaAndFunctionDataQualityModelConfiguration) HasEnableFlatlineDetection() bool {
	return o != nil && o.EnableFlatlineDetection != nil
}

// SetEnableFlatlineDetection gets a reference to the given bool and assigns it to the EnableFlatlineDetection field.
func (o *MonitorFormulaAndFunctionDataQualityModelConfiguration) SetEnableFlatlineDetection(v bool) {
	o.EnableFlatlineDetection = &v
}

// GetFunction returns the Function field value if set, zero value otherwise.
func (o *MonitorFormulaAndFunctionDataQualityModelConfiguration) GetFunction() MonitorFormulaAndFunctionDataQualityDiffFunction {
	if o == nil || o.Function == nil {
		var ret MonitorFormulaAndFunctionDataQualityDiffFunction
		return ret
	}
	return *o.Function
}

// GetFunctionOk returns a tuple with the Function field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionDataQualityModelConfiguration) GetFunctionOk() (*MonitorFormulaAndFunctionDataQualityDiffFunction, bool) {
	if o == nil || o.Function == nil {
		return nil, false
	}
	return o.Function, true
}

// HasFunction returns a boolean if a field has been set.
func (o *MonitorFormulaAndFunctionDataQualityModelConfiguration) HasFunction() bool {
	return o != nil && o.Function != nil
}

// SetFunction gets a reference to the given MonitorFormulaAndFunctionDataQualityDiffFunction and assigns it to the Function field.
func (o *MonitorFormulaAndFunctionDataQualityModelConfiguration) SetFunction(v MonitorFormulaAndFunctionDataQualityDiffFunction) {
	o.Function = &v
}

// GetMinLowerBoundSize returns the MinLowerBoundSize field value if set, zero value otherwise.
func (o *MonitorFormulaAndFunctionDataQualityModelConfiguration) GetMinLowerBoundSize() float64 {
	if o == nil || o.MinLowerBoundSize == nil {
		var ret float64
		return ret
	}
	return *o.MinLowerBoundSize
}

// GetMinLowerBoundSizeOk returns a tuple with the MinLowerBoundSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionDataQualityModelConfiguration) GetMinLowerBoundSizeOk() (*float64, bool) {
	if o == nil || o.MinLowerBoundSize == nil {
		return nil, false
	}
	return o.MinLowerBoundSize, true
}

// HasMinLowerBoundSize returns a boolean if a field has been set.
func (o *MonitorFormulaAndFunctionDataQualityModelConfiguration) HasMinLowerBoundSize() bool {
	return o != nil && o.MinLowerBoundSize != nil
}

// SetMinLowerBoundSize gets a reference to the given float64 and assigns it to the MinLowerBoundSize field.
func (o *MonitorFormulaAndFunctionDataQualityModelConfiguration) SetMinLowerBoundSize(v float64) {
	o.MinLowerBoundSize = &v
}

// GetMinUpperBoundSize returns the MinUpperBoundSize field value if set, zero value otherwise.
func (o *MonitorFormulaAndFunctionDataQualityModelConfiguration) GetMinUpperBoundSize() float64 {
	if o == nil || o.MinUpperBoundSize == nil {
		var ret float64
		return ret
	}
	return *o.MinUpperBoundSize
}

// GetMinUpperBoundSizeOk returns a tuple with the MinUpperBoundSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionDataQualityModelConfiguration) GetMinUpperBoundSizeOk() (*float64, bool) {
	if o == nil || o.MinUpperBoundSize == nil {
		return nil, false
	}
	return o.MinUpperBoundSize, true
}

// HasMinUpperBoundSize returns a boolean if a field has been set.
func (o *MonitorFormulaAndFunctionDataQualityModelConfiguration) HasMinUpperBoundSize() bool {
	return o != nil && o.MinUpperBoundSize != nil
}

// SetMinUpperBoundSize gets a reference to the given float64 and assigns it to the MinUpperBoundSize field.
func (o *MonitorFormulaAndFunctionDataQualityModelConfiguration) SetMinUpperBoundSize(v float64) {
	o.MinUpperBoundSize = &v
}

// GetModelBoundsOverride returns the ModelBoundsOverride field value if set, zero value otherwise.
func (o *MonitorFormulaAndFunctionDataQualityModelConfiguration) GetModelBoundsOverride() MonitorFormulaAndFunctionDataQualityModelBoundsOverride {
	if o == nil || o.ModelBoundsOverride == nil {
		var ret MonitorFormulaAndFunctionDataQualityModelBoundsOverride
		return ret
	}
	return *o.ModelBoundsOverride
}

// GetModelBoundsOverrideOk returns a tuple with the ModelBoundsOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionDataQualityModelConfiguration) GetModelBoundsOverrideOk() (*MonitorFormulaAndFunctionDataQualityModelBoundsOverride, bool) {
	if o == nil || o.ModelBoundsOverride == nil {
		return nil, false
	}
	return o.ModelBoundsOverride, true
}

// HasModelBoundsOverride returns a boolean if a field has been set.
func (o *MonitorFormulaAndFunctionDataQualityModelConfiguration) HasModelBoundsOverride() bool {
	return o != nil && o.ModelBoundsOverride != nil
}

// SetModelBoundsOverride gets a reference to the given MonitorFormulaAndFunctionDataQualityModelBoundsOverride and assigns it to the ModelBoundsOverride field.
func (o *MonitorFormulaAndFunctionDataQualityModelConfiguration) SetModelBoundsOverride(v MonitorFormulaAndFunctionDataQualityModelBoundsOverride) {
	o.ModelBoundsOverride = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorFormulaAndFunctionDataQualityModelConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AutoResolveDays != nil {
		toSerialize["auto_resolve_days"] = o.AutoResolveDays
	}
	if o.EnableFlatlineDetection != nil {
		toSerialize["enable_flatline_detection"] = o.EnableFlatlineDetection
	}
	if o.Function != nil {
		toSerialize["function"] = o.Function
	}
	if o.MinLowerBoundSize != nil {
		toSerialize["min_lower_bound_size"] = o.MinLowerBoundSize
	}
	if o.MinUpperBoundSize != nil {
		toSerialize["min_upper_bound_size"] = o.MinUpperBoundSize
	}
	if o.ModelBoundsOverride != nil {
		toSerialize["model_bounds_override"] = o.ModelBoundsOverride
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorFormulaAndFunctionDataQualityModelConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AutoResolveDays         *int32                                                   `json:"auto_resolve_days,omitempty"`
		EnableFlatlineDetection *bool                                                    `json:"enable_flatline_detection,omitempty"`
		Function                *MonitorFormulaAndFunctionDataQualityDiffFunction        `json:"function,omitempty"`
		MinLowerBoundSize       *float64                                                 `json:"min_lower_bound_size,omitempty"`
		MinUpperBoundSize       *float64                                                 `json:"min_upper_bound_size,omitempty"`
		ModelBoundsOverride     *MonitorFormulaAndFunctionDataQualityModelBoundsOverride `json:"model_bounds_override,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auto_resolve_days", "enable_flatline_detection", "function", "min_lower_bound_size", "min_upper_bound_size", "model_bounds_override"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AutoResolveDays = all.AutoResolveDays
	o.EnableFlatlineDetection = all.EnableFlatlineDetection
	if all.Function != nil && !all.Function.IsValid() {
		hasInvalidField = true
	} else {
		o.Function = all.Function
	}
	o.MinLowerBoundSize = all.MinLowerBoundSize
	o.MinUpperBoundSize = all.MinUpperBoundSize
	if all.ModelBoundsOverride != nil && !all.ModelBoundsOverride.IsValid() {
		hasInvalidField = true
	} else {
		o.ModelBoundsOverride = all.ModelBoundsOverride
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
