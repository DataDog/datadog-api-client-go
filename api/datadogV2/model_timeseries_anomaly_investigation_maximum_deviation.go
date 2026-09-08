// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesAnomalyInvestigationMaximumDeviation Most anomalous point within the detected interval.
type TimeseriesAnomalyInvestigationMaximumDeviation struct {
	// Absolute distance between the observed value and the nearest anomaly boundary.
	DeltaFromBoundary float64 `json:"delta_from_boundary"`
	// Point timestamp in milliseconds since the Unix epoch.
	Timestamp int64 `json:"timestamp"`
	// Observed value at the point.
	Value float64 `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTimeseriesAnomalyInvestigationMaximumDeviation instantiates a new TimeseriesAnomalyInvestigationMaximumDeviation object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTimeseriesAnomalyInvestigationMaximumDeviation(deltaFromBoundary float64, timestamp int64, value float64) *TimeseriesAnomalyInvestigationMaximumDeviation {
	this := TimeseriesAnomalyInvestigationMaximumDeviation{}
	this.DeltaFromBoundary = deltaFromBoundary
	this.Timestamp = timestamp
	this.Value = value
	return &this
}

// NewTimeseriesAnomalyInvestigationMaximumDeviationWithDefaults instantiates a new TimeseriesAnomalyInvestigationMaximumDeviation object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTimeseriesAnomalyInvestigationMaximumDeviationWithDefaults() *TimeseriesAnomalyInvestigationMaximumDeviation {
	this := TimeseriesAnomalyInvestigationMaximumDeviation{}
	return &this
}

// GetDeltaFromBoundary returns the DeltaFromBoundary field value.
func (o *TimeseriesAnomalyInvestigationMaximumDeviation) GetDeltaFromBoundary() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.DeltaFromBoundary
}

// GetDeltaFromBoundaryOk returns a tuple with the DeltaFromBoundary field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationMaximumDeviation) GetDeltaFromBoundaryOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeltaFromBoundary, true
}

// SetDeltaFromBoundary sets field value.
func (o *TimeseriesAnomalyInvestigationMaximumDeviation) SetDeltaFromBoundary(v float64) {
	o.DeltaFromBoundary = v
}

// GetTimestamp returns the Timestamp field value.
func (o *TimeseriesAnomalyInvestigationMaximumDeviation) GetTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationMaximumDeviation) GetTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value.
func (o *TimeseriesAnomalyInvestigationMaximumDeviation) SetTimestamp(v int64) {
	o.Timestamp = v
}

// GetValue returns the Value field value.
func (o *TimeseriesAnomalyInvestigationMaximumDeviation) GetValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationMaximumDeviation) GetValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *TimeseriesAnomalyInvestigationMaximumDeviation) SetValue(v float64) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TimeseriesAnomalyInvestigationMaximumDeviation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["delta_from_boundary"] = o.DeltaFromBoundary
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TimeseriesAnomalyInvestigationMaximumDeviation) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DeltaFromBoundary *float64 `json:"delta_from_boundary"`
		Timestamp         *int64   `json:"timestamp"`
		Value             *float64 `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DeltaFromBoundary == nil {
		return fmt.Errorf("required field delta_from_boundary missing")
	}
	if all.Timestamp == nil {
		return fmt.Errorf("required field timestamp missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"delta_from_boundary", "timestamp", "value"})
	} else {
		return err
	}
	o.DeltaFromBoundary = *all.DeltaFromBoundary
	o.Timestamp = *all.Timestamp
	o.Value = *all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
