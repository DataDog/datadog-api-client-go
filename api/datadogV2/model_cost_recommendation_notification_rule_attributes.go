// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostRecommendationNotificationRuleAttributes Attributes of a recommendation notification rule.
type CostRecommendationNotificationRuleAttributes struct {
	// The Slack channel ID. Empty for `teams` rules.
	ChannelId string `json:"channel_id"`
	// When the rule was created.
	CreatedAt time.Time `json:"created_at"`
	// The UUID of the user who created the rule.
	CreatedBy string `json:"created_by"`
	// Where the digest is sent. `slack` requires `workspace_id` and `channel_id`; `teams` requires the
	// `teams_*` attributes. Defaults to `slack`.
	DestinationType CostRecommendationNotificationRuleDestinationType `json:"destination_type"`
	// Whether the digest is sent.
	Enabled bool `json:"enabled"`
	// How often the digest is sent. `weekly` and `biweekly` use `weekly_execution_day` and `weekly_execution_time`;
	// `every_30` and `every_90` (days) use `interval_execution_date` and `interval_execution_time`.
	Frequency CostRecommendationNotificationRuleFrequency `json:"frequency"`
	// The first date (`YYYY-MM-DD`) of an `every_30` or `every_90` schedule.
	IntervalExecutionDate datadog.NullableString `json:"interval_execution_date"`
	// The time of day (`HH:MM`, in `timezone`) of an `every_30` or `every_90` schedule.
	IntervalExecutionTime datadog.NullableString `json:"interval_execution_time"`
	// Slack member IDs to mention in the digest. Only valid for `slack` rules.
	MentionTargets []string `json:"mention_targets"`
	// Which sections the digest shows. Sent as a whole: sections left out of a request use their defaults, which
	// show everything.
	MessageCustomization CostRecommendationNotificationRuleMessageCustomization `json:"message_customization"`
	// When the rule was last modified.
	ModifiedAt time.Time `json:"modified_at"`
	// The UUID of the user who last modified the rule.
	ModifiedBy string `json:"modified_by"`
	// The name of the rule, unique within the organization.
	Name string `json:"name"`
	// The ID of the organization that owns the rule.
	OrgId int64 `json:"org_id"`
	// The recommendation search query that selects what the digest includes. Empty selects every recommendation.
	ScopeQuery string `json:"scope_query"`
	// The Microsoft Teams channel ID. Empty for `slack` rules.
	TeamsChannelId string `json:"teams_channel_id"`
	// The Microsoft Teams channel name. Empty for `slack` rules.
	TeamsChannelName string `json:"teams_channel_name"`
	// The Microsoft Teams team ID. Empty for `slack` rules.
	TeamsTeamId string `json:"teams_team_id"`
	// The Microsoft Teams team name. Empty for `slack` rules.
	TeamsTeamName string `json:"teams_team_name"`
	// The Microsoft Teams tenant ID. Empty for `slack` rules.
	TeamsTenantId string `json:"teams_tenant_id"`
	// The Microsoft Teams tenant name. Empty for `slack` rules.
	TeamsTenantName string `json:"teams_tenant_name"`
	// The IANA time zone of the schedule.
	Timezone string `json:"timezone"`
	// The day of the week of a `weekly` or `biweekly` schedule, such as `monday`.
	WeeklyExecutionDay datadog.NullableString `json:"weekly_execution_day"`
	// The time of day (`HH:MM`, in `timezone`) of a `weekly` or `biweekly` schedule.
	WeeklyExecutionTime datadog.NullableString `json:"weekly_execution_time"`
	// The Slack workspace ID. Empty for `teams` rules.
	WorkspaceId string `json:"workspace_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCostRecommendationNotificationRuleAttributes instantiates a new CostRecommendationNotificationRuleAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCostRecommendationNotificationRuleAttributes(channelId string, createdAt time.Time, createdBy string, destinationType CostRecommendationNotificationRuleDestinationType, enabled bool, frequency CostRecommendationNotificationRuleFrequency, intervalExecutionDate datadog.NullableString, intervalExecutionTime datadog.NullableString, mentionTargets []string, messageCustomization CostRecommendationNotificationRuleMessageCustomization, modifiedAt time.Time, modifiedBy string, name string, orgId int64, scopeQuery string, teamsChannelId string, teamsChannelName string, teamsTeamId string, teamsTeamName string, teamsTenantId string, teamsTenantName string, timezone string, weeklyExecutionDay datadog.NullableString, weeklyExecutionTime datadog.NullableString, workspaceId string) *CostRecommendationNotificationRuleAttributes {
	this := CostRecommendationNotificationRuleAttributes{}
	this.ChannelId = channelId
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.DestinationType = destinationType
	this.Enabled = enabled
	this.Frequency = frequency
	this.IntervalExecutionDate = intervalExecutionDate
	this.IntervalExecutionTime = intervalExecutionTime
	this.MentionTargets = mentionTargets
	this.MessageCustomization = messageCustomization
	this.ModifiedAt = modifiedAt
	this.ModifiedBy = modifiedBy
	this.Name = name
	this.OrgId = orgId
	this.ScopeQuery = scopeQuery
	this.TeamsChannelId = teamsChannelId
	this.TeamsChannelName = teamsChannelName
	this.TeamsTeamId = teamsTeamId
	this.TeamsTeamName = teamsTeamName
	this.TeamsTenantId = teamsTenantId
	this.TeamsTenantName = teamsTenantName
	this.Timezone = timezone
	this.WeeklyExecutionDay = weeklyExecutionDay
	this.WeeklyExecutionTime = weeklyExecutionTime
	this.WorkspaceId = workspaceId
	return &this
}

// NewCostRecommendationNotificationRuleAttributesWithDefaults instantiates a new CostRecommendationNotificationRuleAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCostRecommendationNotificationRuleAttributesWithDefaults() *CostRecommendationNotificationRuleAttributes {
	this := CostRecommendationNotificationRuleAttributes{}
	return &this
}

// GetChannelId returns the ChannelId field value.
func (o *CostRecommendationNotificationRuleAttributes) GetChannelId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleAttributes) GetChannelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelId, true
}

// SetChannelId sets field value.
func (o *CostRecommendationNotificationRuleAttributes) SetChannelId(v string) {
	o.ChannelId = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *CostRecommendationNotificationRuleAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *CostRecommendationNotificationRuleAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value.
func (o *CostRecommendationNotificationRuleAttributes) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleAttributes) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value.
func (o *CostRecommendationNotificationRuleAttributes) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetDestinationType returns the DestinationType field value.
func (o *CostRecommendationNotificationRuleAttributes) GetDestinationType() CostRecommendationNotificationRuleDestinationType {
	if o == nil {
		var ret CostRecommendationNotificationRuleDestinationType
		return ret
	}
	return o.DestinationType
}

// GetDestinationTypeOk returns a tuple with the DestinationType field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleAttributes) GetDestinationTypeOk() (*CostRecommendationNotificationRuleDestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationType, true
}

// SetDestinationType sets field value.
func (o *CostRecommendationNotificationRuleAttributes) SetDestinationType(v CostRecommendationNotificationRuleDestinationType) {
	o.DestinationType = v
}

// GetEnabled returns the Enabled field value.
func (o *CostRecommendationNotificationRuleAttributes) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *CostRecommendationNotificationRuleAttributes) SetEnabled(v bool) {
	o.Enabled = v
}

// GetFrequency returns the Frequency field value.
func (o *CostRecommendationNotificationRuleAttributes) GetFrequency() CostRecommendationNotificationRuleFrequency {
	if o == nil {
		var ret CostRecommendationNotificationRuleFrequency
		return ret
	}
	return o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleAttributes) GetFrequencyOk() (*CostRecommendationNotificationRuleFrequency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Frequency, true
}

// SetFrequency sets field value.
func (o *CostRecommendationNotificationRuleAttributes) SetFrequency(v CostRecommendationNotificationRuleFrequency) {
	o.Frequency = v
}

// GetIntervalExecutionDate returns the IntervalExecutionDate field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *CostRecommendationNotificationRuleAttributes) GetIntervalExecutionDate() string {
	if o == nil || o.IntervalExecutionDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.IntervalExecutionDate.Get()
}

// GetIntervalExecutionDateOk returns a tuple with the IntervalExecutionDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CostRecommendationNotificationRuleAttributes) GetIntervalExecutionDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IntervalExecutionDate.Get(), o.IntervalExecutionDate.IsSet()
}

// SetIntervalExecutionDate sets field value.
func (o *CostRecommendationNotificationRuleAttributes) SetIntervalExecutionDate(v string) {
	o.IntervalExecutionDate.Set(&v)
}

// GetIntervalExecutionTime returns the IntervalExecutionTime field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *CostRecommendationNotificationRuleAttributes) GetIntervalExecutionTime() string {
	if o == nil || o.IntervalExecutionTime.Get() == nil {
		var ret string
		return ret
	}
	return *o.IntervalExecutionTime.Get()
}

// GetIntervalExecutionTimeOk returns a tuple with the IntervalExecutionTime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CostRecommendationNotificationRuleAttributes) GetIntervalExecutionTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IntervalExecutionTime.Get(), o.IntervalExecutionTime.IsSet()
}

// SetIntervalExecutionTime sets field value.
func (o *CostRecommendationNotificationRuleAttributes) SetIntervalExecutionTime(v string) {
	o.IntervalExecutionTime.Set(&v)
}

// GetMentionTargets returns the MentionTargets field value.
func (o *CostRecommendationNotificationRuleAttributes) GetMentionTargets() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.MentionTargets
}

// GetMentionTargetsOk returns a tuple with the MentionTargets field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleAttributes) GetMentionTargetsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MentionTargets, true
}

// SetMentionTargets sets field value.
func (o *CostRecommendationNotificationRuleAttributes) SetMentionTargets(v []string) {
	o.MentionTargets = v
}

// GetMessageCustomization returns the MessageCustomization field value.
func (o *CostRecommendationNotificationRuleAttributes) GetMessageCustomization() CostRecommendationNotificationRuleMessageCustomization {
	if o == nil {
		var ret CostRecommendationNotificationRuleMessageCustomization
		return ret
	}
	return o.MessageCustomization
}

// GetMessageCustomizationOk returns a tuple with the MessageCustomization field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleAttributes) GetMessageCustomizationOk() (*CostRecommendationNotificationRuleMessageCustomization, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageCustomization, true
}

// SetMessageCustomization sets field value.
func (o *CostRecommendationNotificationRuleAttributes) SetMessageCustomization(v CostRecommendationNotificationRuleMessageCustomization) {
	o.MessageCustomization = v
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *CostRecommendationNotificationRuleAttributes) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *CostRecommendationNotificationRuleAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetModifiedBy returns the ModifiedBy field value.
func (o *CostRecommendationNotificationRuleAttributes) GetModifiedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleAttributes) GetModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedBy, true
}

// SetModifiedBy sets field value.
func (o *CostRecommendationNotificationRuleAttributes) SetModifiedBy(v string) {
	o.ModifiedBy = v
}

// GetName returns the Name field value.
func (o *CostRecommendationNotificationRuleAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CostRecommendationNotificationRuleAttributes) SetName(v string) {
	o.Name = v
}

// GetOrgId returns the OrgId field value.
func (o *CostRecommendationNotificationRuleAttributes) GetOrgId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleAttributes) GetOrgIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value.
func (o *CostRecommendationNotificationRuleAttributes) SetOrgId(v int64) {
	o.OrgId = v
}

// GetScopeQuery returns the ScopeQuery field value.
func (o *CostRecommendationNotificationRuleAttributes) GetScopeQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ScopeQuery
}

// GetScopeQueryOk returns a tuple with the ScopeQuery field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleAttributes) GetScopeQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScopeQuery, true
}

// SetScopeQuery sets field value.
func (o *CostRecommendationNotificationRuleAttributes) SetScopeQuery(v string) {
	o.ScopeQuery = v
}

// GetTeamsChannelId returns the TeamsChannelId field value.
func (o *CostRecommendationNotificationRuleAttributes) GetTeamsChannelId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TeamsChannelId
}

// GetTeamsChannelIdOk returns a tuple with the TeamsChannelId field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleAttributes) GetTeamsChannelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TeamsChannelId, true
}

// SetTeamsChannelId sets field value.
func (o *CostRecommendationNotificationRuleAttributes) SetTeamsChannelId(v string) {
	o.TeamsChannelId = v
}

// GetTeamsChannelName returns the TeamsChannelName field value.
func (o *CostRecommendationNotificationRuleAttributes) GetTeamsChannelName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TeamsChannelName
}

// GetTeamsChannelNameOk returns a tuple with the TeamsChannelName field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleAttributes) GetTeamsChannelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TeamsChannelName, true
}

// SetTeamsChannelName sets field value.
func (o *CostRecommendationNotificationRuleAttributes) SetTeamsChannelName(v string) {
	o.TeamsChannelName = v
}

// GetTeamsTeamId returns the TeamsTeamId field value.
func (o *CostRecommendationNotificationRuleAttributes) GetTeamsTeamId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TeamsTeamId
}

// GetTeamsTeamIdOk returns a tuple with the TeamsTeamId field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleAttributes) GetTeamsTeamIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TeamsTeamId, true
}

// SetTeamsTeamId sets field value.
func (o *CostRecommendationNotificationRuleAttributes) SetTeamsTeamId(v string) {
	o.TeamsTeamId = v
}

// GetTeamsTeamName returns the TeamsTeamName field value.
func (o *CostRecommendationNotificationRuleAttributes) GetTeamsTeamName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TeamsTeamName
}

// GetTeamsTeamNameOk returns a tuple with the TeamsTeamName field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleAttributes) GetTeamsTeamNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TeamsTeamName, true
}

// SetTeamsTeamName sets field value.
func (o *CostRecommendationNotificationRuleAttributes) SetTeamsTeamName(v string) {
	o.TeamsTeamName = v
}

// GetTeamsTenantId returns the TeamsTenantId field value.
func (o *CostRecommendationNotificationRuleAttributes) GetTeamsTenantId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TeamsTenantId
}

// GetTeamsTenantIdOk returns a tuple with the TeamsTenantId field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleAttributes) GetTeamsTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TeamsTenantId, true
}

// SetTeamsTenantId sets field value.
func (o *CostRecommendationNotificationRuleAttributes) SetTeamsTenantId(v string) {
	o.TeamsTenantId = v
}

// GetTeamsTenantName returns the TeamsTenantName field value.
func (o *CostRecommendationNotificationRuleAttributes) GetTeamsTenantName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TeamsTenantName
}

// GetTeamsTenantNameOk returns a tuple with the TeamsTenantName field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleAttributes) GetTeamsTenantNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TeamsTenantName, true
}

// SetTeamsTenantName sets field value.
func (o *CostRecommendationNotificationRuleAttributes) SetTeamsTenantName(v string) {
	o.TeamsTenantName = v
}

// GetTimezone returns the Timezone field value.
func (o *CostRecommendationNotificationRuleAttributes) GetTimezone() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleAttributes) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timezone, true
}

// SetTimezone sets field value.
func (o *CostRecommendationNotificationRuleAttributes) SetTimezone(v string) {
	o.Timezone = v
}

// GetWeeklyExecutionDay returns the WeeklyExecutionDay field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *CostRecommendationNotificationRuleAttributes) GetWeeklyExecutionDay() string {
	if o == nil || o.WeeklyExecutionDay.Get() == nil {
		var ret string
		return ret
	}
	return *o.WeeklyExecutionDay.Get()
}

// GetWeeklyExecutionDayOk returns a tuple with the WeeklyExecutionDay field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CostRecommendationNotificationRuleAttributes) GetWeeklyExecutionDayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WeeklyExecutionDay.Get(), o.WeeklyExecutionDay.IsSet()
}

// SetWeeklyExecutionDay sets field value.
func (o *CostRecommendationNotificationRuleAttributes) SetWeeklyExecutionDay(v string) {
	o.WeeklyExecutionDay.Set(&v)
}

// GetWeeklyExecutionTime returns the WeeklyExecutionTime field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *CostRecommendationNotificationRuleAttributes) GetWeeklyExecutionTime() string {
	if o == nil || o.WeeklyExecutionTime.Get() == nil {
		var ret string
		return ret
	}
	return *o.WeeklyExecutionTime.Get()
}

// GetWeeklyExecutionTimeOk returns a tuple with the WeeklyExecutionTime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CostRecommendationNotificationRuleAttributes) GetWeeklyExecutionTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WeeklyExecutionTime.Get(), o.WeeklyExecutionTime.IsSet()
}

// SetWeeklyExecutionTime sets field value.
func (o *CostRecommendationNotificationRuleAttributes) SetWeeklyExecutionTime(v string) {
	o.WeeklyExecutionTime.Set(&v)
}

// GetWorkspaceId returns the WorkspaceId field value.
func (o *CostRecommendationNotificationRuleAttributes) GetWorkspaceId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleAttributes) GetWorkspaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkspaceId, true
}

// SetWorkspaceId sets field value.
func (o *CostRecommendationNotificationRuleAttributes) SetWorkspaceId(v string) {
	o.WorkspaceId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CostRecommendationNotificationRuleAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["channel_id"] = o.ChannelId
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["created_by"] = o.CreatedBy
	toSerialize["destination_type"] = o.DestinationType
	toSerialize["enabled"] = o.Enabled
	toSerialize["frequency"] = o.Frequency
	toSerialize["interval_execution_date"] = o.IntervalExecutionDate.Get()
	toSerialize["interval_execution_time"] = o.IntervalExecutionTime.Get()
	toSerialize["mention_targets"] = o.MentionTargets
	toSerialize["message_customization"] = o.MessageCustomization
	if o.ModifiedAt.Nanosecond() == 0 {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["modified_by"] = o.ModifiedBy
	toSerialize["name"] = o.Name
	toSerialize["org_id"] = o.OrgId
	toSerialize["scope_query"] = o.ScopeQuery
	toSerialize["teams_channel_id"] = o.TeamsChannelId
	toSerialize["teams_channel_name"] = o.TeamsChannelName
	toSerialize["teams_team_id"] = o.TeamsTeamId
	toSerialize["teams_team_name"] = o.TeamsTeamName
	toSerialize["teams_tenant_id"] = o.TeamsTenantId
	toSerialize["teams_tenant_name"] = o.TeamsTenantName
	toSerialize["timezone"] = o.Timezone
	toSerialize["weekly_execution_day"] = o.WeeklyExecutionDay.Get()
	toSerialize["weekly_execution_time"] = o.WeeklyExecutionTime.Get()
	toSerialize["workspace_id"] = o.WorkspaceId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CostRecommendationNotificationRuleAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ChannelId             *string                                                 `json:"channel_id"`
		CreatedAt             *time.Time                                              `json:"created_at"`
		CreatedBy             *string                                                 `json:"created_by"`
		DestinationType       *CostRecommendationNotificationRuleDestinationType      `json:"destination_type"`
		Enabled               *bool                                                   `json:"enabled"`
		Frequency             *CostRecommendationNotificationRuleFrequency            `json:"frequency"`
		IntervalExecutionDate datadog.NullableString                                  `json:"interval_execution_date"`
		IntervalExecutionTime datadog.NullableString                                  `json:"interval_execution_time"`
		MentionTargets        *[]string                                               `json:"mention_targets"`
		MessageCustomization  *CostRecommendationNotificationRuleMessageCustomization `json:"message_customization"`
		ModifiedAt            *time.Time                                              `json:"modified_at"`
		ModifiedBy            *string                                                 `json:"modified_by"`
		Name                  *string                                                 `json:"name"`
		OrgId                 *int64                                                  `json:"org_id"`
		ScopeQuery            *string                                                 `json:"scope_query"`
		TeamsChannelId        *string                                                 `json:"teams_channel_id"`
		TeamsChannelName      *string                                                 `json:"teams_channel_name"`
		TeamsTeamId           *string                                                 `json:"teams_team_id"`
		TeamsTeamName         *string                                                 `json:"teams_team_name"`
		TeamsTenantId         *string                                                 `json:"teams_tenant_id"`
		TeamsTenantName       *string                                                 `json:"teams_tenant_name"`
		Timezone              *string                                                 `json:"timezone"`
		WeeklyExecutionDay    datadog.NullableString                                  `json:"weekly_execution_day"`
		WeeklyExecutionTime   datadog.NullableString                                  `json:"weekly_execution_time"`
		WorkspaceId           *string                                                 `json:"workspace_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ChannelId == nil {
		return fmt.Errorf("required field channel_id missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.CreatedBy == nil {
		return fmt.Errorf("required field created_by missing")
	}
	if all.DestinationType == nil {
		return fmt.Errorf("required field destination_type missing")
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.Frequency == nil {
		return fmt.Errorf("required field frequency missing")
	}
	if !all.IntervalExecutionDate.IsSet() {
		return fmt.Errorf("required field interval_execution_date missing")
	}
	if !all.IntervalExecutionTime.IsSet() {
		return fmt.Errorf("required field interval_execution_time missing")
	}
	if all.MentionTargets == nil {
		return fmt.Errorf("required field mention_targets missing")
	}
	if all.MessageCustomization == nil {
		return fmt.Errorf("required field message_customization missing")
	}
	if all.ModifiedAt == nil {
		return fmt.Errorf("required field modified_at missing")
	}
	if all.ModifiedBy == nil {
		return fmt.Errorf("required field modified_by missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.OrgId == nil {
		return fmt.Errorf("required field org_id missing")
	}
	if all.ScopeQuery == nil {
		return fmt.Errorf("required field scope_query missing")
	}
	if all.TeamsChannelId == nil {
		return fmt.Errorf("required field teams_channel_id missing")
	}
	if all.TeamsChannelName == nil {
		return fmt.Errorf("required field teams_channel_name missing")
	}
	if all.TeamsTeamId == nil {
		return fmt.Errorf("required field teams_team_id missing")
	}
	if all.TeamsTeamName == nil {
		return fmt.Errorf("required field teams_team_name missing")
	}
	if all.TeamsTenantId == nil {
		return fmt.Errorf("required field teams_tenant_id missing")
	}
	if all.TeamsTenantName == nil {
		return fmt.Errorf("required field teams_tenant_name missing")
	}
	if all.Timezone == nil {
		return fmt.Errorf("required field timezone missing")
	}
	if !all.WeeklyExecutionDay.IsSet() {
		return fmt.Errorf("required field weekly_execution_day missing")
	}
	if !all.WeeklyExecutionTime.IsSet() {
		return fmt.Errorf("required field weekly_execution_time missing")
	}
	if all.WorkspaceId == nil {
		return fmt.Errorf("required field workspace_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"channel_id", "created_at", "created_by", "destination_type", "enabled", "frequency", "interval_execution_date", "interval_execution_time", "mention_targets", "message_customization", "modified_at", "modified_by", "name", "org_id", "scope_query", "teams_channel_id", "teams_channel_name", "teams_team_id", "teams_team_name", "teams_tenant_id", "teams_tenant_name", "timezone", "weekly_execution_day", "weekly_execution_time", "workspace_id"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ChannelId = *all.ChannelId
	o.CreatedAt = *all.CreatedAt
	o.CreatedBy = *all.CreatedBy
	if !all.DestinationType.IsValid() {
		hasInvalidField = true
	} else {
		o.DestinationType = *all.DestinationType
	}
	o.Enabled = *all.Enabled
	if !all.Frequency.IsValid() {
		hasInvalidField = true
	} else {
		o.Frequency = *all.Frequency
	}
	o.IntervalExecutionDate = all.IntervalExecutionDate
	o.IntervalExecutionTime = all.IntervalExecutionTime
	o.MentionTargets = *all.MentionTargets
	if all.MessageCustomization.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.MessageCustomization = *all.MessageCustomization
	o.ModifiedAt = *all.ModifiedAt
	o.ModifiedBy = *all.ModifiedBy
	o.Name = *all.Name
	o.OrgId = *all.OrgId
	o.ScopeQuery = *all.ScopeQuery
	o.TeamsChannelId = *all.TeamsChannelId
	o.TeamsChannelName = *all.TeamsChannelName
	o.TeamsTeamId = *all.TeamsTeamId
	o.TeamsTeamName = *all.TeamsTeamName
	o.TeamsTenantId = *all.TeamsTenantId
	o.TeamsTenantName = *all.TeamsTenantName
	o.Timezone = *all.Timezone
	o.WeeklyExecutionDay = all.WeeklyExecutionDay
	o.WeeklyExecutionTime = all.WeeklyExecutionTime
	o.WorkspaceId = *all.WorkspaceId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
