// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringDatasetDefinitionColumn A column in a dataset definition.
type SecurityMonitoringDatasetDefinitionColumn struct {
	// The name of the column.
	Column string `json:"column"`
	// The type of the column in a dataset definition.
	Type SecurityMonitoringDatasetDefinitionColumnType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringDatasetDefinitionColumn instantiates a new SecurityMonitoringDatasetDefinitionColumn object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringDatasetDefinitionColumn(column string, typeVar SecurityMonitoringDatasetDefinitionColumnType) *SecurityMonitoringDatasetDefinitionColumn {
	this := SecurityMonitoringDatasetDefinitionColumn{}
	this.Column = column
	this.Type = typeVar
	return &this
}

// NewSecurityMonitoringDatasetDefinitionColumnWithDefaults instantiates a new SecurityMonitoringDatasetDefinitionColumn object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringDatasetDefinitionColumnWithDefaults() *SecurityMonitoringDatasetDefinitionColumn {
	this := SecurityMonitoringDatasetDefinitionColumn{}
	return &this
}

// GetColumn returns the Column field value.
func (o *SecurityMonitoringDatasetDefinitionColumn) GetColumn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Column
}

// GetColumnOk returns a tuple with the Column field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetDefinitionColumn) GetColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Column, true
}

// SetColumn sets field value.
func (o *SecurityMonitoringDatasetDefinitionColumn) SetColumn(v string) {
	o.Column = v
}

// GetType returns the Type field value.
func (o *SecurityMonitoringDatasetDefinitionColumn) GetType() SecurityMonitoringDatasetDefinitionColumnType {
	if o == nil {
		var ret SecurityMonitoringDatasetDefinitionColumnType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetDefinitionColumn) GetTypeOk() (*SecurityMonitoringDatasetDefinitionColumnType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SecurityMonitoringDatasetDefinitionColumn) SetType(v SecurityMonitoringDatasetDefinitionColumnType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringDatasetDefinitionColumn) MarshalJSON() ([]byte, error) {
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
func (o *SecurityMonitoringDatasetDefinitionColumn) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Column *string                                        `json:"column"`
		Type   *SecurityMonitoringDatasetDefinitionColumnType `json:"type"`
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
