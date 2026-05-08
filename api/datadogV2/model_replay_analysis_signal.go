// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReplayAnalysisSignal A signal associated with a replay issue, capturing user interaction details.
type ReplayAnalysisSignal struct {
	// Event name associated with the signal.
	Event string `json:"event"`
	// Type of signal captured.
	SignalType string `json:"signal_type"`
	// Absolute timestamp of the signal in milliseconds.
	TimestampMs int64 `json:"timestamp_ms"`
	// User behavior pattern identified for the signal.
	UserPattern string `json:"user_pattern"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewReplayAnalysisSignal instantiates a new ReplayAnalysisSignal object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReplayAnalysisSignal(event string, signalType string, timestampMs int64, userPattern string) *ReplayAnalysisSignal {
	this := ReplayAnalysisSignal{}
	this.Event = event
	this.SignalType = signalType
	this.TimestampMs = timestampMs
	this.UserPattern = userPattern
	return &this
}

// NewReplayAnalysisSignalWithDefaults instantiates a new ReplayAnalysisSignal object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReplayAnalysisSignalWithDefaults() *ReplayAnalysisSignal {
	this := ReplayAnalysisSignal{}
	return &this
}

// GetEvent returns the Event field value.
func (o *ReplayAnalysisSignal) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *ReplayAnalysisSignal) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value.
func (o *ReplayAnalysisSignal) SetEvent(v string) {
	o.Event = v
}

// GetSignalType returns the SignalType field value.
func (o *ReplayAnalysisSignal) GetSignalType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SignalType
}

// GetSignalTypeOk returns a tuple with the SignalType field value
// and a boolean to check if the value has been set.
func (o *ReplayAnalysisSignal) GetSignalTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignalType, true
}

// SetSignalType sets field value.
func (o *ReplayAnalysisSignal) SetSignalType(v string) {
	o.SignalType = v
}

// GetTimestampMs returns the TimestampMs field value.
func (o *ReplayAnalysisSignal) GetTimestampMs() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TimestampMs
}

// GetTimestampMsOk returns a tuple with the TimestampMs field value
// and a boolean to check if the value has been set.
func (o *ReplayAnalysisSignal) GetTimestampMsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimestampMs, true
}

// SetTimestampMs sets field value.
func (o *ReplayAnalysisSignal) SetTimestampMs(v int64) {
	o.TimestampMs = v
}

// GetUserPattern returns the UserPattern field value.
func (o *ReplayAnalysisSignal) GetUserPattern() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UserPattern
}

// GetUserPatternOk returns a tuple with the UserPattern field value
// and a boolean to check if the value has been set.
func (o *ReplayAnalysisSignal) GetUserPatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserPattern, true
}

// SetUserPattern sets field value.
func (o *ReplayAnalysisSignal) SetUserPattern(v string) {
	o.UserPattern = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ReplayAnalysisSignal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["event"] = o.Event
	toSerialize["signal_type"] = o.SignalType
	toSerialize["timestamp_ms"] = o.TimestampMs
	toSerialize["user_pattern"] = o.UserPattern

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ReplayAnalysisSignal) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Event       *string `json:"event"`
		SignalType  *string `json:"signal_type"`
		TimestampMs *int64  `json:"timestamp_ms"`
		UserPattern *string `json:"user_pattern"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Event == nil {
		return fmt.Errorf("required field event missing")
	}
	if all.SignalType == nil {
		return fmt.Errorf("required field signal_type missing")
	}
	if all.TimestampMs == nil {
		return fmt.Errorf("required field timestamp_ms missing")
	}
	if all.UserPattern == nil {
		return fmt.Errorf("required field user_pattern missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"event", "signal_type", "timestamp_ms", "user_pattern"})
	} else {
		return err
	}
	o.Event = *all.Event
	o.SignalType = *all.SignalType
	o.TimestampMs = *all.TimestampMs
	o.UserPattern = *all.UserPattern

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
