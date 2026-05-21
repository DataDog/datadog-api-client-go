// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentZoomMeeting A Zoom meeting associated with an incident.
type IncidentZoomMeeting struct {
	// The Zoom host identifier.
	HostId *string `json:"host_id,omitempty"`
	// The URL to join the meeting.
	JoinUrl string `json:"join_url"`
	// The Zoom meeting identifier.
	MeetingId int64 `json:"meeting_id"`
	// The meeting password.
	Password *string `json:"password,omitempty"`
	// The URL of the meeting recording.
	RecordingUrl *string `json:"recording_url,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentZoomMeeting instantiates a new IncidentZoomMeeting object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentZoomMeeting(joinUrl string, meetingId int64) *IncidentZoomMeeting {
	this := IncidentZoomMeeting{}
	this.JoinUrl = joinUrl
	this.MeetingId = meetingId
	return &this
}

// NewIncidentZoomMeetingWithDefaults instantiates a new IncidentZoomMeeting object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentZoomMeetingWithDefaults() *IncidentZoomMeeting {
	this := IncidentZoomMeeting{}
	return &this
}

// GetHostId returns the HostId field value if set, zero value otherwise.
func (o *IncidentZoomMeeting) GetHostId() string {
	if o == nil || o.HostId == nil {
		var ret string
		return ret
	}
	return *o.HostId
}

// GetHostIdOk returns a tuple with the HostId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentZoomMeeting) GetHostIdOk() (*string, bool) {
	if o == nil || o.HostId == nil {
		return nil, false
	}
	return o.HostId, true
}

// HasHostId returns a boolean if a field has been set.
func (o *IncidentZoomMeeting) HasHostId() bool {
	return o != nil && o.HostId != nil
}

// SetHostId gets a reference to the given string and assigns it to the HostId field.
func (o *IncidentZoomMeeting) SetHostId(v string) {
	o.HostId = &v
}

// GetJoinUrl returns the JoinUrl field value.
func (o *IncidentZoomMeeting) GetJoinUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.JoinUrl
}

// GetJoinUrlOk returns a tuple with the JoinUrl field value
// and a boolean to check if the value has been set.
func (o *IncidentZoomMeeting) GetJoinUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JoinUrl, true
}

// SetJoinUrl sets field value.
func (o *IncidentZoomMeeting) SetJoinUrl(v string) {
	o.JoinUrl = v
}

// GetMeetingId returns the MeetingId field value.
func (o *IncidentZoomMeeting) GetMeetingId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.MeetingId
}

// GetMeetingIdOk returns a tuple with the MeetingId field value
// and a boolean to check if the value has been set.
func (o *IncidentZoomMeeting) GetMeetingIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MeetingId, true
}

// SetMeetingId sets field value.
func (o *IncidentZoomMeeting) SetMeetingId(v int64) {
	o.MeetingId = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *IncidentZoomMeeting) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentZoomMeeting) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *IncidentZoomMeeting) HasPassword() bool {
	return o != nil && o.Password != nil
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *IncidentZoomMeeting) SetPassword(v string) {
	o.Password = &v
}

// GetRecordingUrl returns the RecordingUrl field value if set, zero value otherwise.
func (o *IncidentZoomMeeting) GetRecordingUrl() string {
	if o == nil || o.RecordingUrl == nil {
		var ret string
		return ret
	}
	return *o.RecordingUrl
}

// GetRecordingUrlOk returns a tuple with the RecordingUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentZoomMeeting) GetRecordingUrlOk() (*string, bool) {
	if o == nil || o.RecordingUrl == nil {
		return nil, false
	}
	return o.RecordingUrl, true
}

// HasRecordingUrl returns a boolean if a field has been set.
func (o *IncidentZoomMeeting) HasRecordingUrl() bool {
	return o != nil && o.RecordingUrl != nil
}

// SetRecordingUrl gets a reference to the given string and assigns it to the RecordingUrl field.
func (o *IncidentZoomMeeting) SetRecordingUrl(v string) {
	o.RecordingUrl = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentZoomMeeting) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.HostId != nil {
		toSerialize["host_id"] = o.HostId
	}
	toSerialize["join_url"] = o.JoinUrl
	toSerialize["meeting_id"] = o.MeetingId
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.RecordingUrl != nil {
		toSerialize["recording_url"] = o.RecordingUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentZoomMeeting) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		HostId       *string `json:"host_id,omitempty"`
		JoinUrl      *string `json:"join_url"`
		MeetingId    *int64  `json:"meeting_id"`
		Password     *string `json:"password,omitempty"`
		RecordingUrl *string `json:"recording_url,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.JoinUrl == nil {
		return fmt.Errorf("required field join_url missing")
	}
	if all.MeetingId == nil {
		return fmt.Errorf("required field meeting_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"host_id", "join_url", "meeting_id", "password", "recording_url"})
	} else {
		return err
	}
	o.HostId = all.HostId
	o.JoinUrl = *all.JoinUrl
	o.MeetingId = *all.MeetingId
	o.Password = all.Password
	o.RecordingUrl = all.RecordingUrl

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
