// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReplaySummaryChapter A chapter within a RUM replay summary, representing a distinct segment of user activity.
type ReplaySummaryChapter struct {
	// End time of the chapter in milliseconds.
	EndMs int64 `json:"end_ms"`
	// Start time of the chapter in milliseconds.
	StartMs int64 `json:"start_ms"`
	// Description of user activity during this chapter.
	Text string `json:"text"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewReplaySummaryChapter instantiates a new ReplaySummaryChapter object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReplaySummaryChapter(endMs int64, startMs int64, text string) *ReplaySummaryChapter {
	this := ReplaySummaryChapter{}
	this.EndMs = endMs
	this.StartMs = startMs
	this.Text = text
	return &this
}

// NewReplaySummaryChapterWithDefaults instantiates a new ReplaySummaryChapter object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReplaySummaryChapterWithDefaults() *ReplaySummaryChapter {
	this := ReplaySummaryChapter{}
	return &this
}

// GetEndMs returns the EndMs field value.
func (o *ReplaySummaryChapter) GetEndMs() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.EndMs
}

// GetEndMsOk returns a tuple with the EndMs field value
// and a boolean to check if the value has been set.
func (o *ReplaySummaryChapter) GetEndMsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndMs, true
}

// SetEndMs sets field value.
func (o *ReplaySummaryChapter) SetEndMs(v int64) {
	o.EndMs = v
}

// GetStartMs returns the StartMs field value.
func (o *ReplaySummaryChapter) GetStartMs() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.StartMs
}

// GetStartMsOk returns a tuple with the StartMs field value
// and a boolean to check if the value has been set.
func (o *ReplaySummaryChapter) GetStartMsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartMs, true
}

// SetStartMs sets field value.
func (o *ReplaySummaryChapter) SetStartMs(v int64) {
	o.StartMs = v
}

// GetText returns the Text field value.
func (o *ReplaySummaryChapter) GetText() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *ReplaySummaryChapter) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value.
func (o *ReplaySummaryChapter) SetText(v string) {
	o.Text = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ReplaySummaryChapter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["end_ms"] = o.EndMs
	toSerialize["start_ms"] = o.StartMs
	toSerialize["text"] = o.Text

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ReplaySummaryChapter) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EndMs   *int64  `json:"end_ms"`
		StartMs *int64  `json:"start_ms"`
		Text    *string `json:"text"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.EndMs == nil {
		return fmt.Errorf("required field end_ms missing")
	}
	if all.StartMs == nil {
		return fmt.Errorf("required field start_ms missing")
	}
	if all.Text == nil {
		return fmt.Errorf("required field text missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"end_ms", "start_ms", "text"})
	} else {
		return err
	}
	o.EndMs = *all.EndMs
	o.StartMs = *all.StartMs
	o.Text = *all.Text

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
