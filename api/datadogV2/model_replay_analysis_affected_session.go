// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReplayAnalysisAffectedSession A session affected by a replay analysis issue.
type ReplayAnalysisAffectedSession struct {
	// Unique identifier of the affected session.
	SessionId string `json:"session_id"`
	// Session start timestamp in milliseconds.
	SessionStartTimestampMs datadog.NullableInt64 `json:"session_start_timestamp_ms,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewReplayAnalysisAffectedSession instantiates a new ReplayAnalysisAffectedSession object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReplayAnalysisAffectedSession(sessionId string) *ReplayAnalysisAffectedSession {
	this := ReplayAnalysisAffectedSession{}
	this.SessionId = sessionId
	return &this
}

// NewReplayAnalysisAffectedSessionWithDefaults instantiates a new ReplayAnalysisAffectedSession object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReplayAnalysisAffectedSessionWithDefaults() *ReplayAnalysisAffectedSession {
	this := ReplayAnalysisAffectedSession{}
	return &this
}

// GetSessionId returns the SessionId field value.
func (o *ReplayAnalysisAffectedSession) GetSessionId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value
// and a boolean to check if the value has been set.
func (o *ReplayAnalysisAffectedSession) GetSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionId, true
}

// SetSessionId sets field value.
func (o *ReplayAnalysisAffectedSession) SetSessionId(v string) {
	o.SessionId = v
}

// GetSessionStartTimestampMs returns the SessionStartTimestampMs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReplayAnalysisAffectedSession) GetSessionStartTimestampMs() int64 {
	if o == nil || o.SessionStartTimestampMs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.SessionStartTimestampMs.Get()
}

// GetSessionStartTimestampMsOk returns a tuple with the SessionStartTimestampMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReplayAnalysisAffectedSession) GetSessionStartTimestampMsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SessionStartTimestampMs.Get(), o.SessionStartTimestampMs.IsSet()
}

// HasSessionStartTimestampMs returns a boolean if a field has been set.
func (o *ReplayAnalysisAffectedSession) HasSessionStartTimestampMs() bool {
	return o != nil && o.SessionStartTimestampMs.IsSet()
}

// SetSessionStartTimestampMs gets a reference to the given datadog.NullableInt64 and assigns it to the SessionStartTimestampMs field.
func (o *ReplayAnalysisAffectedSession) SetSessionStartTimestampMs(v int64) {
	o.SessionStartTimestampMs.Set(&v)
}

// SetSessionStartTimestampMsNil sets the value for SessionStartTimestampMs to be an explicit nil.
func (o *ReplayAnalysisAffectedSession) SetSessionStartTimestampMsNil() {
	o.SessionStartTimestampMs.Set(nil)
}

// UnsetSessionStartTimestampMs ensures that no value is present for SessionStartTimestampMs, not even an explicit nil.
func (o *ReplayAnalysisAffectedSession) UnsetSessionStartTimestampMs() {
	o.SessionStartTimestampMs.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ReplayAnalysisAffectedSession) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["session_id"] = o.SessionId
	if o.SessionStartTimestampMs.IsSet() {
		toSerialize["session_start_timestamp_ms"] = o.SessionStartTimestampMs.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ReplayAnalysisAffectedSession) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SessionId               *string               `json:"session_id"`
		SessionStartTimestampMs datadog.NullableInt64 `json:"session_start_timestamp_ms,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.SessionId == nil {
		return fmt.Errorf("required field session_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"session_id", "session_start_timestamp_ms"})
	} else {
		return err
	}
	o.SessionId = *all.SessionId
	o.SessionStartTimestampMs = all.SessionStartTimestampMs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
