// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SegmentAttributes Attributes of a segment.
type SegmentAttributes struct {
	// Time the segment was created.
	CreatedAt time.Time `json:"created_at"`
	// A Datadog user who created, modified, or disabled a segment.
	CreatedBy *SegmentUser `json:"created_by,omitempty"`
	// Defines who belongs to the segment. It contains one or more blocks, up to 20 in total.
	// Use `combination` to combine named blocks with a boolean formula such as `A AND NOT B`.
	DataQuery SegmentDataQuery `json:"data_query"`
	// Description of the segment.
	Description *string `json:"description,omitempty"`
	// Time the segment was disabled, if it was.
	DisabledAt *time.Time `json:"disabled_at,omitempty"`
	// A Datadog user who created, modified, or disabled a segment.
	DisabledBy *SegmentUser `json:"disabled_by,omitempty"`
	// Kind of entity that the segment contains, `user` or `account`.
	Entity *string `json:"entity,omitempty"`
	// Time the segment membership was last computed.
	MaterializedAt datadog.NullableTime `json:"materialized_at,omitempty"`
	// Time the segment was last modified.
	ModifiedAt time.Time `json:"modified_at"`
	// A Datadog user who created, modified, or disabled a segment.
	ModifiedBy *SegmentUser `json:"modified_by,omitempty"`
	// Name of the segment.
	Name string `json:"name"`
	// ID of the organization that owns the segment.
	OrgId *int64 `json:"org_id,omitempty"`
	// Number of entities in the segment when its membership was last computed. `-1` means the count is not
	// available yet.
	RowCount *int64 `json:"row_count,omitempty"`
	// Who created the segment. `0` for a user and `1` for Datadog.
	Source *int64 `json:"source,omitempty"`
	// Tags of the segment.
	Tags datadog.NullableList[string] `json:"tags,omitempty"`
	// Kind of segment, derived from its blocks, such as `event_platform`, `user_store`, or `combination`.
	Type *string `json:"type,omitempty"`
	// Version of the segment, increased on each update.
	Version *int64 `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSegmentAttributes instantiates a new SegmentAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSegmentAttributes(createdAt time.Time, dataQuery SegmentDataQuery, modifiedAt time.Time, name string) *SegmentAttributes {
	this := SegmentAttributes{}
	this.CreatedAt = createdAt
	this.DataQuery = dataQuery
	this.ModifiedAt = modifiedAt
	this.Name = name
	return &this
}

// NewSegmentAttributesWithDefaults instantiates a new SegmentAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSegmentAttributesWithDefaults() *SegmentAttributes {
	this := SegmentAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *SegmentAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SegmentAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *SegmentAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *SegmentAttributes) GetCreatedBy() SegmentUser {
	if o == nil || o.CreatedBy == nil {
		var ret SegmentUser
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentAttributes) GetCreatedByOk() (*SegmentUser, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *SegmentAttributes) HasCreatedBy() bool {
	return o != nil && o.CreatedBy != nil
}

// SetCreatedBy gets a reference to the given SegmentUser and assigns it to the CreatedBy field.
func (o *SegmentAttributes) SetCreatedBy(v SegmentUser) {
	o.CreatedBy = &v
}

// GetDataQuery returns the DataQuery field value.
func (o *SegmentAttributes) GetDataQuery() SegmentDataQuery {
	if o == nil {
		var ret SegmentDataQuery
		return ret
	}
	return o.DataQuery
}

// GetDataQueryOk returns a tuple with the DataQuery field value
// and a boolean to check if the value has been set.
func (o *SegmentAttributes) GetDataQueryOk() (*SegmentDataQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataQuery, true
}

// SetDataQuery sets field value.
func (o *SegmentAttributes) SetDataQuery(v SegmentDataQuery) {
	o.DataQuery = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SegmentAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SegmentAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SegmentAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetDisabledAt returns the DisabledAt field value if set, zero value otherwise.
func (o *SegmentAttributes) GetDisabledAt() time.Time {
	if o == nil || o.DisabledAt == nil {
		var ret time.Time
		return ret
	}
	return *o.DisabledAt
}

// GetDisabledAtOk returns a tuple with the DisabledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentAttributes) GetDisabledAtOk() (*time.Time, bool) {
	if o == nil || o.DisabledAt == nil {
		return nil, false
	}
	return o.DisabledAt, true
}

// HasDisabledAt returns a boolean if a field has been set.
func (o *SegmentAttributes) HasDisabledAt() bool {
	return o != nil && o.DisabledAt != nil
}

// SetDisabledAt gets a reference to the given time.Time and assigns it to the DisabledAt field.
func (o *SegmentAttributes) SetDisabledAt(v time.Time) {
	o.DisabledAt = &v
}

// GetDisabledBy returns the DisabledBy field value if set, zero value otherwise.
func (o *SegmentAttributes) GetDisabledBy() SegmentUser {
	if o == nil || o.DisabledBy == nil {
		var ret SegmentUser
		return ret
	}
	return *o.DisabledBy
}

// GetDisabledByOk returns a tuple with the DisabledBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentAttributes) GetDisabledByOk() (*SegmentUser, bool) {
	if o == nil || o.DisabledBy == nil {
		return nil, false
	}
	return o.DisabledBy, true
}

// HasDisabledBy returns a boolean if a field has been set.
func (o *SegmentAttributes) HasDisabledBy() bool {
	return o != nil && o.DisabledBy != nil
}

// SetDisabledBy gets a reference to the given SegmentUser and assigns it to the DisabledBy field.
func (o *SegmentAttributes) SetDisabledBy(v SegmentUser) {
	o.DisabledBy = &v
}

// GetEntity returns the Entity field value if set, zero value otherwise.
func (o *SegmentAttributes) GetEntity() string {
	if o == nil || o.Entity == nil {
		var ret string
		return ret
	}
	return *o.Entity
}

// GetEntityOk returns a tuple with the Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentAttributes) GetEntityOk() (*string, bool) {
	if o == nil || o.Entity == nil {
		return nil, false
	}
	return o.Entity, true
}

// HasEntity returns a boolean if a field has been set.
func (o *SegmentAttributes) HasEntity() bool {
	return o != nil && o.Entity != nil
}

// SetEntity gets a reference to the given string and assigns it to the Entity field.
func (o *SegmentAttributes) SetEntity(v string) {
	o.Entity = &v
}

// GetMaterializedAt returns the MaterializedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SegmentAttributes) GetMaterializedAt() time.Time {
	if o == nil || o.MaterializedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.MaterializedAt.Get()
}

// GetMaterializedAtOk returns a tuple with the MaterializedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SegmentAttributes) GetMaterializedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaterializedAt.Get(), o.MaterializedAt.IsSet()
}

// HasMaterializedAt returns a boolean if a field has been set.
func (o *SegmentAttributes) HasMaterializedAt() bool {
	return o != nil && o.MaterializedAt.IsSet()
}

// SetMaterializedAt gets a reference to the given datadog.NullableTime and assigns it to the MaterializedAt field.
func (o *SegmentAttributes) SetMaterializedAt(v time.Time) {
	o.MaterializedAt.Set(&v)
}

// SetMaterializedAtNil sets the value for MaterializedAt to be an explicit nil.
func (o *SegmentAttributes) SetMaterializedAtNil() {
	o.MaterializedAt.Set(nil)
}

// UnsetMaterializedAt ensures that no value is present for MaterializedAt, not even an explicit nil.
func (o *SegmentAttributes) UnsetMaterializedAt() {
	o.MaterializedAt.Unset()
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *SegmentAttributes) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *SegmentAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *SegmentAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *SegmentAttributes) GetModifiedBy() SegmentUser {
	if o == nil || o.ModifiedBy == nil {
		var ret SegmentUser
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentAttributes) GetModifiedByOk() (*SegmentUser, bool) {
	if o == nil || o.ModifiedBy == nil {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *SegmentAttributes) HasModifiedBy() bool {
	return o != nil && o.ModifiedBy != nil
}

// SetModifiedBy gets a reference to the given SegmentUser and assigns it to the ModifiedBy field.
func (o *SegmentAttributes) SetModifiedBy(v SegmentUser) {
	o.ModifiedBy = &v
}

// GetName returns the Name field value.
func (o *SegmentAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SegmentAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SegmentAttributes) SetName(v string) {
	o.Name = v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *SegmentAttributes) GetOrgId() int64 {
	if o == nil || o.OrgId == nil {
		var ret int64
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentAttributes) GetOrgIdOk() (*int64, bool) {
	if o == nil || o.OrgId == nil {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *SegmentAttributes) HasOrgId() bool {
	return o != nil && o.OrgId != nil
}

// SetOrgId gets a reference to the given int64 and assigns it to the OrgId field.
func (o *SegmentAttributes) SetOrgId(v int64) {
	o.OrgId = &v
}

// GetRowCount returns the RowCount field value if set, zero value otherwise.
func (o *SegmentAttributes) GetRowCount() int64 {
	if o == nil || o.RowCount == nil {
		var ret int64
		return ret
	}
	return *o.RowCount
}

// GetRowCountOk returns a tuple with the RowCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentAttributes) GetRowCountOk() (*int64, bool) {
	if o == nil || o.RowCount == nil {
		return nil, false
	}
	return o.RowCount, true
}

// HasRowCount returns a boolean if a field has been set.
func (o *SegmentAttributes) HasRowCount() bool {
	return o != nil && o.RowCount != nil
}

// SetRowCount gets a reference to the given int64 and assigns it to the RowCount field.
func (o *SegmentAttributes) SetRowCount(v int64) {
	o.RowCount = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *SegmentAttributes) GetSource() int64 {
	if o == nil || o.Source == nil {
		var ret int64
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentAttributes) GetSourceOk() (*int64, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *SegmentAttributes) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given int64 and assigns it to the Source field.
func (o *SegmentAttributes) SetSource(v int64) {
	o.Source = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SegmentAttributes) GetTags() []string {
	if o == nil || o.Tags.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Tags.Get()
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SegmentAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags.Get(), o.Tags.IsSet()
}

// HasTags returns a boolean if a field has been set.
func (o *SegmentAttributes) HasTags() bool {
	return o != nil && o.Tags.IsSet()
}

// SetTags gets a reference to the given datadog.NullableList[string] and assigns it to the Tags field.
func (o *SegmentAttributes) SetTags(v []string) {
	o.Tags.Set(&v)
}

// SetTagsNil sets the value for Tags to be an explicit nil.
func (o *SegmentAttributes) SetTagsNil() {
	o.Tags.Set(nil)
}

// UnsetTags ensures that no value is present for Tags, not even an explicit nil.
func (o *SegmentAttributes) UnsetTags() {
	o.Tags.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SegmentAttributes) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentAttributes) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SegmentAttributes) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SegmentAttributes) SetType(v string) {
	o.Type = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SegmentAttributes) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentAttributes) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SegmentAttributes) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *SegmentAttributes) SetVersion(v int64) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SegmentAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
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
	if o.Entity != nil {
		toSerialize["entity"] = o.Entity
	}
	if o.MaterializedAt.IsSet() {
		toSerialize["materialized_at"] = o.MaterializedAt.Get()
	}
	if o.ModifiedAt.Nanosecond() == 0 {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.ModifiedBy != nil {
		toSerialize["modified_by"] = o.ModifiedBy
	}
	toSerialize["name"] = o.Name
	if o.OrgId != nil {
		toSerialize["org_id"] = o.OrgId
	}
	if o.RowCount != nil {
		toSerialize["row_count"] = o.RowCount
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Tags.IsSet() {
		toSerialize["tags"] = o.Tags.Get()
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
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
func (o *SegmentAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt      *time.Time                   `json:"created_at"`
		CreatedBy      *SegmentUser                 `json:"created_by,omitempty"`
		DataQuery      *SegmentDataQuery            `json:"data_query"`
		Description    *string                      `json:"description,omitempty"`
		DisabledAt     *time.Time                   `json:"disabled_at,omitempty"`
		DisabledBy     *SegmentUser                 `json:"disabled_by,omitempty"`
		Entity         *string                      `json:"entity,omitempty"`
		MaterializedAt datadog.NullableTime         `json:"materialized_at,omitempty"`
		ModifiedAt     *time.Time                   `json:"modified_at"`
		ModifiedBy     *SegmentUser                 `json:"modified_by,omitempty"`
		Name           *string                      `json:"name"`
		OrgId          *int64                       `json:"org_id,omitempty"`
		RowCount       *int64                       `json:"row_count,omitempty"`
		Source         *int64                       `json:"source,omitempty"`
		Tags           datadog.NullableList[string] `json:"tags,omitempty"`
		Type           *string                      `json:"type,omitempty"`
		Version        *int64                       `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.DataQuery == nil {
		return fmt.Errorf("required field data_query missing")
	}
	if all.ModifiedAt == nil {
		return fmt.Errorf("required field modified_at missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "created_by", "data_query", "description", "disabled_at", "disabled_by", "entity", "materialized_at", "modified_at", "modified_by", "name", "org_id", "row_count", "source", "tags", "type", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = *all.CreatedAt
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
	o.Entity = all.Entity
	o.MaterializedAt = all.MaterializedAt
	o.ModifiedAt = *all.ModifiedAt
	if all.ModifiedBy != nil && all.ModifiedBy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ModifiedBy = all.ModifiedBy
	o.Name = *all.Name
	o.OrgId = all.OrgId
	o.RowCount = all.RowCount
	o.Source = all.Source
	o.Tags = all.Tags
	o.Type = all.Type
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
