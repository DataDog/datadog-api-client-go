// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsRetentionGridCohortCell One cell of the retention grid, holding the result for a single cohort over a single return period.
// Aggregated rows omit the time and count fields.
type ProductAnalyticsRetentionGridCohortCell struct {
	// Number of entities that returned during the period.
	CellCount *int64 `json:"cell_count,omitempty"`
	// Fraction of the cohort that returned, between `0` and `1`.
	CellRate *float64 `json:"cell_rate,omitempty"`
	// Change in the metric relative to the cohort baseline.
	CellRelativeValueChange datadog.NullableFloat64 `json:"cell_relative_value_change,omitempty"`
	// Value of the computed metric, when a metric other than the retention rate is requested.
	CellValue datadog.NullableFloat64 `json:"cell_value,omitempty"`
	// Whether the return period is still open, so the numbers are not yet final.
	IsPartialData *bool `json:"is_partial_data,omitempty"`
	// End of the return period, in epoch milliseconds.
	ReturnPeriodEndTime *int64 `json:"return_period_end_time,omitempty"`
	// Zero-based index of the return period this cell belongs to.
	ReturnPeriodIndex *int64 `json:"return_period_index,omitempty"`
	// Start of the return period, in epoch milliseconds.
	ReturnPeriodStartTime *int64 `json:"return_period_start_time,omitempty"`
	// Whether the row holds one cohort's own numbers, or the weighted roll-up across every cohort.
	Type *ProductAnalyticsRetentionGridCohortType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsRetentionGridCohortCell instantiates a new ProductAnalyticsRetentionGridCohortCell object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsRetentionGridCohortCell() *ProductAnalyticsRetentionGridCohortCell {
	this := ProductAnalyticsRetentionGridCohortCell{}
	return &this
}

// NewProductAnalyticsRetentionGridCohortCellWithDefaults instantiates a new ProductAnalyticsRetentionGridCohortCell object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsRetentionGridCohortCellWithDefaults() *ProductAnalyticsRetentionGridCohortCell {
	this := ProductAnalyticsRetentionGridCohortCell{}
	return &this
}

// GetCellCount returns the CellCount field value if set, zero value otherwise.
func (o *ProductAnalyticsRetentionGridCohortCell) GetCellCount() int64 {
	if o == nil || o.CellCount == nil {
		var ret int64
		return ret
	}
	return *o.CellCount
}

// GetCellCountOk returns a tuple with the CellCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionGridCohortCell) GetCellCountOk() (*int64, bool) {
	if o == nil || o.CellCount == nil {
		return nil, false
	}
	return o.CellCount, true
}

// HasCellCount returns a boolean if a field has been set.
func (o *ProductAnalyticsRetentionGridCohortCell) HasCellCount() bool {
	return o != nil && o.CellCount != nil
}

// SetCellCount gets a reference to the given int64 and assigns it to the CellCount field.
func (o *ProductAnalyticsRetentionGridCohortCell) SetCellCount(v int64) {
	o.CellCount = &v
}

// GetCellRate returns the CellRate field value if set, zero value otherwise.
func (o *ProductAnalyticsRetentionGridCohortCell) GetCellRate() float64 {
	if o == nil || o.CellRate == nil {
		var ret float64
		return ret
	}
	return *o.CellRate
}

// GetCellRateOk returns a tuple with the CellRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionGridCohortCell) GetCellRateOk() (*float64, bool) {
	if o == nil || o.CellRate == nil {
		return nil, false
	}
	return o.CellRate, true
}

// HasCellRate returns a boolean if a field has been set.
func (o *ProductAnalyticsRetentionGridCohortCell) HasCellRate() bool {
	return o != nil && o.CellRate != nil
}

// SetCellRate gets a reference to the given float64 and assigns it to the CellRate field.
func (o *ProductAnalyticsRetentionGridCohortCell) SetCellRate(v float64) {
	o.CellRate = &v
}

// GetCellRelativeValueChange returns the CellRelativeValueChange field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAnalyticsRetentionGridCohortCell) GetCellRelativeValueChange() float64 {
	if o == nil || o.CellRelativeValueChange.Get() == nil {
		var ret float64
		return ret
	}
	return *o.CellRelativeValueChange.Get()
}

// GetCellRelativeValueChangeOk returns a tuple with the CellRelativeValueChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ProductAnalyticsRetentionGridCohortCell) GetCellRelativeValueChangeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CellRelativeValueChange.Get(), o.CellRelativeValueChange.IsSet()
}

// HasCellRelativeValueChange returns a boolean if a field has been set.
func (o *ProductAnalyticsRetentionGridCohortCell) HasCellRelativeValueChange() bool {
	return o != nil && o.CellRelativeValueChange.IsSet()
}

