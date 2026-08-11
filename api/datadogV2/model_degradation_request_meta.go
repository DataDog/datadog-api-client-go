// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DegradationRequestMeta The supported metadata for a degradation request.
type DegradationRequestMeta struct {
	// A unique key used to ensure idempotent requests.
	IdempotencyKey *uuid.UUID `json:"idempotency_key,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDegradationRequestMeta instantiates a new DegradationRequestMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDegradationRequestMeta() *DegradationRequestMeta {
	this := DegradationRequestMeta{}
	return &this
}

// NewDegradationRequestMetaWithDefaults instantiates a new DegradationRequestMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDegradationRequestMetaWithDefaults() *DegradationRequestMeta {
	this := DegradationRequestMeta{}
	return &this
}

// GetIdempotencyKey returns the IdempotencyKey field value if set, zero value otherwise.
func (o *DegradationRequestMeta) GetIdempotencyKey() uuid.UUID {
	if o == nil || o.IdempotencyKey == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.IdempotencyKey
}

// GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DegradationRequestMeta) GetIdempotencyKeyOk() (*uuid.UUID, bool) {
	if o == nil || o.IdempotencyKey == nil {
		return nil, false
	}
	return o.IdempotencyKey, true
}

// HasIdempotencyKey returns a boolean if a field has been set.
func (o *DegradationRequestMeta) HasIdempotencyKey() bool {
	return o != nil && o.IdempotencyKey != nil
}

// SetIdempotencyKey gets a reference to the given uuid.UUID and assigns it to the IdempotencyKey field.
func (o *DegradationRequestMeta) SetIdempotencyKey(v uuid.UUID) {
	o.IdempotencyKey = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DegradationRequestMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.IdempotencyKey != nil {
		toSerialize["idempotency_key"] = o.IdempotencyKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DegradationRequestMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IdempotencyKey *uuid.UUID `json:"idempotency_key,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"idempotency_key"})
	} else {
		return err
	}
	o.IdempotencyKey = all.IdempotencyKey

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
