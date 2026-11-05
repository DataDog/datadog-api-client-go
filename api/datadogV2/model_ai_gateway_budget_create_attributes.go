// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayBudgetCreateAttributes Attributes used to create an AI Gateway budget.
type AiGatewayBudgetCreateAttributes struct {
	// Budget capacity in microcents.
	CapacityMicrocents int64 `json:"capacity_microcents"`
	// Action taken when spend exceeds the budget.
	Enforcement AiGatewayBudgetEnforcement `json:"enforcement"`
	// Human-readable resource name.
	Name string `json:"name"`
	// Dimension rules for a template budget. A nonempty selector creates independent counters for matched instances. Omit or pass an empty selector for a legacy fixed-scope budget.
	Selector []AiGatewayBudgetSelectorEntry `json:"selector,omitempty"`
	// Request tags that must match for the budget to apply.
	Tags map[string]string `json:"tags,omitempty"`
	// Optional Datadog Team scope.
	TeamId *uuid.UUID `json:"team_id,omitempty"`
	// Initialize current-period spend from historical priced LLM Observability spans. Defaults to false. Unsupported for template budgets; failures can be retried with an update.
	TrackExisting *bool `json:"track_existing,omitempty"`
	// Optional user scope.
	UserId *uuid.UUID `json:"user_id,omitempty"`
	// Fraction of capacity at which to issue a warning.
	WarningThreshold *float64 `json:"warning_threshold,omitempty"`
	// UTC accounting period for a budget.
	WindowPolicy AiGatewayBudgetWindowPolicy `json:"window_policy"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayBudgetCreateAttributes instantiates a new AiGatewayBudgetCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayBudgetCreateAttributes(capacityMicrocents int64, enforcement AiGatewayBudgetEnforcement, name string, windowPolicy AiGatewayBudgetWindowPolicy) *AiGatewayBudgetCreateAttributes {
	this := AiGatewayBudgetCreateAttributes{}
	this.CapacityMicrocents = capacityMicrocents
	this.Enforcement = enforcement
	this.Name = name
	var trackExisting bool = false
	this.TrackExisting = &trackExisting
	var warningThreshold float64 = 0.9
	this.WarningThreshold = &warningThreshold
	this.WindowPolicy = windowPolicy
	return &this
}

// NewAiGatewayBudgetCreateAttributesWithDefaults instantiates a new AiGatewayBudgetCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayBudgetCreateAttributesWithDefaults() *AiGatewayBudgetCreateAttributes {
	this := AiGatewayBudgetCreateAttributes{}
	var trackExisting bool = false
	this.TrackExisting = &trackExisting
	var warningThreshold float64 = 0.9
	this.WarningThreshold = &warningThreshold
	return &this
}

// GetCapacityMicrocents returns the CapacityMicrocents field value.
func (o *AiGatewayBudgetCreateAttributes) GetCapacityMicrocents() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.CapacityMicrocents
}

// GetCapacityMicrocentsOk returns a tuple with the CapacityMicrocents field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetCreateAttributes) GetCapacityMicrocentsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CapacityMicrocents, true
}

// SetCapacityMicrocents sets field value.
func (o *AiGatewayBudgetCreateAttributes) SetCapacityMicrocents(v int64) {
	o.CapacityMicrocents = v
}

// GetEnforcement returns the Enforcement field value.
func (o *AiGatewayBudgetCreateAttributes) GetEnforcement() AiGatewayBudgetEnforcement {
	if o == nil {
		var ret AiGatewayBudgetEnforcement
		return ret
	}
	return o.Enforcement
}

// GetEnforcementOk returns a tuple with the Enforcement field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetCreateAttributes) GetEnforcementOk() (*AiGatewayBudgetEnforcement, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enforcement, true
}

// SetEnforcement sets field value.
func (o *AiGatewayBudgetCreateAttributes) SetEnforcement(v AiGatewayBudgetEnforcement) {
	o.Enforcement = v
}

// GetName returns the Name field value.
func (o *AiGatewayBudgetCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *AiGatewayBudgetCreateAttributes) SetName(v string) {
	o.Name = v
}

// GetSelector returns the Selector field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiGatewayBudgetCreateAttributes) GetSelector() []AiGatewayBudgetSelectorEntry {
	if o == nil {
		var ret []AiGatewayBudgetSelectorEntry
		return ret
	}
	return o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AiGatewayBudgetCreateAttributes) GetSelectorOk() (*[]AiGatewayBudgetSelectorEntry, bool) {
	if o == nil || o.Selector == nil {
		return nil, false
	}
	return &o.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *AiGatewayBudgetCreateAttributes) HasSelector() bool {
	return o != nil && o.Selector != nil
}

// SetSelector gets a reference to the given []AiGatewayBudgetSelectorEntry and assigns it to the Selector field.
func (o *AiGatewayBudgetCreateAttributes) SetSelector(v []AiGatewayBudgetSelectorEntry) {
	o.Selector = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *AiGatewayBudgetCreateAttributes) GetTags() map[string]string {
	if o == nil || o.Tags == nil {
		var ret map[string]string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetCreateAttributes) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AiGatewayBudgetCreateAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *AiGatewayBudgetCreateAttributes) SetTags(v map[string]string) {
	o.Tags = v
}

// GetTeamId returns the TeamId field value if set, zero value otherwise.
func (o *AiGatewayBudgetCreateAttributes) GetTeamId() uuid.UUID {
	if o == nil || o.TeamId == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetCreateAttributes) GetTeamIdOk() (*uuid.UUID, bool) {
	if o == nil || o.TeamId == nil {
		return nil, false
	}
	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *AiGatewayBudgetCreateAttributes) HasTeamId() bool {
	return o != nil && o.TeamId != nil
}

// SetTeamId gets a reference to the given uuid.UUID and assigns it to the TeamId field.
func (o *AiGatewayBudgetCreateAttributes) SetTeamId(v uuid.UUID) {
	o.TeamId = &v
}

// GetTrackExisting returns the TrackExisting field value if set, zero value otherwise.
func (o *AiGatewayBudgetCreateAttributes) GetTrackExisting() bool {
	if o == nil || o.TrackExisting == nil {
		var ret bool
		return ret
	}
	return *o.TrackExisting
}

// GetTrackExistingOk returns a tuple with the TrackExisting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetCreateAttributes) GetTrackExistingOk() (*bool, bool) {
	if o == nil || o.TrackExisting == nil {
		return nil, false
	}
	return o.TrackExisting, true
}

// HasTrackExisting returns a boolean if a field has been set.
func (o *AiGatewayBudgetCreateAttributes) HasTrackExisting() bool {
	return o != nil && o.TrackExisting != nil
}

// SetTrackExisting gets a reference to the given bool and assigns it to the TrackExisting field.
func (o *AiGatewayBudgetCreateAttributes) SetTrackExisting(v bool) {
	o.TrackExisting = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *AiGatewayBudgetCreateAttributes) GetUserId() uuid.UUID {
	if o == nil || o.UserId == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetCreateAttributes) GetUserIdOk() (*uuid.UUID, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *AiGatewayBudgetCreateAttributes) HasUserId() bool {
	return o != nil && o.UserId != nil
}

// SetUserId gets a reference to the given uuid.UUID and assigns it to the UserId field.
func (o *AiGatewayBudgetCreateAttributes) SetUserId(v uuid.UUID) {
	o.UserId = &v
}

// GetWarningThreshold returns the WarningThreshold field value if set, zero value otherwise.
func (o *AiGatewayBudgetCreateAttributes) GetWarningThreshold() float64 {
	if o == nil || o.WarningThreshold == nil {
		var ret float64
		return ret
	}
	return *o.WarningThreshold
}

// GetWarningThresholdOk returns a tuple with the WarningThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetCreateAttributes) GetWarningThresholdOk() (*float64, bool) {
	if o == nil || o.WarningThreshold == nil {
		return nil, false
	}
	return o.WarningThreshold, true
}

// HasWarningThreshold returns a boolean if a field has been set.
func (o *AiGatewayBudgetCreateAttributes) HasWarningThreshold() bool {
	return o != nil && o.WarningThreshold != nil
}

// SetWarningThreshold gets a reference to the given float64 and assigns it to the WarningThreshold field.
func (o *AiGatewayBudgetCreateAttributes) SetWarningThreshold(v float64) {
	o.WarningThreshold = &v
}

// GetWindowPolicy returns the WindowPolicy field value.
func (o *AiGatewayBudgetCreateAttributes) GetWindowPolicy() AiGatewayBudgetWindowPolicy {
	if o == nil {
		var ret AiGatewayBudgetWindowPolicy
		return ret
	}
	return o.WindowPolicy
}

// GetWindowPolicyOk returns a tuple with the WindowPolicy field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetCreateAttributes) GetWindowPolicyOk() (*AiGatewayBudgetWindowPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WindowPolicy, true
}

// SetWindowPolicy sets field value.
func (o *AiGatewayBudgetCreateAttributes) SetWindowPolicy(v AiGatewayBudgetWindowPolicy) {
	o.WindowPolicy = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayBudgetCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["capacity_microcents"] = o.CapacityMicrocents
	toSerialize["enforcement"] = o.Enforcement
	toSerialize["name"] = o.Name
	if o.Selector != nil {
		toSerialize["selector"] = o.Selector
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.TeamId != nil {
		toSerialize["team_id"] = o.TeamId
	}
	if o.TrackExisting != nil {
		toSerialize["track_existing"] = o.TrackExisting
	}
	if o.UserId != nil {
		toSerialize["user_id"] = o.UserId
	}
	if o.WarningThreshold != nil {
		toSerialize["warning_threshold"] = o.WarningThreshold
	}
	toSerialize["window_policy"] = o.WindowPolicy

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayBudgetCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CapacityMicrocents *int64                         `json:"capacity_microcents"`
		Enforcement        *AiGatewayBudgetEnforcement    `json:"enforcement"`
		Name               *string                        `json:"name"`
		Selector           []AiGatewayBudgetSelectorEntry `json:"selector,omitempty"`
		Tags               map[string]string              `json:"tags,omitempty"`
		TeamId             *uuid.UUID                     `json:"team_id,omitempty"`
		TrackExisting      *bool                          `json:"track_existing,omitempty"`
		UserId             *uuid.UUID                     `json:"user_id,omitempty"`
		WarningThreshold   *float64                       `json:"warning_threshold,omitempty"`
		WindowPolicy       *AiGatewayBudgetWindowPolicy   `json:"window_policy"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CapacityMicrocents == nil {
		return fmt.Errorf("required field capacity_microcents missing")
	}
	if all.Enforcement == nil {
		return fmt.Errorf("required field enforcement missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.WindowPolicy == nil {
		return fmt.Errorf("required field window_policy missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"capacity_microcents", "enforcement", "name", "selector", "tags", "team_id", "track_existing", "user_id", "warning_threshold", "window_policy"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CapacityMicrocents = *all.CapacityMicrocents
	if !all.Enforcement.IsValid() {
		hasInvalidField = true
	} else {
		o.Enforcement = *all.Enforcement
	}
	o.Name = *all.Name
	o.Selector = all.Selector
	o.Tags = all.Tags
	o.TeamId = all.TeamId
	o.TrackExisting = all.TrackExisting
	o.UserId = all.UserId
	o.WarningThreshold = all.WarningThreshold
	if !all.WindowPolicy.IsValid() {
		hasInvalidField = true
	} else {
		o.WindowPolicy = *all.WindowPolicy
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
