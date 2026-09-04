// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesAnomalyInvestigationInterval Half-open time interval in milliseconds since the Unix epoch.
type TimeseriesAnomalyInvestigationInterval struct {
	// Exclusive end of the interval.
	End int64 `json:"end"`
	// Inclusive start of the interval.
	Start int64 `json:"start"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTimeseriesAnomalyInvestigationInterval instantiates a new TimeseriesAnomalyInvestigationInterval object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTimeseriesAnomalyInvestigationInterval(end int64, start int64) *TimeseriesAnomalyInvestigationInterval {
	this := TimeseriesAnomalyInvestigationInterval{}
	this.End = end
	this.Start = start
	return &this
}

// NewTimeseriesAnomalyInvestigationIntervalWithDefaults instantiates a new TimeseriesAnomalyInvestigationInterval object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTimeseriesAnomalyInvestigationIntervalWithDefaults() *TimeseriesAnomalyInvestigationInterval {
	this := TimeseriesAnomalyInvestigationInterval{}
	return &this
}

// GetEnd returns the End field value.
func (o *TimeseriesAnomalyInvestigationInterval) GetEnd() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationInterval) GetEndOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value.
func (o *TimeseriesAnomalyInvestigationInterval) SetEnd(v int64) {
	o.End = v
}

// GetStart returns the Start field value.
func (o *TimeseriesAnomalyInvestigationInterval) GetStart() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationInterval) GetStartOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value.
func (o *TimeseriesAnomalyInvestigationInterval) SetStart(v int64) {
	o.Start = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TimeseriesAnomalyInvestigationInterval) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["end"] = o.End
	toSerialize["start"] = o.Start

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TimeseriesAnomalyInvestigationInterval) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		End   *int64 `json:"end"`
		Start *int64 `json:"start"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.End == nil {
		return fmt.Errorf("required field end missing")
	}
	if all.Start == nil {
		return fmt.Errorf("required field start missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"end", "start"})
	} else {
		return err
	}
	o.End = *all.End
	o.Start = *all.Start

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
