// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RecommendationsStreamRequestColumnsItems Test-only generated schema for the dashboard versioning proof of concept.
type RecommendationsStreamRequestColumnsItems struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Alias *string `json:"alias,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Field string `json:"field"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Width RecommendationsStreamRequestColumnsItems20270101Width `json:"width"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewRecommendationsStreamRequestColumnsItems instantiates a new RecommendationsStreamRequestColumnsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRecommendationsStreamRequestColumnsItems(field string, width RecommendationsStreamRequestColumnsItems20270101Width) *RecommendationsStreamRequestColumnsItems {
	this := RecommendationsStreamRequestColumnsItems{}
	this.Field = field
	this.Width = width
	return &this
}

// NewRecommendationsStreamRequestColumnsItemsWithDefaults instantiates a new RecommendationsStreamRequestColumnsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRecommendationsStreamRequestColumnsItemsWithDefaults() *RecommendationsStreamRequestColumnsItems {
	this := RecommendationsStreamRequestColumnsItems{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *RecommendationsStreamRequestColumnsItems) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsStreamRequestColumnsItems) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *RecommendationsStreamRequestColumnsItems) HasAlias() bool {
	return o != nil && o.Alias != nil
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *RecommendationsStreamRequestColumnsItems) SetAlias(v string) {
	o.Alias = &v
}

// GetField returns the Field field value.
func (o *RecommendationsStreamRequestColumnsItems) GetField() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *RecommendationsStreamRequestColumnsItems) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value.
func (o *RecommendationsStreamRequestColumnsItems) SetField(v string) {
	o.Field = v
}

// GetWidth returns the Width field value.
func (o *RecommendationsStreamRequestColumnsItems) GetWidth() RecommendationsStreamRequestColumnsItems20270101Width {
	if o == nil {
		var ret RecommendationsStreamRequestColumnsItems20270101Width
		return ret
	}
	return o.Width
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
func (o *RecommendationsStreamRequestColumnsItems) GetWidthOk() (*RecommendationsStreamRequestColumnsItems20270101Width, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Width, true
}

// SetWidth sets field value.
func (o *RecommendationsStreamRequestColumnsItems) SetWidth(v RecommendationsStreamRequestColumnsItems20270101Width) {
	o.Width = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RecommendationsStreamRequestColumnsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
	}
	toSerialize["field"] = o.Field
	toSerialize["width"] = o.Width
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RecommendationsStreamRequestColumnsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Alias *string                                                `json:"alias,omitempty"`
		Field *string                                                `json:"field"`
		Width *RecommendationsStreamRequestColumnsItems20270101Width `json:"width"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Field == nil {
		return fmt.Errorf("required field field missing")
	}
	if all.Width == nil {
		return fmt.Errorf("required field width missing")
	}
	o.Alias = all.Alias
	o.Field = *all.Field
	o.Width = *all.Width

	return nil
}
