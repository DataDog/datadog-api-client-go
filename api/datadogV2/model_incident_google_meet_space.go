// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentGoogleMeetSpace A Google Meet space associated with an incident.
type IncidentGoogleMeetSpace struct {
	// The URL to join the Google Meet space.
	JoinUrl string `json:"join_url"`
	// The meeting code for the space.
	MeetingCode string `json:"meeting_code"`
	// The name of the Google Meet space.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentGoogleMeetSpace instantiates a new IncidentGoogleMeetSpace object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentGoogleMeetSpace(joinUrl string, meetingCode string, name string) *IncidentGoogleMeetSpace {
	this := IncidentGoogleMeetSpace{}
	this.JoinUrl = joinUrl
	this.MeetingCode = meetingCode
	this.Name = name
	return &this
}

// NewIncidentGoogleMeetSpaceWithDefaults instantiates a new IncidentGoogleMeetSpace object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentGoogleMeetSpaceWithDefaults() *IncidentGoogleMeetSpace {
	this := IncidentGoogleMeetSpace{}
	return &this
}

// GetJoinUrl returns the JoinUrl field value.
func (o *IncidentGoogleMeetSpace) GetJoinUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.JoinUrl
}

// GetJoinUrlOk returns a tuple with the JoinUrl field value
// and a boolean to check if the value has been set.
func (o *IncidentGoogleMeetSpace) GetJoinUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JoinUrl, true
}

// SetJoinUrl sets field value.
func (o *IncidentGoogleMeetSpace) SetJoinUrl(v string) {
	o.JoinUrl = v
}

// GetMeetingCode returns the MeetingCode field value.
func (o *IncidentGoogleMeetSpace) GetMeetingCode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MeetingCode
}

// GetMeetingCodeOk returns a tuple with the MeetingCode field value
// and a boolean to check if the value has been set.
func (o *IncidentGoogleMeetSpace) GetMeetingCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MeetingCode, true
}

// SetMeetingCode sets field value.
func (o *IncidentGoogleMeetSpace) SetMeetingCode(v string) {
	o.MeetingCode = v
}

// GetName returns the Name field value.
func (o *IncidentGoogleMeetSpace) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IncidentGoogleMeetSpace) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *IncidentGoogleMeetSpace) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentGoogleMeetSpace) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["join_url"] = o.JoinUrl
	toSerialize["meeting_code"] = o.MeetingCode
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentGoogleMeetSpace) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		JoinUrl     *string `json:"join_url"`
		MeetingCode *string `json:"meeting_code"`
		Name        *string `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.JoinUrl == nil {
		return fmt.Errorf("required field join_url missing")
	}
	if all.MeetingCode == nil {
		return fmt.Errorf("required field meeting_code missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"join_url", "meeting_code", "name"})
	} else {
		return err
	}
	o.JoinUrl = *all.JoinUrl
	o.MeetingCode = *all.MeetingCode
	o.Name = *all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
