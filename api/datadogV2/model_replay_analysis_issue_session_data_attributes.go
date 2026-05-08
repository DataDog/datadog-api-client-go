// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReplayAnalysisIssueSessionDataAttributes Attributes of a session related to a RUM replay analysis issue.
type ReplayAnalysisIssueSessionDataAttributes struct {
	// Session start timestamp in milliseconds.
	SessionStartTimestampMs int64 `json:"session_start_timestamp_ms"`
	// List of signals observed in this session.
	Signals []ReplayAnalysisSignal `json:"signals"`
	// Name of the view where the issue was observed.
	ViewName string `json:"view_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewReplayAnalysisIssueSessionDataAttributes instantiates a new ReplayAnalysisIssueSessionDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReplayAnalysisIssueSessionDataAttributes(sessionStartTimestampMs int64, signals []ReplayAnalysisSignal, viewName string) *ReplayAnalysisIssueSessionDataAttributes {
	this := ReplayAnalysisIssueSessionDataAttributes{}
	this.SessionStartTimestampMs = sessionStartTimestampMs
	this.Signals = signals
	this.ViewName = viewName
	return &this
}

// NewReplayAnalysisIssueSessionDataAttributesWithDefaults instantiates a new ReplayAnalysisIssueSessionDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReplayAnalysisIssueSessionDataAttributesWithDefaults() *ReplayAnalysisIssueSessionDataAttributes {
	this := ReplayAnalysisIssueSessionDataAttributes{}
	return &this
}

// GetSessionStartTimestampMs returns the SessionStartTimestampMs field value.
func (o *ReplayAnalysisIssueSessionDataAttributes) GetSessionStartTimestampMs() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.SessionStartTimestampMs
}

// GetSessionStartTimestampMsOk returns a tuple with the SessionStartTimestampMs field value
// and a boolean to check if the value has been set.
func (o *ReplayAnalysisIssueSessionDataAttributes) GetSessionStartTimestampMsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionStartTimestampMs, true
}

// SetSessionStartTimestampMs sets field value.
func (o *ReplayAnalysisIssueSessionDataAttributes) SetSessionStartTimestampMs(v int64) {
	o.SessionStartTimestampMs = v
}

// GetSignals returns the Signals field value.
func (o *ReplayAnalysisIssueSessionDataAttributes) GetSignals() []ReplayAnalysisSignal {
	if o == nil {
		var ret []ReplayAnalysisSignal
		return ret
	}
	return o.Signals
}

// GetSignalsOk returns a tuple with the Signals field value
// and a boolean to check if the value has been set.
func (o *ReplayAnalysisIssueSessionDataAttributes) GetSignalsOk() (*[]ReplayAnalysisSignal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signals, true
}

// SetSignals sets field value.
func (o *ReplayAnalysisIssueSessionDataAttributes) SetSignals(v []ReplayAnalysisSignal) {
	o.Signals = v
}

// GetViewName returns the ViewName field value.
func (o *ReplayAnalysisIssueSessionDataAttributes) GetViewName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ViewName
}

// GetViewNameOk returns a tuple with the ViewName field value
// and a boolean to check if the value has been set.
func (o *ReplayAnalysisIssueSessionDataAttributes) GetViewNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ViewName, true
}

// SetViewName sets field value.
func (o *ReplayAnalysisIssueSessionDataAttributes) SetViewName(v string) {
	o.ViewName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ReplayAnalysisIssueSessionDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["session_start_timestamp_ms"] = o.SessionStartTimestampMs
	toSerialize["signals"] = o.Signals
	toSerialize["view_name"] = o.ViewName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ReplayAnalysisIssueSessionDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SessionStartTimestampMs *int64                  `json:"session_start_timestamp_ms"`
		Signals                 *[]ReplayAnalysisSignal `json:"signals"`
		ViewName                *string                 `json:"view_name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.SessionStartTimestampMs == nil {
		return fmt.Errorf("required field session_start_timestamp_ms missing")
	}
	if all.Signals == nil {
		return fmt.Errorf("required field signals missing")
	}
	if all.ViewName == nil {
		return fmt.Errorf("required field view_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"session_start_timestamp_ms", "signals", "view_name"})
	} else {
		return err
	}
	o.SessionStartTimestampMs = *all.SessionStartTimestampMs
	o.Signals = *all.Signals
	o.ViewName = *all.ViewName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
