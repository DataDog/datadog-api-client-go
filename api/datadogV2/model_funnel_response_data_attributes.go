// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FunnelResponseDataAttributes
type FunnelResponseDataAttributes struct {
	//
	EndToEndConversionRate *float64 `json:"end_to_end_conversion_rate,omitempty"`
	//
	EndToEndElapsedTime *FunnelResponseElapsedTime `json:"end_to_end_elapsed_time,omitempty"`
	//
	FunnelSteps []FunnelResponseDataAttributesFunnelStepsItems `json:"funnel_steps,omitempty"`
	//
	InitialCount *int64 `json:"initial_count,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFunnelResponseDataAttributes instantiates a new FunnelResponseDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFunnelResponseDataAttributes() *FunnelResponseDataAttributes {
	this := FunnelResponseDataAttributes{}
	return &this
}

// NewFunnelResponseDataAttributesWithDefaults instantiates a new FunnelResponseDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFunnelResponseDataAttributesWithDefaults() *FunnelResponseDataAttributes {
	this := FunnelResponseDataAttributes{}
	return &this
}

// GetEndToEndConversionRate returns the EndToEndConversionRate field value if set, zero value otherwise.
func (o *FunnelResponseDataAttributes) GetEndToEndConversionRate() float64 {
	if o == nil || o.EndToEndConversionRate == nil {
		var ret float64
		return ret
	}
	return *o.EndToEndConversionRate
}

// GetEndToEndConversionRateOk returns a tuple with the EndToEndConversionRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelResponseDataAttributes) GetEndToEndConversionRateOk() (*float64, bool) {
	if o == nil || o.EndToEndConversionRate == nil {
		return nil, false
	}
	return o.EndToEndConversionRate, true
}

// HasEndToEndConversionRate returns a boolean if a field has been set.
func (o *FunnelResponseDataAttributes) HasEndToEndConversionRate() bool {
	return o != nil && o.EndToEndConversionRate != nil
}

// SetEndToEndConversionRate gets a reference to the given float64 and assigns it to the EndToEndConversionRate field.
func (o *FunnelResponseDataAttributes) SetEndToEndConversionRate(v float64) {
	o.EndToEndConversionRate = &v
}

// GetEndToEndElapsedTime returns the EndToEndElapsedTime field value if set, zero value otherwise.
func (o *FunnelResponseDataAttributes) GetEndToEndElapsedTime() FunnelResponseElapsedTime {
	if o == nil || o.EndToEndElapsedTime == nil {
		var ret FunnelResponseElapsedTime
		return ret
	}
	return *o.EndToEndElapsedTime
}

// GetEndToEndElapsedTimeOk returns a tuple with the EndToEndElapsedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelResponseDataAttributes) GetEndToEndElapsedTimeOk() (*FunnelResponseElapsedTime, bool) {
	if o == nil || o.EndToEndElapsedTime == nil {
		return nil, false
	}
	return o.EndToEndElapsedTime, true
}

// HasEndToEndElapsedTime returns a boolean if a field has been set.
func (o *FunnelResponseDataAttributes) HasEndToEndElapsedTime() bool {
	return o != nil && o.EndToEndElapsedTime != nil
}

// SetEndToEndElapsedTime gets a reference to the given FunnelResponseElapsedTime and assigns it to the EndToEndElapsedTime field.
func (o *FunnelResponseDataAttributes) SetEndToEndElapsedTime(v FunnelResponseElapsedTime) {
	o.EndToEndElapsedTime = &v
}

// GetFunnelSteps returns the FunnelSteps field value if set, zero value otherwise.
func (o *FunnelResponseDataAttributes) GetFunnelSteps() []FunnelResponseDataAttributesFunnelStepsItems {
	if o == nil || o.FunnelSteps == nil {
		var ret []FunnelResponseDataAttributesFunnelStepsItems
		return ret
	}
	return o.FunnelSteps
}

// GetFunnelStepsOk returns a tuple with the FunnelSteps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelResponseDataAttributes) GetFunnelStepsOk() (*[]FunnelResponseDataAttributesFunnelStepsItems, bool) {
	if o == nil || o.FunnelSteps == nil {
		return nil, false
	}
	return &o.FunnelSteps, true
}

// HasFunnelSteps returns a boolean if a field has been set.
func (o *FunnelResponseDataAttributes) HasFunnelSteps() bool {
	return o != nil && o.FunnelSteps != nil
}

// SetFunnelSteps gets a reference to the given []FunnelResponseDataAttributesFunnelStepsItems and assigns it to the FunnelSteps field.
func (o *FunnelResponseDataAttributes) SetFunnelSteps(v []FunnelResponseDataAttributesFunnelStepsItems) {
	o.FunnelSteps = v
}

// GetInitialCount returns the InitialCount field value if set, zero value otherwise.
func (o *FunnelResponseDataAttributes) GetInitialCount() int64 {
	if o == nil || o.InitialCount == nil {
		var ret int64
		return ret
	}
	return *o.InitialCount
}

// GetInitialCountOk returns a tuple with the InitialCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelResponseDataAttributes) GetInitialCountOk() (*int64, bool) {
	if o == nil || o.InitialCount == nil {
		return nil, false
	}
	return o.InitialCount, true
}

// HasInitialCount returns a boolean if a field has been set.
func (o *FunnelResponseDataAttributes) HasInitialCount() bool {
	return o != nil && o.InitialCount != nil
}

// SetInitialCount gets a reference to the given int64 and assigns it to the InitialCount field.
func (o *FunnelResponseDataAttributes) SetInitialCount(v int64) {
	o.InitialCount = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FunnelResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.EndToEndConversionRate != nil {
		toSerialize["end_to_end_conversion_rate"] = o.EndToEndConversionRate
	}
	if o.EndToEndElapsedTime != nil {
		toSerialize["end_to_end_elapsed_time"] = o.EndToEndElapsedTime
	}
	if o.FunnelSteps != nil {
		toSerialize["funnel_steps"] = o.FunnelSteps
	}
	if o.InitialCount != nil {
		toSerialize["initial_count"] = o.InitialCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FunnelResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EndToEndConversionRate *float64                                       `json:"end_to_end_conversion_rate,omitempty"`
		EndToEndElapsedTime    *FunnelResponseElapsedTime                     `json:"end_to_end_elapsed_time,omitempty"`
		FunnelSteps            []FunnelResponseDataAttributesFunnelStepsItems `json:"funnel_steps,omitempty"`
		InitialCount           *int64                                         `json:"initial_count,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"end_to_end_conversion_rate", "end_to_end_elapsed_time", "funnel_steps", "initial_count"})
	} else {
		return err
	}

	hasInvalidField := false
	o.EndToEndConversionRate = all.EndToEndConversionRate
	if all.EndToEndElapsedTime != nil && all.EndToEndElapsedTime.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.EndToEndElapsedTime = all.EndToEndElapsedTime
	o.FunnelSteps = all.FunnelSteps
	o.InitialCount = all.InitialCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
