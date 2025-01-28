// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ActionInbox Action of the inbox rule
type ActionInbox struct {
	// Free text to add a reason description.
	ReasonDescription *string `json:"reason_description,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewActionInbox instantiates a new ActionInbox object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewActionInbox() *ActionInbox {
	this := ActionInbox{}
	return &this
}

// NewActionInboxWithDefaults instantiates a new ActionInbox object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewActionInboxWithDefaults() *ActionInbox {
	this := ActionInbox{}
	return &this
}

// GetReasonDescription returns the ReasonDescription field value if set, zero value otherwise.
func (o *ActionInbox) GetReasonDescription() string {
	if o == nil || o.ReasonDescription == nil {
		var ret string
		return ret
	}
	return *o.ReasonDescription
}

// GetReasonDescriptionOk returns a tuple with the ReasonDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionInbox) GetReasonDescriptionOk() (*string, bool) {
	if o == nil || o.ReasonDescription == nil {
		return nil, false
	}
	return o.ReasonDescription, true
}

// HasReasonDescription returns a boolean if a field has been set.
func (o *ActionInbox) HasReasonDescription() bool {
	return o != nil && o.ReasonDescription != nil
}

// SetReasonDescription gets a reference to the given string and assigns it to the ReasonDescription field.
func (o *ActionInbox) SetReasonDescription(v string) {
	o.ReasonDescription = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ActionInbox) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ReasonDescription != nil {
		toSerialize["reason_description"] = o.ReasonDescription
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ActionInbox) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ReasonDescription *string `json:"reason_description,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"reason_description"})
	} else {
		return err
	}
	o.ReasonDescription = all.ReasonDescription

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
