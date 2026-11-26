// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesRequestStyle Test-only generated schema for the dashboard versioning proof of concept.
type TimeseriesRequestStyle struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	ColorOrder *TimeseriesColorOrder `json:"color_order,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	HasValueLabels *bool `json:"has_value_labels,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	LineType *WidgetLineType `json:"line_type,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	LineWidth *WidgetLineWidth `json:"line_width,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	OrderBy *WidgetStyleOrderBy `json:"order_by,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	OrderReverse *bool `json:"order_reverse,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Palette *string `json:"palette,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ValueLabelsDensity *ValueLabelsDensity `json:"value_labels_density,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTimeseriesRequestStyle instantiates a new TimeseriesRequestStyle object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTimeseriesRequestStyle() *TimeseriesRequestStyle {
	this := TimeseriesRequestStyle{}
	return &this
}

// NewTimeseriesRequestStyleWithDefaults instantiates a new TimeseriesRequestStyle object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTimeseriesRequestStyleWithDefaults() *TimeseriesRequestStyle {
	this := TimeseriesRequestStyle{}
	return &this
}

// GetColorOrder returns the ColorOrder field value if set, zero value otherwise.
func (o *TimeseriesRequestStyle) GetColorOrder() TimeseriesColorOrder {
	if o == nil || o.ColorOrder == nil {
		var ret TimeseriesColorOrder
		return ret
	}
	return *o.ColorOrder
}

// GetColorOrderOk returns a tuple with the ColorOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesRequestStyle) GetColorOrderOk() (*TimeseriesColorOrder, bool) {
	if o == nil || o.ColorOrder == nil {
		return nil, false
	}
	return o.ColorOrder, true
}

// HasColorOrder returns a boolean if a field has been set.
func (o *TimeseriesRequestStyle) HasColorOrder() bool {
	return o != nil && o.ColorOrder != nil
}

// SetColorOrder gets a reference to the given TimeseriesColorOrder and assigns it to the ColorOrder field.
func (o *TimeseriesRequestStyle) SetColorOrder(v TimeseriesColorOrder) {
	o.ColorOrder = &v
}

// GetHasValueLabels returns the HasValueLabels field value if set, zero value otherwise.
func (o *TimeseriesRequestStyle) GetHasValueLabels() bool {
	if o == nil || o.HasValueLabels == nil {
		var ret bool
		return ret
	}
	return *o.HasValueLabels
}

// GetHasValueLabelsOk returns a tuple with the HasValueLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesRequestStyle) GetHasValueLabelsOk() (*bool, bool) {
	if o == nil || o.HasValueLabels == nil {
		return nil, false
	}
	return o.HasValueLabels, true
}

// HasHasValueLabels returns a boolean if a field has been set.
func (o *TimeseriesRequestStyle) HasHasValueLabels() bool {
	return o != nil && o.HasValueLabels != nil
}

// SetHasValueLabels gets a reference to the given bool and assigns it to the HasValueLabels field.
func (o *TimeseriesRequestStyle) SetHasValueLabels(v bool) {
	o.HasValueLabels = &v
}

// GetLineType returns the LineType field value if set, zero value otherwise.
func (o *TimeseriesRequestStyle) GetLineType() WidgetLineType {
	if o == nil || o.LineType == nil {
		var ret WidgetLineType
		return ret
	}
	return *o.LineType
}

// GetLineTypeOk returns a tuple with the LineType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesRequestStyle) GetLineTypeOk() (*WidgetLineType, bool) {
	if o == nil || o.LineType == nil {
		return nil, false
	}
	return o.LineType, true
}

// HasLineType returns a boolean if a field has been set.
func (o *TimeseriesRequestStyle) HasLineType() bool {
	return o != nil && o.LineType != nil
}

// SetLineType gets a reference to the given WidgetLineType and assigns it to the LineType field.
func (o *TimeseriesRequestStyle) SetLineType(v WidgetLineType) {
	o.LineType = &v
}

// GetLineWidth returns the LineWidth field value if set, zero value otherwise.
func (o *TimeseriesRequestStyle) GetLineWidth() WidgetLineWidth {
	if o == nil || o.LineWidth == nil {
		var ret WidgetLineWidth
		return ret
	}
	return *o.LineWidth
}

// GetLineWidthOk returns a tuple with the LineWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesRequestStyle) GetLineWidthOk() (*WidgetLineWidth, bool) {
	if o == nil || o.LineWidth == nil {
		return nil, false
	}
	return o.LineWidth, true
}

// HasLineWidth returns a boolean if a field has been set.
func (o *TimeseriesRequestStyle) HasLineWidth() bool {
	return o != nil && o.LineWidth != nil
}

// SetLineWidth gets a reference to the given WidgetLineWidth and assigns it to the LineWidth field.
func (o *TimeseriesRequestStyle) SetLineWidth(v WidgetLineWidth) {
	o.LineWidth = &v
}

// GetOrderBy returns the OrderBy field value if set, zero value otherwise.
func (o *TimeseriesRequestStyle) GetOrderBy() WidgetStyleOrderBy {
	if o == nil || o.OrderBy == nil {
		var ret WidgetStyleOrderBy
		return ret
	}
	return *o.OrderBy
}

// GetOrderByOk returns a tuple with the OrderBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesRequestStyle) GetOrderByOk() (*WidgetStyleOrderBy, bool) {
	if o == nil || o.OrderBy == nil {
		return nil, false
	}
	return o.OrderBy, true
}

// HasOrderBy returns a boolean if a field has been set.
func (o *TimeseriesRequestStyle) HasOrderBy() bool {
	return o != nil && o.OrderBy != nil
}

