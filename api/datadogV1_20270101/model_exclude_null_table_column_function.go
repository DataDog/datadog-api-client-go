// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExcludeNullTableColumnFunction Test-only generated schema for the dashboard versioning proof of concept.
type ExcludeNullTableColumnFunction struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name ExcludeNullTableColumnFunctionName `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewExcludeNullTableColumnFunction instantiates a new ExcludeNullTableColumnFunction object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExcludeNullTableColumnFunction(name ExcludeNullTableColumnFunctionName) *ExcludeNullTableColumnFunction {
	this := ExcludeNullTableColumnFunction{}
	this.Name = name
	return &this
}

// NewExcludeNullTableColumnFunctionWithDefaults instantiates a new ExcludeNullTableColumnFunction object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExcludeNullTableColumnFunctionWithDefaults() *ExcludeNullTableColumnFunction {
	this := ExcludeNullTableColumnFunction{}
	return &this
}

// GetName returns the Name field value.
func (o *ExcludeNullTableColumnFunction) GetName() ExcludeNullTableColumnFunctionName {
	if o == nil {
		var ret ExcludeNullTableColumnFunctionName
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExcludeNullTableColumnFunction) GetNameOk() (*ExcludeNullTableColumnFunctionName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ExcludeNullTableColumnFunction) SetName(v ExcludeNullTableColumnFunctionName) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExcludeNullTableColumnFunction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExcludeNullTableColumnFunction) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name *ExcludeNullTableColumnFunctionName `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}

	hasInvalidField := false
	if !all.Name.IsValid() {
		hasInvalidField = true
	} else {
		o.Name = *all.Name
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
