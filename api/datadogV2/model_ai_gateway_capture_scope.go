// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayCaptureScope All selected filters combine with AND, with one value per filter. The `model` and `provider` tags apply to each actual direct or routed attempt, including fallback destinations. Routes are not selectable filters. At least one filter is required.
// If any matching override uses `metadata_only`, that mode applies. Otherwise, if any matching override uses `full_capture`, that mode applies. The organization default applies only when no override matches. Overrides have no priority or specificity ordering.
type AiGatewayCaptureScope struct {
	// Exact-match request tags with lowercase keys. The reserved `model` tag is the canonical model ID, and the reserved `provider` tag is the configured connection slug. The server sets both. Custom tags come from `x-dd-tag` headers. Keys and values must not have leading or trailing white space. Values are case-sensitive.
	Tags map[string]string `json:"tags,omitempty"`
	// UUID of a Datadog team. Matches the authorized team selected by the `x-ai-gateway-team` request header. Requests without this header do not match team-scoped overrides. Open-membership teams cannot scope overrides.
	TeamId *uuid.UUID `json:"team_id,omitempty"`
	// Authenticated Datadog user UUID in this organization.
	UserId *uuid.UUID `json:"user_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewAiGatewayCaptureScope instantiates a new AiGatewayCaptureScope object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayCaptureScope() *AiGatewayCaptureScope {
	this := AiGatewayCaptureScope{}
	return &this
}

// NewAiGatewayCaptureScopeWithDefaults instantiates a new AiGatewayCaptureScope object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayCaptureScopeWithDefaults() *AiGatewayCaptureScope {
	this := AiGatewayCaptureScope{}
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *AiGatewayCaptureScope) GetTags() map[string]string {
	if o == nil || o.Tags == nil {
		var ret map[string]string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayCaptureScope) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AiGatewayCaptureScope) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *AiGatewayCaptureScope) SetTags(v map[string]string) {
	o.Tags = v
}

// GetTeamId returns the TeamId field value if set, zero value otherwise.
func (o *AiGatewayCaptureScope) GetTeamId() uuid.UUID {
	if o == nil || o.TeamId == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayCaptureScope) GetTeamIdOk() (*uuid.UUID, bool) {
	if o == nil || o.TeamId == nil {
		return nil, false
	}
	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *AiGatewayCaptureScope) HasTeamId() bool {
	return o != nil && o.TeamId != nil
}

// SetTeamId gets a reference to the given uuid.UUID and assigns it to the TeamId field.
func (o *AiGatewayCaptureScope) SetTeamId(v uuid.UUID) {
	o.TeamId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *AiGatewayCaptureScope) GetUserId() uuid.UUID {
	if o == nil || o.UserId == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayCaptureScope) GetUserIdOk() (*uuid.UUID, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *AiGatewayCaptureScope) HasUserId() bool {
	return o != nil && o.UserId != nil
}

// SetUserId gets a reference to the given uuid.UUID and assigns it to the UserId field.
func (o *AiGatewayCaptureScope) SetUserId(v uuid.UUID) {
	o.UserId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayCaptureScope) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
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
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayCaptureScope) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Tags   map[string]string `json:"tags,omitempty"`
		TeamId *uuid.UUID        `json:"team_id,omitempty"`
		UserId *uuid.UUID        `json:"user_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.Tags = all.Tags
	o.TeamId = all.TeamId
	o.UserId = all.UserId

	return nil
}
