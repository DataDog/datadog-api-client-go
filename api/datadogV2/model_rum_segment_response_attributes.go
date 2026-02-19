// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumSegmentResponseAttributes Attributes of a segment in a response.
type RumSegmentResponseAttributes struct {
	// The creation timestamp in RFC 3339 format.
	CreatedAt time.Time `json:"created_at"`
	// A user who performed an action on a segment.
	CreatedBy RumSegmentUser `json:"created_by"`
	// Query definition for the segment. Contains one or more query blocks and an optional combination formula.
	DataQuery RumSegmentDataQuery `json:"data_query"`
	// A description of the segment.
	Description string `json:"description"`
	// The last modification timestamp in RFC 3339 format.
	ModifiedAt time.Time `json:"modified_at"`
	// A user who performed an action on a segment.
	ModifiedBy RumSegmentUser `json:"modified_by"`
	// The name of the segment.
	Name string `json:"name"`
	// The organization identifier.
	OrgId int64 `json:"org_id"`
	// The number of users in the segment.
	RowCount int64 `json:"row_count"`
	// The source of a segment.
	Source RumSegmentSource `json:"source"`
	// A list of tags for the segment.
	Tags []string `json:"tags"`
	// The type of a segment based on its data query configuration.
	Type RumSegmentSegmentType `json:"type"`
	// The unique identifier of the segment.
	Uuid string `json:"uuid"`
	// The version number of the segment.
	Version int64 `json:"version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumSegmentResponseAttributes instantiates a new RumSegmentResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumSegmentResponseAttributes(createdAt time.Time, createdBy RumSegmentUser, dataQuery RumSegmentDataQuery, description string, modifiedAt time.Time, modifiedBy RumSegmentUser, name string, orgId int64, rowCount int64, source RumSegmentSource, tags []string, typeVar RumSegmentSegmentType, uuid string, version int64) *RumSegmentResponseAttributes {
	this := RumSegmentResponseAttributes{}
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.DataQuery = dataQuery
	this.Description = description
	this.ModifiedAt = modifiedAt
	this.ModifiedBy = modifiedBy
	this.Name = name
	this.OrgId = orgId
	this.RowCount = rowCount
	this.Source = source
	this.Tags = tags
	this.Type = typeVar
	this.Uuid = uuid
	this.Version = version
	return &this
}

// NewRumSegmentResponseAttributesWithDefaults instantiates a new RumSegmentResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumSegmentResponseAttributesWithDefaults() *RumSegmentResponseAttributes {
	this := RumSegmentResponseAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *RumSegmentResponseAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *RumSegmentResponseAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *RumSegmentResponseAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value.
func (o *RumSegmentResponseAttributes) GetCreatedBy() RumSegmentUser {
	if o == nil {
		var ret RumSegmentUser
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *RumSegmentResponseAttributes) GetCreatedByOk() (*RumSegmentUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value.
func (o *RumSegmentResponseAttributes) SetCreatedBy(v RumSegmentUser) {
	o.CreatedBy = v
}

// GetDataQuery returns the DataQuery field value.
func (o *RumSegmentResponseAttributes) GetDataQuery() RumSegmentDataQuery {
	if o == nil {
		var ret RumSegmentDataQuery
		return ret
	}
	return o.DataQuery
}

// GetDataQueryOk returns a tuple with the DataQuery field value
// and a boolean to check if the value has been set.
func (o *RumSegmentResponseAttributes) GetDataQueryOk() (*RumSegmentDataQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataQuery, true
}

// SetDataQuery sets field value.
func (o *RumSegmentResponseAttributes) SetDataQuery(v RumSegmentDataQuery) {
	o.DataQuery = v
}

// GetDescription returns the Description field value.
func (o *RumSegmentResponseAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *RumSegmentResponseAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *RumSegmentResponseAttributes) SetDescription(v string) {
	o.Description = v
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *RumSegmentResponseAttributes) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *RumSegmentResponseAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *RumSegmentResponseAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetModifiedBy returns the ModifiedBy field value.
func (o *RumSegmentResponseAttributes) GetModifiedBy() RumSegmentUser {
	if o == nil {
		var ret RumSegmentUser
		return ret
	}
	return o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value
// and a boolean to check if the value has been set.
func (o *RumSegmentResponseAttributes) GetModifiedByOk() (*RumSegmentUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedBy, true
}

// SetModifiedBy sets field value.
func (o *RumSegmentResponseAttributes) SetModifiedBy(v RumSegmentUser) {
	o.ModifiedBy = v
}

// GetName returns the Name field value.
func (o *RumSegmentResponseAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RumSegmentResponseAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *RumSegmentResponseAttributes) SetName(v string) {
	o.Name = v
}

// GetOrgId returns the OrgId field value.
func (o *RumSegmentResponseAttributes) GetOrgId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *RumSegmentResponseAttributes) GetOrgIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value.
func (o *RumSegmentResponseAttributes) SetOrgId(v int64) {
	o.OrgId = v
}

// GetRowCount returns the RowCount field value.
func (o *RumSegmentResponseAttributes) GetRowCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.RowCount
}

// GetRowCountOk returns a tuple with the RowCount field value
// and a boolean to check if the value has been set.
func (o *RumSegmentResponseAttributes) GetRowCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RowCount, true
}

// SetRowCount sets field value.
func (o *RumSegmentResponseAttributes) SetRowCount(v int64) {
	o.RowCount = v
}

// GetSource returns the Source field value.
func (o *RumSegmentResponseAttributes) GetSource() RumSegmentSource {
	if o == nil {
		var ret RumSegmentSource
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *RumSegmentResponseAttributes) GetSourceOk() (*RumSegmentSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *RumSegmentResponseAttributes) SetSource(v RumSegmentSource) {
	o.Source = v
}

// GetTags returns the Tags field value.
func (o *RumSegmentResponseAttributes) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *RumSegmentResponseAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value.
func (o *RumSegmentResponseAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetType returns the Type field value.
func (o *RumSegmentResponseAttributes) GetType() RumSegmentSegmentType {
	if o == nil {
		var ret RumSegmentSegmentType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RumSegmentResponseAttributes) GetTypeOk() (*RumSegmentSegmentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *RumSegmentResponseAttributes) SetType(v RumSegmentSegmentType) {
	o.Type = v
}

// GetUuid returns the Uuid field value.
func (o *RumSegmentResponseAttributes) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *RumSegmentResponseAttributes) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value.
func (o *RumSegmentResponseAttributes) SetUuid(v string) {
	o.Uuid = v
}

// GetVersion returns the Version field value.
func (o *RumSegmentResponseAttributes) GetVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *RumSegmentResponseAttributes) GetVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *RumSegmentResponseAttributes) SetVersion(v int64) {
	o.Version = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumSegmentResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["created_by"] = o.CreatedBy
	toSerialize["data_query"] = o.DataQuery
	toSerialize["description"] = o.Description
	if o.ModifiedAt.Nanosecond() == 0 {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["modified_by"] = o.ModifiedBy
	toSerialize["name"] = o.Name
	toSerialize["org_id"] = o.OrgId
	toSerialize["row_count"] = o.RowCount
	toSerialize["source"] = o.Source
	toSerialize["tags"] = o.Tags
	toSerialize["type"] = o.Type
	toSerialize["uuid"] = o.Uuid
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumSegmentResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt   *time.Time             `json:"created_at"`
		CreatedBy   *RumSegmentUser        `json:"created_by"`
		DataQuery   *RumSegmentDataQuery   `json:"data_query"`
		Description *string                `json:"description"`
		ModifiedAt  *time.Time             `json:"modified_at"`
		ModifiedBy  *RumSegmentUser        `json:"modified_by"`
		Name        *string                `json:"name"`
		OrgId       *int64                 `json:"org_id"`
		RowCount    *int64                 `json:"row_count"`
		Source      *RumSegmentSource      `json:"source"`
		Tags        *[]string              `json:"tags"`
		Type        *RumSegmentSegmentType `json:"type"`
		Uuid        *string                `json:"uuid"`
		Version     *int64                 `json:"version"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.CreatedBy == nil {
		return fmt.Errorf("required field created_by missing")
	}
	if all.DataQuery == nil {
		return fmt.Errorf("required field data_query missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.ModifiedAt == nil {
		return fmt.Errorf("required field modified_at missing")
	}
	if all.ModifiedBy == nil {
		return fmt.Errorf("required field modified_by missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.OrgId == nil {
		return fmt.Errorf("required field org_id missing")
	}
	if all.RowCount == nil {
		return fmt.Errorf("required field row_count missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	if all.Tags == nil {
		return fmt.Errorf("required field tags missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Uuid == nil {
		return fmt.Errorf("required field uuid missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "created_by", "data_query", "description", "modified_at", "modified_by", "name", "org_id", "row_count", "source", "tags", "type", "uuid", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = *all.CreatedAt
	if all.CreatedBy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CreatedBy = *all.CreatedBy
	if all.DataQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DataQuery = *all.DataQuery
	o.Description = *all.Description
	o.ModifiedAt = *all.ModifiedAt
	if all.ModifiedBy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ModifiedBy = *all.ModifiedBy
	o.Name = *all.Name
	o.OrgId = *all.OrgId
	o.RowCount = *all.RowCount
	if !all.Source.IsValid() {
		hasInvalidField = true
	} else {
		o.Source = *all.Source
	}
	o.Tags = *all.Tags
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.Uuid = *all.Uuid
	o.Version = *all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
