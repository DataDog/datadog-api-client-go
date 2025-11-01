// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventEmailAddressResponseDataRelationships
type EventEmailAddressResponseDataRelationships struct {
	//
	CreatedBy EventEmailAddressResponseDataRelationshipsUser `json:"created_by"`
	//
	RevokedBy *EventEmailAddressResponseDataRelationshipsUser `json:"revoked_by,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEventEmailAddressResponseDataRelationships instantiates a new EventEmailAddressResponseDataRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEventEmailAddressResponseDataRelationships(createdBy EventEmailAddressResponseDataRelationshipsUser) *EventEmailAddressResponseDataRelationships {
	this := EventEmailAddressResponseDataRelationships{}
	this.CreatedBy = createdBy
	return &this
}

// NewEventEmailAddressResponseDataRelationshipsWithDefaults instantiates a new EventEmailAddressResponseDataRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEventEmailAddressResponseDataRelationshipsWithDefaults() *EventEmailAddressResponseDataRelationships {
	this := EventEmailAddressResponseDataRelationships{}
	return &this
}

// GetCreatedBy returns the CreatedBy field value.
func (o *EventEmailAddressResponseDataRelationships) GetCreatedBy() EventEmailAddressResponseDataRelationshipsUser {
	if o == nil {
		var ret EventEmailAddressResponseDataRelationshipsUser
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *EventEmailAddressResponseDataRelationships) GetCreatedByOk() (*EventEmailAddressResponseDataRelationshipsUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value.
func (o *EventEmailAddressResponseDataRelationships) SetCreatedBy(v EventEmailAddressResponseDataRelationshipsUser) {
	o.CreatedBy = v
}

// GetRevokedBy returns the RevokedBy field value if set, zero value otherwise.
func (o *EventEmailAddressResponseDataRelationships) GetRevokedBy() EventEmailAddressResponseDataRelationshipsUser {
	if o == nil || o.RevokedBy == nil {
		var ret EventEmailAddressResponseDataRelationshipsUser
		return ret
	}
	return *o.RevokedBy
}

// GetRevokedByOk returns a tuple with the RevokedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventEmailAddressResponseDataRelationships) GetRevokedByOk() (*EventEmailAddressResponseDataRelationshipsUser, bool) {
	if o == nil || o.RevokedBy == nil {
		return nil, false
	}
	return o.RevokedBy, true
}

// HasRevokedBy returns a boolean if a field has been set.
func (o *EventEmailAddressResponseDataRelationships) HasRevokedBy() bool {
	return o != nil && o.RevokedBy != nil
}

// SetRevokedBy gets a reference to the given EventEmailAddressResponseDataRelationshipsUser and assigns it to the RevokedBy field.
func (o *EventEmailAddressResponseDataRelationships) SetRevokedBy(v EventEmailAddressResponseDataRelationshipsUser) {
	o.RevokedBy = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EventEmailAddressResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["created_by"] = o.CreatedBy
	if o.RevokedBy != nil {
		toSerialize["revoked_by"] = o.RevokedBy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EventEmailAddressResponseDataRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedBy *EventEmailAddressResponseDataRelationshipsUser `json:"created_by"`
		RevokedBy *EventEmailAddressResponseDataRelationshipsUser `json:"revoked_by,omitempty"`
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
	if all.RevokedBy != nil && all.RevokedBy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.RevokedBy = all.RevokedBy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
