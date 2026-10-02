// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsMetricSQLModelV2DTODataAttributes Details of the metric SQL model.
type ExperimentsMetricSQLModelV2DTODataAttributes struct {
	// Time when this resource was certified.
	CertifiedAt datadog.NullableTime `json:"certified_at,omitempty"`
	// Time when this resource was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// SQL column used to partition the source data by date.
	DatePartitionColumn datadog.NullableString `json:"date_partition_column,omitempty"`
	// Text that explains the metric SQL model.
	Description datadog.NullableString `json:"description,omitempty"`
	// Read-only measure ID. Pass it as warehouse_metric_measure.id when the metric operation is count.
	EventCountMeasureId *string `json:"event_count_measure_id,omitempty"`
	// Number of experiments that reference this resource.
	ExperimentCount datadog.NullableInt64 `json:"experiment_count,omitempty"`
	// Whether this resource has been certified.
	IsCertified *bool `json:"is_certified,omitempty"`
	// Measures available from the SQL model's result columns.
	Measures []ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems `json:"measures,omitempty"`
	// Number of metrics that use this SQL model.
	MetricCount datadog.NullableInt64 `json:"metric_count,omitempty"`
	// Metadata retained for resources imported from another system.
	MigrationMetadata interface{} `json:"migration_metadata,omitempty"`
	// Display name of the metric SQL model.
	Name *string `json:"name,omitempty"`
	// Property columns exposed by the SQL model.
	Properties []ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems `json:"properties,omitempty"`
	// SQL query that produces the model's source data.
	Sql *string `json:"sql,omitempty"`
	// Subject types mapped to columns in the SQL model.
	SubjectTypes []ExperimentsMetricSQLModelV2DTODataAttributesSubjectTypesItems `json:"subject_types,omitempty"`
	// SQL column that supplies the event timestamp.
	TimestampColumn *string `json:"timestamp_column,omitempty"`
	// Time when this resource was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsMetricSQLModelV2DTODataAttributes instantiates a new ExperimentsMetricSQLModelV2DTODataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsMetricSQLModelV2DTODataAttributes() *ExperimentsMetricSQLModelV2DTODataAttributes {
	this := ExperimentsMetricSQLModelV2DTODataAttributes{}
	return &this
}

// NewExperimentsMetricSQLModelV2DTODataAttributesWithDefaults instantiates a new ExperimentsMetricSQLModelV2DTODataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsMetricSQLModelV2DTODataAttributesWithDefaults() *ExperimentsMetricSQLModelV2DTODataAttributes {
	this := ExperimentsMetricSQLModelV2DTODataAttributes{}
	return &this
}

// GetCertifiedAt returns the CertifiedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) GetCertifiedAt() time.Time {
	if o == nil || o.CertifiedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CertifiedAt.Get()
}

// GetCertifiedAtOk returns a tuple with the CertifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) GetCertifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CertifiedAt.Get(), o.CertifiedAt.IsSet()
}

// HasCertifiedAt returns a boolean if a field has been set.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) HasCertifiedAt() bool {
	return o != nil && o.CertifiedAt.IsSet()
}

// SetCertifiedAt gets a reference to the given datadog.NullableTime and assigns it to the CertifiedAt field.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) SetCertifiedAt(v time.Time) {
	o.CertifiedAt.Set(&v)
}

// SetCertifiedAtNil sets the value for CertifiedAt to be an explicit nil.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) SetCertifiedAtNil() {
	o.CertifiedAt.Set(nil)
}

// UnsetCertifiedAt ensures that no value is present for CertifiedAt, not even an explicit nil.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) UnsetCertifiedAt() {
	o.CertifiedAt.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDatePartitionColumn returns the DatePartitionColumn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) GetDatePartitionColumn() string {
	if o == nil || o.DatePartitionColumn.Get() == nil {
		var ret string
		return ret
	}
	return *o.DatePartitionColumn.Get()
}

// GetDatePartitionColumnOk returns a tuple with the DatePartitionColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) GetDatePartitionColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DatePartitionColumn.Get(), o.DatePartitionColumn.IsSet()
}

