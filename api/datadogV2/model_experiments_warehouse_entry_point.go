// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsWarehouseEntryPoint Optional Warehouse measure that scopes analyzed subjects.
type ExperimentsWarehouseEntryPoint struct {
	// Complete ordered Warehouse entry-point filter set.
	Filters []ExperimentsWarehouseExposureFilter `json:"filters"`
	// Warehouse measure UUID.
	MeasureId string `json:"measure_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsWarehouseEntryPoint instantiates a new ExperimentsWarehouseEntryPoint object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsWarehouseEntryPoint(filters []ExperimentsWarehouseExposureFilter, measureId string) *ExperimentsWarehouseEntryPoint {
	this := ExperimentsWarehouseEntryPoint{}
	this.Filters = filters
	this.MeasureId = measureId
	return &this
}

// NewExperimentsWarehouseEntryPointWithDefaults instantiates a new ExperimentsWarehouseEntryPoint object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsWarehouseEntryPointWithDefaults() *ExperimentsWarehouseEntryPoint {
	this := ExperimentsWarehouseEntryPoint{}
	return &this
}

// GetFilters returns the Filters field value.
func (o *ExperimentsWarehouseEntryPoint) GetFilters() []ExperimentsWarehouseExposureFilter {
	if o == nil {
		var ret []ExperimentsWarehouseExposureFilter
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value
// and a boolean to check if the value has been set.
func (o *ExperimentsWarehouseEntryPoint) GetFiltersOk() (*[]ExperimentsWarehouseExposureFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filters, true
}

// SetFilters sets field value.
func (o *ExperimentsWarehouseEntryPoint) SetFilters(v []ExperimentsWarehouseExposureFilter) {
	o.Filters = v
}

// GetMeasureId returns the MeasureId field value.
func (o *ExperimentsWarehouseEntryPoint) GetMeasureId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MeasureId
}

// GetMeasureIdOk returns a tuple with the MeasureId field value
// and a boolean to check if the value has been set.
func (o *ExperimentsWarehouseEntryPoint) GetMeasureIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MeasureId, true
}

// SetMeasureId sets field value.
func (o *ExperimentsWarehouseEntryPoint) SetMeasureId(v string) {
	o.MeasureId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsWarehouseEntryPoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["filters"] = o.Filters
	toSerialize["measure_id"] = o.MeasureId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsWarehouseEntryPoint) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Filters   *[]ExperimentsWarehouseExposureFilter `json:"filters"`
		MeasureId *string                               `json:"measure_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Filters == nil {
		return fmt.Errorf("required field filters missing")
	}
	if all.MeasureId == nil {
		return fmt.Errorf("required field measure_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"filters", "measure_id"})
	} else {
		return err
	}
	o.Filters = *all.Filters
	o.MeasureId = *all.MeasureId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

// NullableExperimentsWarehouseEntryPoint handles when a null is used for ExperimentsWarehouseEntryPoint.
type NullableExperimentsWarehouseEntryPoint struct {
	value *ExperimentsWarehouseEntryPoint
	isSet bool
}

// Get returns the associated value.
func (v NullableExperimentsWarehouseEntryPoint) Get() *ExperimentsWarehouseEntryPoint {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableExperimentsWarehouseEntryPoint) Set(val *ExperimentsWarehouseEntryPoint) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableExperimentsWarehouseEntryPoint) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableExperimentsWarehouseEntryPoint) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableExperimentsWarehouseEntryPoint initializes the struct as if Set has been called.
func NewNullableExperimentsWarehouseEntryPoint(val *ExperimentsWarehouseEntryPoint) *NullableExperimentsWarehouseEntryPoint {
	return &NullableExperimentsWarehouseEntryPoint{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableExperimentsWarehouseEntryPoint) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableExperimentsWarehouseEntryPoint) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
