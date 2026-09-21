// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MatchingSignalAttributes Attributes of a matching security signal.
type MatchingSignalAttributes struct {
	// The tracker ID linking the signal back to the originating event. Distinct from `id`, which identifies the matching signal itself.
	EventTrackerId string `json:"event_tracker_id"`
	// The severity of the signal.
	Severity string `json:"severity"`
	// The title of the signal.
	Title string `json:"title"`
	// The Unix timestamp (in milliseconds) at which the signal was triggered.
	TriggerTimeMs int64 `json:"trigger_time_ms"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMatchingSignalAttributes instantiates a new MatchingSignalAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMatchingSignalAttributes(eventTrackerId string, severity string, title string, triggerTimeMs int64) *MatchingSignalAttributes {
	this := MatchingSignalAttributes{}
	this.EventTrackerId = eventTrackerId
	this.Severity = severity
	this.Title = title
	this.TriggerTimeMs = triggerTimeMs
	return &this
}

// NewMatchingSignalAttributesWithDefaults instantiates a new MatchingSignalAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMatchingSignalAttributesWithDefaults() *MatchingSignalAttributes {
	this := MatchingSignalAttributes{}
	return &this
}

// GetEventTrackerId returns the EventTrackerId field value.
func (o *MatchingSignalAttributes) GetEventTrackerId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EventTrackerId
}

// GetEventTrackerIdOk returns a tuple with the EventTrackerId field value
// and a boolean to check if the value has been set.
func (o *MatchingSignalAttributes) GetEventTrackerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventTrackerId, true
}

// SetEventTrackerId sets field value.
func (o *MatchingSignalAttributes) SetEventTrackerId(v string) {
	o.EventTrackerId = v
}

// GetSeverity returns the Severity field value.
func (o *MatchingSignalAttributes) GetSeverity() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *MatchingSignalAttributes) GetSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value.
func (o *MatchingSignalAttributes) SetSeverity(v string) {
	o.Severity = v
}

// GetTitle returns the Title field value.
func (o *MatchingSignalAttributes) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *MatchingSignalAttributes) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *MatchingSignalAttributes) SetTitle(v string) {
	o.Title = v
}

// GetTriggerTimeMs returns the TriggerTimeMs field value.
func (o *MatchingSignalAttributes) GetTriggerTimeMs() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TriggerTimeMs
}

// GetTriggerTimeMsOk returns a tuple with the TriggerTimeMs field value
// and a boolean to check if the value has been set.
func (o *MatchingSignalAttributes) GetTriggerTimeMsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TriggerTimeMs, true
}

// SetTriggerTimeMs sets field value.
func (o *MatchingSignalAttributes) SetTriggerTimeMs(v int64) {
	o.TriggerTimeMs = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MatchingSignalAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["event_tracker_id"] = o.EventTrackerId
	toSerialize["severity"] = o.Severity
	toSerialize["title"] = o.Title
	toSerialize["trigger_time_ms"] = o.TriggerTimeMs

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MatchingSignalAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EventTrackerId *string `json:"event_tracker_id"`
		Severity       *string `json:"severity"`
		Title          *string `json:"title"`
		TriggerTimeMs  *int64  `json:"trigger_time_ms"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.EventTrackerId == nil {
		return fmt.Errorf("required field event_tracker_id missing")
	}
	if all.Severity == nil {
		return fmt.Errorf("required field severity missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	if all.TriggerTimeMs == nil {
		return fmt.Errorf("required field trigger_time_ms missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"event_tracker_id", "severity", "title", "trigger_time_ms"})
	} else {
		return err
	}
	o.EventTrackerId = *all.EventTrackerId
	o.Severity = *all.Severity
	o.Title = *all.Title
	o.TriggerTimeMs = *all.TriggerTimeMs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
