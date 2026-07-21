// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DdsqlTabularQueryTimeWindow Time window scoping the underlying data sources, expressed in Unix milliseconds
// since the epoch. Inclusive on `from_timestamp`, exclusive on `to_timestamp`.
// Results from static tables (for example, `dd.hosts`) are not affected by the
// time window, but the field must still be provided.
type DdsqlTabularQueryTimeWindow struct {
	// Start of the query window (inclusive), in Unix milliseconds since the epoch.
	FromTimestamp int64 `json:"from_timestamp"`
	// End of the query window (exclusive), in Unix milliseconds since the epoch.
	ToTimestamp int64 `json:"to_timestamp"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDdsqlTabularQueryTimeWindow instantiates a new DdsqlTabularQueryTimeWindow object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDdsqlTabularQueryTimeWindow(fromTimestamp int64, toTimestamp int64) *DdsqlTabularQueryTimeWindow {
	this := DdsqlTabularQueryTimeWindow{}
	this.FromTimestamp = fromTimestamp
	this.ToTimestamp = toTimestamp
	return &this
}

// NewDdsqlTabularQueryTimeWindowWithDefaults instantiates a new DdsqlTabularQueryTimeWindow object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDdsqlTabularQueryTimeWindowWithDefaults() *DdsqlTabularQueryTimeWindow {
	this := DdsqlTabularQueryTimeWindow{}
	return &this
}

// GetFromTimestamp returns the FromTimestamp field value.
func (o *DdsqlTabularQueryTimeWindow) GetFromTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.FromTimestamp
}

// GetFromTimestampOk returns a tuple with the FromTimestamp field value
// and a boolean to check if the value has been set.
func (o *DdsqlTabularQueryTimeWindow) GetFromTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromTimestamp, true
}

// SetFromTimestamp sets field value.
func (o *DdsqlTabularQueryTimeWindow) SetFromTimestamp(v int64) {
	o.FromTimestamp = v
}

// GetToTimestamp returns the ToTimestamp field value.
func (o *DdsqlTabularQueryTimeWindow) GetToTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ToTimestamp
}

// GetToTimestampOk returns a tuple with the ToTimestamp field value
// and a boolean to check if the value has been set.
func (o *DdsqlTabularQueryTimeWindow) GetToTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToTimestamp, true
}

// SetToTimestamp sets field value.
func (o *DdsqlTabularQueryTimeWindow) SetToTimestamp(v int64) {
	o.ToTimestamp = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DdsqlTabularQueryTimeWindow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["from_timestamp"] = o.FromTimestamp
	toSerialize["to_timestamp"] = o.ToTimestamp

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DdsqlTabularQueryTimeWindow) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FromTimestamp *int64 `json:"from_timestamp"`
		ToTimestamp   *int64 `json:"to_timestamp"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.FromTimestamp == nil {
		return fmt.Errorf("required field from_timestamp missing")
	}
	if all.ToTimestamp == nil {
		return fmt.Errorf("required field to_timestamp missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"from_timestamp", "to_timestamp"})
	} else {
		return err
	}
	o.FromTimestamp = *all.FromTimestamp
	o.ToTimestamp = *all.ToTimestamp

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
