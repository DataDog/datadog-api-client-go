// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsUpdateMetricSQLModelV2RequestDataAttributes Complete column mappings and query used to replace the metric SQL model.
type ExperimentsUpdateMetricSQLModelV2RequestDataAttributes struct {
	// SQL column used to partition the source data by date.
	DatePartitionColumn datadog.NullableString `json:"date_partition_column,omitempty"`
	// Text that explains the metric SQL model.
	Description datadog.NullableString `json:"description,omitempty"`
	// Measures available from the SQL model's result columns.
	Measures []ExperimentsCreateMetricSQLModelV2RequestDataAttributesMeasuresItems `json:"measures,omitempty"`
	// Metadata retained for resources imported from another system.
	MigrationMetadata interface{} `json:"migration_metadata,omitempty"`
	// Display name of the metric SQL model.
	Name string `json:"name"`
	// Property columns exposed by the SQL model.
	Properties []ExperimentsMetricSQLModelPropertyInput `json:"properties,omitempty"`
	// SQL query that produces the model's source data.
	Sql string `json:"sql"`
	// Subject types mapped to columns in the SQL model.
	SubjectTypes []ExperimentsCreateExposureSQLModelV2RequestDataAttributesSubjectTypesItems `json:"subject_types"`
	// SQL column that supplies the event timestamp.
	TimestampColumn string `json:"timestamp_column"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsUpdateMetricSQLModelV2RequestDataAttributes instantiates a new ExperimentsUpdateMetricSQLModelV2RequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsUpdateMetricSQLModelV2RequestDataAttributes(name string, sql string, subjectTypes []ExperimentsCreateExposureSQLModelV2RequestDataAttributesSubjectTypesItems, timestampColumn string) *ExperimentsUpdateMetricSQLModelV2RequestDataAttributes {
	this := ExperimentsUpdateMetricSQLModelV2RequestDataAttributes{}
	this.Name = name
	this.Sql = sql
	this.SubjectTypes = subjectTypes
	this.TimestampColumn = timestampColumn
	return &this
}

// NewExperimentsUpdateMetricSQLModelV2RequestDataAttributesWithDefaults instantiates a new ExperimentsUpdateMetricSQLModelV2RequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsUpdateMetricSQLModelV2RequestDataAttributesWithDefaults() *ExperimentsUpdateMetricSQLModelV2RequestDataAttributes {
	this := ExperimentsUpdateMetricSQLModelV2RequestDataAttributes{}
	return &this
}

// GetDatePartitionColumn returns the DatePartitionColumn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsUpdateMetricSQLModelV2RequestDataAttributes) GetDatePartitionColumn() string {
	if o == nil || o.DatePartitionColumn.Get() == nil {
		var ret string
		return ret
	}
	return *o.DatePartitionColumn.Get()
}

// GetDatePartitionColumnOk returns a tuple with the DatePartitionColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsUpdateMetricSQLModelV2RequestDataAttributes) GetDatePartitionColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DatePartitionColumn.Get(), o.DatePartitionColumn.IsSet()
}

// HasDatePartitionColumn returns a boolean if a field has been set.
func (o *ExperimentsUpdateMetricSQLModelV2RequestDataAttributes) HasDatePartitionColumn() bool {
	return o != nil && o.DatePartitionColumn.IsSet()
}

// SetDatePartitionColumn gets a reference to the given datadog.NullableString and assigns it to the DatePartitionColumn field.
func (o *ExperimentsUpdateMetricSQLModelV2RequestDataAttributes) SetDatePartitionColumn(v string) {
	o.DatePartitionColumn.Set(&v)
}

// SetDatePartitionColumnNil sets the value for DatePartitionColumn to be an explicit nil.
func (o *ExperimentsUpdateMetricSQLModelV2RequestDataAttributes) SetDatePartitionColumnNil() {
	o.DatePartitionColumn.Set(nil)
}

// UnsetDatePartitionColumn ensures that no value is present for DatePartitionColumn, not even an explicit nil.
func (o *ExperimentsUpdateMetricSQLModelV2RequestDataAttributes) UnsetDatePartitionColumn() {
	o.DatePartitionColumn.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsUpdateMetricSQLModelV2RequestDataAttributes) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsUpdateMetricSQLModelV2RequestDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ExperimentsUpdateMetricSQLModelV2RequestDataAttributes) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given datadog.NullableString and assigns it to the Description field.
func (o *ExperimentsUpdateMetricSQLModelV2RequestDataAttributes) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *ExperimentsUpdateMetricSQLModelV2RequestDataAttributes) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *ExperimentsUpdateMetricSQLModelV2RequestDataAttributes) UnsetDescription() {
	o.Description.Unset()
}

// GetMeasures returns the Measures field value if set, zero value otherwise.
func (o *ExperimentsUpdateMetricSQLModelV2RequestDataAttributes) GetMeasures() []ExperimentsCreateMetricSQLModelV2RequestDataAttributesMeasuresItems {
	if o == nil || o.Measures == nil {
		var ret []ExperimentsCreateMetricSQLModelV2RequestDataAttributesMeasuresItems
		return ret
	}
	return o.Measures
}

// GetMeasuresOk returns a tuple with the Measures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsUpdateMetricSQLModelV2RequestDataAttributes) GetMeasuresOk() (*[]ExperimentsCreateMetricSQLModelV2RequestDataAttributesMeasuresItems, bool) {
	if o == nil || o.Measures == nil {
		return nil, false
	}
	return &o.Measures, true
}

// HasMeasures returns a boolean if a field has been set.
func (o *ExperimentsUpdateMetricSQLModelV2RequestDataAttributes) HasMeasures() bool {
	return o != nil && o.Measures != nil
}

// SetMeasures gets a reference to the given []ExperimentsCreateMetricSQLModelV2RequestDataAttributesMeasuresItems and assigns it to the Measures field.
func (o *ExperimentsUpdateMetricSQLModelV2RequestDataAttributes) SetMeasures(v []ExperimentsCreateMetricSQLModelV2RequestDataAttributesMeasuresItems) {
	o.Measures = v
}

// GetMigrationMetadata returns the MigrationMetadata field value if set, zero value otherwise.
func (o *ExperimentsUpdateMetricSQLModelV2RequestDataAttributes) GetMigrationMetadata() interface{} {
	if o == nil || o.MigrationMetadata == nil {
		var ret interface{}
		return ret
	}
	return o.MigrationMetadata
}

// GetMigrationMetadataOk returns a tuple with the MigrationMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsUpdateMetricSQLModelV2RequestDataAttributes) GetMigrationMetadataOk() (*interface{}, bool) {
	if o == nil || o.MigrationMetadata == nil {
		return nil, false
	}
	return &o.MigrationMetadata, true
}

// HasMigrationMetadata returns a boolean if a field has been set.
func (o *ExperimentsUpdateMetricSQLModelV2RequestDataAttributes) HasMigrationMetadata() bool {
	return o != nil && o.MigrationMetadata != nil
}

// SetMigrationMetadata gets a reference to the given interface{} and assigns it to the MigrationMetadata field.
func (o *ExperimentsUpdateMetricSQLModelV2RequestDataAttributes) SetMigrationMetadata(v interface{}) {
	o.MigrationMetadata = v
}

// GetName returns the Name field value.
func (o *ExperimentsUpdateMetricSQLModelV2RequestDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExperimentsUpdateMetricSQLModelV2RequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ExperimentsUpdateMetricSQLModelV2RequestDataAttributes) SetName(v string) {
	o.Name = v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *ExperimentsUpdateMetricSQLModelV2RequestDataAttributes) GetProperties() []ExperimentsMetricSQLModelPropertyInput {
	if o == nil || o.Properties == nil {
		var ret []ExperimentsMetricSQLModelPropertyInput
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsUpdateMetricSQLModelV2RequestDataAttributes) GetPropertiesOk() (*[]ExperimentsMetricSQLModelPropertyInput, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return &o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *ExperimentsUpdateMetricSQLModelV2RequestDataAttributes) HasProperties() bool {
	return o != nil && o.Properties != nil
}

// SetProperties gets a reference to the given []ExperimentsMetricSQLModelPropertyInput and assigns it to the Properties field.
func (o *ExperimentsUpdateMetricSQLModelV2RequestDataAttributes) SetProperties(v []ExperimentsMetricSQLModelPropertyInput) {
	o.Properties = v
}

// GetSql returns the Sql field value.
func (o *ExperimentsUpdateMetricSQLModelV2RequestDataAttributes) GetSql() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Sql
}

// GetSqlOk returns a tuple with the Sql field value
// and a boolean to check if the value has been set.
func (o *ExperimentsUpdateMetricSQLModelV2RequestDataAttributes) GetSqlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sql, true
}

// SetSql sets field value.
func (o *ExperimentsUpdateMetricSQLModelV2RequestDataAttributes) SetSql(v string) {
	o.Sql = v
}

// GetSubjectTypes returns the SubjectTypes field value.
func (o *ExperimentsUpdateMetricSQLModelV2RequestDataAttributes) GetSubjectTypes() []ExperimentsCreateExposureSQLModelV2RequestDataAttributesSubjectTypesItems {
	if o == nil {
		var ret []ExperimentsCreateExposureSQLModelV2RequestDataAttributesSubjectTypesItems
		return ret
	}
	return o.SubjectTypes
}

// GetSubjectTypesOk returns a tuple with the SubjectTypes field value
// and a boolean to check if the value has been set.
func (o *ExperimentsUpdateMetricSQLModelV2RequestDataAttributes) GetSubjectTypesOk() (*[]ExperimentsCreateExposureSQLModelV2RequestDataAttributesSubjectTypesItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubjectTypes, true
}

// SetSubjectTypes sets field value.
func (o *ExperimentsUpdateMetricSQLModelV2RequestDataAttributes) SetSubjectTypes(v []ExperimentsCreateExposureSQLModelV2RequestDataAttributesSubjectTypesItems) {
	o.SubjectTypes = v
}

// GetTimestampColumn returns the TimestampColumn field value.
func (o *ExperimentsUpdateMetricSQLModelV2RequestDataAttributes) GetTimestampColumn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TimestampColumn
}

// GetTimestampColumnOk returns a tuple with the TimestampColumn field value
// and a boolean to check if the value has been set.
func (o *ExperimentsUpdateMetricSQLModelV2RequestDataAttributes) GetTimestampColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimestampColumn, true
}

// SetTimestampColumn sets field value.
func (o *ExperimentsUpdateMetricSQLModelV2RequestDataAttributes) SetTimestampColumn(v string) {
	o.TimestampColumn = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsUpdateMetricSQLModelV2RequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DatePartitionColumn.IsSet() {
		toSerialize["date_partition_column"] = o.DatePartitionColumn.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Measures != nil {
		toSerialize["measures"] = o.Measures
	}
	if o.MigrationMetadata != nil {
		toSerialize["migration_metadata"] = o.MigrationMetadata
	}
	toSerialize["name"] = o.Name
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	toSerialize["sql"] = o.Sql
	toSerialize["subject_types"] = o.SubjectTypes
	toSerialize["timestamp_column"] = o.TimestampColumn

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsUpdateMetricSQLModelV2RequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DatePartitionColumn datadog.NullableString                                                       `json:"date_partition_column,omitempty"`
		Description         datadog.NullableString                                                       `json:"description,omitempty"`
		Measures            []ExperimentsCreateMetricSQLModelV2RequestDataAttributesMeasuresItems        `json:"measures,omitempty"`
		MigrationMetadata   interface{}                                                                  `json:"migration_metadata,omitempty"`
		Name                *string                                                                      `json:"name"`
		Properties          []ExperimentsMetricSQLModelPropertyInput                                     `json:"properties,omitempty"`
		Sql                 *string                                                                      `json:"sql"`
		SubjectTypes        *[]ExperimentsCreateExposureSQLModelV2RequestDataAttributesSubjectTypesItems `json:"subject_types"`
		TimestampColumn     *string                                                                      `json:"timestamp_column"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Sql == nil {
		return fmt.Errorf("required field sql missing")
	}
	if all.SubjectTypes == nil {
		return fmt.Errorf("required field subject_types missing")
	}
	if all.TimestampColumn == nil {
		return fmt.Errorf("required field timestamp_column missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"date_partition_column", "description", "measures", "migration_metadata", "name", "properties", "sql", "subject_types", "timestamp_column"})
	} else {
		return err
	}
	o.DatePartitionColumn = all.DatePartitionColumn
	o.Description = all.Description
	o.Measures = all.Measures
	o.MigrationMetadata = all.MigrationMetadata
	o.Name = *all.Name
	o.Properties = all.Properties
	o.Sql = *all.Sql
	o.SubjectTypes = *all.SubjectTypes
	o.TimestampColumn = *all.TimestampColumn

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
