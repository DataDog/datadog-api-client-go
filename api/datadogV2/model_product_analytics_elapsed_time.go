// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsElapsedTime Elapsed time statistics (min/max/avg in milliseconds).
type ProductAnalyticsElapsedTime struct {
	// Average elapsed time to reach the next step, in milliseconds.
	Avg int64 `json:"avg"`
	// Maximum elapsed time to reach the next step, in milliseconds.
	Max int64 `json:"max"`
	// Minimum elapsed time to reach the next step, in milliseconds.
	Min int64 `json:"min"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsElapsedTime instantiates a new ProductAnalyticsElapsedTime object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsElapsedTime(avg int64, max int64, min int64) *ProductAnalyticsElapsedTime {
	this := ProductAnalyticsElapsedTime{}
	this.Avg = avg
	this.Max = max
	this.Min = min
	return &this
}

// NewProductAnalyticsElapsedTimeWithDefaults instantiates a new ProductAnalyticsElapsedTime object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsElapsedTimeWithDefaults() *ProductAnalyticsElapsedTime {
	this := ProductAnalyticsElapsedTime{}
	return &this
}

// GetAvg returns the Avg field value.
func (o *ProductAnalyticsElapsedTime) GetAvg() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Avg
}

// GetAvgOk returns a tuple with the Avg field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsElapsedTime) GetAvgOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Avg, true
}

// SetAvg sets field value.
func (o *ProductAnalyticsElapsedTime) SetAvg(v int64) {
	o.Avg = v
}

// GetMax returns the Max field value.
func (o *ProductAnalyticsElapsedTime) GetMax() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Max
}

// GetMaxOk returns a tuple with the Max field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsElapsedTime) GetMaxOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Max, true
}

// SetMax sets field value.
func (o *ProductAnalyticsElapsedTime) SetMax(v int64) {
	o.Max = v
}

// GetMin returns the Min field value.
func (o *ProductAnalyticsElapsedTime) GetMin() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Min
}

// GetMinOk returns a tuple with the Min field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsElapsedTime) GetMinOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Min, true
}

// SetMin sets field value.
func (o *ProductAnalyticsElapsedTime) SetMin(v int64) {
	o.Min = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsElapsedTime) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["avg"] = o.Avg
	toSerialize["max"] = o.Max
	toSerialize["min"] = o.Min

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsElapsedTime) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Avg *int64 `json:"avg"`
		Max *int64 `json:"max"`
		Min *int64 `json:"min"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Avg == nil {
		return fmt.Errorf("required field avg missing")
	}
	if all.Max == nil {
		return fmt.Errorf("required field max missing")
	}
	if all.Min == nil {
		return fmt.Errorf("required field min missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"avg", "max", "min"})
	} else {
		return err
	}
	o.Avg = *all.Avg
	o.Max = *all.Max
	o.Min = *all.Min

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
