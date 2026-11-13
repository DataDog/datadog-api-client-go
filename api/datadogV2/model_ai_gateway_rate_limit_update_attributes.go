// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayRateLimitUpdateAttributes Attributes used to update an AI Gateway rate limit. Only sent fields are changed.
type AiGatewayRateLimitUpdateAttributes struct {
	// Spend limit in microcents per rolling minute. Set to 0 to remove the cost limit. Null is ignored. At least one limit must remain.
	CostMicrocentsPerMinute datadog.NullableInt64 `json:"cost_microcents_per_minute,omitempty"`
	// Whether this rate limit is active.
	Enabled *bool `json:"enabled,omitempty"`
	// Set to 0 to remove the request limit. Null is ignored. At least one limit must remain.
	InputTokensPerMinute datadog.NullableInt64 `json:"input_tokens_per_minute,omitempty"`
	// Action taken when request rate exceeds the limit.
	Mode *AiGatewayRateLimitMode `json:"mode,omitempty"`
	// Rate limit name.
	Name *string `json:"name,omitempty"`
	// Set to 0 to remove the request limit. Null is ignored. At least one limit must remain.
	OutputTokensPerMinute datadog.NullableInt64 `json:"output_tokens_per_minute,omitempty"`
	// Set to 0 to remove the request limit. Null is ignored. At least one limit must remain.
	RequestsPerMinute datadog.NullableInt64 `json:"requests_per_minute,omitempty"`
	// Replaces all scope tags. An empty object clears them. Null is ignored.
	Tags map[string]string `json:"tags,omitempty"`
	// Replaces the Team scope. An empty string clears it. Null is ignored.
	TeamId datadog.NullableString `json:"team_id,omitempty"`
	// Replaces the user scope. An empty string clears it. Null is ignored.
	UserId datadog.NullableString `json:"user_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayRateLimitUpdateAttributes instantiates a new AiGatewayRateLimitUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayRateLimitUpdateAttributes() *AiGatewayRateLimitUpdateAttributes {
	this := AiGatewayRateLimitUpdateAttributes{}
	return &this
}

// NewAiGatewayRateLimitUpdateAttributesWithDefaults instantiates a new AiGatewayRateLimitUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayRateLimitUpdateAttributesWithDefaults() *AiGatewayRateLimitUpdateAttributes {
	this := AiGatewayRateLimitUpdateAttributes{}
	return &this
}

// GetCostMicrocentsPerMinute returns the CostMicrocentsPerMinute field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiGatewayRateLimitUpdateAttributes) GetCostMicrocentsPerMinute() int64 {
	if o == nil || o.CostMicrocentsPerMinute.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CostMicrocentsPerMinute.Get()
}

// GetCostMicrocentsPerMinuteOk returns a tuple with the CostMicrocentsPerMinute field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AiGatewayRateLimitUpdateAttributes) GetCostMicrocentsPerMinuteOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CostMicrocentsPerMinute.Get(), o.CostMicrocentsPerMinute.IsSet()
}

// HasCostMicrocentsPerMinute returns a boolean if a field has been set.
func (o *AiGatewayRateLimitUpdateAttributes) HasCostMicrocentsPerMinute() bool {
	return o != nil && o.CostMicrocentsPerMinute.IsSet()
}

// SetCostMicrocentsPerMinute gets a reference to the given datadog.NullableInt64 and assigns it to the CostMicrocentsPerMinute field.
func (o *AiGatewayRateLimitUpdateAttributes) SetCostMicrocentsPerMinute(v int64) {
	o.CostMicrocentsPerMinute.Set(&v)
}

// SetCostMicrocentsPerMinuteNil sets the value for CostMicrocentsPerMinute to be an explicit nil.
func (o *AiGatewayRateLimitUpdateAttributes) SetCostMicrocentsPerMinuteNil() {
	o.CostMicrocentsPerMinute.Set(nil)
}

// UnsetCostMicrocentsPerMinute ensures that no value is present for CostMicrocentsPerMinute, not even an explicit nil.
func (o *AiGatewayRateLimitUpdateAttributes) UnsetCostMicrocentsPerMinute() {
	o.CostMicrocentsPerMinute.Unset()
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AiGatewayRateLimitUpdateAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitUpdateAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AiGatewayRateLimitUpdateAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AiGatewayRateLimitUpdateAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetInputTokensPerMinute returns the InputTokensPerMinute field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiGatewayRateLimitUpdateAttributes) GetInputTokensPerMinute() int64 {
	if o == nil || o.InputTokensPerMinute.Get() == nil {
		var ret int64
		return ret
	}
	return *o.InputTokensPerMinute.Get()
}

// GetInputTokensPerMinuteOk returns a tuple with the InputTokensPerMinute field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AiGatewayRateLimitUpdateAttributes) GetInputTokensPerMinuteOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.InputTokensPerMinute.Get(), o.InputTokensPerMinute.IsSet()
}

// HasInputTokensPerMinute returns a boolean if a field has been set.
func (o *AiGatewayRateLimitUpdateAttributes) HasInputTokensPerMinute() bool {
	return o != nil && o.InputTokensPerMinute.IsSet()
}

// SetInputTokensPerMinute gets a reference to the given datadog.NullableInt64 and assigns it to the InputTokensPerMinute field.
func (o *AiGatewayRateLimitUpdateAttributes) SetInputTokensPerMinute(v int64) {
	o.InputTokensPerMinute.Set(&v)
}

// SetInputTokensPerMinuteNil sets the value for InputTokensPerMinute to be an explicit nil.
func (o *AiGatewayRateLimitUpdateAttributes) SetInputTokensPerMinuteNil() {
	o.InputTokensPerMinute.Set(nil)
}

// UnsetInputTokensPerMinute ensures that no value is present for InputTokensPerMinute, not even an explicit nil.
func (o *AiGatewayRateLimitUpdateAttributes) UnsetInputTokensPerMinute() {
	o.InputTokensPerMinute.Unset()
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *AiGatewayRateLimitUpdateAttributes) GetMode() AiGatewayRateLimitMode {
	if o == nil || o.Mode == nil {
		var ret AiGatewayRateLimitMode
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitUpdateAttributes) GetModeOk() (*AiGatewayRateLimitMode, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *AiGatewayRateLimitUpdateAttributes) HasMode() bool {
	return o != nil && o.Mode != nil
}

// SetMode gets a reference to the given AiGatewayRateLimitMode and assigns it to the Mode field.
func (o *AiGatewayRateLimitUpdateAttributes) SetMode(v AiGatewayRateLimitMode) {
	o.Mode = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AiGatewayRateLimitUpdateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AiGatewayRateLimitUpdateAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AiGatewayRateLimitUpdateAttributes) SetName(v string) {
	o.Name = &v
}

// GetOutputTokensPerMinute returns the OutputTokensPerMinute field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiGatewayRateLimitUpdateAttributes) GetOutputTokensPerMinute() int64 {
	if o == nil || o.OutputTokensPerMinute.Get() == nil {
		var ret int64
		return ret
	}
	return *o.OutputTokensPerMinute.Get()
}

// GetOutputTokensPerMinuteOk returns a tuple with the OutputTokensPerMinute field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AiGatewayRateLimitUpdateAttributes) GetOutputTokensPerMinuteOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.OutputTokensPerMinute.Get(), o.OutputTokensPerMinute.IsSet()
}

// HasOutputTokensPerMinute returns a boolean if a field has been set.
func (o *AiGatewayRateLimitUpdateAttributes) HasOutputTokensPerMinute() bool {
	return o != nil && o.OutputTokensPerMinute.IsSet()
}

// SetOutputTokensPerMinute gets a reference to the given datadog.NullableInt64 and assigns it to the OutputTokensPerMinute field.
func (o *AiGatewayRateLimitUpdateAttributes) SetOutputTokensPerMinute(v int64) {
	o.OutputTokensPerMinute.Set(&v)
}

// SetOutputTokensPerMinuteNil sets the value for OutputTokensPerMinute to be an explicit nil.
func (o *AiGatewayRateLimitUpdateAttributes) SetOutputTokensPerMinuteNil() {
	o.OutputTokensPerMinute.Set(nil)
}

// UnsetOutputTokensPerMinute ensures that no value is present for OutputTokensPerMinute, not even an explicit nil.
func (o *AiGatewayRateLimitUpdateAttributes) UnsetOutputTokensPerMinute() {
	o.OutputTokensPerMinute.Unset()
}

// GetRequestsPerMinute returns the RequestsPerMinute field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiGatewayRateLimitUpdateAttributes) GetRequestsPerMinute() int64 {
	if o == nil || o.RequestsPerMinute.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RequestsPerMinute.Get()
}

// GetRequestsPerMinuteOk returns a tuple with the RequestsPerMinute field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AiGatewayRateLimitUpdateAttributes) GetRequestsPerMinuteOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestsPerMinute.Get(), o.RequestsPerMinute.IsSet()
}

// HasRequestsPerMinute returns a boolean if a field has been set.
func (o *AiGatewayRateLimitUpdateAttributes) HasRequestsPerMinute() bool {
	return o != nil && o.RequestsPerMinute.IsSet()
}

// SetRequestsPerMinute gets a reference to the given datadog.NullableInt64 and assigns it to the RequestsPerMinute field.
func (o *AiGatewayRateLimitUpdateAttributes) SetRequestsPerMinute(v int64) {
	o.RequestsPerMinute.Set(&v)
}

// SetRequestsPerMinuteNil sets the value for RequestsPerMinute to be an explicit nil.
func (o *AiGatewayRateLimitUpdateAttributes) SetRequestsPerMinuteNil() {
	o.RequestsPerMinute.Set(nil)
}

// UnsetRequestsPerMinute ensures that no value is present for RequestsPerMinute, not even an explicit nil.
func (o *AiGatewayRateLimitUpdateAttributes) UnsetRequestsPerMinute() {
	o.RequestsPerMinute.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiGatewayRateLimitUpdateAttributes) GetTags() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AiGatewayRateLimitUpdateAttributes) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AiGatewayRateLimitUpdateAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *AiGatewayRateLimitUpdateAttributes) SetTags(v map[string]string) {
	o.Tags = v
}

// GetTeamId returns the TeamId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiGatewayRateLimitUpdateAttributes) GetTeamId() string {
	if o == nil || o.TeamId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TeamId.Get()
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AiGatewayRateLimitUpdateAttributes) GetTeamIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TeamId.Get(), o.TeamId.IsSet()
}

// HasTeamId returns a boolean if a field has been set.
func (o *AiGatewayRateLimitUpdateAttributes) HasTeamId() bool {
	return o != nil && o.TeamId.IsSet()
}

// SetTeamId gets a reference to the given datadog.NullableString and assigns it to the TeamId field.
func (o *AiGatewayRateLimitUpdateAttributes) SetTeamId(v string) {
	o.TeamId.Set(&v)
}

// SetTeamIdNil sets the value for TeamId to be an explicit nil.
func (o *AiGatewayRateLimitUpdateAttributes) SetTeamIdNil() {
	o.TeamId.Set(nil)
}

// UnsetTeamId ensures that no value is present for TeamId, not even an explicit nil.
func (o *AiGatewayRateLimitUpdateAttributes) UnsetTeamId() {
	o.TeamId.Unset()
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiGatewayRateLimitUpdateAttributes) GetUserId() string {
	if o == nil || o.UserId.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AiGatewayRateLimitUpdateAttributes) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *AiGatewayRateLimitUpdateAttributes) HasUserId() bool {
	return o != nil && o.UserId.IsSet()
}

// SetUserId gets a reference to the given datadog.NullableString and assigns it to the UserId field.
func (o *AiGatewayRateLimitUpdateAttributes) SetUserId(v string) {
	o.UserId.Set(&v)
}

// SetUserIdNil sets the value for UserId to be an explicit nil.
func (o *AiGatewayRateLimitUpdateAttributes) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil.
func (o *AiGatewayRateLimitUpdateAttributes) UnsetUserId() {
	o.UserId.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayRateLimitUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CostMicrocentsPerMinute.IsSet() {
		toSerialize["cost_microcents_per_minute"] = o.CostMicrocentsPerMinute.Get()
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.InputTokensPerMinute.IsSet() {
		toSerialize["input_tokens_per_minute"] = o.InputTokensPerMinute.Get()
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OutputTokensPerMinute.IsSet() {
		toSerialize["output_tokens_per_minute"] = o.OutputTokensPerMinute.Get()
	}
	if o.RequestsPerMinute.IsSet() {
		toSerialize["requests_per_minute"] = o.RequestsPerMinute.Get()
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.TeamId.IsSet() {
		toSerialize["team_id"] = o.TeamId.Get()
	}
	if o.UserId.IsSet() {
		toSerialize["user_id"] = o.UserId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayRateLimitUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CostMicrocentsPerMinute datadog.NullableInt64   `json:"cost_microcents_per_minute,omitempty"`
		Enabled                 *bool                   `json:"enabled,omitempty"`
		InputTokensPerMinute    datadog.NullableInt64   `json:"input_tokens_per_minute,omitempty"`
		Mode                    *AiGatewayRateLimitMode `json:"mode,omitempty"`
		Name                    *string                 `json:"name,omitempty"`
		OutputTokensPerMinute   datadog.NullableInt64   `json:"output_tokens_per_minute,omitempty"`
		RequestsPerMinute       datadog.NullableInt64   `json:"requests_per_minute,omitempty"`
		Tags                    map[string]string       `json:"tags,omitempty"`
		TeamId                  datadog.NullableString  `json:"team_id,omitempty"`
		UserId                  datadog.NullableString  `json:"user_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cost_microcents_per_minute", "enabled", "input_tokens_per_minute", "mode", "name", "output_tokens_per_minute", "requests_per_minute", "tags", "team_id", "user_id"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CostMicrocentsPerMinute = all.CostMicrocentsPerMinute
	o.Enabled = all.Enabled
	o.InputTokensPerMinute = all.InputTokensPerMinute
	if all.Mode != nil && !all.Mode.IsValid() {
		hasInvalidField = true
	} else {
		o.Mode = all.Mode
	}
	o.Name = all.Name
	o.OutputTokensPerMinute = all.OutputTokensPerMinute
	o.RequestsPerMinute = all.RequestsPerMinute
	o.Tags = all.Tags
	o.TeamId = all.TeamId
	o.UserId = all.UserId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
