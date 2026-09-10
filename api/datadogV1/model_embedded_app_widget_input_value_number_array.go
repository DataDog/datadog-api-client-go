// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EmbeddedAppWidgetInputValueNumberArray An array of numeric values passed to the embedded app.
type EmbeddedAppWidgetInputValueNumberArray struct {
	Items []float64

	// UnparsedObject contains the raw value of the array if there was an error when deserializing into the struct
	UnparsedObject []interface{} `json:"-"`
}

// NewEmbeddedAppWidgetInputValueNumberArray instantiates a new EmbeddedAppWidgetInputValueNumberArray object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEmbeddedAppWidgetInputValueNumberArray() *EmbeddedAppWidgetInputValueNumberArray {
	this := EmbeddedAppWidgetInputValueNumberArray{}
	return &this
}

// NewEmbeddedAppWidgetInputValueNumberArrayWithDefaults instantiates a new EmbeddedAppWidgetInputValueNumberArray object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEmbeddedAppWidgetInputValueNumberArrayWithDefaults() *EmbeddedAppWidgetInputValueNumberArray {
	this := EmbeddedAppWidgetInputValueNumberArray{}
	return &this
}

// MarshalJSON serializes the struct using spec logic.
func (o EmbeddedAppWidgetInputValueNumberArray) MarshalJSON() ([]byte, error) {
	toSerialize := make([]interface{}, len(o.Items))
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EmbeddedAppWidgetInputValueNumberArray) UnmarshalJSON(bytes []byte) (err error) {
	if err = datadog.Unmarshal(bytes, &o.Items); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
