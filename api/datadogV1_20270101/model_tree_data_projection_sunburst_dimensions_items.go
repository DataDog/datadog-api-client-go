// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TreeDataProjectionSunburstDimensionsItems Test-only generated schema for the dashboard versioning proof of concept.
type TreeDataProjectionSunburstDimensionsItems struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Column string `json:"column"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Dimension TreeDataProjectionSunburstDimensionsItemsDimension `json:"dimension"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	NumberFormat *WidgetNumberFormat `json:"number_format,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTreeDataProjectionSunburstDimensionsItems instantiates a new TreeDataProjectionSunburstDimensionsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTreeDataProjectionSunburstDimensionsItems(column string, dimension TreeDataProjectionSunburstDimensionsItemsDimension) *TreeDataProjectionSunburstDimensionsItems {
	this := TreeDataProjectionSunburstDimensionsItems{}
	this.Column = column
	this.Dimension = dimension
	return &this
}

// NewTreeDataProjectionSunburstDimensionsItemsWithDefaults instantiates a new TreeDataProjectionSunburstDimensionsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTreeDataProjectionSunburstDimensionsItemsWithDefaults() *TreeDataProjectionSunburstDimensionsItems {
	this := TreeDataProjectionSunburstDimensionsItems{}
	return &this
}

// GetColumn returns the Column field value.
func (o *TreeDataProjectionSunburstDimensionsItems) GetColumn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Column
}

// GetColumnOk returns a tuple with the Column field value
// and a boolean to check if the value has been set.
func (o *TreeDataProjectionSunburstDimensionsItems) GetColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Column, true
}

// SetColumn sets field value.
func (o *TreeDataProjectionSunburstDimensionsItems) SetColumn(v string) {
	o.Column = v
}

// GetDimension returns the Dimension field value.
func (o *TreeDataProjectionSunburstDimensionsItems) GetDimension() TreeDataProjectionSunburstDimensionsItemsDimension {
	if o == nil {
		var ret TreeDataProjectionSunburstDimensionsItemsDimension
		return ret
	}
	return o.Dimension
}

// GetDimensionOk returns a tuple with the Dimension field value
// and a boolean to check if the value has been set.
func (o *TreeDataProjectionSunburstDimensionsItems) GetDimensionOk() (*TreeDataProjectionSunburstDimensionsItemsDimension, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dimension, true
}

// SetDimension sets field value.
func (o *TreeDataProjectionSunburstDimensionsItems) SetDimension(v TreeDataProjectionSunburstDimensionsItemsDimension) {
	o.Dimension = v
}

// GetNumberFormat returns the NumberFormat field value if set, zero value otherwise.
func (o *TreeDataProjectionSunburstDimensionsItems) GetNumberFormat() WidgetNumberFormat {
	if o == nil || o.NumberFormat == nil {
		var ret WidgetNumberFormat
		return ret
	}
	return *o.NumberFormat
}

// GetNumberFormatOk returns a tuple with the NumberFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TreeDataProjectionSunburstDimensionsItems) GetNumberFormatOk() (*WidgetNumberFormat, bool) {
	if o == nil || o.NumberFormat == nil {
		return nil, false
	}
	return o.NumberFormat, true
}

// HasNumberFormat returns a boolean if a field has been set.
func (o *TreeDataProjectionSunburstDimensionsItems) HasNumberFormat() bool {
	return o != nil && o.NumberFormat != nil
}

// SetNumberFormat gets a reference to the given WidgetNumberFormat and assigns it to the NumberFormat field.
func (o *TreeDataProjectionSunburstDimensionsItems) SetNumberFormat(v WidgetNumberFormat) {
	o.NumberFormat = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TreeDataProjectionSunburstDimensionsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["column"] = o.Column
	toSerialize["dimension"] = o.Dimension
	if o.NumberFormat != nil {
		toSerialize["number_format"] = o.NumberFormat
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TreeDataProjectionSunburstDimensionsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Column       *string                                             `json:"column"`
		Dimension    *TreeDataProjectionSunburstDimensionsItemsDimension `json:"dimension"`
		NumberFormat *WidgetNumberFormat                                 `json:"number_format,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Column == nil {
		return fmt.Errorf("required field column missing")
	}
	if all.Dimension == nil {
		return fmt.Errorf("required field dimension missing")
	}

	hasInvalidField := false
	o.Column = *all.Column
	if !all.Dimension.IsValid() {
		hasInvalidField = true
	} else {
		o.Dimension = *all.Dimension
	}
	if all.NumberFormat != nil && all.NumberFormat.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.NumberFormat = all.NumberFormat

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
