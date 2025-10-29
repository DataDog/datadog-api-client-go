// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GetCohortResponseDataAttributesCohortsItemsValuesItems
type GetCohortResponseDataAttributesCohortsItemsValuesItems struct {
	//
	AbsoluteValue *int64 `json:"absolute_value,omitempty"`
	//
	EndTime *int64 `json:"end_time,omitempty"`
	//
	RelativeValue *float64 `json:"relative_value,omitempty"`
	//
	StartTime *int64 `json:"start_time,omitempty"`
	//
	Window *int64 `json:"window,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGetCohortResponseDataAttributesCohortsItemsValuesItems instantiates a new GetCohortResponseDataAttributesCohortsItemsValuesItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGetCohortResponseDataAttributesCohortsItemsValuesItems() *GetCohortResponseDataAttributesCohortsItemsValuesItems {
	this := GetCohortResponseDataAttributesCohortsItemsValuesItems{}
	return &this
}

// NewGetCohortResponseDataAttributesCohortsItemsValuesItemsWithDefaults instantiates a new GetCohortResponseDataAttributesCohortsItemsValuesItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGetCohortResponseDataAttributesCohortsItemsValuesItemsWithDefaults() *GetCohortResponseDataAttributesCohortsItemsValuesItems {
	this := GetCohortResponseDataAttributesCohortsItemsValuesItems{}
	return &this
}

// GetAbsoluteValue returns the AbsoluteValue field value if set, zero value otherwise.
func (o *GetCohortResponseDataAttributesCohortsItemsValuesItems) GetAbsoluteValue() int64 {
	if o == nil || o.AbsoluteValue == nil {
		var ret int64
		return ret
	}
	return *o.AbsoluteValue
}

// GetAbsoluteValueOk returns a tuple with the AbsoluteValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCohortResponseDataAttributesCohortsItemsValuesItems) GetAbsoluteValueOk() (*int64, bool) {
	if o == nil || o.AbsoluteValue == nil {
		return nil, false
	}
	return o.AbsoluteValue, true
}

// HasAbsoluteValue returns a boolean if a field has been set.
func (o *GetCohortResponseDataAttributesCohortsItemsValuesItems) HasAbsoluteValue() bool {
	return o != nil && o.AbsoluteValue != nil
}

// SetAbsoluteValue gets a reference to the given int64 and assigns it to the AbsoluteValue field.
func (o *GetCohortResponseDataAttributesCohortsItemsValuesItems) SetAbsoluteValue(v int64) {
	o.AbsoluteValue = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *GetCohortResponseDataAttributesCohortsItemsValuesItems) GetEndTime() int64 {
	if o == nil || o.EndTime == nil {
		var ret int64
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCohortResponseDataAttributesCohortsItemsValuesItems) GetEndTimeOk() (*int64, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *GetCohortResponseDataAttributesCohortsItemsValuesItems) HasEndTime() bool {
	return o != nil && o.EndTime != nil
}

// SetEndTime gets a reference to the given int64 and assigns it to the EndTime field.
func (o *GetCohortResponseDataAttributesCohortsItemsValuesItems) SetEndTime(v int64) {
	o.EndTime = &v
}

// GetRelativeValue returns the RelativeValue field value if set, zero value otherwise.
func (o *GetCohortResponseDataAttributesCohortsItemsValuesItems) GetRelativeValue() float64 {
	if o == nil || o.RelativeValue == nil {
		var ret float64
		return ret
	}
	return *o.RelativeValue
}

// GetRelativeValueOk returns a tuple with the RelativeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCohortResponseDataAttributesCohortsItemsValuesItems) GetRelativeValueOk() (*float64, bool) {
	if o == nil || o.RelativeValue == nil {
		return nil, false
	}
	return o.RelativeValue, true
}

// HasRelativeValue returns a boolean if a field has been set.
func (o *GetCohortResponseDataAttributesCohortsItemsValuesItems) HasRelativeValue() bool {
	return o != nil && o.RelativeValue != nil
}

// SetRelativeValue gets a reference to the given float64 and assigns it to the RelativeValue field.
func (o *GetCohortResponseDataAttributesCohortsItemsValuesItems) SetRelativeValue(v float64) {
	o.RelativeValue = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *GetCohortResponseDataAttributesCohortsItemsValuesItems) GetStartTime() int64 {
	if o == nil || o.StartTime == nil {
		var ret int64
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCohortResponseDataAttributesCohortsItemsValuesItems) GetStartTimeOk() (*int64, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *GetCohortResponseDataAttributesCohortsItemsValuesItems) HasStartTime() bool {
	return o != nil && o.StartTime != nil
}

// SetStartTime gets a reference to the given int64 and assigns it to the StartTime field.
func (o *GetCohortResponseDataAttributesCohortsItemsValuesItems) SetStartTime(v int64) {
	o.StartTime = &v
}

// GetWindow returns the Window field value if set, zero value otherwise.
func (o *GetCohortResponseDataAttributesCohortsItemsValuesItems) GetWindow() int64 {
	if o == nil || o.Window == nil {
		var ret int64
		return ret
	}
	return *o.Window
}

// GetWindowOk returns a tuple with the Window field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCohortResponseDataAttributesCohortsItemsValuesItems) GetWindowOk() (*int64, bool) {
	if o == nil || o.Window == nil {
		return nil, false
	}
	return o.Window, true
}

// HasWindow returns a boolean if a field has been set.
func (o *GetCohortResponseDataAttributesCohortsItemsValuesItems) HasWindow() bool {
	return o != nil && o.Window != nil
}

// SetWindow gets a reference to the given int64 and assigns it to the Window field.
func (o *GetCohortResponseDataAttributesCohortsItemsValuesItems) SetWindow(v int64) {
	o.Window = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GetCohortResponseDataAttributesCohortsItemsValuesItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AbsoluteValue != nil {
		toSerialize["absolute_value"] = o.AbsoluteValue
	}
	if o.EndTime != nil {
		toSerialize["end_time"] = o.EndTime
	}
	if o.RelativeValue != nil {
		toSerialize["relative_value"] = o.RelativeValue
	}
	if o.StartTime != nil {
		toSerialize["start_time"] = o.StartTime
	}
	if o.Window != nil {
		toSerialize["window"] = o.Window
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GetCohortResponseDataAttributesCohortsItemsValuesItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AbsoluteValue *int64   `json:"absolute_value,omitempty"`
		EndTime       *int64   `json:"end_time,omitempty"`
		RelativeValue *float64 `json:"relative_value,omitempty"`
		StartTime     *int64   `json:"start_time,omitempty"`
		Window        *int64   `json:"window,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"absolute_value", "end_time", "relative_value", "start_time", "window"})
	} else {
		return err
	}
	o.AbsoluteValue = all.AbsoluteValue
	o.EndTime = all.EndTime
	o.RelativeValue = all.RelativeValue
	o.StartTime = all.StartTime
	o.Window = all.Window

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
