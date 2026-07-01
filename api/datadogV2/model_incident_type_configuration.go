// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentTypeConfiguration The incident-type-scoped behavior settings. All fields are optional on update. Any field omitted from a PATCH request keeps its current value. This object is read-only on the incident type resource itself and is only mutated through the update (PATCH) endpoint.
type IncidentTypeConfiguration struct {
	// Whether incidents of this type can be deleted.
	AllowIncidentDeletion *bool `json:"allow_incident_deletion,omitempty"`
	// Whether automation workflows can be triggered for incidents of this type.
	AllowWorkflows *bool `json:"allow_workflows,omitempty"`
	// An optional message shown to users when they declare an incident of this type.
	CreateMessage *string `json:"create_message,omitempty"`
	// Whether the out-of-the-box postmortem template is disabled for incidents of this type.
	DisableOutOfTheBoxPostmortemTemplate *bool `json:"disable_out_of_the_box_postmortem_template,omitempty"`
	// Whether responders can edit incident timestamps for incidents of this type.
	EditableTimestamps *bool `json:"editable_timestamps,omitempty"`
	// Whether responders can create private incidents of this type. This is an opt-in setting, distinct from `private_incidents_by_default`, which controls whether incidents are created private automatically.
	PrivateIncidents *bool `json:"private_incidents,omitempty"`
	// Whether incidents of this type are created as private by default.
	PrivateIncidentsByDefault *bool `json:"private_incidents_by_default,omitempty"`
	// When set to `servicenow`, incidents will display the ServiceNow record ID instead of the public ID. If no ServiceNow integration exists, the public ID will be displayed.
	SlugSource *IncidentTypeSlugSource `json:"slug_source,omitempty"`
	// Whether incidents of this type are treated as test incidents.
	TestIncidents *bool `json:"test_incidents,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentTypeConfiguration instantiates a new IncidentTypeConfiguration object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentTypeConfiguration() *IncidentTypeConfiguration {
	this := IncidentTypeConfiguration{}
	var allowIncidentDeletion bool = false
	this.AllowIncidentDeletion = &allowIncidentDeletion
	var allowWorkflows bool = true
	this.AllowWorkflows = &allowWorkflows
	var disableOutOfTheBoxPostmortemTemplate bool = false
	this.DisableOutOfTheBoxPostmortemTemplate = &disableOutOfTheBoxPostmortemTemplate
	var editableTimestamps bool = false
	this.EditableTimestamps = &editableTimestamps
	var privateIncidents bool = false
	this.PrivateIncidents = &privateIncidents
	var privateIncidentsByDefault bool = false
	this.PrivateIncidentsByDefault = &privateIncidentsByDefault
	var slugSource IncidentTypeSlugSource = INCIDENTTYPESLUGSOURCE_DEFAULT
	this.SlugSource = &slugSource
	var testIncidents bool = true
	this.TestIncidents = &testIncidents
	return &this
}

// NewIncidentTypeConfigurationWithDefaults instantiates a new IncidentTypeConfiguration object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentTypeConfigurationWithDefaults() *IncidentTypeConfiguration {
	this := IncidentTypeConfiguration{}
	var allowIncidentDeletion bool = false
	this.AllowIncidentDeletion = &allowIncidentDeletion
	var allowWorkflows bool = true
	this.AllowWorkflows = &allowWorkflows
	var disableOutOfTheBoxPostmortemTemplate bool = false
	this.DisableOutOfTheBoxPostmortemTemplate = &disableOutOfTheBoxPostmortemTemplate
	var editableTimestamps bool = false
	this.EditableTimestamps = &editableTimestamps
	var privateIncidents bool = false
	this.PrivateIncidents = &privateIncidents
	var privateIncidentsByDefault bool = false
	this.PrivateIncidentsByDefault = &privateIncidentsByDefault
	var slugSource IncidentTypeSlugSource = INCIDENTTYPESLUGSOURCE_DEFAULT
	this.SlugSource = &slugSource
	var testIncidents bool = true
	this.TestIncidents = &testIncidents
	return &this
}

// GetAllowIncidentDeletion returns the AllowIncidentDeletion field value if set, zero value otherwise.
func (o *IncidentTypeConfiguration) GetAllowIncidentDeletion() bool {
	if o == nil || o.AllowIncidentDeletion == nil {
		var ret bool
		return ret
	}
	return *o.AllowIncidentDeletion
}

// GetAllowIncidentDeletionOk returns a tuple with the AllowIncidentDeletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTypeConfiguration) GetAllowIncidentDeletionOk() (*bool, bool) {
	if o == nil || o.AllowIncidentDeletion == nil {
		return nil, false
	}
	return o.AllowIncidentDeletion, true
}

// HasAllowIncidentDeletion returns a boolean if a field has been set.
func (o *IncidentTypeConfiguration) HasAllowIncidentDeletion() bool {
	return o != nil && o.AllowIncidentDeletion != nil
}

// SetAllowIncidentDeletion gets a reference to the given bool and assigns it to the AllowIncidentDeletion field.
func (o *IncidentTypeConfiguration) SetAllowIncidentDeletion(v bool) {
	o.AllowIncidentDeletion = &v
}

// GetAllowWorkflows returns the AllowWorkflows field value if set, zero value otherwise.
func (o *IncidentTypeConfiguration) GetAllowWorkflows() bool {
	if o == nil || o.AllowWorkflows == nil {
		var ret bool
		return ret
	}
	return *o.AllowWorkflows
}

// GetAllowWorkflowsOk returns a tuple with the AllowWorkflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTypeConfiguration) GetAllowWorkflowsOk() (*bool, bool) {
	if o == nil || o.AllowWorkflows == nil {
		return nil, false
	}
	return o.AllowWorkflows, true
}

// HasAllowWorkflows returns a boolean if a field has been set.
func (o *IncidentTypeConfiguration) HasAllowWorkflows() bool {
	return o != nil && o.AllowWorkflows != nil
}

// SetAllowWorkflows gets a reference to the given bool and assigns it to the AllowWorkflows field.
func (o *IncidentTypeConfiguration) SetAllowWorkflows(v bool) {
	o.AllowWorkflows = &v
}

// GetCreateMessage returns the CreateMessage field value if set, zero value otherwise.
func (o *IncidentTypeConfiguration) GetCreateMessage() string {
	if o == nil || o.CreateMessage == nil {
		var ret string
		return ret
	}
	return *o.CreateMessage
}

// GetCreateMessageOk returns a tuple with the CreateMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTypeConfiguration) GetCreateMessageOk() (*string, bool) {
	if o == nil || o.CreateMessage == nil {
		return nil, false
	}
	return o.CreateMessage, true
}

// HasCreateMessage returns a boolean if a field has been set.
func (o *IncidentTypeConfiguration) HasCreateMessage() bool {
	return o != nil && o.CreateMessage != nil
}

// SetCreateMessage gets a reference to the given string and assigns it to the CreateMessage field.
func (o *IncidentTypeConfiguration) SetCreateMessage(v string) {
	o.CreateMessage = &v
}

// GetDisableOutOfTheBoxPostmortemTemplate returns the DisableOutOfTheBoxPostmortemTemplate field value if set, zero value otherwise.
func (o *IncidentTypeConfiguration) GetDisableOutOfTheBoxPostmortemTemplate() bool {
	if o == nil || o.DisableOutOfTheBoxPostmortemTemplate == nil {
		var ret bool
		return ret
	}
	return *o.DisableOutOfTheBoxPostmortemTemplate
}

// GetDisableOutOfTheBoxPostmortemTemplateOk returns a tuple with the DisableOutOfTheBoxPostmortemTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTypeConfiguration) GetDisableOutOfTheBoxPostmortemTemplateOk() (*bool, bool) {
	if o == nil || o.DisableOutOfTheBoxPostmortemTemplate == nil {
		return nil, false
	}
	return o.DisableOutOfTheBoxPostmortemTemplate, true
}

// HasDisableOutOfTheBoxPostmortemTemplate returns a boolean if a field has been set.
func (o *IncidentTypeConfiguration) HasDisableOutOfTheBoxPostmortemTemplate() bool {
	return o != nil && o.DisableOutOfTheBoxPostmortemTemplate != nil
}

// SetDisableOutOfTheBoxPostmortemTemplate gets a reference to the given bool and assigns it to the DisableOutOfTheBoxPostmortemTemplate field.
func (o *IncidentTypeConfiguration) SetDisableOutOfTheBoxPostmortemTemplate(v bool) {
	o.DisableOutOfTheBoxPostmortemTemplate = &v
}

// GetEditableTimestamps returns the EditableTimestamps field value if set, zero value otherwise.
func (o *IncidentTypeConfiguration) GetEditableTimestamps() bool {
	if o == nil || o.EditableTimestamps == nil {
		var ret bool
		return ret
	}
	return *o.EditableTimestamps
}

// GetEditableTimestampsOk returns a tuple with the EditableTimestamps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTypeConfiguration) GetEditableTimestampsOk() (*bool, bool) {
	if o == nil || o.EditableTimestamps == nil {
		return nil, false
	}
	return o.EditableTimestamps, true
}

// HasEditableTimestamps returns a boolean if a field has been set.
func (o *IncidentTypeConfiguration) HasEditableTimestamps() bool {
	return o != nil && o.EditableTimestamps != nil
}

// SetEditableTimestamps gets a reference to the given bool and assigns it to the EditableTimestamps field.
func (o *IncidentTypeConfiguration) SetEditableTimestamps(v bool) {
	o.EditableTimestamps = &v
}

// GetPrivateIncidents returns the PrivateIncidents field value if set, zero value otherwise.
func (o *IncidentTypeConfiguration) GetPrivateIncidents() bool {
	if o == nil || o.PrivateIncidents == nil {
		var ret bool
		return ret
	}
	return *o.PrivateIncidents
}

// GetPrivateIncidentsOk returns a tuple with the PrivateIncidents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTypeConfiguration) GetPrivateIncidentsOk() (*bool, bool) {
	if o == nil || o.PrivateIncidents == nil {
		return nil, false
	}
	return o.PrivateIncidents, true
}

// HasPrivateIncidents returns a boolean if a field has been set.
func (o *IncidentTypeConfiguration) HasPrivateIncidents() bool {
	return o != nil && o.PrivateIncidents != nil
}

// SetPrivateIncidents gets a reference to the given bool and assigns it to the PrivateIncidents field.
func (o *IncidentTypeConfiguration) SetPrivateIncidents(v bool) {
	o.PrivateIncidents = &v
}

// GetPrivateIncidentsByDefault returns the PrivateIncidentsByDefault field value if set, zero value otherwise.
func (o *IncidentTypeConfiguration) GetPrivateIncidentsByDefault() bool {
	if o == nil || o.PrivateIncidentsByDefault == nil {
		var ret bool
		return ret
	}
	return *o.PrivateIncidentsByDefault
}

// GetPrivateIncidentsByDefaultOk returns a tuple with the PrivateIncidentsByDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTypeConfiguration) GetPrivateIncidentsByDefaultOk() (*bool, bool) {
	if o == nil || o.PrivateIncidentsByDefault == nil {
		return nil, false
	}
	return o.PrivateIncidentsByDefault, true
}

// HasPrivateIncidentsByDefault returns a boolean if a field has been set.
func (o *IncidentTypeConfiguration) HasPrivateIncidentsByDefault() bool {
	return o != nil && o.PrivateIncidentsByDefault != nil
}

// SetPrivateIncidentsByDefault gets a reference to the given bool and assigns it to the PrivateIncidentsByDefault field.
func (o *IncidentTypeConfiguration) SetPrivateIncidentsByDefault(v bool) {
	o.PrivateIncidentsByDefault = &v
}

// GetSlugSource returns the SlugSource field value if set, zero value otherwise.
func (o *IncidentTypeConfiguration) GetSlugSource() IncidentTypeSlugSource {
	if o == nil || o.SlugSource == nil {
		var ret IncidentTypeSlugSource
		return ret
	}
	return *o.SlugSource
}

// GetSlugSourceOk returns a tuple with the SlugSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTypeConfiguration) GetSlugSourceOk() (*IncidentTypeSlugSource, bool) {
	if o == nil || o.SlugSource == nil {
		return nil, false
	}
	return o.SlugSource, true
}

// HasSlugSource returns a boolean if a field has been set.
func (o *IncidentTypeConfiguration) HasSlugSource() bool {
	return o != nil && o.SlugSource != nil
}

// SetSlugSource gets a reference to the given IncidentTypeSlugSource and assigns it to the SlugSource field.
func (o *IncidentTypeConfiguration) SetSlugSource(v IncidentTypeSlugSource) {
	o.SlugSource = &v
}

// GetTestIncidents returns the TestIncidents field value if set, zero value otherwise.
func (o *IncidentTypeConfiguration) GetTestIncidents() bool {
	if o == nil || o.TestIncidents == nil {
		var ret bool
		return ret
	}
	return *o.TestIncidents
}

// GetTestIncidentsOk returns a tuple with the TestIncidents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTypeConfiguration) GetTestIncidentsOk() (*bool, bool) {
	if o == nil || o.TestIncidents == nil {
		return nil, false
	}
	return o.TestIncidents, true
}

// HasTestIncidents returns a boolean if a field has been set.
func (o *IncidentTypeConfiguration) HasTestIncidents() bool {
	return o != nil && o.TestIncidents != nil
}

// SetTestIncidents gets a reference to the given bool and assigns it to the TestIncidents field.
func (o *IncidentTypeConfiguration) SetTestIncidents(v bool) {
	o.TestIncidents = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentTypeConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AllowIncidentDeletion != nil {
		toSerialize["allow_incident_deletion"] = o.AllowIncidentDeletion
	}
	if o.AllowWorkflows != nil {
		toSerialize["allow_workflows"] = o.AllowWorkflows
	}
	if o.CreateMessage != nil {
		toSerialize["create_message"] = o.CreateMessage
	}
	if o.DisableOutOfTheBoxPostmortemTemplate != nil {
		toSerialize["disable_out_of_the_box_postmortem_template"] = o.DisableOutOfTheBoxPostmortemTemplate
	}
	if o.EditableTimestamps != nil {
		toSerialize["editable_timestamps"] = o.EditableTimestamps
	}
	if o.PrivateIncidents != nil {
		toSerialize["private_incidents"] = o.PrivateIncidents
	}
	if o.PrivateIncidentsByDefault != nil {
		toSerialize["private_incidents_by_default"] = o.PrivateIncidentsByDefault
	}
	if o.SlugSource != nil {
		toSerialize["slug_source"] = o.SlugSource
	}
	if o.TestIncidents != nil {
		toSerialize["test_incidents"] = o.TestIncidents
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentTypeConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AllowIncidentDeletion                *bool                   `json:"allow_incident_deletion,omitempty"`
		AllowWorkflows                       *bool                   `json:"allow_workflows,omitempty"`
		CreateMessage                        *string                 `json:"create_message,omitempty"`
		DisableOutOfTheBoxPostmortemTemplate *bool                   `json:"disable_out_of_the_box_postmortem_template,omitempty"`
		EditableTimestamps                   *bool                   `json:"editable_timestamps,omitempty"`
		PrivateIncidents                     *bool                   `json:"private_incidents,omitempty"`
		PrivateIncidentsByDefault            *bool                   `json:"private_incidents_by_default,omitempty"`
		SlugSource                           *IncidentTypeSlugSource `json:"slug_source,omitempty"`
		TestIncidents                        *bool                   `json:"test_incidents,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"allow_incident_deletion", "allow_workflows", "create_message", "disable_out_of_the_box_postmortem_template", "editable_timestamps", "private_incidents", "private_incidents_by_default", "slug_source", "test_incidents"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AllowIncidentDeletion = all.AllowIncidentDeletion
	o.AllowWorkflows = all.AllowWorkflows
	o.CreateMessage = all.CreateMessage
	o.DisableOutOfTheBoxPostmortemTemplate = all.DisableOutOfTheBoxPostmortemTemplate
	o.EditableTimestamps = all.EditableTimestamps
	o.PrivateIncidents = all.PrivateIncidents
	o.PrivateIncidentsByDefault = all.PrivateIncidentsByDefault
	if all.SlugSource != nil && !all.SlugSource.IsValid() {
		hasInvalidField = true
	} else {
		o.SlugSource = all.SlugSource
	}
	o.TestIncidents = all.TestIncidents

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
