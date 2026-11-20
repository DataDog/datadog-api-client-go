// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableTrendCellDisplayModeOptions Test-only generated schema for the dashboard versioning proof of concept.
type TableTrendCellDisplayModeOptions struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	TrendType TableTrendType `json:"trend_type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	YScale TableTrendCellDisplayModeOptionsYScale `json:"y_scale"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTableTrendCellDisplayModeOptions instantiates a new TableTrendCellDisplayModeOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTableTrendCellDisplayModeOptions(trendType TableTrendType, yScale TableTrendCellDisplayModeOptionsYScale) *TableTrendCellDisplayModeOptions {
	this := TableTrendCellDisplayModeOptions{}
	this.TrendType = trendType
	this.YScale = yScale
	return &this
}

// NewTableTrendCellDisplayModeOptionsWithDefaults instantiates a new TableTrendCellDisplayModeOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTableTrendCellDisplayModeOptionsWithDefaults() *TableTrendCellDisplayModeOptions {
	this := TableTrendCellDisplayModeOptions{}
	return &this
}

// GetTrendType returns the TrendType field value.
func (o *TableTrendCellDisplayModeOptions) GetTrendType() TableTrendType {
	if o == nil {
		var ret TableTrendType
		return ret
	}
	return o.TrendType
}

// GetTrendTypeOk returns a tuple with the TrendType field value
// and a boolean to check if the value has been set.
func (o *TableTrendCellDisplayModeOptions) GetTrendTypeOk() (*TableTrendType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrendType, true
}

// SetTrendType sets field value.
func (o *TableTrendCellDisplayModeOptions) SetTrendType(v TableTrendType) {
	o.TrendType = v
}

// GetYScale returns the YScale field value.
func (o *TableTrendCellDisplayModeOptions) GetYScale() TableTrendCellDisplayModeOptionsYScale {
	if o == nil {
		var ret TableTrendCellDisplayModeOptionsYScale
		return ret
	}
	return o.YScale
}

// GetYScaleOk returns a tuple with the YScale field value
// and a boolean to check if the value has been set.
func (o *TableTrendCellDisplayModeOptions) GetYScaleOk() (*TableTrendCellDisplayModeOptionsYScale, bool) {
	if o == nil {
		return nil, false
	}
	return &o.YScale, true
}

// SetYScale sets field value.
func (o *TableTrendCellDisplayModeOptions) SetYScale(v TableTrendCellDisplayModeOptionsYScale) {
	o.YScale = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TableTrendCellDisplayModeOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["trend_type"] = o.TrendType
	toSerialize["y_scale"] = o.YScale
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TableTrendCellDisplayModeOptions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TrendType *TableTrendType                         `json:"trend_type"`
		YScale    *TableTrendCellDisplayModeOptionsYScale `json:"y_scale"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.TrendType == nil {
		return fmt.Errorf("required field trend_type missing")
	}
	if all.YScale == nil {
		return fmt.Errorf("required field y_scale missing")
	}

	hasInvalidField := false
	if !all.TrendType.IsValid() {
		hasInvalidField = true
	} else {
		o.TrendType = *all.TrendType
	}
	if !all.YScale.IsValid() {
		hasInvalidField = true
	} else {
		o.YScale = *all.YScale
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
