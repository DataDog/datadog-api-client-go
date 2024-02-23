// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorFreshness The freshness of the monitor, indicating if the monitor evaluation is up to date. **This feature is currently in private beta.**
type MonitorFreshness struct {
	// The error message if the monitor freshness was not able to be determined.
	Error datadog.NullableString `json:"error,omitempty"`
	// Whether the monitor evaluation is up to date.
	IsFresh *bool `json:"is_fresh,omitempty"`
	// The timestamp of the last evaluation.
	LastEvaluated *time.Time `json:"last_evaluated,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewMonitorFreshness instantiates a new MonitorFreshness object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorFreshness() *MonitorFreshness {
	this := MonitorFreshness{}
	return &this
}

// NewMonitorFreshnessWithDefaults instantiates a new MonitorFreshness object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorFreshnessWithDefaults() *MonitorFreshness {
	this := MonitorFreshness{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonitorFreshness) GetError() string {
	if o == nil || o.Error.Get() == nil {
		var ret string
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonitorFreshness) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *MonitorFreshness) HasError() bool {
	return o != nil && o.Error.IsSet()
}

// SetError gets a reference to the given datadog.NullableString and assigns it to the Error field.
func (o *MonitorFreshness) SetError(v string) {
	o.Error.Set(&v)
}

// SetErrorNil sets the value for Error to be an explicit nil.
func (o *MonitorFreshness) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil.
func (o *MonitorFreshness) UnsetError() {
	o.Error.Unset()
}

// GetIsFresh returns the IsFresh field value if set, zero value otherwise.
func (o *MonitorFreshness) GetIsFresh() bool {
	if o == nil || o.IsFresh == nil {
		var ret bool
		return ret
	}
	return *o.IsFresh
}

// GetIsFreshOk returns a tuple with the IsFresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorFreshness) GetIsFreshOk() (*bool, bool) {
	if o == nil || o.IsFresh == nil {
		return nil, false
	}
	return o.IsFresh, true
}

// HasIsFresh returns a boolean if a field has been set.
func (o *MonitorFreshness) HasIsFresh() bool {
	return o != nil && o.IsFresh != nil
}

// SetIsFresh gets a reference to the given bool and assigns it to the IsFresh field.
func (o *MonitorFreshness) SetIsFresh(v bool) {
	o.IsFresh = &v
}

// GetLastEvaluated returns the LastEvaluated field value if set, zero value otherwise.
func (o *MonitorFreshness) GetLastEvaluated() time.Time {
	if o == nil || o.LastEvaluated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastEvaluated
}

// GetLastEvaluatedOk returns a tuple with the LastEvaluated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorFreshness) GetLastEvaluatedOk() (*time.Time, bool) {
	if o == nil || o.LastEvaluated == nil {
		return nil, false
	}
	return o.LastEvaluated, true
}

// HasLastEvaluated returns a boolean if a field has been set.
func (o *MonitorFreshness) HasLastEvaluated() bool {
	return o != nil && o.LastEvaluated != nil
}

// SetLastEvaluated gets a reference to the given time.Time and assigns it to the LastEvaluated field.
func (o *MonitorFreshness) SetLastEvaluated(v time.Time) {
	o.LastEvaluated = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorFreshness) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	if o.IsFresh != nil {
		toSerialize["is_fresh"] = o.IsFresh
	}
	if o.LastEvaluated != nil {
		if o.LastEvaluated.Nanosecond() == 0 {
			toSerialize["last_evaluated"] = o.LastEvaluated.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["last_evaluated"] = o.LastEvaluated.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorFreshness) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Error         datadog.NullableString `json:"error,omitempty"`
		IsFresh       *bool                  `json:"is_fresh,omitempty"`
		LastEvaluated *time.Time             `json:"last_evaluated,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"error", "is_fresh", "last_evaluated"})
	} else {
		return err
	}
	o.Error = all.Error
	o.IsFresh = all.IsFresh
	o.LastEvaluated = all.LastEvaluated

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
