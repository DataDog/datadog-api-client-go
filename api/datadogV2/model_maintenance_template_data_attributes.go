// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MaintenanceTemplateDataAttributes The attributes of a maintenance template.
type MaintenanceTemplateDataAttributes struct {
	// The description shown when a maintenance created from this template is completed.
	CompletedDescription *string `json:"completed_description,omitempty"`
	// The IDs of the components affected by a maintenance created from this template.
	ComponentIds []string `json:"component_ids,omitempty"`
	// Timestamp of when the maintenance template was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The description shown while a maintenance created from this template is in progress.
	InProgressDescription *string `json:"in_progress_description,omitempty"`
	// The title used for a maintenance created from this template.
	MaintenanceTitle *string `json:"maintenance_title,omitempty"`
	// Timestamp of when the maintenance template was last modified.
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	// The name of the maintenance template.
	Name *string `json:"name,omitempty"`
	// The description shown when a maintenance created from this template is scheduled.
	ScheduledDescription *string `json:"scheduled_description,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMaintenanceTemplateDataAttributes instantiates a new MaintenanceTemplateDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMaintenanceTemplateDataAttributes() *MaintenanceTemplateDataAttributes {
	this := MaintenanceTemplateDataAttributes{}
	return &this
}

// NewMaintenanceTemplateDataAttributesWithDefaults instantiates a new MaintenanceTemplateDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMaintenanceTemplateDataAttributesWithDefaults() *MaintenanceTemplateDataAttributes {
	this := MaintenanceTemplateDataAttributes{}
	return &this
}

// GetCompletedDescription returns the CompletedDescription field value if set, zero value otherwise.
func (o *MaintenanceTemplateDataAttributes) GetCompletedDescription() string {
	if o == nil || o.CompletedDescription == nil {
		var ret string
		return ret
	}
	return *o.CompletedDescription
}

// GetCompletedDescriptionOk returns a tuple with the CompletedDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceTemplateDataAttributes) GetCompletedDescriptionOk() (*string, bool) {
	if o == nil || o.CompletedDescription == nil {
		return nil, false
	}
	return o.CompletedDescription, true
}

// HasCompletedDescription returns a boolean if a field has been set.
func (o *MaintenanceTemplateDataAttributes) HasCompletedDescription() bool {
	return o != nil && o.CompletedDescription != nil
}

// SetCompletedDescription gets a reference to the given string and assigns it to the CompletedDescription field.
func (o *MaintenanceTemplateDataAttributes) SetCompletedDescription(v string) {
	o.CompletedDescription = &v
}

// GetComponentIds returns the ComponentIds field value if set, zero value otherwise.
func (o *MaintenanceTemplateDataAttributes) GetComponentIds() []string {
	if o == nil || o.ComponentIds == nil {
		var ret []string
		return ret
	}
	return o.ComponentIds
}

// GetComponentIdsOk returns a tuple with the ComponentIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceTemplateDataAttributes) GetComponentIdsOk() (*[]string, bool) {
	if o == nil || o.ComponentIds == nil {
		return nil, false
	}
	return &o.ComponentIds, true
}

// HasComponentIds returns a boolean if a field has been set.
func (o *MaintenanceTemplateDataAttributes) HasComponentIds() bool {
	return o != nil && o.ComponentIds != nil
}

// SetComponentIds gets a reference to the given []string and assigns it to the ComponentIds field.
func (o *MaintenanceTemplateDataAttributes) SetComponentIds(v []string) {
	o.ComponentIds = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *MaintenanceTemplateDataAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceTemplateDataAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *MaintenanceTemplateDataAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *MaintenanceTemplateDataAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetInProgressDescription returns the InProgressDescription field value if set, zero value otherwise.
func (o *MaintenanceTemplateDataAttributes) GetInProgressDescription() string {
	if o == nil || o.InProgressDescription == nil {
		var ret string
		return ret
	}
	return *o.InProgressDescription
}

// GetInProgressDescriptionOk returns a tuple with the InProgressDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceTemplateDataAttributes) GetInProgressDescriptionOk() (*string, bool) {
	if o == nil || o.InProgressDescription == nil {
		return nil, false
	}
	return o.InProgressDescription, true
}

// HasInProgressDescription returns a boolean if a field has been set.
func (o *MaintenanceTemplateDataAttributes) HasInProgressDescription() bool {
	return o != nil && o.InProgressDescription != nil
}

// SetInProgressDescription gets a reference to the given string and assigns it to the InProgressDescription field.
func (o *MaintenanceTemplateDataAttributes) SetInProgressDescription(v string) {
	o.InProgressDescription = &v
}

// GetMaintenanceTitle returns the MaintenanceTitle field value if set, zero value otherwise.
func (o *MaintenanceTemplateDataAttributes) GetMaintenanceTitle() string {
	if o == nil || o.MaintenanceTitle == nil {
		var ret string
		return ret
	}
	return *o.MaintenanceTitle
}

// GetMaintenanceTitleOk returns a tuple with the MaintenanceTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceTemplateDataAttributes) GetMaintenanceTitleOk() (*string, bool) {
	if o == nil || o.MaintenanceTitle == nil {
		return nil, false
	}
	return o.MaintenanceTitle, true
}

// HasMaintenanceTitle returns a boolean if a field has been set.
func (o *MaintenanceTemplateDataAttributes) HasMaintenanceTitle() bool {
	return o != nil && o.MaintenanceTitle != nil
}

// SetMaintenanceTitle gets a reference to the given string and assigns it to the MaintenanceTitle field.
func (o *MaintenanceTemplateDataAttributes) SetMaintenanceTitle(v string) {
	o.MaintenanceTitle = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *MaintenanceTemplateDataAttributes) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceTemplateDataAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *MaintenanceTemplateDataAttributes) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt != nil
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *MaintenanceTemplateDataAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MaintenanceTemplateDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceTemplateDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MaintenanceTemplateDataAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MaintenanceTemplateDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetScheduledDescription returns the ScheduledDescription field value if set, zero value otherwise.
func (o *MaintenanceTemplateDataAttributes) GetScheduledDescription() string {
	if o == nil || o.ScheduledDescription == nil {
		var ret string
		return ret
	}
	return *o.ScheduledDescription
}

// GetScheduledDescriptionOk returns a tuple with the ScheduledDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceTemplateDataAttributes) GetScheduledDescriptionOk() (*string, bool) {
	if o == nil || o.ScheduledDescription == nil {
		return nil, false
	}
	return o.ScheduledDescription, true
}

// HasScheduledDescription returns a boolean if a field has been set.
func (o *MaintenanceTemplateDataAttributes) HasScheduledDescription() bool {
	return o != nil && o.ScheduledDescription != nil
}

// SetScheduledDescription gets a reference to the given string and assigns it to the ScheduledDescription field.
func (o *MaintenanceTemplateDataAttributes) SetScheduledDescription(v string) {
	o.ScheduledDescription = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MaintenanceTemplateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CompletedDescription != nil {
		toSerialize["completed_description"] = o.CompletedDescription
	}
	if o.ComponentIds != nil {
		toSerialize["component_ids"] = o.ComponentIds
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.InProgressDescription != nil {
		toSerialize["in_progress_description"] = o.InProgressDescription
	}
	if o.MaintenanceTitle != nil {
		toSerialize["maintenance_title"] = o.MaintenanceTitle
	}
	if o.ModifiedAt != nil {
		if o.ModifiedAt.Nanosecond() == 0 {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ScheduledDescription != nil {
		toSerialize["scheduled_description"] = o.ScheduledDescription
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MaintenanceTemplateDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CompletedDescription  *string    `json:"completed_description,omitempty"`
		ComponentIds          []string   `json:"component_ids,omitempty"`
		CreatedAt             *time.Time `json:"created_at,omitempty"`
		InProgressDescription *string    `json:"in_progress_description,omitempty"`
		MaintenanceTitle      *string    `json:"maintenance_title,omitempty"`
		ModifiedAt            *time.Time `json:"modified_at,omitempty"`
		Name                  *string    `json:"name,omitempty"`
		ScheduledDescription  *string    `json:"scheduled_description,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"completed_description", "component_ids", "created_at", "in_progress_description", "maintenance_title", "modified_at", "name", "scheduled_description"})
	} else {
		return err
	}
	o.CompletedDescription = all.CompletedDescription
	o.ComponentIds = all.ComponentIds
	o.CreatedAt = all.CreatedAt
	o.InProgressDescription = all.InProgressDescription
	o.MaintenanceTitle = all.MaintenanceTitle
	o.ModifiedAt = all.ModifiedAt
	o.Name = all.Name
	o.ScheduledDescription = all.ScheduledDescription

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
