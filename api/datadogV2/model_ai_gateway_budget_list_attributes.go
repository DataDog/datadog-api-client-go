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

// AiGatewayBudgetListAttributes Budget configuration including spend and capacity for the active accounting window.
type AiGatewayBudgetListAttributes struct {
	// Budget capacity in microcents.
	CapacityMicrocents int64 `json:"capacity_microcents"`
	// Time when the resource was created.
	CreatedAt time.Time `json:"created_at"`
	// Identifier of the principal that created the resource.
	CreatedBy string `json:"created_by"`
	// Capacity enforced for the active window, in microcents.
	CurrentPeriodCapacityMicrocents int64 `json:"current_period_capacity_microcents"`
	// Spend currently counted against the active window, in microcents.
	CurrentSpendMicrocents int64 `json:"current_spend_microcents"`
	// Whether the resource is enabled.
	Enabled bool `json:"enabled"`
	// Action taken when spend exceeds the budget.
	Enforcement AiGatewayBudgetEnforcement `json:"enforcement"`
	// Datadog monitor ID provisioned for this budget. Null when monitor provisioning was skipped or failed.
	MonitorId datadog.NullableInt64 `json:"monitor_id"`
	// Human-readable resource name.
	Name string `json:"name"`
	// Whether provider or model scope metadata was withheld. Omit unchanged scope fields on an unrelated update to preserve their saved values.
	ScopeRestricted bool `json:"scope_restricted"`
	// Dimension rules for a template budget. A nonempty selector creates independent counters for matched instances. Omit or pass an empty selector for a legacy fixed-scope budget.
	Selector []AiGatewayBudgetSelectorEntry `json:"selector,omitempty"`
	// Request tags that must match for the budget to apply.
	Tags map[string]string `json:"tags"`
	// Optional Datadog Team scope.
	TeamId *uuid.UUID `json:"team_id,omitempty"`
	// Time when the resource was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// User scope. Omitted when the budget is not user-scoped.
	UserId *uuid.UUID `json:"user_id,omitempty"`
	// Fraction of capacity at which to issue a warning.
	WarningThreshold float64 `json:"warning_threshold"`
	// End of the active accounting window.
	WindowEnd time.Time `json:"window_end"`
	// UTC accounting period for a budget.
	WindowPolicy AiGatewayBudgetWindowPolicy `json:"window_policy"`
	// Start of the active accounting window.
	WindowStart time.Time `json:"window_start"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayBudgetListAttributes instantiates a new AiGatewayBudgetListAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayBudgetListAttributes(capacityMicrocents int64, createdAt time.Time, createdBy string, currentPeriodCapacityMicrocents int64, currentSpendMicrocents int64, enabled bool, enforcement AiGatewayBudgetEnforcement, monitorId datadog.NullableInt64, name string, scopeRestricted bool, tags map[string]string, updatedAt time.Time, warningThreshold float64, windowEnd time.Time, windowPolicy AiGatewayBudgetWindowPolicy, windowStart time.Time) *AiGatewayBudgetListAttributes {
	this := AiGatewayBudgetListAttributes{}
	this.CapacityMicrocents = capacityMicrocents
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.CurrentPeriodCapacityMicrocents = currentPeriodCapacityMicrocents
	this.CurrentSpendMicrocents = currentSpendMicrocents
	this.Enabled = enabled
	this.Enforcement = enforcement
	this.MonitorId = monitorId
	this.Name = name
	this.ScopeRestricted = scopeRestricted
	this.Tags = tags
	this.UpdatedAt = updatedAt
	this.WarningThreshold = warningThreshold
	this.WindowEnd = windowEnd
	this.WindowPolicy = windowPolicy
	this.WindowStart = windowStart
	return &this
}

// NewAiGatewayBudgetListAttributesWithDefaults instantiates a new AiGatewayBudgetListAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayBudgetListAttributesWithDefaults() *AiGatewayBudgetListAttributes {
	this := AiGatewayBudgetListAttributes{}
	return &this
}

// GetCapacityMicrocents returns the CapacityMicrocents field value.
func (o *AiGatewayBudgetListAttributes) GetCapacityMicrocents() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.CapacityMicrocents
}

// GetCapacityMicrocentsOk returns a tuple with the CapacityMicrocents field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetListAttributes) GetCapacityMicrocentsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CapacityMicrocents, true
}

// SetCapacityMicrocents sets field value.
func (o *AiGatewayBudgetListAttributes) SetCapacityMicrocents(v int64) {
	o.CapacityMicrocents = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *AiGatewayBudgetListAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetListAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *AiGatewayBudgetListAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value.
func (o *AiGatewayBudgetListAttributes) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetListAttributes) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value.
func (o *AiGatewayBudgetListAttributes) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCurrentPeriodCapacityMicrocents returns the CurrentPeriodCapacityMicrocents field value.
func (o *AiGatewayBudgetListAttributes) GetCurrentPeriodCapacityMicrocents() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.CurrentPeriodCapacityMicrocents
}

// GetCurrentPeriodCapacityMicrocentsOk returns a tuple with the CurrentPeriodCapacityMicrocents field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetListAttributes) GetCurrentPeriodCapacityMicrocentsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentPeriodCapacityMicrocents, true
}

// SetCurrentPeriodCapacityMicrocents sets field value.
func (o *AiGatewayBudgetListAttributes) SetCurrentPeriodCapacityMicrocents(v int64) {
	o.CurrentPeriodCapacityMicrocents = v
}

// GetCurrentSpendMicrocents returns the CurrentSpendMicrocents field value.
func (o *AiGatewayBudgetListAttributes) GetCurrentSpendMicrocents() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.CurrentSpendMicrocents
}

// GetCurrentSpendMicrocentsOk returns a tuple with the CurrentSpendMicrocents field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetListAttributes) GetCurrentSpendMicrocentsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentSpendMicrocents, true
}

// SetCurrentSpendMicrocents sets field value.
func (o *AiGatewayBudgetListAttributes) SetCurrentSpendMicrocents(v int64) {
	o.CurrentSpendMicrocents = v
}

// GetEnabled returns the Enabled field value.
func (o *AiGatewayBudgetListAttributes) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetListAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *AiGatewayBudgetListAttributes) SetEnabled(v bool) {
	o.Enabled = v
}

// GetEnforcement returns the Enforcement field value.
func (o *AiGatewayBudgetListAttributes) GetEnforcement() AiGatewayBudgetEnforcement {
	if o == nil {
		var ret AiGatewayBudgetEnforcement
		return ret
	}
	return o.Enforcement
}

// GetEnforcementOk returns a tuple with the Enforcement field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetListAttributes) GetEnforcementOk() (*AiGatewayBudgetEnforcement, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enforcement, true
}

// SetEnforcement sets field value.
func (o *AiGatewayBudgetListAttributes) SetEnforcement(v AiGatewayBudgetEnforcement) {
	o.Enforcement = v
}

// GetMonitorId returns the MonitorId field value.
// If the value is explicit nil, the zero value for int64 will be returned.
func (o *AiGatewayBudgetListAttributes) GetMonitorId() int64 {
	if o == nil || o.MonitorId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.MonitorId.Get()
}

// GetMonitorIdOk returns a tuple with the MonitorId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AiGatewayBudgetListAttributes) GetMonitorIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MonitorId.Get(), o.MonitorId.IsSet()
}

// SetMonitorId sets field value.
func (o *AiGatewayBudgetListAttributes) SetMonitorId(v int64) {
	o.MonitorId.Set(&v)
}

// GetName returns the Name field value.
func (o *AiGatewayBudgetListAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetListAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *AiGatewayBudgetListAttributes) SetName(v string) {
	o.Name = v
}

// GetScopeRestricted returns the ScopeRestricted field value.
func (o *AiGatewayBudgetListAttributes) GetScopeRestricted() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.ScopeRestricted
}

// GetScopeRestrictedOk returns a tuple with the ScopeRestricted field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetListAttributes) GetScopeRestrictedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScopeRestricted, true
}

// SetScopeRestricted sets field value.
func (o *AiGatewayBudgetListAttributes) SetScopeRestricted(v bool) {
	o.ScopeRestricted = v
}

// GetSelector returns the Selector field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiGatewayBudgetListAttributes) GetSelector() []AiGatewayBudgetSelectorEntry {
	if o == nil {
		var ret []AiGatewayBudgetSelectorEntry
		return ret
	}
	return o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AiGatewayBudgetListAttributes) GetSelectorOk() (*[]AiGatewayBudgetSelectorEntry, bool) {
	if o == nil || o.Selector == nil {
		return nil, false
	}
	return &o.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *AiGatewayBudgetListAttributes) HasSelector() bool {
	return o != nil && o.Selector != nil
}

// SetSelector gets a reference to the given []AiGatewayBudgetSelectorEntry and assigns it to the Selector field.
func (o *AiGatewayBudgetListAttributes) SetSelector(v []AiGatewayBudgetSelectorEntry) {
	o.Selector = v
}

// GetTags returns the Tags field value.
func (o *AiGatewayBudgetListAttributes) GetTags() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetListAttributes) GetTagsOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value.
func (o *AiGatewayBudgetListAttributes) SetTags(v map[string]string) {
	o.Tags = v
}

// GetTeamId returns the TeamId field value if set, zero value otherwise.
func (o *AiGatewayBudgetListAttributes) GetTeamId() uuid.UUID {
	if o == nil || o.TeamId == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetListAttributes) GetTeamIdOk() (*uuid.UUID, bool) {
	if o == nil || o.TeamId == nil {
		return nil, false
	}
	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *AiGatewayBudgetListAttributes) HasTeamId() bool {
	return o != nil && o.TeamId != nil
}

// SetTeamId gets a reference to the given uuid.UUID and assigns it to the TeamId field.
func (o *AiGatewayBudgetListAttributes) SetTeamId(v uuid.UUID) {
	o.TeamId = &v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *AiGatewayBudgetListAttributes) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetListAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *AiGatewayBudgetListAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *AiGatewayBudgetListAttributes) GetUserId() uuid.UUID {
	if o == nil || o.UserId == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetListAttributes) GetUserIdOk() (*uuid.UUID, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *AiGatewayBudgetListAttributes) HasUserId() bool {
	return o != nil && o.UserId != nil
}

// SetUserId gets a reference to the given uuid.UUID and assigns it to the UserId field.
func (o *AiGatewayBudgetListAttributes) SetUserId(v uuid.UUID) {
	o.UserId = &v
}

// GetWarningThreshold returns the WarningThreshold field value.
func (o *AiGatewayBudgetListAttributes) GetWarningThreshold() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.WarningThreshold
}

// GetWarningThresholdOk returns a tuple with the WarningThreshold field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetListAttributes) GetWarningThresholdOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WarningThreshold, true
}

// SetWarningThreshold sets field value.
func (o *AiGatewayBudgetListAttributes) SetWarningThreshold(v float64) {
	o.WarningThreshold = v
}

// GetWindowEnd returns the WindowEnd field value.
func (o *AiGatewayBudgetListAttributes) GetWindowEnd() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.WindowEnd
}

// GetWindowEndOk returns a tuple with the WindowEnd field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetListAttributes) GetWindowEndOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WindowEnd, true
}

// SetWindowEnd sets field value.
func (o *AiGatewayBudgetListAttributes) SetWindowEnd(v time.Time) {
	o.WindowEnd = v
}

// GetWindowPolicy returns the WindowPolicy field value.
func (o *AiGatewayBudgetListAttributes) GetWindowPolicy() AiGatewayBudgetWindowPolicy {
	if o == nil {
		var ret AiGatewayBudgetWindowPolicy
		return ret
	}
	return o.WindowPolicy
}

// GetWindowPolicyOk returns a tuple with the WindowPolicy field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetListAttributes) GetWindowPolicyOk() (*AiGatewayBudgetWindowPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WindowPolicy, true
}

// SetWindowPolicy sets field value.
func (o *AiGatewayBudgetListAttributes) SetWindowPolicy(v AiGatewayBudgetWindowPolicy) {
	o.WindowPolicy = v
}

// GetWindowStart returns the WindowStart field value.
func (o *AiGatewayBudgetListAttributes) GetWindowStart() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.WindowStart
}

// GetWindowStartOk returns a tuple with the WindowStart field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetListAttributes) GetWindowStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WindowStart, true
}

// SetWindowStart sets field value.
func (o *AiGatewayBudgetListAttributes) SetWindowStart(v time.Time) {
	o.WindowStart = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayBudgetListAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["capacity_microcents"] = o.CapacityMicrocents
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["created_by"] = o.CreatedBy
	toSerialize["current_period_capacity_microcents"] = o.CurrentPeriodCapacityMicrocents
	toSerialize["current_spend_microcents"] = o.CurrentSpendMicrocents
	toSerialize["enabled"] = o.Enabled
	toSerialize["enforcement"] = o.Enforcement
	toSerialize["monitor_id"] = o.MonitorId.Get()
	toSerialize["name"] = o.Name
	toSerialize["scope_restricted"] = o.ScopeRestricted
	if o.Selector != nil {
		toSerialize["selector"] = o.Selector
	}
	toSerialize["tags"] = o.Tags
	if o.TeamId != nil {
		toSerialize["team_id"] = o.TeamId
	}
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.UserId != nil {
		toSerialize["user_id"] = o.UserId
	}
	toSerialize["warning_threshold"] = o.WarningThreshold
	if o.WindowEnd.Nanosecond() == 0 {
		toSerialize["window_end"] = o.WindowEnd.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["window_end"] = o.WindowEnd.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["window_policy"] = o.WindowPolicy
	if o.WindowStart.Nanosecond() == 0 {
		toSerialize["window_start"] = o.WindowStart.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["window_start"] = o.WindowStart.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayBudgetListAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CapacityMicrocents              *int64                         `json:"capacity_microcents"`
		CreatedAt                       *time.Time                     `json:"created_at"`
		CreatedBy                       *string                        `json:"created_by"`
		CurrentPeriodCapacityMicrocents *int64                         `json:"current_period_capacity_microcents"`
		CurrentSpendMicrocents          *int64                         `json:"current_spend_microcents"`
		Enabled                         *bool                          `json:"enabled"`
		Enforcement                     *AiGatewayBudgetEnforcement    `json:"enforcement"`
		MonitorId                       datadog.NullableInt64          `json:"monitor_id"`
		Name                            *string                        `json:"name"`
		ScopeRestricted                 *bool                          `json:"scope_restricted"`
		Selector                        []AiGatewayBudgetSelectorEntry `json:"selector,omitempty"`
		Tags                            *map[string]string             `json:"tags"`
		TeamId                          *uuid.UUID                     `json:"team_id,omitempty"`
		UpdatedAt                       *time.Time                     `json:"updated_at"`
		UserId                          *uuid.UUID                     `json:"user_id,omitempty"`
		WarningThreshold                *float64                       `json:"warning_threshold"`
		WindowEnd                       *time.Time                     `json:"window_end"`
		WindowPolicy                    *AiGatewayBudgetWindowPolicy   `json:"window_policy"`
		WindowStart                     *time.Time                     `json:"window_start"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CapacityMicrocents == nil {
		return fmt.Errorf("required field capacity_microcents missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.CreatedBy == nil {
		return fmt.Errorf("required field created_by missing")
	}
	if all.CurrentPeriodCapacityMicrocents == nil {
		return fmt.Errorf("required field current_period_capacity_microcents missing")
	}
	if all.CurrentSpendMicrocents == nil {
		return fmt.Errorf("required field current_spend_microcents missing")
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.Enforcement == nil {
		return fmt.Errorf("required field enforcement missing")
	}
	if !all.MonitorId.IsSet() {
		return fmt.Errorf("required field monitor_id missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.ScopeRestricted == nil {
		return fmt.Errorf("required field scope_restricted missing")
	}
	if all.Tags == nil {
		return fmt.Errorf("required field tags missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	if all.WarningThreshold == nil {
		return fmt.Errorf("required field warning_threshold missing")
	}
	if all.WindowEnd == nil {
		return fmt.Errorf("required field window_end missing")
	}
	if all.WindowPolicy == nil {
		return fmt.Errorf("required field window_policy missing")
	}
	if all.WindowStart == nil {
		return fmt.Errorf("required field window_start missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"capacity_microcents", "created_at", "created_by", "current_period_capacity_microcents", "current_spend_microcents", "enabled", "enforcement", "monitor_id", "name", "scope_restricted", "selector", "tags", "team_id", "updated_at", "user_id", "warning_threshold", "window_end", "window_policy", "window_start"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CapacityMicrocents = *all.CapacityMicrocents
	o.CreatedAt = *all.CreatedAt
	o.CreatedBy = *all.CreatedBy
	o.CurrentPeriodCapacityMicrocents = *all.CurrentPeriodCapacityMicrocents
	o.CurrentSpendMicrocents = *all.CurrentSpendMicrocents
	o.Enabled = *all.Enabled
	if !all.Enforcement.IsValid() {
		hasInvalidField = true
	} else {
		o.Enforcement = *all.Enforcement
	}
	o.MonitorId = all.MonitorId
	o.Name = *all.Name
	o.ScopeRestricted = *all.ScopeRestricted
	o.Selector = all.Selector
	o.Tags = *all.Tags
	o.TeamId = all.TeamId
	o.UpdatedAt = *all.UpdatedAt
	o.UserId = all.UserId
	o.WarningThreshold = *all.WarningThreshold
	o.WindowEnd = *all.WindowEnd
	if !all.WindowPolicy.IsValid() {
		hasInvalidField = true
	} else {
		o.WindowPolicy = *all.WindowPolicy
	}
	o.WindowStart = *all.WindowStart

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
