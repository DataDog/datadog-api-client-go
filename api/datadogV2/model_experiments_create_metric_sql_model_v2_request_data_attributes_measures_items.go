// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsCreateMetricSQLModelV2RequestDataAttributesMeasuresItems Column in the SQL model that supplies values for a metric measure.
type ExperimentsCreateMetricSQLModelV2RequestDataAttributesMeasuresItems struct {
	// SQL result column that contains the measure values.
	ColumnName string `json:"column_name"`
	// Data type of a column in the SQL model.
	ColumnType ExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnType `json:"column_type"`
	// Description of the measure.
	Description datadog.NullableString `json:"description,omitempty"`
	// Metadata associated with migration of this resource.
	MigrationMetadata interface{} `json:"migration_metadata,omitempty"`
	// Display name of the measure.
	Name datadog.NullableString `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsCreateMetricSQLModelV2RequestDataAttributesMeasuresItems instantiates a new ExperimentsCreateMetricSQLModelV2RequestDataAttributesMeasuresItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsCreateMetricSQLModelV2RequestDataAttributesMeasuresItems(columnName string, columnType ExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnType) *ExperimentsCreateMetricSQLModelV2RequestDataAttributesMeasuresItems {
	this := ExperimentsCreateMetricSQLModelV2RequestDataAttributesMeasuresItems{}
	this.ColumnName = columnName
	this.ColumnType = columnType
	return &this
}

// NewExperimentsCreateMetricSQLModelV2RequestDataAttributesMeasuresItemsWithDefaults instantiates a new ExperimentsCreateMetricSQLModelV2RequestDataAttributesMeasuresItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsCreateMetricSQLModelV2RequestDataAttributesMeasuresItemsWithDefaults() *ExperimentsCreateMetricSQLModelV2RequestDataAttributesMeasuresItems {
	this := ExperimentsCreateMetricSQLModelV2RequestDataAttributesMeasuresItems{}
	return &this
}

// GetColumnName returns the ColumnName field value.
func (o *ExperimentsCreateMetricSQLModelV2RequestDataAttributesMeasuresItems) GetColumnName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ColumnName
}

// GetColumnNameOk returns a tuple with the ColumnName field value
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateMetricSQLModelV2RequestDataAttributesMeasuresItems) GetColumnNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ColumnName, true
}

// SetColumnName sets field value.
func (o *ExperimentsCreateMetricSQLModelV2RequestDataAttributesMeasuresItems) SetColumnName(v string) {
	o.ColumnName = v
}

// GetColumnType returns the ColumnType field value.
func (o *ExperimentsCreateMetricSQLModelV2RequestDataAttributesMeasuresItems) GetColumnType() ExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnType {
	if o == nil {
		var ret ExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnType
		return ret
	}
	return o.ColumnType
}

// GetColumnTypeOk returns a tuple with the ColumnType field value
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateMetricSQLModelV2RequestDataAttributesMeasuresItems) GetColumnTypeOk() (*ExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ColumnType, true
}

// SetColumnType sets field value.
func (o *ExperimentsCreateMetricSQLModelV2RequestDataAttributesMeasuresItems) SetColumnType(v ExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnType) {
	o.ColumnType = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsCreateMetricSQLModelV2RequestDataAttributesMeasuresItems) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsCreateMetricSQLModelV2RequestDataAttributesMeasuresItems) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ExperimentsCreateMetricSQLModelV2RequestDataAttributesMeasuresItems) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given datadog.NullableString and assigns it to the Description field.
func (o *ExperimentsCreateMetricSQLModelV2RequestDataAttributesMeasuresItems) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *ExperimentsCreateMetricSQLModelV2RequestDataAttributesMeasuresItems) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *ExperimentsCreateMetricSQLModelV2RequestDataAttributesMeasuresItems) UnsetDescription() {
	o.Description.Unset()
}

// GetMigrationMetadata returns the MigrationMetadata field value if set, zero value otherwise.
func (o *ExperimentsCreateMetricSQLModelV2RequestDataAttributesMeasuresItems) GetMigrationMetadata() interface{} {
	if o == nil || o.MigrationMetadata == nil {
		var ret interface{}
		return ret
	}
	return o.MigrationMetadata
}

// GetMigrationMetadataOk returns a tuple with the MigrationMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateMetricSQLModelV2RequestDataAttributesMeasuresItems) GetMigrationMetadataOk() (*interface{}, bool) {
	if o == nil || o.MigrationMetadata == nil {
		return nil, false
	}
	return &o.MigrationMetadata, true
}

// HasMigrationMetadata returns a boolean if a field has been set.
func (o *ExperimentsCreateMetricSQLModelV2RequestDataAttributesMeasuresItems) HasMigrationMetadata() bool {
	return o != nil && o.MigrationMetadata != nil
}

// SetMigrationMetadata gets a reference to the given interface{} and assigns it to the MigrationMetadata field.
func (o *ExperimentsCreateMetricSQLModelV2RequestDataAttributesMeasuresItems) SetMigrationMetadata(v interface{}) {
	o.MigrationMetadata = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsCreateMetricSQLModelV2RequestDataAttributesMeasuresItems) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsCreateMetricSQLModelV2RequestDataAttributesMeasuresItems) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ExperimentsCreateMetricSQLModelV2RequestDataAttributesMeasuresItems) HasName() bool {
	return o != nil && o.Name.IsSet()
}

// SetName gets a reference to the given datadog.NullableString and assigns it to the Name field.
func (o *ExperimentsCreateMetricSQLModelV2RequestDataAttributesMeasuresItems) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil.
func (o *ExperimentsCreateMetricSQLModelV2RequestDataAttributesMeasuresItems) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil.
func (o *ExperimentsCreateMetricSQLModelV2RequestDataAttributesMeasuresItems) UnsetName() {
	o.Name.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsCreateMetricSQLModelV2RequestDataAttributesMeasuresItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["column_name"] = o.ColumnName
	toSerialize["column_type"] = o.ColumnType
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.MigrationMetadata != nil {
		toSerialize["migration_metadata"] = o.MigrationMetadata
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsCreateMetricSQLModelV2RequestDataAttributesMeasuresItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ColumnName        *string                                                                  `json:"column_name"`
		ColumnType        *ExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnType `json:"column_type"`
		Description       datadog.NullableString                                                   `json:"description,omitempty"`
		MigrationMetadata interface{}                                                              `json:"migration_metadata,omitempty"`
		Name              datadog.NullableString                                                   `json:"name,omitempty"`
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
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"column_name", "column_type", "description", "migration_metadata", "name"})
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
	o.Description = all.Description
	o.MigrationMetadata = all.MigrationMetadata
	o.Name = all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
