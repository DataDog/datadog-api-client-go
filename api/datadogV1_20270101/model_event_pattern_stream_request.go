// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventPatternStreamRequest Test-only generated schema for the dashboard versioning proof of concept.
type EventPatternStreamRequest struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Columns []EventPatternStreamRequestColumnsItems `json:"columns"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Query EventPatternStreamQuery `json:"query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ResponseFormat EventPatternStreamRequestResponseFormat `json:"response_format"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewEventPatternStreamRequest instantiates a new EventPatternStreamRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEventPatternStreamRequest(columns []EventPatternStreamRequestColumnsItems, query EventPatternStreamQuery, responseFormat EventPatternStreamRequestResponseFormat) *EventPatternStreamRequest {
	this := EventPatternStreamRequest{}
	this.Columns = columns
	this.Query = query
	this.ResponseFormat = responseFormat
	return &this
}

// NewEventPatternStreamRequestWithDefaults instantiates a new EventPatternStreamRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEventPatternStreamRequestWithDefaults() *EventPatternStreamRequest {
	this := EventPatternStreamRequest{}
	return &this
}

// GetColumns returns the Columns field value.
func (o *EventPatternStreamRequest) GetColumns() []EventPatternStreamRequestColumnsItems {
	if o == nil {
		var ret []EventPatternStreamRequestColumnsItems
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value
// and a boolean to check if the value has been set.
func (o *EventPatternStreamRequest) GetColumnsOk() (*[]EventPatternStreamRequestColumnsItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Columns, true
}

// SetColumns sets field value.
func (o *EventPatternStreamRequest) SetColumns(v []EventPatternStreamRequestColumnsItems) {
	o.Columns = v
}

// GetQuery returns the Query field value.
func (o *EventPatternStreamRequest) GetQuery() EventPatternStreamQuery {
	if o == nil {
		var ret EventPatternStreamQuery
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *EventPatternStreamRequest) GetQueryOk() (*EventPatternStreamQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *EventPatternStreamRequest) SetQuery(v EventPatternStreamQuery) {
	o.Query = v
}

// GetResponseFormat returns the ResponseFormat field value.
func (o *EventPatternStreamRequest) GetResponseFormat() EventPatternStreamRequestResponseFormat {
	if o == nil {
		var ret EventPatternStreamRequestResponseFormat
		return ret
	}
	return o.ResponseFormat
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value
// and a boolean to check if the value has been set.
func (o *EventPatternStreamRequest) GetResponseFormatOk() (*EventPatternStreamRequestResponseFormat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseFormat, true
}

// SetResponseFormat sets field value.
func (o *EventPatternStreamRequest) SetResponseFormat(v EventPatternStreamRequestResponseFormat) {
	o.ResponseFormat = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EventPatternStreamRequest) MarshalJSON() ([]byte, error) {
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
func (o *EventPatternStreamRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Columns        *[]EventPatternStreamRequestColumnsItems `json:"columns"`
		Query          *EventPatternStreamQuery                 `json:"query"`
		ResponseFormat *EventPatternStreamRequestResponseFormat `json:"response_format"`
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
