/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// SecurityMonitoringRuleOptions Options on rules.
type SecurityMonitoringRuleOptions struct {
	// The detection method.
	DetectionMethod *SecurityMonitoringRuleDetectionMethod `json:"detectionMethod,omitempty"`
	// A time window is specified to match when at least one of the cases matches true. This is a sliding window
	// and evaluates in real time.
	EvaluationWindow *SecurityMonitoringRuleEvaluationWindow `json:"evaluationWindow,omitempty"`
	// Hardcoded evaluator type.
	HardcodedEvaluatorType *SecurityMonitoringRuleHardcodedEvaluatorType `json:"hardcodedEvaluatorType,omitempty"`
	// Options on impossible travel rules.
	ImpossibleTravelOptions *SecurityMonitoringRuleImpossibleTravelOptions `json:"impossibleTravelOptions,omitempty"`
	// Once a signal is generated, the signal will remain “open” if a case is matched at least once within
	// this keep alive window.
	KeepAlive *SecurityMonitoringRuleKeepAlive `json:"keepAlive,omitempty"`
	// A signal will “close” regardless of the query being matched once the time exceeds the maximum duration.
	// This time is calculated from the first seen timestamp.
	MaxSignalDuration *SecurityMonitoringRuleMaxSignalDuration `json:"maxSignalDuration,omitempty"`
	// Options on new value rules.
	NewValueOptions *SecurityMonitoringRuleNewValueOptions `json:"newValueOptions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewSecurityMonitoringRuleOptions instantiates a new SecurityMonitoringRuleOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityMonitoringRuleOptions() *SecurityMonitoringRuleOptions {
	this := SecurityMonitoringRuleOptions{}
	return &this
}

// NewSecurityMonitoringRuleOptionsWithDefaults instantiates a new SecurityMonitoringRuleOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityMonitoringRuleOptionsWithDefaults() *SecurityMonitoringRuleOptions {
	this := SecurityMonitoringRuleOptions{}
	return &this
}

// GetDetectionMethod returns the DetectionMethod field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleOptions) GetDetectionMethod() SecurityMonitoringRuleDetectionMethod {
	if o == nil || o.DetectionMethod == nil {
		var ret SecurityMonitoringRuleDetectionMethod
		return ret
	}
	return *o.DetectionMethod
}

// GetDetectionMethodOk returns a tuple with the DetectionMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleOptions) GetDetectionMethodOk() (*SecurityMonitoringRuleDetectionMethod, bool) {
	if o == nil || o.DetectionMethod == nil {
		return nil, false
	}
	return o.DetectionMethod, true
}

// HasDetectionMethod returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleOptions) HasDetectionMethod() bool {
	if o != nil && o.DetectionMethod != nil {
		return true
	}

	return false
}

// SetDetectionMethod gets a reference to the given SecurityMonitoringRuleDetectionMethod and assigns it to the DetectionMethod field.
func (o *SecurityMonitoringRuleOptions) SetDetectionMethod(v SecurityMonitoringRuleDetectionMethod) {
	o.DetectionMethod = &v
}

// GetEvaluationWindow returns the EvaluationWindow field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleOptions) GetEvaluationWindow() SecurityMonitoringRuleEvaluationWindow {
	if o == nil || o.EvaluationWindow == nil {
		var ret SecurityMonitoringRuleEvaluationWindow
		return ret
	}
	return *o.EvaluationWindow
}

// GetEvaluationWindowOk returns a tuple with the EvaluationWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleOptions) GetEvaluationWindowOk() (*SecurityMonitoringRuleEvaluationWindow, bool) {
	if o == nil || o.EvaluationWindow == nil {
		return nil, false
	}
	return o.EvaluationWindow, true
}

// HasEvaluationWindow returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleOptions) HasEvaluationWindow() bool {
	if o != nil && o.EvaluationWindow != nil {
		return true
	}

	return false
}

// SetEvaluationWindow gets a reference to the given SecurityMonitoringRuleEvaluationWindow and assigns it to the EvaluationWindow field.
func (o *SecurityMonitoringRuleOptions) SetEvaluationWindow(v SecurityMonitoringRuleEvaluationWindow) {
	o.EvaluationWindow = &v
}

// GetHardcodedEvaluatorType returns the HardcodedEvaluatorType field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleOptions) GetHardcodedEvaluatorType() SecurityMonitoringRuleHardcodedEvaluatorType {
	if o == nil || o.HardcodedEvaluatorType == nil {
		var ret SecurityMonitoringRuleHardcodedEvaluatorType
		return ret
	}
	return *o.HardcodedEvaluatorType
}

// GetHardcodedEvaluatorTypeOk returns a tuple with the HardcodedEvaluatorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleOptions) GetHardcodedEvaluatorTypeOk() (*SecurityMonitoringRuleHardcodedEvaluatorType, bool) {
	if o == nil || o.HardcodedEvaluatorType == nil {
		return nil, false
	}
	return o.HardcodedEvaluatorType, true
}

// HasHardcodedEvaluatorType returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleOptions) HasHardcodedEvaluatorType() bool {
	if o != nil && o.HardcodedEvaluatorType != nil {
		return true
	}

	return false
}

// SetHardcodedEvaluatorType gets a reference to the given SecurityMonitoringRuleHardcodedEvaluatorType and assigns it to the HardcodedEvaluatorType field.
func (o *SecurityMonitoringRuleOptions) SetHardcodedEvaluatorType(v SecurityMonitoringRuleHardcodedEvaluatorType) {
	o.HardcodedEvaluatorType = &v
}

// GetImpossibleTravelOptions returns the ImpossibleTravelOptions field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleOptions) GetImpossibleTravelOptions() SecurityMonitoringRuleImpossibleTravelOptions {
	if o == nil || o.ImpossibleTravelOptions == nil {
		var ret SecurityMonitoringRuleImpossibleTravelOptions
		return ret
	}
	return *o.ImpossibleTravelOptions
}

// GetImpossibleTravelOptionsOk returns a tuple with the ImpossibleTravelOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleOptions) GetImpossibleTravelOptionsOk() (*SecurityMonitoringRuleImpossibleTravelOptions, bool) {
	if o == nil || o.ImpossibleTravelOptions == nil {
		return nil, false
	}
	return o.ImpossibleTravelOptions, true
}

// HasImpossibleTravelOptions returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleOptions) HasImpossibleTravelOptions() bool {
	if o != nil && o.ImpossibleTravelOptions != nil {
		return true
	}

	return false
}

// SetImpossibleTravelOptions gets a reference to the given SecurityMonitoringRuleImpossibleTravelOptions and assigns it to the ImpossibleTravelOptions field.
func (o *SecurityMonitoringRuleOptions) SetImpossibleTravelOptions(v SecurityMonitoringRuleImpossibleTravelOptions) {
	o.ImpossibleTravelOptions = &v
}

// GetKeepAlive returns the KeepAlive field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleOptions) GetKeepAlive() SecurityMonitoringRuleKeepAlive {
	if o == nil || o.KeepAlive == nil {
		var ret SecurityMonitoringRuleKeepAlive
		return ret
	}
	return *o.KeepAlive
}

// GetKeepAliveOk returns a tuple with the KeepAlive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleOptions) GetKeepAliveOk() (*SecurityMonitoringRuleKeepAlive, bool) {
	if o == nil || o.KeepAlive == nil {
		return nil, false
	}
	return o.KeepAlive, true
}

// HasKeepAlive returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleOptions) HasKeepAlive() bool {
	if o != nil && o.KeepAlive != nil {
		return true
	}

	return false
}

// SetKeepAlive gets a reference to the given SecurityMonitoringRuleKeepAlive and assigns it to the KeepAlive field.
func (o *SecurityMonitoringRuleOptions) SetKeepAlive(v SecurityMonitoringRuleKeepAlive) {
	o.KeepAlive = &v
}

// GetMaxSignalDuration returns the MaxSignalDuration field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleOptions) GetMaxSignalDuration() SecurityMonitoringRuleMaxSignalDuration {
	if o == nil || o.MaxSignalDuration == nil {
		var ret SecurityMonitoringRuleMaxSignalDuration
		return ret
	}
	return *o.MaxSignalDuration
}

// GetMaxSignalDurationOk returns a tuple with the MaxSignalDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleOptions) GetMaxSignalDurationOk() (*SecurityMonitoringRuleMaxSignalDuration, bool) {
	if o == nil || o.MaxSignalDuration == nil {
		return nil, false
	}
	return o.MaxSignalDuration, true
}

// HasMaxSignalDuration returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleOptions) HasMaxSignalDuration() bool {
	if o != nil && o.MaxSignalDuration != nil {
		return true
	}

	return false
}

// SetMaxSignalDuration gets a reference to the given SecurityMonitoringRuleMaxSignalDuration and assigns it to the MaxSignalDuration field.
func (o *SecurityMonitoringRuleOptions) SetMaxSignalDuration(v SecurityMonitoringRuleMaxSignalDuration) {
	o.MaxSignalDuration = &v
}

// GetNewValueOptions returns the NewValueOptions field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleOptions) GetNewValueOptions() SecurityMonitoringRuleNewValueOptions {
	if o == nil || o.NewValueOptions == nil {
		var ret SecurityMonitoringRuleNewValueOptions
		return ret
	}
	return *o.NewValueOptions
}

// GetNewValueOptionsOk returns a tuple with the NewValueOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleOptions) GetNewValueOptionsOk() (*SecurityMonitoringRuleNewValueOptions, bool) {
	if o == nil || o.NewValueOptions == nil {
		return nil, false
	}
	return o.NewValueOptions, true
}

// HasNewValueOptions returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleOptions) HasNewValueOptions() bool {
	if o != nil && o.NewValueOptions != nil {
		return true
	}

	return false
}

// SetNewValueOptions gets a reference to the given SecurityMonitoringRuleNewValueOptions and assigns it to the NewValueOptions field.
func (o *SecurityMonitoringRuleOptions) SetNewValueOptions(v SecurityMonitoringRuleNewValueOptions) {
	o.NewValueOptions = &v
}

func (o SecurityMonitoringRuleOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.DetectionMethod != nil {
		toSerialize["detectionMethod"] = o.DetectionMethod
	}
	if o.EvaluationWindow != nil {
		toSerialize["evaluationWindow"] = o.EvaluationWindow
	}
	if o.HardcodedEvaluatorType != nil {
		toSerialize["hardcodedEvaluatorType"] = o.HardcodedEvaluatorType
	}
	if o.ImpossibleTravelOptions != nil {
		toSerialize["impossibleTravelOptions"] = o.ImpossibleTravelOptions
	}
	if o.KeepAlive != nil {
		toSerialize["keepAlive"] = o.KeepAlive
	}
	if o.MaxSignalDuration != nil {
		toSerialize["maxSignalDuration"] = o.MaxSignalDuration
	}
	if o.NewValueOptions != nil {
		toSerialize["newValueOptions"] = o.NewValueOptions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

func (o *SecurityMonitoringRuleOptions) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		DetectionMethod         *SecurityMonitoringRuleDetectionMethod         `json:"detectionMethod,omitempty"`
		EvaluationWindow        *SecurityMonitoringRuleEvaluationWindow        `json:"evaluationWindow,omitempty"`
		HardcodedEvaluatorType  *SecurityMonitoringRuleHardcodedEvaluatorType  `json:"hardcodedEvaluatorType,omitempty"`
		ImpossibleTravelOptions *SecurityMonitoringRuleImpossibleTravelOptions `json:"impossibleTravelOptions,omitempty"`
		KeepAlive               *SecurityMonitoringRuleKeepAlive               `json:"keepAlive,omitempty"`
		MaxSignalDuration       *SecurityMonitoringRuleMaxSignalDuration       `json:"maxSignalDuration,omitempty"`
		NewValueOptions         *SecurityMonitoringRuleNewValueOptions         `json:"newValueOptions,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if v := all.DetectionMethod; v != nil && !v.IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if v := all.EvaluationWindow; v != nil && !v.IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if v := all.HardcodedEvaluatorType; v != nil && !v.IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if v := all.KeepAlive; v != nil && !v.IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if v := all.MaxSignalDuration; v != nil && !v.IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.DetectionMethod = all.DetectionMethod
	o.EvaluationWindow = all.EvaluationWindow
	o.HardcodedEvaluatorType = all.HardcodedEvaluatorType
	if all.ImpossibleTravelOptions != nil && all.ImpossibleTravelOptions.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.ImpossibleTravelOptions = all.ImpossibleTravelOptions
	o.KeepAlive = all.KeepAlive
	o.MaxSignalDuration = all.MaxSignalDuration
	if all.NewValueOptions != nil && all.NewValueOptions.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.NewValueOptions = all.NewValueOptions
	return nil
}
