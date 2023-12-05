// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FullAPIKeyLastUsedDate Attributes for the last time the specific API key was used.
type FullAPIKeyLastUsedDate struct {
	// The description of the what the API key was used for.
	Description *string `json:"description,omitempty"`
	// The data and time of when the API key was last used.
	Timestamp datadog.NullableString `json:"timestamp,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewFullAPIKeyLastUsedDate instantiates a new FullAPIKeyLastUsedDate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFullAPIKeyLastUsedDate() *FullAPIKeyLastUsedDate {
	this := FullAPIKeyLastUsedDate{}
	return &this
}

// NewFullAPIKeyLastUsedDateWithDefaults instantiates a new FullAPIKeyLastUsedDate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFullAPIKeyLastUsedDateWithDefaults() *FullAPIKeyLastUsedDate {
	this := FullAPIKeyLastUsedDate{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FullAPIKeyLastUsedDate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullAPIKeyLastUsedDate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FullAPIKeyLastUsedDate) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FullAPIKeyLastUsedDate) SetDescription(v string) {
	o.Description = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FullAPIKeyLastUsedDate) GetTimestamp() string {
	if o == nil || o.Timestamp.Get() == nil {
		var ret string
		return ret
	}
	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FullAPIKeyLastUsedDate) GetTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// HasTimestamp returns a boolean if a field has been set.
func (o *FullAPIKeyLastUsedDate) HasTimestamp() bool {
	return o != nil && o.Timestamp.IsSet()
}

// SetTimestamp gets a reference to the given datadog.NullableString and assigns it to the Timestamp field.
func (o *FullAPIKeyLastUsedDate) SetTimestamp(v string) {
	o.Timestamp.Set(&v)
}

// SetTimestampNil sets the value for Timestamp to be an explicit nil.
func (o *FullAPIKeyLastUsedDate) SetTimestampNil() {
	o.Timestamp.Set(nil)
}

// UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil.
func (o *FullAPIKeyLastUsedDate) UnsetTimestamp() {
	o.Timestamp.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o FullAPIKeyLastUsedDate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Timestamp.IsSet() {
		toSerialize["timestamp"] = o.Timestamp.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FullAPIKeyLastUsedDate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string                `json:"description,omitempty"`
		Timestamp   datadog.NullableString `json:"timestamp,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "timestamp"})
	} else {
		return err
	}
	o.Description = all.Description
	o.Timestamp = all.Timestamp

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
