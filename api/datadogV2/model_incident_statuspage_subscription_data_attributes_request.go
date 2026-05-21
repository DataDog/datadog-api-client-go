// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentStatuspageSubscriptionDataAttributesRequest Attributes for creating an email subscription.
type IncidentStatuspageSubscriptionDataAttributesRequest struct {
	// The email address to subscribe.
	Email string `json:"email"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentStatuspageSubscriptionDataAttributesRequest instantiates a new IncidentStatuspageSubscriptionDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentStatuspageSubscriptionDataAttributesRequest(email string) *IncidentStatuspageSubscriptionDataAttributesRequest {
	this := IncidentStatuspageSubscriptionDataAttributesRequest{}
	this.Email = email
	return &this
}

// NewIncidentStatuspageSubscriptionDataAttributesRequestWithDefaults instantiates a new IncidentStatuspageSubscriptionDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentStatuspageSubscriptionDataAttributesRequestWithDefaults() *IncidentStatuspageSubscriptionDataAttributesRequest {
	this := IncidentStatuspageSubscriptionDataAttributesRequest{}
	return &this
}

// GetEmail returns the Email field value.
func (o *IncidentStatuspageSubscriptionDataAttributesRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *IncidentStatuspageSubscriptionDataAttributesRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value.
func (o *IncidentStatuspageSubscriptionDataAttributesRequest) SetEmail(v string) {
	o.Email = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentStatuspageSubscriptionDataAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["email"] = o.Email

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentStatuspageSubscriptionDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Email *string `json:"email"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Email == nil {
		return fmt.Errorf("required field email missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"email"})
	} else {
		return err
	}
	o.Email = *all.Email

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
