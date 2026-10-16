// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringBatchMigrationAttributes Attributes of a batch migration.
type SecurityMonitoringBatchMigrationAttributes struct {
	// The creation date of the batch migration, in RFC 3339 format.
	CreatedAt time.Time `json:"createdAt"`
	// The UUID of the user or service account who created the batch migration.
	CreatedBy *string `json:"createdBy,omitempty"`
	// The description of the batch migration.
	Description *string `json:"description,omitempty"`
	// The origin of a batch migration. Always `batch_migration` for batch migrations
	// created through this API.
	InputSource SecurityMonitoringBatchMigrationInputSource `json:"inputSource"`
	// The last update date of the batch migration, in RFC 3339 format.
	LastUpdatedAt time.Time `json:"lastUpdatedAt"`
	// The UUID of the user or service account who last updated the batch migration,
	// or the name of a system actor.
	LastUpdatedBy *string `json:"lastUpdatedBy,omitempty"`
	// The name of the batch migration.
	Name string `json:"name"`
	// The ID of the organization that owns the batch migration.
	OrgId int64 `json:"orgId"`
	// The number of resources of a batch migration in each status.
	ResourceStatusCounts SecurityMonitoringBatchMigrationResourceStatusCounts `json:"resourceStatusCounts"`
	// The status of a batch migration. A batch migration is created as `draft` and moves to
	// `in_progress` when started. It moves to `completed` once no resource is `pending` or
	// `in_progress`, or to `canceled` when canceled.
	Status SecurityMonitoringBatchMigrationStatus `json:"status"`
	// Tags of the batch migration. Up to 50 tags, each up to 256 characters.
	Tags []string `json:"tags"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringBatchMigrationAttributes instantiates a new SecurityMonitoringBatchMigrationAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringBatchMigrationAttributes(createdAt time.Time, inputSource SecurityMonitoringBatchMigrationInputSource, lastUpdatedAt time.Time, name string, orgId int64, resourceStatusCounts SecurityMonitoringBatchMigrationResourceStatusCounts, status SecurityMonitoringBatchMigrationStatus, tags []string) *SecurityMonitoringBatchMigrationAttributes {
	this := SecurityMonitoringBatchMigrationAttributes{}
	this.CreatedAt = createdAt
	this.InputSource = inputSource
	this.LastUpdatedAt = lastUpdatedAt
	this.Name = name
	this.OrgId = orgId
	this.ResourceStatusCounts = resourceStatusCounts
	this.Status = status
	this.Tags = tags
	return &this
}

// NewSecurityMonitoringBatchMigrationAttributesWithDefaults instantiates a new SecurityMonitoringBatchMigrationAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringBatchMigrationAttributesWithDefaults() *SecurityMonitoringBatchMigrationAttributes {
	this := SecurityMonitoringBatchMigrationAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *SecurityMonitoringBatchMigrationAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *SecurityMonitoringBatchMigrationAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *SecurityMonitoringBatchMigrationAttributes) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationAttributes) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *SecurityMonitoringBatchMigrationAttributes) HasCreatedBy() bool {
	return o != nil && o.CreatedBy != nil
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *SecurityMonitoringBatchMigrationAttributes) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SecurityMonitoringBatchMigrationAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SecurityMonitoringBatchMigrationAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SecurityMonitoringBatchMigrationAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetInputSource returns the InputSource field value.
func (o *SecurityMonitoringBatchMigrationAttributes) GetInputSource() SecurityMonitoringBatchMigrationInputSource {
	if o == nil {
		var ret SecurityMonitoringBatchMigrationInputSource
		return ret
	}
	return o.InputSource
}

// GetInputSourceOk returns a tuple with the InputSource field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationAttributes) GetInputSourceOk() (*SecurityMonitoringBatchMigrationInputSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InputSource, true
}

// SetInputSource sets field value.
func (o *SecurityMonitoringBatchMigrationAttributes) SetInputSource(v SecurityMonitoringBatchMigrationInputSource) {
	o.InputSource = v
}

// GetLastUpdatedAt returns the LastUpdatedAt field value.
func (o *SecurityMonitoringBatchMigrationAttributes) GetLastUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.LastUpdatedAt
}

// GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationAttributes) GetLastUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedAt, true
}

// SetLastUpdatedAt sets field value.
func (o *SecurityMonitoringBatchMigrationAttributes) SetLastUpdatedAt(v time.Time) {
	o.LastUpdatedAt = v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value if set, zero value otherwise.
func (o *SecurityMonitoringBatchMigrationAttributes) GetLastUpdatedBy() string {
	if o == nil || o.LastUpdatedBy == nil {
		var ret string
		return ret
	}
	return *o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationAttributes) GetLastUpdatedByOk() (*string, bool) {
	if o == nil || o.LastUpdatedBy == nil {
		return nil, false
	}
	return o.LastUpdatedBy, true
}

// HasLastUpdatedBy returns a boolean if a field has been set.
func (o *SecurityMonitoringBatchMigrationAttributes) HasLastUpdatedBy() bool {
	return o != nil && o.LastUpdatedBy != nil
}

// SetLastUpdatedBy gets a reference to the given string and assigns it to the LastUpdatedBy field.
func (o *SecurityMonitoringBatchMigrationAttributes) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = &v
}

// GetName returns the Name field value.
func (o *SecurityMonitoringBatchMigrationAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SecurityMonitoringBatchMigrationAttributes) SetName(v string) {
	o.Name = v
}

// GetOrgId returns the OrgId field value.
func (o *SecurityMonitoringBatchMigrationAttributes) GetOrgId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationAttributes) GetOrgIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value.
func (o *SecurityMonitoringBatchMigrationAttributes) SetOrgId(v int64) {
	o.OrgId = v
}

// GetResourceStatusCounts returns the ResourceStatusCounts field value.
func (o *SecurityMonitoringBatchMigrationAttributes) GetResourceStatusCounts() SecurityMonitoringBatchMigrationResourceStatusCounts {
	if o == nil {
		var ret SecurityMonitoringBatchMigrationResourceStatusCounts
		return ret
	}
	return o.ResourceStatusCounts
}

// GetResourceStatusCountsOk returns a tuple with the ResourceStatusCounts field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationAttributes) GetResourceStatusCountsOk() (*SecurityMonitoringBatchMigrationResourceStatusCounts, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceStatusCounts, true
}

// SetResourceStatusCounts sets field value.
func (o *SecurityMonitoringBatchMigrationAttributes) SetResourceStatusCounts(v SecurityMonitoringBatchMigrationResourceStatusCounts) {
	o.ResourceStatusCounts = v
}

// GetStatus returns the Status field value.
func (o *SecurityMonitoringBatchMigrationAttributes) GetStatus() SecurityMonitoringBatchMigrationStatus {
	if o == nil {
		var ret SecurityMonitoringBatchMigrationStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationAttributes) GetStatusOk() (*SecurityMonitoringBatchMigrationStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *SecurityMonitoringBatchMigrationAttributes) SetStatus(v SecurityMonitoringBatchMigrationStatus) {
	o.Status = v
}

// GetTags returns the Tags field value.
func (o *SecurityMonitoringBatchMigrationAttributes) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value.
func (o *SecurityMonitoringBatchMigrationAttributes) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringBatchMigrationAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	toSerialize["inputSource"] = o.InputSource
	if o.LastUpdatedAt.Nanosecond() == 0 {
		toSerialize["lastUpdatedAt"] = o.LastUpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["lastUpdatedAt"] = o.LastUpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.LastUpdatedBy != nil {
		toSerialize["lastUpdatedBy"] = o.LastUpdatedBy
	}
	toSerialize["name"] = o.Name
	toSerialize["orgId"] = o.OrgId
	toSerialize["resourceStatusCounts"] = o.ResourceStatusCounts
	toSerialize["status"] = o.Status
	toSerialize["tags"] = o.Tags

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringBatchMigrationAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt            *time.Time                                            `json:"createdAt"`
		CreatedBy            *string                                               `json:"createdBy,omitempty"`
		Description          *string                                               `json:"description,omitempty"`
		InputSource          *SecurityMonitoringBatchMigrationInputSource          `json:"inputSource"`
		LastUpdatedAt        *time.Time                                            `json:"lastUpdatedAt"`
		LastUpdatedBy        *string                                               `json:"lastUpdatedBy,omitempty"`
		Name                 *string                                               `json:"name"`
		OrgId                *int64                                                `json:"orgId"`
		ResourceStatusCounts *SecurityMonitoringBatchMigrationResourceStatusCounts `json:"resourceStatusCounts"`
		Status               *SecurityMonitoringBatchMigrationStatus               `json:"status"`
		Tags                 *[]string                                             `json:"tags"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	if all.InputSource == nil {
		return fmt.Errorf("required field inputSource missing")
	}
	if all.LastUpdatedAt == nil {
		return fmt.Errorf("required field lastUpdatedAt missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.OrgId == nil {
		return fmt.Errorf("required field orgId missing")
	}
	if all.ResourceStatusCounts == nil {
		return fmt.Errorf("required field resourceStatusCounts missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.Tags == nil {
		return fmt.Errorf("required field tags missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"createdAt", "createdBy", "description", "inputSource", "lastUpdatedAt", "lastUpdatedBy", "name", "orgId", "resourceStatusCounts", "status", "tags"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = *all.CreatedAt
	o.CreatedBy = all.CreatedBy
	o.Description = all.Description
	if !all.InputSource.IsValid() {
		hasInvalidField = true
	} else {
		o.InputSource = *all.InputSource
	}
	o.LastUpdatedAt = *all.LastUpdatedAt
	o.LastUpdatedBy = all.LastUpdatedBy
	o.Name = *all.Name
	o.OrgId = *all.OrgId
	if all.ResourceStatusCounts.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ResourceStatusCounts = *all.ResourceStatusCounts
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.Tags = *all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
