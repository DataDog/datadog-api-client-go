// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentCommunicationContentHandle A handle used for sending a communication.
type IncidentCommunicationContentHandle struct {
	// Timestamp when the handle was added.
	CreatedAt *string `json:"created_at,omitempty"`
	// The display name for the handle.
	DisplayName *string `json:"display_name,omitempty"`
	// The notification handle.
	Handle string `json:"handle"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentCommunicationContentHandle instantiates a new IncidentCommunicationContentHandle object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentCommunicationContentHandle(handle string) *IncidentCommunicationContentHandle {
	this := IncidentCommunicationContentHandle{}
	this.Handle = handle
	return &this
}

// NewIncidentCommunicationContentHandleWithDefaults instantiates a new IncidentCommunicationContentHandle object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentCommunicationContentHandleWithDefaults() *IncidentCommunicationContentHandle {
	this := IncidentCommunicationContentHandle{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *IncidentCommunicationContentHandle) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentCommunicationContentHandle) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *IncidentCommunicationContentHandle) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *IncidentCommunicationContentHandle) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *IncidentCommunicationContentHandle) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentCommunicationContentHandle) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *IncidentCommunicationContentHandle) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *IncidentCommunicationContentHandle) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetHandle returns the Handle field value.
func (o *IncidentCommunicationContentHandle) GetHandle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Handle
}

// GetHandleOk returns a tuple with the Handle field value
// and a boolean to check if the value has been set.
func (o *IncidentCommunicationContentHandle) GetHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Handle, true
}

// SetHandle sets field value.
func (o *IncidentCommunicationContentHandle) SetHandle(v string) {
	o.Handle = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentCommunicationContentHandle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	toSerialize["handle"] = o.Handle

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentCommunicationContentHandle) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt   *string `json:"created_at,omitempty"`
		DisplayName *string `json:"display_name,omitempty"`
		Handle      *string `json:"handle"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Handle == nil {
		return fmt.Errorf("required field handle missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "display_name", "handle"})
	} else {
		return err
	}
	o.CreatedAt = all.CreatedAt
	o.DisplayName = all.DisplayName
	o.Handle = *all.Handle

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
