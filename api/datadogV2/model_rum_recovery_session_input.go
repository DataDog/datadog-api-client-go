// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRecoverySessionInput Archived session to recover.
type RumRecoverySessionInput struct {
	// Application UUID.
	ApplicationId uuid.UUID `json:"application_id"`
	// Session UUID.
	SessionId uuid.UUID `json:"session_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumRecoverySessionInput instantiates a new RumRecoverySessionInput object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumRecoverySessionInput(applicationId uuid.UUID, sessionId uuid.UUID) *RumRecoverySessionInput {
	this := RumRecoverySessionInput{}
	this.ApplicationId = applicationId
	this.SessionId = sessionId
	return &this
}

// NewRumRecoverySessionInputWithDefaults instantiates a new RumRecoverySessionInput object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumRecoverySessionInputWithDefaults() *RumRecoverySessionInput {
	this := RumRecoverySessionInput{}
	return &this
}

// GetApplicationId returns the ApplicationId field value.
func (o *RumRecoverySessionInput) GetApplicationId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value
// and a boolean to check if the value has been set.
func (o *RumRecoverySessionInput) GetApplicationIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationId, true
}

// SetApplicationId sets field value.
func (o *RumRecoverySessionInput) SetApplicationId(v uuid.UUID) {
	o.ApplicationId = v
}

// GetSessionId returns the SessionId field value.
func (o *RumRecoverySessionInput) GetSessionId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value
// and a boolean to check if the value has been set.
func (o *RumRecoverySessionInput) GetSessionIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionId, true
}

// SetSessionId sets field value.
func (o *RumRecoverySessionInput) SetSessionId(v uuid.UUID) {
	o.SessionId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumRecoverySessionInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["application_id"] = o.ApplicationId
	toSerialize["session_id"] = o.SessionId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumRecoverySessionInput) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApplicationId *uuid.UUID `json:"application_id"`
		SessionId     *uuid.UUID `json:"session_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ApplicationId == nil {
		return fmt.Errorf("required field application_id missing")
	}
	if all.SessionId == nil {
		return fmt.Errorf("required field session_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"application_id", "session_id"})
	} else {
		return err
	}
	o.ApplicationId = *all.ApplicationId
	o.SessionId = *all.SessionId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