// HasDatePartitionColumn returns a boolean if a field has been set.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) HasDatePartitionColumn() bool {
	return o != nil && o.DatePartitionColumn.IsSet()
}

// SetDatePartitionColumn gets a reference to the given datadog.NullableString and assigns it to the DatePartitionColumn field.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) SetDatePartitionColumn(v string) {
	o.DatePartitionColumn.Set(&v)
}

// SetDatePartitionColumnNil sets the value for DatePartitionColumn to be an explicit nil.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) SetDatePartitionColumnNil() {
	o.DatePartitionColumn.Set(nil)
}

// UnsetDatePartitionColumn ensures that no value is present for DatePartitionColumn, not even an explicit nil.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) UnsetDatePartitionColumn() {
	o.DatePartitionColumn.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given datadog.NullableString and assigns it to the Description field.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) UnsetDescription() {
	o.Description.Unset()
}

// GetEventCountMeasureId returns the EventCountMeasureId field value if set, zero value otherwise.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) GetEventCountMeasureId() string {
	if o == nil || o.EventCountMeasureId == nil {
		var ret string
		return ret
	}
	return *o.EventCountMeasureId
}

// GetEventCountMeasureIdOk returns a tuple with the EventCountMeasureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) GetEventCountMeasureIdOk() (*string, bool) {
	if o == nil || o.EventCountMeasureId == nil {
		return nil, false
	}
	return o.EventCountMeasureId, true
}

// HasEventCountMeasureId returns a boolean if a field has been set.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) HasEventCountMeasureId() bool {
	return o != nil && o.EventCountMeasureId != nil
}

// SetEventCountMeasureId gets a reference to the given string and assigns it to the EventCountMeasureId field.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) SetEventCountMeasureId(v string) {
	o.EventCountMeasureId = &v
}

// GetExperimentCount returns the ExperimentCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) GetExperimentCount() int64 {
	if o == nil || o.ExperimentCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ExperimentCount.Get()
}

// GetExperimentCountOk returns a tuple with the ExperimentCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) GetExperimentCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExperimentCount.Get(), o.ExperimentCount.IsSet()
}

// HasExperimentCount returns a boolean if a field has been set.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) HasExperimentCount() bool {
	return o != nil && o.ExperimentCount.IsSet()
}

// SetExperimentCount gets a reference to the given datadog.NullableInt64 and assigns it to the ExperimentCount field.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) SetExperimentCount(v int64) {
	o.ExperimentCount.Set(&v)
}

// SetExperimentCountNil sets the value for ExperimentCount to be an explicit nil.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) SetExperimentCountNil() {
	o.ExperimentCount.Set(nil)
}

// UnsetExperimentCount ensures that no value is present for ExperimentCount, not even an explicit nil.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) UnsetExperimentCount() {
	o.ExperimentCount.Unset()
}

// GetIsCertified returns the IsCertified field value if set, zero value otherwise.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) GetIsCertified() bool {
	if o == nil || o.IsCertified == nil {
		var ret bool
		return ret
	}
	return *o.IsCertified
}

// GetIsCertifiedOk returns a tuple with the IsCertified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) GetIsCertifiedOk() (*bool, bool) {
	if o == nil || o.IsCertified == nil {
		return nil, false
	}
	return o.IsCertified, true
}

// HasIsCertified returns a boolean if a field has been set.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) HasIsCertified() bool {
	return o != nil && o.IsCertified != nil
}

// SetIsCertified gets a reference to the given bool and assigns it to the IsCertified field.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) SetIsCertified(v bool) {
	o.IsCertified = &v
}

// GetMeasures returns the Measures field value if set, zero value otherwise.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) GetMeasures() []ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems {
	if o == nil || o.Measures == nil {
		var ret []ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems
		return ret
	}
	return o.Measures
}

// GetMeasuresOk returns a tuple with the Measures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) GetMeasuresOk() (*[]ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems, bool) {
	if o == nil || o.Measures == nil {
		return nil, false
	}
	return &o.Measures, true
}

