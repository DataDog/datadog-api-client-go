// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentJiraTemplateFieldConfiguration Configuration for a Jira field mapping.
type IncidentJiraTemplateFieldConfiguration struct {
	// Custom value for outbound synchronization.
	CustomOutboundValue interface{} `json:"custom_outbound_value,omitempty"`
	// The incident field to map to.
	IncidentField datadog.NullableString `json:"incident_field,omitempty"`
	// The Jira field key.
	JiraFieldKey string `json:"jira_field_key"`
	// The type of the Jira field.
	JiraFieldType datadog.NullableString `json:"jira_field_type,omitempty"`
	// The direction of synchronization.
	SyncDirection string `json:"sync_direction"`
	// Mapping of values between incident and Jira fields.
	ValueMapping map[string]string `json:"value_mapping,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentJiraTemplateFieldConfiguration instantiates a new IncidentJiraTemplateFieldConfiguration object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentJiraTemplateFieldConfiguration(jiraFieldKey string, syncDirection string) *IncidentJiraTemplateFieldConfiguration {
	this := IncidentJiraTemplateFieldConfiguration{}
	this.JiraFieldKey = jiraFieldKey
	this.SyncDirection = syncDirection
	return &this
}

// NewIncidentJiraTemplateFieldConfigurationWithDefaults instantiates a new IncidentJiraTemplateFieldConfiguration object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentJiraTemplateFieldConfigurationWithDefaults() *IncidentJiraTemplateFieldConfiguration {
	this := IncidentJiraTemplateFieldConfiguration{}
	return &this
}

// GetCustomOutboundValue returns the CustomOutboundValue field value if set, zero value otherwise.
func (o *IncidentJiraTemplateFieldConfiguration) GetCustomOutboundValue() interface{} {
	if o == nil || o.CustomOutboundValue == nil {
		var ret interface{}
		return ret
	}
	return o.CustomOutboundValue
}

// GetCustomOutboundValueOk returns a tuple with the CustomOutboundValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentJiraTemplateFieldConfiguration) GetCustomOutboundValueOk() (*interface{}, bool) {
	if o == nil || o.CustomOutboundValue == nil {
		return nil, false
	}
	return &o.CustomOutboundValue, true
}

// HasCustomOutboundValue returns a boolean if a field has been set.
func (o *IncidentJiraTemplateFieldConfiguration) HasCustomOutboundValue() bool {
	return o != nil && o.CustomOutboundValue != nil
}

// SetCustomOutboundValue gets a reference to the given interface{} and assigns it to the CustomOutboundValue field.
func (o *IncidentJiraTemplateFieldConfiguration) SetCustomOutboundValue(v interface{}) {
	o.CustomOutboundValue = v
}

// GetIncidentField returns the IncidentField field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentJiraTemplateFieldConfiguration) GetIncidentField() string {
	if o == nil || o.IncidentField.Get() == nil {
		var ret string
		return ret
	}
	return *o.IncidentField.Get()
}

// GetIncidentFieldOk returns a tuple with the IncidentField field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentJiraTemplateFieldConfiguration) GetIncidentFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IncidentField.Get(), o.IncidentField.IsSet()
}

// HasIncidentField returns a boolean if a field has been set.
func (o *IncidentJiraTemplateFieldConfiguration) HasIncidentField() bool {
	return o != nil && o.IncidentField.IsSet()
}

// SetIncidentField gets a reference to the given datadog.NullableString and assigns it to the IncidentField field.
func (o *IncidentJiraTemplateFieldConfiguration) SetIncidentField(v string) {
	o.IncidentField.Set(&v)
}

// SetIncidentFieldNil sets the value for IncidentField to be an explicit nil.
func (o *IncidentJiraTemplateFieldConfiguration) SetIncidentFieldNil() {
	o.IncidentField.Set(nil)
}

// UnsetIncidentField ensures that no value is present for IncidentField, not even an explicit nil.
func (o *IncidentJiraTemplateFieldConfiguration) UnsetIncidentField() {
	o.IncidentField.Unset()
}

// GetJiraFieldKey returns the JiraFieldKey field value.
func (o *IncidentJiraTemplateFieldConfiguration) GetJiraFieldKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.JiraFieldKey
}

// GetJiraFieldKeyOk returns a tuple with the JiraFieldKey field value
// and a boolean to check if the value has been set.
func (o *IncidentJiraTemplateFieldConfiguration) GetJiraFieldKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JiraFieldKey, true
}

// SetJiraFieldKey sets field value.
func (o *IncidentJiraTemplateFieldConfiguration) SetJiraFieldKey(v string) {
	o.JiraFieldKey = v
}

// GetJiraFieldType returns the JiraFieldType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentJiraTemplateFieldConfiguration) GetJiraFieldType() string {
	if o == nil || o.JiraFieldType.Get() == nil {
		var ret string
		return ret
	}
	return *o.JiraFieldType.Get()
}

// GetJiraFieldTypeOk returns a tuple with the JiraFieldType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentJiraTemplateFieldConfiguration) GetJiraFieldTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.JiraFieldType.Get(), o.JiraFieldType.IsSet()
}

// HasJiraFieldType returns a boolean if a field has been set.
func (o *IncidentJiraTemplateFieldConfiguration) HasJiraFieldType() bool {
	return o != nil && o.JiraFieldType.IsSet()
}

// SetJiraFieldType gets a reference to the given datadog.NullableString and assigns it to the JiraFieldType field.
func (o *IncidentJiraTemplateFieldConfiguration) SetJiraFieldType(v string) {
	o.JiraFieldType.Set(&v)
}

// SetJiraFieldTypeNil sets the value for JiraFieldType to be an explicit nil.
func (o *IncidentJiraTemplateFieldConfiguration) SetJiraFieldTypeNil() {
	o.JiraFieldType.Set(nil)
}

// UnsetJiraFieldType ensures that no value is present for JiraFieldType, not even an explicit nil.
func (o *IncidentJiraTemplateFieldConfiguration) UnsetJiraFieldType() {
	o.JiraFieldType.Unset()
}

// GetSyncDirection returns the SyncDirection field value.
func (o *IncidentJiraTemplateFieldConfiguration) GetSyncDirection() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SyncDirection
}

// GetSyncDirectionOk returns a tuple with the SyncDirection field value
// and a boolean to check if the value has been set.
func (o *IncidentJiraTemplateFieldConfiguration) GetSyncDirectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SyncDirection, true
}

// SetSyncDirection sets field value.
func (o *IncidentJiraTemplateFieldConfiguration) SetSyncDirection(v string) {
	o.SyncDirection = v
}

// GetValueMapping returns the ValueMapping field value if set, zero value otherwise.
func (o *IncidentJiraTemplateFieldConfiguration) GetValueMapping() map[string]string {
	if o == nil || o.ValueMapping == nil {
		var ret map[string]string
		return ret
	}
	return o.ValueMapping
}

// GetValueMappingOk returns a tuple with the ValueMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentJiraTemplateFieldConfiguration) GetValueMappingOk() (*map[string]string, bool) {
	if o == nil || o.ValueMapping == nil {
		return nil, false
	}
	return &o.ValueMapping, true
}

// HasValueMapping returns a boolean if a field has been set.
func (o *IncidentJiraTemplateFieldConfiguration) HasValueMapping() bool {
	return o != nil && o.ValueMapping != nil
}

// SetValueMapping gets a reference to the given map[string]string and assigns it to the ValueMapping field.
func (o *IncidentJiraTemplateFieldConfiguration) SetValueMapping(v map[string]string) {
	o.ValueMapping = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentJiraTemplateFieldConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CustomOutboundValue != nil {
		toSerialize["custom_outbound_value"] = o.CustomOutboundValue
	}
	if o.IncidentField.IsSet() {
		toSerialize["incident_field"] = o.IncidentField.Get()
	}
	toSerialize["jira_field_key"] = o.JiraFieldKey
	if o.JiraFieldType.IsSet() {
		toSerialize["jira_field_type"] = o.JiraFieldType.Get()
	}
	toSerialize["sync_direction"] = o.SyncDirection
	if o.ValueMapping != nil {
		toSerialize["value_mapping"] = o.ValueMapping
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentJiraTemplateFieldConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CustomOutboundValue interface{}            `json:"custom_outbound_value,omitempty"`
		IncidentField       datadog.NullableString `json:"incident_field,omitempty"`
		JiraFieldKey        *string                `json:"jira_field_key"`
		JiraFieldType       datadog.NullableString `json:"jira_field_type,omitempty"`
		SyncDirection       *string                `json:"sync_direction"`
		ValueMapping        map[string]string      `json:"value_mapping,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.JiraFieldKey == nil {
		return fmt.Errorf("required field jira_field_key missing")
	}
	if all.SyncDirection == nil {
		return fmt.Errorf("required field sync_direction missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"custom_outbound_value", "incident_field", "jira_field_key", "jira_field_type", "sync_direction", "value_mapping"})
	} else {
		return err
	}
	o.CustomOutboundValue = all.CustomOutboundValue
	o.IncidentField = all.IncidentField
	o.JiraFieldKey = *all.JiraFieldKey
	o.JiraFieldType = all.JiraFieldType
	o.SyncDirection = *all.SyncDirection
	o.ValueMapping = all.ValueMapping

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
