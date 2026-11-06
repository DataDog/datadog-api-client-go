// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRecoverySessionsCreateAttributes Recover 1 to 100 explicit archived sessions. Session IDs must be unique within the request.
type RumRecoverySessionsCreateAttributes struct {
	// Targeted recovery. Omitted request_type defaults to sessions.
	RequestType *RumRecoverySessionsRequestType `json:"request_type,omitempty"`
	// Sessions to recover.
	Sessions []RumRecoverySessionInput `json:"sessions"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumRecoverySessionsCreateAttributes instantiates a new RumRecoverySessionsCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumRecoverySessionsCreateAttributes(sessions []RumRecoverySessionInput) *RumRecoverySessionsCreateAttributes {
	this := RumRecoverySessionsCreateAttributes{}
	var requestType RumRecoverySessionsRequestType = RUMRECOVERYSESSIONSREQUESTTYPE_SESSIONS
	this.RequestType = &requestType
	this.Sessions = sessions
	return &this
}

// NewRumRecoverySessionsCreateAttributesWithDefaults instantiates a new RumRecoverySessionsCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumRecoverySessionsCreateAttributesWithDefaults() *RumRecoverySessionsCreateAttributes {
	this := RumRecoverySessionsCreateAttributes{}
	var requestType RumRecoverySessionsRequestType = RUMRECOVERYSESSIONSREQUESTTYPE_SESSIONS
	this.RequestType = &requestType
	return &this
}

// GetRequestType returns the RequestType field value if set, zero value otherwise.
func (o *RumRecoverySessionsCreateAttributes) GetRequestType() RumRecoverySessionsRequestType {
	if o == nil || o.RequestType == nil {
		var ret RumRecoverySessionsRequestType
		return ret
	}
	return *o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumRecoverySessionsCreateAttributes) GetRequestTypeOk() (*RumRecoverySessionsRequestType, bool) {
	if o == nil || o.RequestType == nil {
		return nil, false
	}
	return o.RequestType, true
}

// HasRequestType returns a boolean if a field has been set.
func (o *RumRecoverySessionsCreateAttributes) HasRequestType() bool {
	return o != nil && o.RequestType != nil
}

// SetRequestType gets a reference to the given RumRecoverySessionsRequestType and assigns it to the RequestType field.
func (o *RumRecoverySessionsCreateAttributes) SetRequestType(v RumRecoverySessionsRequestType) {
	o.RequestType = &v
}

// GetSessions returns the Sessions field value.
func (o *RumRecoverySessionsCreateAttributes) GetSessions() []RumRecoverySessionInput {
	if o == nil {
		var ret []RumRecoverySessionInput
		return ret
	}
	return o.Sessions
}

// GetSessionsOk returns a tuple with the Sessions field value
// and a boolean to check if the value has been set.
func (o *RumRecoverySessionsCreateAttributes) GetSessionsOk() (*[]RumRecoverySessionInput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sessions, true
}

// SetSessions sets field value.
func (o *RumRecoverySessionsCreateAttributes) SetSessions(v []RumRecoverySessionInput) {
	o.Sessions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumRecoverySessionsCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.RequestType != nil {
		toSerialize["request_type"] = o.RequestType
	}
	toSerialize["sessions"] = o.Sessions

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumRecoverySessionsCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RequestType *RumRecoverySessionsRequestType `json:"request_type,omitempty"`
		Sessions    *[]RumRecoverySessionInput      `json:"sessions"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Sessions == nil {
		return fmt.Errorf("required field sessions missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"request_type", "sessions"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.RequestType != nil && !all.RequestType.IsValid() {
		hasInvalidField = true
	} else {
		o.RequestType = all.RequestType
	}
	o.Sessions = *all.Sessions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
