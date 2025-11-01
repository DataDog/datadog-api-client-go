// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SegmentDataAttributes
type SegmentDataAttributes struct {
	//
	CreatedAt *time.Time `json:"created_at,omitempty"`
	//
	CreatedBy *SegmentDataSource `json:"created_by,omitempty"`
	//
	DataQuery SegmentDataAttributesDataQuery `json:"data_query"`
	//
	Description *string `json:"description,omitempty"`
	//
	DisabledAt *time.Time `json:"disabled_at,omitempty"`
	//
	DisabledBy *SegmentDataSource `json:"disabled_by,omitempty"`
	//
	MaterializationRowCount *int64 `json:"materialization_row_count,omitempty"`
	//
	MaterializedAt *string `json:"materialized_at,omitempty"`
	//
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	//
	ModifiedBy *SegmentDataSource `json:"modified_by,omitempty"`
	//
	Name string `json:"name"`
	//
	OrgId *int64 `json:"org_id,omitempty"`
	//
	Source *int64 `json:"source,omitempty"`
	//
	Tags []string `json:"tags,omitempty"`
	//
	Version *int64 `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSegmentDataAttributes instantiates a new SegmentDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSegmentDataAttributes(dataQuery SegmentDataAttributesDataQuery, name string) *SegmentDataAttributes {
	this := SegmentDataAttributes{}
	this.DataQuery = dataQuery
	this.Name = name
	return &this
}

// NewSegmentDataAttributesWithDefaults instantiates a new SegmentDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSegmentDataAttributesWithDefaults() *SegmentDataAttributes {
	this := SegmentDataAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SegmentDataAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentDataAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SegmentDataAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *SegmentDataAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *SegmentDataAttributes) GetCreatedBy() SegmentDataSource {
	if o == nil || o.CreatedBy == nil {
		var ret SegmentDataSource
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentDataAttributes) GetCreatedByOk() (*SegmentDataSource, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *SegmentDataAttributes) HasCreatedBy() bool {
	return o != nil && o.CreatedBy != nil
}

// SetCreatedBy gets a reference to the given SegmentDataSource and assigns it to the CreatedBy field.
func (o *SegmentDataAttributes) SetCreatedBy(v SegmentDataSource) {
	o.CreatedBy = &v
}

// GetDataQuery returns the DataQuery field value.
func (o *SegmentDataAttributes) GetDataQuery() SegmentDataAttributesDataQuery {
	if o == nil {
		var ret SegmentDataAttributesDataQuery
		return ret
	}
	return o.DataQuery
}

// GetDataQueryOk returns a tuple with the DataQuery field value
// and a boolean to check if the value has been set.
func (o *SegmentDataAttributes) GetDataQueryOk() (*SegmentDataAttributesDataQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataQuery, true
}

// SetDataQuery sets field value.
func (o *SegmentDataAttributes) SetDataQuery(v SegmentDataAttributesDataQuery) {
	o.DataQuery = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SegmentDataAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SegmentDataAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SegmentDataAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetDisabledAt returns the DisabledAt field value if set, zero value otherwise.
func (o *SegmentDataAttributes) GetDisabledAt() time.Time {
	if o == nil || o.DisabledAt == nil {
		var ret time.Time
		return ret
	}
	return *o.DisabledAt
}

// GetDisabledAtOk returns a tuple with the DisabledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentDataAttributes) GetDisabledAtOk() (*time.Time, bool) {
	if o == nil || o.DisabledAt == nil {
		return nil, false
	}
	return o.DisabledAt, true
}

// HasDisabledAt returns a boolean if a field has been set.
func (o *SegmentDataAttributes) HasDisabledAt() bool {
	return o != nil && o.DisabledAt != nil
}

// SetDisabledAt gets a reference to the given time.Time and assigns it to the DisabledAt field.
func (o *SegmentDataAttributes) SetDisabledAt(v time.Time) {
	o.DisabledAt = &v
}

// GetDisabledBy returns the DisabledBy field value if set, zero value otherwise.
func (o *SegmentDataAttributes) GetDisabledBy() SegmentDataSource {
	if o == nil || o.DisabledBy == nil {
		var ret SegmentDataSource
		return ret
	}
	return *o.DisabledBy
}

// GetDisabledByOk returns a tuple with the DisabledBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentDataAttributes) GetDisabledByOk() (*SegmentDataSource, bool) {
	if o == nil || o.DisabledBy == nil {
		return nil, false
	}
	return o.DisabledBy, true
}

// HasDisabledBy returns a boolean if a field has been set.
func (o *SegmentDataAttributes) HasDisabledBy() bool {
	return o != nil && o.DisabledBy != nil
}

// SetDisabledBy gets a reference to the given SegmentDataSource and assigns it to the DisabledBy field.
func (o *SegmentDataAttributes) SetDisabledBy(v SegmentDataSource) {
	o.DisabledBy = &v
}

// GetMaterializationRowCount returns the MaterializationRowCount field value if set, zero value otherwise.
func (o *SegmentDataAttributes) GetMaterializationRowCount() int64 {
	if o == nil || o.MaterializationRowCount == nil {
		var ret int64
		return ret
	}
	return *o.MaterializationRowCount
}

// GetMaterializationRowCountOk returns a tuple with the MaterializationRowCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentDataAttributes) GetMaterializationRowCountOk() (*int64, bool) {
	if o == nil || o.MaterializationRowCount == nil {
		return nil, false
	}
	return o.MaterializationRowCount, true
}

// HasMaterializationRowCount returns a boolean if a field has been set.
func (o *SegmentDataAttributes) HasMaterializationRowCount() bool {
	return o != nil && o.MaterializationRowCount != nil
}

// SetMaterializationRowCount gets a reference to the given int64 and assigns it to the MaterializationRowCount field.
func (o *SegmentDataAttributes) SetMaterializationRowCount(v int64) {
	o.MaterializationRowCount = &v
}

// GetMaterializedAt returns the MaterializedAt field value if set, zero value otherwise.
func (o *SegmentDataAttributes) GetMaterializedAt() string {
	if o == nil || o.MaterializedAt == nil {
		var ret string
		return ret
	}
	return *o.MaterializedAt
}

// GetMaterializedAtOk returns a tuple with the MaterializedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentDataAttributes) GetMaterializedAtOk() (*string, bool) {
	if o == nil || o.MaterializedAt == nil {
		return nil, false
	}
	return o.MaterializedAt, true
}

// HasMaterializedAt returns a boolean if a field has been set.
func (o *SegmentDataAttributes) HasMaterializedAt() bool {
	return o != nil && o.MaterializedAt != nil
}

// SetMaterializedAt gets a reference to the given string and assigns it to the MaterializedAt field.
func (o *SegmentDataAttributes) SetMaterializedAt(v string) {
	o.MaterializedAt = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *SegmentDataAttributes) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentDataAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *SegmentDataAttributes) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt != nil
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *SegmentDataAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *SegmentDataAttributes) GetModifiedBy() SegmentDataSource {
	if o == nil || o.ModifiedBy == nil {
		var ret SegmentDataSource
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentDataAttributes) GetModifiedByOk() (*SegmentDataSource, bool) {
	if o == nil || o.ModifiedBy == nil {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *SegmentDataAttributes) HasModifiedBy() bool {
	return o != nil && o.ModifiedBy != nil
}

// SetModifiedBy gets a reference to the given SegmentDataSource and assigns it to the ModifiedBy field.
func (o *SegmentDataAttributes) SetModifiedBy(v SegmentDataSource) {
	o.ModifiedBy = &v
}

// GetName returns the Name field value.
func (o *SegmentDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SegmentDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SegmentDataAttributes) SetName(v string) {
	o.Name = v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *SegmentDataAttributes) GetOrgId() int64 {
	if o == nil || o.OrgId == nil {
		var ret int64
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentDataAttributes) GetOrgIdOk() (*int64, bool) {
	if o == nil || o.OrgId == nil {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *SegmentDataAttributes) HasOrgId() bool {
	return o != nil && o.OrgId != nil
}

// SetOrgId gets a reference to the given int64 and assigns it to the OrgId field.
func (o *SegmentDataAttributes) SetOrgId(v int64) {
	o.OrgId = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *SegmentDataAttributes) GetSource() int64 {
	if o == nil || o.Source == nil {
		var ret int64
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentDataAttributes) GetSourceOk() (*int64, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *SegmentDataAttributes) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given int64 and assigns it to the Source field.
func (o *SegmentDataAttributes) SetSource(v int64) {
	o.Source = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SegmentDataAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentDataAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SegmentDataAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *SegmentDataAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SegmentDataAttributes) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentDataAttributes) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SegmentDataAttributes) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *SegmentDataAttributes) SetVersion(v int64) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SegmentDataAttributes) MarshalJSON() ([]byte, error) {
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
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	toSerialize["data_query"] = o.DataQuery
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisabledAt != nil {
		if o.DisabledAt.Nanosecond() == 0 {
			toSerialize["disabled_at"] = o.DisabledAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["disabled_at"] = o.DisabledAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.DisabledBy != nil {
		toSerialize["disabled_by"] = o.DisabledBy
	}
	if o.MaterializationRowCount != nil {
		toSerialize["materialization_row_count"] = o.MaterializationRowCount
	}
	if o.MaterializedAt != nil {
		toSerialize["materialized_at"] = o.MaterializedAt
	}
	if o.ModifiedAt != nil {
		if o.ModifiedAt.Nanosecond() == 0 {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.ModifiedBy != nil {
		toSerialize["modified_by"] = o.ModifiedBy
	}
	toSerialize["name"] = o.Name
	if o.OrgId != nil {
		toSerialize["org_id"] = o.OrgId
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SegmentDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt               *time.Time                      `json:"created_at,omitempty"`
		CreatedBy               *SegmentDataSource              `json:"created_by,omitempty"`
		DataQuery               *SegmentDataAttributesDataQuery `json:"data_query"`
		Description             *string                         `json:"description,omitempty"`
		DisabledAt              *time.Time                      `json:"disabled_at,omitempty"`
		DisabledBy              *SegmentDataSource              `json:"disabled_by,omitempty"`
		MaterializationRowCount *int64                          `json:"materialization_row_count,omitempty"`
		MaterializedAt          *string                         `json:"materialized_at,omitempty"`
		ModifiedAt              *time.Time                      `json:"modified_at,omitempty"`
		ModifiedBy              *SegmentDataSource              `json:"modified_by,omitempty"`
		Name                    *string                         `json:"name"`
		OrgId                   *int64                          `json:"org_id,omitempty"`
		Source                  *int64                          `json:"source,omitempty"`
		Tags                    []string                        `json:"tags,omitempty"`
		Version                 *int64                          `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataQuery == nil {
		return fmt.Errorf("required field data_query missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "created_by", "data_query", "description", "disabled_at", "disabled_by", "materialization_row_count", "materialized_at", "modified_at", "modified_by", "name", "org_id", "source", "tags", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = all.CreatedAt
	if all.CreatedBy != nil && all.CreatedBy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CreatedBy = all.CreatedBy
	if all.DataQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DataQuery = *all.DataQuery
	o.Description = all.Description
	o.DisabledAt = all.DisabledAt
	if all.DisabledBy != nil && all.DisabledBy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DisabledBy = all.DisabledBy
	o.MaterializationRowCount = all.MaterializationRowCount
	o.MaterializedAt = all.MaterializedAt
	o.ModifiedAt = all.ModifiedAt
	if all.ModifiedBy != nil && all.ModifiedBy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ModifiedBy = all.ModifiedBy
	o.Name = *all.Name
	o.OrgId = all.OrgId
	o.Source = all.Source
	o.Tags = all.Tags
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