// HasMeasures returns a boolean if a field has been set.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) HasMeasures() bool {
	return o != nil && o.Measures != nil
}

// SetMeasures gets a reference to the given []ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems and assigns it to the Measures field.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) SetMeasures(v []ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems) {
	o.Measures = v
}

// GetMetricCount returns the MetricCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) GetMetricCount() int64 {
	if o == nil || o.MetricCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.MetricCount.Get()
}

// GetMetricCountOk returns a tuple with the MetricCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) GetMetricCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetricCount.Get(), o.MetricCount.IsSet()
}

// HasMetricCount returns a boolean if a field has been set.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) HasMetricCount() bool {
	return o != nil && o.MetricCount.IsSet()
}

// SetMetricCount gets a reference to the given datadog.NullableInt64 and assigns it to the MetricCount field.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) SetMetricCount(v int64) {
	o.MetricCount.Set(&v)
}

// SetMetricCountNil sets the value for MetricCount to be an explicit nil.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) SetMetricCountNil() {
	o.MetricCount.Set(nil)
}

// UnsetMetricCount ensures that no value is present for MetricCount, not even an explicit nil.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) UnsetMetricCount() {
	o.MetricCount.Unset()
}

// GetMigrationMetadata returns the MigrationMetadata field value if set, zero value otherwise.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) GetMigrationMetadata() interface{} {
	if o == nil || o.MigrationMetadata == nil {
		var ret interface{}
		return ret
	}
	return o.MigrationMetadata
}

// GetMigrationMetadataOk returns a tuple with the MigrationMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) GetMigrationMetadataOk() (*interface{}, bool) {
	if o == nil || o.MigrationMetadata == nil {
		return nil, false
	}
	return &o.MigrationMetadata, true
}

// HasMigrationMetadata returns a boolean if a field has been set.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) HasMigrationMetadata() bool {
	return o != nil && o.MigrationMetadata != nil
}

// SetMigrationMetadata gets a reference to the given interface{} and assigns it to the MigrationMetadata field.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) SetMigrationMetadata(v interface{}) {
	o.MigrationMetadata = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) SetName(v string) {
	o.Name = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) GetProperties() []ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems {
	if o == nil || o.Properties == nil {
		var ret []ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) GetPropertiesOk() (*[]ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return &o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) HasProperties() bool {
	return o != nil && o.Properties != nil
}

// SetProperties gets a reference to the given []ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems and assigns it to the Properties field.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) SetProperties(v []ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems) {
	o.Properties = v
}

// GetSql returns the Sql field value if set, zero value otherwise.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) GetSql() string {
	if o == nil || o.Sql == nil {
		var ret string
		return ret
	}
	return *o.Sql
}

// GetSqlOk returns a tuple with the Sql field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) GetSqlOk() (*string, bool) {
	if o == nil || o.Sql == nil {
		return nil, false
	}
	return o.Sql, true
}

// HasSql returns a boolean if a field has been set.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) HasSql() bool {
	return o != nil && o.Sql != nil
}

// SetSql gets a reference to the given string and assigns it to the Sql field.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) SetSql(v string) {
	o.Sql = &v
}

// GetSubjectTypes returns the SubjectTypes field value if set, zero value otherwise.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) GetSubjectTypes() []ExperimentsMetricSQLModelV2DTODataAttributesSubjectTypesItems {
	if o == nil || o.SubjectTypes == nil {
		var ret []ExperimentsMetricSQLModelV2DTODataAttributesSubjectTypesItems
		return ret
	}
	return o.SubjectTypes
}

// GetSubjectTypesOk returns a tuple with the SubjectTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) GetSubjectTypesOk() (*[]ExperimentsMetricSQLModelV2DTODataAttributesSubjectTypesItems, bool) {
	if o == nil || o.SubjectTypes == nil {
		return nil, false
	}
	return &o.SubjectTypes, true
}

// HasSubjectTypes returns a boolean if a field has been set.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) HasSubjectTypes() bool {
	return o != nil && o.SubjectTypes != nil
}

