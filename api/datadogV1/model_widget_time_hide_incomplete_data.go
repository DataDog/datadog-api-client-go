// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetTimeHideIncompleteData Widget time setting with hide incomplete cost data option.
type WidgetTimeHideIncompleteData struct {
	// Whether to hide incomplete cost data in the widget.
	HideIncompleteCostData bool `json:"hide_incomplete_cost_data"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewWidgetTimeHideIncompleteData instantiates a new WidgetTimeHideIncompleteData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWidgetTimeHideIncompleteData(hideIncompleteCostData bool) *WidgetTimeHideIncompleteData {
	this := WidgetTimeHideIncompleteData{}
	this.HideIncompleteCostData = hideIncompleteCostData
	return &this
}

// NewWidgetTimeHideIncompleteDataWithDefaults instantiates a new WidgetTimeHideIncompleteData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWidgetTimeHideIncompleteDataWithDefaults() *WidgetTimeHideIncompleteData {
	this := WidgetTimeHideIncompleteData{}
	return &this
}

// GetHideIncompleteCostData returns the HideIncompleteCostData field value.
func (o *WidgetTimeHideIncompleteData) GetHideIncompleteCostData() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.HideIncompleteCostData
}

// GetHideIncompleteCostDataOk returns a tuple with the HideIncompleteCostData field value
// and a boolean to check if the value has been set.
func (o *WidgetTimeHideIncompleteData) GetHideIncompleteCostDataOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HideIncompleteCostData, true
}

// SetHideIncompleteCostData sets field value.
func (o *WidgetTimeHideIncompleteData) SetHideIncompleteCostData(v bool) {
	o.HideIncompleteCostData = v
}

// MarshalJSON serializes the struct using spec logic.
func (o WidgetTimeHideIncompleteData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["hide_incomplete_cost_data"] = o.HideIncompleteCostData
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WidgetTimeHideIncompleteData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		HideIncompleteCostData *bool `json:"hide_incomplete_cost_data"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.HideIncompleteCostData == nil {
		return fmt.Errorf("required field hide_incomplete_cost_data missing")
	}
	o.HideIncompleteCostData = *all.HideIncompleteCostData

	return nil
}
