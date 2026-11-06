// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumArchiveSessionsSearchAttributes Archived-session query. A window wider than configured archive retention is narrowed before searching.
type RumArchiveSessionsSearchAttributes struct {
	// Opaque next-page cursor. Keep the query, time range, and hide_retained value unchanged.
	Cursor *string `json:"cursor,omitempty"`
	// Start of the range in Unix milliseconds. Must be positive and earlier than to.
	From int64 `json:"from"`
	// When true, exclude sessions already present in hot RUM storage. When false, those sessions can appear in the results. Keep this value unchanged when continuing with a cursor.
	HideRetained *bool `json:"hide_retained,omitempty"`
	// Page size, at most 1000 (default 50). Values at or below zero use 50.
	Limit *int32 `json:"limit,omitempty"`
	// RUM search query. Must not be empty.
	Query string `json:"query"`
	// End of the range in Unix milliseconds.
	To int64 `json:"to"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumArchiveSessionsSearchAttributes instantiates a new RumArchiveSessionsSearchAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumArchiveSessionsSearchAttributes(from int64, query string, to int64) *RumArchiveSessionsSearchAttributes {
	this := RumArchiveSessionsSearchAttributes{}
	this.From = from
	var hideRetained bool = false
	this.HideRetained = &hideRetained
	var limit int32 = 50
	this.Limit = &limit
	this.Query = query
	this.To = to
	return &this
}

// NewRumArchiveSessionsSearchAttributesWithDefaults instantiates a new RumArchiveSessionsSearchAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumArchiveSessionsSearchAttributesWithDefaults() *RumArchiveSessionsSearchAttributes {
	this := RumArchiveSessionsSearchAttributes{}
	var hideRetained bool = false
	this.HideRetained = &hideRetained
	var limit int32 = 50
	this.Limit = &limit
	return &this
}

// GetCursor returns the Cursor field value if set, zero value otherwise.
func (o *RumArchiveSessionsSearchAttributes) GetCursor() string {
	if o == nil || o.Cursor == nil {
		var ret string
		return ret
	}
	return *o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumArchiveSessionsSearchAttributes) GetCursorOk() (*string, bool) {
	if o == nil || o.Cursor == nil {
		return nil, false
	}
	return o.Cursor, true
}

// HasCursor returns a boolean if a field has been set.
func (o *RumArchiveSessionsSearchAttributes) HasCursor() bool {
	return o != nil && o.Cursor != nil
}

// SetCursor gets a reference to the given string and assigns it to the Cursor field.
func (o *RumArchiveSessionsSearchAttributes) SetCursor(v string) {
	o.Cursor = &v
}

// GetFrom returns the From field value.
func (o *RumArchiveSessionsSearchAttributes) GetFrom() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *RumArchiveSessionsSearchAttributes) GetFromOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value.
func (o *RumArchiveSessionsSearchAttributes) SetFrom(v int64) {
	o.From = v
}

// GetHideRetained returns the HideRetained field value if set, zero value otherwise.
func (o *RumArchiveSessionsSearchAttributes) GetHideRetained() bool {
	if o == nil || o.HideRetained == nil {
		var ret bool
		return ret
	}
	return *o.HideRetained
}

// GetHideRetainedOk returns a tuple with the HideRetained field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumArchiveSessionsSearchAttributes) GetHideRetainedOk() (*bool, bool) {
	if o == nil || o.HideRetained == nil {
		return nil, false
	}
	return o.HideRetained, true
}

// HasHideRetained returns a boolean if a field has been set.
func (o *RumArchiveSessionsSearchAttributes) HasHideRetained() bool {
	return o != nil && o.HideRetained != nil
}

// SetHideRetained gets a reference to the given bool and assigns it to the HideRetained field.
func (o *RumArchiveSessionsSearchAttributes) SetHideRetained(v bool) {
	o.HideRetained = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *RumArchiveSessionsSearchAttributes) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumArchiveSessionsSearchAttributes) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *RumArchiveSessionsSearchAttributes) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *RumArchiveSessionsSearchAttributes) SetLimit(v int32) {
	o.Limit = &v
}

// GetQuery returns the Query field value.
func (o *RumArchiveSessionsSearchAttributes) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *RumArchiveSessionsSearchAttributes) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *RumArchiveSessionsSearchAttributes) SetQuery(v string) {
	o.Query = v
}

// GetTo returns the To field value.
func (o *RumArchiveSessionsSearchAttributes) GetTo() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *RumArchiveSessionsSearchAttributes) GetToOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value.
func (o *RumArchiveSessionsSearchAttributes) SetTo(v int64) {
	o.To = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumArchiveSessionsSearchAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Cursor != nil {
		toSerialize["cursor"] = o.Cursor
	}
	toSerialize["from"] = o.From
	if o.HideRetained != nil {
		toSerialize["hide_retained"] = o.HideRetained
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	toSerialize["query"] = o.Query
	toSerialize["to"] = o.To

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumArchiveSessionsSearchAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Cursor       *string `json:"cursor,omitempty"`
		From         *int64  `json:"from"`
		HideRetained *bool   `json:"hide_retained,omitempty"`
		Limit        *int32  `json:"limit,omitempty"`
		Query        *string `json:"query"`
		To           *int64  `json:"to"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.From == nil {
		return fmt.Errorf("required field from missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	if all.To == nil {
		return fmt.Errorf("required field to missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cursor", "from", "hide_retained", "limit", "query", "to"})
	} else {
		return err
	}
	o.Cursor = all.Cursor
	o.From = *all.From
	o.HideRetained = all.HideRetained
	o.Limit = all.Limit
	o.Query = *all.Query
	o.To = *all.To

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