// SetCellRelativeValueChange gets a reference to the given datadog.NullableFloat64 and assigns it to the CellRelativeValueChange field.
func (o *ProductAnalyticsRetentionGridCohortCell) SetCellRelativeValueChange(v float64) {
	o.CellRelativeValueChange.Set(&v)
}

// SetCellRelativeValueChangeNil sets the value for CellRelativeValueChange to be an explicit nil.
func (o *ProductAnalyticsRetentionGridCohortCell) SetCellRelativeValueChangeNil() {
	o.CellRelativeValueChange.Set(nil)
}

// UnsetCellRelativeValueChange ensures that no value is present for CellRelativeValueChange, not even an explicit nil.
func (o *ProductAnalyticsRetentionGridCohortCell) UnsetCellRelativeValueChange() {
	o.CellRelativeValueChange.Unset()
}

// GetCellValue returns the CellValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAnalyticsRetentionGridCohortCell) GetCellValue() float64 {
	if o == nil || o.CellValue.Get() == nil {
		var ret float64
		return ret
	}
	return *o.CellValue.Get()
}

// GetCellValueOk returns a tuple with the CellValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ProductAnalyticsRetentionGridCohortCell) GetCellValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CellValue.Get(), o.CellValue.IsSet()
}

// HasCellValue returns a boolean if a field has been set.
func (o *ProductAnalyticsRetentionGridCohortCell) HasCellValue() bool {
	return o != nil && o.CellValue.IsSet()
}

// SetCellValue gets a reference to the given datadog.NullableFloat64 and assigns it to the CellValue field.
func (o *ProductAnalyticsRetentionGridCohortCell) SetCellValue(v float64) {
	o.CellValue.Set(&v)
}

// SetCellValueNil sets the value for CellValue to be an explicit nil.
func (o *ProductAnalyticsRetentionGridCohortCell) SetCellValueNil() {
	o.CellValue.Set(nil)
}

// UnsetCellValue ensures that no value is present for CellValue, not even an explicit nil.
func (o *ProductAnalyticsRetentionGridCohortCell) UnsetCellValue() {
	o.CellValue.Unset()
}

// GetIsPartialData returns the IsPartialData field value if set, zero value otherwise.
func (o *ProductAnalyticsRetentionGridCohortCell) GetIsPartialData() bool {
	if o == nil || o.IsPartialData == nil {
		var ret bool
		return ret
	}
	return *o.IsPartialData
}

// GetIsPartialDataOk returns a tuple with the IsPartialData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionGridCohortCell) GetIsPartialDataOk() (*bool, bool) {
	if o == nil || o.IsPartialData == nil {
		return nil, false
	}
	return o.IsPartialData, true
}

// HasIsPartialData returns a boolean if a field has been set.
func (o *ProductAnalyticsRetentionGridCohortCell) HasIsPartialData() bool {
	return o != nil && o.IsPartialData != nil
}

// SetIsPartialData gets a reference to the given bool and assigns it to the IsPartialData field.
func (o *ProductAnalyticsRetentionGridCohortCell) SetIsPartialData(v bool) {
	o.IsPartialData = &v
}

// GetReturnPeriodEndTime returns the ReturnPeriodEndTime field value if set, zero value otherwise.
func (o *ProductAnalyticsRetentionGridCohortCell) GetReturnPeriodEndTime() int64 {
	if o == nil || o.ReturnPeriodEndTime == nil {
		var ret int64
		return ret
	}
	return *o.ReturnPeriodEndTime
}

// GetReturnPeriodEndTimeOk returns a tuple with the ReturnPeriodEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionGridCohortCell) GetReturnPeriodEndTimeOk() (*int64, bool) {
	if o == nil || o.ReturnPeriodEndTime == nil {
		return nil, false
	}
	return o.ReturnPeriodEndTime, true
}

// HasReturnPeriodEndTime returns a boolean if a field has been set.
func (o *ProductAnalyticsRetentionGridCohortCell) HasReturnPeriodEndTime() bool {
	return o != nil && o.ReturnPeriodEndTime != nil
}

// SetReturnPeriodEndTime gets a reference to the given int64 and assigns it to the ReturnPeriodEndTime field.
func (o *ProductAnalyticsRetentionGridCohortCell) SetReturnPeriodEndTime(v int64) {
	o.ReturnPeriodEndTime = &v
}

// GetReturnPeriodIndex returns the ReturnPeriodIndex field value if set, zero value otherwise.
func (o *ProductAnalyticsRetentionGridCohortCell) GetReturnPeriodIndex() int64 {
	if o == nil || o.ReturnPeriodIndex == nil {
		var ret int64
		return ret
	}
	return *o.ReturnPeriodIndex
}

// GetReturnPeriodIndexOk returns a tuple with the ReturnPeriodIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionGridCohortCell) GetReturnPeriodIndexOk() (*int64, bool) {
	if o == nil || o.ReturnPeriodIndex == nil {
		return nil, false
	}
	return o.ReturnPeriodIndex, true
}