// SetSubjectTypes gets a reference to the given []ExperimentsMetricSQLModelV2DTODataAttributesSubjectTypesItems and assigns it to the SubjectTypes field.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) SetSubjectTypes(v []ExperimentsMetricSQLModelV2DTODataAttributesSubjectTypesItems) {
	o.SubjectTypes = v
}

// GetTimestampColumn returns the TimestampColumn field value if set, zero value otherwise.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) GetTimestampColumn() string {
	if o == nil || o.TimestampColumn == nil {
		var ret string
		return ret
	}
	return *o.TimestampColumn
}

// GetTimestampColumnOk returns a tuple with the TimestampColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) GetTimestampColumnOk() (*string, bool) {
	if o == nil || o.TimestampColumn == nil {
		return nil, false
	}
	return o.TimestampColumn, true
}

// HasTimestampColumn returns a boolean if a field has been set.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) HasTimestampColumn() bool {
	return o != nil && o.TimestampColumn != nil
}

// SetTimestampColumn gets a reference to the given string and assigns it to the TimestampColumn field.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) SetTimestampColumn(v string) {
	o.TimestampColumn = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsMetricSQLModelV2DTODataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CertifiedAt.IsSet() {
		toSerialize["certified_at"] = o.CertifiedAt.Get()
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
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.EventCountMeasureId != nil {
		toSerialize["event_count_measure_id"] = o.EventCountMeasureId
	}
	if o.ExperimentCount.IsSet() {
		toSerialize["experiment_count"] = o.ExperimentCount.Get()
	}
	if o.IsCertified != nil {
		toSerialize["is_certified"] = o.IsCertified
	}
	if o.Measures != nil {
		toSerialize["measures"] = o.Measures
	}
	if o.MetricCount.IsSet() {
		toSerialize["metric_count"] = o.MetricCount.Get()
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsMetricSQLModelV2DTODataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CertifiedAt         datadog.NullableTime                                            `json:"certified_at,omitempty"`
		CreatedAt           *time.Time                                                      `json:"created_at,omitempty"`
		DatePartitionColumn datadog.NullableString                                          `json:"date_partition_column,omitempty"`
		Description         datadog.NullableString                                          `json:"description,omitempty"`
		EventCountMeasureId *string                                                         `json:"event_count_measure_id,omitempty"`
		ExperimentCount     datadog.NullableInt64                                           `json:"experiment_count,omitempty"`
		IsCertified         *bool                                                           `json:"is_certified,omitempty"`
		Measures            []ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems     `json:"measures,omitempty"`
		MetricCount         datadog.NullableInt64                                           `json:"metric_count,omitempty"`
		MigrationMetadata   interface{}                                                     `json:"migration_metadata,omitempty"`
		Name                *string                                                         `json:"name,omitempty"`
		Properties          []ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems     `json:"properties,omitempty"`
		Sql                 *string                                                         `json:"sql,omitempty"`
		SubjectTypes        []ExperimentsMetricSQLModelV2DTODataAttributesSubjectTypesItems `json:"subject_types,omitempty"`
		TimestampColumn     *string                                                         `json:"timestamp_column,omitempty"`
		UpdatedAt           *time.Time                                                      `json:"updated_at,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"certified_at", "created_at", "date_partition_column", "description", "event_count_measure_id", "experiment_count", "is_certified", "measures", "metric_count", "migration_metadata", "name", "properties", "sql", "subject_types", "timestamp_column", "updated_at"})
	} else {
		return err
	}
	o.CertifiedAt = all.CertifiedAt
	o.CreatedAt = all.CreatedAt
	o.DatePartitionColumn = all.DatePartitionColumn
	o.Description = all.Description
	o.EventCountMeasureId = all.EventCountMeasureId
	o.ExperimentCount = all.ExperimentCount
	o.IsCertified = all.IsCertified
	o.Measures = all.Measures
	o.MetricCount = all.MetricCount
	o.MigrationMetadata = all.MigrationMetadata
	o.Name = all.Name
	o.Properties = all.Properties
	o.Sql = all.Sql
	o.SubjectTypes = all.SubjectTypes
	o.TimestampColumn = all.TimestampColumn
	o.UpdatedAt = all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
