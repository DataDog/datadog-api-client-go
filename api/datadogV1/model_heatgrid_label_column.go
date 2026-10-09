// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridLabelColumn Configuration of the group label column.
type HeatgridLabelColumn struct {
	// Width of the label column.
	Width HeatgridLabelColumnWidth `json:"width"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewHeatgridLabelColumn instantiates a new HeatgridLabelColumn object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHeatgridLabelColumn(width HeatgridLabelColumnWidth) *HeatgridLabelColumn {
	this := HeatgridLabelColumn{}
	this.Width = width
	return &this
}

// NewHeatgridLabelColumnWithDefaults instantiates a new HeatgridLabelColumn object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHeatgridLabelColumnWithDefaults() *HeatgridLabelColumn {
	this := HeatgridLabelColumn{}
	return &this
}

// GetWidth returns the Width field value.
func (o *HeatgridLabelColumn) GetWidth() HeatgridLabelColumnWidth {
	if o == nil {
		var ret HeatgridLabelColumnWidth
		return ret
	}
	return o.Width
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
func (o *HeatgridLabelColumn) GetWidthOk() (*HeatgridLabelColumnWidth, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Width, true
}

// SetWidth sets field value.
func (o *HeatgridLabelColumn) SetWidth(v HeatgridLabelColumnWidth) {
	o.Width = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HeatgridLabelColumn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["width"] = o.Width
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HeatgridLabelColumn) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Width *HeatgridLabelColumnWidth `json:"width"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Width == nil {
		return fmt.Errorf("required field width missing")
	}

	hasInvalidField := false
	if !all.Width.IsValid() {
		hasInvalidField = true
	} else {
		o.Width = *all.Width
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
