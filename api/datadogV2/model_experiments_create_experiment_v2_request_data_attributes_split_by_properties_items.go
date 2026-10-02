// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsCreateExperimentV2RequestDataAttributesSplitByPropertiesItems Complete Datadog split-by selection. Identify each property by column_name. Omit this field to copy organization defaults.
type ExperimentsCreateExperimentV2RequestDataAttributesSplitByPropertiesItems struct {
	// Exposure field that identifies the property.
	ColumnName string `json:"column_name"`
	// Data type of the column evaluated by the entry-point filter.
	ColumnType *ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationEntryPointFiltersItemsItemsColumnType `json:"column_type,omitempty"`
	// Optional display name. Defaults to column_name for a new property.
	Name *string `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsCreateExperimentV2RequestDataAttributesSplitByPropertiesItems instantiates a new ExperimentsCreateExperimentV2RequestDataAttributesSplitByPropertiesItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsCreateExperimentV2RequestDataAttributesSplitByPropertiesItems(columnName string) *ExperimentsCreateExperimentV2RequestDataAttributesSplitByPropertiesItems {
	this := ExperimentsCreateExperimentV2RequestDataAttributesSplitByPropertiesItems{}
	this.ColumnName = columnName
	return &this
}

// NewExperimentsCreateExperimentV2RequestDataAttributesSplitByPropertiesItemsWithDefaults instantiates a new ExperimentsCreateExperimentV2RequestDataAttributesSplitByPropertiesItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsCreateExperimentV2RequestDataAttributesSplitByPropertiesItemsWithDefaults() *ExperimentsCreateExperimentV2RequestDataAttributesSplitByPropertiesItems {
	this := ExperimentsCreateExperimentV2RequestDataAttributesSplitByPropertiesItems{}
	return &this
}

// GetColumnName returns the ColumnName field value.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesSplitByPropertiesItems) GetColumnName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ColumnName
}

// GetColumnNameOk returns a tuple with the ColumnName field value
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesSplitByPropertiesItems) GetColumnNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ColumnName, true
}

// SetColumnName sets field value.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesSplitByPropertiesItems) SetColumnName(v string) {
	o.ColumnName = v
}

// GetColumnType returns the ColumnType field value if set, zero value otherwise.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesSplitByPropertiesItems) GetColumnType() ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationEntryPointFiltersItemsItemsColumnType {
	if o == nil || o.ColumnType == nil {
		var ret ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationEntryPointFiltersItemsItemsColumnType
		return ret
	}
	return *o.ColumnType
}

// GetColumnTypeOk returns a tuple with the ColumnType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesSplitByPropertiesItems) GetColumnTypeOk() (*ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationEntryPointFiltersItemsItemsColumnType, bool) {
	if o == nil || o.ColumnType == nil {
		return nil, false
	}
	return o.ColumnType, true
}

// HasColumnType returns a boolean if a field has been set.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesSplitByPropertiesItems) HasColumnType() bool {
	return o != nil && o.ColumnType != nil
}

// SetColumnType gets a reference to the given ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationEntryPointFiltersItemsItemsColumnType and assigns it to the ColumnType field.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesSplitByPropertiesItems) SetColumnType(v ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationEntryPointFiltersItemsItemsColumnType) {
	o.ColumnType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesSplitByPropertiesItems) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesSplitByPropertiesItems) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesSplitByPropertiesItems) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesSplitByPropertiesItems) SetName(v string) {
	o.Name = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsCreateExperimentV2RequestDataAttributesSplitByPropertiesItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["column_name"] = o.ColumnName
	if o.ColumnType != nil {
		toSerialize["column_type"] = o.ColumnType
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesSplitByPropertiesItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ColumnName *string                                                                                                          `json:"column_name"`
		ColumnType *ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationEntryPointFiltersItemsItemsColumnType `json:"column_type,omitempty"`
		Name       *string                                                                                                          `json:"name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ColumnName == nil {
		return fmt.Errorf("required field column_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"column_name", "column_type", "name"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ColumnName = *all.ColumnName
	if all.ColumnType != nil && !all.ColumnType.IsValid() {
		hasInvalidField = true
	} else {
		o.ColumnType = all.ColumnType
	}
	o.Name = all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