// SetOrderBy gets a reference to the given WidgetStyleOrderBy and assigns it to the OrderBy field.
func (o *TimeseriesRequestStyle) SetOrderBy(v WidgetStyleOrderBy) {
	o.OrderBy = &v
}

// GetOrderReverse returns the OrderReverse field value if set, zero value otherwise.
func (o *TimeseriesRequestStyle) GetOrderReverse() bool {
	if o == nil || o.OrderReverse == nil {
		var ret bool
		return ret
	}
	return *o.OrderReverse
}

// GetOrderReverseOk returns a tuple with the OrderReverse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesRequestStyle) GetOrderReverseOk() (*bool, bool) {
	if o == nil || o.OrderReverse == nil {
		return nil, false
	}
	return o.OrderReverse, true
}

// HasOrderReverse returns a boolean if a field has been set.
func (o *TimeseriesRequestStyle) HasOrderReverse() bool {
	return o != nil && o.OrderReverse != nil
}

// SetOrderReverse gets a reference to the given bool and assigns it to the OrderReverse field.
func (o *TimeseriesRequestStyle) SetOrderReverse(v bool) {
	o.OrderReverse = &v
}

// GetPalette returns the Palette field value if set, zero value otherwise.
func (o *TimeseriesRequestStyle) GetPalette() string {
	if o == nil || o.Palette == nil {
		var ret string
		return ret
	}
	return *o.Palette
}

// GetPaletteOk returns a tuple with the Palette field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesRequestStyle) GetPaletteOk() (*string, bool) {
	if o == nil || o.Palette == nil {
		return nil, false
	}
	return o.Palette, true
}

// HasPalette returns a boolean if a field has been set.
func (o *TimeseriesRequestStyle) HasPalette() bool {
	return o != nil && o.Palette != nil
}

// SetPalette gets a reference to the given string and assigns it to the Palette field.
func (o *TimeseriesRequestStyle) SetPalette(v string) {
	o.Palette = &v
}

// GetValueLabelsDensity returns the ValueLabelsDensity field value if set, zero value otherwise.
func (o *TimeseriesRequestStyle) GetValueLabelsDensity() ValueLabelsDensity {
	if o == nil || o.ValueLabelsDensity == nil {
		var ret ValueLabelsDensity
		return ret
	}
	return *o.ValueLabelsDensity
}

// GetValueLabelsDensityOk returns a tuple with the ValueLabelsDensity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesRequestStyle) GetValueLabelsDensityOk() (*ValueLabelsDensity, bool) {
	if o == nil || o.ValueLabelsDensity == nil {
		return nil, false
	}
	return o.ValueLabelsDensity, true
}

// HasValueLabelsDensity returns a boolean if a field has been set.
func (o *TimeseriesRequestStyle) HasValueLabelsDensity() bool {
	return o != nil && o.ValueLabelsDensity != nil
}

// SetValueLabelsDensity gets a reference to the given ValueLabelsDensity and assigns it to the ValueLabelsDensity field.
func (o *TimeseriesRequestStyle) SetValueLabelsDensity(v ValueLabelsDensity) {
	o.ValueLabelsDensity = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TimeseriesRequestStyle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ColorOrder != nil {
		toSerialize["color_order"] = o.ColorOrder
	}
	if o.HasValueLabels != nil {
		toSerialize["has_value_labels"] = o.HasValueLabels
	}
	if o.LineType != nil {
		toSerialize["line_type"] = o.LineType
	}
	if o.LineWidth != nil {
		toSerialize["line_width"] = o.LineWidth
	}
	if o.OrderBy != nil {
		toSerialize["order_by"] = o.OrderBy
	}
	if o.OrderReverse != nil {
		toSerialize["order_reverse"] = o.OrderReverse
	}
	if o.Palette != nil {
		toSerialize["palette"] = o.Palette
	}
	if o.ValueLabelsDensity != nil {
		toSerialize["value_labels_density"] = o.ValueLabelsDensity
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TimeseriesRequestStyle) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ColorOrder         *TimeseriesColorOrder `json:"color_order,omitempty"`
		HasValueLabels     *bool                 `json:"has_value_labels,omitempty"`
		LineType           *WidgetLineType       `json:"line_type,omitempty"`
		LineWidth          *WidgetLineWidth      `json:"line_width,omitempty"`
		OrderBy            *WidgetStyleOrderBy   `json:"order_by,omitempty"`
		OrderReverse       *bool                 `json:"order_reverse,omitempty"`
		Palette            *string               `json:"palette,omitempty"`
		ValueLabelsDensity *ValueLabelsDensity   `json:"value_labels_density,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	hasInvalidField := false
	if all.ColorOrder != nil && !all.ColorOrder.IsValid() {
		hasInvalidField = true
	} else {
		o.ColorOrder = all.ColorOrder
	}
	o.HasValueLabels = all.HasValueLabels
	if all.LineType != nil && !all.LineType.IsValid() {
		hasInvalidField = true
	} else {
		o.LineType = all.LineType
	}
	if all.LineWidth != nil && !all.LineWidth.IsValid() {
		hasInvalidField = true
	} else {
		o.LineWidth = all.LineWidth
	}
	if all.OrderBy != nil && !all.OrderBy.IsValid() {
		hasInvalidField = true
	} else {
		o.OrderBy = all.OrderBy
	}
	o.OrderReverse = all.OrderReverse
	o.Palette = all.Palette
	if all.ValueLabelsDensity != nil && !all.ValueLabelsDensity.IsValid() {
		hasInvalidField = true
	} else {
		o.ValueLabelsDensity = all.ValueLabelsDensity
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
