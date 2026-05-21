// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentStatuspagePreferencesDataAttributes Attributes of subscription preferences.
type IncidentStatuspagePreferencesDataAttributes struct {
	// Whether the user is subscribed.
	Subscribed bool `json:"subscribed"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentStatuspagePreferencesDataAttributes instantiates a new IncidentStatuspagePreferencesDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentStatuspagePreferencesDataAttributes(subscribed bool) *IncidentStatuspagePreferencesDataAttributes {
	this := IncidentStatuspagePreferencesDataAttributes{}
	this.Subscribed = subscribed
	return &this
}

// NewIncidentStatuspagePreferencesDataAttributesWithDefaults instantiates a new IncidentStatuspagePreferencesDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentStatuspagePreferencesDataAttributesWithDefaults() *IncidentStatuspagePreferencesDataAttributes {
	this := IncidentStatuspagePreferencesDataAttributes{}
	return &this
}

// GetSubscribed returns the Subscribed field value.
func (o *IncidentStatuspagePreferencesDataAttributes) GetSubscribed() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Subscribed
}

// GetSubscribedOk returns a tuple with the Subscribed field value
// and a boolean to check if the value has been set.
func (o *IncidentStatuspagePreferencesDataAttributes) GetSubscribedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subscribed, true
}

// SetSubscribed sets field value.
func (o *IncidentStatuspagePreferencesDataAttributes) SetSubscribed(v bool) {
	o.Subscribed = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentStatuspagePreferencesDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["subscribed"] = o.Subscribed

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentStatuspagePreferencesDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Subscribed *bool `json:"subscribed"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Subscribed == nil {
		return fmt.Errorf("required field subscribed missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"subscribed"})
	} else {
		return err
	}
	o.Subscribed = *all.Subscribed

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
