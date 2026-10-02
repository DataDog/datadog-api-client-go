// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsSubjectTypeV2DTODataAttributes Details of the subject type.
type ExperimentsSubjectTypeV2DTODataAttributes struct {
	// Time when this resource was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Number of experiments that reference this resource.
	ExperimentCount datadog.NullableInt64 `json:"experiment_count,omitempty"`
	// Number of exposure sources that reference this subject type.
	ExposureSourceCount datadog.NullableInt64 `json:"exposure_source_count,omitempty"`
	// Whether this is the organization's default subject type.
	IsDefault *bool `json:"is_default,omitempty"`
	// Number of metric SQL models that reference this subject type.
	MetricSqlModelCount datadog.NullableInt64 `json:"metric_sql_model_count,omitempty"`
	// Metadata retained for resources imported from another system.
	MigrationMetadata interface{} `json:"migration_metadata,omitempty"`
	// Display name of the subject type.
	Name *string `json:"name,omitempty"`
	// Product Analytics attribute used to identify subjects of this type.
	ProductAnalyticsAttribute *string `json:"product_analytics_attribute,omitempty"`
	// Number of protocols that reference this subject type.
	ProtocolCount datadog.NullableInt64 `json:"protocol_count,omitempty"`
	// Time when this resource was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Warehouse columns that identify subjects of this type.
	WarehouseColumnNames []string `json:"warehouse_column_names,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsSubjectTypeV2DTODataAttributes instantiates a new ExperimentsSubjectTypeV2DTODataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsSubjectTypeV2DTODataAttributes() *ExperimentsSubjectTypeV2DTODataAttributes {
	this := ExperimentsSubjectTypeV2DTODataAttributes{}
	return &this
}

// NewExperimentsSubjectTypeV2DTODataAttributesWithDefaults instantiates a new ExperimentsSubjectTypeV2DTODataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsSubjectTypeV2DTODataAttributesWithDefaults() *ExperimentsSubjectTypeV2DTODataAttributes {
	this := ExperimentsSubjectTypeV2DTODataAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetExperimentCount returns the ExperimentCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsSubjectTypeV2DTODataAttributes) GetExperimentCount() int64 {
	if o == nil || o.ExperimentCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ExperimentCount.Get()
}

// GetExperimentCountOk returns a tuple with the ExperimentCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) GetExperimentCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExperimentCount.Get(), o.ExperimentCount.IsSet()
}

// HasExperimentCount returns a boolean if a field has been set.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) HasExperimentCount() bool {
	return o != nil && o.ExperimentCount.IsSet()
}

// SetExperimentCount gets a reference to the given datadog.NullableInt64 and assigns it to the ExperimentCount field.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) SetExperimentCount(v int64) {
	o.ExperimentCount.Set(&v)
}

// SetExperimentCountNil sets the value for ExperimentCount to be an explicit nil.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) SetExperimentCountNil() {
	o.ExperimentCount.Set(nil)
}

// UnsetExperimentCount ensures that no value is present for ExperimentCount, not even an explicit nil.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) UnsetExperimentCount() {
	o.ExperimentCount.Unset()
}

// GetExposureSourceCount returns the ExposureSourceCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsSubjectTypeV2DTODataAttributes) GetExposureSourceCount() int64 {
	if o == nil || o.ExposureSourceCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ExposureSourceCount.Get()
}

// GetExposureSourceCountOk returns a tuple with the ExposureSourceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) GetExposureSourceCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExposureSourceCount.Get(), o.ExposureSourceCount.IsSet()
}

// HasExposureSourceCount returns a boolean if a field has been set.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) HasExposureSourceCount() bool {
	return o != nil && o.ExposureSourceCount.IsSet()
}

// SetExposureSourceCount gets a reference to the given datadog.NullableInt64 and assigns it to the ExposureSourceCount field.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) SetExposureSourceCount(v int64) {
	o.ExposureSourceCount.Set(&v)
}

// SetExposureSourceCountNil sets the value for ExposureSourceCount to be an explicit nil.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) SetExposureSourceCountNil() {
	o.ExposureSourceCount.Set(nil)
}

// UnsetExposureSourceCount ensures that no value is present for ExposureSourceCount, not even an explicit nil.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) UnsetExposureSourceCount() {
	o.ExposureSourceCount.Unset()
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) GetIsDefault() bool {
	if o == nil || o.IsDefault == nil {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) GetIsDefaultOk() (*bool, bool) {
	if o == nil || o.IsDefault == nil {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) HasIsDefault() bool {
	return o != nil && o.IsDefault != nil
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetMetricSqlModelCount returns the MetricSqlModelCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsSubjectTypeV2DTODataAttributes) GetMetricSqlModelCount() int64 {
	if o == nil || o.MetricSqlModelCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.MetricSqlModelCount.Get()
}

// GetMetricSqlModelCountOk returns a tuple with the MetricSqlModelCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) GetMetricSqlModelCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetricSqlModelCount.Get(), o.MetricSqlModelCount.IsSet()
}

// HasMetricSqlModelCount returns a boolean if a field has been set.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) HasMetricSqlModelCount() bool {
	return o != nil && o.MetricSqlModelCount.IsSet()
}

// SetMetricSqlModelCount gets a reference to the given datadog.NullableInt64 and assigns it to the MetricSqlModelCount field.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) SetMetricSqlModelCount(v int64) {
	o.MetricSqlModelCount.Set(&v)
}

// SetMetricSqlModelCountNil sets the value for MetricSqlModelCount to be an explicit nil.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) SetMetricSqlModelCountNil() {
	o.MetricSqlModelCount.Set(nil)
}

// UnsetMetricSqlModelCount ensures that no value is present for MetricSqlModelCount, not even an explicit nil.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) UnsetMetricSqlModelCount() {
	o.MetricSqlModelCount.Unset()
}

// GetMigrationMetadata returns the MigrationMetadata field value if set, zero value otherwise.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) GetMigrationMetadata() interface{} {
	if o == nil || o.MigrationMetadata == nil {
		var ret interface{}
		return ret
	}
	return o.MigrationMetadata
}

// GetMigrationMetadataOk returns a tuple with the MigrationMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) GetMigrationMetadataOk() (*interface{}, bool) {
	if o == nil || o.MigrationMetadata == nil {
		return nil, false
	}
	return &o.MigrationMetadata, true
}

// HasMigrationMetadata returns a boolean if a field has been set.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) HasMigrationMetadata() bool {
	return o != nil && o.MigrationMetadata != nil
}

// SetMigrationMetadata gets a reference to the given interface{} and assigns it to the MigrationMetadata field.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) SetMigrationMetadata(v interface{}) {
	o.MigrationMetadata = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) SetName(v string) {
	o.Name = &v
}

// GetProductAnalyticsAttribute returns the ProductAnalyticsAttribute field value if set, zero value otherwise.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) GetProductAnalyticsAttribute() string {
	if o == nil || o.ProductAnalyticsAttribute == nil {
		var ret string
		return ret
	}
	return *o.ProductAnalyticsAttribute
}

// GetProductAnalyticsAttributeOk returns a tuple with the ProductAnalyticsAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) GetProductAnalyticsAttributeOk() (*string, bool) {
	if o == nil || o.ProductAnalyticsAttribute == nil {
		return nil, false
	}
	return o.ProductAnalyticsAttribute, true
}

// HasProductAnalyticsAttribute returns a boolean if a field has been set.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) HasProductAnalyticsAttribute() bool {
	return o != nil && o.ProductAnalyticsAttribute != nil
}

// SetProductAnalyticsAttribute gets a reference to the given string and assigns it to the ProductAnalyticsAttribute field.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) SetProductAnalyticsAttribute(v string) {
	o.ProductAnalyticsAttribute = &v
}

// GetProtocolCount returns the ProtocolCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsSubjectTypeV2DTODataAttributes) GetProtocolCount() int64 {
	if o == nil || o.ProtocolCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ProtocolCount.Get()
}

// GetProtocolCountOk returns a tuple with the ProtocolCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) GetProtocolCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProtocolCount.Get(), o.ProtocolCount.IsSet()
}

// HasProtocolCount returns a boolean if a field has been set.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) HasProtocolCount() bool {
	return o != nil && o.ProtocolCount.IsSet()
}

// SetProtocolCount gets a reference to the given datadog.NullableInt64 and assigns it to the ProtocolCount field.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) SetProtocolCount(v int64) {
	o.ProtocolCount.Set(&v)
}

// SetProtocolCountNil sets the value for ProtocolCount to be an explicit nil.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) SetProtocolCountNil() {
	o.ProtocolCount.Set(nil)
}

// UnsetProtocolCount ensures that no value is present for ProtocolCount, not even an explicit nil.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) UnsetProtocolCount() {
	o.ProtocolCount.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetWarehouseColumnNames returns the WarehouseColumnNames field value if set, zero value otherwise.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) GetWarehouseColumnNames() []string {
	if o == nil || o.WarehouseColumnNames == nil {
		var ret []string
		return ret
	}
	return o.WarehouseColumnNames
}

// GetWarehouseColumnNamesOk returns a tuple with the WarehouseColumnNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) GetWarehouseColumnNamesOk() (*[]string, bool) {
	if o == nil || o.WarehouseColumnNames == nil {
		return nil, false
	}
	return &o.WarehouseColumnNames, true
}

// HasWarehouseColumnNames returns a boolean if a field has been set.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) HasWarehouseColumnNames() bool {
	return o != nil && o.WarehouseColumnNames != nil
}

// SetWarehouseColumnNames gets a reference to the given []string and assigns it to the WarehouseColumnNames field.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) SetWarehouseColumnNames(v []string) {
	o.WarehouseColumnNames = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsSubjectTypeV2DTODataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.ExperimentCount.IsSet() {
		toSerialize["experiment_count"] = o.ExperimentCount.Get()
	}
	if o.ExposureSourceCount.IsSet() {
		toSerialize["exposure_source_count"] = o.ExposureSourceCount.Get()
	}
	if o.IsDefault != nil {
		toSerialize["is_default"] = o.IsDefault
	}
	if o.MetricSqlModelCount.IsSet() {
		toSerialize["metric_sql_model_count"] = o.MetricSqlModelCount.Get()
	}
	if o.MigrationMetadata != nil {
		toSerialize["migration_metadata"] = o.MigrationMetadata
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ProductAnalyticsAttribute != nil {
		toSerialize["product_analytics_attribute"] = o.ProductAnalyticsAttribute
	}
	if o.ProtocolCount.IsSet() {
		toSerialize["protocol_count"] = o.ProtocolCount.Get()
	}
	if o.UpdatedAt != nil {
		if o.UpdatedAt.Nanosecond() == 0 {
			toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.WarehouseColumnNames != nil {
		toSerialize["warehouse_column_names"] = o.WarehouseColumnNames
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsSubjectTypeV2DTODataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt                 *time.Time            `json:"created_at,omitempty"`
		ExperimentCount           datadog.NullableInt64 `json:"experiment_count,omitempty"`
		ExposureSourceCount       datadog.NullableInt64 `json:"exposure_source_count,omitempty"`
		IsDefault                 *bool                 `json:"is_default,omitempty"`
		MetricSqlModelCount       datadog.NullableInt64 `json:"metric_sql_model_count,omitempty"`
		MigrationMetadata         interface{}           `json:"migration_metadata,omitempty"`
		Name                      *string               `json:"name,omitempty"`
		ProductAnalyticsAttribute *string               `json:"product_analytics_attribute,omitempty"`
		ProtocolCount             datadog.NullableInt64 `json:"protocol_count,omitempty"`
		UpdatedAt                 *time.Time            `json:"updated_at,omitempty"`
		WarehouseColumnNames      []string              `json:"warehouse_column_names,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "experiment_count", "exposure_source_count", "is_default", "metric_sql_model_count", "migration_metadata", "name", "product_analytics_attribute", "protocol_count", "updated_at", "warehouse_column_names"})
	} else {
		return err
	}
	o.CreatedAt = all.CreatedAt
	o.ExperimentCount = all.ExperimentCount
	o.ExposureSourceCount = all.ExposureSourceCount
	o.IsDefault = all.IsDefault
	o.MetricSqlModelCount = all.MetricSqlModelCount
	o.MigrationMetadata = all.MigrationMetadata
	o.Name = all.Name
	o.ProductAnalyticsAttribute = all.ProductAnalyticsAttribute
	o.ProtocolCount = all.ProtocolCount
	o.UpdatedAt = all.UpdatedAt
	o.WarehouseColumnNames = all.WarehouseColumnNames

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
