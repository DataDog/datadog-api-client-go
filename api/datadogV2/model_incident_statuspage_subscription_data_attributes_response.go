// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentStatuspageSubscriptionDataAttributesResponse Attributes of an email subscription.
type IncidentStatuspageSubscriptionDataAttributesResponse struct {
	// Whether the subscription has been confirmed.
	Confirmed bool `json:"confirmed"`
	// Timestamp when the subscription was created.
	CreatedAt time.Time `json:"created_at"`
	// The subscribed email address.
	Email string `json:"email"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentStatuspageSubscriptionDataAttributesResponse instantiates a new IncidentStatuspageSubscriptionDataAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentStatuspageSubscriptionDataAttributesResponse(confirmed bool, createdAt time.Time, email string) *IncidentStatuspageSubscriptionDataAttributesResponse {
	this := IncidentStatuspageSubscriptionDataAttributesResponse{}
	this.Confirmed = confirmed
	this.CreatedAt = createdAt
	this.Email = email
	return &this
}

// NewIncidentStatuspageSubscriptionDataAttributesResponseWithDefaults instantiates a new IncidentStatuspageSubscriptionDataAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentStatuspageSubscriptionDataAttributesResponseWithDefaults() *IncidentStatuspageSubscriptionDataAttributesResponse {
	this := IncidentStatuspageSubscriptionDataAttributesResponse{}
	return &this
}

// GetConfirmed returns the Confirmed field value.
func (o *IncidentStatuspageSubscriptionDataAttributesResponse) GetConfirmed() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Confirmed
}

// GetConfirmedOk returns a tuple with the Confirmed field value
// and a boolean to check if the value has been set.
func (o *IncidentStatuspageSubscriptionDataAttributesResponse) GetConfirmedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Confirmed, true
}

// SetConfirmed sets field value.
func (o *IncidentStatuspageSubscriptionDataAttributesResponse) SetConfirmed(v bool) {
	o.Confirmed = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *IncidentStatuspageSubscriptionDataAttributesResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *IncidentStatuspageSubscriptionDataAttributesResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *IncidentStatuspageSubscriptionDataAttributesResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetEmail returns the Email field value.
func (o *IncidentStatuspageSubscriptionDataAttributesResponse) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *IncidentStatuspageSubscriptionDataAttributesResponse) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value.
func (o *IncidentStatuspageSubscriptionDataAttributesResponse) SetEmail(v string) {
	o.Email = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentStatuspageSubscriptionDataAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["confirmed"] = o.Confirmed
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["email"] = o.Email

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentStatuspageSubscriptionDataAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Confirmed *bool      `json:"confirmed"`
		CreatedAt *time.Time `json:"created_at"`
		Email     *string    `json:"email"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Confirmed == nil {
		return fmt.Errorf("required field confirmed missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.Email == nil {
		return fmt.Errorf("required field email missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"confirmed", "created_at", "email"})
	} else {
		return err
	}
	o.Confirmed = *all.Confirmed
	o.CreatedAt = *all.CreatedAt
	o.Email = *all.Email

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
