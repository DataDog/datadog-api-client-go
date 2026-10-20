// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SavedFilterPermissionDeniedAttributes The caller lacks edit permission on one or more feature flags that reference this saved filter.
type SavedFilterPermissionDeniedAttributes struct {
	// Feature flags preventing this update.
	DeniedFlags []SavedFilterFlagReference `json:"denied_flags"`
	// The reason the saved filter update was rejected.
	Message string `json:"message"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSavedFilterPermissionDeniedAttributes instantiates a new SavedFilterPermissionDeniedAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSavedFilterPermissionDeniedAttributes(deniedFlags []SavedFilterFlagReference, message string) *SavedFilterPermissionDeniedAttributes {
	this := SavedFilterPermissionDeniedAttributes{}
	this.DeniedFlags = deniedFlags
	this.Message = message
	return &this
}

// NewSavedFilterPermissionDeniedAttributesWithDefaults instantiates a new SavedFilterPermissionDeniedAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSavedFilterPermissionDeniedAttributesWithDefaults() *SavedFilterPermissionDeniedAttributes {
	this := SavedFilterPermissionDeniedAttributes{}
	return &this
}

// GetDeniedFlags returns the DeniedFlags field value.
func (o *SavedFilterPermissionDeniedAttributes) GetDeniedFlags() []SavedFilterFlagReference {
	if o == nil {
		var ret []SavedFilterFlagReference
		return ret
	}
	return o.DeniedFlags
}

// GetDeniedFlagsOk returns a tuple with the DeniedFlags field value
// and a boolean to check if the value has been set.
func (o *SavedFilterPermissionDeniedAttributes) GetDeniedFlagsOk() (*[]SavedFilterFlagReference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeniedFlags, true
}

// SetDeniedFlags sets field value.
func (o *SavedFilterPermissionDeniedAttributes) SetDeniedFlags(v []SavedFilterFlagReference) {
	o.DeniedFlags = v
}

// GetMessage returns the Message field value.
func (o *SavedFilterPermissionDeniedAttributes) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SavedFilterPermissionDeniedAttributes) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value.
func (o *SavedFilterPermissionDeniedAttributes) SetMessage(v string) {
	o.Message = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SavedFilterPermissionDeniedAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["denied_flags"] = o.DeniedFlags
	toSerialize["message"] = o.Message

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SavedFilterPermissionDeniedAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DeniedFlags *[]SavedFilterFlagReference `json:"denied_flags"`
		Message     *string                     `json:"message"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DeniedFlags == nil {
		return fmt.Errorf("required field denied_flags missing")
	}
	if all.Message == nil {
		return fmt.Errorf("required field message missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"denied_flags", "message"})
	} else {
		return err
	}
	o.DeniedFlags = *all.DeniedFlags
	o.Message = *all.Message

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
