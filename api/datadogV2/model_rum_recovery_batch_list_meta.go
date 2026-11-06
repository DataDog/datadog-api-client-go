// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRecoveryBatchListMeta Batch-list pagination metadata.
type RumRecoveryBatchListMeta struct {
	// Opaque next-page cursor. Omitted on the final page.
	NextCursor *string `json:"next_cursor,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumRecoveryBatchListMeta instantiates a new RumRecoveryBatchListMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumRecoveryBatchListMeta() *RumRecoveryBatchListMeta {
	this := RumRecoveryBatchListMeta{}
	return &this
}

// NewRumRecoveryBatchListMetaWithDefaults instantiates a new RumRecoveryBatchListMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumRecoveryBatchListMetaWithDefaults() *RumRecoveryBatchListMeta {
	this := RumRecoveryBatchListMeta{}
	return &this
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise.
func (o *RumRecoveryBatchListMeta) GetNextCursor() string {
	if o == nil || o.NextCursor == nil {
		var ret string
		return ret
	}
	return *o.NextCursor
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumRecoveryBatchListMeta) GetNextCursorOk() (*string, bool) {
	if o == nil || o.NextCursor == nil {
		return nil, false
	}
	return o.NextCursor, true
}

// HasNextCursor returns a boolean if a field has been set.
func (o *RumRecoveryBatchListMeta) HasNextCursor() bool {
	return o != nil && o.NextCursor != nil
}

// SetNextCursor gets a reference to the given string and assigns it to the NextCursor field.
func (o *RumRecoveryBatchListMeta) SetNextCursor(v string) {
	o.NextCursor = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumRecoveryBatchListMeta) MarshalJSON() ([]byte, error) {
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
func (o *RumRecoveryBatchListMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NextCursor *string `json:"next_cursor,omitempty"`
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
