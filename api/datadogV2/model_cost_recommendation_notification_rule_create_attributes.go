// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostRecommendationNotificationRuleCreateAttributes Attributes of a new recommendation notification rule. Schedule attributes depend on `frequency` and target
// attributes depend on `destination_type`.
type CostRecommendationNotificationRuleCreateAttributes struct {
	// The Slack channel ID, starting with `C` or `G`. Required for `slack` rules.
	ChannelId *string `json:"channel_id,omitempty"`
	// Where the digest is sent. `slack` requires `workspace_id` and `channel_id`; `teams` requires the
	// `teams_*` attributes. Defaults to `slack`.
	DestinationType *CostRecommendationNotificationRuleDestinationType `json:"destination_type,omitempty"`
	// Whether the digest is sent. Defaults to `true`.
	Enabled *bool `json:"enabled,omitempty"`
	// How often the digest is sent. `weekly` and `biweekly` use `weekly_execution_day` and `weekly_execution_time`;
	// `every_30` and `every_90` (days) use `interval_execution_date` and `interval_execution_time`.
	Frequency CostRecommendationNotificationRuleFrequency `json:"frequency"`
	// The first date (`YYYY-MM-DD`) of an `every_30` or `every_90` schedule.
	IntervalExecutionDate *string `json:"interval_execution_date,omitempty"`
	// The time of day (`HH:MM`, in `timezone`) of an `every_30` or `every_90` schedule.
	IntervalExecutionTime *string `json:"interval_execution_time,omitempty"`
	// Slack member IDs to mention in the digest. Only valid for `slack` rules.
	MentionTargets []string `json:"mention_targets,omitempty"`
	// Which sections the digest shows. Sent as a whole: sections left out of a request use their defaults, which
	// show everything.
	MessageCustomization *CostRecommendationNotificationRuleMessageCustomization `json:"message_customization,omitempty"`
	// The name of the rule, unique within the organization.
	Name string `json:"name"`
	// The recommendation search query that selects what the digest includes. Empty selects every recommendation.
	ScopeQuery *string `json:"scope_query,omitempty"`
	// The Microsoft Teams channel ID. Required for `teams` rules.
	TeamsChannelId *string `json:"teams_channel_id,omitempty"`
	// The Microsoft Teams channel name. Required for `teams` rules.
	TeamsChannelName *string `json:"teams_channel_name,omitempty"`
	// The Microsoft Teams team ID. Required for `teams` rules.
	TeamsTeamId *string `json:"teams_team_id,omitempty"`
	// The Microsoft Teams team name. Required for `teams` rules.
	TeamsTeamName *string `json:"teams_team_name,omitempty"`
	// The Microsoft Teams tenant ID. Required for `teams` rules.
	TeamsTenantId *string `json:"teams_tenant_id,omitempty"`
	// The Microsoft Teams tenant name. Required for `teams` rules.
	TeamsTenantName *string `json:"teams_tenant_name,omitempty"`
	// The IANA time zone of the schedule.
	Timezone string `json:"timezone"`
	// The day of the week of a `weekly` or `biweekly` schedule, such as `monday`.
	WeeklyExecutionDay *string `json:"weekly_execution_day,omitempty"`
	// The time of day (`HH:MM`, in `timezone`) of a `weekly` or `biweekly` schedule.
	WeeklyExecutionTime *string `json:"weekly_execution_time,omitempty"`
	// The Slack workspace ID, starting with `T`. Required for `slack` rules.
	WorkspaceId *string `json:"workspace_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCostRecommendationNotificationRuleCreateAttributes instantiates a new CostRecommendationNotificationRuleCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCostRecommendationNotificationRuleCreateAttributes(frequency CostRecommendationNotificationRuleFrequency, name string, timezone string) *CostRecommendationNotificationRuleCreateAttributes {
	this := CostRecommendationNotificationRuleCreateAttributes{}
	this.Frequency = frequency
	this.Name = name
	this.Timezone = timezone
	return &this
}

// NewCostRecommendationNotificationRuleCreateAttributesWithDefaults instantiates a new CostRecommendationNotificationRuleCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCostRecommendationNotificationRuleCreateAttributesWithDefaults() *CostRecommendationNotificationRuleCreateAttributes {
	this := CostRecommendationNotificationRuleCreateAttributes{}
	return &this
}

// GetChannelId returns the ChannelId field value if set, zero value otherwise.
func (o *CostRecommendationNotificationRuleCreateAttributes) GetChannelId() string {
	if o == nil || o.ChannelId == nil {
		var ret string
		return ret
	}
	return *o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleCreateAttributes) GetChannelIdOk() (*string, bool) {
	if o == nil || o.ChannelId == nil {
		return nil, false
	}
	return o.ChannelId, true
}

// HasChannelId returns a boolean if a field has been set.
func (o *CostRecommendationNotificationRuleCreateAttributes) HasChannelId() bool {
	return o != nil && o.ChannelId != nil
}

// SetChannelId gets a reference to the given string and assigns it to the ChannelId field.
func (o *CostRecommendationNotificationRuleCreateAttributes) SetChannelId(v string) {
	o.ChannelId = &v
}

// GetDestinationType returns the DestinationType field value if set, zero value otherwise.
func (o *CostRecommendationNotificationRuleCreateAttributes) GetDestinationType() CostRecommendationNotificationRuleDestinationType {
	if o == nil || o.DestinationType == nil {
		var ret CostRecommendationNotificationRuleDestinationType
		return ret
	}
	return *o.DestinationType
}

// GetDestinationTypeOk returns a tuple with the DestinationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleCreateAttributes) GetDestinationTypeOk() (*CostRecommendationNotificationRuleDestinationType, bool) {
	if o == nil || o.DestinationType == nil {
		return nil, false
	}
	return o.DestinationType, true
}

// HasDestinationType returns a boolean if a field has been set.
func (o *CostRecommendationNotificationRuleCreateAttributes) HasDestinationType() bool {
	return o != nil && o.DestinationType != nil
}

// SetDestinationType gets a reference to the given CostRecommendationNotificationRuleDestinationType and assigns it to the DestinationType field.
func (o *CostRecommendationNotificationRuleCreateAttributes) SetDestinationType(v CostRecommendationNotificationRuleDestinationType) {
	o.DestinationType = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CostRecommendationNotificationRuleCreateAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleCreateAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CostRecommendationNotificationRuleCreateAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CostRecommendationNotificationRuleCreateAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetFrequency returns the Frequency field value.
func (o *CostRecommendationNotificationRuleCreateAttributes) GetFrequency() CostRecommendationNotificationRuleFrequency {
	if o == nil {
		var ret CostRecommendationNotificationRuleFrequency
		return ret
	}
	return o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleCreateAttributes) GetFrequencyOk() (*CostRecommendationNotificationRuleFrequency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Frequency, true
}

// SetFrequency sets field value.
func (o *CostRecommendationNotificationRuleCreateAttributes) SetFrequency(v CostRecommendationNotificationRuleFrequency) {
	o.Frequency = v
}

// GetIntervalExecutionDate returns the IntervalExecutionDate field value if set, zero value otherwise.
func (o *CostRecommendationNotificationRuleCreateAttributes) GetIntervalExecutionDate() string {
	if o == nil || o.IntervalExecutionDate == nil {
		var ret string
		return ret
	}
	return *o.IntervalExecutionDate
}

// GetIntervalExecutionDateOk returns a tuple with the IntervalExecutionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleCreateAttributes) GetIntervalExecutionDateOk() (*string, bool) {
	if o == nil || o.IntervalExecutionDate == nil {
		return nil, false
	}
	return o.IntervalExecutionDate, true
}

// HasIntervalExecutionDate returns a boolean if a field has been set.
func (o *CostRecommendationNotificationRuleCreateAttributes) HasIntervalExecutionDate() bool {
	return o != nil && o.IntervalExecutionDate != nil
}

// SetIntervalExecutionDate gets a reference to the given string and assigns it to the IntervalExecutionDate field.
func (o *CostRecommendationNotificationRuleCreateAttributes) SetIntervalExecutionDate(v string) {
	o.IntervalExecutionDate = &v
}

// GetIntervalExecutionTime returns the IntervalExecutionTime field value if set, zero value otherwise.
func (o *CostRecommendationNotificationRuleCreateAttributes) GetIntervalExecutionTime() string {
	if o == nil || o.IntervalExecutionTime == nil {
		var ret string
		return ret
	}
	return *o.IntervalExecutionTime
}

// GetIntervalExecutionTimeOk returns a tuple with the IntervalExecutionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleCreateAttributes) GetIntervalExecutionTimeOk() (*string, bool) {
	if o == nil || o.IntervalExecutionTime == nil {
		return nil, false
	}
	return o.IntervalExecutionTime, true
}

// HasIntervalExecutionTime returns a boolean if a field has been set.
func (o *CostRecommendationNotificationRuleCreateAttributes) HasIntervalExecutionTime() bool {
	return o != nil && o.IntervalExecutionTime != nil
}

// SetIntervalExecutionTime gets a reference to the given string and assigns it to the IntervalExecutionTime field.
func (o *CostRecommendationNotificationRuleCreateAttributes) SetIntervalExecutionTime(v string) {
	o.IntervalExecutionTime = &v
}

// GetMentionTargets returns the MentionTargets field value if set, zero value otherwise.
func (o *CostRecommendationNotificationRuleCreateAttributes) GetMentionTargets() []string {
	if o == nil || o.MentionTargets == nil {
		var ret []string
		return ret
	}
	return o.MentionTargets
}

// GetMentionTargetsOk returns a tuple with the MentionTargets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleCreateAttributes) GetMentionTargetsOk() (*[]string, bool) {
	if o == nil || o.MentionTargets == nil {
		return nil, false
	}
	return &o.MentionTargets, true
}

// HasMentionTargets returns a boolean if a field has been set.
func (o *CostRecommendationNotificationRuleCreateAttributes) HasMentionTargets() bool {
	return o != nil && o.MentionTargets != nil
}

// SetMentionTargets gets a reference to the given []string and assigns it to the MentionTargets field.
func (o *CostRecommendationNotificationRuleCreateAttributes) SetMentionTargets(v []string) {
	o.MentionTargets = v
}

// GetMessageCustomization returns the MessageCustomization field value if set, zero value otherwise.
func (o *CostRecommendationNotificationRuleCreateAttributes) GetMessageCustomization() CostRecommendationNotificationRuleMessageCustomization {
	if o == nil || o.MessageCustomization == nil {
		var ret CostRecommendationNotificationRuleMessageCustomization
		return ret
	}
	return *o.MessageCustomization
}

// GetMessageCustomizationOk returns a tuple with the MessageCustomization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleCreateAttributes) GetMessageCustomizationOk() (*CostRecommendationNotificationRuleMessageCustomization, bool) {
	if o == nil || o.MessageCustomization == nil {
		return nil, false
	}
	return o.MessageCustomization, true
}

// HasMessageCustomization returns a boolean if a field has been set.
func (o *CostRecommendationNotificationRuleCreateAttributes) HasMessageCustomization() bool {
	return o != nil && o.MessageCustomization != nil
}

// SetMessageCustomization gets a reference to the given CostRecommendationNotificationRuleMessageCustomization and assigns it to the MessageCustomization field.
func (o *CostRecommendationNotificationRuleCreateAttributes) SetMessageCustomization(v CostRecommendationNotificationRuleMessageCustomization) {
	o.MessageCustomization = &v
}

// GetName returns the Name field value.
func (o *CostRecommendationNotificationRuleCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CostRecommendationNotificationRuleCreateAttributes) SetName(v string) {
	o.Name = v
}

// GetScopeQuery returns the ScopeQuery field value if set, zero value otherwise.
func (o *CostRecommendationNotificationRuleCreateAttributes) GetScopeQuery() string {
	if o == nil || o.ScopeQuery == nil {
		var ret string
		return ret
	}
	return *o.ScopeQuery
}

// GetScopeQueryOk returns a tuple with the ScopeQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleCreateAttributes) GetScopeQueryOk() (*string, bool) {
	if o == nil || o.ScopeQuery == nil {
		return nil, false
	}
	return o.ScopeQuery, true
}

// HasScopeQuery returns a boolean if a field has been set.
func (o *CostRecommendationNotificationRuleCreateAttributes) HasScopeQuery() bool {
	return o != nil && o.ScopeQuery != nil
}

// SetScopeQuery gets a reference to the given string and assigns it to the ScopeQuery field.
func (o *CostRecommendationNotificationRuleCreateAttributes) SetScopeQuery(v string) {
	o.ScopeQuery = &v
}

// GetTeamsChannelId returns the TeamsChannelId field value if set, zero value otherwise.
func (o *CostRecommendationNotificationRuleCreateAttributes) GetTeamsChannelId() string {
	if o == nil || o.TeamsChannelId == nil {
		var ret string
		return ret
	}
	return *o.TeamsChannelId
}

// GetTeamsChannelIdOk returns a tuple with the TeamsChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleCreateAttributes) GetTeamsChannelIdOk() (*string, bool) {
	if o == nil || o.TeamsChannelId == nil {
		return nil, false
	}
	return o.TeamsChannelId, true
}

// HasTeamsChannelId returns a boolean if a field has been set.
func (o *CostRecommendationNotificationRuleCreateAttributes) HasTeamsChannelId() bool {
	return o != nil && o.TeamsChannelId != nil
}

// SetTeamsChannelId gets a reference to the given string and assigns it to the TeamsChannelId field.
func (o *CostRecommendationNotificationRuleCreateAttributes) SetTeamsChannelId(v string) {
	o.TeamsChannelId = &v
}

// GetTeamsChannelName returns the TeamsChannelName field value if set, zero value otherwise.
func (o *CostRecommendationNotificationRuleCreateAttributes) GetTeamsChannelName() string {
	if o == nil || o.TeamsChannelName == nil {
		var ret string
		return ret
	}
	return *o.TeamsChannelName
}

// GetTeamsChannelNameOk returns a tuple with the TeamsChannelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleCreateAttributes) GetTeamsChannelNameOk() (*string, bool) {
	if o == nil || o.TeamsChannelName == nil {
		return nil, false
	}
	return o.TeamsChannelName, true
}

// HasTeamsChannelName returns a boolean if a field has been set.
func (o *CostRecommendationNotificationRuleCreateAttributes) HasTeamsChannelName() bool {
	return o != nil && o.TeamsChannelName != nil
}

// SetTeamsChannelName gets a reference to the given string and assigns it to the TeamsChannelName field.
func (o *CostRecommendationNotificationRuleCreateAttributes) SetTeamsChannelName(v string) {
	o.TeamsChannelName = &v
}

// GetTeamsTeamId returns the TeamsTeamId field value if set, zero value otherwise.
func (o *CostRecommendationNotificationRuleCreateAttributes) GetTeamsTeamId() string {
	if o == nil || o.TeamsTeamId == nil {
		var ret string
		return ret
	}
	return *o.TeamsTeamId
}

// GetTeamsTeamIdOk returns a tuple with the TeamsTeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleCreateAttributes) GetTeamsTeamIdOk() (*string, bool) {
	if o == nil || o.TeamsTeamId == nil {
		return nil, false
	}
	return o.TeamsTeamId, true
}

// HasTeamsTeamId returns a boolean if a field has been set.
func (o *CostRecommendationNotificationRuleCreateAttributes) HasTeamsTeamId() bool {
	return o != nil && o.TeamsTeamId != nil
}

// SetTeamsTeamId gets a reference to the given string and assigns it to the TeamsTeamId field.
func (o *CostRecommendationNotificationRuleCreateAttributes) SetTeamsTeamId(v string) {
	o.TeamsTeamId = &v
}

// GetTeamsTeamName returns the TeamsTeamName field value if set, zero value otherwise.
func (o *CostRecommendationNotificationRuleCreateAttributes) GetTeamsTeamName() string {
	if o == nil || o.TeamsTeamName == nil {
		var ret string
		return ret
	}
	return *o.TeamsTeamName
}

// GetTeamsTeamNameOk returns a tuple with the TeamsTeamName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleCreateAttributes) GetTeamsTeamNameOk() (*string, bool) {
	if o == nil || o.TeamsTeamName == nil {
		return nil, false
	}
	return o.TeamsTeamName, true
}

// HasTeamsTeamName returns a boolean if a field has been set.
func (o *CostRecommendationNotificationRuleCreateAttributes) HasTeamsTeamName() bool {
	return o != nil && o.TeamsTeamName != nil
}

// SetTeamsTeamName gets a reference to the given string and assigns it to the TeamsTeamName field.
func (o *CostRecommendationNotificationRuleCreateAttributes) SetTeamsTeamName(v string) {
	o.TeamsTeamName = &v
}

// GetTeamsTenantId returns the TeamsTenantId field value if set, zero value otherwise.
func (o *CostRecommendationNotificationRuleCreateAttributes) GetTeamsTenantId() string {
	if o == nil || o.TeamsTenantId == nil {
		var ret string
		return ret
	}
	return *o.TeamsTenantId
}

// GetTeamsTenantIdOk returns a tuple with the TeamsTenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleCreateAttributes) GetTeamsTenantIdOk() (*string, bool) {
	if o == nil || o.TeamsTenantId == nil {
		return nil, false
	}
	return o.TeamsTenantId, true
}

// HasTeamsTenantId returns a boolean if a field has been set.
func (o *CostRecommendationNotificationRuleCreateAttributes) HasTeamsTenantId() bool {
	return o != nil && o.TeamsTenantId != nil
}

// SetTeamsTenantId gets a reference to the given string and assigns it to the TeamsTenantId field.
func (o *CostRecommendationNotificationRuleCreateAttributes) SetTeamsTenantId(v string) {
	o.TeamsTenantId = &v
}

// GetTeamsTenantName returns the TeamsTenantName field value if set, zero value otherwise.
func (o *CostRecommendationNotificationRuleCreateAttributes) GetTeamsTenantName() string {
	if o == nil || o.TeamsTenantName == nil {
		var ret string
		return ret
	}
	return *o.TeamsTenantName
}

// GetTeamsTenantNameOk returns a tuple with the TeamsTenantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleCreateAttributes) GetTeamsTenantNameOk() (*string, bool) {
	if o == nil || o.TeamsTenantName == nil {
		return nil, false
	}
	return o.TeamsTenantName, true
}

// HasTeamsTenantName returns a boolean if a field has been set.
func (o *CostRecommendationNotificationRuleCreateAttributes) HasTeamsTenantName() bool {
	return o != nil && o.TeamsTenantName != nil
}

// SetTeamsTenantName gets a reference to the given string and assigns it to the TeamsTenantName field.
func (o *CostRecommendationNotificationRuleCreateAttributes) SetTeamsTenantName(v string) {
	o.TeamsTenantName = &v
}

// GetTimezone returns the Timezone field value.
func (o *CostRecommendationNotificationRuleCreateAttributes) GetTimezone() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleCreateAttributes) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timezone, true
}

// SetTimezone sets field value.
func (o *CostRecommendationNotificationRuleCreateAttributes) SetTimezone(v string) {
	o.Timezone = v
}

// GetWeeklyExecutionDay returns the WeeklyExecutionDay field value if set, zero value otherwise.
func (o *CostRecommendationNotificationRuleCreateAttributes) GetWeeklyExecutionDay() string {
	if o == nil || o.WeeklyExecutionDay == nil {
		var ret string
		return ret
	}
	return *o.WeeklyExecutionDay
}

// GetWeeklyExecutionDayOk returns a tuple with the WeeklyExecutionDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleCreateAttributes) GetWeeklyExecutionDayOk() (*string, bool) {
	if o == nil || o.WeeklyExecutionDay == nil {
		return nil, false
	}
	return o.WeeklyExecutionDay, true
}

// HasWeeklyExecutionDay returns a boolean if a field has been set.
func (o *CostRecommendationNotificationRuleCreateAttributes) HasWeeklyExecutionDay() bool {
	return o != nil && o.WeeklyExecutionDay != nil
}

// SetWeeklyExecutionDay gets a reference to the given string and assigns it to the WeeklyExecutionDay field.
func (o *CostRecommendationNotificationRuleCreateAttributes) SetWeeklyExecutionDay(v string) {
	o.WeeklyExecutionDay = &v
}

// GetWeeklyExecutionTime returns the WeeklyExecutionTime field value if set, zero value otherwise.
func (o *CostRecommendationNotificationRuleCreateAttributes) GetWeeklyExecutionTime() string {
	if o == nil || o.WeeklyExecutionTime == nil {
		var ret string
		return ret
	}
	return *o.WeeklyExecutionTime
}

// GetWeeklyExecutionTimeOk returns a tuple with the WeeklyExecutionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleCreateAttributes) GetWeeklyExecutionTimeOk() (*string, bool) {
	if o == nil || o.WeeklyExecutionTime == nil {
		return nil, false
	}
	return o.WeeklyExecutionTime, true
}

// HasWeeklyExecutionTime returns a boolean if a field has been set.
func (o *CostRecommendationNotificationRuleCreateAttributes) HasWeeklyExecutionTime() bool {
	return o != nil && o.WeeklyExecutionTime != nil
}

// SetWeeklyExecutionTime gets a reference to the given string and assigns it to the WeeklyExecutionTime field.
func (o *CostRecommendationNotificationRuleCreateAttributes) SetWeeklyExecutionTime(v string) {
	o.WeeklyExecutionTime = &v
}

// GetWorkspaceId returns the WorkspaceId field value if set, zero value otherwise.
func (o *CostRecommendationNotificationRuleCreateAttributes) GetWorkspaceId() string {
	if o == nil || o.WorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleCreateAttributes) GetWorkspaceIdOk() (*string, bool) {
	if o == nil || o.WorkspaceId == nil {
		return nil, false
	}
	return o.WorkspaceId, true
}

// HasWorkspaceId returns a boolean if a field has been set.
func (o *CostRecommendationNotificationRuleCreateAttributes) HasWorkspaceId() bool {
	return o != nil && o.WorkspaceId != nil
}

// SetWorkspaceId gets a reference to the given string and assigns it to the WorkspaceId field.
func (o *CostRecommendationNotificationRuleCreateAttributes) SetWorkspaceId(v string) {
	o.WorkspaceId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CostRecommendationNotificationRuleCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ChannelId != nil {
		toSerialize["channel_id"] = o.ChannelId
	}
	if o.DestinationType != nil {
		toSerialize["destination_type"] = o.DestinationType
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["frequency"] = o.Frequency
	if o.IntervalExecutionDate != nil {
		toSerialize["interval_execution_date"] = o.IntervalExecutionDate
	}
	if o.IntervalExecutionTime != nil {
		toSerialize["interval_execution_time"] = o.IntervalExecutionTime
	}
	if o.MentionTargets != nil {
		toSerialize["mention_targets"] = o.MentionTargets
	}
	if o.MessageCustomization != nil {
		toSerialize["message_customization"] = o.MessageCustomization
	}
	toSerialize["name"] = o.Name
	if o.ScopeQuery != nil {
		toSerialize["scope_query"] = o.ScopeQuery
	}
	if o.TeamsChannelId != nil {
		toSerialize["teams_channel_id"] = o.TeamsChannelId
	}
	if o.TeamsChannelName != nil {
		toSerialize["teams_channel_name"] = o.TeamsChannelName
	}
	if o.TeamsTeamId != nil {
		toSerialize["teams_team_id"] = o.TeamsTeamId
	}
	if o.TeamsTeamName != nil {
		toSerialize["teams_team_name"] = o.TeamsTeamName
	}
	if o.TeamsTenantId != nil {
		toSerialize["teams_tenant_id"] = o.TeamsTenantId
	}
	if o.TeamsTenantName != nil {
		toSerialize["teams_tenant_name"] = o.TeamsTenantName
	}
	toSerialize["timezone"] = o.Timezone
	if o.WeeklyExecutionDay != nil {
		toSerialize["weekly_execution_day"] = o.WeeklyExecutionDay
	}
	if o.WeeklyExecutionTime != nil {
		toSerialize["weekly_execution_time"] = o.WeeklyExecutionTime
	}
	if o.WorkspaceId != nil {
		toSerialize["workspace_id"] = o.WorkspaceId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CostRecommendationNotificationRuleCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ChannelId             *string                                                 `json:"channel_id,omitempty"`
		DestinationType       *CostRecommendationNotificationRuleDestinationType      `json:"destination_type,omitempty"`
		Enabled               *bool                                                   `json:"enabled,omitempty"`
		Frequency             *CostRecommendationNotificationRuleFrequency            `json:"frequency"`
		IntervalExecutionDate *string                                                 `json:"interval_execution_date,omitempty"`
		IntervalExecutionTime *string                                                 `json:"interval_execution_time,omitempty"`
		MentionTargets        []string                                                `json:"mention_targets,omitempty"`
		MessageCustomization  *CostRecommendationNotificationRuleMessageCustomization `json:"message_customization,omitempty"`
		Name                  *string                                                 `json:"name"`
		ScopeQuery            *string                                                 `json:"scope_query,omitempty"`
		TeamsChannelId        *string                                                 `json:"teams_channel_id,omitempty"`
		TeamsChannelName      *string                                                 `json:"teams_channel_name,omitempty"`
		TeamsTeamId           *string                                                 `json:"teams_team_id,omitempty"`
		TeamsTeamName         *string                                                 `json:"teams_team_name,omitempty"`
		TeamsTenantId         *string                                                 `json:"teams_tenant_id,omitempty"`
		TeamsTenantName       *string                                                 `json:"teams_tenant_name,omitempty"`
		Timezone              *string                                                 `json:"timezone"`
		WeeklyExecutionDay    *string                                                 `json:"weekly_execution_day,omitempty"`
		WeeklyExecutionTime   *string                                                 `json:"weekly_execution_time,omitempty"`
		WorkspaceId           *string                                                 `json:"workspace_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Frequency == nil {
		return fmt.Errorf("required field frequency missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Timezone == nil {
		return fmt.Errorf("required field timezone missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"channel_id", "destination_type", "enabled", "frequency", "interval_execution_date", "interval_execution_time", "mention_targets", "message_customization", "name", "scope_query", "teams_channel_id", "teams_channel_name", "teams_team_id", "teams_team_name", "teams_tenant_id", "teams_tenant_name", "timezone", "weekly_execution_day", "weekly_execution_time", "workspace_id"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ChannelId = all.ChannelId
	if all.DestinationType != nil && !all.DestinationType.IsValid() {
		hasInvalidField = true
	} else {
		o.DestinationType = all.DestinationType
	}
	o.Enabled = all.Enabled
	if !all.Frequency.IsValid() {
		hasInvalidField = true
	} else {
		o.Frequency = *all.Frequency
	}
	o.IntervalExecutionDate = all.IntervalExecutionDate
	o.IntervalExecutionTime = all.IntervalExecutionTime
	o.MentionTargets = all.MentionTargets
	if all.MessageCustomization != nil && all.MessageCustomization.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.MessageCustomization = all.MessageCustomization
	o.Name = *all.Name
	o.ScopeQuery = all.ScopeQuery
	o.TeamsChannelId = all.TeamsChannelId
	o.TeamsChannelName = all.TeamsChannelName
	o.TeamsTeamId = all.TeamsTeamId
	o.TeamsTeamName = all.TeamsTeamName
	o.TeamsTenantId = all.TeamsTenantId
	o.TeamsTenantName = all.TeamsTenantName
	o.Timezone = *all.Timezone
	o.WeeklyExecutionDay = all.WeeklyExecutionDay
	o.WeeklyExecutionTime = all.WeeklyExecutionTime
	o.WorkspaceId = all.WorkspaceId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
