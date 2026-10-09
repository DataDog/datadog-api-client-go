// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridLegend Legend configuration for the heatgrid widget.
type HeatgridLegend struct {
	// Whether to display the legend caption.
	ShowCaption *bool `json:"show_caption,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewHeatgridLegend instantiates a new HeatgridLegend object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHeatgridLegend() *HeatgridLegend {
	this := HeatgridLegend{}
	return &this
}

// NewHeatgridLegendWithDefaults instantiates a new HeatgridLegend object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHeatgridLegendWithDefaults() *HeatgridLegend {
	this := HeatgridLegend{}
	return &this
}

// GetShowCaption returns the ShowCaption field value if set, zero value otherwise.
func (o *HeatgridLegend) GetShowCaption() bool {
	if o == nil || o.ShowCaption == nil {
		var ret bool
		return ret
	}
	return *o.ShowCaption
}

// GetShowCaptionOk returns a tuple with the ShowCaption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatgridLegend) GetShowCaptionOk() (*bool, bool) {
	if o == nil || o.ShowCaption == nil {
		return nil, false
	}
	return o.ShowCaption, true
}

// HasShowCaption returns a boolean if a field has been set.
func (o *HeatgridLegend) HasShowCaption() bool {
	return o != nil && o.ShowCaption != nil
}

// SetShowCaption gets a reference to the given bool and assigns it to the ShowCaption field.
func (o *HeatgridLegend) SetShowCaption(v bool) {
	o.ShowCaption = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o HeatgridLegend) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ShowCaption != nil {
		toSerialize["show_caption"] = o.ShowCaption
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HeatgridLegend) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ShowCaption *bool `json:"show_caption,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.ShowCaption = all.ShowCaption

	return nil
}
