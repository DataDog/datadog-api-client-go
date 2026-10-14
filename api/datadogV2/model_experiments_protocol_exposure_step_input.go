// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsProtocolExposureStepInput A traffic exposure step in the protocol template.
type ExperimentsProtocolExposureStepInput struct {
	// Fraction of traffic exposed during this rollout step.
	ExposureRatio *float64 `json:"exposure_ratio,omitempty"`
	// Index of the group that contains this rollout step.
	GroupedStepIndex datadog.NullableInt64 `json:"grouped_step_index,omitempty"`
	// Duration of this rollout step, in milliseconds.
	IntervalMs datadog.NullableInt64 `json:"interval_ms,omitempty"`
	// Whether this schedule entry represents a pause.
	IsPauseRecord *bool `json:"is_pause_record,omitempty"`
	// Position of this entry in the ordered configuration.
	OrderPosition *int64 `json:"order_position,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsProtocolExposureStepInput instantiates a new ExperimentsProtocolExposureStepInput object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsProtocolExposureStepInput() *ExperimentsProtocolExposureStepInput {
	this := ExperimentsProtocolExposureStepInput{}
	return &this
}

// NewExperimentsProtocolExposureStepInputWithDefaults instantiates a new ExperimentsProtocolExposureStepInput object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsProtocolExposureStepInputWithDefaults() *ExperimentsProtocolExposureStepInput {
	this := ExperimentsProtocolExposureStepInput{}
	return &this
}

// GetExposureRatio returns the ExposureRatio field value if set, zero value otherwise.
func (o *ExperimentsProtocolExposureStepInput) GetExposureRatio() float64 {
	if o == nil || o.ExposureRatio == nil {
		var ret float64
		return ret
	}
	return *o.ExposureRatio
}

// GetExposureRatioOk returns a tuple with the ExposureRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsProtocolExposureStepInput) GetExposureRatioOk() (*float64, bool) {
	if o == nil || o.ExposureRatio == nil {
		return nil, false
	}
	return o.ExposureRatio, true
}

// HasExposureRatio returns a boolean if a field has been set.
func (o *ExperimentsProtocolExposureStepInput) HasExposureRatio() bool {
	return o != nil && o.ExposureRatio != nil
}

// SetExposureRatio gets a reference to the given float64 and assigns it to the ExposureRatio field.
func (o *ExperimentsProtocolExposureStepInput) SetExposureRatio(v float64) {
	o.ExposureRatio = &v
}

// GetGroupedStepIndex returns the GroupedStepIndex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsProtocolExposureStepInput) GetGroupedStepIndex() int64 {
	if o == nil || o.GroupedStepIndex.Get() == nil {
		var ret int64
		return ret
	}
	return *o.GroupedStepIndex.Get()
}

// GetGroupedStepIndexOk returns a tuple with the GroupedStepIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsProtocolExposureStepInput) GetGroupedStepIndexOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupedStepIndex.Get(), o.GroupedStepIndex.IsSet()
}

// HasGroupedStepIndex returns a boolean if a field has been set.
func (o *ExperimentsProtocolExposureStepInput) HasGroupedStepIndex() bool {
	return o != nil && o.GroupedStepIndex.IsSet()
}

// SetGroupedStepIndex gets a reference to the given datadog.NullableInt64 and assigns it to the GroupedStepIndex field.
func (o *ExperimentsProtocolExposureStepInput) SetGroupedStepIndex(v int64) {
	o.GroupedStepIndex.Set(&v)
}

// SetGroupedStepIndexNil sets the value for GroupedStepIndex to be an explicit nil.
func (o *ExperimentsProtocolExposureStepInput) SetGroupedStepIndexNil() {
	o.GroupedStepIndex.Set(nil)
}

// UnsetGroupedStepIndex ensures that no value is present for GroupedStepIndex, not even an explicit nil.
func (o *ExperimentsProtocolExposureStepInput) UnsetGroupedStepIndex() {
	o.GroupedStepIndex.Unset()
}

// GetIntervalMs returns the IntervalMs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsProtocolExposureStepInput) GetIntervalMs() int64 {
	if o == nil || o.IntervalMs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.IntervalMs.Get()
}

// GetIntervalMsOk returns a tuple with the IntervalMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsProtocolExposureStepInput) GetIntervalMsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.IntervalMs.Get(), o.IntervalMs.IsSet()
}

// HasIntervalMs returns a boolean if a field has been set.
func (o *ExperimentsProtocolExposureStepInput) HasIntervalMs() bool {
	return o != nil && o.IntervalMs.IsSet()
}

// SetIntervalMs gets a reference to the given datadog.NullableInt64 and assigns it to the IntervalMs field.
func (o *ExperimentsProtocolExposureStepInput) SetIntervalMs(v int64) {
	o.IntervalMs.Set(&v)
}

// SetIntervalMsNil sets the value for IntervalMs to be an explicit nil.
func (o *ExperimentsProtocolExposureStepInput) SetIntervalMsNil() {
	o.IntervalMs.Set(nil)
}

// UnsetIntervalMs ensures that no value is present for IntervalMs, not even an explicit nil.
func (o *ExperimentsProtocolExposureStepInput) UnsetIntervalMs() {
	o.IntervalMs.Unset()
}

// GetIsPauseRecord returns the IsPauseRecord field value if set, zero value otherwise.
func (o *ExperimentsProtocolExposureStepInput) GetIsPauseRecord() bool {
	if o == nil || o.IsPauseRecord == nil {
		var ret bool
		return ret
	}
	return *o.IsPauseRecord
}

// GetIsPauseRecordOk returns a tuple with the IsPauseRecord field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsProtocolExposureStepInput) GetIsPauseRecordOk() (*bool, bool) {
	if o == nil || o.IsPauseRecord == nil {
		return nil, false
	}
	return o.IsPauseRecord, true
}

// HasIsPauseRecord returns a boolean if a field has been set.
func (o *ExperimentsProtocolExposureStepInput) HasIsPauseRecord() bool {
	return o != nil && o.IsPauseRecord != nil
}

// SetIsPauseRecord gets a reference to the given bool and assigns it to the IsPauseRecord field.
func (o *ExperimentsProtocolExposureStepInput) SetIsPauseRecord(v bool) {
	o.IsPauseRecord = &v
}

// GetOrderPosition returns the OrderPosition field value if set, zero value otherwise.
func (o *ExperimentsProtocolExposureStepInput) GetOrderPosition() int64 {
	if o == nil || o.OrderPosition == nil {
		var ret int64
		return ret
	}
	return *o.OrderPosition
}

// GetOrderPositionOk returns a tuple with the OrderPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsProtocolExposureStepInput) GetOrderPositionOk() (*int64, bool) {
	if o == nil || o.OrderPosition == nil {
		return nil, false
	}
	return o.OrderPosition, true
}

// HasOrderPosition returns a boolean if a field has been set.
func (o *ExperimentsProtocolExposureStepInput) HasOrderPosition() bool {
	return o != nil && o.OrderPosition != nil
}

// SetOrderPosition gets a reference to the given int64 and assigns it to the OrderPosition field.
func (o *ExperimentsProtocolExposureStepInput) SetOrderPosition(v int64) {
	o.OrderPosition = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsProtocolExposureStepInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ExposureRatio != nil {
		toSerialize["exposure_ratio"] = o.ExposureRatio
	}
	if o.GroupedStepIndex.IsSet() {
		toSerialize["grouped_step_index"] = o.GroupedStepIndex.Get()
	}
	if o.IntervalMs.IsSet() {
		toSerialize["interval_ms"] = o.IntervalMs.Get()
	}
	if o.IsPauseRecord != nil {
		toSerialize["is_pause_record"] = o.IsPauseRecord
	}
	if o.OrderPosition != nil {
		toSerialize["order_position"] = o.OrderPosition
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsProtocolExposureStepInput) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ExposureRatio    *float64              `json:"exposure_ratio,omitempty"`
		GroupedStepIndex datadog.NullableInt64 `json:"grouped_step_index,omitempty"`
		IntervalMs       datadog.NullableInt64 `json:"interval_ms,omitempty"`
		IsPauseRecord    *bool                 `json:"is_pause_record,omitempty"`
		OrderPosition    *int64                `json:"order_position,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"exposure_ratio", "grouped_step_index", "interval_ms", "is_pause_record", "order_position"})
	} else {
		return err
	}
	o.ExposureRatio = all.ExposureRatio
	o.GroupedStepIndex = all.GroupedStepIndex
	o.IntervalMs = all.IntervalMs
	o.IsPauseRecord = all.IsPauseRecord
	o.OrderPosition = all.OrderPosition

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
