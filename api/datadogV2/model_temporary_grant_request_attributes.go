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

// TemporaryGrantRequestAttributes Request history and the linked grant state. List and detail reads include grant_state for approved requests. Decision responses record the decision. Read the request for current grant state.
type TemporaryGrantRequestAttributes struct {
	// When the request was submitted.
	CreatedAt time.Time `json:"created_at"`
	// When the decision or cancellation was recorded.
	DecidedAt *time.Time `json:"decided_at,omitempty"`
	// UUID of the reviewer, or of the requester if the request was canceled.
	DecidedByUserId *uuid.UUID `json:"decided_by_user_id,omitempty"`
	// The requester identity resolved by the server.
	EntityKey string `json:"entity_key"`
	// Immutable exclusive expiry, normalized to UTC microsecond precision.
	ExpiresAt time.Time `json:"expires_at"`
	// Current state of the grant associated with an approved request. Replaced means a newer grant superseded this grant. The newer grant might no longer be active. This field is absent for requests that weren't approved.
	GrantState *TemporaryGrantRequestGrantState `json:"grant_state,omitempty"`
	// Reason for requesting temporary credits. Trimmed and limited to 4096 UTF-8 bytes.
	Justification string `json:"justification"`
	// The quota namespace. Only AI credits are supported.
	QuotaNamespace string `json:"quota_namespace"`
	// UUID of the authenticated requester.
	RequestedByUserId uuid.UUID `json:"requested_by_user_id"`
	// The durable request decision, independent of whether its grant is still active.
	Status TemporaryGrantRequestStatus `json:"status"`
	// Target monthly AI credit limit in positive whole credits, not additional credits. Fractional credits are not accepted.
	Total int64 `json:"total"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTemporaryGrantRequestAttributes instantiates a new TemporaryGrantRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTemporaryGrantRequestAttributes(createdAt time.Time, entityKey string, expiresAt time.Time, justification string, quotaNamespace string, requestedByUserId uuid.UUID, status TemporaryGrantRequestStatus, total int64) *TemporaryGrantRequestAttributes {
	this := TemporaryGrantRequestAttributes{}
	this.CreatedAt = createdAt
	this.EntityKey = entityKey
	this.ExpiresAt = expiresAt
	this.Justification = justification
	this.QuotaNamespace = quotaNamespace
	this.RequestedByUserId = requestedByUserId
	this.Status = status
	this.Total = total
	return &this
}

// NewTemporaryGrantRequestAttributesWithDefaults instantiates a new TemporaryGrantRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTemporaryGrantRequestAttributesWithDefaults() *TemporaryGrantRequestAttributes {
	this := TemporaryGrantRequestAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *TemporaryGrantRequestAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TemporaryGrantRequestAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *TemporaryGrantRequestAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDecidedAt returns the DecidedAt field value if set, zero value otherwise.
func (o *TemporaryGrantRequestAttributes) GetDecidedAt() time.Time {
	if o == nil || o.DecidedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.DecidedAt
}

// GetDecidedAtOk returns a tuple with the DecidedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemporaryGrantRequestAttributes) GetDecidedAtOk() (*time.Time, bool) {
	if o == nil || o.DecidedAt == nil {
		return nil, false
	}
	return o.DecidedAt, true
}

// HasDecidedAt returns a boolean if a field has been set.
func (o *TemporaryGrantRequestAttributes) HasDecidedAt() bool {
	return o != nil && o.DecidedAt != nil
}

// SetDecidedAt gets a reference to the given time.Time and assigns it to the DecidedAt field.
func (o *TemporaryGrantRequestAttributes) SetDecidedAt(v time.Time) {
	o.DecidedAt = &v
}

// GetDecidedByUserId returns the DecidedByUserId field value if set, zero value otherwise.
func (o *TemporaryGrantRequestAttributes) GetDecidedByUserId() uuid.UUID {
	if o == nil || o.DecidedByUserId == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.DecidedByUserId
}

// GetDecidedByUserIdOk returns a tuple with the DecidedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemporaryGrantRequestAttributes) GetDecidedByUserIdOk() (*uuid.UUID, bool) {
	if o == nil || o.DecidedByUserId == nil {
		return nil, false
	}
	return o.DecidedByUserId, true
}

// HasDecidedByUserId returns a boolean if a field has been set.
func (o *TemporaryGrantRequestAttributes) HasDecidedByUserId() bool {
	return o != nil && o.DecidedByUserId != nil
}

// SetDecidedByUserId gets a reference to the given uuid.UUID and assigns it to the DecidedByUserId field.
func (o *TemporaryGrantRequestAttributes) SetDecidedByUserId(v uuid.UUID) {
	o.DecidedByUserId = &v
}

// GetEntityKey returns the EntityKey field value.
func (o *TemporaryGrantRequestAttributes) GetEntityKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EntityKey
}

// GetEntityKeyOk returns a tuple with the EntityKey field value
// and a boolean to check if the value has been set.
func (o *TemporaryGrantRequestAttributes) GetEntityKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityKey, true
}

// SetEntityKey sets field value.
func (o *TemporaryGrantRequestAttributes) SetEntityKey(v string) {
	o.EntityKey = v
}

// GetExpiresAt returns the ExpiresAt field value.
func (o *TemporaryGrantRequestAttributes) GetExpiresAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *TemporaryGrantRequestAttributes) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value.
func (o *TemporaryGrantRequestAttributes) SetExpiresAt(v time.Time) {
	o.ExpiresAt = v
}

// GetGrantState returns the GrantState field value if set, zero value otherwise.
func (o *TemporaryGrantRequestAttributes) GetGrantState() TemporaryGrantRequestGrantState {
	if o == nil || o.GrantState == nil {
		var ret TemporaryGrantRequestGrantState
		return ret
	}
	return *o.GrantState
}

// GetGrantStateOk returns a tuple with the GrantState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemporaryGrantRequestAttributes) GetGrantStateOk() (*TemporaryGrantRequestGrantState, bool) {
	if o == nil || o.GrantState == nil {
		return nil, false
	}
	return o.GrantState, true
}

// HasGrantState returns a boolean if a field has been set.
func (o *TemporaryGrantRequestAttributes) HasGrantState() bool {
	return o != nil && o.GrantState != nil
}

// SetGrantState gets a reference to the given TemporaryGrantRequestGrantState and assigns it to the GrantState field.
func (o *TemporaryGrantRequestAttributes) SetGrantState(v TemporaryGrantRequestGrantState) {
	o.GrantState = &v
}

// GetJustification returns the Justification field value.
func (o *TemporaryGrantRequestAttributes) GetJustification() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Justification
}

// GetJustificationOk returns a tuple with the Justification field value
// and a boolean to check if the value has been set.
func (o *TemporaryGrantRequestAttributes) GetJustificationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Justification, true
}

// SetJustification sets field value.
func (o *TemporaryGrantRequestAttributes) SetJustification(v string) {
	o.Justification = v
}

// GetQuotaNamespace returns the QuotaNamespace field value.
func (o *TemporaryGrantRequestAttributes) GetQuotaNamespace() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.QuotaNamespace
}

// GetQuotaNamespaceOk returns a tuple with the QuotaNamespace field value
// and a boolean to check if the value has been set.
func (o *TemporaryGrantRequestAttributes) GetQuotaNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QuotaNamespace, true
}

// SetQuotaNamespace sets field value.
func (o *TemporaryGrantRequestAttributes) SetQuotaNamespace(v string) {
	o.QuotaNamespace = v
}

// GetRequestedByUserId returns the RequestedByUserId field value.
func (o *TemporaryGrantRequestAttributes) GetRequestedByUserId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.RequestedByUserId
}

// GetRequestedByUserIdOk returns a tuple with the RequestedByUserId field value
// and a boolean to check if the value has been set.
func (o *TemporaryGrantRequestAttributes) GetRequestedByUserIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestedByUserId, true
}

// SetRequestedByUserId sets field value.
func (o *TemporaryGrantRequestAttributes) SetRequestedByUserId(v uuid.UUID) {
	o.RequestedByUserId = v
}

// GetStatus returns the Status field value.
func (o *TemporaryGrantRequestAttributes) GetStatus() TemporaryGrantRequestStatus {
	if o == nil {
		var ret TemporaryGrantRequestStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TemporaryGrantRequestAttributes) GetStatusOk() (*TemporaryGrantRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *TemporaryGrantRequestAttributes) SetStatus(v TemporaryGrantRequestStatus) {
	o.Status = v
}

// GetTotal returns the Total field value.
func (o *TemporaryGrantRequestAttributes) GetTotal() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *TemporaryGrantRequestAttributes) GetTotalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value.
func (o *TemporaryGrantRequestAttributes) SetTotal(v int64) {
	o.Total = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TemporaryGrantRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.DecidedAt != nil {
		if o.DecidedAt.Nanosecond() == 0 {
			toSerialize["decided_at"] = o.DecidedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["decided_at"] = o.DecidedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.DecidedByUserId != nil {
		toSerialize["decided_by_user_id"] = o.DecidedByUserId
	}
	toSerialize["entity_key"] = o.EntityKey
	if o.ExpiresAt.Nanosecond() == 0 {
		toSerialize["expires_at"] = o.ExpiresAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["expires_at"] = o.ExpiresAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.GrantState != nil {
		toSerialize["grant_state"] = o.GrantState
	}
	toSerialize["justification"] = o.Justification
	toSerialize["quota_namespace"] = o.QuotaNamespace
	toSerialize["requested_by_user_id"] = o.RequestedByUserId
	toSerialize["status"] = o.Status
	toSerialize["total"] = o.Total

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TemporaryGrantRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt         *time.Time                       `json:"created_at"`
		DecidedAt         *time.Time                       `json:"decided_at,omitempty"`
		DecidedByUserId   *uuid.UUID                       `json:"decided_by_user_id,omitempty"`
		EntityKey         *string                          `json:"entity_key"`
		ExpiresAt         *time.Time                       `json:"expires_at"`
		GrantState        *TemporaryGrantRequestGrantState `json:"grant_state,omitempty"`
		Justification     *string                          `json:"justification"`
		QuotaNamespace    *string                          `json:"quota_namespace"`
		RequestedByUserId *uuid.UUID                       `json:"requested_by_user_id"`
		Status            *TemporaryGrantRequestStatus     `json:"status"`
		Total             *int64                           `json:"total"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.EntityKey == nil {
		return fmt.Errorf("required field entity_key missing")
	}
	if all.ExpiresAt == nil {
		return fmt.Errorf("required field expires_at missing")
	}
	if all.Justification == nil {
		return fmt.Errorf("required field justification missing")
	}
	if all.QuotaNamespace == nil {
		return fmt.Errorf("required field quota_namespace missing")
	}
	if all.RequestedByUserId == nil {
		return fmt.Errorf("required field requested_by_user_id missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.Total == nil {
		return fmt.Errorf("required field total missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "decided_at", "decided_by_user_id", "entity_key", "expires_at", "grant_state", "justification", "quota_namespace", "requested_by_user_id", "status", "total"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = *all.CreatedAt
	o.DecidedAt = all.DecidedAt
	o.DecidedByUserId = all.DecidedByUserId
	o.EntityKey = *all.EntityKey
	o.ExpiresAt = *all.ExpiresAt
	if all.GrantState != nil && !all.GrantState.IsValid() {
		hasInvalidField = true
	} else {
		o.GrantState = all.GrantState
	}
	o.Justification = *all.Justification
	o.QuotaNamespace = *all.QuotaNamespace
	o.RequestedByUserId = *all.RequestedByUserId
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.Total = *all.Total

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
