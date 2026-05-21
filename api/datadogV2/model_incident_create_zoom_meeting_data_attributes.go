// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentCreateZoomMeetingDataAttributes Attributes for creating a Zoom meeting.
type IncidentCreateZoomMeetingDataAttributes struct {
	// The topic of the Zoom meeting.
	Topic string `json:"topic"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentCreateZoomMeetingDataAttributes instantiates a new IncidentCreateZoomMeetingDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentCreateZoomMeetingDataAttributes(topic string) *IncidentCreateZoomMeetingDataAttributes {
	this := IncidentCreateZoomMeetingDataAttributes{}
	this.Topic = topic
	return &this
}

// NewIncidentCreateZoomMeetingDataAttributesWithDefaults instantiates a new IncidentCreateZoomMeetingDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentCreateZoomMeetingDataAttributesWithDefaults() *IncidentCreateZoomMeetingDataAttributes {
	this := IncidentCreateZoomMeetingDataAttributes{}
	return &this
}

// GetTopic returns the Topic field value.
func (o *IncidentCreateZoomMeetingDataAttributes) GetTopic() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Topic
}

// GetTopicOk returns a tuple with the Topic field value
// and a boolean to check if the value has been set.
func (o *IncidentCreateZoomMeetingDataAttributes) GetTopicOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Topic, true
}

// SetTopic sets field value.
func (o *IncidentCreateZoomMeetingDataAttributes) SetTopic(v string) {
	o.Topic = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentCreateZoomMeetingDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["topic"] = o.Topic

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentCreateZoomMeetingDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Topic *string `json:"topic"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Topic == nil {
		return fmt.Errorf("required field topic missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"topic"})
	} else {
		return err
	}
	o.Topic = *all.Topic

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
