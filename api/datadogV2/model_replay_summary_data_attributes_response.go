// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReplaySummaryDataAttributesResponse Attributes of a RUM replay summary response.
type ReplaySummaryDataAttributesResponse struct {
	// List of chapters breaking down the session into distinct activity segments.
	Chapters []ReplaySummaryChapter `json:"chapters"`
	// Whether the session contained sufficient user activity to generate a summary.
	HasEnoughActivity bool `json:"has_enough_activity"`
	// Whether the session exceeded the event count limit for summary generation.
	HasTooManyEvents bool `json:"has_too_many_events"`
	// AI-generated summary of the replay session.
	Summary string `json:"summary"`
	// Version of the prompt used to generate the summary.
	Version string `json:"version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewReplaySummaryDataAttributesResponse instantiates a new ReplaySummaryDataAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReplaySummaryDataAttributesResponse(chapters []ReplaySummaryChapter, hasEnoughActivity bool, hasTooManyEvents bool, summary string, version string) *ReplaySummaryDataAttributesResponse {
	this := ReplaySummaryDataAttributesResponse{}
	this.Chapters = chapters
	this.HasEnoughActivity = hasEnoughActivity
	this.HasTooManyEvents = hasTooManyEvents
	this.Summary = summary
	this.Version = version
	return &this
}

// NewReplaySummaryDataAttributesResponseWithDefaults instantiates a new ReplaySummaryDataAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReplaySummaryDataAttributesResponseWithDefaults() *ReplaySummaryDataAttributesResponse {
	this := ReplaySummaryDataAttributesResponse{}
	return &this
}

// GetChapters returns the Chapters field value.
func (o *ReplaySummaryDataAttributesResponse) GetChapters() []ReplaySummaryChapter {
	if o == nil {
		var ret []ReplaySummaryChapter
		return ret
	}
	return o.Chapters
}

// GetChaptersOk returns a tuple with the Chapters field value
// and a boolean to check if the value has been set.
func (o *ReplaySummaryDataAttributesResponse) GetChaptersOk() (*[]ReplaySummaryChapter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chapters, true
}

// SetChapters sets field value.
func (o *ReplaySummaryDataAttributesResponse) SetChapters(v []ReplaySummaryChapter) {
	o.Chapters = v
}

// GetHasEnoughActivity returns the HasEnoughActivity field value.
func (o *ReplaySummaryDataAttributesResponse) GetHasEnoughActivity() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.HasEnoughActivity
}

// GetHasEnoughActivityOk returns a tuple with the HasEnoughActivity field value
// and a boolean to check if the value has been set.
func (o *ReplaySummaryDataAttributesResponse) GetHasEnoughActivityOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasEnoughActivity, true
}

// SetHasEnoughActivity sets field value.
func (o *ReplaySummaryDataAttributesResponse) SetHasEnoughActivity(v bool) {
	o.HasEnoughActivity = v
}

// GetHasTooManyEvents returns the HasTooManyEvents field value.
func (o *ReplaySummaryDataAttributesResponse) GetHasTooManyEvents() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.HasTooManyEvents
}

// GetHasTooManyEventsOk returns a tuple with the HasTooManyEvents field value
// and a boolean to check if the value has been set.
func (o *ReplaySummaryDataAttributesResponse) GetHasTooManyEventsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasTooManyEvents, true
}

// SetHasTooManyEvents sets field value.
func (o *ReplaySummaryDataAttributesResponse) SetHasTooManyEvents(v bool) {
	o.HasTooManyEvents = v
}

// GetSummary returns the Summary field value.
func (o *ReplaySummaryDataAttributesResponse) GetSummary() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *ReplaySummaryDataAttributesResponse) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value.
func (o *ReplaySummaryDataAttributesResponse) SetSummary(v string) {
	o.Summary = v
}

// GetVersion returns the Version field value.
func (o *ReplaySummaryDataAttributesResponse) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ReplaySummaryDataAttributesResponse) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *ReplaySummaryDataAttributesResponse) SetVersion(v string) {
	o.Version = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ReplaySummaryDataAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["chapters"] = o.Chapters
	toSerialize["has_enough_activity"] = o.HasEnoughActivity
	toSerialize["has_too_many_events"] = o.HasTooManyEvents
	toSerialize["summary"] = o.Summary
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ReplaySummaryDataAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Chapters          *[]ReplaySummaryChapter `json:"chapters"`
		HasEnoughActivity *bool                   `json:"has_enough_activity"`
		HasTooManyEvents  *bool                   `json:"has_too_many_events"`
		Summary           *string                 `json:"summary"`
		Version           *string                 `json:"version"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Chapters == nil {
		return fmt.Errorf("required field chapters missing")
	}
	if all.HasEnoughActivity == nil {
		return fmt.Errorf("required field has_enough_activity missing")
	}
	if all.HasTooManyEvents == nil {
		return fmt.Errorf("required field has_too_many_events missing")
	}
	if all.Summary == nil {
		return fmt.Errorf("required field summary missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"chapters", "has_enough_activity", "has_too_many_events", "summary", "version"})
	} else {
		return err
	}
	o.Chapters = *all.Chapters
	o.HasEnoughActivity = *all.HasEnoughActivity
	o.HasTooManyEvents = *all.HasTooManyEvents
	o.Summary = *all.Summary
	o.Version = *all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
