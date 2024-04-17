// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringNotificationRuleResponseAttributes Attributes of the notification rule.
type SecurityMonitoringNotificationRuleResponseAttributes struct {
	// Timestamp of creation of the notification rule in milliseconds.
	CreationDate *int64 `json:"creation_date,omitempty"`
	// The author of the notification rule.
	Creator *SecurityMonitoringCreator `json:"creator,omitempty"`
	// Whether the notification rule is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The name of the notification rule.
	Name *string `json:"name,omitempty"`
	// Selectors describing the notification rule.
	Selectors *SecurityMonitoringNotificationRuleSelectors `json:"selectors,omitempty"`
	// Set of targets to notify.
	Targets []string `json:"targets,omitempty"`
	// Timestamp of last modification of the notification rule in milliseconds.
	UpdateDate *int64 `json:"update_date,omitempty"`
	// The editor of the notification rule.
	Updater *SecurityMonitoringUpdater `json:"updater,omitempty"`
	// The version of the rule being updated.
	Version *int32 `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewSecurityMonitoringNotificationRuleResponseAttributes instantiates a new SecurityMonitoringNotificationRuleResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringNotificationRuleResponseAttributes() *SecurityMonitoringNotificationRuleResponseAttributes {
	this := SecurityMonitoringNotificationRuleResponseAttributes{}
	return &this
}

// NewSecurityMonitoringNotificationRuleResponseAttributesWithDefaults instantiates a new SecurityMonitoringNotificationRuleResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringNotificationRuleResponseAttributesWithDefaults() *SecurityMonitoringNotificationRuleResponseAttributes {
	this := SecurityMonitoringNotificationRuleResponseAttributes{}
	return &this
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *SecurityMonitoringNotificationRuleResponseAttributes) GetCreationDate() int64 {
	if o == nil || o.CreationDate == nil {
		var ret int64
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringNotificationRuleResponseAttributes) GetCreationDateOk() (*int64, bool) {
	if o == nil || o.CreationDate == nil {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *SecurityMonitoringNotificationRuleResponseAttributes) HasCreationDate() bool {
	return o != nil && o.CreationDate != nil
}

// SetCreationDate gets a reference to the given int64 and assigns it to the CreationDate field.
func (o *SecurityMonitoringNotificationRuleResponseAttributes) SetCreationDate(v int64) {
	o.CreationDate = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *SecurityMonitoringNotificationRuleResponseAttributes) GetCreator() SecurityMonitoringCreator {
	if o == nil || o.Creator == nil {
		var ret SecurityMonitoringCreator
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringNotificationRuleResponseAttributes) GetCreatorOk() (*SecurityMonitoringCreator, bool) {
	if o == nil || o.Creator == nil {
		return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *SecurityMonitoringNotificationRuleResponseAttributes) HasCreator() bool {
	return o != nil && o.Creator != nil
}

// SetCreator gets a reference to the given SecurityMonitoringCreator and assigns it to the Creator field.
func (o *SecurityMonitoringNotificationRuleResponseAttributes) SetCreator(v SecurityMonitoringCreator) {
	o.Creator = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SecurityMonitoringNotificationRuleResponseAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringNotificationRuleResponseAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SecurityMonitoringNotificationRuleResponseAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SecurityMonitoringNotificationRuleResponseAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SecurityMonitoringNotificationRuleResponseAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringNotificationRuleResponseAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SecurityMonitoringNotificationRuleResponseAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SecurityMonitoringNotificationRuleResponseAttributes) SetName(v string) {
	o.Name = &v
}

// GetSelectors returns the Selectors field value if set, zero value otherwise.
func (o *SecurityMonitoringNotificationRuleResponseAttributes) GetSelectors() SecurityMonitoringNotificationRuleSelectors {
	if o == nil || o.Selectors == nil {
		var ret SecurityMonitoringNotificationRuleSelectors
		return ret
	}
	return *o.Selectors
}

// GetSelectorsOk returns a tuple with the Selectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringNotificationRuleResponseAttributes) GetSelectorsOk() (*SecurityMonitoringNotificationRuleSelectors, bool) {
	if o == nil || o.Selectors == nil {
		return nil, false
	}
	return o.Selectors, true
}

// HasSelectors returns a boolean if a field has been set.
func (o *SecurityMonitoringNotificationRuleResponseAttributes) HasSelectors() bool {
	return o != nil && o.Selectors != nil
}

// SetSelectors gets a reference to the given SecurityMonitoringNotificationRuleSelectors and assigns it to the Selectors field.
func (o *SecurityMonitoringNotificationRuleResponseAttributes) SetSelectors(v SecurityMonitoringNotificationRuleSelectors) {
	o.Selectors = &v
}

// GetTargets returns the Targets field value if set, zero value otherwise.
func (o *SecurityMonitoringNotificationRuleResponseAttributes) GetTargets() []string {
	if o == nil || o.Targets == nil {
		var ret []string
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringNotificationRuleResponseAttributes) GetTargetsOk() (*[]string, bool) {
	if o == nil || o.Targets == nil {
		return nil, false
	}
	return &o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *SecurityMonitoringNotificationRuleResponseAttributes) HasTargets() bool {
	return o != nil && o.Targets != nil
}

// SetTargets gets a reference to the given []string and assigns it to the Targets field.
func (o *SecurityMonitoringNotificationRuleResponseAttributes) SetTargets(v []string) {
	o.Targets = v
}

// GetUpdateDate returns the UpdateDate field value if set, zero value otherwise.
func (o *SecurityMonitoringNotificationRuleResponseAttributes) GetUpdateDate() int64 {
	if o == nil || o.UpdateDate == nil {
		var ret int64
		return ret
	}
	return *o.UpdateDate
}

// GetUpdateDateOk returns a tuple with the UpdateDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringNotificationRuleResponseAttributes) GetUpdateDateOk() (*int64, bool) {
	if o == nil || o.UpdateDate == nil {
		return nil, false
	}
	return o.UpdateDate, true
}

// HasUpdateDate returns a boolean if a field has been set.
func (o *SecurityMonitoringNotificationRuleResponseAttributes) HasUpdateDate() bool {
	return o != nil && o.UpdateDate != nil
}

// SetUpdateDate gets a reference to the given int64 and assigns it to the UpdateDate field.
func (o *SecurityMonitoringNotificationRuleResponseAttributes) SetUpdateDate(v int64) {
	o.UpdateDate = &v
}

// GetUpdater returns the Updater field value if set, zero value otherwise.
func (o *SecurityMonitoringNotificationRuleResponseAttributes) GetUpdater() SecurityMonitoringUpdater {
	if o == nil || o.Updater == nil {
		var ret SecurityMonitoringUpdater
		return ret
	}
	return *o.Updater
}

// GetUpdaterOk returns a tuple with the Updater field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringNotificationRuleResponseAttributes) GetUpdaterOk() (*SecurityMonitoringUpdater, bool) {
	if o == nil || o.Updater == nil {
		return nil, false
	}
	return o.Updater, true
}

// HasUpdater returns a boolean if a field has been set.
func (o *SecurityMonitoringNotificationRuleResponseAttributes) HasUpdater() bool {
	return o != nil && o.Updater != nil
}

// SetUpdater gets a reference to the given SecurityMonitoringUpdater and assigns it to the Updater field.
func (o *SecurityMonitoringNotificationRuleResponseAttributes) SetUpdater(v SecurityMonitoringUpdater) {
	o.Updater = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SecurityMonitoringNotificationRuleResponseAttributes) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringNotificationRuleResponseAttributes) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SecurityMonitoringNotificationRuleResponseAttributes) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *SecurityMonitoringNotificationRuleResponseAttributes) SetVersion(v int32) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringNotificationRuleResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreationDate != nil {
		toSerialize["creation_date"] = o.CreationDate
	}
	if o.Creator != nil {
		toSerialize["creator"] = o.Creator
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Selectors != nil {
		toSerialize["selectors"] = o.Selectors
	}
	if o.Targets != nil {
		toSerialize["targets"] = o.Targets
	}
	if o.UpdateDate != nil {
		toSerialize["update_date"] = o.UpdateDate
	}
	if o.Updater != nil {
		toSerialize["updater"] = o.Updater
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
func (o *SecurityMonitoringNotificationRuleResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreationDate *int64                                       `json:"creation_date,omitempty"`
		Creator      *SecurityMonitoringCreator                   `json:"creator,omitempty"`
		Enabled      *bool                                        `json:"enabled,omitempty"`
		Name         *string                                      `json:"name,omitempty"`
		Selectors    *SecurityMonitoringNotificationRuleSelectors `json:"selectors,omitempty"`
		Targets      []string                                     `json:"targets,omitempty"`
		UpdateDate   *int64                                       `json:"update_date,omitempty"`
		Updater      *SecurityMonitoringUpdater                   `json:"updater,omitempty"`
		Version      *int32                                       `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"creation_date", "creator", "enabled", "name", "selectors", "targets", "update_date", "updater", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreationDate = all.CreationDate
	if all.Creator != nil && all.Creator.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Creator = all.Creator
	o.Enabled = all.Enabled
	o.Name = all.Name
	if all.Selectors != nil && all.Selectors.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Selectors = all.Selectors
	o.Targets = all.Targets
	o.UpdateDate = all.UpdateDate
	if all.Updater != nil && all.Updater.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Updater = all.Updater
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
