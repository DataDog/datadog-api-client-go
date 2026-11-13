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

// AiGatewayRateLimitAttributes Attributes of an AI Gateway rate limit.
type AiGatewayRateLimitAttributes struct {
	// Maximum priced spend per rolling minute, in microcents of 1e-8 USD each, shared by all matching traffic, or per counter for a template rate limit. Null when this rule does not limit cost.
	CostMicrocentsPerMinute datadog.NullableInt64 `json:"cost_microcents_per_minute,omitempty"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at"`
	// Email of the user who created this rate limit.
	CreatedBy string `json:"created_by"`
	// Whether this rate limit is active.
	Enabled bool `json:"enabled"`
	// Maximum input tokens per rolling minute shared by all matching traffic, or per counter for a template rate limit. Null when this rule does not limit input tokens.
	InputTokensPerMinute datadog.NullableInt64 `json:"input_tokens_per_minute,omitempty"`
	// Whether the selector gives each matching value its own allowance, that is, whether it has at least one per_value entry. Returned even when scope_restricted hides the selector.
	IsTemplate bool `json:"is_template"`
	// Action taken when request rate exceeds the limit.
	Mode AiGatewayRateLimitMode `json:"mode"`
	// Datadog monitor that alerts on this rule. Null until one is provisioned.
	MonitorId datadog.NullableInt64 `json:"monitor_id,omitempty"`
	// Rate limit name.
	Name string `json:"name"`
	// Maximum output tokens per rolling minute shared by all matching traffic, or per counter for a template rate limit. Null when this rule does not limit output tokens.
	OutputTokensPerMinute datadog.NullableInt64 `json:"output_tokens_per_minute,omitempty"`
	// Maximum requests per rolling minute shared by all matching traffic, or per counter for a template rate limit. Null when this rule does not limit requests.
	RequestsPerMinute datadog.NullableInt64 `json:"requests_per_minute,omitempty"`
	// Increases on every change to the definition.
	Revision int64 `json:"revision"`
	// Whether provider or model scope metadata was withheld. Omit unchanged scope fields on an unrelated update to preserve their saved values.
	ScopeRestricted bool `json:"scope_restricted"`
	// Dimension rules for a template rate limit, which gives each matched user, Team, or value its own allowance with the rule's limits. Can't be combined with user_id, team_id, or tags, and can't be changed after creation. Omit for a rule with one shared allowance.
	Selector []AiGatewayRateLimitSelectorEntry `json:"selector,omitempty"`
	// Optional scope tags, such as provider, model, ml_app, or custom dimensions. A request matches when it has all of these tags.
	Tags map[string]string `json:"tags"`
	// Optional Team UUID. Count only requests attributed to this Datadog Team.
	TeamId datadog.NullableUUID `json:"team_id,omitempty"`
	// Dimensions of the selector's per_value entries, such as user_id or tag:env, each matching value of which gets its own allowance. Empty unless is_template is true. Returned even when scope_restricted hides the selector.
	TemplateSplitDims []string `json:"template_split_dims"`
	// Last modification timestamp.
	UpdatedAt time.Time `json:"updated_at"`
	// Email of the user who last modified this rate limit.
	UpdatedBy string `json:"updated_by"`
	// Optional user UUID. Count only this user's requests.
	UserId datadog.NullableUUID `json:"user_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayRateLimitAttributes instantiates a new AiGatewayRateLimitAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayRateLimitAttributes(createdAt time.Time, createdBy string, enabled bool, isTemplate bool, mode AiGatewayRateLimitMode, name string, revision int64, scopeRestricted bool, tags map[string]string, templateSplitDims []string, updatedAt time.Time, updatedBy string) *AiGatewayRateLimitAttributes {
	this := AiGatewayRateLimitAttributes{}
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.Enabled = enabled
	this.IsTemplate = isTemplate
	this.Mode = mode
	this.Name = name
	this.Revision = revision
	this.ScopeRestricted = scopeRestricted
	this.Tags = tags
	this.TemplateSplitDims = templateSplitDims
	this.UpdatedAt = updatedAt
	this.UpdatedBy = updatedBy
	return &this
}

// NewAiGatewayRateLimitAttributesWithDefaults instantiates a new AiGatewayRateLimitAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayRateLimitAttributesWithDefaults() *AiGatewayRateLimitAttributes {
	this := AiGatewayRateLimitAttributes{}
	return &this
}

// GetCostMicrocentsPerMinute returns the CostMicrocentsPerMinute field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiGatewayRateLimitAttributes) GetCostMicrocentsPerMinute() int64 {
	if o == nil || o.CostMicrocentsPerMinute.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CostMicrocentsPerMinute.Get()
}

// GetCostMicrocentsPerMinuteOk returns a tuple with the CostMicrocentsPerMinute field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AiGatewayRateLimitAttributes) GetCostMicrocentsPerMinuteOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CostMicrocentsPerMinute.Get(), o.CostMicrocentsPerMinute.IsSet()
}

// HasCostMicrocentsPerMinute returns a boolean if a field has been set.
func (o *AiGatewayRateLimitAttributes) HasCostMicrocentsPerMinute() bool {
	return o != nil && o.CostMicrocentsPerMinute.IsSet()
}

// SetCostMicrocentsPerMinute gets a reference to the given datadog.NullableInt64 and assigns it to the CostMicrocentsPerMinute field.
func (o *AiGatewayRateLimitAttributes) SetCostMicrocentsPerMinute(v int64) {
	o.CostMicrocentsPerMinute.Set(&v)
}

// SetCostMicrocentsPerMinuteNil sets the value for CostMicrocentsPerMinute to be an explicit nil.
func (o *AiGatewayRateLimitAttributes) SetCostMicrocentsPerMinuteNil() {
	o.CostMicrocentsPerMinute.Set(nil)
}

// UnsetCostMicrocentsPerMinute ensures that no value is present for CostMicrocentsPerMinute, not even an explicit nil.
func (o *AiGatewayRateLimitAttributes) UnsetCostMicrocentsPerMinute() {
	o.CostMicrocentsPerMinute.Unset()
}

// GetCreatedAt returns the CreatedAt field value.
func (o *AiGatewayRateLimitAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *AiGatewayRateLimitAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value.
func (o *AiGatewayRateLimitAttributes) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitAttributes) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value.
func (o *AiGatewayRateLimitAttributes) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetEnabled returns the Enabled field value.
func (o *AiGatewayRateLimitAttributes) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *AiGatewayRateLimitAttributes) SetEnabled(v bool) {
	o.Enabled = v
}

// GetInputTokensPerMinute returns the InputTokensPerMinute field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiGatewayRateLimitAttributes) GetInputTokensPerMinute() int64 {
	if o == nil || o.InputTokensPerMinute.Get() == nil {
		var ret int64
		return ret
	}
	return *o.InputTokensPerMinute.Get()
}

// GetInputTokensPerMinuteOk returns a tuple with the InputTokensPerMinute field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AiGatewayRateLimitAttributes) GetInputTokensPerMinuteOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.InputTokensPerMinute.Get(), o.InputTokensPerMinute.IsSet()
}

// HasInputTokensPerMinute returns a boolean if a field has been set.
func (o *AiGatewayRateLimitAttributes) HasInputTokensPerMinute() bool {
	return o != nil && o.InputTokensPerMinute.IsSet()
}

// SetInputTokensPerMinute gets a reference to the given datadog.NullableInt64 and assigns it to the InputTokensPerMinute field.
func (o *AiGatewayRateLimitAttributes) SetInputTokensPerMinute(v int64) {
	o.InputTokensPerMinute.Set(&v)
}

// SetInputTokensPerMinuteNil sets the value for InputTokensPerMinute to be an explicit nil.
func (o *AiGatewayRateLimitAttributes) SetInputTokensPerMinuteNil() {
	o.InputTokensPerMinute.Set(nil)
}

// UnsetInputTokensPerMinute ensures that no value is present for InputTokensPerMinute, not even an explicit nil.
func (o *AiGatewayRateLimitAttributes) UnsetInputTokensPerMinute() {
	o.InputTokensPerMinute.Unset()
}

// GetIsTemplate returns the IsTemplate field value.
func (o *AiGatewayRateLimitAttributes) GetIsTemplate() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsTemplate
}

// GetIsTemplateOk returns a tuple with the IsTemplate field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitAttributes) GetIsTemplateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsTemplate, true
}

// SetIsTemplate sets field value.
func (o *AiGatewayRateLimitAttributes) SetIsTemplate(v bool) {
	o.IsTemplate = v
}

// GetMode returns the Mode field value.
func (o *AiGatewayRateLimitAttributes) GetMode() AiGatewayRateLimitMode {
	if o == nil {
		var ret AiGatewayRateLimitMode
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitAttributes) GetModeOk() (*AiGatewayRateLimitMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *AiGatewayRateLimitAttributes) SetMode(v AiGatewayRateLimitMode) {
	o.Mode = v
}

// GetMonitorId returns the MonitorId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiGatewayRateLimitAttributes) GetMonitorId() int64 {
	if o == nil || o.MonitorId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.MonitorId.Get()
}

// GetMonitorIdOk returns a tuple with the MonitorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AiGatewayRateLimitAttributes) GetMonitorIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MonitorId.Get(), o.MonitorId.IsSet()
}

// HasMonitorId returns a boolean if a field has been set.
func (o *AiGatewayRateLimitAttributes) HasMonitorId() bool {
	return o != nil && o.MonitorId.IsSet()
}

// SetMonitorId gets a reference to the given datadog.NullableInt64 and assigns it to the MonitorId field.
func (o *AiGatewayRateLimitAttributes) SetMonitorId(v int64) {
	o.MonitorId.Set(&v)
}

// SetMonitorIdNil sets the value for MonitorId to be an explicit nil.
func (o *AiGatewayRateLimitAttributes) SetMonitorIdNil() {
	o.MonitorId.Set(nil)
}

// UnsetMonitorId ensures that no value is present for MonitorId, not even an explicit nil.
func (o *AiGatewayRateLimitAttributes) UnsetMonitorId() {
	o.MonitorId.Unset()
}

// GetName returns the Name field value.
func (o *AiGatewayRateLimitAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *AiGatewayRateLimitAttributes) SetName(v string) {
	o.Name = v
}

// GetOutputTokensPerMinute returns the OutputTokensPerMinute field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiGatewayRateLimitAttributes) GetOutputTokensPerMinute() int64 {
	if o == nil || o.OutputTokensPerMinute.Get() == nil {
		var ret int64
		return ret
	}
	return *o.OutputTokensPerMinute.Get()
}

// GetOutputTokensPerMinuteOk returns a tuple with the OutputTokensPerMinute field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AiGatewayRateLimitAttributes) GetOutputTokensPerMinuteOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.OutputTokensPerMinute.Get(), o.OutputTokensPerMinute.IsSet()
}

// HasOutputTokensPerMinute returns a boolean if a field has been set.
func (o *AiGatewayRateLimitAttributes) HasOutputTokensPerMinute() bool {
	return o != nil && o.OutputTokensPerMinute.IsSet()
}

// SetOutputTokensPerMinute gets a reference to the given datadog.NullableInt64 and assigns it to the OutputTokensPerMinute field.
func (o *AiGatewayRateLimitAttributes) SetOutputTokensPerMinute(v int64) {
	o.OutputTokensPerMinute.Set(&v)
}

// SetOutputTokensPerMinuteNil sets the value for OutputTokensPerMinute to be an explicit nil.
func (o *AiGatewayRateLimitAttributes) SetOutputTokensPerMinuteNil() {
	o.OutputTokensPerMinute.Set(nil)
}

// UnsetOutputTokensPerMinute ensures that no value is present for OutputTokensPerMinute, not even an explicit nil.
func (o *AiGatewayRateLimitAttributes) UnsetOutputTokensPerMinute() {
	o.OutputTokensPerMinute.Unset()
}

// GetRequestsPerMinute returns the RequestsPerMinute field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiGatewayRateLimitAttributes) GetRequestsPerMinute() int64 {
	if o == nil || o.RequestsPerMinute.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RequestsPerMinute.Get()
}

// GetRequestsPerMinuteOk returns a tuple with the RequestsPerMinute field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AiGatewayRateLimitAttributes) GetRequestsPerMinuteOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestsPerMinute.Get(), o.RequestsPerMinute.IsSet()
}

// HasRequestsPerMinute returns a boolean if a field has been set.
func (o *AiGatewayRateLimitAttributes) HasRequestsPerMinute() bool {
	return o != nil && o.RequestsPerMinute.IsSet()
}

// SetRequestsPerMinute gets a reference to the given datadog.NullableInt64 and assigns it to the RequestsPerMinute field.
func (o *AiGatewayRateLimitAttributes) SetRequestsPerMinute(v int64) {
	o.RequestsPerMinute.Set(&v)
}

// SetRequestsPerMinuteNil sets the value for RequestsPerMinute to be an explicit nil.
func (o *AiGatewayRateLimitAttributes) SetRequestsPerMinuteNil() {
	o.RequestsPerMinute.Set(nil)
}

// UnsetRequestsPerMinute ensures that no value is present for RequestsPerMinute, not even an explicit nil.
func (o *AiGatewayRateLimitAttributes) UnsetRequestsPerMinute() {
	o.RequestsPerMinute.Unset()
}

// GetRevision returns the Revision field value.
func (o *AiGatewayRateLimitAttributes) GetRevision() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitAttributes) GetRevisionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revision, true
}

// SetRevision sets field value.
func (o *AiGatewayRateLimitAttributes) SetRevision(v int64) {
	o.Revision = v
}

// GetScopeRestricted returns the ScopeRestricted field value.
func (o *AiGatewayRateLimitAttributes) GetScopeRestricted() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.ScopeRestricted
}

// GetScopeRestrictedOk returns a tuple with the ScopeRestricted field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitAttributes) GetScopeRestrictedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScopeRestricted, true
}

// SetScopeRestricted sets field value.
func (o *AiGatewayRateLimitAttributes) SetScopeRestricted(v bool) {
	o.ScopeRestricted = v
}

// GetSelector returns the Selector field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiGatewayRateLimitAttributes) GetSelector() []AiGatewayRateLimitSelectorEntry {
	if o == nil {
		var ret []AiGatewayRateLimitSelectorEntry
		return ret
	}
	return o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AiGatewayRateLimitAttributes) GetSelectorOk() (*[]AiGatewayRateLimitSelectorEntry, bool) {
	if o == nil || o.Selector == nil {
		return nil, false
	}
	return &o.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *AiGatewayRateLimitAttributes) HasSelector() bool {
	return o != nil && o.Selector != nil
}

// SetSelector gets a reference to the given []AiGatewayRateLimitSelectorEntry and assigns it to the Selector field.
func (o *AiGatewayRateLimitAttributes) SetSelector(v []AiGatewayRateLimitSelectorEntry) {
	o.Selector = v
}

// GetTags returns the Tags field value.
func (o *AiGatewayRateLimitAttributes) GetTags() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitAttributes) GetTagsOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value.
func (o *AiGatewayRateLimitAttributes) SetTags(v map[string]string) {
	o.Tags = v
}

// GetTeamId returns the TeamId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiGatewayRateLimitAttributes) GetTeamId() uuid.UUID {
	if o == nil || o.TeamId.Get() == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.TeamId.Get()
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AiGatewayRateLimitAttributes) GetTeamIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return o.TeamId.Get(), o.TeamId.IsSet()
}

// HasTeamId returns a boolean if a field has been set.
func (o *AiGatewayRateLimitAttributes) HasTeamId() bool {
	return o != nil && o.TeamId.IsSet()
}

// SetTeamId gets a reference to the given datadog.NullableUUID and assigns it to the TeamId field.
func (o *AiGatewayRateLimitAttributes) SetTeamId(v uuid.UUID) {
	o.TeamId.Set(&v)
}

// SetTeamIdNil sets the value for TeamId to be an explicit nil.
func (o *AiGatewayRateLimitAttributes) SetTeamIdNil() {
	o.TeamId.Set(nil)
}

// UnsetTeamId ensures that no value is present for TeamId, not even an explicit nil.
func (o *AiGatewayRateLimitAttributes) UnsetTeamId() {
	o.TeamId.Unset()
}

// GetTemplateSplitDims returns the TemplateSplitDims field value.
func (o *AiGatewayRateLimitAttributes) GetTemplateSplitDims() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TemplateSplitDims
}

// GetTemplateSplitDimsOk returns a tuple with the TemplateSplitDims field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitAttributes) GetTemplateSplitDimsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateSplitDims, true
}

// SetTemplateSplitDims sets field value.
func (o *AiGatewayRateLimitAttributes) SetTemplateSplitDims(v []string) {
	o.TemplateSplitDims = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *AiGatewayRateLimitAttributes) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *AiGatewayRateLimitAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetUpdatedBy returns the UpdatedBy field value.
func (o *AiGatewayRateLimitAttributes) GetUpdatedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitAttributes) GetUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedBy, true
}

// SetUpdatedBy sets field value.
func (o *AiGatewayRateLimitAttributes) SetUpdatedBy(v string) {
	o.UpdatedBy = v
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiGatewayRateLimitAttributes) GetUserId() uuid.UUID {
	if o == nil || o.UserId.Get() == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AiGatewayRateLimitAttributes) GetUserIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *AiGatewayRateLimitAttributes) HasUserId() bool {
	return o != nil && o.UserId.IsSet()
}

// SetUserId gets a reference to the given datadog.NullableUUID and assigns it to the UserId field.
func (o *AiGatewayRateLimitAttributes) SetUserId(v uuid.UUID) {
	o.UserId.Set(&v)
}

// SetUserIdNil sets the value for UserId to be an explicit nil.
func (o *AiGatewayRateLimitAttributes) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil.
func (o *AiGatewayRateLimitAttributes) UnsetUserId() {
	o.UserId.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayRateLimitAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CostMicrocentsPerMinute.IsSet() {
		toSerialize["cost_microcents_per_minute"] = o.CostMicrocentsPerMinute.Get()
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["created_by"] = o.CreatedBy
	toSerialize["enabled"] = o.Enabled
	if o.InputTokensPerMinute.IsSet() {
		toSerialize["input_tokens_per_minute"] = o.InputTokensPerMinute.Get()
	}
	toSerialize["is_template"] = o.IsTemplate
	toSerialize["mode"] = o.Mode
	if o.MonitorId.IsSet() {
		toSerialize["monitor_id"] = o.MonitorId.Get()
	}
	toSerialize["name"] = o.Name
	if o.OutputTokensPerMinute.IsSet() {
		toSerialize["output_tokens_per_minute"] = o.OutputTokensPerMinute.Get()
	}
	if o.RequestsPerMinute.IsSet() {
		toSerialize["requests_per_minute"] = o.RequestsPerMinute.Get()
	}
	toSerialize["revision"] = o.Revision
	toSerialize["scope_restricted"] = o.ScopeRestricted
	if o.Selector != nil {
		toSerialize["selector"] = o.Selector
	}
	toSerialize["tags"] = o.Tags
	if o.TeamId.IsSet() {
		toSerialize["team_id"] = o.TeamId.Get()
	}
	toSerialize["template_split_dims"] = o.TemplateSplitDims
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["updated_by"] = o.UpdatedBy
	if o.UserId.IsSet() {
		toSerialize["user_id"] = o.UserId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayRateLimitAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CostMicrocentsPerMinute datadog.NullableInt64             `json:"cost_microcents_per_minute,omitempty"`
		CreatedAt               *time.Time                        `json:"created_at"`
		CreatedBy               *string                           `json:"created_by"`
		Enabled                 *bool                             `json:"enabled"`
		InputTokensPerMinute    datadog.NullableInt64             `json:"input_tokens_per_minute,omitempty"`
		IsTemplate              *bool                             `json:"is_template"`
		Mode                    *AiGatewayRateLimitMode           `json:"mode"`
		MonitorId               datadog.NullableInt64             `json:"monitor_id,omitempty"`
		Name                    *string                           `json:"name"`
		OutputTokensPerMinute   datadog.NullableInt64             `json:"output_tokens_per_minute,omitempty"`
		RequestsPerMinute       datadog.NullableInt64             `json:"requests_per_minute,omitempty"`
		Revision                *int64                            `json:"revision"`
		ScopeRestricted         *bool                             `json:"scope_restricted"`
		Selector                []AiGatewayRateLimitSelectorEntry `json:"selector,omitempty"`
		Tags                    *map[string]string                `json:"tags"`
		TeamId                  datadog.NullableUUID              `json:"team_id,omitempty"`
		TemplateSplitDims       *[]string                         `json:"template_split_dims"`
		UpdatedAt               *time.Time                        `json:"updated_at"`
		UpdatedBy               *string                           `json:"updated_by"`
		UserId                  datadog.NullableUUID              `json:"user_id,omitempty"`
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
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.IsTemplate == nil {
		return fmt.Errorf("required field is_template missing")
	}
	if all.Mode == nil {
		return fmt.Errorf("required field mode missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Revision == nil {
		return fmt.Errorf("required field revision missing")
	}
	if all.ScopeRestricted == nil {
		return fmt.Errorf("required field scope_restricted missing")
	}
	if all.Tags == nil {
		return fmt.Errorf("required field tags missing")
	}
	if all.TemplateSplitDims == nil {
		return fmt.Errorf("required field template_split_dims missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	if all.UpdatedBy == nil {
		return fmt.Errorf("required field updated_by missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cost_microcents_per_minute", "created_at", "created_by", "enabled", "input_tokens_per_minute", "is_template", "mode", "monitor_id", "name", "output_tokens_per_minute", "requests_per_minute", "revision", "scope_restricted", "selector", "tags", "team_id", "template_split_dims", "updated_at", "updated_by", "user_id"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CostMicrocentsPerMinute = all.CostMicrocentsPerMinute
	o.CreatedAt = *all.CreatedAt
	o.CreatedBy = *all.CreatedBy
	o.Enabled = *all.Enabled
	o.InputTokensPerMinute = all.InputTokensPerMinute
	o.IsTemplate = *all.IsTemplate
	if !all.Mode.IsValid() {
		hasInvalidField = true
	} else {
		o.Mode = *all.Mode
	}
	o.MonitorId = all.MonitorId
	o.Name = *all.Name
	o.OutputTokensPerMinute = all.OutputTokensPerMinute
	o.RequestsPerMinute = all.RequestsPerMinute
	o.Revision = *all.Revision
	o.ScopeRestricted = *all.ScopeRestricted
	o.Selector = all.Selector
	o.Tags = *all.Tags
	o.TeamId = all.TeamId
	o.TemplateSplitDims = *all.TemplateSplitDims
	o.UpdatedAt = *all.UpdatedAt
	o.UpdatedBy = *all.UpdatedBy
	o.UserId = all.UserId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
