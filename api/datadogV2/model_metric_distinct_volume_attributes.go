// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricDistinctVolumeAttributes Object containing the definition of a metric's distinct volume.
type MetricDistinctVolumeAttributes struct {
	// Distinct volume for the given metric. This value is `null` for standard metrics in organizations on
	// Metric Name Pricing.
	DistinctVolume datadog.NullableInt64 `json:"distinct_volume,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMetricDistinctVolumeAttributes instantiates a new MetricDistinctVolumeAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMetricDistinctVolumeAttributes() *MetricDistinctVolumeAttributes {
	this := MetricDistinctVolumeAttributes{}
	return &this
}

// NewMetricDistinctVolumeAttributesWithDefaults instantiates a new MetricDistinctVolumeAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMetricDistinctVolumeAttributesWithDefaults() *MetricDistinctVolumeAttributes {
	this := MetricDistinctVolumeAttributes{}
	return &this
}

// GetDistinctVolume returns the DistinctVolume field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetricDistinctVolumeAttributes) GetDistinctVolume() int64 {
	if o == nil || o.DistinctVolume.Get() == nil {
		var ret int64
		return ret
	}
	return *o.DistinctVolume.Get()
}

// GetDistinctVolumeOk returns a tuple with the DistinctVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MetricDistinctVolumeAttributes) GetDistinctVolumeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DistinctVolume.Get(), o.DistinctVolume.IsSet()
}

// HasDistinctVolume returns a boolean if a field has been set.
func (o *MetricDistinctVolumeAttributes) HasDistinctVolume() bool {
	return o != nil && o.DistinctVolume.IsSet()
}

// SetDistinctVolume gets a reference to the given datadog.NullableInt64 and assigns it to the DistinctVolume field.
func (o *MetricDistinctVolumeAttributes) SetDistinctVolume(v int64) {
	o.DistinctVolume.Set(&v)
}

// SetDistinctVolumeNil sets the value for DistinctVolume to be an explicit nil.
func (o *MetricDistinctVolumeAttributes) SetDistinctVolumeNil() {
	o.DistinctVolume.Set(nil)
}

// UnsetDistinctVolume ensures that no value is present for DistinctVolume, not even an explicit nil.
func (o *MetricDistinctVolumeAttributes) UnsetDistinctVolume() {
	o.DistinctVolume.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o MetricDistinctVolumeAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DistinctVolume.IsSet() {
		toSerialize["distinct_volume"] = o.DistinctVolume.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MetricDistinctVolumeAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DistinctVolume datadog.NullableInt64 `json:"distinct_volume,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"distinct_volume"})
	} else {
		return err
	}
	o.DistinctVolume = all.DistinctVolume

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
