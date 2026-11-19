// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SegmentSummaryAttributes Attributes of a segment in a list. The segment definition is not included.
type SegmentSummaryAttributes struct {
	// Time the segment was created.
	CreatedAt time.Time `json:"created_at"`
	// A Datadog user who created, modified, or disabled a segment.
	CreatedBy *SegmentUser `json:"created_by,omitempty"`
	// Description of the segment.
	Description *string `json:"description,omitempty"`
	// Kind of entity that the segment contains, `user` or `account`.
	Entity *string `json:"entity,omitempty"`
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
	// Kind of segment, derived from its blocks.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSegmentSummaryAttributes instantiates a new SegmentSummaryAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSegmentSummaryAttributes(createdAt time.Time, modifiedAt time.Time, name string) *SegmentSummaryAttributes {
	this := SegmentSummaryAttributes{}
	this.CreatedAt = createdAt
	this.ModifiedAt = modifiedAt
	this.Name = name
	return &this
}

// NewSegmentSummaryAttributesWithDefaults instantiates a new SegmentSummaryAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSegmentSummaryAttributesWithDefaults() *SegmentSummaryAttributes {
	this := SegmentSummaryAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *SegmentSummaryAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SegmentSummaryAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *SegmentSummaryAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *SegmentSummaryAttributes) GetCreatedBy() SegmentUser {
	if o == nil || o.CreatedBy == nil {
		var ret SegmentUser
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentSummaryAttributes) GetCreatedByOk() (*SegmentUser, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *SegmentSummaryAttributes) HasCreatedBy() bool {
	return o != nil && o.CreatedBy != nil
}

// SetCreatedBy gets a reference to the given SegmentUser and assigns it to the CreatedBy field.
func (o *SegmentSummaryAttributes) SetCreatedBy(v SegmentUser) {
	o.CreatedBy = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SegmentSummaryAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentSummaryAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SegmentSummaryAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SegmentSummaryAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetEntity returns the Entity field value if set, zero value otherwise.
func (o *SegmentSummaryAttributes) GetEntity() string {
	if o == nil || o.Entity == nil {
		var ret string
		return ret
	}
	return *o.Entity
}

// GetEntityOk returns a tuple with the Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentSummaryAttributes) GetEntityOk() (*string, bool) {
	if o == nil || o.Entity == nil {
		return nil, false
	}
	return o.Entity, true
}

// HasEntity returns a boolean if a field has been set.
func (o *SegmentSummaryAttributes) HasEntity() bool {
	return o != nil && o.Entity != nil
}

// SetEntity gets a reference to the given string and assigns it to the Entity field.
func (o *SegmentSummaryAttributes) SetEntity(v string) {
	o.Entity = &v
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *SegmentSummaryAttributes) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *SegmentSummaryAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *SegmentSummaryAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *SegmentSummaryAttributes) GetModifiedBy() SegmentUser {
	if o == nil || o.ModifiedBy == nil {
		var ret SegmentUser
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentSummaryAttributes) GetModifiedByOk() (*SegmentUser, bool) {
	if o == nil || o.ModifiedBy == nil {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *SegmentSummaryAttributes) HasModifiedBy() bool {
	return o != nil && o.ModifiedBy != nil
}

// SetModifiedBy gets a reference to the given SegmentUser and assigns it to the ModifiedBy field.
func (o *SegmentSummaryAttributes) SetModifiedBy(v SegmentUser) {
	o.ModifiedBy = &v
}

// GetName returns the Name field value.
func (o *SegmentSummaryAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SegmentSummaryAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SegmentSummaryAttributes) SetName(v string) {
	o.Name = v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *SegmentSummaryAttributes) GetOrgId() int64 {
	if o == nil || o.OrgId == nil {
		var ret int64
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentSummaryAttributes) GetOrgIdOk() (*int64, bool) {
	if o == nil || o.OrgId == nil {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *SegmentSummaryAttributes) HasOrgId() bool {
	return o != nil && o.OrgId != nil
}

// SetOrgId gets a reference to the given int64 and assigns it to the OrgId field.
func (o *SegmentSummaryAttributes) SetOrgId(v int64) {
	o.OrgId = &v
}

// GetRowCount returns the RowCount field value if set, zero value otherwise.
func (o *SegmentSummaryAttributes) GetRowCount() int64 {
	if o == nil || o.RowCount == nil {
		var ret int64
		return ret
	}
	return *o.RowCount
}

// GetRowCountOk returns a tuple with the RowCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentSummaryAttributes) GetRowCountOk() (*int64, bool) {
	if o == nil || o.RowCount == nil {
		return nil, false
	}
	return o.RowCount, true
}

// HasRowCount returns a boolean if a field has been set.
func (o *SegmentSummaryAttributes) HasRowCount() bool {
	return o != nil && o.RowCount != nil
}

// SetRowCount gets a reference to the given int64 and assigns it to the RowCount field.
func (o *SegmentSummaryAttributes) SetRowCount(v int64) {
	o.RowCount = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *SegmentSummaryAttributes) GetSource() int64 {
	if o == nil || o.Source == nil {
		var ret int64
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentSummaryAttributes) GetSourceOk() (*int64, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *SegmentSummaryAttributes) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given int64 and assigns it to the Source field.
func (o *SegmentSummaryAttributes) SetSource(v int64) {
	o.Source = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SegmentSummaryAttributes) GetTags() []string {
	if o == nil || o.Tags.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Tags.Get()
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SegmentSummaryAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags.Get(), o.Tags.IsSet()
}

// HasTags returns a boolean if a field has been set.
func (o *SegmentSummaryAttributes) HasTags() bool {
	return o != nil && o.Tags.IsSet()
}

// SetTags gets a reference to the given datadog.NullableList[string] and assigns it to the Tags field.
func (o *SegmentSummaryAttributes) SetTags(v []string) {
	o.Tags.Set(&v)
}

// SetTagsNil sets the value for Tags to be an explicit nil.
func (o *SegmentSummaryAttributes) SetTagsNil() {
	o.Tags.Set(nil)
}

// UnsetTags ensures that no value is present for Tags, not even an explicit nil.
func (o *SegmentSummaryAttributes) UnsetTags() {
	o.Tags.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SegmentSummaryAttributes) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentSummaryAttributes) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SegmentSummaryAttributes) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SegmentSummaryAttributes) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SegmentSummaryAttributes) MarshalJSON() ([]byte, error) {
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
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Entity != nil {
		toSerialize["entity"] = o.Entity
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SegmentSummaryAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt   *time.Time                   `json:"created_at"`
		CreatedBy   *SegmentUser                 `json:"created_by,omitempty"`
		Description *string                      `json:"description,omitempty"`
		Entity      *string                      `json:"entity,omitempty"`
		ModifiedAt  *time.Time                   `json:"modified_at"`
		ModifiedBy  *SegmentUser                 `json:"modified_by,omitempty"`
		Name        *string                      `json:"name"`
		OrgId       *int64                       `json:"org_id,omitempty"`
		RowCount    *int64                       `json:"row_count,omitempty"`
		Source      *int64                       `json:"source,omitempty"`
		Tags        datadog.NullableList[string] `json:"tags,omitempty"`
		Type        *string                      `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.ModifiedAt == nil {
		return fmt.Errorf("required field modified_at missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "created_by", "description", "entity", "modified_at", "modified_by", "name", "org_id", "row_count", "source", "tags", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = *all.CreatedAt
	if all.CreatedBy != nil && all.CreatedBy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CreatedBy = all.CreatedBy
	o.Description = all.Description
	o.Entity = all.Entity
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

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
