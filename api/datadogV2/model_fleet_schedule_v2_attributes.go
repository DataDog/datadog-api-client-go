// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetScheduleV2Attributes Attributes of a fleet schedule in the v2 API response.
type FleetScheduleV2Attributes struct {
	// RFC3339 timestamp when the schedule was created.
	CreatedAt *string `json:"created_at,omitempty"`
	// User handle of the person who created the schedule.
	CreatedBy *string `json:"created_by,omitempty"`
	// Whether this is the default schedule for the organization.
	IsDefault *bool `json:"is_default,omitempty"`
	// Human-readable name for the schedule.
	Name *string `json:"name,omitempty"`
	// RFC3339 timestamp of the next scheduled maintenance window start time.
	// Absent when the next run time cannot be computed.
	NextRun *string `json:"next_run,omitempty"`
	// Notification configuration attached to a schedule.
	//
	// Included when available. If the notification rule cannot be retrieved, this field is
	// omitted and the schedule is still returned. If the notification rule is retrieved but its
	// handles cannot be resolved, it is still included with an empty `handles` array.
	NotificationRule *FleetScheduleV2NotificationRule `json:"notification_rule,omitempty"`
	// Query used to filter and select target hosts for scheduled deployments.
	Query *string `json:"query,omitempty"`
	// Defines the recurrence pattern for the schedule.
	Rule *FleetScheduleV2RecurrenceRule `json:"rule,omitempty"`
	// The status of the schedule.
	// - `active`: The schedule is active and will create deployments according to its recurrence rule.
	// - `inactive`: The schedule is inactive and will not create any deployments.
	Status *FleetScheduleStatus `json:"status,omitempty"`
	// RFC3339 timestamp when the schedule was last updated.
	UpdatedAt *string `json:"updated_at,omitempty"`
	// User handle of the person who last updated the schedule.
	UpdatedBy *string `json:"updated_by,omitempty"`
	// Number of major versions behind the latest to target for upgrades.
	// - 0: Always upgrade to the latest version.
	// - 1: Upgrade to latest minus 1 major version.
	// - 2: Upgrade to latest minus 2 major versions.
	VersionToLatest *int64 `json:"version_to_latest,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetScheduleV2Attributes instantiates a new FleetScheduleV2Attributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetScheduleV2Attributes() *FleetScheduleV2Attributes {
	this := FleetScheduleV2Attributes{}
	return &this
}

// NewFleetScheduleV2AttributesWithDefaults instantiates a new FleetScheduleV2Attributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetScheduleV2AttributesWithDefaults() *FleetScheduleV2Attributes {
	this := FleetScheduleV2Attributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *FleetScheduleV2Attributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetScheduleV2Attributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *FleetScheduleV2Attributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *FleetScheduleV2Attributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *FleetScheduleV2Attributes) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetScheduleV2Attributes) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *FleetScheduleV2Attributes) HasCreatedBy() bool {
	return o != nil && o.CreatedBy != nil
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *FleetScheduleV2Attributes) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *FleetScheduleV2Attributes) GetIsDefault() bool {
	if o == nil || o.IsDefault == nil {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetScheduleV2Attributes) GetIsDefaultOk() (*bool, bool) {
	if o == nil || o.IsDefault == nil {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *FleetScheduleV2Attributes) HasIsDefault() bool {
	return o != nil && o.IsDefault != nil
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *FleetScheduleV2Attributes) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FleetScheduleV2Attributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetScheduleV2Attributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FleetScheduleV2Attributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FleetScheduleV2Attributes) SetName(v string) {
	o.Name = &v
}

// GetNextRun returns the NextRun field value if set, zero value otherwise.
func (o *FleetScheduleV2Attributes) GetNextRun() string {
	if o == nil || o.NextRun == nil {
		var ret string
		return ret
	}
	return *o.NextRun
}

// GetNextRunOk returns a tuple with the NextRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetScheduleV2Attributes) GetNextRunOk() (*string, bool) {
	if o == nil || o.NextRun == nil {
		return nil, false
	}
	return o.NextRun, true
}

// HasNextRun returns a boolean if a field has been set.
func (o *FleetScheduleV2Attributes) HasNextRun() bool {
	return o != nil && o.NextRun != nil
}

// SetNextRun gets a reference to the given string and assigns it to the NextRun field.
func (o *FleetScheduleV2Attributes) SetNextRun(v string) {
	o.NextRun = &v
}

// GetNotificationRule returns the NotificationRule field value if set, zero value otherwise.
func (o *FleetScheduleV2Attributes) GetNotificationRule() FleetScheduleV2NotificationRule {
	if o == nil || o.NotificationRule == nil {
		var ret FleetScheduleV2NotificationRule
		return ret
	}
	return *o.NotificationRule
}

// GetNotificationRuleOk returns a tuple with the NotificationRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetScheduleV2Attributes) GetNotificationRuleOk() (*FleetScheduleV2NotificationRule, bool) {
	if o == nil || o.NotificationRule == nil {
		return nil, false
	}
	return o.NotificationRule, true
}

// HasNotificationRule returns a boolean if a field has been set.
func (o *FleetScheduleV2Attributes) HasNotificationRule() bool {
	return o != nil && o.NotificationRule != nil
}

// SetNotificationRule gets a reference to the given FleetScheduleV2NotificationRule and assigns it to the NotificationRule field.
func (o *FleetScheduleV2Attributes) SetNotificationRule(v FleetScheduleV2NotificationRule) {
	o.NotificationRule = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *FleetScheduleV2Attributes) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetScheduleV2Attributes) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *FleetScheduleV2Attributes) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *FleetScheduleV2Attributes) SetQuery(v string) {
	o.Query = &v
}

// GetRule returns the Rule field value if set, zero value otherwise.
func (o *FleetScheduleV2Attributes) GetRule() FleetScheduleV2RecurrenceRule {
	if o == nil || o.Rule == nil {
		var ret FleetScheduleV2RecurrenceRule
		return ret
	}
	return *o.Rule
}

// GetRuleOk returns a tuple with the Rule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetScheduleV2Attributes) GetRuleOk() (*FleetScheduleV2RecurrenceRule, bool) {
	if o == nil || o.Rule == nil {
		return nil, false
	}
	return o.Rule, true
}

// HasRule returns a boolean if a field has been set.
func (o *FleetScheduleV2Attributes) HasRule() bool {
	return o != nil && o.Rule != nil
}

// SetRule gets a reference to the given FleetScheduleV2RecurrenceRule and assigns it to the Rule field.
func (o *FleetScheduleV2Attributes) SetRule(v FleetScheduleV2RecurrenceRule) {
	o.Rule = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FleetScheduleV2Attributes) GetStatus() FleetScheduleStatus {
	if o == nil || o.Status == nil {
		var ret FleetScheduleStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetScheduleV2Attributes) GetStatusOk() (*FleetScheduleStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FleetScheduleV2Attributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given FleetScheduleStatus and assigns it to the Status field.
func (o *FleetScheduleV2Attributes) SetStatus(v FleetScheduleStatus) {
	o.Status = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *FleetScheduleV2Attributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetScheduleV2Attributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *FleetScheduleV2Attributes) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *FleetScheduleV2Attributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *FleetScheduleV2Attributes) GetUpdatedBy() string {
	if o == nil || o.UpdatedBy == nil {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetScheduleV2Attributes) GetUpdatedByOk() (*string, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *FleetScheduleV2Attributes) HasUpdatedBy() bool {
	return o != nil && o.UpdatedBy != nil
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *FleetScheduleV2Attributes) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

// GetVersionToLatest returns the VersionToLatest field value if set, zero value otherwise.
func (o *FleetScheduleV2Attributes) GetVersionToLatest() int64 {
	if o == nil || o.VersionToLatest == nil {
		var ret int64
		return ret
	}
	return *o.VersionToLatest
}

// GetVersionToLatestOk returns a tuple with the VersionToLatest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetScheduleV2Attributes) GetVersionToLatestOk() (*int64, bool) {
	if o == nil || o.VersionToLatest == nil {
		return nil, false
	}
	return o.VersionToLatest, true
}

// HasVersionToLatest returns a boolean if a field has been set.
func (o *FleetScheduleV2Attributes) HasVersionToLatest() bool {
	return o != nil && o.VersionToLatest != nil
}

// SetVersionToLatest gets a reference to the given int64 and assigns it to the VersionToLatest field.
func (o *FleetScheduleV2Attributes) SetVersionToLatest(v int64) {
	o.VersionToLatest = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetScheduleV2Attributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.IsDefault != nil {
		toSerialize["is_default"] = o.IsDefault
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NextRun != nil {
		toSerialize["next_run"] = o.NextRun
	}
	if o.NotificationRule != nil {
		toSerialize["notification_rule"] = o.NotificationRule
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.Rule != nil {
		toSerialize["rule"] = o.Rule
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.UpdatedBy != nil {
		toSerialize["updated_by"] = o.UpdatedBy
	}
	if o.VersionToLatest != nil {
		toSerialize["version_to_latest"] = o.VersionToLatest
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetScheduleV2Attributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt        *string                          `json:"created_at,omitempty"`
		CreatedBy        *string                          `json:"created_by,omitempty"`
		IsDefault        *bool                            `json:"is_default,omitempty"`
		Name             *string                          `json:"name,omitempty"`
		NextRun          *string                          `json:"next_run,omitempty"`
		NotificationRule *FleetScheduleV2NotificationRule `json:"notification_rule,omitempty"`
		Query            *string                          `json:"query,omitempty"`
		Rule             *FleetScheduleV2RecurrenceRule   `json:"rule,omitempty"`
		Status           *FleetScheduleStatus             `json:"status,omitempty"`
		UpdatedAt        *string                          `json:"updated_at,omitempty"`
		UpdatedBy        *string                          `json:"updated_by,omitempty"`
		VersionToLatest  *int64                           `json:"version_to_latest,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "created_by", "is_default", "name", "next_run", "notification_rule", "query", "rule", "status", "updated_at", "updated_by", "version_to_latest"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = all.CreatedAt
	o.CreatedBy = all.CreatedBy
	o.IsDefault = all.IsDefault
	o.Name = all.Name
	o.NextRun = all.NextRun
	if all.NotificationRule != nil && all.NotificationRule.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.NotificationRule = all.NotificationRule
	o.Query = all.Query
	if all.Rule != nil && all.Rule.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Rule = all.Rule
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	o.UpdatedAt = all.UpdatedAt
	o.UpdatedBy = all.UpdatedBy
	o.VersionToLatest = all.VersionToLatest

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
