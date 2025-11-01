// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FunnelResponseElapsedTime
type FunnelResponseElapsedTime struct {
	//
	Avg *int64 `json:"avg,omitempty"`
	//
	Max *int64 `json:"max,omitempty"`
	//
	Min *int64 `json:"min,omitempty"`
	//
	P5 *int64 `json:"p5,omitempty"`
	//
	P50 *int64 `json:"p50,omitempty"`
	//
	P95 *int64 `json:"p95,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFunnelResponseElapsedTime instantiates a new FunnelResponseElapsedTime object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFunnelResponseElapsedTime() *FunnelResponseElapsedTime {
	this := FunnelResponseElapsedTime{}
	return &this
}

// NewFunnelResponseElapsedTimeWithDefaults instantiates a new FunnelResponseElapsedTime object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFunnelResponseElapsedTimeWithDefaults() *FunnelResponseElapsedTime {
	this := FunnelResponseElapsedTime{}
	return &this
}

// GetAvg returns the Avg field value if set, zero value otherwise.
func (o *FunnelResponseElapsedTime) GetAvg() int64 {
	if o == nil || o.Avg == nil {
		var ret int64
		return ret
	}
	return *o.Avg
}

// GetAvgOk returns a tuple with the Avg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelResponseElapsedTime) GetAvgOk() (*int64, bool) {
	if o == nil || o.Avg == nil {
		return nil, false
	}
	return o.Avg, true
}

// HasAvg returns a boolean if a field has been set.
func (o *FunnelResponseElapsedTime) HasAvg() bool {
	return o != nil && o.Avg != nil
}

// SetAvg gets a reference to the given int64 and assigns it to the Avg field.
func (o *FunnelResponseElapsedTime) SetAvg(v int64) {
	o.Avg = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *FunnelResponseElapsedTime) GetMax() int64 {
	if o == nil || o.Max == nil {
		var ret int64
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelResponseElapsedTime) GetMaxOk() (*int64, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *FunnelResponseElapsedTime) HasMax() bool {
	return o != nil && o.Max != nil
}

// SetMax gets a reference to the given int64 and assigns it to the Max field.
func (o *FunnelResponseElapsedTime) SetMax(v int64) {
	o.Max = &v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *FunnelResponseElapsedTime) GetMin() int64 {
	if o == nil || o.Min == nil {
		var ret int64
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelResponseElapsedTime) GetMinOk() (*int64, bool) {
	if o == nil || o.Min == nil {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *FunnelResponseElapsedTime) HasMin() bool {
	return o != nil && o.Min != nil
}

// SetMin gets a reference to the given int64 and assigns it to the Min field.
func (o *FunnelResponseElapsedTime) SetMin(v int64) {
	o.Min = &v
}

// GetP5 returns the P5 field value if set, zero value otherwise.
func (o *FunnelResponseElapsedTime) GetP5() int64 {
	if o == nil || o.P5 == nil {
		var ret int64
		return ret
	}
	return *o.P5
}

// GetP5Ok returns a tuple with the P5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelResponseElapsedTime) GetP5Ok() (*int64, bool) {
	if o == nil || o.P5 == nil {
		return nil, false
	}
	return o.P5, true
}

// HasP5 returns a boolean if a field has been set.
func (o *FunnelResponseElapsedTime) HasP5() bool {
	return o != nil && o.P5 != nil
}

// SetP5 gets a reference to the given int64 and assigns it to the P5 field.
func (o *FunnelResponseElapsedTime) SetP5(v int64) {
	o.P5 = &v
}

// GetP50 returns the P50 field value if set, zero value otherwise.
func (o *FunnelResponseElapsedTime) GetP50() int64 {
	if o == nil || o.P50 == nil {
		var ret int64
		return ret
	}
	return *o.P50
}

// GetP50Ok returns a tuple with the P50 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelResponseElapsedTime) GetP50Ok() (*int64, bool) {
	if o == nil || o.P50 == nil {
		return nil, false
	}
	return o.P50, true
}

// HasP50 returns a boolean if a field has been set.
func (o *FunnelResponseElapsedTime) HasP50() bool {
	return o != nil && o.P50 != nil
}

// SetP50 gets a reference to the given int64 and assigns it to the P50 field.
func (o *FunnelResponseElapsedTime) SetP50(v int64) {
	o.P50 = &v
}

// GetP95 returns the P95 field value if set, zero value otherwise.
func (o *FunnelResponseElapsedTime) GetP95() int64 {
	if o == nil || o.P95 == nil {
		var ret int64
		return ret
	}
	return *o.P95
}

// GetP95Ok returns a tuple with the P95 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelResponseElapsedTime) GetP95Ok() (*int64, bool) {
	if o == nil || o.P95 == nil {
		return nil, false
	}
	return o.P95, true
}

// HasP95 returns a boolean if a field has been set.
func (o *FunnelResponseElapsedTime) HasP95() bool {
	return o != nil && o.P95 != nil
}

// SetP95 gets a reference to the given int64 and assigns it to the P95 field.
func (o *FunnelResponseElapsedTime) SetP95(v int64) {
	o.P95 = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FunnelResponseElapsedTime) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Avg != nil {
		toSerialize["avg"] = o.Avg
	}
	if o.Max != nil {
		toSerialize["max"] = o.Max
	}
	if o.Min != nil {
		toSerialize["min"] = o.Min
	}
	if o.P5 != nil {
		toSerialize["p5"] = o.P5
	}
	if o.P50 != nil {
		toSerialize["p50"] = o.P50
	}
	if o.P95 != nil {
		toSerialize["p95"] = o.P95
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FunnelResponseElapsedTime) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Avg *int64 `json:"avg,omitempty"`
		Max *int64 `json:"max,omitempty"`
		Min *int64 `json:"min,omitempty"`
		P5  *int64 `json:"p5,omitempty"`
		P50 *int64 `json:"p50,omitempty"`
		P95 *int64 `json:"p95,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"avg", "max", "min", "p5", "p50", "p95"})
	} else {
		return err
	}
	o.Avg = all.Avg
	o.Max = all.Max
	o.Min = all.Min
	o.P5 = all.P5
	o.P50 = all.P50
	o.P95 = all.P95

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
