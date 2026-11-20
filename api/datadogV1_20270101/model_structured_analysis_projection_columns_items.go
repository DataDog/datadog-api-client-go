// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StructuredAnalysisProjectionColumnsItems Test-only generated schema for the dashboard versioning proof of concept.
type StructuredAnalysisProjectionColumnsItems struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Alias *string `json:"alias,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Column string `json:"column"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewStructuredAnalysisProjectionColumnsItems instantiates a new StructuredAnalysisProjectionColumnsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStructuredAnalysisProjectionColumnsItems(column string) *StructuredAnalysisProjectionColumnsItems {
	this := StructuredAnalysisProjectionColumnsItems{}
	this.Column = column
	return &this
}

// NewStructuredAnalysisProjectionColumnsItemsWithDefaults instantiates a new StructuredAnalysisProjectionColumnsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStructuredAnalysisProjectionColumnsItemsWithDefaults() *StructuredAnalysisProjectionColumnsItems {
	this := StructuredAnalysisProjectionColumnsItems{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *StructuredAnalysisProjectionColumnsItems) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredAnalysisProjectionColumnsItems) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *StructuredAnalysisProjectionColumnsItems) HasAlias() bool {
	return o != nil && o.Alias != nil
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *StructuredAnalysisProjectionColumnsItems) SetAlias(v string) {
	o.Alias = &v
}

// GetColumn returns the Column field value.
func (o *StructuredAnalysisProjectionColumnsItems) GetColumn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Column
}

// GetColumnOk returns a tuple with the Column field value
// and a boolean to check if the value has been set.
func (o *StructuredAnalysisProjectionColumnsItems) GetColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Column, true
}

// SetColumn sets field value.
func (o *StructuredAnalysisProjectionColumnsItems) SetColumn(v string) {
	o.Column = v
}

// MarshalJSON serializes the struct using spec logic.
func (o StructuredAnalysisProjectionColumnsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
	}
	toSerialize["column"] = o.Column
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StructuredAnalysisProjectionColumnsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Alias  *string `json:"alias,omitempty"`
		Column *string `json:"column"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Column == nil {
		return fmt.Errorf("required field column missing")
	}
	o.Alias = all.Alias
	o.Column = *all.Column

	return nil
}
