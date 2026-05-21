// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentZoomConfigurationDataAttributesResponse Attributes of a Zoom configuration.
type IncidentZoomConfigurationDataAttributesResponse struct {
	// Timestamp when the configuration was created.
	CreatedAt time.Time `json:"created_at"`
	// Whether manual meeting creation is enabled.
	ManualMeetingCreation bool `json:"manual_meeting_creation"`
	// Whether meeting chat timeline sync is enabled.
	MeetingChatTimelineSync bool `json:"meeting_chat_timeline_sync"`
	// Timestamp when the configuration was last modified.
	ModifiedAt time.Time `json:"modified_at"`
	// Whether post-meeting summary is enabled.
	PostMeetingSummary bool `json:"post_meeting_summary"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentZoomConfigurationDataAttributesResponse instantiates a new IncidentZoomConfigurationDataAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentZoomConfigurationDataAttributesResponse(createdAt time.Time, manualMeetingCreation bool, meetingChatTimelineSync bool, modifiedAt time.Time, postMeetingSummary bool) *IncidentZoomConfigurationDataAttributesResponse {
	this := IncidentZoomConfigurationDataAttributesResponse{}
	this.CreatedAt = createdAt
	this.ManualMeetingCreation = manualMeetingCreation
	this.MeetingChatTimelineSync = meetingChatTimelineSync
	this.ModifiedAt = modifiedAt
	this.PostMeetingSummary = postMeetingSummary
	return &this
}

// NewIncidentZoomConfigurationDataAttributesResponseWithDefaults instantiates a new IncidentZoomConfigurationDataAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentZoomConfigurationDataAttributesResponseWithDefaults() *IncidentZoomConfigurationDataAttributesResponse {
	this := IncidentZoomConfigurationDataAttributesResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *IncidentZoomConfigurationDataAttributesResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *IncidentZoomConfigurationDataAttributesResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *IncidentZoomConfigurationDataAttributesResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetManualMeetingCreation returns the ManualMeetingCreation field value.
func (o *IncidentZoomConfigurationDataAttributesResponse) GetManualMeetingCreation() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.ManualMeetingCreation
}

// GetManualMeetingCreationOk returns a tuple with the ManualMeetingCreation field value
// and a boolean to check if the value has been set.
func (o *IncidentZoomConfigurationDataAttributesResponse) GetManualMeetingCreationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ManualMeetingCreation, true
}

// SetManualMeetingCreation sets field value.
func (o *IncidentZoomConfigurationDataAttributesResponse) SetManualMeetingCreation(v bool) {
	o.ManualMeetingCreation = v
}

// GetMeetingChatTimelineSync returns the MeetingChatTimelineSync field value.
func (o *IncidentZoomConfigurationDataAttributesResponse) GetMeetingChatTimelineSync() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.MeetingChatTimelineSync
}

// GetMeetingChatTimelineSyncOk returns a tuple with the MeetingChatTimelineSync field value
// and a boolean to check if the value has been set.
func (o *IncidentZoomConfigurationDataAttributesResponse) GetMeetingChatTimelineSyncOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MeetingChatTimelineSync, true
}

// SetMeetingChatTimelineSync sets field value.
func (o *IncidentZoomConfigurationDataAttributesResponse) SetMeetingChatTimelineSync(v bool) {
	o.MeetingChatTimelineSync = v
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *IncidentZoomConfigurationDataAttributesResponse) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *IncidentZoomConfigurationDataAttributesResponse) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *IncidentZoomConfigurationDataAttributesResponse) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetPostMeetingSummary returns the PostMeetingSummary field value.
func (o *IncidentZoomConfigurationDataAttributesResponse) GetPostMeetingSummary() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.PostMeetingSummary
}

// GetPostMeetingSummaryOk returns a tuple with the PostMeetingSummary field value
// and a boolean to check if the value has been set.
func (o *IncidentZoomConfigurationDataAttributesResponse) GetPostMeetingSummaryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostMeetingSummary, true
}

// SetPostMeetingSummary sets field value.
func (o *IncidentZoomConfigurationDataAttributesResponse) SetPostMeetingSummary(v bool) {
	o.PostMeetingSummary = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentZoomConfigurationDataAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["manual_meeting_creation"] = o.ManualMeetingCreation
	toSerialize["meeting_chat_timeline_sync"] = o.MeetingChatTimelineSync
	if o.ModifiedAt.Nanosecond() == 0 {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["post_meeting_summary"] = o.PostMeetingSummary

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentZoomConfigurationDataAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt               *time.Time `json:"created_at"`
		ManualMeetingCreation   *bool      `json:"manual_meeting_creation"`
		MeetingChatTimelineSync *bool      `json:"meeting_chat_timeline_sync"`
		ModifiedAt              *time.Time `json:"modified_at"`
		PostMeetingSummary      *bool      `json:"post_meeting_summary"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.ManualMeetingCreation == nil {
		return fmt.Errorf("required field manual_meeting_creation missing")
	}
	if all.MeetingChatTimelineSync == nil {
		return fmt.Errorf("required field meeting_chat_timeline_sync missing")
	}
	if all.ModifiedAt == nil {
		return fmt.Errorf("required field modified_at missing")
	}
	if all.PostMeetingSummary == nil {
		return fmt.Errorf("required field post_meeting_summary missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "manual_meeting_creation", "meeting_chat_timeline_sync", "modified_at", "post_meeting_summary"})
	} else {
		return err
	}
	o.CreatedAt = *all.CreatedAt
	o.ManualMeetingCreation = *all.ManualMeetingCreation
	o.MeetingChatTimelineSync = *all.MeetingChatTimelineSync
	o.ModifiedAt = *all.ModifiedAt
	o.PostMeetingSummary = *all.PostMeetingSummary

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
