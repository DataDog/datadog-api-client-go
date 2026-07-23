// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateMaintenanceTemplateRequestDataAttributes The attributes for creating a maintenance template.
type CreateMaintenanceTemplateRequestDataAttributes struct {
	// The description shown when a maintenance created from this template is completed.
	CompletedDescription *string `json:"completed_description,omitempty"`
	// The IDs of the components affected by a maintenance created from this template.
	ComponentIds []string `json:"component_ids,omitempty"`
	// The description shown while a maintenance created from this template is in progress.
	InProgressDescription *string `json:"in_progress_description,omitempty"`
	// The title used for a maintenance created from this template.
	MaintenanceTitle *string `json:"maintenance_title,omitempty"`
	// The name of the maintenance template.
	Name string `json:"name"`
	// The description shown when a maintenance created from this template is scheduled.
	ScheduledDescription *string `json:"scheduled_description,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateMaintenanceTemplateRequestDataAttributes instantiates a new CreateMaintenanceTemplateRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateMaintenanceTemplateRequestDataAttributes(name string) *CreateMaintenanceTemplateRequestDataAttributes {
	this := CreateMaintenanceTemplateRequestDataAttributes{}
	this.Name = name
	return &this
}

// NewCreateMaintenanceTemplateRequestDataAttributesWithDefaults instantiates a new CreateMaintenanceTemplateRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateMaintenanceTemplateRequestDataAttributesWithDefaults() *CreateMaintenanceTemplateRequestDataAttributes {
	this := CreateMaintenanceTemplateRequestDataAttributes{}
	return &this
}

// GetCompletedDescription returns the CompletedDescription field value if set, zero value otherwise.
func (o *CreateMaintenanceTemplateRequestDataAttributes) GetCompletedDescription() string {
	if o == nil || o.CompletedDescription == nil {
		var ret string
		return ret
	}
	return *o.CompletedDescription
}

// GetCompletedDescriptionOk returns a tuple with the CompletedDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMaintenanceTemplateRequestDataAttributes) GetCompletedDescriptionOk() (*string, bool) {
	if o == nil || o.CompletedDescription == nil {
		return nil, false
	}
	return o.CompletedDescription, true
}

// HasCompletedDescription returns a boolean if a field has been set.
func (o *CreateMaintenanceTemplateRequestDataAttributes) HasCompletedDescription() bool {
	return o != nil && o.CompletedDescription != nil
}

// SetCompletedDescription gets a reference to the given string and assigns it to the CompletedDescription field.
func (o *CreateMaintenanceTemplateRequestDataAttributes) SetCompletedDescription(v string) {
	o.CompletedDescription = &v
}

// GetComponentIds returns the ComponentIds field value if set, zero value otherwise.
func (o *CreateMaintenanceTemplateRequestDataAttributes) GetComponentIds() []string {
	if o == nil || o.ComponentIds == nil {
		var ret []string
		return ret
	}
	return o.ComponentIds
}

// GetComponentIdsOk returns a tuple with the ComponentIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMaintenanceTemplateRequestDataAttributes) GetComponentIdsOk() (*[]string, bool) {
	if o == nil || o.ComponentIds == nil {
		return nil, false
	}
	return &o.ComponentIds, true
}

// HasComponentIds returns a boolean if a field has been set.
func (o *CreateMaintenanceTemplateRequestDataAttributes) HasComponentIds() bool {
	return o != nil && o.ComponentIds != nil
}

// SetComponentIds gets a reference to the given []string and assigns it to the ComponentIds field.
func (o *CreateMaintenanceTemplateRequestDataAttributes) SetComponentIds(v []string) {
	o.ComponentIds = v
}

// GetInProgressDescription returns the InProgressDescription field value if set, zero value otherwise.
func (o *CreateMaintenanceTemplateRequestDataAttributes) GetInProgressDescription() string {
	if o == nil || o.InProgressDescription == nil {
		var ret string
		return ret
	}
	return *o.InProgressDescription
}

// GetInProgressDescriptionOk returns a tuple with the InProgressDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMaintenanceTemplateRequestDataAttributes) GetInProgressDescriptionOk() (*string, bool) {
	if o == nil || o.InProgressDescription == nil {
		return nil, false
	}
	return o.InProgressDescription, true
}

// HasInProgressDescription returns a boolean if a field has been set.
func (o *CreateMaintenanceTemplateRequestDataAttributes) HasInProgressDescription() bool {
	return o != nil && o.InProgressDescription != nil
}

// SetInProgressDescription gets a reference to the given string and assigns it to the InProgressDescription field.
func (o *CreateMaintenanceTemplateRequestDataAttributes) SetInProgressDescription(v string) {
	o.InProgressDescription = &v
}

// GetMaintenanceTitle returns the MaintenanceTitle field value if set, zero value otherwise.
func (o *CreateMaintenanceTemplateRequestDataAttributes) GetMaintenanceTitle() string {
	if o == nil || o.MaintenanceTitle == nil {
		var ret string
		return ret
	}
	return *o.MaintenanceTitle
}

// GetMaintenanceTitleOk returns a tuple with the MaintenanceTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMaintenanceTemplateRequestDataAttributes) GetMaintenanceTitleOk() (*string, bool) {
	if o == nil || o.MaintenanceTitle == nil {
		return nil, false
	}
	return o.MaintenanceTitle, true
}

// HasMaintenanceTitle returns a boolean if a field has been set.
func (o *CreateMaintenanceTemplateRequestDataAttributes) HasMaintenanceTitle() bool {
	return o != nil && o.MaintenanceTitle != nil
}

// SetMaintenanceTitle gets a reference to the given string and assigns it to the MaintenanceTitle field.
func (o *CreateMaintenanceTemplateRequestDataAttributes) SetMaintenanceTitle(v string) {
	o.MaintenanceTitle = &v
}

// GetName returns the Name field value.
func (o *CreateMaintenanceTemplateRequestDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateMaintenanceTemplateRequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CreateMaintenanceTemplateRequestDataAttributes) SetName(v string) {
	o.Name = v
}

// GetScheduledDescription returns the ScheduledDescription field value if set, zero value otherwise.
func (o *CreateMaintenanceTemplateRequestDataAttributes) GetScheduledDescription() string {
	if o == nil || o.ScheduledDescription == nil {
		var ret string
		return ret
	}
	return *o.ScheduledDescription
}

// GetScheduledDescriptionOk returns a tuple with the ScheduledDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMaintenanceTemplateRequestDataAttributes) GetScheduledDescriptionOk() (*string, bool) {
	if o == nil || o.ScheduledDescription == nil {
		return nil, false
	}
	return o.ScheduledDescription, true
}

// HasScheduledDescription returns a boolean if a field has been set.
func (o *CreateMaintenanceTemplateRequestDataAttributes) HasScheduledDescription() bool {
	return o != nil && o.ScheduledDescription != nil
}

// SetScheduledDescription gets a reference to the given string and assigns it to the ScheduledDescription field.
func (o *CreateMaintenanceTemplateRequestDataAttributes) SetScheduledDescription(v string) {
	o.ScheduledDescription = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateMaintenanceTemplateRequestDataAttributes) MarshalJSON() ([]byte, error) {
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
	if o.InProgressDescription != nil {
		toSerialize["in_progress_description"] = o.InProgressDescription
	}
	if o.MaintenanceTitle != nil {
		toSerialize["maintenance_title"] = o.MaintenanceTitle
	}
	toSerialize["name"] = o.Name
	if o.ScheduledDescription != nil {
		toSerialize["scheduled_description"] = o.ScheduledDescription
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateMaintenanceTemplateRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CompletedDescription  *string  `json:"completed_description,omitempty"`
		ComponentIds          []string `json:"component_ids,omitempty"`
		InProgressDescription *string  `json:"in_progress_description,omitempty"`
		MaintenanceTitle      *string  `json:"maintenance_title,omitempty"`
		Name                  *string  `json:"name"`
		ScheduledDescription  *string  `json:"scheduled_description,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"completed_description", "component_ids", "in_progress_description", "maintenance_title", "name", "scheduled_description"})
	} else {
		return err
	}
	o.CompletedDescription = all.CompletedDescription
	o.ComponentIds = all.ComponentIds
	o.InProgressDescription = all.InProgressDescription
	o.MaintenanceTitle = all.MaintenanceTitle
	o.Name = *all.Name
	o.ScheduledDescription = all.ScheduledDescription

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
