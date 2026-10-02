// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsUpdateExposureSQLModelV2RequestDataAttributes Complete column mappings and query used to replace the exposure SQL model.
type ExperimentsUpdateExposureSQLModelV2RequestDataAttributes struct {
	// SQL column used to partition the source data by date.
	DatePartitionColumn datadog.NullableString `json:"date_partition_column,omitempty"`
	// SQL column that identifies the experiment for each exposure.
	ExperimentColumn string `json:"experiment_column"`
	// Metadata retained for resources imported from another system.
	MigrationMetadata interface{} `json:"migration_metadata,omitempty"`
	// Display name of the exposure SQL model.
	Name string `json:"name"`
	// Property columns exposed by the SQL model.
	Properties []ExperimentsSQLModelPropertyInput `json:"properties,omitempty"`
	// SQL query that produces the model's source data.
	Sql string `json:"sql"`
	// Subject types mapped to columns in the SQL model.
	SubjectTypes []ExperimentsCreateExposureSQLModelV2RequestDataAttributesSubjectTypesItems `json:"subject_types"`
	// SQL column that supplies the event timestamp.
	TimestampColumn string `json:"timestamp_column"`
	// SQL column that identifies the variant for each exposure.
	VariantColumn string `json:"variant_column"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsUpdateExposureSQLModelV2RequestDataAttributes instantiates a new ExperimentsUpdateExposureSQLModelV2RequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsUpdateExposureSQLModelV2RequestDataAttributes(experimentColumn string, name string, sql string, subjectTypes []ExperimentsCreateExposureSQLModelV2RequestDataAttributesSubjectTypesItems, timestampColumn string, variantColumn string) *ExperimentsUpdateExposureSQLModelV2RequestDataAttributes {
	this := ExperimentsUpdateExposureSQLModelV2RequestDataAttributes{}
	this.ExperimentColumn = experimentColumn
	this.Name = name
	this.Sql = sql
	this.SubjectTypes = subjectTypes
	this.TimestampColumn = timestampColumn
	this.VariantColumn = variantColumn
	return &this
}

// NewExperimentsUpdateExposureSQLModelV2RequestDataAttributesWithDefaults instantiates a new ExperimentsUpdateExposureSQLModelV2RequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsUpdateExposureSQLModelV2RequestDataAttributesWithDefaults() *ExperimentsUpdateExposureSQLModelV2RequestDataAttributes {
	this := ExperimentsUpdateExposureSQLModelV2RequestDataAttributes{}
	return &this
}

// GetDatePartitionColumn returns the DatePartitionColumn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsUpdateExposureSQLModelV2RequestDataAttributes) GetDatePartitionColumn() string {
	if o == nil || o.DatePartitionColumn.Get() == nil {
		var ret string
		return ret
	}
	return *o.DatePartitionColumn.Get()
}

// GetDatePartitionColumnOk returns a tuple with the DatePartitionColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsUpdateExposureSQLModelV2RequestDataAttributes) GetDatePartitionColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DatePartitionColumn.Get(), o.DatePartitionColumn.IsSet()
}

// HasDatePartitionColumn returns a boolean if a field has been set.
func (o *ExperimentsUpdateExposureSQLModelV2RequestDataAttributes) HasDatePartitionColumn() bool {
	return o != nil && o.DatePartitionColumn.IsSet()
}

// SetDatePartitionColumn gets a reference to the given datadog.NullableString and assigns it to the DatePartitionColumn field.
func (o *ExperimentsUpdateExposureSQLModelV2RequestDataAttributes) SetDatePartitionColumn(v string) {
	o.DatePartitionColumn.Set(&v)
}

// SetDatePartitionColumnNil sets the value for DatePartitionColumn to be an explicit nil.
func (o *ExperimentsUpdateExposureSQLModelV2RequestDataAttributes) SetDatePartitionColumnNil() {
	o.DatePartitionColumn.Set(nil)
}

// UnsetDatePartitionColumn ensures that no value is present for DatePartitionColumn, not even an explicit nil.
func (o *ExperimentsUpdateExposureSQLModelV2RequestDataAttributes) UnsetDatePartitionColumn() {
	o.DatePartitionColumn.Unset()
}

// GetExperimentColumn returns the ExperimentColumn field value.
func (o *ExperimentsUpdateExposureSQLModelV2RequestDataAttributes) GetExperimentColumn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ExperimentColumn
}

// GetExperimentColumnOk returns a tuple with the ExperimentColumn field value
// and a boolean to check if the value has been set.
func (o *ExperimentsUpdateExposureSQLModelV2RequestDataAttributes) GetExperimentColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExperimentColumn, true
}

// SetExperimentColumn sets field value.
func (o *ExperimentsUpdateExposureSQLModelV2RequestDataAttributes) SetExperimentColumn(v string) {
	o.ExperimentColumn = v
}

// GetMigrationMetadata returns the MigrationMetadata field value if set, zero value otherwise.
func (o *ExperimentsUpdateExposureSQLModelV2RequestDataAttributes) GetMigrationMetadata() interface{} {
	if o == nil || o.MigrationMetadata == nil {
		var ret interface{}
		return ret
	}
	return o.MigrationMetadata
}

// GetMigrationMetadataOk returns a tuple with the MigrationMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsUpdateExposureSQLModelV2RequestDataAttributes) GetMigrationMetadataOk() (*interface{}, bool) {
	if o == nil || o.MigrationMetadata == nil {
		return nil, false
	}
	return &o.MigrationMetadata, true
}

// HasMigrationMetadata returns a boolean if a field has been set.
func (o *ExperimentsUpdateExposureSQLModelV2RequestDataAttributes) HasMigrationMetadata() bool {
	return o != nil && o.MigrationMetadata != nil
}

// SetMigrationMetadata gets a reference to the given interface{} and assigns it to the MigrationMetadata field.
func (o *ExperimentsUpdateExposureSQLModelV2RequestDataAttributes) SetMigrationMetadata(v interface{}) {
	o.MigrationMetadata = v
}

// GetName returns the Name field value.
func (o *ExperimentsUpdateExposureSQLModelV2RequestDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExperimentsUpdateExposureSQLModelV2RequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ExperimentsUpdateExposureSQLModelV2RequestDataAttributes) SetName(v string) {
	o.Name = v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *ExperimentsUpdateExposureSQLModelV2RequestDataAttributes) GetProperties() []ExperimentsSQLModelPropertyInput {
	if o == nil || o.Properties == nil {
		var ret []ExperimentsSQLModelPropertyInput
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsUpdateExposureSQLModelV2RequestDataAttributes) GetPropertiesOk() (*[]ExperimentsSQLModelPropertyInput, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return &o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *ExperimentsUpdateExposureSQLModelV2RequestDataAttributes) HasProperties() bool {
	return o != nil && o.Properties != nil
}

// SetProperties gets a reference to the given []ExperimentsSQLModelPropertyInput and assigns it to the Properties field.
func (o *ExperimentsUpdateExposureSQLModelV2RequestDataAttributes) SetProperties(v []ExperimentsSQLModelPropertyInput) {
	o.Properties = v
}

// GetSql returns the Sql field value.
func (o *ExperimentsUpdateExposureSQLModelV2RequestDataAttributes) GetSql() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Sql
}

// GetSqlOk returns a tuple with the Sql field value
// and a boolean to check if the value has been set.
func (o *ExperimentsUpdateExposureSQLModelV2RequestDataAttributes) GetSqlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sql, true
}

// SetSql sets field value.
func (o *ExperimentsUpdateExposureSQLModelV2RequestDataAttributes) SetSql(v string) {
	o.Sql = v
}

// GetSubjectTypes returns the SubjectTypes field value.
func (o *ExperimentsUpdateExposureSQLModelV2RequestDataAttributes) GetSubjectTypes() []ExperimentsCreateExposureSQLModelV2RequestDataAttributesSubjectTypesItems {
	if o == nil {
		var ret []ExperimentsCreateExposureSQLModelV2RequestDataAttributesSubjectTypesItems
		return ret
	}
	return o.SubjectTypes
}

// GetSubjectTypesOk returns a tuple with the SubjectTypes field value
// and a boolean to check if the value has been set.
func (o *ExperimentsUpdateExposureSQLModelV2RequestDataAttributes) GetSubjectTypesOk() (*[]ExperimentsCreateExposureSQLModelV2RequestDataAttributesSubjectTypesItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubjectTypes, true
}

// SetSubjectTypes sets field value.
func (o *ExperimentsUpdateExposureSQLModelV2RequestDataAttributes) SetSubjectTypes(v []ExperimentsCreateExposureSQLModelV2RequestDataAttributesSubjectTypesItems) {
	o.SubjectTypes = v
}

// GetTimestampColumn returns the TimestampColumn field value.
func (o *ExperimentsUpdateExposureSQLModelV2RequestDataAttributes) GetTimestampColumn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TimestampColumn
}

// GetTimestampColumnOk returns a tuple with the TimestampColumn field value
// and a boolean to check if the value has been set.
func (o *ExperimentsUpdateExposureSQLModelV2RequestDataAttributes) GetTimestampColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimestampColumn, true
}

// SetTimestampColumn sets field value.
func (o *ExperimentsUpdateExposureSQLModelV2RequestDataAttributes) SetTimestampColumn(v string) {
	o.TimestampColumn = v
}

// GetVariantColumn returns the VariantColumn field value.
func (o *ExperimentsUpdateExposureSQLModelV2RequestDataAttributes) GetVariantColumn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.VariantColumn
}

// GetVariantColumnOk returns a tuple with the VariantColumn field value
// and a boolean to check if the value has been set.
func (o *ExperimentsUpdateExposureSQLModelV2RequestDataAttributes) GetVariantColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VariantColumn, true
}

// SetVariantColumn sets field value.
func (o *ExperimentsUpdateExposureSQLModelV2RequestDataAttributes) SetVariantColumn(v string) {
	o.VariantColumn = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsUpdateExposureSQLModelV2RequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DatePartitionColumn.IsSet() {
		toSerialize["date_partition_column"] = o.DatePartitionColumn.Get()
	}
	toSerialize["experiment_column"] = o.ExperimentColumn
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
	toSerialize["variant_column"] = o.VariantColumn

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsUpdateExposureSQLModelV2RequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DatePartitionColumn datadog.NullableString                                                       `json:"date_partition_column,omitempty"`
		ExperimentColumn    *string                                                                      `json:"experiment_column"`
		MigrationMetadata   interface{}                                                                  `json:"migration_metadata,omitempty"`
		Name                *string                                                                      `json:"name"`
		Properties          []ExperimentsSQLModelPropertyInput                                           `json:"properties,omitempty"`
		Sql                 *string                                                                      `json:"sql"`
		SubjectTypes        *[]ExperimentsCreateExposureSQLModelV2RequestDataAttributesSubjectTypesItems `json:"subject_types"`
		TimestampColumn     *string                                                                      `json:"timestamp_column"`
		VariantColumn       *string                                                                      `json:"variant_column"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ExperimentColumn == nil {
		return fmt.Errorf("required field experiment_column missing")
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
	if all.VariantColumn == nil {
		return fmt.Errorf("required field variant_column missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"date_partition_column", "experiment_column", "migration_metadata", "name", "properties", "sql", "subject_types", "timestamp_column", "variant_column"})
	} else {
		return err
	}
	o.DatePartitionColumn = all.DatePartitionColumn
	o.ExperimentColumn = *all.ExperimentColumn
	o.MigrationMetadata = all.MigrationMetadata
	o.Name = *all.Name
	o.Properties = all.Properties
	o.Sql = *all.Sql
	o.SubjectTypes = *all.SubjectTypes
	o.TimestampColumn = *all.TimestampColumn
	o.VariantColumn = *all.VariantColumn

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
