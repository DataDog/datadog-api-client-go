// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GuidedTableSort Sort configuration for a guided table.
type GuidedTableSort struct {
	// Columns to sort by, in order of application.
	Columns []WidgetFieldSort `json:"columns"`
	// Maximum number of rows to show after sorting.
	Limit int64 `json:"limit"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGuidedTableSort instantiates a new GuidedTableSort object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGuidedTableSort(columns []WidgetFieldSort, limit int64) *GuidedTableSort {
	this := GuidedTableSort{}
	this.Columns = columns
	this.Limit = limit
	return &this
}

// NewGuidedTableSortWithDefaults instantiates a new GuidedTableSort object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGuidedTableSortWithDefaults() *GuidedTableSort {
	this := GuidedTableSort{}
	return &this
}

// GetColumns returns the Columns field value.
func (o *GuidedTableSort) GetColumns() []WidgetFieldSort {
	if o == nil {
		var ret []WidgetFieldSort
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value
// and a boolean to check if the value has been set.
func (o *GuidedTableSort) GetColumnsOk() (*[]WidgetFieldSort, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Columns, true
}

// SetColumns sets field value.
func (o *GuidedTableSort) SetColumns(v []WidgetFieldSort) {
	o.Columns = v
}

// GetLimit returns the Limit field value.
func (o *GuidedTableSort) GetLimit() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *GuidedTableSort) GetLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value.
func (o *GuidedTableSort) SetLimit(v int64) {
	o.Limit = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GuidedTableSort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["columns"] = o.Columns
	toSerialize["limit"] = o.Limit

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GuidedTableSort) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Columns *[]WidgetFieldSort `json:"columns"`
		Limit   *int64             `json:"limit"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Columns == nil {
		return fmt.Errorf("required field columns missing")
	}
	if all.Limit == nil {
		return fmt.Errorf("required field limit missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"columns", "limit"})
	} else {
		return err
	}
	o.Columns = *all.Columns
	o.Limit = *all.Limit

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
