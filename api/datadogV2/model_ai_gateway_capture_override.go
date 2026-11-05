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

// AiGatewayCaptureOverride Saved data capture override with its request filters, capture mode, revision, and timestamps.
type AiGatewayCaptureOverride struct {
	// Timestamp when the override was created.
	CreatedAt time.Time `json:"created_at"`
	// UUID identifying the override within the organization.
	Id uuid.UUID `json:"id"`
	// Items Metadata-only capture adds, each listed once. An item requires the items whose information it already reveals: `tool_output_size`, `tool_names` and `conversation_shape` require `tool_usage`, and `input_size_by_role` requires `tool_output_size`. A set that leaves out a required item is rejected. When several Metadata-only overrides match a request, it gets only the items all of them select.
	MetadataCapture []AiGatewayMetadataCaptureItem `json:"metadata_capture,omitempty"`
	// Capture mode for matching traffic.
	Mode AiGatewayCaptureOverrideMode `json:"mode"`
	// Display name of the override. Names need not be unique; identify overrides by `id`.
	Name string `json:"name"`
	// Current override revision, required for replacement or deletion.
	Revision int64 `json:"revision"`
	// Visible request filters. Provider and model tags are withheld when the caller cannot view them. An empty object or empty tags can represent a fully restricted scope; scope_restricted identifies redaction.
	Scope AiGatewayCaptureResponseScope `json:"scope"`
	// Whether provider or model scope metadata was withheld. Omit unchanged scope fields on an unrelated update to preserve their saved values.
	ScopeRestricted bool `json:"scope_restricted"`
	// Timestamp when the override was last modified.
	UpdatedAt time.Time `json:"updated_at"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayCaptureOverride instantiates a new AiGatewayCaptureOverride object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayCaptureOverride(createdAt time.Time, id uuid.UUID, mode AiGatewayCaptureOverrideMode, name string, revision int64, scope AiGatewayCaptureResponseScope, scopeRestricted bool, updatedAt time.Time) *AiGatewayCaptureOverride {
	this := AiGatewayCaptureOverride{}
	this.CreatedAt = createdAt
	this.Id = id
	this.Mode = mode
	this.Name = name
	this.Revision = revision
	this.Scope = scope
	this.ScopeRestricted = scopeRestricted
	this.UpdatedAt = updatedAt
	return &this
}

// NewAiGatewayCaptureOverrideWithDefaults instantiates a new AiGatewayCaptureOverride object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayCaptureOverrideWithDefaults() *AiGatewayCaptureOverride {
	this := AiGatewayCaptureOverride{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *AiGatewayCaptureOverride) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AiGatewayCaptureOverride) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *AiGatewayCaptureOverride) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetId returns the Id field value.
func (o *AiGatewayCaptureOverride) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AiGatewayCaptureOverride) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *AiGatewayCaptureOverride) SetId(v uuid.UUID) {
	o.Id = v
}

// GetMetadataCapture returns the MetadataCapture field value if set, zero value otherwise.
func (o *AiGatewayCaptureOverride) GetMetadataCapture() []AiGatewayMetadataCaptureItem {
	if o == nil || o.MetadataCapture == nil {
		var ret []AiGatewayMetadataCaptureItem
		return ret
	}
	return o.MetadataCapture
}

// GetMetadataCaptureOk returns a tuple with the MetadataCapture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayCaptureOverride) GetMetadataCaptureOk() (*[]AiGatewayMetadataCaptureItem, bool) {
	if o == nil || o.MetadataCapture == nil {
		return nil, false
	}
	return &o.MetadataCapture, true
}

// HasMetadataCapture returns a boolean if a field has been set.
func (o *AiGatewayCaptureOverride) HasMetadataCapture() bool {
	return o != nil && o.MetadataCapture != nil
}

// SetMetadataCapture gets a reference to the given []AiGatewayMetadataCaptureItem and assigns it to the MetadataCapture field.
func (o *AiGatewayCaptureOverride) SetMetadataCapture(v []AiGatewayMetadataCaptureItem) {
	o.MetadataCapture = v
}

// GetMode returns the Mode field value.
func (o *AiGatewayCaptureOverride) GetMode() AiGatewayCaptureOverrideMode {
	if o == nil {
		var ret AiGatewayCaptureOverrideMode
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *AiGatewayCaptureOverride) GetModeOk() (*AiGatewayCaptureOverrideMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *AiGatewayCaptureOverride) SetMode(v AiGatewayCaptureOverrideMode) {
	o.Mode = v
}

// GetName returns the Name field value.
func (o *AiGatewayCaptureOverride) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiGatewayCaptureOverride) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *AiGatewayCaptureOverride) SetName(v string) {
	o.Name = v
}

// GetRevision returns the Revision field value.
func (o *AiGatewayCaptureOverride) GetRevision() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value
// and a boolean to check if the value has been set.
func (o *AiGatewayCaptureOverride) GetRevisionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revision, true
}

// SetRevision sets field value.
func (o *AiGatewayCaptureOverride) SetRevision(v int64) {
	o.Revision = v
}

// GetScope returns the Scope field value.
func (o *AiGatewayCaptureOverride) GetScope() AiGatewayCaptureResponseScope {
	if o == nil {
		var ret AiGatewayCaptureResponseScope
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *AiGatewayCaptureOverride) GetScopeOk() (*AiGatewayCaptureResponseScope, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value.
func (o *AiGatewayCaptureOverride) SetScope(v AiGatewayCaptureResponseScope) {
	o.Scope = v
}

// GetScopeRestricted returns the ScopeRestricted field value.
func (o *AiGatewayCaptureOverride) GetScopeRestricted() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.ScopeRestricted
}

// GetScopeRestrictedOk returns a tuple with the ScopeRestricted field value
// and a boolean to check if the value has been set.
func (o *AiGatewayCaptureOverride) GetScopeRestrictedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScopeRestricted, true
}

// SetScopeRestricted sets field value.
func (o *AiGatewayCaptureOverride) SetScopeRestricted(v bool) {
	o.ScopeRestricted = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *AiGatewayCaptureOverride) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *AiGatewayCaptureOverride) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *AiGatewayCaptureOverride) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayCaptureOverride) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["id"] = o.Id
	if o.MetadataCapture != nil {
		toSerialize["metadata_capture"] = o.MetadataCapture
	}
	toSerialize["mode"] = o.Mode
	toSerialize["name"] = o.Name
	toSerialize["revision"] = o.Revision
	toSerialize["scope"] = o.Scope
	toSerialize["scope_restricted"] = o.ScopeRestricted
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayCaptureOverride) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt       *time.Time                     `json:"created_at"`
		Id              *uuid.UUID                     `json:"id"`
		MetadataCapture []AiGatewayMetadataCaptureItem `json:"metadata_capture,omitempty"`
		Mode            *AiGatewayCaptureOverrideMode  `json:"mode"`
		Name            *string                        `json:"name"`
		Revision        *int64                         `json:"revision"`
		Scope           *AiGatewayCaptureResponseScope `json:"scope"`
		ScopeRestricted *bool                          `json:"scope_restricted"`
		UpdatedAt       *time.Time                     `json:"updated_at"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
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
	if all.Scope == nil {
		return fmt.Errorf("required field scope missing")
	}
	if all.ScopeRestricted == nil {
		return fmt.Errorf("required field scope_restricted missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "id", "metadata_capture", "mode", "name", "revision", "scope", "scope_restricted", "updated_at"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = *all.CreatedAt
	o.Id = *all.Id
	o.MetadataCapture = all.MetadataCapture
	if !all.Mode.IsValid() {
		hasInvalidField = true
	} else {
		o.Mode = *all.Mode
	}
	o.Name = *all.Name
	o.Revision = *all.Revision
	if all.Scope.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Scope = *all.Scope
	o.ScopeRestricted = *all.ScopeRestricted
	o.UpdatedAt = *all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
