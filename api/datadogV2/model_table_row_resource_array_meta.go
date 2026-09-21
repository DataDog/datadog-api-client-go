// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableRowResourceArrayMeta Metadata about the rows requested, including which ones were not found.
type TableRowResourceArrayMeta struct {
	// Number of requested rows that were found and returned in `data`.
	FoundCount int64 `json:"found_count"`
	// Row IDs from the request that do not exist in the reference table. Empty when every requested row was found.
	NotFound []string `json:"not_found"`
	// Number of row IDs supplied in the `row_id` query parameter.
	RequestedCount int64 `json:"requested_count"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTableRowResourceArrayMeta instantiates a new TableRowResourceArrayMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTableRowResourceArrayMeta(foundCount int64, notFound []string, requestedCount int64) *TableRowResourceArrayMeta {
	this := TableRowResourceArrayMeta{}
	this.FoundCount = foundCount
	this.NotFound = notFound
	this.RequestedCount = requestedCount
	return &this
}

// NewTableRowResourceArrayMetaWithDefaults instantiates a new TableRowResourceArrayMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTableRowResourceArrayMetaWithDefaults() *TableRowResourceArrayMeta {
	this := TableRowResourceArrayMeta{}
	return &this
}

// GetFoundCount returns the FoundCount field value.
func (o *TableRowResourceArrayMeta) GetFoundCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.FoundCount
}

// GetFoundCountOk returns a tuple with the FoundCount field value
// and a boolean to check if the value has been set.
func (o *TableRowResourceArrayMeta) GetFoundCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FoundCount, true
}

// SetFoundCount sets field value.
func (o *TableRowResourceArrayMeta) SetFoundCount(v int64) {
	o.FoundCount = v
}

// GetNotFound returns the NotFound field value.
func (o *TableRowResourceArrayMeta) GetNotFound() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.NotFound
}

// GetNotFoundOk returns a tuple with the NotFound field value
// and a boolean to check if the value has been set.
func (o *TableRowResourceArrayMeta) GetNotFoundOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotFound, true
}

// SetNotFound sets field value.
func (o *TableRowResourceArrayMeta) SetNotFound(v []string) {
	o.NotFound = v
}

// GetRequestedCount returns the RequestedCount field value.
func (o *TableRowResourceArrayMeta) GetRequestedCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.RequestedCount
}

// GetRequestedCountOk returns a tuple with the RequestedCount field value
// and a boolean to check if the value has been set.
func (o *TableRowResourceArrayMeta) GetRequestedCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestedCount, true
}

// SetRequestedCount sets field value.
func (o *TableRowResourceArrayMeta) SetRequestedCount(v int64) {
	o.RequestedCount = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TableRowResourceArrayMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["found_count"] = o.FoundCount
	toSerialize["not_found"] = o.NotFound
	toSerialize["requested_count"] = o.RequestedCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TableRowResourceArrayMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FoundCount     *int64    `json:"found_count"`
		NotFound       *[]string `json:"not_found"`
		RequestedCount *int64    `json:"requested_count"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.FoundCount == nil {
		return fmt.Errorf("required field found_count missing")
	}
	if all.NotFound == nil {
		return fmt.Errorf("required field not_found missing")
	}
	if all.RequestedCount == nil {
		return fmt.Errorf("required field requested_count missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"found_count", "not_found", "requested_count"})
	} else {
		return err
	}
	o.FoundCount = *all.FoundCount
	o.NotFound = *all.NotFound
	o.RequestedCount = *all.RequestedCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
