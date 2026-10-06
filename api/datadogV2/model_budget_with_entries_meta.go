// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BudgetWithEntriesMeta Additional information about errors encountered while retrieving budget cost data.
type BudgetWithEntriesMeta struct {
	// A user-facing explanation of why budget cost data could not be retrieved.
	Error string `json:"error"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBudgetWithEntriesMeta instantiates a new BudgetWithEntriesMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBudgetWithEntriesMeta(error string) *BudgetWithEntriesMeta {
	this := BudgetWithEntriesMeta{}
	this.Error = error
	return &this
}

// NewBudgetWithEntriesMetaWithDefaults instantiates a new BudgetWithEntriesMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBudgetWithEntriesMetaWithDefaults() *BudgetWithEntriesMeta {
	this := BudgetWithEntriesMeta{}
	return &this
}

// GetError returns the Error field value.
func (o *BudgetWithEntriesMeta) GetError() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *BudgetWithEntriesMeta) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value.
func (o *BudgetWithEntriesMeta) SetError(v string) {
	o.Error = v
}

// MarshalJSON serializes the struct using spec logic.
func (o BudgetWithEntriesMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["error"] = o.Error

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BudgetWithEntriesMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Error *string `json:"error"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Error == nil {
		return fmt.Errorf("required field error missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"error"})
	} else {
		return err
	}
	o.Error = *all.Error

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
