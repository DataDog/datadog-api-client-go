// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CommitmentsUtilizationTimeseriesResponseAttributes Attributes of a timeseries utilization response.
type CommitmentsUtilizationTimeseriesResponseAttributes struct {
	// Metadata reporting that the cost data lacks the unused costs of some commitments, as with Azure exports
	// scoped below the billing profile. Utilization is then overstated, and realized savings and effective
	// savings rate are upper bounds. Omitted when the unused costs are available.
	Meta *CommitmentsUnusedCostMeta `json:"meta,omitempty"`
	// The query that produces the response values, in the same format as dashboard and notebook widget
	// requests, so a client can graph the values in a widget. Each formula's `id` names the response value
	// that it produces.
	QueryDefinition *CommitmentsQueryDefinition `json:"query_definition,omitempty"`
	// Timeseries data as a map of series names to their corresponding value arrays.
	Series map[string][]float64 `json:"series"`
	// Unix timestamps in seconds for the timeseries data points.
	Times []int64 `json:"times"`
	// Unit metadata for a numeric metric. Unset fields are omitted, and the unit is `null` when it is unknown.
	Unit NullableCommitmentsUnit `json:"unit,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCommitmentsUtilizationTimeseriesResponseAttributes instantiates a new CommitmentsUtilizationTimeseriesResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCommitmentsUtilizationTimeseriesResponseAttributes(series map[string][]float64, times []int64) *CommitmentsUtilizationTimeseriesResponseAttributes {
	this := CommitmentsUtilizationTimeseriesResponseAttributes{}
	this.Series = series
	this.Times = times
	return &this
}

// NewCommitmentsUtilizationTimeseriesResponseAttributesWithDefaults instantiates a new CommitmentsUtilizationTimeseriesResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCommitmentsUtilizationTimeseriesResponseAttributesWithDefaults() *CommitmentsUtilizationTimeseriesResponseAttributes {
	this := CommitmentsUtilizationTimeseriesResponseAttributes{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *CommitmentsUtilizationTimeseriesResponseAttributes) GetMeta() CommitmentsUnusedCostMeta {
	if o == nil || o.Meta == nil {
		var ret CommitmentsUnusedCostMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsUtilizationTimeseriesResponseAttributes) GetMetaOk() (*CommitmentsUnusedCostMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *CommitmentsUtilizationTimeseriesResponseAttributes) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given CommitmentsUnusedCostMeta and assigns it to the Meta field.
func (o *CommitmentsUtilizationTimeseriesResponseAttributes) SetMeta(v CommitmentsUnusedCostMeta) {
	o.Meta = &v
}

// GetQueryDefinition returns the QueryDefinition field value if set, zero value otherwise.
func (o *CommitmentsUtilizationTimeseriesResponseAttributes) GetQueryDefinition() CommitmentsQueryDefinition {
	if o == nil || o.QueryDefinition == nil {
		var ret CommitmentsQueryDefinition
		return ret
	}
	return *o.QueryDefinition
}

// GetQueryDefinitionOk returns a tuple with the QueryDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsUtilizationTimeseriesResponseAttributes) GetQueryDefinitionOk() (*CommitmentsQueryDefinition, bool) {
	if o == nil || o.QueryDefinition == nil {
		return nil, false
	}
	return o.QueryDefinition, true
}

// HasQueryDefinition returns a boolean if a field has been set.
func (o *CommitmentsUtilizationTimeseriesResponseAttributes) HasQueryDefinition() bool {
	return o != nil && o.QueryDefinition != nil
}

// SetQueryDefinition gets a reference to the given CommitmentsQueryDefinition and assigns it to the QueryDefinition field.
func (o *CommitmentsUtilizationTimeseriesResponseAttributes) SetQueryDefinition(v CommitmentsQueryDefinition) {
	o.QueryDefinition = &v
}

// GetSeries returns the Series field value.
func (o *CommitmentsUtilizationTimeseriesResponseAttributes) GetSeries() map[string][]float64 {
	if o == nil {
		var ret map[string][]float64
		return ret
	}
	return o.Series
}

// GetSeriesOk returns a tuple with the Series field value
// and a boolean to check if the value has been set.
func (o *CommitmentsUtilizationTimeseriesResponseAttributes) GetSeriesOk() (*map[string][]float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Series, true
}

// SetSeries sets field value.
func (o *CommitmentsUtilizationTimeseriesResponseAttributes) SetSeries(v map[string][]float64) {
	o.Series = v
}

// GetTimes returns the Times field value.
func (o *CommitmentsUtilizationTimeseriesResponseAttributes) GetTimes() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.Times
}

// GetTimesOk returns a tuple with the Times field value
// and a boolean to check if the value has been set.
func (o *CommitmentsUtilizationTimeseriesResponseAttributes) GetTimesOk() (*[]int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Times, true
}

// SetTimes sets field value.
func (o *CommitmentsUtilizationTimeseriesResponseAttributes) SetTimes(v []int64) {
	o.Times = v
}

// GetUnit returns the Unit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommitmentsUtilizationTimeseriesResponseAttributes) GetUnit() CommitmentsUnit {
	if o == nil || o.Unit.Get() == nil {
		var ret CommitmentsUnit
		return ret
	}
	return *o.Unit.Get()
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CommitmentsUtilizationTimeseriesResponseAttributes) GetUnitOk() (*CommitmentsUnit, bool) {
	if o == nil {
		return nil, false
	}
	return o.Unit.Get(), o.Unit.IsSet()
}

// HasUnit returns a boolean if a field has been set.
func (o *CommitmentsUtilizationTimeseriesResponseAttributes) HasUnit() bool {
	return o != nil && o.Unit.IsSet()
}

// SetUnit gets a reference to the given NullableCommitmentsUnit and assigns it to the Unit field.
func (o *CommitmentsUtilizationTimeseriesResponseAttributes) SetUnit(v CommitmentsUnit) {
	o.Unit.Set(&v)
}

// SetUnitNil sets the value for Unit to be an explicit nil.
func (o *CommitmentsUtilizationTimeseriesResponseAttributes) SetUnitNil() {
	o.Unit.Set(nil)
}

// UnsetUnit ensures that no value is present for Unit, not even an explicit nil.
func (o *CommitmentsUtilizationTimeseriesResponseAttributes) UnsetUnit() {
	o.Unit.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o CommitmentsUtilizationTimeseriesResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.QueryDefinition != nil {
		toSerialize["query_definition"] = o.QueryDefinition
	}
	toSerialize["series"] = o.Series
	toSerialize["times"] = o.Times
	if o.Unit.IsSet() {
		toSerialize["unit"] = o.Unit.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CommitmentsUtilizationTimeseriesResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Meta            *CommitmentsUnusedCostMeta  `json:"meta,omitempty"`
		QueryDefinition *CommitmentsQueryDefinition `json:"query_definition,omitempty"`
		Series          *map[string][]float64       `json:"series"`
		Times           *[]int64                    `json:"times"`
		Unit            NullableCommitmentsUnit     `json:"unit,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Series == nil {
		return fmt.Errorf("required field series missing")
	}
	if all.Times == nil {
		return fmt.Errorf("required field times missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"meta", "query_definition", "series", "times", "unit"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta
	if all.QueryDefinition != nil && all.QueryDefinition.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.QueryDefinition = all.QueryDefinition
	o.Series = *all.Series
	o.Times = *all.Times
	o.Unit = all.Unit

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
