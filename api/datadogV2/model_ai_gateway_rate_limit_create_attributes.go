// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayRateLimitCreateAttributes Attributes used to create an AI Gateway rate limit.
type AiGatewayRateLimitCreateAttributes struct {
	// Maximum input tokens per rolling minute.
	InputTokensPerMinute datadog.NullableInt64 `json:"input_tokens_per_minute,omitempty"`
	// Action taken when request rate exceeds the limit.
	Mode *AiGatewayRateLimitMode `json:"mode,omitempty"`
	// Rate limit name. Must not be empty.
	Name string `json:"name"`
	// Maximum output tokens per rolling minute.
	OutputTokensPerMinute datadog.NullableInt64 `json:"output_tokens_per_minute,omitempty"`
	// Maximum requests per rolling minute. At least one of `requests_per_minute`, `input_tokens_per_minute`, and `output_tokens_per_minute` is required.
	RequestsPerMinute datadog.NullableInt64 `json:"requests_per_minute,omitempty"`
	// Dimension rules for a template rate limit, which gives each matched user, Team, or value its own allowance with the rule's limits. Can't be combined with user_id, team_id, or tags, and can't be changed after creation. Omit for a rule with one shared allowance.
	Selector []AiGatewayRateLimitSelectorEntry `json:"selector,omitempty"`
	// Optional scope tags, such as `provider`, `model`, `ml_app`, or custom dimensions. A request matches when it has all of these tags. The reserved metadata keys `name`, `budget_name`, and `span_filter` are rejected.
	Tags map[string]string `json:"tags,omitempty"`
	// Optional UUID. Count only requests attributed to this Datadog Team.
	TeamId *uuid.UUID `json:"team_id,omitempty"`
	// Optional UUID. Count only this user's requests.
	UserId *uuid.UUID `json:"user_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayRateLimitCreateAttributes instantiates a new AiGatewayRateLimitCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayRateLimitCreateAttributes(name string) *AiGatewayRateLimitCreateAttributes {
	this := AiGatewayRateLimitCreateAttributes{}
	this.Name = name
	return &this
}

// NewAiGatewayRateLimitCreateAttributesWithDefaults instantiates a new AiGatewayRateLimitCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayRateLimitCreateAttributesWithDefaults() *AiGatewayRateLimitCreateAttributes {
	this := AiGatewayRateLimitCreateAttributes{}
	return &this
}

// GetInputTokensPerMinute returns the InputTokensPerMinute field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiGatewayRateLimitCreateAttributes) GetInputTokensPerMinute() int64 {
	if o == nil || o.InputTokensPerMinute.Get() == nil {
		var ret int64
		return ret
	}
	return *o.InputTokensPerMinute.Get()
}

// GetInputTokensPerMinuteOk returns a tuple with the InputTokensPerMinute field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AiGatewayRateLimitCreateAttributes) GetInputTokensPerMinuteOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.InputTokensPerMinute.Get(), o.InputTokensPerMinute.IsSet()
}

// HasInputTokensPerMinute returns a boolean if a field has been set.
func (o *AiGatewayRateLimitCreateAttributes) HasInputTokensPerMinute() bool {
	return o != nil && o.InputTokensPerMinute.IsSet()
}

// SetInputTokensPerMinute gets a reference to the given datadog.NullableInt64 and assigns it to the InputTokensPerMinute field.
func (o *AiGatewayRateLimitCreateAttributes) SetInputTokensPerMinute(v int64) {
	o.InputTokensPerMinute.Set(&v)
}

// SetInputTokensPerMinuteNil sets the value for InputTokensPerMinute to be an explicit nil.
func (o *AiGatewayRateLimitCreateAttributes) SetInputTokensPerMinuteNil() {
	o.InputTokensPerMinute.Set(nil)
}

// UnsetInputTokensPerMinute ensures that no value is present for InputTokensPerMinute, not even an explicit nil.
func (o *AiGatewayRateLimitCreateAttributes) UnsetInputTokensPerMinute() {
	o.InputTokensPerMinute.Unset()
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *AiGatewayRateLimitCreateAttributes) GetMode() AiGatewayRateLimitMode {
	if o == nil || o.Mode == nil {
		var ret AiGatewayRateLimitMode
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitCreateAttributes) GetModeOk() (*AiGatewayRateLimitMode, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *AiGatewayRateLimitCreateAttributes) HasMode() bool {
	return o != nil && o.Mode != nil
}

// SetMode gets a reference to the given AiGatewayRateLimitMode and assigns it to the Mode field.
func (o *AiGatewayRateLimitCreateAttributes) SetMode(v AiGatewayRateLimitMode) {
	o.Mode = &v
}

// GetName returns the Name field value.
func (o *AiGatewayRateLimitCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *AiGatewayRateLimitCreateAttributes) SetName(v string) {
	o.Name = v
}

// GetOutputTokensPerMinute returns the OutputTokensPerMinute field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiGatewayRateLimitCreateAttributes) GetOutputTokensPerMinute() int64 {
	if o == nil || o.OutputTokensPerMinute.Get() == nil {
		var ret int64
		return ret
	}
	return *o.OutputTokensPerMinute.Get()
}

// GetOutputTokensPerMinuteOk returns a tuple with the OutputTokensPerMinute field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AiGatewayRateLimitCreateAttributes) GetOutputTokensPerMinuteOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.OutputTokensPerMinute.Get(), o.OutputTokensPerMinute.IsSet()
}

// HasOutputTokensPerMinute returns a boolean if a field has been set.
func (o *AiGatewayRateLimitCreateAttributes) HasOutputTokensPerMinute() bool {
	return o != nil && o.OutputTokensPerMinute.IsSet()
}

// SetOutputTokensPerMinute gets a reference to the given datadog.NullableInt64 and assigns it to the OutputTokensPerMinute field.
func (o *AiGatewayRateLimitCreateAttributes) SetOutputTokensPerMinute(v int64) {
	o.OutputTokensPerMinute.Set(&v)
}

// SetOutputTokensPerMinuteNil sets the value for OutputTokensPerMinute to be an explicit nil.
func (o *AiGatewayRateLimitCreateAttributes) SetOutputTokensPerMinuteNil() {
	o.OutputTokensPerMinute.Set(nil)
}

// UnsetOutputTokensPerMinute ensures that no value is present for OutputTokensPerMinute, not even an explicit nil.
func (o *AiGatewayRateLimitCreateAttributes) UnsetOutputTokensPerMinute() {
	o.OutputTokensPerMinute.Unset()
}

// GetRequestsPerMinute returns the RequestsPerMinute field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiGatewayRateLimitCreateAttributes) GetRequestsPerMinute() int64 {
	if o == nil || o.RequestsPerMinute.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RequestsPerMinute.Get()
}

// GetRequestsPerMinuteOk returns a tuple with the RequestsPerMinute field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AiGatewayRateLimitCreateAttributes) GetRequestsPerMinuteOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestsPerMinute.Get(), o.RequestsPerMinute.IsSet()
}

// HasRequestsPerMinute returns a boolean if a field has been set.
func (o *AiGatewayRateLimitCreateAttributes) HasRequestsPerMinute() bool {
	return o != nil && o.RequestsPerMinute.IsSet()
}

// SetRequestsPerMinute gets a reference to the given datadog.NullableInt64 and assigns it to the RequestsPerMinute field.
func (o *AiGatewayRateLimitCreateAttributes) SetRequestsPerMinute(v int64) {
	o.RequestsPerMinute.Set(&v)
}

// SetRequestsPerMinuteNil sets the value for RequestsPerMinute to be an explicit nil.
func (o *AiGatewayRateLimitCreateAttributes) SetRequestsPerMinuteNil() {
	o.RequestsPerMinute.Set(nil)
}

// UnsetRequestsPerMinute ensures that no value is present for RequestsPerMinute, not even an explicit nil.
func (o *AiGatewayRateLimitCreateAttributes) UnsetRequestsPerMinute() {
	o.RequestsPerMinute.Unset()
}

// GetSelector returns the Selector field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiGatewayRateLimitCreateAttributes) GetSelector() []AiGatewayRateLimitSelectorEntry {
	if o == nil {
		var ret []AiGatewayRateLimitSelectorEntry
		return ret
	}
	return o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AiGatewayRateLimitCreateAttributes) GetSelectorOk() (*[]AiGatewayRateLimitSelectorEntry, bool) {
	if o == nil || o.Selector == nil {
		return nil, false
	}
	return &o.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *AiGatewayRateLimitCreateAttributes) HasSelector() bool {
	return o != nil && o.Selector != nil
}

// SetSelector gets a reference to the given []AiGatewayRateLimitSelectorEntry and assigns it to the Selector field.
func (o *AiGatewayRateLimitCreateAttributes) SetSelector(v []AiGatewayRateLimitSelectorEntry) {
	o.Selector = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *AiGatewayRateLimitCreateAttributes) GetTags() map[string]string {
	if o == nil || o.Tags == nil {
		var ret map[string]string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitCreateAttributes) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AiGatewayRateLimitCreateAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *AiGatewayRateLimitCreateAttributes) SetTags(v map[string]string) {
	o.Tags = v
}

// GetTeamId returns the TeamId field value if set, zero value otherwise.
func (o *AiGatewayRateLimitCreateAttributes) GetTeamId() uuid.UUID {
	if o == nil || o.TeamId == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitCreateAttributes) GetTeamIdOk() (*uuid.UUID, bool) {
	if o == nil || o.TeamId == nil {
		return nil, false
	}
	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *AiGatewayRateLimitCreateAttributes) HasTeamId() bool {
	return o != nil && o.TeamId != nil
}

// SetTeamId gets a reference to the given uuid.UUID and assigns it to the TeamId field.
func (o *AiGatewayRateLimitCreateAttributes) SetTeamId(v uuid.UUID) {
	o.TeamId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *AiGatewayRateLimitCreateAttributes) GetUserId() uuid.UUID {
	if o == nil || o.UserId == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitCreateAttributes) GetUserIdOk() (*uuid.UUID, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *AiGatewayRateLimitCreateAttributes) HasUserId() bool {
	return o != nil && o.UserId != nil
}

// SetUserId gets a reference to the given uuid.UUID and assigns it to the UserId field.
func (o *AiGatewayRateLimitCreateAttributes) SetUserId(v uuid.UUID) {
	o.UserId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayRateLimitCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.InputTokensPerMinute.IsSet() {
		toSerialize["input_tokens_per_minute"] = o.InputTokensPerMinute.Get()
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	toSerialize["name"] = o.Name
	if o.OutputTokensPerMinute.IsSet() {
		toSerialize["output_tokens_per_minute"] = o.OutputTokensPerMinute.Get()
	}
	if o.RequestsPerMinute.IsSet() {
		toSerialize["requests_per_minute"] = o.RequestsPerMinute.Get()
	}
	if o.Selector != nil {
		toSerialize["selector"] = o.Selector
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.TeamId != nil {
		toSerialize["team_id"] = o.TeamId
	}
	if o.UserId != nil {
		toSerialize["user_id"] = o.UserId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayRateLimitCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		InputTokensPerMinute  datadog.NullableInt64             `json:"input_tokens_per_minute,omitempty"`
		Mode                  *AiGatewayRateLimitMode           `json:"mode,omitempty"`
		Name                  *string                           `json:"name"`
		OutputTokensPerMinute datadog.NullableInt64             `json:"output_tokens_per_minute,omitempty"`
		RequestsPerMinute     datadog.NullableInt64             `json:"requests_per_minute,omitempty"`
		Selector              []AiGatewayRateLimitSelectorEntry `json:"selector,omitempty"`
		Tags                  map[string]string                 `json:"tags,omitempty"`
		TeamId                *uuid.UUID                        `json:"team_id,omitempty"`
		UserId                *uuid.UUID                        `json:"user_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"input_tokens_per_minute", "mode", "name", "output_tokens_per_minute", "requests_per_minute", "selector", "tags", "team_id", "user_id"})
	} else {
		return err
	}

	hasInvalidField := false
	o.InputTokensPerMinute = all.InputTokensPerMinute
	if all.Mode != nil && !all.Mode.IsValid() {
		hasInvalidField = true
	} else {
		o.Mode = all.Mode
	}
	o.Name = *all.Name
	o.OutputTokensPerMinute = all.OutputTokensPerMinute
	o.RequestsPerMinute = all.RequestsPerMinute
	o.Selector = all.Selector
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
