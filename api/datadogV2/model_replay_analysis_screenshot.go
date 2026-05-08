// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReplayAnalysisScreenshot A screenshot captured during a replay session.
type ReplayAnalysisScreenshot struct {
	// Filename or key identifier of the screenshot.
	ScreenshotKey string `json:"screenshot_key"`
	// Unique identifier of the session where the screenshot was taken.
	SessionId string `json:"session_id"`
	// Timestamp of the screenshot in milliseconds.
	TimestampMs int64 `json:"timestamp_ms"`
	// Unique identifier of the view where the screenshot was taken.
	ViewId string `json:"view_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewReplayAnalysisScreenshot instantiates a new ReplayAnalysisScreenshot object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReplayAnalysisScreenshot(screenshotKey string, sessionId string, timestampMs int64, viewId string) *ReplayAnalysisScreenshot {
	this := ReplayAnalysisScreenshot{}
	this.ScreenshotKey = screenshotKey
	this.SessionId = sessionId
	this.TimestampMs = timestampMs
	this.ViewId = viewId
	return &this
}

// NewReplayAnalysisScreenshotWithDefaults instantiates a new ReplayAnalysisScreenshot object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReplayAnalysisScreenshotWithDefaults() *ReplayAnalysisScreenshot {
	this := ReplayAnalysisScreenshot{}
	return &this
}

// GetScreenshotKey returns the ScreenshotKey field value.
func (o *ReplayAnalysisScreenshot) GetScreenshotKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ScreenshotKey
}

// GetScreenshotKeyOk returns a tuple with the ScreenshotKey field value
// and a boolean to check if the value has been set.
func (o *ReplayAnalysisScreenshot) GetScreenshotKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScreenshotKey, true
}

// SetScreenshotKey sets field value.
func (o *ReplayAnalysisScreenshot) SetScreenshotKey(v string) {
	o.ScreenshotKey = v
}

// GetSessionId returns the SessionId field value.
func (o *ReplayAnalysisScreenshot) GetSessionId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value
// and a boolean to check if the value has been set.
func (o *ReplayAnalysisScreenshot) GetSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionId, true
}

// SetSessionId sets field value.
func (o *ReplayAnalysisScreenshot) SetSessionId(v string) {
	o.SessionId = v
}

// GetTimestampMs returns the TimestampMs field value.
func (o *ReplayAnalysisScreenshot) GetTimestampMs() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TimestampMs
}

// GetTimestampMsOk returns a tuple with the TimestampMs field value
// and a boolean to check if the value has been set.
func (o *ReplayAnalysisScreenshot) GetTimestampMsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimestampMs, true
}

// SetTimestampMs sets field value.
func (o *ReplayAnalysisScreenshot) SetTimestampMs(v int64) {
	o.TimestampMs = v
}

// GetViewId returns the ViewId field value.
func (o *ReplayAnalysisScreenshot) GetViewId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ViewId
}

// GetViewIdOk returns a tuple with the ViewId field value
// and a boolean to check if the value has been set.
func (o *ReplayAnalysisScreenshot) GetViewIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ViewId, true
}

// SetViewId sets field value.
func (o *ReplayAnalysisScreenshot) SetViewId(v string) {
	o.ViewId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ReplayAnalysisScreenshot) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["screenshot_key"] = o.ScreenshotKey
	toSerialize["session_id"] = o.SessionId
	toSerialize["timestamp_ms"] = o.TimestampMs
	toSerialize["view_id"] = o.ViewId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ReplayAnalysisScreenshot) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ScreenshotKey *string `json:"screenshot_key"`
		SessionId     *string `json:"session_id"`
		TimestampMs   *int64  `json:"timestamp_ms"`
		ViewId        *string `json:"view_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ScreenshotKey == nil {
		return fmt.Errorf("required field screenshot_key missing")
	}
	if all.SessionId == nil {
		return fmt.Errorf("required field session_id missing")
	}
	if all.TimestampMs == nil {
		return fmt.Errorf("required field timestamp_ms missing")
	}
	if all.ViewId == nil {
		return fmt.Errorf("required field view_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"screenshot_key", "session_id", "timestamp_ms", "view_id"})
	} else {
		return err
	}
	o.ScreenshotKey = *all.ScreenshotKey
	o.SessionId = *all.SessionId
	o.TimestampMs = *all.TimestampMs
	o.ViewId = *all.ViewId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
