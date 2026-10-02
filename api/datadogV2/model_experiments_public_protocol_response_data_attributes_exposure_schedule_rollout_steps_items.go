// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsPublicProtocolResponseDataAttributesExposureScheduleRolloutStepsItems One step in the protocol's traffic exposure schedule.
type ExperimentsPublicProtocolResponseDataAttributesExposureScheduleRolloutStepsItems struct {
	// Fraction of traffic exposed during this rollout step.
	ExposureRatio *float64 `json:"exposure_ratio,omitempty"`
	// Index of the group that contains this rollout step.
	GroupedStepIndex *int64 `json:"grouped_step_index,omitempty"`
	// Duration of this rollout step, in milliseconds.
	IntervalMs *int64 `json:"interval_ms,omitempty"`
	// Whether this schedule entry represents a pause.
	IsPauseRecord *bool `json:"is_pause_record,omitempty"`
	// Position of this entry in the ordered configuration.
	OrderPosition *int64 `json:"order_position,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsPublicProtocolResponseDataAttributesExposureScheduleRolloutStepsItems instantiates a new ExperimentsPublicProtocolResponseDataAttributesExposureScheduleRolloutStepsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsPublicProtocolResponseDataAttributesExposureScheduleRolloutStepsItems() *ExperimentsPublicProtocolResponseDataAttributesExposureScheduleRolloutStepsItems {
	this := ExperimentsPublicProtocolResponseDataAttributesExposureScheduleRolloutStepsItems{}
	return &this
}

// NewExperimentsPublicProtocolResponseDataAttributesExposureScheduleRolloutStepsItemsWithDefaults instantiates a new ExperimentsPublicProtocolResponseDataAttributesExposureScheduleRolloutStepsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsPublicProtocolResponseDataAttributesExposureScheduleRolloutStepsItemsWithDefaults() *ExperimentsPublicProtocolResponseDataAttributesExposureScheduleRolloutStepsItems {
	this := ExperimentsPublicProtocolResponseDataAttributesExposureScheduleRolloutStepsItems{}
	return &this
}

// GetExposureRatio returns the ExposureRatio field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributesExposureScheduleRolloutStepsItems) GetExposureRatio() float64 {
	if o == nil || o.ExposureRatio == nil {
		var ret float64
		return ret
	}
	return *o.ExposureRatio
}

// GetExposureRatioOk returns a tuple with the ExposureRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesExposureScheduleRolloutStepsItems) GetExposureRatioOk() (*float64, bool) {
	if o == nil || o.ExposureRatio == nil {
		return nil, false
	}
	return o.ExposureRatio, true
}

// HasExposureRatio returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesExposureScheduleRolloutStepsItems) HasExposureRatio() bool {
	return o != nil && o.ExposureRatio != nil
}

// SetExposureRatio gets a reference to the given float64 and assigns it to the ExposureRatio field.
func (o *ExperimentsPublicProtocolResponseDataAttributesExposureScheduleRolloutStepsItems) SetExposureRatio(v float64) {
	o.ExposureRatio = &v
}

// GetGroupedStepIndex returns the GroupedStepIndex field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributesExposureScheduleRolloutStepsItems) GetGroupedStepIndex() int64 {
	if o == nil || o.GroupedStepIndex == nil {
		var ret int64
		return ret
	}
	return *o.GroupedStepIndex
}

// GetGroupedStepIndexOk returns a tuple with the GroupedStepIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesExposureScheduleRolloutStepsItems) GetGroupedStepIndexOk() (*int64, bool) {
	if o == nil || o.GroupedStepIndex == nil {
		return nil, false
	}
	return o.GroupedStepIndex, true
}

// HasGroupedStepIndex returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesExposureScheduleRolloutStepsItems) HasGroupedStepIndex() bool {
	return o != nil && o.GroupedStepIndex != nil
}

// SetGroupedStepIndex gets a reference to the given int64 and assigns it to the GroupedStepIndex field.
func (o *ExperimentsPublicProtocolResponseDataAttributesExposureScheduleRolloutStepsItems) SetGroupedStepIndex(v int64) {
	o.GroupedStepIndex = &v
}

// GetIntervalMs returns the IntervalMs field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributesExposureScheduleRolloutStepsItems) GetIntervalMs() int64 {
	if o == nil || o.IntervalMs == nil {
		var ret int64
		return ret
	}
	return *o.IntervalMs
}

// GetIntervalMsOk returns a tuple with the IntervalMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesExposureScheduleRolloutStepsItems) GetIntervalMsOk() (*int64, bool) {
	if o == nil || o.IntervalMs == nil {
		return nil, false
	}
	return o.IntervalMs, true
}

// HasIntervalMs returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesExposureScheduleRolloutStepsItems) HasIntervalMs() bool {
	return o != nil && o.IntervalMs != nil
}

// SetIntervalMs gets a reference to the given int64 and assigns it to the IntervalMs field.
func (o *ExperimentsPublicProtocolResponseDataAttributesExposureScheduleRolloutStepsItems) SetIntervalMs(v int64) {
	o.IntervalMs = &v
}

// GetIsPauseRecord returns the IsPauseRecord field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributesExposureScheduleRolloutStepsItems) GetIsPauseRecord() bool {
	if o == nil || o.IsPauseRecord == nil {
		var ret bool
		return ret
	}
	return *o.IsPauseRecord
}

// GetIsPauseRecordOk returns a tuple with the IsPauseRecord field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesExposureScheduleRolloutStepsItems) GetIsPauseRecordOk() (*bool, bool) {
	if o == nil || o.IsPauseRecord == nil {
		return nil, false
	}
	return o.IsPauseRecord, true
}

// HasIsPauseRecord returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesExposureScheduleRolloutStepsItems) HasIsPauseRecord() bool {
	return o != nil && o.IsPauseRecord != nil
}

// SetIsPauseRecord gets a reference to the given bool and assigns it to the IsPauseRecord field.
func (o *ExperimentsPublicProtocolResponseDataAttributesExposureScheduleRolloutStepsItems) SetIsPauseRecord(v bool) {
	o.IsPauseRecord = &v
}

// GetOrderPosition returns the OrderPosition field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributesExposureScheduleRolloutStepsItems) GetOrderPosition() int64 {
	if o == nil || o.OrderPosition == nil {
		var ret int64
		return ret
	}
	return *o.OrderPosition
}

// GetOrderPositionOk returns a tuple with the OrderPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesExposureScheduleRolloutStepsItems) GetOrderPositionOk() (*int64, bool) {
	if o == nil || o.OrderPosition == nil {
		return nil, false
	}
	return o.OrderPosition, true
}

// HasOrderPosition returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesExposureScheduleRolloutStepsItems) HasOrderPosition() bool {
	return o != nil && o.OrderPosition != nil
}

// SetOrderPosition gets a reference to the given int64 and assigns it to the OrderPosition field.
func (o *ExperimentsPublicProtocolResponseDataAttributesExposureScheduleRolloutStepsItems) SetOrderPosition(v int64) {
	o.OrderPosition = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsPublicProtocolResponseDataAttributesExposureScheduleRolloutStepsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ExposureRatio != nil {
		toSerialize["exposure_ratio"] = o.ExposureRatio
	}
	if o.GroupedStepIndex != nil {
		toSerialize["grouped_step_index"] = o.GroupedStepIndex
	}
	if o.IntervalMs != nil {
		toSerialize["interval_ms"] = o.IntervalMs
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
func (o *ExperimentsPublicProtocolResponseDataAttributesExposureScheduleRolloutStepsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ExposureRatio    *float64 `json:"exposure_ratio,omitempty"`
		GroupedStepIndex *int64   `json:"grouped_step_index,omitempty"`
		IntervalMs       *int64   `json:"interval_ms,omitempty"`
		IsPauseRecord    *bool    `json:"is_pause_record,omitempty"`
		OrderPosition    *int64   `json:"order_position,omitempty"`
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
