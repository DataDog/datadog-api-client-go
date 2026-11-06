// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricIngestedIndexedVolumeAttributes Object containing the definition of a metric's ingested and indexed volume.
type MetricIngestedIndexedVolumeAttributes struct {
	// Estimated average hourly number of indexed time series for the given metric over the last hour. For organizations on Metric Name Pricing, this represents the estimated sum of indexed data points over the last hour.
	IndexedVolume *int64 `json:"indexed_volume,omitempty"`
	// Change in `indexed_volume` compared with the previous hour: the last hour's value minus the value for
	// the hour before. Negative when volume decreased.
	IndexedVolumeDelta *int64 `json:"indexed_volume_delta,omitempty"`
	// Estimated average hourly number of ingested time series for the given metric over the last hour. This value is `0` for metrics not configured with Metrics Without Limits. For organizations on Metric Name Pricing, this represents the estimated sum of ingested data points over the last hour.
	IngestedVolume *int64 `json:"ingested_volume,omitempty"`
	// Change in `ingested_volume` compared with the previous hour: the last hour's value minus the value for
	// the hour before. Negative when volume decreased.
	IngestedVolumeDelta *int64 `json:"ingested_volume_delta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMetricIngestedIndexedVolumeAttributes instantiates a new MetricIngestedIndexedVolumeAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMetricIngestedIndexedVolumeAttributes() *MetricIngestedIndexedVolumeAttributes {
	this := MetricIngestedIndexedVolumeAttributes{}
	return &this
}

// NewMetricIngestedIndexedVolumeAttributesWithDefaults instantiates a new MetricIngestedIndexedVolumeAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMetricIngestedIndexedVolumeAttributesWithDefaults() *MetricIngestedIndexedVolumeAttributes {
	this := MetricIngestedIndexedVolumeAttributes{}
	return &this
}

// GetIndexedVolume returns the IndexedVolume field value if set, zero value otherwise.
func (o *MetricIngestedIndexedVolumeAttributes) GetIndexedVolume() int64 {
	if o == nil || o.IndexedVolume == nil {
		var ret int64
		return ret
	}
	return *o.IndexedVolume
}

// GetIndexedVolumeOk returns a tuple with the IndexedVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricIngestedIndexedVolumeAttributes) GetIndexedVolumeOk() (*int64, bool) {
	if o == nil || o.IndexedVolume == nil {
		return nil, false
	}
	return o.IndexedVolume, true
}

// HasIndexedVolume returns a boolean if a field has been set.
func (o *MetricIngestedIndexedVolumeAttributes) HasIndexedVolume() bool {
	return o != nil && o.IndexedVolume != nil
}

// SetIndexedVolume gets a reference to the given int64 and assigns it to the IndexedVolume field.
func (o *MetricIngestedIndexedVolumeAttributes) SetIndexedVolume(v int64) {
	o.IndexedVolume = &v
}

// GetIndexedVolumeDelta returns the IndexedVolumeDelta field value if set, zero value otherwise.
func (o *MetricIngestedIndexedVolumeAttributes) GetIndexedVolumeDelta() int64 {
	if o == nil || o.IndexedVolumeDelta == nil {
		var ret int64
		return ret
	}
	return *o.IndexedVolumeDelta
}

// GetIndexedVolumeDeltaOk returns a tuple with the IndexedVolumeDelta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricIngestedIndexedVolumeAttributes) GetIndexedVolumeDeltaOk() (*int64, bool) {
	if o == nil || o.IndexedVolumeDelta == nil {
		return nil, false
	}
	return o.IndexedVolumeDelta, true
}

// HasIndexedVolumeDelta returns a boolean if a field has been set.
func (o *MetricIngestedIndexedVolumeAttributes) HasIndexedVolumeDelta() bool {
	return o != nil && o.IndexedVolumeDelta != nil
}

// SetIndexedVolumeDelta gets a reference to the given int64 and assigns it to the IndexedVolumeDelta field.
func (o *MetricIngestedIndexedVolumeAttributes) SetIndexedVolumeDelta(v int64) {
	o.IndexedVolumeDelta = &v
}

// GetIngestedVolume returns the IngestedVolume field value if set, zero value otherwise.
func (o *MetricIngestedIndexedVolumeAttributes) GetIngestedVolume() int64 {
	if o == nil || o.IngestedVolume == nil {
		var ret int64
		return ret
	}
	return *o.IngestedVolume
}

// GetIngestedVolumeOk returns a tuple with the IngestedVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricIngestedIndexedVolumeAttributes) GetIngestedVolumeOk() (*int64, bool) {
	if o == nil || o.IngestedVolume == nil {
		return nil, false
	}
	return o.IngestedVolume, true
}

// HasIngestedVolume returns a boolean if a field has been set.
func (o *MetricIngestedIndexedVolumeAttributes) HasIngestedVolume() bool {
	return o != nil && o.IngestedVolume != nil
}

// SetIngestedVolume gets a reference to the given int64 and assigns it to the IngestedVolume field.
func (o *MetricIngestedIndexedVolumeAttributes) SetIngestedVolume(v int64) {
	o.IngestedVolume = &v
}

// GetIngestedVolumeDelta returns the IngestedVolumeDelta field value if set, zero value otherwise.
func (o *MetricIngestedIndexedVolumeAttributes) GetIngestedVolumeDelta() int64 {
	if o == nil || o.IngestedVolumeDelta == nil {
		var ret int64
		return ret
	}
	return *o.IngestedVolumeDelta
}

// GetIngestedVolumeDeltaOk returns a tuple with the IngestedVolumeDelta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricIngestedIndexedVolumeAttributes) GetIngestedVolumeDeltaOk() (*int64, bool) {
	if o == nil || o.IngestedVolumeDelta == nil {
		return nil, false
	}
	return o.IngestedVolumeDelta, true
}

// HasIngestedVolumeDelta returns a boolean if a field has been set.
func (o *MetricIngestedIndexedVolumeAttributes) HasIngestedVolumeDelta() bool {
	return o != nil && o.IngestedVolumeDelta != nil
}

// SetIngestedVolumeDelta gets a reference to the given int64 and assigns it to the IngestedVolumeDelta field.
func (o *MetricIngestedIndexedVolumeAttributes) SetIngestedVolumeDelta(v int64) {
	o.IngestedVolumeDelta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MetricIngestedIndexedVolumeAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.IndexedVolume != nil {
		toSerialize["indexed_volume"] = o.IndexedVolume
	}
	if o.IndexedVolumeDelta != nil {
		toSerialize["indexed_volume_delta"] = o.IndexedVolumeDelta
	}
	if o.IngestedVolume != nil {
		toSerialize["ingested_volume"] = o.IngestedVolume
	}
	if o.IngestedVolumeDelta != nil {
		toSerialize["ingested_volume_delta"] = o.IngestedVolumeDelta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MetricIngestedIndexedVolumeAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IndexedVolume       *int64 `json:"indexed_volume,omitempty"`
		IndexedVolumeDelta  *int64 `json:"indexed_volume_delta,omitempty"`
		IngestedVolume      *int64 `json:"ingested_volume,omitempty"`
		IngestedVolumeDelta *int64 `json:"ingested_volume_delta,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"indexed_volume", "indexed_volume_delta", "ingested_volume", "ingested_volume_delta"})
	} else {
		return err
	}
	o.IndexedVolume = all.IndexedVolume
	o.IndexedVolumeDelta = all.IndexedVolumeDelta
	o.IngestedVolume = all.IngestedVolume
	o.IngestedVolumeDelta = all.IngestedVolumeDelta

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
