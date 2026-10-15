// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SavedFilterAttributes Saved filter details, including the full targeting rules and conditions.
type SavedFilterAttributes struct {
	// The time the saved filter was archived. Omitted when the filter is not archived.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
	// The time the saved filter was created.
	CreatedAt time.Time `json:"created_at"`
	// The ID of the user who created the saved filter.
	CreatedBy string `json:"created_by"`
	// How the saved filter was authored. RULES indicates rule-based authoring and LIST indicates list import. Both types return targeting rules and conditions.
	CreationType SavedFilterCreationType `json:"creation_type"`
	// A description of the saved filter. Omitted when no description is set.
	Description *string `json:"description,omitempty"`
	// Optional JSON metadata supplied by migration tools for tracking the source of a saved filter.
	MigrationMetadata interface{} `json:"migration_metadata,omitempty"`
	// The name of the saved filter.
	Name string `json:"name"`
	// Targeting rules combined with OR. Conditions within each rule are combined with AND. Both saved lists and saved rules use this structure.
	TargetingRules []SavedFilterTargetingRule `json:"targeting_rules"`
	// The time the saved filter was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSavedFilterAttributes instantiates a new SavedFilterAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSavedFilterAttributes(createdAt time.Time, createdBy string, creationType SavedFilterCreationType, name string, targetingRules []SavedFilterTargetingRule, updatedAt time.Time) *SavedFilterAttributes {
	this := SavedFilterAttributes{}
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.CreationType = creationType
	this.Name = name
	this.TargetingRules = targetingRules
	this.UpdatedAt = updatedAt
	return &this
}

// NewSavedFilterAttributesWithDefaults instantiates a new SavedFilterAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSavedFilterAttributesWithDefaults() *SavedFilterAttributes {
	this := SavedFilterAttributes{}
	return &this
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise.
func (o *SavedFilterAttributes) GetArchivedAt() time.Time {
	if o == nil || o.ArchivedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ArchivedAt
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedFilterAttributes) GetArchivedAtOk() (*time.Time, bool) {
	if o == nil || o.ArchivedAt == nil {
		return nil, false
	}
	return o.ArchivedAt, true
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *SavedFilterAttributes) HasArchivedAt() bool {
	return o != nil && o.ArchivedAt != nil
}

// SetArchivedAt gets a reference to the given time.Time and assigns it to the ArchivedAt field.
func (o *SavedFilterAttributes) SetArchivedAt(v time.Time) {
	o.ArchivedAt = &v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *SavedFilterAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SavedFilterAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *SavedFilterAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value.
func (o *SavedFilterAttributes) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *SavedFilterAttributes) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value.
func (o *SavedFilterAttributes) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreationType returns the CreationType field value.
func (o *SavedFilterAttributes) GetCreationType() SavedFilterCreationType {
	if o == nil {
		var ret SavedFilterCreationType
		return ret
	}
	return o.CreationType
}

// GetCreationTypeOk returns a tuple with the CreationType field value
// and a boolean to check if the value has been set.
func (o *SavedFilterAttributes) GetCreationTypeOk() (*SavedFilterCreationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreationType, true
}

// SetCreationType sets field value.
func (o *SavedFilterAttributes) SetCreationType(v SavedFilterCreationType) {
	o.CreationType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SavedFilterAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedFilterAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SavedFilterAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SavedFilterAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetMigrationMetadata returns the MigrationMetadata field value if set, zero value otherwise.
func (o *SavedFilterAttributes) GetMigrationMetadata() interface{} {
	if o == nil || o.MigrationMetadata == nil {
		var ret interface{}
		return ret
	}
	return o.MigrationMetadata
}

// GetMigrationMetadataOk returns a tuple with the MigrationMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedFilterAttributes) GetMigrationMetadataOk() (*interface{}, bool) {
	if o == nil || o.MigrationMetadata == nil {
		return nil, false
	}
	return &o.MigrationMetadata, true
}

// HasMigrationMetadata returns a boolean if a field has been set.
func (o *SavedFilterAttributes) HasMigrationMetadata() bool {
	return o != nil && o.MigrationMetadata != nil
}

// SetMigrationMetadata gets a reference to the given interface{} and assigns it to the MigrationMetadata field.
func (o *SavedFilterAttributes) SetMigrationMetadata(v interface{}) {
	o.MigrationMetadata = v
}

// GetName returns the Name field value.
func (o *SavedFilterAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SavedFilterAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SavedFilterAttributes) SetName(v string) {
	o.Name = v
}

// GetTargetingRules returns the TargetingRules field value.
func (o *SavedFilterAttributes) GetTargetingRules() []SavedFilterTargetingRule {
	if o == nil {
		var ret []SavedFilterTargetingRule
		return ret
	}
	return o.TargetingRules
}

// GetTargetingRulesOk returns a tuple with the TargetingRules field value
// and a boolean to check if the value has been set.
func (o *SavedFilterAttributes) GetTargetingRulesOk() (*[]SavedFilterTargetingRule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetingRules, true
}

// SetTargetingRules sets field value.
func (o *SavedFilterAttributes) SetTargetingRules(v []SavedFilterTargetingRule) {
	o.TargetingRules = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *SavedFilterAttributes) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *SavedFilterAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *SavedFilterAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SavedFilterAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ArchivedAt != nil {
		if o.ArchivedAt.Nanosecond() == 0 {
			toSerialize["archived_at"] = o.ArchivedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["archived_at"] = o.ArchivedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["created_by"] = o.CreatedBy
	toSerialize["creation_type"] = o.CreationType
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.MigrationMetadata != nil {
		toSerialize["migration_metadata"] = o.MigrationMetadata
	}
	toSerialize["name"] = o.Name
	toSerialize["targeting_rules"] = o.TargetingRules
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SavedFilterAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ArchivedAt        *time.Time                  `json:"archived_at,omitempty"`
		CreatedAt         *time.Time                  `json:"created_at"`
		CreatedBy         *string                     `json:"created_by"`
		CreationType      *SavedFilterCreationType    `json:"creation_type"`
		Description       *string                     `json:"description,omitempty"`
		MigrationMetadata interface{}                 `json:"migration_metadata,omitempty"`
		Name              *string                     `json:"name"`
		TargetingRules    *[]SavedFilterTargetingRule `json:"targeting_rules"`
		UpdatedAt         *time.Time                  `json:"updated_at"`
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
	if all.CreationType == nil {
		return fmt.Errorf("required field creation_type missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.TargetingRules == nil {
		return fmt.Errorf("required field targeting_rules missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"archived_at", "created_at", "created_by", "creation_type", "description", "migration_metadata", "name", "targeting_rules", "updated_at"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ArchivedAt = all.ArchivedAt
	o.CreatedAt = *all.CreatedAt
	o.CreatedBy = *all.CreatedBy
	if !all.CreationType.IsValid() {
		hasInvalidField = true
	} else {
		o.CreationType = *all.CreationType
	}
	o.Description = all.Description
	o.MigrationMetadata = all.MigrationMetadata
	o.Name = *all.Name
	o.TargetingRules = *all.TargetingRules
	o.UpdatedAt = *all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
