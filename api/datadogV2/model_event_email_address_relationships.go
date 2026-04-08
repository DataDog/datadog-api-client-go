// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventEmailAddressRelationships Relationships associated with an event email address resource.
type EventEmailAddressRelationships struct {
	// Relationship to the user who created the email address.
	CreatedBy EventEmailAddressCreatedByRelationship `json:"created_by"`
	// Relationship to the user who revoked the email address.
	RevokedBy NullableEventEmailAddressRevokedByRelationship `json:"revoked_by,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEventEmailAddressRelationships instantiates a new EventEmailAddressRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEventEmailAddressRelationships(createdBy EventEmailAddressCreatedByRelationship) *EventEmailAddressRelationships {
	this := EventEmailAddressRelationships{}
	this.CreatedBy = createdBy
	return &this
}

// NewEventEmailAddressRelationshipsWithDefaults instantiates a new EventEmailAddressRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEventEmailAddressRelationshipsWithDefaults() *EventEmailAddressRelationships {
	this := EventEmailAddressRelationships{}
	return &this
}

// GetCreatedBy returns the CreatedBy field value.
func (o *EventEmailAddressRelationships) GetCreatedBy() EventEmailAddressCreatedByRelationship {
	if o == nil {
		var ret EventEmailAddressCreatedByRelationship
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *EventEmailAddressRelationships) GetCreatedByOk() (*EventEmailAddressCreatedByRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value.
func (o *EventEmailAddressRelationships) SetCreatedBy(v EventEmailAddressCreatedByRelationship) {
	o.CreatedBy = v
}

// GetRevokedBy returns the RevokedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventEmailAddressRelationships) GetRevokedBy() EventEmailAddressRevokedByRelationship {
	if o == nil || o.RevokedBy.Get() == nil {
		var ret EventEmailAddressRevokedByRelationship
		return ret
	}
	return *o.RevokedBy.Get()
}

// GetRevokedByOk returns a tuple with the RevokedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EventEmailAddressRelationships) GetRevokedByOk() (*EventEmailAddressRevokedByRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RevokedBy.Get(), o.RevokedBy.IsSet()
}

// HasRevokedBy returns a boolean if a field has been set.
func (o *EventEmailAddressRelationships) HasRevokedBy() bool {
	return o != nil && o.RevokedBy.IsSet()
}

// SetRevokedBy gets a reference to the given NullableEventEmailAddressRevokedByRelationship and assigns it to the RevokedBy field.
func (o *EventEmailAddressRelationships) SetRevokedBy(v EventEmailAddressRevokedByRelationship) {
	o.RevokedBy.Set(&v)
}

// SetRevokedByNil sets the value for RevokedBy to be an explicit nil.
func (o *EventEmailAddressRelationships) SetRevokedByNil() {
	o.RevokedBy.Set(nil)
}

// UnsetRevokedBy ensures that no value is present for RevokedBy, not even an explicit nil.
func (o *EventEmailAddressRelationships) UnsetRevokedBy() {
	o.RevokedBy.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o EventEmailAddressRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["created_by"] = o.CreatedBy
	if o.RevokedBy.IsSet() {
		toSerialize["revoked_by"] = o.RevokedBy.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EventEmailAddressRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedBy *EventEmailAddressCreatedByRelationship        `json:"created_by"`
		RevokedBy NullableEventEmailAddressRevokedByRelationship `json:"revoked_by,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedBy == nil {
		return fmt.Errorf("required field created_by missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_by", "revoked_by"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CreatedBy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CreatedBy = *all.CreatedBy
	o.RevokedBy = all.RevokedBy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
