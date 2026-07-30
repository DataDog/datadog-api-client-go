// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMOperationResponseAttributes Attributes of a RUM operation response.
type RUMOperationResponseAttributes struct {
	// The RUM application ID the operation belongs to.
	ApplicationId datadog.NullableUUID `json:"application_id,omitempty"`
	// The category of the RUM operation.
	Category datadog.NullableString `json:"category,omitempty"`
	// The timestamp when the RUM operation was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// A Datadog user referenced by a RUM operation.
	CreatedBy *RUMOperationUser `json:"created_by,omitempty"`
	// A description of the RUM operation.
	Description datadog.NullableString `json:"description,omitempty"`
	// A human-readable display name for the RUM operation.
	DisplayName *string `json:"display_name,omitempty"`
	// The list of feature IDs associated with the RUM operation.
	FeatureIds []string `json:"feature_ids,omitempty"`
	// The definition of a RUM operation's journey, used to detect it from RUM events.
	JourneyRum RUMOperationJourneyRum `json:"journey_rum"`
	// The unique name of the RUM operation. Must not contain spaces.
	Name string `json:"name"`
	// The ID of the organization the RUM operation belongs to.
	OrgId *int64 `json:"org_id,omitempty"`
	// A list of tags associated with the RUM operation.
	Tags []string `json:"tags"`
	// The timestamp when the RUM operation was last updated.
	UpdatedAt datadog.NullableTime `json:"updated_at,omitempty"`
	// A Datadog user referenced by a RUM operation.
	UpdatedBy *RUMOperationUser `json:"updated_by,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRUMOperationResponseAttributes instantiates a new RUMOperationResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRUMOperationResponseAttributes(journeyRum RUMOperationJourneyRum, name string, tags []string) *RUMOperationResponseAttributes {
	this := RUMOperationResponseAttributes{}
	this.JourneyRum = journeyRum
	this.Name = name
	this.Tags = tags
	return &this
}

// NewRUMOperationResponseAttributesWithDefaults instantiates a new RUMOperationResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRUMOperationResponseAttributesWithDefaults() *RUMOperationResponseAttributes {
	this := RUMOperationResponseAttributes{}
	return &this
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RUMOperationResponseAttributes) GetApplicationId() uuid.UUID {
	if o == nil || o.ApplicationId.Get() == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.ApplicationId.Get()
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *RUMOperationResponseAttributes) GetApplicationIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApplicationId.Get(), o.ApplicationId.IsSet()
}

// HasApplicationId returns a boolean if a field has been set.
func (o *RUMOperationResponseAttributes) HasApplicationId() bool {
	return o != nil && o.ApplicationId.IsSet()
}

// SetApplicationId gets a reference to the given datadog.NullableUUID and assigns it to the ApplicationId field.
func (o *RUMOperationResponseAttributes) SetApplicationId(v uuid.UUID) {
	o.ApplicationId.Set(&v)
}

// SetApplicationIdNil sets the value for ApplicationId to be an explicit nil.
func (o *RUMOperationResponseAttributes) SetApplicationIdNil() {
	o.ApplicationId.Set(nil)
}

// UnsetApplicationId ensures that no value is present for ApplicationId, not even an explicit nil.
func (o *RUMOperationResponseAttributes) UnsetApplicationId() {
	o.ApplicationId.Unset()
}

// GetCategory returns the Category field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RUMOperationResponseAttributes) GetCategory() string {
	if o == nil || o.Category.Get() == nil {
		var ret string
		return ret
	}
	return *o.Category.Get()
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *RUMOperationResponseAttributes) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Category.Get(), o.Category.IsSet()
}

// HasCategory returns a boolean if a field has been set.
func (o *RUMOperationResponseAttributes) HasCategory() bool {
	return o != nil && o.Category.IsSet()
}

// SetCategory gets a reference to the given datadog.NullableString and assigns it to the Category field.
func (o *RUMOperationResponseAttributes) SetCategory(v string) {
	o.Category.Set(&v)
}

// SetCategoryNil sets the value for Category to be an explicit nil.
func (o *RUMOperationResponseAttributes) SetCategoryNil() {
	o.Category.Set(nil)
}

// UnsetCategory ensures that no value is present for Category, not even an explicit nil.
func (o *RUMOperationResponseAttributes) UnsetCategory() {
	o.Category.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RUMOperationResponseAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMOperationResponseAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RUMOperationResponseAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *RUMOperationResponseAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *RUMOperationResponseAttributes) GetCreatedBy() RUMOperationUser {
	if o == nil || o.CreatedBy == nil {
		var ret RUMOperationUser
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMOperationResponseAttributes) GetCreatedByOk() (*RUMOperationUser, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *RUMOperationResponseAttributes) HasCreatedBy() bool {
	return o != nil && o.CreatedBy != nil
}

// SetCreatedBy gets a reference to the given RUMOperationUser and assigns it to the CreatedBy field.
func (o *RUMOperationResponseAttributes) SetCreatedBy(v RUMOperationUser) {
	o.CreatedBy = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RUMOperationResponseAttributes) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *RUMOperationResponseAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *RUMOperationResponseAttributes) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given datadog.NullableString and assigns it to the Description field.
func (o *RUMOperationResponseAttributes) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *RUMOperationResponseAttributes) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *RUMOperationResponseAttributes) UnsetDescription() {
	o.Description.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *RUMOperationResponseAttributes) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMOperationResponseAttributes) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *RUMOperationResponseAttributes) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *RUMOperationResponseAttributes) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetFeatureIds returns the FeatureIds field value if set, zero value otherwise.
func (o *RUMOperationResponseAttributes) GetFeatureIds() []string {
	if o == nil || o.FeatureIds == nil {
		var ret []string
		return ret
	}
	return o.FeatureIds
}

// GetFeatureIdsOk returns a tuple with the FeatureIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMOperationResponseAttributes) GetFeatureIdsOk() (*[]string, bool) {
	if o == nil || o.FeatureIds == nil {
		return nil, false
	}
	return &o.FeatureIds, true
}

// HasFeatureIds returns a boolean if a field has been set.
func (o *RUMOperationResponseAttributes) HasFeatureIds() bool {
	return o != nil && o.FeatureIds != nil
}

// SetFeatureIds gets a reference to the given []string and assigns it to the FeatureIds field.
func (o *RUMOperationResponseAttributes) SetFeatureIds(v []string) {
	o.FeatureIds = v
}

// GetJourneyRum returns the JourneyRum field value.
func (o *RUMOperationResponseAttributes) GetJourneyRum() RUMOperationJourneyRum {
	if o == nil {
		var ret RUMOperationJourneyRum
		return ret
	}
	return o.JourneyRum
}

// GetJourneyRumOk returns a tuple with the JourneyRum field value
// and a boolean to check if the value has been set.
func (o *RUMOperationResponseAttributes) GetJourneyRumOk() (*RUMOperationJourneyRum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JourneyRum, true
}

// SetJourneyRum sets field value.
func (o *RUMOperationResponseAttributes) SetJourneyRum(v RUMOperationJourneyRum) {
	o.JourneyRum = v
}

// GetName returns the Name field value.
func (o *RUMOperationResponseAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RUMOperationResponseAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *RUMOperationResponseAttributes) SetName(v string) {
	o.Name = v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *RUMOperationResponseAttributes) GetOrgId() int64 {
	if o == nil || o.OrgId == nil {
		var ret int64
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMOperationResponseAttributes) GetOrgIdOk() (*int64, bool) {
	if o == nil || o.OrgId == nil {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *RUMOperationResponseAttributes) HasOrgId() bool {
	return o != nil && o.OrgId != nil
}

// SetOrgId gets a reference to the given int64 and assigns it to the OrgId field.
func (o *RUMOperationResponseAttributes) SetOrgId(v int64) {
	o.OrgId = &v
}

// GetTags returns the Tags field value.
func (o *RUMOperationResponseAttributes) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *RUMOperationResponseAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value.
func (o *RUMOperationResponseAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RUMOperationResponseAttributes) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *RUMOperationResponseAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *RUMOperationResponseAttributes) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt.IsSet()
}

// SetUpdatedAt gets a reference to the given datadog.NullableTime and assigns it to the UpdatedAt field.
func (o *RUMOperationResponseAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil.
func (o *RUMOperationResponseAttributes) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil.
func (o *RUMOperationResponseAttributes) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *RUMOperationResponseAttributes) GetUpdatedBy() RUMOperationUser {
	if o == nil || o.UpdatedBy == nil {
		var ret RUMOperationUser
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMOperationResponseAttributes) GetUpdatedByOk() (*RUMOperationUser, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *RUMOperationResponseAttributes) HasUpdatedBy() bool {
	return o != nil && o.UpdatedBy != nil
}

// SetUpdatedBy gets a reference to the given RUMOperationUser and assigns it to the UpdatedBy field.
func (o *RUMOperationResponseAttributes) SetUpdatedBy(v RUMOperationUser) {
	o.UpdatedBy = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RUMOperationResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ApplicationId.IsSet() {
		toSerialize["application_id"] = o.ApplicationId.Get()
	}
	if o.Category.IsSet() {
		toSerialize["category"] = o.Category.Get()
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
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.FeatureIds != nil {
		toSerialize["feature_ids"] = o.FeatureIds
	}
	toSerialize["journey_rum"] = o.JourneyRum
	toSerialize["name"] = o.Name
	if o.OrgId != nil {
		toSerialize["org_id"] = o.OrgId
	}
	toSerialize["tags"] = o.Tags
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
	if o.UpdatedBy != nil {
		toSerialize["updated_by"] = o.UpdatedBy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RUMOperationResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApplicationId datadog.NullableUUID    `json:"application_id,omitempty"`
		Category      datadog.NullableString  `json:"category,omitempty"`
		CreatedAt     *time.Time              `json:"created_at,omitempty"`
		CreatedBy     *RUMOperationUser       `json:"created_by,omitempty"`
		Description   datadog.NullableString  `json:"description,omitempty"`
		DisplayName   *string                 `json:"display_name,omitempty"`
		FeatureIds    []string                `json:"feature_ids,omitempty"`
		JourneyRum    *RUMOperationJourneyRum `json:"journey_rum"`
		Name          *string                 `json:"name"`
		OrgId         *int64                  `json:"org_id,omitempty"`
		Tags          *[]string               `json:"tags"`
		UpdatedAt     datadog.NullableTime    `json:"updated_at,omitempty"`
		UpdatedBy     *RUMOperationUser       `json:"updated_by,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.JourneyRum == nil {
		return fmt.Errorf("required field journey_rum missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Tags == nil {
		return fmt.Errorf("required field tags missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"application_id", "category", "created_at", "created_by", "description", "display_name", "feature_ids", "journey_rum", "name", "org_id", "tags", "updated_at", "updated_by"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ApplicationId = all.ApplicationId
	o.Category = all.Category
	o.CreatedAt = all.CreatedAt
	if all.CreatedBy != nil && all.CreatedBy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CreatedBy = all.CreatedBy
	o.Description = all.Description
	o.DisplayName = all.DisplayName
	o.FeatureIds = all.FeatureIds
	if all.JourneyRum.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.JourneyRum = *all.JourneyRum
	o.Name = *all.Name
	o.OrgId = all.OrgId
	o.Tags = *all.Tags
	o.UpdatedAt = all.UpdatedAt
	if all.UpdatedBy != nil && all.UpdatedBy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.UpdatedBy = all.UpdatedBy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
