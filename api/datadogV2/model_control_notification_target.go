// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ControlNotificationTarget A destination that receives notifications for an event type.
type ControlNotificationTarget struct {
	// The destination handle, such as an email address, Slack channel, or user handle.
	Handle string `json:"handle"`
	// The type of notification destination.
	Type ControlNotificationTargetType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewControlNotificationTarget instantiates a new ControlNotificationTarget object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewControlNotificationTarget(handle string, typeVar ControlNotificationTargetType) *ControlNotificationTarget {
	this := ControlNotificationTarget{}
	this.Handle = handle
	this.Type = typeVar
	return &this
}

// NewControlNotificationTargetWithDefaults instantiates a new ControlNotificationTarget object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewControlNotificationTargetWithDefaults() *ControlNotificationTarget {
	this := ControlNotificationTarget{}
	return &this
}

// GetHandle returns the Handle field value.
func (o *ControlNotificationTarget) GetHandle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Handle
}

// GetHandleOk returns a tuple with the Handle field value
// and a boolean to check if the value has been set.
func (o *ControlNotificationTarget) GetHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Handle, true
}

// SetHandle sets field value.
func (o *ControlNotificationTarget) SetHandle(v string) {
	o.Handle = v
}

// GetType returns the Type field value.
func (o *ControlNotificationTarget) GetType() ControlNotificationTargetType {
	if o == nil {
		var ret ControlNotificationTargetType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ControlNotificationTarget) GetTypeOk() (*ControlNotificationTargetType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ControlNotificationTarget) SetType(v ControlNotificationTargetType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ControlNotificationTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["handle"] = o.Handle
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ControlNotificationTarget) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Handle *string                        `json:"handle"`
		Type   *ControlNotificationTargetType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Handle == nil {
		return fmt.Errorf("required field handle missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"handle", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Handle = *all.Handle
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
