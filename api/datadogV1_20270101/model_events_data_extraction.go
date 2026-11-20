// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventsDataExtraction Test-only generated schema for the dashboard versioning proof of concept.
type EventsDataExtraction struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Column string `json:"column"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Extraction EventsDataExtractionExtraction `json:"extraction"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewEventsDataExtraction instantiates a new EventsDataExtraction object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEventsDataExtraction(column string, extraction EventsDataExtractionExtraction) *EventsDataExtraction {
	this := EventsDataExtraction{}
	this.Column = column
	this.Extraction = extraction
	return &this
}

// NewEventsDataExtractionWithDefaults instantiates a new EventsDataExtraction object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEventsDataExtractionWithDefaults() *EventsDataExtraction {
	this := EventsDataExtraction{}
	return &this
}

// GetColumn returns the Column field value.
func (o *EventsDataExtraction) GetColumn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Column
}

// GetColumnOk returns a tuple with the Column field value
// and a boolean to check if the value has been set.
func (o *EventsDataExtraction) GetColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Column, true
}

// SetColumn sets field value.
func (o *EventsDataExtraction) SetColumn(v string) {
	o.Column = v
}

// GetExtraction returns the Extraction field value.
func (o *EventsDataExtraction) GetExtraction() EventsDataExtractionExtraction {
	if o == nil {
		var ret EventsDataExtractionExtraction
		return ret
	}
	return o.Extraction
}

// GetExtractionOk returns a tuple with the Extraction field value
// and a boolean to check if the value has been set.
func (o *EventsDataExtraction) GetExtractionOk() (*EventsDataExtractionExtraction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Extraction, true
}

// SetExtraction sets field value.
func (o *EventsDataExtraction) SetExtraction(v EventsDataExtractionExtraction) {
	o.Extraction = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EventsDataExtraction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["column"] = o.Column
	toSerialize["extraction"] = o.Extraction
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EventsDataExtraction) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Column     *string                         `json:"column"`
		Extraction *EventsDataExtractionExtraction `json:"extraction"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Column == nil {
		return fmt.Errorf("required field column missing")
	}
	if all.Extraction == nil {
		return fmt.Errorf("required field extraction missing")
	}

	hasInvalidField := false
	o.Column = *all.Column
	if all.Extraction.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Extraction = *all.Extraction

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
