// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsExposureSQLModelV2DTODataAttributesItems Property column available from the exposure SQL model.
type ExperimentsExposureSQLModelV2DTODataAttributesItems struct {
	// SQL result column that contains this property.
	ColumnName *string `json:"column_name,omitempty"`
	// Data type of the property column.
	ColumnType *string `json:"column_type,omitempty"`
	// Description of the exposure property.
	Description *string `json:"description,omitempty"`
	// Identifier of the exposure property.
	Id *string `json:"id,omitempty"`
	// Metadata associated with migration of this resource.
	MigrationMetadata interface{} `json:"migration_metadata,omitempty"`
	// Display name of the exposure property.
	Name *string `json:"name,omitempty"`
	// Suffix used for this property column in the analysis pipeline.
	PipelineColumnSuffix *string `json:"pipeline_column_suffix,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsExposureSQLModelV2DTODataAttributesItems instantiates a new ExperimentsExposureSQLModelV2DTODataAttributesItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsExposureSQLModelV2DTODataAttributesItems() *ExperimentsExposureSQLModelV2DTODataAttributesItems {
	this := ExperimentsExposureSQLModelV2DTODataAttributesItems{}
	return &this
}

// NewExperimentsExposureSQLModelV2DTODataAttributesItemsWithDefaults instantiates a new ExperimentsExposureSQLModelV2DTODataAttributesItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsExposureSQLModelV2DTODataAttributesItemsWithDefaults() *ExperimentsExposureSQLModelV2DTODataAttributesItems {
	this := ExperimentsExposureSQLModelV2DTODataAttributesItems{}
	return &this
}

// GetColumnName returns the ColumnName field value if set, zero value otherwise.
func (o *ExperimentsExposureSQLModelV2DTODataAttributesItems) GetColumnName() string {
	if o == nil || o.ColumnName == nil {
		var ret string
		return ret
	}
	return *o.ColumnName
}

// GetColumnNameOk returns a tuple with the ColumnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExposureSQLModelV2DTODataAttributesItems) GetColumnNameOk() (*string, bool) {
	if o == nil || o.ColumnName == nil {
		return nil, false
	}
	return o.ColumnName, true
}

// HasColumnName returns a boolean if a field has been set.
func (o *ExperimentsExposureSQLModelV2DTODataAttributesItems) HasColumnName() bool {
	return o != nil && o.ColumnName != nil
}

// SetColumnName gets a reference to the given string and assigns it to the ColumnName field.
func (o *ExperimentsExposureSQLModelV2DTODataAttributesItems) SetColumnName(v string) {
	o.ColumnName = &v
}

// GetColumnType returns the ColumnType field value if set, zero value otherwise.
func (o *ExperimentsExposureSQLModelV2DTODataAttributesItems) GetColumnType() string {
	if o == nil || o.ColumnType == nil {
		var ret string
		return ret
	}
	return *o.ColumnType
}

// GetColumnTypeOk returns a tuple with the ColumnType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExposureSQLModelV2DTODataAttributesItems) GetColumnTypeOk() (*string, bool) {
	if o == nil || o.ColumnType == nil {
		return nil, false
	}
	return o.ColumnType, true
}

// HasColumnType returns a boolean if a field has been set.
func (o *ExperimentsExposureSQLModelV2DTODataAttributesItems) HasColumnType() bool {
	return o != nil && o.ColumnType != nil
}

// SetColumnType gets a reference to the given string and assigns it to the ColumnType field.
func (o *ExperimentsExposureSQLModelV2DTODataAttributesItems) SetColumnType(v string) {
	o.ColumnType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ExperimentsExposureSQLModelV2DTODataAttributesItems) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExposureSQLModelV2DTODataAttributesItems) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ExperimentsExposureSQLModelV2DTODataAttributesItems) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ExperimentsExposureSQLModelV2DTODataAttributesItems) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ExperimentsExposureSQLModelV2DTODataAttributesItems) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExposureSQLModelV2DTODataAttributesItems) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ExperimentsExposureSQLModelV2DTODataAttributesItems) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ExperimentsExposureSQLModelV2DTODataAttributesItems) SetId(v string) {
	o.Id = &v
}

// GetMigrationMetadata returns the MigrationMetadata field value if set, zero value otherwise.
func (o *ExperimentsExposureSQLModelV2DTODataAttributesItems) GetMigrationMetadata() interface{} {
	if o == nil || o.MigrationMetadata == nil {
		var ret interface{}
		return ret
	}
	return o.MigrationMetadata
}

// GetMigrationMetadataOk returns a tuple with the MigrationMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExposureSQLModelV2DTODataAttributesItems) GetMigrationMetadataOk() (*interface{}, bool) {
	if o == nil || o.MigrationMetadata == nil {
		return nil, false
	}
	return &o.MigrationMetadata, true
}

// HasMigrationMetadata returns a boolean if a field has been set.
func (o *ExperimentsExposureSQLModelV2DTODataAttributesItems) HasMigrationMetadata() bool {
	return o != nil && o.MigrationMetadata != nil
}

// SetMigrationMetadata gets a reference to the given interface{} and assigns it to the MigrationMetadata field.
func (o *ExperimentsExposureSQLModelV2DTODataAttributesItems) SetMigrationMetadata(v interface{}) {
	o.MigrationMetadata = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExperimentsExposureSQLModelV2DTODataAttributesItems) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExposureSQLModelV2DTODataAttributesItems) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExperimentsExposureSQLModelV2DTODataAttributesItems) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExperimentsExposureSQLModelV2DTODataAttributesItems) SetName(v string) {
	o.Name = &v
}

// GetPipelineColumnSuffix returns the PipelineColumnSuffix field value if set, zero value otherwise.
func (o *ExperimentsExposureSQLModelV2DTODataAttributesItems) GetPipelineColumnSuffix() string {
	if o == nil || o.PipelineColumnSuffix == nil {
		var ret string
		return ret
	}
	return *o.PipelineColumnSuffix
}

// GetPipelineColumnSuffixOk returns a tuple with the PipelineColumnSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExposureSQLModelV2DTODataAttributesItems) GetPipelineColumnSuffixOk() (*string, bool) {
	if o == nil || o.PipelineColumnSuffix == nil {
		return nil, false
	}
	return o.PipelineColumnSuffix, true
}

// HasPipelineColumnSuffix returns a boolean if a field has been set.
func (o *ExperimentsExposureSQLModelV2DTODataAttributesItems) HasPipelineColumnSuffix() bool {
	return o != nil && o.PipelineColumnSuffix != nil
}

// SetPipelineColumnSuffix gets a reference to the given string and assigns it to the PipelineColumnSuffix field.
func (o *ExperimentsExposureSQLModelV2DTODataAttributesItems) SetPipelineColumnSuffix(v string) {
	o.PipelineColumnSuffix = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsExposureSQLModelV2DTODataAttributesItems) MarshalJSON() ([]byte, error) {
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
	if o.PipelineColumnSuffix != nil {
		toSerialize["pipeline_column_suffix"] = o.PipelineColumnSuffix
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsExposureSQLModelV2DTODataAttributesItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ColumnName           *string     `json:"column_name,omitempty"`
		ColumnType           *string     `json:"column_type,omitempty"`
		Description          *string     `json:"description,omitempty"`
		Id                   *string     `json:"id,omitempty"`
		MigrationMetadata    interface{} `json:"migration_metadata,omitempty"`
		Name                 *string     `json:"name,omitempty"`
		PipelineColumnSuffix *string     `json:"pipeline_column_suffix,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"column_name", "column_type", "description", "id", "migration_metadata", "name", "pipeline_column_suffix"})
	} else {
		return err
	}
	o.ColumnName = all.ColumnName
	o.ColumnType = all.ColumnType
	o.Description = all.Description
	o.Id = all.Id
	o.MigrationMetadata = all.MigrationMetadata
	o.Name = all.Name
	o.PipelineColumnSuffix = all.PipelineColumnSuffix

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
