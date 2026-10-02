// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsMetricSQLModelPropertyInput A property column defined by a metric SQL model.
type ExperimentsMetricSQLModelPropertyInput struct {
	// Name of the SQL result column that supplies this property.
	ColumnName string `json:"column_name"`
	// Data type of a column in the SQL model.
	ColumnType ExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnType `json:"column_type"`
	// Optional text that explains what this property represents.
	Description datadog.NullableString `json:"description,omitempty"`
	// Opaque metadata preserved when this property is migrated.
	MigrationMetadata interface{} `json:"migration_metadata,omitempty"`
	// Name used to identify the property in the model.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsMetricSQLModelPropertyInput instantiates a new ExperimentsMetricSQLModelPropertyInput object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsMetricSQLModelPropertyInput(columnName string, columnType ExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnType, name string) *ExperimentsMetricSQLModelPropertyInput {
	this := ExperimentsMetricSQLModelPropertyInput{}
	this.ColumnName = columnName
	this.ColumnType = columnType
	this.Name = name
	return &this
}

// NewExperimentsMetricSQLModelPropertyInputWithDefaults instantiates a new ExperimentsMetricSQLModelPropertyInput object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsMetricSQLModelPropertyInputWithDefaults() *ExperimentsMetricSQLModelPropertyInput {
	this := ExperimentsMetricSQLModelPropertyInput{}
	return &this
}

// GetColumnName returns the ColumnName field value.
func (o *ExperimentsMetricSQLModelPropertyInput) GetColumnName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ColumnName
}

// GetColumnNameOk returns a tuple with the ColumnName field value
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricSQLModelPropertyInput) GetColumnNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ColumnName, true
}

// SetColumnName sets field value.
func (o *ExperimentsMetricSQLModelPropertyInput) SetColumnName(v string) {
	o.ColumnName = v
}

// GetColumnType returns the ColumnType field value.
func (o *ExperimentsMetricSQLModelPropertyInput) GetColumnType() ExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnType {
	if o == nil {
		var ret ExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnType
		return ret
	}
	return o.ColumnType
}

// GetColumnTypeOk returns a tuple with the ColumnType field value
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricSQLModelPropertyInput) GetColumnTypeOk() (*ExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ColumnType, true
}

// SetColumnType sets field value.
func (o *ExperimentsMetricSQLModelPropertyInput) SetColumnType(v ExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnType) {
	o.ColumnType = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsMetricSQLModelPropertyInput) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsMetricSQLModelPropertyInput) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ExperimentsMetricSQLModelPropertyInput) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given datadog.NullableString and assigns it to the Description field.
func (o *ExperimentsMetricSQLModelPropertyInput) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *ExperimentsMetricSQLModelPropertyInput) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *ExperimentsMetricSQLModelPropertyInput) UnsetDescription() {
	o.Description.Unset()
}

// GetMigrationMetadata returns the MigrationMetadata field value if set, zero value otherwise.
func (o *ExperimentsMetricSQLModelPropertyInput) GetMigrationMetadata() interface{} {
	if o == nil || o.MigrationMetadata == nil {
		var ret interface{}
		return ret
	}
	return o.MigrationMetadata
}

// GetMigrationMetadataOk returns a tuple with the MigrationMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricSQLModelPropertyInput) GetMigrationMetadataOk() (*interface{}, bool) {
	if o == nil || o.MigrationMetadata == nil {
		return nil, false
	}
	return &o.MigrationMetadata, true
}

// HasMigrationMetadata returns a boolean if a field has been set.
func (o *ExperimentsMetricSQLModelPropertyInput) HasMigrationMetadata() bool {
	return o != nil && o.MigrationMetadata != nil
}

// SetMigrationMetadata gets a reference to the given interface{} and assigns it to the MigrationMetadata field.
func (o *ExperimentsMetricSQLModelPropertyInput) SetMigrationMetadata(v interface{}) {
	o.MigrationMetadata = v
}

// GetName returns the Name field value.
func (o *ExperimentsMetricSQLModelPropertyInput) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricSQLModelPropertyInput) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ExperimentsMetricSQLModelPropertyInput) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsMetricSQLModelPropertyInput) MarshalJSON() ([]byte, error) {
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
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsMetricSQLModelPropertyInput) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ColumnName        *string                                                                  `json:"column_name"`
		ColumnType        *ExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnType `json:"column_type"`
		Description       datadog.NullableString                                                   `json:"description,omitempty"`
		MigrationMetadata interface{}                                                              `json:"migration_metadata,omitempty"`
		Name              *string                                                                  `json:"name"`
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
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
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
	o.Name = *all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
