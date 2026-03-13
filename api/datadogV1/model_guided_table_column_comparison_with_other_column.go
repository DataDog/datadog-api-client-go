// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GuidedTableColumnComparisonWithOtherColumn Comparison of a column value against another column.
type GuidedTableColumnComparisonWithOtherColumn struct {
	// Name of the column to compare against.
	Column string `json:"column"`
	//
	Type GuidedTableColumnComparisonWithOtherColumnType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGuidedTableColumnComparisonWithOtherColumn instantiates a new GuidedTableColumnComparisonWithOtherColumn object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGuidedTableColumnComparisonWithOtherColumn(column string, typeVar GuidedTableColumnComparisonWithOtherColumnType) *GuidedTableColumnComparisonWithOtherColumn {
	this := GuidedTableColumnComparisonWithOtherColumn{}
	this.Column = column
	this.Type = typeVar
	return &this
}

// NewGuidedTableColumnComparisonWithOtherColumnWithDefaults instantiates a new GuidedTableColumnComparisonWithOtherColumn object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGuidedTableColumnComparisonWithOtherColumnWithDefaults() *GuidedTableColumnComparisonWithOtherColumn {
	this := GuidedTableColumnComparisonWithOtherColumn{}
	return &this
}

// GetColumn returns the Column field value.
func (o *GuidedTableColumnComparisonWithOtherColumn) GetColumn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Column
}

// GetColumnOk returns a tuple with the Column field value
// and a boolean to check if the value has been set.
func (o *GuidedTableColumnComparisonWithOtherColumn) GetColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Column, true
}

// SetColumn sets field value.
func (o *GuidedTableColumnComparisonWithOtherColumn) SetColumn(v string) {
	o.Column = v
}

// GetType returns the Type field value.
func (o *GuidedTableColumnComparisonWithOtherColumn) GetType() GuidedTableColumnComparisonWithOtherColumnType {
	if o == nil {
		var ret GuidedTableColumnComparisonWithOtherColumnType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GuidedTableColumnComparisonWithOtherColumn) GetTypeOk() (*GuidedTableColumnComparisonWithOtherColumnType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *GuidedTableColumnComparisonWithOtherColumn) SetType(v GuidedTableColumnComparisonWithOtherColumnType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GuidedTableColumnComparisonWithOtherColumn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["column"] = o.Column
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GuidedTableColumnComparisonWithOtherColumn) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Column *string                                         `json:"column"`
		Type   *GuidedTableColumnComparisonWithOtherColumnType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Column == nil {
		return fmt.Errorf("required field column missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"column", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Column = *all.Column
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
