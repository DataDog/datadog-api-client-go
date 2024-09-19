// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.


package datadogV1

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"

)


// EventListResponse An event list response.
type EventListResponse struct {
	// An array of events.
	Events []Event `json:"events,omitempty"`
	// A status.
	Status *string `json:"status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewEventListResponse instantiates a new EventListResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEventListResponse() *EventListResponse {
	this := EventListResponse{}
	return &this
}

// NewEventListResponseWithDefaults instantiates a new EventListResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEventListResponseWithDefaults() *EventListResponse {
	this := EventListResponse{}
	return &this
}
// GetEvents returns the Events field value if set, zero value otherwise.
func (o *EventListResponse) GetEvents() []Event {
	if o == nil || o.Events == nil {
		var ret []Event
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventListResponse) GetEventsOk() (*[]Event, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return &o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *EventListResponse) HasEvents() bool {
	return o != nil && o.Events != nil
}

// SetEvents gets a reference to the given []Event and assigns it to the Events field.
func (o *EventListResponse) SetEvents(v []Event) {
	o.Events = v
}


// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EventListResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventListResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EventListResponse) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *EventListResponse) SetStatus(v string) {
	o.Status = &v
}



// MarshalJSON serializes the struct using spec logic.
func (o EventListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EventListResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Events []Event `json:"events,omitempty"`
		Status *string `json:"status,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{ "events", "status",  })
	} else {
		return err
	}
	o.Events = all.Events
	o.Status = all.Status

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
