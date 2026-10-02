// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsExposureSQLModelV2DTODataAttributes Query and column mappings used to read experiment assignment data.
type ExperimentsExposureSQLModelV2DTODataAttributes struct {
	// Time when the exposure SQL model was archived.
	ArchivedAt datadog.NullableTime `json:"archived_at,omitempty"`
	// Time when the exposure SQL model was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Column used to identify date partitions in the exposure data.
	DatePartitionColumn datadog.NullableString `json:"date_partition_column,omitempty"`
	// SQL result column that contains the experiment key.
	ExperimentColumn *string `json:"experiment_column,omitempty"`
	// Number of experiments associated with the exposure SQL model.
	ExperimentCount datadog.NullableInt64 `json:"experiment_count,omitempty"`
	// Metadata associated with migration of this resource.
	MigrationMetadata interface{} `json:"migration_metadata,omitempty"`
	// Display name of the exposure SQL model.
	Name *string `json:"name,omitempty"`
	// Property columns available for filtering or splitting exposure data.
	Properties []ExperimentsExposureSQLModelV2DTODataAttributesItems `json:"properties,omitempty"`
	// SQL query that supplies the experiment assignment data.
	Sql *string `json:"sql,omitempty"`
	// Mappings between subject types and their identifier columns.
	SubjectTypes []ExperimentsCreateExposureSQLModelV2RequestDataAttributesSubjectTypesItems `json:"subject_types,omitempty"`
	// SQL result column that contains the assignment timestamp.
	TimestampColumn *string `json:"timestamp_column,omitempty"`
	// Time when the exposure SQL model was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// SQL result column that contains the assigned variant.
	VariantColumn *string `json:"variant_column,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsExposureSQLModelV2DTODataAttributes instantiates a new ExperimentsExposureSQLModelV2DTODataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsExposureSQLModelV2DTODataAttributes() *ExperimentsExposureSQLModelV2DTODataAttributes {
	this := ExperimentsExposureSQLModelV2DTODataAttributes{}
	return &this
}

// NewExperimentsExposureSQLModelV2DTODataAttributesWithDefaults instantiates a new ExperimentsExposureSQLModelV2DTODataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsExposureSQLModelV2DTODataAttributesWithDefaults() *ExperimentsExposureSQLModelV2DTODataAttributes {
	this := ExperimentsExposureSQLModelV2DTODataAttributes{}
	return &this
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) GetArchivedAt() time.Time {
	if o == nil || o.ArchivedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ArchivedAt.Get()
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) GetArchivedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ArchivedAt.Get(), o.ArchivedAt.IsSet()
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) HasArchivedAt() bool {
	return o != nil && o.ArchivedAt.IsSet()
}

// SetArchivedAt gets a reference to the given datadog.NullableTime and assigns it to the ArchivedAt field.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) SetArchivedAt(v time.Time) {
	o.ArchivedAt.Set(&v)
}

// SetArchivedAtNil sets the value for ArchivedAt to be an explicit nil.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) SetArchivedAtNil() {
	o.ArchivedAt.Set(nil)
}

// UnsetArchivedAt ensures that no value is present for ArchivedAt, not even an explicit nil.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) UnsetArchivedAt() {
	o.ArchivedAt.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDatePartitionColumn returns the DatePartitionColumn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) GetDatePartitionColumn() string {
	if o == nil || o.DatePartitionColumn.Get() == nil {
		var ret string
		return ret
	}
	return *o.DatePartitionColumn.Get()
}

// GetDatePartitionColumnOk returns a tuple with the DatePartitionColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) GetDatePartitionColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DatePartitionColumn.Get(), o.DatePartitionColumn.IsSet()
}

// HasDatePartitionColumn returns a boolean if a field has been set.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) HasDatePartitionColumn() bool {
	return o != nil && o.DatePartitionColumn.IsSet()
}

// SetDatePartitionColumn gets a reference to the given datadog.NullableString and assigns it to the DatePartitionColumn field.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) SetDatePartitionColumn(v string) {
	o.DatePartitionColumn.Set(&v)
}

// SetDatePartitionColumnNil sets the value for DatePartitionColumn to be an explicit nil.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) SetDatePartitionColumnNil() {
	o.DatePartitionColumn.Set(nil)
}

// UnsetDatePartitionColumn ensures that no value is present for DatePartitionColumn, not even an explicit nil.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) UnsetDatePartitionColumn() {
	o.DatePartitionColumn.Unset()
}

// GetExperimentColumn returns the ExperimentColumn field value if set, zero value otherwise.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) GetExperimentColumn() string {
	if o == nil || o.ExperimentColumn == nil {
		var ret string
		return ret
	}
	return *o.ExperimentColumn
}

// GetExperimentColumnOk returns a tuple with the ExperimentColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) GetExperimentColumnOk() (*string, bool) {
	if o == nil || o.ExperimentColumn == nil {
		return nil, false
	}
	return o.ExperimentColumn, true
}

// HasExperimentColumn returns a boolean if a field has been set.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) HasExperimentColumn() bool {
	return o != nil && o.ExperimentColumn != nil
}

// SetExperimentColumn gets a reference to the given string and assigns it to the ExperimentColumn field.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) SetExperimentColumn(v string) {
	o.ExperimentColumn = &v
}

// GetExperimentCount returns the ExperimentCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) GetExperimentCount() int64 {
	if o == nil || o.ExperimentCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ExperimentCount.Get()
}

// GetExperimentCountOk returns a tuple with the ExperimentCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) GetExperimentCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExperimentCount.Get(), o.ExperimentCount.IsSet()
}

// HasExperimentCount returns a boolean if a field has been set.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) HasExperimentCount() bool {
	return o != nil && o.ExperimentCount.IsSet()
}

// SetExperimentCount gets a reference to the given datadog.NullableInt64 and assigns it to the ExperimentCount field.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) SetExperimentCount(v int64) {
	o.ExperimentCount.Set(&v)
}

// SetExperimentCountNil sets the value for ExperimentCount to be an explicit nil.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) SetExperimentCountNil() {
	o.ExperimentCount.Set(nil)
}

// UnsetExperimentCount ensures that no value is present for ExperimentCount, not even an explicit nil.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) UnsetExperimentCount() {
	o.ExperimentCount.Unset()
}

// GetMigrationMetadata returns the MigrationMetadata field value if set, zero value otherwise.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) GetMigrationMetadata() interface{} {
	if o == nil || o.MigrationMetadata == nil {
		var ret interface{}
		return ret
	}
	return o.MigrationMetadata
}

// GetMigrationMetadataOk returns a tuple with the MigrationMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) GetMigrationMetadataOk() (*interface{}, bool) {
	if o == nil || o.MigrationMetadata == nil {
		return nil, false
	}
	return &o.MigrationMetadata, true
}

// HasMigrationMetadata returns a boolean if a field has been set.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) HasMigrationMetadata() bool {
	return o != nil && o.MigrationMetadata != nil
}

// SetMigrationMetadata gets a reference to the given interface{} and assigns it to the MigrationMetadata field.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) SetMigrationMetadata(v interface{}) {
	o.MigrationMetadata = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) SetName(v string) {
	o.Name = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) GetProperties() []ExperimentsExposureSQLModelV2DTODataAttributesItems {
	if o == nil || o.Properties == nil {
		var ret []ExperimentsExposureSQLModelV2DTODataAttributesItems
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) GetPropertiesOk() (*[]ExperimentsExposureSQLModelV2DTODataAttributesItems, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return &o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) HasProperties() bool {
	return o != nil && o.Properties != nil
}

// SetProperties gets a reference to the given []ExperimentsExposureSQLModelV2DTODataAttributesItems and assigns it to the Properties field.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) SetProperties(v []ExperimentsExposureSQLModelV2DTODataAttributesItems) {
	o.Properties = v
}

// GetSql returns the Sql field value if set, zero value otherwise.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) GetSql() string {
	if o == nil || o.Sql == nil {
		var ret string
		return ret
	}
	return *o.Sql
}

// GetSqlOk returns a tuple with the Sql field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) GetSqlOk() (*string, bool) {
	if o == nil || o.Sql == nil {
		return nil, false
	}
	return o.Sql, true
}

// HasSql returns a boolean if a field has been set.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) HasSql() bool {
	return o != nil && o.Sql != nil
}

// SetSql gets a reference to the given string and assigns it to the Sql field.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) SetSql(v string) {
	o.Sql = &v
}

// GetSubjectTypes returns the SubjectTypes field value if set, zero value otherwise.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) GetSubjectTypes() []ExperimentsCreateExposureSQLModelV2RequestDataAttributesSubjectTypesItems {
	if o == nil || o.SubjectTypes == nil {
		var ret []ExperimentsCreateExposureSQLModelV2RequestDataAttributesSubjectTypesItems
		return ret
	}
	return o.SubjectTypes
}

// GetSubjectTypesOk returns a tuple with the SubjectTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) GetSubjectTypesOk() (*[]ExperimentsCreateExposureSQLModelV2RequestDataAttributesSubjectTypesItems, bool) {
	if o == nil || o.SubjectTypes == nil {
		return nil, false
	}
	return &o.SubjectTypes, true
}

// HasSubjectTypes returns a boolean if a field has been set.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) HasSubjectTypes() bool {
	return o != nil && o.SubjectTypes != nil
}

// SetSubjectTypes gets a reference to the given []ExperimentsCreateExposureSQLModelV2RequestDataAttributesSubjectTypesItems and assigns it to the SubjectTypes field.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) SetSubjectTypes(v []ExperimentsCreateExposureSQLModelV2RequestDataAttributesSubjectTypesItems) {
	o.SubjectTypes = v
}

// GetTimestampColumn returns the TimestampColumn field value if set, zero value otherwise.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) GetTimestampColumn() string {
	if o == nil || o.TimestampColumn == nil {
		var ret string
		return ret
	}
	return *o.TimestampColumn
}

// GetTimestampColumnOk returns a tuple with the TimestampColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) GetTimestampColumnOk() (*string, bool) {
	if o == nil || o.TimestampColumn == nil {
		return nil, false
	}
	return o.TimestampColumn, true
}

// HasTimestampColumn returns a boolean if a field has been set.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) HasTimestampColumn() bool {
	return o != nil && o.TimestampColumn != nil
}

// SetTimestampColumn gets a reference to the given string and assigns it to the TimestampColumn field.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) SetTimestampColumn(v string) {
	o.TimestampColumn = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetVariantColumn returns the VariantColumn field value if set, zero value otherwise.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) GetVariantColumn() string {
	if o == nil || o.VariantColumn == nil {
		var ret string
		return ret
	}
	return *o.VariantColumn
}

// GetVariantColumnOk returns a tuple with the VariantColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) GetVariantColumnOk() (*string, bool) {
	if o == nil || o.VariantColumn == nil {
		return nil, false
	}
	return o.VariantColumn, true
}

// HasVariantColumn returns a boolean if a field has been set.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) HasVariantColumn() bool {
	return o != nil && o.VariantColumn != nil
}

// SetVariantColumn gets a reference to the given string and assigns it to the VariantColumn field.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) SetVariantColumn(v string) {
	o.VariantColumn = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsExposureSQLModelV2DTODataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ArchivedAt.IsSet() {
		toSerialize["archived_at"] = o.ArchivedAt.Get()
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.DatePartitionColumn.IsSet() {
		toSerialize["date_partition_column"] = o.DatePartitionColumn.Get()
	}
	if o.ExperimentColumn != nil {
		toSerialize["experiment_column"] = o.ExperimentColumn
	}
	if o.ExperimentCount.IsSet() {
		toSerialize["experiment_count"] = o.ExperimentCount.Get()
	}
	if o.MigrationMetadata != nil {
		toSerialize["migration_metadata"] = o.MigrationMetadata
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.Sql != nil {
		toSerialize["sql"] = o.Sql
	}
	if o.SubjectTypes != nil {
		toSerialize["subject_types"] = o.SubjectTypes
	}
	if o.TimestampColumn != nil {
		toSerialize["timestamp_column"] = o.TimestampColumn
	}
	if o.UpdatedAt != nil {
		if o.UpdatedAt.Nanosecond() == 0 {
			toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.VariantColumn != nil {
		toSerialize["variant_column"] = o.VariantColumn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsExposureSQLModelV2DTODataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ArchivedAt          datadog.NullableTime                                                        `json:"archived_at,omitempty"`
		CreatedAt           *time.Time                                                                  `json:"created_at,omitempty"`
		DatePartitionColumn datadog.NullableString                                                      `json:"date_partition_column,omitempty"`
		ExperimentColumn    *string                                                                     `json:"experiment_column,omitempty"`
		ExperimentCount     datadog.NullableInt64                                                       `json:"experiment_count,omitempty"`
		MigrationMetadata   interface{}                                                                 `json:"migration_metadata,omitempty"`
		Name                *string                                                                     `json:"name,omitempty"`
		Properties          []ExperimentsExposureSQLModelV2DTODataAttributesItems                       `json:"properties,omitempty"`
		Sql                 *string                                                                     `json:"sql,omitempty"`
		SubjectTypes        []ExperimentsCreateExposureSQLModelV2RequestDataAttributesSubjectTypesItems `json:"subject_types,omitempty"`
		TimestampColumn     *string                                                                     `json:"timestamp_column,omitempty"`
		UpdatedAt           *time.Time                                                                  `json:"updated_at,omitempty"`
		VariantColumn       *string                                                                     `json:"variant_column,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"archived_at", "created_at", "date_partition_column", "experiment_column", "experiment_count", "migration_metadata", "name", "properties", "sql", "subject_types", "timestamp_column", "updated_at", "variant_column"})
	} else {
		return err
	}
	o.ArchivedAt = all.ArchivedAt
	o.CreatedAt = all.CreatedAt
	o.DatePartitionColumn = all.DatePartitionColumn
	o.ExperimentColumn = all.ExperimentColumn
	o.ExperimentCount = all.ExperimentCount
	o.MigrationMetadata = all.MigrationMetadata
	o.Name = all.Name
	o.Properties = all.Properties
	o.Sql = all.Sql
	o.SubjectTypes = all.SubjectTypes
	o.TimestampColumn = all.TimestampColumn
	o.UpdatedAt = all.UpdatedAt
	o.VariantColumn = all.VariantColumn

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
