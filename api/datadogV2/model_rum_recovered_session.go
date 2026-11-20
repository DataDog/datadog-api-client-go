// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMRecoveredSession Progress for one targeted session.
type RUMRecoveredSession struct {
	// Application UUID.
	ApplicationId uuid.UUID `json:"application_id"`
	// Failure details, if present.
	ErrorMessage *string `json:"error_message,omitempty"`
	// Recovered event count.
	EventsRecovered int64 `json:"events_recovered"`
	// Session identifier.
	SessionId uuid.UUID `json:"session_id"`
	// Session availability or recovery state. `RECOVERED` indicates that the session is ready to query.
	Status RUMRecoverySessionStatus `json:"status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRUMRecoveredSession instantiates a new RUMRecoveredSession object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRUMRecoveredSession(applicationId uuid.UUID, eventsRecovered int64, sessionId uuid.UUID, status RUMRecoverySessionStatus) *RUMRecoveredSession {
	this := RUMRecoveredSession{}
	this.ApplicationId = applicationId
	this.EventsRecovered = eventsRecovered
	this.SessionId = sessionId
	this.Status = status
	return &this
}

// NewRUMRecoveredSessionWithDefaults instantiates a new RUMRecoveredSession object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRUMRecoveredSessionWithDefaults() *RUMRecoveredSession {
	this := RUMRecoveredSession{}
	return &this
}

// GetApplicationId returns the ApplicationId field value.
func (o *RUMRecoveredSession) GetApplicationId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value
// and a boolean to check if the value has been set.
func (o *RUMRecoveredSession) GetApplicationIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationId, true
}

// SetApplicationId sets field value.
func (o *RUMRecoveredSession) SetApplicationId(v uuid.UUID) {
	o.ApplicationId = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *RUMRecoveredSession) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMRecoveredSession) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *RUMRecoveredSession) HasErrorMessage() bool {
	return o != nil && o.ErrorMessage != nil
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *RUMRecoveredSession) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetEventsRecovered returns the EventsRecovered field value.
func (o *RUMRecoveredSession) GetEventsRecovered() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.EventsRecovered
}

// GetEventsRecoveredOk returns a tuple with the EventsRecovered field value
// and a boolean to check if the value has been set.
func (o *RUMRecoveredSession) GetEventsRecoveredOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventsRecovered, true
}

// SetEventsRecovered sets field value.
func (o *RUMRecoveredSession) SetEventsRecovered(v int64) {
	o.EventsRecovered = v
}

// GetSessionId returns the SessionId field value.
func (o *RUMRecoveredSession) GetSessionId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value
// and a boolean to check if the value has been set.
func (o *RUMRecoveredSession) GetSessionIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionId, true
}

// SetSessionId sets field value.
func (o *RUMRecoveredSession) SetSessionId(v uuid.UUID) {
	o.SessionId = v
}

// GetStatus returns the Status field value.
func (o *RUMRecoveredSession) GetStatus() RUMRecoverySessionStatus {
	if o == nil {
		var ret RUMRecoverySessionStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RUMRecoveredSession) GetStatusOk() (*RUMRecoverySessionStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *RUMRecoveredSession) SetStatus(v RUMRecoverySessionStatus) {
	o.Status = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RUMRecoveredSession) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["application_id"] = o.ApplicationId
	if o.ErrorMessage != nil {
		toSerialize["error_message"] = o.ErrorMessage
	}
	toSerialize["events_recovered"] = o.EventsRecovered
	toSerialize["session_id"] = o.SessionId
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RUMRecoveredSession) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApplicationId   *uuid.UUID                `json:"application_id"`
		ErrorMessage    *string                   `json:"error_message,omitempty"`
		EventsRecovered *int64                    `json:"events_recovered"`
		SessionId       *uuid.UUID                `json:"session_id"`
		Status          *RUMRecoverySessionStatus `json:"status"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ApplicationId == nil {
		return fmt.Errorf("required field application_id missing")
	}
	if all.EventsRecovered == nil {
		return fmt.Errorf("required field events_recovered missing")
	}
	if all.SessionId == nil {
		return fmt.Errorf("required field session_id missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"application_id", "error_message", "events_recovered", "session_id", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ApplicationId = *all.ApplicationId
	o.ErrorMessage = all.ErrorMessage
	o.EventsRecovered = *all.EventsRecovered
	o.SessionId = *all.SessionId
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
