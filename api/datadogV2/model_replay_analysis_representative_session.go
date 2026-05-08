// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReplayAnalysisRepresentativeSession A representative session illustrating a replay analysis issue.
type ReplayAnalysisRepresentativeSession struct {
	// Category of the issue observed in this session.
	IssueCategory string `json:"issue_category"`
	// AI-generated analysis details for a replay issue.
	LlmAnalysisDetails ReplayAnalysisLLMDetails `json:"llm_analysis_details"`
	// A screenshot captured during a replay session.
	Screenshot *ReplayAnalysisScreenshot `json:"screenshot,omitempty"`
	// Unique identifier of the representative session.
	SessionId string `json:"session_id"`
	// Session start timestamp in milliseconds.
	SessionStartTimestampMs int64 `json:"session_start_timestamp_ms"`
	// List of signals observed in the representative session.
	Signals []ReplayAnalysisSignal `json:"signals"`
	// Name of the view where the issue was observed.
	ViewName datadog.NullableString `json:"view_name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewReplayAnalysisRepresentativeSession instantiates a new ReplayAnalysisRepresentativeSession object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReplayAnalysisRepresentativeSession(issueCategory string, llmAnalysisDetails ReplayAnalysisLLMDetails, sessionId string, sessionStartTimestampMs int64, signals []ReplayAnalysisSignal) *ReplayAnalysisRepresentativeSession {
	this := ReplayAnalysisRepresentativeSession{}
	this.IssueCategory = issueCategory
	this.LlmAnalysisDetails = llmAnalysisDetails
	this.SessionId = sessionId
	this.SessionStartTimestampMs = sessionStartTimestampMs
	this.Signals = signals
	return &this
}

// NewReplayAnalysisRepresentativeSessionWithDefaults instantiates a new ReplayAnalysisRepresentativeSession object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReplayAnalysisRepresentativeSessionWithDefaults() *ReplayAnalysisRepresentativeSession {
	this := ReplayAnalysisRepresentativeSession{}
	return &this
}

// GetIssueCategory returns the IssueCategory field value.
func (o *ReplayAnalysisRepresentativeSession) GetIssueCategory() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.IssueCategory
}

// GetIssueCategoryOk returns a tuple with the IssueCategory field value
// and a boolean to check if the value has been set.
func (o *ReplayAnalysisRepresentativeSession) GetIssueCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssueCategory, true
}

// SetIssueCategory sets field value.
func (o *ReplayAnalysisRepresentativeSession) SetIssueCategory(v string) {
	o.IssueCategory = v
}

// GetLlmAnalysisDetails returns the LlmAnalysisDetails field value.
func (o *ReplayAnalysisRepresentativeSession) GetLlmAnalysisDetails() ReplayAnalysisLLMDetails {
	if o == nil {
		var ret ReplayAnalysisLLMDetails
		return ret
	}
	return o.LlmAnalysisDetails
}

// GetLlmAnalysisDetailsOk returns a tuple with the LlmAnalysisDetails field value
// and a boolean to check if the value has been set.
func (o *ReplayAnalysisRepresentativeSession) GetLlmAnalysisDetailsOk() (*ReplayAnalysisLLMDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LlmAnalysisDetails, true
}

// SetLlmAnalysisDetails sets field value.
func (o *ReplayAnalysisRepresentativeSession) SetLlmAnalysisDetails(v ReplayAnalysisLLMDetails) {
	o.LlmAnalysisDetails = v
}

// GetScreenshot returns the Screenshot field value if set, zero value otherwise.
func (o *ReplayAnalysisRepresentativeSession) GetScreenshot() ReplayAnalysisScreenshot {
	if o == nil || o.Screenshot == nil {
		var ret ReplayAnalysisScreenshot
		return ret
	}
	return *o.Screenshot
}

// GetScreenshotOk returns a tuple with the Screenshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplayAnalysisRepresentativeSession) GetScreenshotOk() (*ReplayAnalysisScreenshot, bool) {
	if o == nil || o.Screenshot == nil {
		return nil, false
	}
	return o.Screenshot, true
}

// HasScreenshot returns a boolean if a field has been set.
func (o *ReplayAnalysisRepresentativeSession) HasScreenshot() bool {
	return o != nil && o.Screenshot != nil
}

// SetScreenshot gets a reference to the given ReplayAnalysisScreenshot and assigns it to the Screenshot field.
func (o *ReplayAnalysisRepresentativeSession) SetScreenshot(v ReplayAnalysisScreenshot) {
	o.Screenshot = &v
}

// GetSessionId returns the SessionId field value.
func (o *ReplayAnalysisRepresentativeSession) GetSessionId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value
// and a boolean to check if the value has been set.
func (o *ReplayAnalysisRepresentativeSession) GetSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionId, true
}

// SetSessionId sets field value.
func (o *ReplayAnalysisRepresentativeSession) SetSessionId(v string) {
	o.SessionId = v
}

// GetSessionStartTimestampMs returns the SessionStartTimestampMs field value.
func (o *ReplayAnalysisRepresentativeSession) GetSessionStartTimestampMs() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.SessionStartTimestampMs
}

// GetSessionStartTimestampMsOk returns a tuple with the SessionStartTimestampMs field value
// and a boolean to check if the value has been set.
func (o *ReplayAnalysisRepresentativeSession) GetSessionStartTimestampMsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionStartTimestampMs, true
}

// SetSessionStartTimestampMs sets field value.
func (o *ReplayAnalysisRepresentativeSession) SetSessionStartTimestampMs(v int64) {
	o.SessionStartTimestampMs = v
}

// GetSignals returns the Signals field value.
func (o *ReplayAnalysisRepresentativeSession) GetSignals() []ReplayAnalysisSignal {
	if o == nil {
		var ret []ReplayAnalysisSignal
		return ret
	}
	return o.Signals
}

// GetSignalsOk returns a tuple with the Signals field value
// and a boolean to check if the value has been set.
func (o *ReplayAnalysisRepresentativeSession) GetSignalsOk() (*[]ReplayAnalysisSignal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signals, true
}

// SetSignals sets field value.
func (o *ReplayAnalysisRepresentativeSession) SetSignals(v []ReplayAnalysisSignal) {
	o.Signals = v
}

// GetViewName returns the ViewName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReplayAnalysisRepresentativeSession) GetViewName() string {
	if o == nil || o.ViewName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ViewName.Get()
}

// GetViewNameOk returns a tuple with the ViewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReplayAnalysisRepresentativeSession) GetViewNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ViewName.Get(), o.ViewName.IsSet()
}

// HasViewName returns a boolean if a field has been set.
func (o *ReplayAnalysisRepresentativeSession) HasViewName() bool {
	return o != nil && o.ViewName.IsSet()
}

// SetViewName gets a reference to the given datadog.NullableString and assigns it to the ViewName field.
func (o *ReplayAnalysisRepresentativeSession) SetViewName(v string) {
	o.ViewName.Set(&v)
}

// SetViewNameNil sets the value for ViewName to be an explicit nil.
func (o *ReplayAnalysisRepresentativeSession) SetViewNameNil() {
	o.ViewName.Set(nil)
}

// UnsetViewName ensures that no value is present for ViewName, not even an explicit nil.
func (o *ReplayAnalysisRepresentativeSession) UnsetViewName() {
	o.ViewName.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ReplayAnalysisRepresentativeSession) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["issue_category"] = o.IssueCategory
	toSerialize["llm_analysis_details"] = o.LlmAnalysisDetails
	if o.Screenshot != nil {
		toSerialize["screenshot"] = o.Screenshot
	}
	toSerialize["session_id"] = o.SessionId
	toSerialize["session_start_timestamp_ms"] = o.SessionStartTimestampMs
	toSerialize["signals"] = o.Signals
	if o.ViewName.IsSet() {
		toSerialize["view_name"] = o.ViewName.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ReplayAnalysisRepresentativeSession) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IssueCategory           *string                   `json:"issue_category"`
		LlmAnalysisDetails      *ReplayAnalysisLLMDetails `json:"llm_analysis_details"`
		Screenshot              *ReplayAnalysisScreenshot `json:"screenshot,omitempty"`
		SessionId               *string                   `json:"session_id"`
		SessionStartTimestampMs *int64                    `json:"session_start_timestamp_ms"`
		Signals                 *[]ReplayAnalysisSignal   `json:"signals"`
		ViewName                datadog.NullableString    `json:"view_name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.IssueCategory == nil {
		return fmt.Errorf("required field issue_category missing")
	}
	if all.LlmAnalysisDetails == nil {
		return fmt.Errorf("required field llm_analysis_details missing")
	}
	if all.SessionId == nil {
		return fmt.Errorf("required field session_id missing")
	}
	if all.SessionStartTimestampMs == nil {
		return fmt.Errorf("required field session_start_timestamp_ms missing")
	}
	if all.Signals == nil {
		return fmt.Errorf("required field signals missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"issue_category", "llm_analysis_details", "screenshot", "session_id", "session_start_timestamp_ms", "signals", "view_name"})
	} else {
		return err
	}

	hasInvalidField := false
	o.IssueCategory = *all.IssueCategory
	if all.LlmAnalysisDetails.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LlmAnalysisDetails = *all.LlmAnalysisDetails
	if all.Screenshot != nil && all.Screenshot.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Screenshot = all.Screenshot
	o.SessionId = *all.SessionId
	o.SessionStartTimestampMs = *all.SessionStartTimestampMs
	o.Signals = *all.Signals
	o.ViewName = all.ViewName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
