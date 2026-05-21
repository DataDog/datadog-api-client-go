// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentZoomConfigurationDataAttributesRequest Attributes for creating or updating a Zoom configuration.
type IncidentZoomConfigurationDataAttributesRequest struct {
	// Whether manual meeting creation is enabled.
	ManualMeetingCreation *bool `json:"manual_meeting_creation,omitempty"`
	// Whether meeting chat timeline sync is enabled.
	MeetingChatTimelineSync *bool `json:"meeting_chat_timeline_sync,omitempty"`
	// Whether post-meeting summary is enabled.
	PostMeetingSummary *bool `json:"post_meeting_summary,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentZoomConfigurationDataAttributesRequest instantiates a new IncidentZoomConfigurationDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentZoomConfigurationDataAttributesRequest() *IncidentZoomConfigurationDataAttributesRequest {
	this := IncidentZoomConfigurationDataAttributesRequest{}
	return &this
}

// NewIncidentZoomConfigurationDataAttributesRequestWithDefaults instantiates a new IncidentZoomConfigurationDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentZoomConfigurationDataAttributesRequestWithDefaults() *IncidentZoomConfigurationDataAttributesRequest {
	this := IncidentZoomConfigurationDataAttributesRequest{}
	return &this
}

// GetManualMeetingCreation returns the ManualMeetingCreation field value if set, zero value otherwise.
func (o *IncidentZoomConfigurationDataAttributesRequest) GetManualMeetingCreation() bool {
	if o == nil || o.ManualMeetingCreation == nil {
		var ret bool
		return ret
	}
	return *o.ManualMeetingCreation
}

// GetManualMeetingCreationOk returns a tuple with the ManualMeetingCreation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentZoomConfigurationDataAttributesRequest) GetManualMeetingCreationOk() (*bool, bool) {
	if o == nil || o.ManualMeetingCreation == nil {
		return nil, false
	}
	return o.ManualMeetingCreation, true
}

// HasManualMeetingCreation returns a boolean if a field has been set.
func (o *IncidentZoomConfigurationDataAttributesRequest) HasManualMeetingCreation() bool {
	return o != nil && o.ManualMeetingCreation != nil
}

// SetManualMeetingCreation gets a reference to the given bool and assigns it to the ManualMeetingCreation field.
func (o *IncidentZoomConfigurationDataAttributesRequest) SetManualMeetingCreation(v bool) {
	o.ManualMeetingCreation = &v
}

// GetMeetingChatTimelineSync returns the MeetingChatTimelineSync field value if set, zero value otherwise.
func (o *IncidentZoomConfigurationDataAttributesRequest) GetMeetingChatTimelineSync() bool {
	if o == nil || o.MeetingChatTimelineSync == nil {
		var ret bool
		return ret
	}
	return *o.MeetingChatTimelineSync
}

// GetMeetingChatTimelineSyncOk returns a tuple with the MeetingChatTimelineSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentZoomConfigurationDataAttributesRequest) GetMeetingChatTimelineSyncOk() (*bool, bool) {
	if o == nil || o.MeetingChatTimelineSync == nil {
		return nil, false
	}
	return o.MeetingChatTimelineSync, true
}

// HasMeetingChatTimelineSync returns a boolean if a field has been set.
func (o *IncidentZoomConfigurationDataAttributesRequest) HasMeetingChatTimelineSync() bool {
	return o != nil && o.MeetingChatTimelineSync != nil
}

// SetMeetingChatTimelineSync gets a reference to the given bool and assigns it to the MeetingChatTimelineSync field.
func (o *IncidentZoomConfigurationDataAttributesRequest) SetMeetingChatTimelineSync(v bool) {
	o.MeetingChatTimelineSync = &v
}

// GetPostMeetingSummary returns the PostMeetingSummary field value if set, zero value otherwise.
func (o *IncidentZoomConfigurationDataAttributesRequest) GetPostMeetingSummary() bool {
	if o == nil || o.PostMeetingSummary == nil {
		var ret bool
		return ret
	}
	return *o.PostMeetingSummary
}

// GetPostMeetingSummaryOk returns a tuple with the PostMeetingSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentZoomConfigurationDataAttributesRequest) GetPostMeetingSummaryOk() (*bool, bool) {
	if o == nil || o.PostMeetingSummary == nil {
		return nil, false
	}
	return o.PostMeetingSummary, true
}

// HasPostMeetingSummary returns a boolean if a field has been set.
func (o *IncidentZoomConfigurationDataAttributesRequest) HasPostMeetingSummary() bool {
	return o != nil && o.PostMeetingSummary != nil
}

// SetPostMeetingSummary gets a reference to the given bool and assigns it to the PostMeetingSummary field.
func (o *IncidentZoomConfigurationDataAttributesRequest) SetPostMeetingSummary(v bool) {
	o.PostMeetingSummary = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentZoomConfigurationDataAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ManualMeetingCreation != nil {
		toSerialize["manual_meeting_creation"] = o.ManualMeetingCreation
	}
	if o.MeetingChatTimelineSync != nil {
		toSerialize["meeting_chat_timeline_sync"] = o.MeetingChatTimelineSync
	}
	if o.PostMeetingSummary != nil {
		toSerialize["post_meeting_summary"] = o.PostMeetingSummary
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentZoomConfigurationDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ManualMeetingCreation   *bool `json:"manual_meeting_creation,omitempty"`
		MeetingChatTimelineSync *bool `json:"meeting_chat_timeline_sync,omitempty"`
		PostMeetingSummary      *bool `json:"post_meeting_summary,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"manual_meeting_creation", "meeting_chat_timeline_sync", "post_meeting_summary"})
	} else {
		return err
	}
	o.ManualMeetingCreation = all.ManualMeetingCreation
	o.MeetingChatTimelineSync = all.MeetingChatTimelineSync
	o.PostMeetingSummary = all.PostMeetingSummary

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
