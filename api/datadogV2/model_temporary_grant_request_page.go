// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TemporaryGrantRequestPage UUID cursor pagination. A full final page can return a cursor followed by an empty page.
type TemporaryGrantRequestPage struct {
	// Use this UUID as page[cursor] to retrieve the next page. Omitted when the page is not full.
	NextCursor *uuid.UUID `json:"next_cursor,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTemporaryGrantRequestPage instantiates a new TemporaryGrantRequestPage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTemporaryGrantRequestPage() *TemporaryGrantRequestPage {
	this := TemporaryGrantRequestPage{}
	return &this
}

// NewTemporaryGrantRequestPageWithDefaults instantiates a new TemporaryGrantRequestPage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTemporaryGrantRequestPageWithDefaults() *TemporaryGrantRequestPage {
	this := TemporaryGrantRequestPage{}
	return &this
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise.
func (o *TemporaryGrantRequestPage) GetNextCursor() uuid.UUID {
	if o == nil || o.NextCursor == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.NextCursor
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemporaryGrantRequestPage) GetNextCursorOk() (*uuid.UUID, bool) {
	if o == nil || o.NextCursor == nil {
		return nil, false
	}
	return o.NextCursor, true
}

// HasNextCursor returns a boolean if a field has been set.
func (o *TemporaryGrantRequestPage) HasNextCursor() bool {
	return o != nil && o.NextCursor != nil
}

// SetNextCursor gets a reference to the given uuid.UUID and assigns it to the NextCursor field.
func (o *TemporaryGrantRequestPage) SetNextCursor(v uuid.UUID) {
	o.NextCursor = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TemporaryGrantRequestPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.NextCursor != nil {
		toSerialize["next_cursor"] = o.NextCursor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TemporaryGrantRequestPage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NextCursor *uuid.UUID `json:"next_cursor,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"next_cursor"})
	} else {
		return err
	}
	o.NextCursor = all.NextCursor

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
