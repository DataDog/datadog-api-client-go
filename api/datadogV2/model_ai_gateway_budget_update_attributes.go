// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayBudgetUpdateAttributes Mutable attributes of an AI Gateway budget.
type AiGatewayBudgetUpdateAttributes struct {
	// Whether a capacity change also applies to the active accounting period. Requires capacity_microcents. Unsupported for template budgets.
	ApplyToCurrentPeriod *bool `json:"apply_to_current_period,omitempty"`
	// Default capacity for future accounting periods, in microcents.
	CapacityMicrocents *int64 `json:"capacity_microcents,omitempty"`
	// Whether the budget is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Action taken when spend exceeds the budget.
	Enforcement *AiGatewayBudgetEnforcement `json:"enforcement,omitempty"`
	// Human-readable resource name.
	Name *string `json:"name,omitempty"`
	// Optional reason recorded with an active-period capacity adjustment.
	Reason *string `json:"reason,omitempty"`
	// Initialize current-period spend from historical priced LLM Observability spans. Defaults to false. Unsupported for template budgets; failures can be retried with an update.
	TrackExisting *bool `json:"track_existing,omitempty"`
	// Fraction of capacity at which to issue a warning.
	WarningThreshold *float64 `json:"warning_threshold,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayBudgetUpdateAttributes instantiates a new AiGatewayBudgetUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayBudgetUpdateAttributes() *AiGatewayBudgetUpdateAttributes {
	this := AiGatewayBudgetUpdateAttributes{}
	var applyToCurrentPeriod bool = false
	this.ApplyToCurrentPeriod = &applyToCurrentPeriod
	var trackExisting bool = false
	this.TrackExisting = &trackExisting
	return &this
}

// NewAiGatewayBudgetUpdateAttributesWithDefaults instantiates a new AiGatewayBudgetUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayBudgetUpdateAttributesWithDefaults() *AiGatewayBudgetUpdateAttributes {
	this := AiGatewayBudgetUpdateAttributes{}
	var applyToCurrentPeriod bool = false
	this.ApplyToCurrentPeriod = &applyToCurrentPeriod
	var trackExisting bool = false
	this.TrackExisting = &trackExisting
	return &this
}

// GetApplyToCurrentPeriod returns the ApplyToCurrentPeriod field value if set, zero value otherwise.
func (o *AiGatewayBudgetUpdateAttributes) GetApplyToCurrentPeriod() bool {
	if o == nil || o.ApplyToCurrentPeriod == nil {
		var ret bool
		return ret
	}
	return *o.ApplyToCurrentPeriod
}

// GetApplyToCurrentPeriodOk returns a tuple with the ApplyToCurrentPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetUpdateAttributes) GetApplyToCurrentPeriodOk() (*bool, bool) {
	if o == nil || o.ApplyToCurrentPeriod == nil {
		return nil, false
	}
	return o.ApplyToCurrentPeriod, true
}

// HasApplyToCurrentPeriod returns a boolean if a field has been set.
func (o *AiGatewayBudgetUpdateAttributes) HasApplyToCurrentPeriod() bool {
	return o != nil && o.ApplyToCurrentPeriod != nil
}

// SetApplyToCurrentPeriod gets a reference to the given bool and assigns it to the ApplyToCurrentPeriod field.
func (o *AiGatewayBudgetUpdateAttributes) SetApplyToCurrentPeriod(v bool) {
	o.ApplyToCurrentPeriod = &v
}

// GetCapacityMicrocents returns the CapacityMicrocents field value if set, zero value otherwise.
func (o *AiGatewayBudgetUpdateAttributes) GetCapacityMicrocents() int64 {
	if o == nil || o.CapacityMicrocents == nil {
		var ret int64
		return ret
	}
	return *o.CapacityMicrocents
}

// GetCapacityMicrocentsOk returns a tuple with the CapacityMicrocents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetUpdateAttributes) GetCapacityMicrocentsOk() (*int64, bool) {
	if o == nil || o.CapacityMicrocents == nil {
		return nil, false
	}
	return o.CapacityMicrocents, true
}

// HasCapacityMicrocents returns a boolean if a field has been set.
func (o *AiGatewayBudgetUpdateAttributes) HasCapacityMicrocents() bool {
	return o != nil && o.CapacityMicrocents != nil
}

// SetCapacityMicrocents gets a reference to the given int64 and assigns it to the CapacityMicrocents field.
func (o *AiGatewayBudgetUpdateAttributes) SetCapacityMicrocents(v int64) {
	o.CapacityMicrocents = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AiGatewayBudgetUpdateAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetUpdateAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AiGatewayBudgetUpdateAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AiGatewayBudgetUpdateAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEnforcement returns the Enforcement field value if set, zero value otherwise.
func (o *AiGatewayBudgetUpdateAttributes) GetEnforcement() AiGatewayBudgetEnforcement {
	if o == nil || o.Enforcement == nil {
		var ret AiGatewayBudgetEnforcement
		return ret
	}
	return *o.Enforcement
}

// GetEnforcementOk returns a tuple with the Enforcement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetUpdateAttributes) GetEnforcementOk() (*AiGatewayBudgetEnforcement, bool) {
	if o == nil || o.Enforcement == nil {
		return nil, false
	}
	return o.Enforcement, true
}

// HasEnforcement returns a boolean if a field has been set.
func (o *AiGatewayBudgetUpdateAttributes) HasEnforcement() bool {
	return o != nil && o.Enforcement != nil
}

// SetEnforcement gets a reference to the given AiGatewayBudgetEnforcement and assigns it to the Enforcement field.
func (o *AiGatewayBudgetUpdateAttributes) SetEnforcement(v AiGatewayBudgetEnforcement) {
	o.Enforcement = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AiGatewayBudgetUpdateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AiGatewayBudgetUpdateAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AiGatewayBudgetUpdateAttributes) SetName(v string) {
	o.Name = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *AiGatewayBudgetUpdateAttributes) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetUpdateAttributes) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *AiGatewayBudgetUpdateAttributes) HasReason() bool {
	return o != nil && o.Reason != nil
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *AiGatewayBudgetUpdateAttributes) SetReason(v string) {
	o.Reason = &v
}

// GetTrackExisting returns the TrackExisting field value if set, zero value otherwise.
func (o *AiGatewayBudgetUpdateAttributes) GetTrackExisting() bool {
	if o == nil || o.TrackExisting == nil {
		var ret bool
		return ret
	}
	return *o.TrackExisting
}

// GetTrackExistingOk returns a tuple with the TrackExisting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetUpdateAttributes) GetTrackExistingOk() (*bool, bool) {
	if o == nil || o.TrackExisting == nil {
		return nil, false
	}
	return o.TrackExisting, true
}

// HasTrackExisting returns a boolean if a field has been set.
func (o *AiGatewayBudgetUpdateAttributes) HasTrackExisting() bool {
	return o != nil && o.TrackExisting != nil
}

// SetTrackExisting gets a reference to the given bool and assigns it to the TrackExisting field.
func (o *AiGatewayBudgetUpdateAttributes) SetTrackExisting(v bool) {
	o.TrackExisting = &v
}

// GetWarningThreshold returns the WarningThreshold field value if set, zero value otherwise.
func (o *AiGatewayBudgetUpdateAttributes) GetWarningThreshold() float64 {
	if o == nil || o.WarningThreshold == nil {
		var ret float64
		return ret
	}
	return *o.WarningThreshold
}

// GetWarningThresholdOk returns a tuple with the WarningThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetUpdateAttributes) GetWarningThresholdOk() (*float64, bool) {
	if o == nil || o.WarningThreshold == nil {
		return nil, false
	}
	return o.WarningThreshold, true
}

// HasWarningThreshold returns a boolean if a field has been set.
func (o *AiGatewayBudgetUpdateAttributes) HasWarningThreshold() bool {
	return o != nil && o.WarningThreshold != nil
}

// SetWarningThreshold gets a reference to the given float64 and assigns it to the WarningThreshold field.
func (o *AiGatewayBudgetUpdateAttributes) SetWarningThreshold(v float64) {
	o.WarningThreshold = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayBudgetUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ApplyToCurrentPeriod != nil {
		toSerialize["apply_to_current_period"] = o.ApplyToCurrentPeriod
	}
	if o.CapacityMicrocents != nil {
		toSerialize["capacity_microcents"] = o.CapacityMicrocents
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Enforcement != nil {
		toSerialize["enforcement"] = o.Enforcement
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.TrackExisting != nil {
		toSerialize["track_existing"] = o.TrackExisting
	}
	if o.WarningThreshold != nil {
		toSerialize["warning_threshold"] = o.WarningThreshold
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayBudgetUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApplyToCurrentPeriod *bool                       `json:"apply_to_current_period,omitempty"`
		CapacityMicrocents   *int64                      `json:"capacity_microcents,omitempty"`
		Enabled              *bool                       `json:"enabled,omitempty"`
		Enforcement          *AiGatewayBudgetEnforcement `json:"enforcement,omitempty"`
		Name                 *string                     `json:"name,omitempty"`
		Reason               *string                     `json:"reason,omitempty"`
		TrackExisting        *bool                       `json:"track_existing,omitempty"`
		WarningThreshold     *float64                    `json:"warning_threshold,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"apply_to_current_period", "capacity_microcents", "enabled", "enforcement", "name", "reason", "track_existing", "warning_threshold"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ApplyToCurrentPeriod = all.ApplyToCurrentPeriod
	o.CapacityMicrocents = all.CapacityMicrocents
	o.Enabled = all.Enabled
	if all.Enforcement != nil && !all.Enforcement.IsValid() {
		hasInvalidField = true
	} else {
		o.Enforcement = all.Enforcement
	}
	o.Name = all.Name
	o.Reason = all.Reason
	o.TrackExisting = all.TrackExisting
	o.WarningThreshold = all.WarningThreshold

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
