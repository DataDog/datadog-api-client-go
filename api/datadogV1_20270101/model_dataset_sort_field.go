// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatasetSortField Test-only generated schema for the dashboard versioning proof of concept.
type DatasetSortField struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Order QuerySortOrder `json:"order"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewDatasetSortField instantiates a new DatasetSortField object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatasetSortField(name string, order QuerySortOrder) *DatasetSortField {
	this := DatasetSortField{}
	this.Name = name
	this.Order = order
	return &this
}

// NewDatasetSortFieldWithDefaults instantiates a new DatasetSortField object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatasetSortFieldWithDefaults() *DatasetSortField {
	this := DatasetSortField{}
	return &this
}

// GetName returns the Name field value.
func (o *DatasetSortField) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DatasetSortField) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *DatasetSortField) SetName(v string) {
	o.Name = v
}

// GetOrder returns the Order field value.
func (o *DatasetSortField) GetOrder() QuerySortOrder {
	if o == nil {
		var ret QuerySortOrder
		return ret
	}
	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *DatasetSortField) GetOrderOk() (*QuerySortOrder, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value.
func (o *DatasetSortField) SetOrder(v QuerySortOrder) {
	o.Order = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatasetSortField) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["order"] = o.Order
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DatasetSortField) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name  *string         `json:"name"`
		Order *QuerySortOrder `json:"order"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Order == nil {
		return fmt.Errorf("required field order missing")
	}

	hasInvalidField := false
	o.Name = *all.Name
	if !all.Order.IsValid() {
		hasInvalidField = true
	} else {
		o.Order = *all.Order
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
