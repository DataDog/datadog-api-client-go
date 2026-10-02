// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems A measure available from a metric SQL model column.
type ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems struct {
	// Name of the SQL result column that supplies this measure.
	ColumnName *string `json:"column_name,omitempty"`
	// Data type of a column in the SQL model.
	ColumnType *ExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnType `json:"column_type,omitempty"`
	// Text that explains the measure.
	Description *string `json:"description,omitempty"`
	// ID of the measure.
	Id *string `json:"id,omitempty"`
	// Metadata retained for resources imported from another system.
	MigrationMetadata interface{} `json:"migration_metadata,omitempty"`
	// Display name of the measure.
	Name *string `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems instantiates a new ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems() *ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems {
	this := ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems{}
	return &this
}

// NewExperimentsMetricSQLModelV2DTODataAttributesMeasuresItemsWithDefaults instantiates a new ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsMetricSQLModelV2DTODataAttributesMeasuresItemsWithDefaults() *ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems {
	this := ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems{}
	return &this
}

// GetColumnName returns the ColumnName field value if set, zero value otherwise.
func (o *ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems) GetColumnName() string {
	if o == nil || o.ColumnName == nil {
		var ret string
		return ret
	}
	return *o.ColumnName
}

// GetColumnNameOk returns a tuple with the ColumnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems) GetColumnNameOk() (*string, bool) {
	if o == nil || o.ColumnName == nil {
		return nil, false
	}
	return o.ColumnName, true
}

// HasColumnName returns a boolean if a field has been set.
func (o *ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems) HasColumnName() bool {
	return o != nil && o.ColumnName != nil
}

// SetColumnName gets a reference to the given string and assigns it to the ColumnName field.
func (o *ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems) SetColumnName(v string) {
	o.ColumnName = &v
}

// GetColumnType returns the ColumnType field value if set, zero value otherwise.
func (o *ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems) GetColumnType() ExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnType {
	if o == nil || o.ColumnType == nil {
		var ret ExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnType
		return ret
	}
	return *o.ColumnType
}

// GetColumnTypeOk returns a tuple with the ColumnType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems) GetColumnTypeOk() (*ExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnType, bool) {
	if o == nil || o.ColumnType == nil {
		return nil, false
	}
	return o.ColumnType, true
}

// HasColumnType returns a boolean if a field has been set.
func (o *ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems) HasColumnType() bool {
	return o != nil && o.ColumnType != nil
}

// SetColumnType gets a reference to the given ExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnType and assigns it to the ColumnType field.
func (o *ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems) SetColumnType(v ExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnType) {
	o.ColumnType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems) SetId(v string) {
	o.Id = &v
}

// GetMigrationMetadata returns the MigrationMetadata field value if set, zero value otherwise.
func (o *ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems) GetMigrationMetadata() interface{} {
	if o == nil || o.MigrationMetadata == nil {
		var ret interface{}
		return ret
	}
	return o.MigrationMetadata
}

// GetMigrationMetadataOk returns a tuple with the MigrationMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems) GetMigrationMetadataOk() (*interface{}, bool) {
	if o == nil || o.MigrationMetadata == nil {
		return nil, false
	}
	return &o.MigrationMetadata, true
}

// HasMigrationMetadata returns a boolean if a field has been set.
func (o *ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems) HasMigrationMetadata() bool {
	return o != nil && o.MigrationMetadata != nil
}

// SetMigrationMetadata gets a reference to the given interface{} and assigns it to the MigrationMetadata field.
func (o *ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems) SetMigrationMetadata(v interface{}) {
	o.MigrationMetadata = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems) SetName(v string) {
	o.Name = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ColumnName != nil {
		toSerialize["column_name"] = o.ColumnName
	}
	if o.ColumnType != nil {
		toSerialize["column_type"] = o.ColumnType
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.MigrationMetadata != nil {
		toSerialize["migration_metadata"] = o.MigrationMetadata
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
func (o *ExperimentsMetricSQLModelV2DTODataAttributesMeasuresItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ColumnName        *string                                                                  `json:"column_name,omitempty"`
		ColumnType        *ExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnType `json:"column_type,omitempty"`
		Description       *string                                                                  `json:"description,omitempty"`
		Id                *string                                                                  `json:"id,omitempty"`
		MigrationMetadata interface{}                                                              `json:"migration_metadata,omitempty"`
		Name              *string                                                                  `json:"name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"column_name", "column_type", "description", "id", "migration_metadata", "name"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ColumnName = all.ColumnName
	if all.ColumnType != nil && !all.ColumnType.IsValid() {
		hasInvalidField = true
	} else {
		o.ColumnType = all.ColumnType
	}
	o.Description = all.Description
	o.Id = all.Id
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
