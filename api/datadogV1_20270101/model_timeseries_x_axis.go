// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesXAxis Test-only generated schema for the dashboard versioning proof of concept.
type TimeseriesXAxis struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	DateFormat *TimeseriesXAxisDateFormat `json:"date_format,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Label *string `json:"label,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTimeseriesXAxis instantiates a new TimeseriesXAxis object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTimeseriesXAxis() *TimeseriesXAxis {
	this := TimeseriesXAxis{}
	return &this
}

// NewTimeseriesXAxisWithDefaults instantiates a new TimeseriesXAxis object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTimeseriesXAxisWithDefaults() *TimeseriesXAxis {
	this := TimeseriesXAxis{}
	return &this
}

// GetDateFormat returns the DateFormat field value if set, zero value otherwise.
func (o *TimeseriesXAxis) GetDateFormat() TimeseriesXAxisDateFormat {
	if o == nil || o.DateFormat == nil {
		var ret TimeseriesXAxisDateFormat
		return ret
	}
	return *o.DateFormat
}

// GetDateFormatOk returns a tuple with the DateFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesXAxis) GetDateFormatOk() (*TimeseriesXAxisDateFormat, bool) {
	if o == nil || o.DateFormat == nil {
		return nil, false
	}
	return o.DateFormat, true
}

// HasDateFormat returns a boolean if a field has been set.
func (o *TimeseriesXAxis) HasDateFormat() bool {
	return o != nil && o.DateFormat != nil
}

// SetDateFormat gets a reference to the given TimeseriesXAxisDateFormat and assigns it to the DateFormat field.
func (o *TimeseriesXAxis) SetDateFormat(v TimeseriesXAxisDateFormat) {
	o.DateFormat = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *TimeseriesXAxis) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesXAxis) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *TimeseriesXAxis) HasLabel() bool {
	return o != nil && o.Label != nil
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *TimeseriesXAxis) SetLabel(v string) {
	o.Label = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TimeseriesXAxis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DateFormat != nil {
		toSerialize["date_format"] = o.DateFormat
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TimeseriesXAxis) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DateFormat *TimeseriesXAxisDateFormat `json:"date_format,omitempty"`
		Label      *string                    `json:"label,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	hasInvalidField := false
	if all.DateFormat != nil && all.DateFormat.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DateFormat = all.DateFormat
	o.Label = all.Label

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
