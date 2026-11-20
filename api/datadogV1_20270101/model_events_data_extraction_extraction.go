// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventsDataExtractionExtraction Test-only generated schema for the dashboard versioning proof of concept.
type EventsDataExtractionExtraction struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Pattern string `json:"pattern"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type EventsDataExtractionExtractionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewEventsDataExtractionExtraction instantiates a new EventsDataExtractionExtraction object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEventsDataExtractionExtraction(pattern string, typeVar EventsDataExtractionExtractionType) *EventsDataExtractionExtraction {
	this := EventsDataExtractionExtraction{}
	this.Pattern = pattern
	this.Type = typeVar
	return &this
}

// NewEventsDataExtractionExtractionWithDefaults instantiates a new EventsDataExtractionExtraction object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEventsDataExtractionExtractionWithDefaults() *EventsDataExtractionExtraction {
	this := EventsDataExtractionExtraction{}
	return &this
}

// GetPattern returns the Pattern field value.
func (o *EventsDataExtractionExtraction) GetPattern() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value
// and a boolean to check if the value has been set.
func (o *EventsDataExtractionExtraction) GetPatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pattern, true
}

// SetPattern sets field value.
func (o *EventsDataExtractionExtraction) SetPattern(v string) {
	o.Pattern = v
}

// GetType returns the Type field value.
func (o *EventsDataExtractionExtraction) GetType() EventsDataExtractionExtractionType {
	if o == nil {
		var ret EventsDataExtractionExtractionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EventsDataExtractionExtraction) GetTypeOk() (*EventsDataExtractionExtractionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *EventsDataExtractionExtraction) SetType(v EventsDataExtractionExtractionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EventsDataExtractionExtraction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["pattern"] = o.Pattern
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EventsDataExtractionExtraction) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Pattern *string                             `json:"pattern"`
		Type    *EventsDataExtractionExtractionType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Pattern == nil {
		return fmt.Errorf("required field pattern missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.Pattern = *all.Pattern
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
