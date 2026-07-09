// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatasetListQuerySortField A single sort directive for a `DatasetListQuery`.
type DatasetListQuerySortField struct {
	// Name of the field to sort on.
	Name string `json:"name"`
	// Direction of sort.
	Order QuerySortOrder `json:"order"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDatasetListQuerySortField instantiates a new DatasetListQuerySortField object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatasetListQuerySortField(name string, order QuerySortOrder) *DatasetListQuerySortField {
	this := DatasetListQuerySortField{}
	this.Name = name
	this.Order = order
	return &this
}

// NewDatasetListQuerySortFieldWithDefaults instantiates a new DatasetListQuerySortField object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatasetListQuerySortFieldWithDefaults() *DatasetListQuerySortField {
	this := DatasetListQuerySortField{}
	var order QuerySortOrder = QUERYSORTORDER_DESC
	this.Order = order
	return &this
}

// GetName returns the Name field value.
func (o *DatasetListQuerySortField) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DatasetListQuerySortField) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *DatasetListQuerySortField) SetName(v string) {
	o.Name = v
}

// GetOrder returns the Order field value.
func (o *DatasetListQuerySortField) GetOrder() QuerySortOrder {
	if o == nil {
		var ret QuerySortOrder
		return ret
	}
	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *DatasetListQuerySortField) GetOrderOk() (*QuerySortOrder, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value.
func (o *DatasetListQuerySortField) SetOrder(v QuerySortOrder) {
	o.Order = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatasetListQuerySortField) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["order"] = o.Order

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DatasetListQuerySortField) UnmarshalJSON(bytes []byte) (err error) {
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
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "order"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = *all.Name
	if !all.Order.IsValid() {
		hasInvalidField = true
	} else {
		o.Order = *all.Order
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
