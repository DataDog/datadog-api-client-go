// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsDatadogEntryPointFilter Complete Datadog OR-of-ANDs entry-point filter expression.
type ExperimentsDatadogEntryPointFilter struct {
	// Exposure field evaluated by the entry-point filter.
	Column string `json:"column"`
	// Data type of the column evaluated by the entry-point filter.
	ColumnType ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationEntryPointFiltersItemsItemsColumnType `json:"column_type"`
	// Comparison applied by the Datadog entry-point filter.
	Operation ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationEntryPointFiltersItemsItemsOperation `json:"operation"`
	// Comparison values used by the filter operation.
	Values []string `json:"values"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsDatadogEntryPointFilter instantiates a new ExperimentsDatadogEntryPointFilter object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsDatadogEntryPointFilter(column string, columnType ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationEntryPointFiltersItemsItemsColumnType, operation ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationEntryPointFiltersItemsItemsOperation, values []string) *ExperimentsDatadogEntryPointFilter {
	this := ExperimentsDatadogEntryPointFilter{}
	this.Column = column
	this.ColumnType = columnType
	this.Operation = operation
	this.Values = values
	return &this
}

// NewExperimentsDatadogEntryPointFilterWithDefaults instantiates a new ExperimentsDatadogEntryPointFilter object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsDatadogEntryPointFilterWithDefaults() *ExperimentsDatadogEntryPointFilter {
	this := ExperimentsDatadogEntryPointFilter{}
	return &this
}

// GetColumn returns the Column field value.
func (o *ExperimentsDatadogEntryPointFilter) GetColumn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Column
}

// GetColumnOk returns a tuple with the Column field value
// and a boolean to check if the value has been set.
func (o *ExperimentsDatadogEntryPointFilter) GetColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Column, true
}

// SetColumn sets field value.
func (o *ExperimentsDatadogEntryPointFilter) SetColumn(v string) {
	o.Column = v
}

// GetColumnType returns the ColumnType field value.
func (o *ExperimentsDatadogEntryPointFilter) GetColumnType() ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationEntryPointFiltersItemsItemsColumnType {
	if o == nil {
		var ret ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationEntryPointFiltersItemsItemsColumnType
		return ret
	}
	return o.ColumnType
}

// GetColumnTypeOk returns a tuple with the ColumnType field value
// and a boolean to check if the value has been set.
func (o *ExperimentsDatadogEntryPointFilter) GetColumnTypeOk() (*ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationEntryPointFiltersItemsItemsColumnType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ColumnType, true
}

// SetColumnType sets field value.
func (o *ExperimentsDatadogEntryPointFilter) SetColumnType(v ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationEntryPointFiltersItemsItemsColumnType) {
	o.ColumnType = v
}

// GetOperation returns the Operation field value.
func (o *ExperimentsDatadogEntryPointFilter) GetOperation() ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationEntryPointFiltersItemsItemsOperation {
	if o == nil {
		var ret ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationEntryPointFiltersItemsItemsOperation
		return ret
	}
	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *ExperimentsDatadogEntryPointFilter) GetOperationOk() (*ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationEntryPointFiltersItemsItemsOperation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value.
func (o *ExperimentsDatadogEntryPointFilter) SetOperation(v ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationEntryPointFiltersItemsItemsOperation) {
	o.Operation = v
}

// GetValues returns the Values field value.
func (o *ExperimentsDatadogEntryPointFilter) GetValues() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *ExperimentsDatadogEntryPointFilter) GetValuesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Values, true
}

// SetValues sets field value.
func (o *ExperimentsDatadogEntryPointFilter) SetValues(v []string) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsDatadogEntryPointFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["column"] = o.Column
	toSerialize["column_type"] = o.ColumnType
	toSerialize["operation"] = o.Operation
	toSerialize["values"] = o.Values

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsDatadogEntryPointFilter) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Column     *string                                                                                                          `json:"column"`
		ColumnType *ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationEntryPointFiltersItemsItemsColumnType `json:"column_type"`
		Operation  *ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationEntryPointFiltersItemsItemsOperation  `json:"operation"`
		Values     *[]string                                                                                                        `json:"values"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Column == nil {
		return fmt.Errorf("required field column missing")
	}
	if all.ColumnType == nil {
		return fmt.Errorf("required field column_type missing")
	}
	if all.Operation == nil {
		return fmt.Errorf("required field operation missing")
	}
	if all.Values == nil {
		return fmt.Errorf("required field values missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"column", "column_type", "operation", "values"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Column = *all.Column
	if !all.ColumnType.IsValid() {
		hasInvalidField = true
	} else {
		o.ColumnType = *all.ColumnType
	}
	if !all.Operation.IsValid() {
		hasInvalidField = true
	} else {
		o.Operation = *all.Operation
	}
	o.Values = *all.Values

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
