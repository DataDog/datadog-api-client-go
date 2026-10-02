// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsExperimentV2DTODataAttributesSplitByPropertiesItems Property used to split experiment results into analysis dimensions.
type ExperimentsExperimentV2DTODataAttributesSplitByPropertiesItems struct {
	// Exposure field or Warehouse column used for the analysis dimension.
	ColumnName string `json:"column_name"`
	// Type of the Datadog exposure field or Warehouse column.
	ColumnType ExperimentsPatchExperimentV2ResponseDataAttributesSplitByPropertiesItemsColumnType `json:"column_type"`
	// Read-only property ID. Omit it from POST and PATCH; writes identify properties by column_name.
	Id string `json:"id"`
	// Display name for the analysis dimension.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsExperimentV2DTODataAttributesSplitByPropertiesItems instantiates a new ExperimentsExperimentV2DTODataAttributesSplitByPropertiesItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsExperimentV2DTODataAttributesSplitByPropertiesItems(columnName string, columnType ExperimentsPatchExperimentV2ResponseDataAttributesSplitByPropertiesItemsColumnType, id string, name string) *ExperimentsExperimentV2DTODataAttributesSplitByPropertiesItems {
	this := ExperimentsExperimentV2DTODataAttributesSplitByPropertiesItems{}
	this.ColumnName = columnName
	this.ColumnType = columnType
	this.Id = id
	this.Name = name
	return &this
}

// NewExperimentsExperimentV2DTODataAttributesSplitByPropertiesItemsWithDefaults instantiates a new ExperimentsExperimentV2DTODataAttributesSplitByPropertiesItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsExperimentV2DTODataAttributesSplitByPropertiesItemsWithDefaults() *ExperimentsExperimentV2DTODataAttributesSplitByPropertiesItems {
	this := ExperimentsExperimentV2DTODataAttributesSplitByPropertiesItems{}
	return &this
}

// GetColumnName returns the ColumnName field value.
func (o *ExperimentsExperimentV2DTODataAttributesSplitByPropertiesItems) GetColumnName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ColumnName
}

// GetColumnNameOk returns a tuple with the ColumnName field value
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentV2DTODataAttributesSplitByPropertiesItems) GetColumnNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ColumnName, true
}

// SetColumnName sets field value.
func (o *ExperimentsExperimentV2DTODataAttributesSplitByPropertiesItems) SetColumnName(v string) {
	o.ColumnName = v
}

// GetColumnType returns the ColumnType field value.
func (o *ExperimentsExperimentV2DTODataAttributesSplitByPropertiesItems) GetColumnType() ExperimentsPatchExperimentV2ResponseDataAttributesSplitByPropertiesItemsColumnType {
	if o == nil {
		var ret ExperimentsPatchExperimentV2ResponseDataAttributesSplitByPropertiesItemsColumnType
		return ret
	}
	return o.ColumnType
}

// GetColumnTypeOk returns a tuple with the ColumnType field value
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentV2DTODataAttributesSplitByPropertiesItems) GetColumnTypeOk() (*ExperimentsPatchExperimentV2ResponseDataAttributesSplitByPropertiesItemsColumnType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ColumnType, true
}

// SetColumnType sets field value.
func (o *ExperimentsExperimentV2DTODataAttributesSplitByPropertiesItems) SetColumnType(v ExperimentsPatchExperimentV2ResponseDataAttributesSplitByPropertiesItemsColumnType) {
	o.ColumnType = v
}

// GetId returns the Id field value.
func (o *ExperimentsExperimentV2DTODataAttributesSplitByPropertiesItems) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentV2DTODataAttributesSplitByPropertiesItems) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ExperimentsExperimentV2DTODataAttributesSplitByPropertiesItems) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value.
func (o *ExperimentsExperimentV2DTODataAttributesSplitByPropertiesItems) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentV2DTODataAttributesSplitByPropertiesItems) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ExperimentsExperimentV2DTODataAttributesSplitByPropertiesItems) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsExperimentV2DTODataAttributesSplitByPropertiesItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["column_name"] = o.ColumnName
	toSerialize["column_type"] = o.ColumnType
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsExperimentV2DTODataAttributesSplitByPropertiesItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ColumnName *string                                                                             `json:"column_name"`
		ColumnType *ExperimentsPatchExperimentV2ResponseDataAttributesSplitByPropertiesItemsColumnType `json:"column_type"`
		Id         *string                                                                             `json:"id"`
		Name       *string                                                                             `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ColumnName == nil {
		return fmt.Errorf("required field column_name missing")
	}
	if all.ColumnType == nil {
		return fmt.Errorf("required field column_type missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"column_name", "column_type", "id", "name"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ColumnName = *all.ColumnName
	if !all.ColumnType.IsValid() {
		hasInvalidField = true
	} else {
		o.ColumnType = *all.ColumnType
	}
	o.Id = *all.Id
	o.Name = *all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
