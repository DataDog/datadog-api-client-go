// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumStreamRequest Test-only generated schema for the dashboard versioning proof of concept.
type RumStreamRequest struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Columns []RumStreamRequestColumnsItems `json:"columns"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Query RumStreamQuery `json:"query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ResponseFormat RumStreamRequestResponseFormat `json:"response_format"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewRumStreamRequest instantiates a new RumStreamRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumStreamRequest(columns []RumStreamRequestColumnsItems, query RumStreamQuery, responseFormat RumStreamRequestResponseFormat) *RumStreamRequest {
	this := RumStreamRequest{}
	this.Columns = columns
	this.Query = query
	this.ResponseFormat = responseFormat
	return &this
}

// NewRumStreamRequestWithDefaults instantiates a new RumStreamRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumStreamRequestWithDefaults() *RumStreamRequest {
	this := RumStreamRequest{}
	return &this
}

// GetColumns returns the Columns field value.
func (o *RumStreamRequest) GetColumns() []RumStreamRequestColumnsItems {
	if o == nil {
		var ret []RumStreamRequestColumnsItems
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value
// and a boolean to check if the value has been set.
func (o *RumStreamRequest) GetColumnsOk() (*[]RumStreamRequestColumnsItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Columns, true
}

// SetColumns sets field value.
func (o *RumStreamRequest) SetColumns(v []RumStreamRequestColumnsItems) {
	o.Columns = v
}

// GetQuery returns the Query field value.
func (o *RumStreamRequest) GetQuery() RumStreamQuery {
	if o == nil {
		var ret RumStreamQuery
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *RumStreamRequest) GetQueryOk() (*RumStreamQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *RumStreamRequest) SetQuery(v RumStreamQuery) {
	o.Query = v
}

// GetResponseFormat returns the ResponseFormat field value.
func (o *RumStreamRequest) GetResponseFormat() RumStreamRequestResponseFormat {
	if o == nil {
		var ret RumStreamRequestResponseFormat
		return ret
	}
	return o.ResponseFormat
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value
// and a boolean to check if the value has been set.
func (o *RumStreamRequest) GetResponseFormatOk() (*RumStreamRequestResponseFormat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseFormat, true
}

// SetResponseFormat sets field value.
func (o *RumStreamRequest) SetResponseFormat(v RumStreamRequestResponseFormat) {
	o.ResponseFormat = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumStreamRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["columns"] = o.Columns
	toSerialize["query"] = o.Query
	toSerialize["response_format"] = o.ResponseFormat
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumStreamRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Columns        *[]RumStreamRequestColumnsItems `json:"columns"`
		Query          *RumStreamQuery                 `json:"query"`
		ResponseFormat *RumStreamRequestResponseFormat `json:"response_format"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Columns == nil {
		return fmt.Errorf("required field columns missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	if all.ResponseFormat == nil {
		return fmt.Errorf("required field response_format missing")
	}

	hasInvalidField := false
	o.Columns = *all.Columns
	if all.Query.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Query = *all.Query
	if !all.ResponseFormat.IsValid() {
		hasInvalidField = true
	} else {
		o.ResponseFormat = *all.ResponseFormat
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
