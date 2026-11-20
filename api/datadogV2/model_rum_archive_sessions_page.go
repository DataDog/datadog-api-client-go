// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMArchiveSessionsPage Search pagination.
type RUMArchiveSessionsPage struct {
	// Opaque continuation cursor. Omitted when no further page is available.
	NextCursor *string `json:"next_cursor,omitempty"`
	// Search result count. Can be zero on continuation pages even when sessions are returned. Use `next_cursor` to determine whether more pages are available.
	TotalCount int64 `json:"total_count"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRUMArchiveSessionsPage instantiates a new RUMArchiveSessionsPage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRUMArchiveSessionsPage(totalCount int64) *RUMArchiveSessionsPage {
	this := RUMArchiveSessionsPage{}
	this.TotalCount = totalCount
	return &this
}

// NewRUMArchiveSessionsPageWithDefaults instantiates a new RUMArchiveSessionsPage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRUMArchiveSessionsPageWithDefaults() *RUMArchiveSessionsPage {
	this := RUMArchiveSessionsPage{}
	return &this
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise.
func (o *RUMArchiveSessionsPage) GetNextCursor() string {
	if o == nil || o.NextCursor == nil {
		var ret string
		return ret
	}
	return *o.NextCursor
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMArchiveSessionsPage) GetNextCursorOk() (*string, bool) {
	if o == nil || o.NextCursor == nil {
		return nil, false
	}
	return o.NextCursor, true
}

// HasNextCursor returns a boolean if a field has been set.
func (o *RUMArchiveSessionsPage) HasNextCursor() bool {
	return o != nil && o.NextCursor != nil
}

// SetNextCursor gets a reference to the given string and assigns it to the NextCursor field.
func (o *RUMArchiveSessionsPage) SetNextCursor(v string) {
	o.NextCursor = &v
}

// GetTotalCount returns the TotalCount field value.
func (o *RUMArchiveSessionsPage) GetTotalCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *RUMArchiveSessionsPage) GetTotalCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value.
func (o *RUMArchiveSessionsPage) SetTotalCount(v int64) {
	o.TotalCount = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RUMArchiveSessionsPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.NextCursor != nil {
		toSerialize["next_cursor"] = o.NextCursor
	}
	toSerialize["total_count"] = o.TotalCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RUMArchiveSessionsPage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NextCursor *string `json:"next_cursor,omitempty"`
		TotalCount *int64  `json:"total_count"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.TotalCount == nil {
		return fmt.Errorf("required field total_count missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"next_cursor", "total_count"})
	} else {
		return err
	}
	o.NextCursor = all.NextCursor
	o.TotalCount = *all.TotalCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