// HasReturnPeriodIndex returns a boolean if a field has been set.
func (o *ProductAnalyticsRetentionGridCohortCell) HasReturnPeriodIndex() bool {
	return o != nil && o.ReturnPeriodIndex != nil
}

// SetReturnPeriodIndex gets a reference to the given int64 and assigns it to the ReturnPeriodIndex field.
func (o *ProductAnalyticsRetentionGridCohortCell) SetReturnPeriodIndex(v int64) {
	o.ReturnPeriodIndex = &v
}

// GetReturnPeriodStartTime returns the ReturnPeriodStartTime field value if set, zero value otherwise.
func (o *ProductAnalyticsRetentionGridCohortCell) GetReturnPeriodStartTime() int64 {
	if o == nil || o.ReturnPeriodStartTime == nil {
		var ret int64
		return ret
	}
	return *o.ReturnPeriodStartTime
}

// GetReturnPeriodStartTimeOk returns a tuple with the ReturnPeriodStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionGridCohortCell) GetReturnPeriodStartTimeOk() (*int64, bool) {
	if o == nil || o.ReturnPeriodStartTime == nil {
		return nil, false
	}
	return o.ReturnPeriodStartTime, true
}

// HasReturnPeriodStartTime returns a boolean if a field has been set.
func (o *ProductAnalyticsRetentionGridCohortCell) HasReturnPeriodStartTime() bool {
	return o != nil && o.ReturnPeriodStartTime != nil
}

// SetReturnPeriodStartTime gets a reference to the given int64 and assigns it to the ReturnPeriodStartTime field.
func (o *ProductAnalyticsRetentionGridCohortCell) SetReturnPeriodStartTime(v int64) {
	o.ReturnPeriodStartTime = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ProductAnalyticsRetentionGridCohortCell) GetType() ProductAnalyticsRetentionGridCohortType {
	if o == nil || o.Type == nil {
		var ret ProductAnalyticsRetentionGridCohortType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionGridCohortCell) GetTypeOk() (*ProductAnalyticsRetentionGridCohortType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ProductAnalyticsRetentionGridCohortCell) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given ProductAnalyticsRetentionGridCohortType and assigns it to the Type field.
func (o *ProductAnalyticsRetentionGridCohortCell) SetType(v ProductAnalyticsRetentionGridCohortType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsRetentionGridCohortCell) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CellCount != nil {
		toSerialize["cell_count"] = o.CellCount
	}
	if o.CellRate != nil {
		toSerialize["cell_rate"] = o.CellRate
	}
	if o.CellRelativeValueChange.IsSet() {
		toSerialize["cell_relative_value_change"] = o.CellRelativeValueChange.Get()
	}
	if o.CellValue.IsSet() {
		toSerialize["cell_value"] = o.CellValue.Get()
	}
	if o.IsPartialData != nil {
		toSerialize["is_partial_data"] = o.IsPartialData
	}
	if o.ReturnPeriodEndTime != nil {
		toSerialize["return_period_end_time"] = o.ReturnPeriodEndTime
	}
	if o.ReturnPeriodIndex != nil {
		toSerialize["return_period_index"] = o.ReturnPeriodIndex
	}
	if o.ReturnPeriodStartTime != nil {
		toSerialize["return_period_start_time"] = o.ReturnPeriodStartTime
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsRetentionGridCohortCell) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CellCount               *int64                                   `json:"cell_count,omitempty"`
		CellRate                *float64                                 `json:"cell_rate,omitempty"`
		CellRelativeValueChange datadog.NullableFloat64                  `json:"cell_relative_value_change,omitempty"`
		CellValue               datadog.NullableFloat64                  `json:"cell_value,omitempty"`
		IsPartialData           *bool                                    `json:"is_partial_data,omitempty"`
		ReturnPeriodEndTime     *int64                                   `json:"return_period_end_time,omitempty"`
		ReturnPeriodIndex       *int64                                   `json:"return_period_index,omitempty"`
		ReturnPeriodStartTime   *int64                                   `json:"return_period_start_time,omitempty"`
		Type                    *ProductAnalyticsRetentionGridCohortType `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cell_count", "cell_rate", "cell_relative_value_change", "cell_value", "is_partial_data", "return_period_end_time", "return_period_index", "return_period_start_time", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CellCount = all.CellCount
	o.CellRate = all.CellRate
	o.CellRelativeValueChange = all.CellRelativeValueChange
	o.CellValue = all.CellValue
	o.IsPartialData = all.IsPartialData
	o.ReturnPeriodEndTime = all.ReturnPeriodEndTime
	o.ReturnPeriodIndex = all.ReturnPeriodIndex
	o.ReturnPeriodStartTime = all.ReturnPeriodStartTime
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
